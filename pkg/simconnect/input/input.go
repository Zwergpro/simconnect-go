//go:build windows

// Package input implements the InputEvents SimConnect API category.
package input

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"sync"
	"unsafe"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// Session is the subset of client.Client methods used by this package.
type Session interface {
	NextRequestID() uint32
	AddWaiter(uint32) (<-chan core.RequestResult, error)
	RemoveWaiter(uint32)
	AddDataSub(uint32, func(core.Message)) error
	RemoveDataSub(uint32)
	TrackSend(uint32)
	Bindings() *bindings.SimConnect
	RegisterHandler(core.RecvID, func(core.Message))
	RegisterCloseHook(func())
	ChannelBuffer() int
	Context() context.Context
}

// Input exposes input event enumeration, get/set, and subscription functions.
type Input struct {
	c Session

	mu               sync.Mutex
	inputSubs        map[uint64][]chan core.InputEventSubscriptionMessage
	paramWaiters     map[uint64]chan core.RequestResult
	controllerWaiter chan core.RequestResult
}

// New creates an Input client and registers specialized message handlers.
func New(c Session) *Input {
	s := &Input{
		c:            c,
		inputSubs:    make(map[uint64][]chan core.InputEventSubscriptionMessage),
		paramWaiters: make(map[uint64]chan core.RequestResult),
	}
	c.RegisterHandler(core.RecvIDSubscribeInputEvent, func(msg core.Message) {
		m, ok := msg.(core.InputEventSubscriptionMessage)
		if !ok {
			return
		}
		s.mu.Lock()
		subs := append([]chan core.InputEventSubscriptionMessage(nil), s.inputSubs[m.Hash]...)
		if m.Hash != 0 {
			subs = append(subs, s.inputSubs[0]...)
		}
		s.mu.Unlock()
		for _, ch := range subs {
			s.sendInputEventSub(ch, m)
		}
	})
	c.RegisterHandler(core.RecvIDEnumerateInputEventParams, func(msg core.Message) {
		m, ok := msg.(core.InputEventParamsMessage)
		if !ok {
			return
		}
		s.mu.Lock()
		if w, ok := s.paramWaiters[m.Hash]; ok {
			delete(s.paramWaiters, m.Hash)
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		s.mu.Unlock()
	})
	c.RegisterHandler(core.RecvIDControllersList, func(msg core.Message) {
		m, ok := msg.(core.ControllersListMessage)
		if !ok {
			return
		}
		s.mu.Lock()
		if w := s.controllerWaiter; w != nil {
			s.controllerWaiter = nil
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		s.mu.Unlock()
	})
	c.RegisterCloseHook(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		for hash, subs := range s.inputSubs {
			for _, ch := range subs {
				safeClose(ch)
			}
			delete(s.inputSubs, hash)
		}
		for hash, w := range s.paramWaiters {
			w <- core.RequestResult{Err: core.ErrClosed}
			close(w)
			delete(s.paramWaiters, hash)
		}
		if s.controllerWaiter != nil {
			s.controllerWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(s.controllerWaiter)
			s.controllerWaiter = nil
		}
	})
	return s
}

func (s *Input) EnumerateInputEvents(ctx context.Context) (core.InputEventListMessage, error) {
	var zero core.InputEventListMessage

	requestID := s.c.NextRequestID()
	waiter, err := s.c.AddWaiter(requestID)
	if err != nil {
		return zero, err
	}

	packets := make(chan core.InputEventListMessage, 16)
	handler := func(msg core.Message) {
		inputMsg, ok := msg.(core.InputEventListMessage)
		if !ok {
			return
		}
		select {
		case packets <- inputMsg:
		case <-s.c.Context().Done():
		}
	}

	if err := s.c.AddDataSub(requestID, handler); err != nil {
		s.c.RemoveWaiter(requestID)
		return zero, err
	}
	defer s.c.RemoveDataSub(requestID)

	if err := s.c.Bindings().EnumerateInputEvents(bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveWaiter(requestID)
		return zero, err
	}
	s.c.TrackSend(requestID)

	for {
		select {
		case result, ok := <-waiter:
			if ok && result.Err != nil {
				return zero, result.Err
			}
			waiter = nil
		case packet := <-packets:
			if zero.OutOf == 0 || packet.EntryNumber < zero.EntryNumber {
				zero.FacilityListMeta = packet.FacilityListMeta
			}
			zero.Events = append(zero.Events, packet.Events...)
			zero.ArraySize = uint32(len(zero.Events))
			if packet.OutOf == 0 || packet.EntryNumber+1 >= packet.OutOf {
				return zero, nil
			}
		case <-ctx.Done():
			s.c.RemoveWaiter(requestID)
			return zero, ctx.Err()
		}
	}
}

func (s *Input) GetInputEvent(ctx context.Context, hash uint64) (core.InputEventValueMessage, error) {
	var zero core.InputEventValueMessage
	requestID := s.c.NextRequestID()
	waiter, err := s.c.AddWaiter(requestID)
	if err != nil {
		return zero, err
	}
	if err := s.c.Bindings().GetInputEvent(bindings.SIMCONNECT_DATA_REQUEST_ID(requestID), hash); err != nil {
		s.c.RemoveWaiter(requestID)
		return zero, err
	}
	s.c.TrackSend(requestID)

	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.c.RemoveWaiter(requestID)
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	msg, ok := result.Msg.(core.InputEventValueMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected input event value, got %T", core.ErrDecode, result.Msg)
	}
	return msg, nil
}

func (s *Input) SetInputEventDouble(hash uint64, value float64) error {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(value))
	return s.c.Bindings().SetInputEvent(hash, uint32(len(buf)), unsafe.Pointer(&buf[0]))
}

