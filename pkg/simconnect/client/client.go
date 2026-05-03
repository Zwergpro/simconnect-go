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

// Client is the core SimConnect session. Domain packages (ai, camera, etc.) wrap
// a *Client to add API surface; all share the same ID allocators and dispatch loop.
type Client struct {
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

// Sim is the public session facade for a SimConnect connection.
type Sim = Client

func Dial(ctx context.Context, appName string, opts ...Option) (*Client, error) {
	cfg := defaultClientConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	raw, err := bindings.Open(appName, bindings.HWND(cfg.hwnd), cfg.eventID, bindings.HANDLE(cfg.eventHandle), cfg.configIndex)
	if err != nil {
		return nil, err
	}

	runCtx, cancel := context.WithCancel(ctx)
	c := &Client{
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
		close(c.done)
		return c, nil
	}

	go c.run()
	return c, nil
}

// Open connects to SimConnect and starts the session dispatch loop.
func Open(ctx context.Context, appName string, opts ...Option) (*Sim, error) {
	return Dial(ctx, appName, opts...)
}

func (c *Client) Errors() <-chan error     { return c.errs }
func (c *Client) Context() context.Context { return c.ctx }
func (c *Client) ChannelBuffer() int       { return c.cfg.channelBuffer }

// Bindings exposes the low-level binding session for facet packages.
func (c *Client) Bindings() *bindings.SimConnect { return c.raw }

func (c *Client) NextRequestID() uint32    { return c.reqIDs.Next() }
func (c *Client) NextDefinitionID() uint32 { return c.defIDs.Next() }
func (c *Client) NextEventID() uint32      { return c.eventIDs.Next() }

// RegisterHandler registers fn to receive messages of the given recv ID.
// Handlers run in the dispatch goroutine and must not block.
func (c *Client) RegisterHandler(id core.RecvID, fn func(core.Message)) {
	c.mu.Lock()
	c.handlers[id] = append(c.handlers[id], fn)
	c.mu.Unlock()
}

// RegisterCloseHook registers a cleanup function called during Close,
// after the context is cancelled and the dispatch loop has stopped.
func (c *Client) RegisterCloseHook(fn func()) {
	c.mu.Lock()
	c.closeHooks = append(c.closeHooks, fn)
	c.mu.Unlock()
}

// AddWaiter registers a one-shot result channel for the given request ID.
func (c *Client) AddWaiter(reqID uint32) (<-chan core.RequestResult, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return nil, core.ErrClosed
	}
	ch := make(chan core.RequestResult, 1)
	c.waiters[reqID] = ch
	return ch, nil
}

// RemoveWaiter removes and closes the waiter for reqID (used on cancellation).
func (c *Client) RemoveWaiter(reqID uint32) {
	c.mu.Lock()
	if waiter, ok := c.waiters[reqID]; ok {
		delete(c.waiters, reqID)
		close(waiter)
	}
	c.mu.Unlock()
}

// AddDataSub registers a continuous data handler keyed by requestID.
func (c *Client) AddDataSub(reqID uint32, fn func(core.Message)) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return core.ErrClosed
	}
	c.dataSubs[reqID] = append(c.dataSubs[reqID], fn)
	return nil
}

// RemoveDataSub removes all data handlers for reqID.
func (c *Client) RemoveDataSub(reqID uint32) {
	c.mu.Lock()
	delete(c.dataSubs, reqID)
	c.mu.Unlock()
}

// TrackSend maps the last sent packet ID to reqID for exception correlation.
func (c *Client) TrackSend(reqID uint32) {
	sendID, err := c.raw.GetLastSentPacketID()
	if err != nil {
		return
	}
	c.mu.Lock()
	c.sendToReq[sendID] = reqID
	c.mu.Unlock()
}

func (c *Client) ReportError(err error) {
	select {
	case c.errs <- err:
	case <-c.ctx.Done():
	}
}

func (c *Client) Close() error {
	c.closeOnce.Do(func() {
		c.cancel()
		if !c.cfg.manualDispatch {
			<-c.done
		}

		c.mu.Lock()
		c.closed = true
		hooks := append([]func(){}, c.closeHooks...)
		for id, waiter := range c.waiters {
			waiter <- core.RequestResult{Err: core.ErrClosed}
			close(waiter)
			delete(c.waiters, id)
		}
		c.dataSubs = map[uint32][]func(core.Message){}
		c.mu.Unlock()

		for _, fn := range hooks {
			fn()
		}

		c.closeErr = c.raw.Close()
		close(c.errs)
	})
	return c.closeErr
}

func (c *Client) run() {
	defer close(c.done)
	ticker := time.NewTicker(c.cfg.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			if err := c.Poll(); err != nil {
				c.ReportError(err)
			}
		}
	}
}

func (c *Client) Poll() error {
	for {
		raw, size, err := c.raw.GetNextDispatch()
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
		c.dispatch(msg)
	}
}

// DispatchProc is the callback type for CallDispatch.
type DispatchProc func(msg core.Message, context uintptr)

// GetNextMessage retrieves the next pending SimConnect message without blocking.
// Returns (nil, false, nil) when the queue is empty.
func (c *Client) GetNextMessage() (core.Message, bool, error) {
	raw, size, err := c.raw.GetNextDispatch()
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
func (c *Client) CallDispatch(fn DispatchProc, context uintptr) error {
	return c.raw.CallDispatch(func(raw *bindings.SIMCONNECT_RECV, size uint32, ctx uintptr) {
		msg, err := decodeMessage(raw, size)
		if err != nil {
			c.ReportError(err)
			return
		}
		fn(msg, ctx)
	}, context)
}

// GetLastSentPacketID returns the packet ID of the most recently sent call.
func (c *Client) GetLastSentPacketID() (uint32, error) {
	return c.raw.GetLastSentPacketID()
}

// RequestResponseTimes returns response-time measurements for the last count packets.
func (c *Client) RequestResponseTimes(count uint32) ([]float32, error) {
	return c.raw.RequestResponseTimes(count)
}
