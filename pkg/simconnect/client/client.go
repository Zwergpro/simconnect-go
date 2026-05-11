//go:build windows

// Package client provides the core SimConnect session, message dispatch,
// and ID allocation shared by all domain packages.
package client

import (
	"context"
	"sync"
	"time"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// Sim is the core SimConnect session. Domain packages (ai, camera, ets.) wrap
// a *Sim to add API surface; all share the same ID allocators and dispatch loop.
type Sim struct {
	raw *bindings.SimConnect
	cfg clientConfig

	ctx    context.Context
	cancel context.CancelFunc
	done   chan struct{}

	closeOnce sync.Once
	closeErr  error

	errs chan error

	reqIDs   *idAllocator
	defIDs   *idAllocator
	eventIDs *idAllocator

	mu         sync.Mutex
	closed     bool
	waiters    map[uint32]chan core.RequestResult
	sendToReq  map[uint32]uint32
	dataSubs   map[uint32][]func(core.Message)
	handlers   map[core.RecvID][]func(core.Message)
	closeHooks []func()

	// Lazily-initialized facet clients, guarded by facetMu.
	facetMu    sync.Mutex
	facetCache facets
}

func Dial(ctx context.Context, appName string, opts ...Option) (*Sim, error) {
	cfg := defaultClientConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	if err := bindings.LoadDLL(); err != nil {
		return nil, err
	}

	raw, err := bindings.Open(appName, bindings.HWND(cfg.hwnd), cfg.eventID, bindings.HANDLE(cfg.eventHandle), cfg.configIndex)
	if err != nil {
		return nil, err
	}

	runCtx, cancel := context.WithCancel(ctx)
	sim := &Sim{
		raw:       raw,
		cfg:       cfg,
		ctx:       runCtx,
		cancel:    cancel,
		done:      make(chan struct{}),
		errs:      make(chan error, cfg.channelBuffer),
		reqIDs:    newIDAllocator(1),
		defIDs:    newIDAllocator(1),
		eventIDs:  newIDAllocator(1),
		waiters:   make(map[uint32]chan core.RequestResult),
		sendToReq: make(map[uint32]uint32),
		dataSubs:  make(map[uint32][]func(core.Message)),
		handlers:  make(map[core.RecvID][]func(core.Message)),
	}

	if cfg.manualDispatch {
		close(sim.done)
		return sim, nil
	}

	go sim.run()
	return sim, nil
}

// Open connects to SimConnect and starts the session dispatch loop.
func Open(ctx context.Context, appName string, opts ...Option) (*Sim, error) {
	return Dial(ctx, appName, opts...)
}

func (s *Sim) Errors() <-chan error     { return s.errs }
func (s *Sim) Context() context.Context { return s.ctx }
func (s *Sim) ChannelBuffer() int       { return s.cfg.channelBuffer }

// Bindings exposes the low-level binding session for facet packages.
func (s *Sim) Bindings() *bindings.SimConnect { return s.raw }

func (s *Sim) NextRequestID() uint32    { return s.reqIDs.Next() }
func (s *Sim) NextDefinitionID() uint32 { return s.defIDs.Next() }
func (s *Sim) NextEventID() uint32      { return s.eventIDs.Next() }

// RegisterHandler registers fn to receive messages of the given recv ID.
// Handlers run in the dispatch goroutine and must not block.
func (s *Sim) RegisterHandler(id core.RecvID, fn func(core.Message)) {
	s.mu.Lock()
	s.handlers[id] = append(s.handlers[id], fn)
	s.mu.Unlock()
}

// RegisterCloseHook registers a cleanup function called during Close,
// after the context is cancelled and the dispatch loop has stopped.
func (s *Sim) RegisterCloseHook(fn func()) {
	s.mu.Lock()
	s.closeHooks = append(s.closeHooks, fn)
	s.mu.Unlock()
}

// AddWaiter registers a one-shot result channel for the given request ID.
func (s *Sim) AddWaiter(reqID uint32) (<-chan core.RequestResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return nil, core.ErrClosed
	}
	ch := make(chan core.RequestResult, 1)
	s.waiters[reqID] = ch
	return ch, nil
}