func (s *Input) SetInputEventString(hash uint64, value string) error {
	buf := append([]byte(value), 0)
	return s.c.Bindings().SetInputEvent(hash, uint32(len(buf)), unsafe.Pointer(&buf[0]))
}

func (s *Input) EnumerateInputEventParams(ctx context.Context, hash uint64) (core.InputEventParamsMessage, error) {
	var zero core.InputEventParamsMessage
	waiter, err := s.addParamWaiter(hash)
	if err != nil {
		return zero, err
	}
	if err := s.c.Bindings().EnumerateInputEventParams(hash); err != nil {
		s.removeParamWaiter(hash)
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.removeParamWaiter(hash)
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	msg, ok := result.Msg.(core.InputEventParamsMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected input event params, got %T", core.ErrDecode, result.Msg)
	}
	return msg, nil
}

func (s *Input) SubscribeInputEvent(ctx context.Context, hash uint64) (<-chan core.InputEventSubscriptionMessage, error) {
	ch := make(chan core.InputEventSubscriptionMessage, s.c.ChannelBuffer())
	s.mu.Lock()
	s.inputSubs[hash] = append(s.inputSubs[hash], ch)
	s.mu.Unlock()

	if err := s.c.Bindings().SubscribeInputEvent(hash); err != nil {
		s.removeInputEventSub(hash, ch)
		return nil, err
	}

	go func() {
		<-ctx.Done()
		s.removeInputEventSub(hash, ch)
	}()
	return ch, nil
}

func (s *Input) UnsubscribeInputEvent(hash uint64) error {
	return s.c.Bindings().UnsubscribeInputEvent(hash)
}

func (s *Input) EnumerateControllers(ctx context.Context) (core.ControllersListMessage, error) {
	var zero core.ControllersListMessage
	waiter, err := s.addControllerWaiter()
	if err != nil {
		return zero, err
	}
	if err := s.c.Bindings().EnumerateControllers(); err != nil {
		s.removeControllerWaiter()
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.removeControllerWaiter()
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	msg, ok := result.Msg.(core.ControllersListMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected controllers list, got %T", core.ErrDecode, result.Msg)
	}
	return msg, nil
}

func (s *Input) addParamWaiter(hash uint64) (<-chan core.RequestResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ch := make(chan core.RequestResult, 1)
	s.paramWaiters[hash] = ch
	return ch, nil
}

func (s *Input) removeParamWaiter(hash uint64) {
	s.mu.Lock()
	if w, ok := s.paramWaiters[hash]; ok {
		delete(s.paramWaiters, hash)
		close(w)
	}
	s.mu.Unlock()
}

func (s *Input) addControllerWaiter() (<-chan core.RequestResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.controllerWaiter != nil {
		return nil, fmt.Errorf("%w: controller enumeration already pending", core.ErrDecode)
	}
	ch := make(chan core.RequestResult, 1)
	s.controllerWaiter = ch
	return ch, nil
}

func (s *Input) removeControllerWaiter() {
	s.mu.Lock()
	if s.controllerWaiter != nil {
		close(s.controllerWaiter)
		s.controllerWaiter = nil
	}
	s.mu.Unlock()
}

func (s *Input) removeInputEventSub(hash uint64, ch chan core.InputEventSubscriptionMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	subs := s.inputSubs[hash]
	for i, sub := range subs {
		if sub == ch {
			safeClose(ch)
			s.inputSubs[hash] = append(subs[:i], subs[i+1:]...)
			break
		}
	}
	if len(s.inputSubs[hash]) == 0 {
		delete(s.inputSubs, hash)
		_ = s.c.Bindings().UnsubscribeInputEvent(hash)
	}
}

func (s *Input) sendInputEventSub(ch chan core.InputEventSubscriptionMessage, msg core.InputEventSubscriptionMessage) {
	defer func() { _ = recover() }()
	select {
	case ch <- msg:
	case <-s.c.Context().Done():
	}
}

func safeClose(ch chan core.InputEventSubscriptionMessage) {
	defer func() { _ = recover() }()
	close(ch)
}
