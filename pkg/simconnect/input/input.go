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

// Session is the subset of client.Sim methods used by this package.
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
	session Session

	mu               sync.Mutex
	inputSubs        map[uint64][]chan core.InputEventSubscriptionMessage
	paramWaiters     map[uint64]chan core.RequestResult
	controllerWaiter chan core.RequestResult
}

// New creates an Input client and registers specialized message handlers.
func New(s Session) *Input {
	i := &Input{
		session:      s,
		inputSubs:    make(map[uint64][]chan core.InputEventSubscriptionMessage),
		paramWaiters: make(map[uint64]chan core.RequestResult),
	}
	s.RegisterHandler(core.RecvIDSubscribeInputEvent, func(msg core.Message) {
		m, ok := msg.(core.InputEventSubscriptionMessage)
		if !ok {
			return
		}
		i.mu.Lock()
		subs := append([]chan core.InputEventSubscriptionMessage(nil), i.inputSubs[m.Hash]...)
		if m.Hash != 0 {
			subs = append(subs, i.inputSubs[0]...)
		}
		i.mu.Unlock()
		for _, ch := range subs {
			i.sendInputEventSub(ch, m)
		}
	})
	s.RegisterHandler(core.RecvIDEnumerateInputEventParams, func(msg core.Message) {
		m, ok := msg.(core.InputEventParamsMessage)
		if !ok {
			return
		}
		i.mu.Lock()
		if w, ok := i.paramWaiters[m.Hash]; ok {
			delete(i.paramWaiters, m.Hash)
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		i.mu.Unlock()
	})
	s.RegisterHandler(core.RecvIDControllersList, func(msg core.Message) {
		m, ok := msg.(core.ControllersListMessage)
		if !ok {
			return
		}
		i.mu.Lock()
		if w := i.controllerWaiter; w != nil {
			i.controllerWaiter = nil
			w <- core.RequestResult{Msg: m}
			close(w)
		}
		i.mu.Unlock()
	})
	s.RegisterCloseHook(func() {
		i.mu.Lock()
		defer i.mu.Unlock()
		for hash, subs := range i.inputSubs {
			for _, ch := range subs {
				safeClose(ch)
			}
			delete(i.inputSubs, hash)
		}
		for hash, w := range i.paramWaiters {
			w <- core.RequestResult{Err: core.ErrClosed}
			close(w)
			delete(i.paramWaiters, hash)
		}
		if i.controllerWaiter != nil {
			i.controllerWaiter <- core.RequestResult{Err: core.ErrClosed}
			close(i.controllerWaiter)
			i.controllerWaiter = nil
		}
	})
	return i
}

func (i *Input) EnumerateInputEvents(ctx context.Context) (core.InputEventListMessage, error) {
	var zero core.InputEventListMessage

	requestID := i.session.NextRequestID()
	waiter, err := i.session.AddWaiter(requestID)
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
		case <-i.session.Context().Done():
		}
	}

	if err := i.session.AddDataSub(requestID, handler); err != nil {
		i.session.RemoveWaiter(requestID)
		return zero, err
	}
	defer i.session.RemoveDataSub(requestID)

	if err := i.session.Bindings().EnumerateInputEvents(bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		i.session.RemoveWaiter(requestID)
		return zero, err
	}
	i.session.TrackSend(requestID)

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
			i.session.RemoveWaiter(requestID)
			return zero, ctx.Err()
		}
	}
}

func (i *Input) GetInputEvent(ctx context.Context, hash uint64) (core.InputEventValueMessage, error) {
	var zero core.InputEventValueMessage
	requestID := i.session.NextRequestID()
	waiter, err := i.session.AddWaiter(requestID)
	if err != nil {
		return zero, err
	}
	if err := i.session.Bindings().GetInputEvent(bindings.SIMCONNECT_DATA_REQUEST_ID(requestID), hash); err != nil {
		i.session.RemoveWaiter(requestID)
		return zero, err
	}
	i.session.TrackSend(requestID)

	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		i.session.RemoveWaiter(requestID)
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

func (i *Input) SetInputEventDouble(hash uint64, value float64) error {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(value))
	return i.session.Bindings().SetInputEvent(hash, uint32(len(buf)), unsafe.Pointer(&buf[0]))
}

func (i *Input) SetInputEventString(hash uint64, value string) error {
	buf := append([]byte(value), 0)
	return i.session.Bindings().SetInputEvent(hash, uint32(len(buf)), unsafe.Pointer(&buf[0]))
}

func (i *Input) EnumerateInputEventParams(ctx context.Context, hash uint64) (core.InputEventParamsMessage, error) {
	var zero core.InputEventParamsMessage
	waiter, err := i.addParamWaiter(hash)
	if err != nil {
		return zero, err
	}
	if err := i.session.Bindings().EnumerateInputEventParams(hash); err != nil {
		i.removeParamWaiter(hash)
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		i.removeParamWaiter(hash)
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

func (i *Input) SubscribeInputEvent(ctx context.Context, hash uint64) (<-chan core.InputEventSubscriptionMessage, error) {
	ch := make(chan core.InputEventSubscriptionMessage, i.session.ChannelBuffer())
	i.mu.Lock()
	i.inputSubs[hash] = append(i.inputSubs[hash], ch)
	i.mu.Unlock()

	if err := i.session.Bindings().SubscribeInputEvent(hash); err != nil {
		i.removeInputEventSub(hash, ch)
		return nil, err
	}

	go func() {
		<-ctx.Done()
		i.removeInputEventSub(hash, ch)
	}()
	return ch, nil
}

func (i *Input) UnsubscribeInputEvent(hash uint64) error {
	return i.session.Bindings().UnsubscribeInputEvent(hash)
}

func (i *Input) EnumerateControllers(ctx context.Context) (core.ControllersListMessage, error) {
	var zero core.ControllersListMessage
	waiter, err := i.addControllerWaiter()
	if err != nil {
		return zero, err
	}
	if err := i.session.Bindings().EnumerateControllers(); err != nil {
		i.removeControllerWaiter()
		return zero, err
	}
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		i.removeControllerWaiter()
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

func (i *Input) addParamWaiter(hash uint64) (<-chan core.RequestResult, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	ch := make(chan core.RequestResult, 1)
	i.paramWaiters[hash] = ch
	return ch, nil
}

func (i *Input) removeParamWaiter(hash uint64) {
	i.mu.Lock()
	if w, ok := i.paramWaiters[hash]; ok {
		delete(i.paramWaiters, hash)
		close(w)
	}
	i.mu.Unlock()
}

func (i *Input) addControllerWaiter() (<-chan core.RequestResult, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.controllerWaiter != nil {
		return nil, fmt.Errorf("%w: controller enumeration already pending", core.ErrDecode)
	}
	ch := make(chan core.RequestResult, 1)
	i.controllerWaiter = ch
	return ch, nil
}

func (i *Input) removeControllerWaiter() {
	i.mu.Lock()
	if i.controllerWaiter != nil {
		close(i.controllerWaiter)
		i.controllerWaiter = nil
	}
	i.mu.Unlock()
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
		_ = s.session.Bindings().UnsubscribeInputEvent(hash)
	}
}

func (i *Input) sendInputEventSub(ch chan core.InputEventSubscriptionMessage, msg core.InputEventSubscriptionMessage) {
	defer func() { _ = recover() }()
	select {
	case ch <- msg:
	case <-i.session.Context().Done():
	}
}

func safeClose(ch chan core.InputEventSubscriptionMessage) {
	defer func() { _ = recover() }()
	close(ch)
}