// RemoveWaiter removes and closes the waiter for reqID (used on cancellation).
func (s *Sim) RemoveWaiter(reqID uint32) {
	s.mu.Lock()
	if waiter, ok := s.waiters[reqID]; ok {
		delete(s.waiters, reqID)
		close(waiter)
	}
	s.mu.Unlock()
}

// AddDataSub registers a continuous data handler keyed by requestID.
func (s *Sim) AddDataSub(reqID uint32, fn func(core.Message)) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return core.ErrClosed
	}
	s.dataSubs[reqID] = append(s.dataSubs[reqID], fn)
	return nil
}

// RemoveDataSub removes all data handlers for reqID.
func (s *Sim) RemoveDataSub(reqID uint32) {
	s.mu.Lock()
	delete(s.dataSubs, reqID)
	s.mu.Unlock()
}

// TrackSend maps the last sent packet ID to reqID for exception correlation.
func (s *Sim) TrackSend(reqID uint32) {
	sendID, err := s.raw.GetLastSentPacketID()
	if err != nil {
		return
	}
	s.mu.Lock()
	s.sendToReq[sendID] = reqID
	s.mu.Unlock()
}

func (s *Sim) ReportError(err error) {
	select {
	case s.errs <- err:
	case <-s.ctx.Done():
	}
}

func (s *Sim) Close() error {
	s.closeOnce.Do(func() {
		s.cancel()
		if !s.cfg.manualDispatch {
			<-s.done
		}

		s.mu.Lock()
		s.closed = true
		hooks := append([]func(){}, s.closeHooks...)
		for id, waiter := range s.waiters {
			waiter <- core.RequestResult{Err: core.ErrClosed}
			close(waiter)
			delete(s.waiters, id)
		}
		s.dataSubs = map[uint32][]func(core.Message){}
		s.mu.Unlock()

		for _, fn := range hooks {
			fn()
		}

		s.closeErr = s.raw.Close()
		close(s.errs)
	})
	return s.closeErr
}

func (s *Sim) run() {
	defer close(s.done)
	ticker := time.NewTicker(s.cfg.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			if err := s.Poll(); err != nil {
				s.ReportError(err)
			}
		}
	}
}

func (s *Sim) Poll() error {
	for {
		raw, size, err := s.raw.GetNextDispatch()
		if err != nil {
			return err
		}
		if raw == nil {
			return nil
		}
		msg, err := decodeMessage(raw, size)
		if err != nil {
			return err
		}
		s.dispatch(msg)
	}
}

// DispatchProc is the callback type for CallDispatch.
type DispatchProc func(msg core.Message, context uintptr)

// GetNextMessage retrieves the next pending SimConnect message without blocking.
// Returns (nil, false, nil) when the queue is empty.
func (s *Sim) GetNextMessage() (core.Message, bool, error) {
	raw, size, err := s.raw.GetNextDispatch()
	if err != nil || raw == nil {
		return nil, false, err
	}
	msg, err := decodeMessage(raw, size)
	if err != nil {
		return nil, false, err
	}
	return msg, true, nil
}

// CallDispatch drains the message queue invoking fn for each decoded message.
func (s *Sim) CallDispatch(fn DispatchProc, context uintptr) error {
	return s.raw.CallDispatch(func(raw *bindings.SIMCONNECT_RECV, size uint32, ctx uintptr) {
		msg, err := decodeMessage(raw, size)
		if err != nil {
			s.ReportError(err)
			return
		}
		fn(msg, ctx)
	}, context)
}

// GetLastSentPacketID returns the packet ID of the most recently sent call.
func (s *Sim) GetLastSentPacketID() (uint32, error) {
	return s.raw.GetLastSentPacketID()
}

// RequestResponseTimes returns response-time measurements for the last count packets.
func (s *Sim) RequestResponseTimes(count uint32) ([]float32, error) {
	return s.raw.RequestResponseTimes(count)
}
