//go:build windows

// Package general implements the General SimConnect API category.
package general

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
	NextEventID() uint32
	NextRequestID() uint32
	AddWaiter(uint32) (<-chan core.RequestResult, error)
	RemoveWaiter(uint32)
	TrackSend(uint32)
	Bindings() *bindings.SimConnect
	RegisterHandler(core.RecvID, func(core.Message))
	RegisterCloseHook(func())
	ChannelBuffer() int
	Context() context.Context
}

// General exposes general SimConnect functions: system events, state requests,
// action parameters, enumeration, and notification groups.
type General struct {
	session Session

	mu        sync.Mutex
	eventSubs map[uint32][]chan core.ClientEvent
}

// New creates a General and registers client-event handlers.
func New(s Session) *General {
	g := &General{
		session:   s,
		eventSubs: make(map[uint32][]chan core.ClientEvent),
	}
	dispatch := func(event core.ClientEvent) {
		g.mu.Lock()
		subs := append([]chan core.ClientEvent(nil), g.eventSubs[event.EventID]...)
		g.mu.Unlock()
		for _, ch := range subs {
			g.sendEvent(ch, event)
		}
	}
	s.RegisterHandler(core.RecvIDEvent, func(msg core.Message) {
		if m, ok := msg.(core.ClientEvent); ok {
			dispatch(m)
		}
	})
	s.RegisterHandler(core.RecvIDEventEX1, func(msg core.Message) {
		if m, ok := msg.(core.ClientEventEX1); ok {
			dispatch(m.ClientEvent)
		}
	})
	s.RegisterHandler(core.RecvIDEventFilename, func(msg core.Message) {
		if m, ok := msg.(core.FilenameEvent); ok {
			dispatch(m.ClientEvent)
		}
	})
	s.RegisterHandler(core.RecvIDEventObjectAddRemove, func(msg core.Message) {
		if m, ok := msg.(core.ObjectAddRemoveEvent); ok {
			dispatch(m.ClientEvent)
		}
	})
	s.RegisterHandler(core.RecvIDEventFrame, func(msg core.Message) {
		if m, ok := msg.(core.FrameEvent); ok {
			dispatch(m.ClientEvent)
		}
	})
	s.RegisterCloseHook(func() {
		g.mu.Lock()
		defer g.mu.Unlock()
		for id, subs := range g.eventSubs {
			for _, ch := range subs {
				close(ch)
			}
			delete(g.eventSubs, id)
		}
	})
	return g
}

// --- Action parameters ---

type ActionParam func(*[]byte)

func PackActionParams(params ...ActionParam) []byte {
	var data []byte
	for _, param := range params {
		param(&data)
	}
	return data
}

func ActionBool(value bool) ActionParam {
	return func(data *[]byte) {
		if value {
			*data = append(*data, 1)
		} else {
			*data = append(*data, 0)
		}
	}
}

func ActionFloat32(value float32) ActionParam {
	return func(data *[]byte) {
		var buf [4]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(value))
		*data = append(*data, buf[:]...)
	}
}

func ActionString256(value string) ActionParam {
	return func(data *[]byte) {
		var buf [256]byte
		copy(buf[:], value)
		*data = append(*data, buf[:]...)
	}
}

// --- System events ---

type SystemEventSubscription struct {
	EventID uint32
	Events  <-chan core.ClientEvent
}

// SubscribeSystemEvent subscribes to a named system event (e.g. "Frame", "1sec").
func (g *General) SubscribeSystemEvent(ctx context.Context, name string) (<-chan core.ClientEvent, error) {
	eventID := core.EventID(g.session.NextEventID())
	return g.subscribeSystemEvent(ctx, eventID, name)
}

// SubscribeSystemEventWithID subscribes and returns the event handle alongside the channel.
func (g *General) SubscribeSystemEventWithID(ctx context.Context, name string) (*SystemEventSubscription, error) {
	eventID := core.EventID(g.session.NextEventID())
	ch, err := g.subscribeSystemEvent(ctx, eventID, name)
	if err != nil {
		return nil, err
	}
	return &SystemEventSubscription{EventID: uint32(eventID), Events: ch}, nil
}

func (g *General) subscribeSystemEvent(ctx context.Context, eventID core.EventID, name string) (<-chan core.ClientEvent, error) {
	if err := g.session.Bindings().SubscribeToSystemEvent(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID), name); err != nil {
		return nil, err
	}
	ch := make(chan core.ClientEvent, g.session.ChannelBuffer())
	g.mu.Lock()
	g.eventSubs[uint32(eventID)] = append(g.eventSubs[uint32(eventID)], ch)
	g.mu.Unlock()

	go func() {
		<-ctx.Done()
		_ = g.session.Bindings().UnsubscribeFromSystemEvent(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID))
		g.removeEventSub(uint32(eventID), ch)
	}()
	return ch, nil
}

func (g *General) SetSystemEventState(eventID uint32, state core.State) error {
	return g.session.Bindings().SetSystemEventState(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID), bindings.SIMCONNECT_STATE(state))
}

func (g *General) SetNotificationGroupPriority(groupID core.NotificationGroupID, priority uint32) error {
	return g.session.Bindings().SetNotificationGroupPriority(bindings.SIMCONNECT_NOTIFICATION_GROUP_ID(groupID), priority)
}

// RequestSystemState requests a named simulator state value (e.g. "AircraftLoaded").
func (g *General) RequestSystemState(ctx context.Context, state string) (core.SystemStateMessage, error) {
	var zero core.SystemStateMessage
	requestID := g.session.NextRequestID()
	waiter, err := g.session.AddWaiter(requestID)
	if err != nil {
		return zero, err
	}
	if err := g.session.Bindings().RequestSystemState(bindings.SIMCONNECT_DATA_REQUEST_ID(requestID), state); err != nil {
		g.session.RemoveWaiter(requestID)
		return zero, err
	}
	g.session.TrackSend(requestID)

	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		g.session.RemoveWaiter(requestID)
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	msg, ok := result.Msg.(core.SystemStateMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected system state, got %T", core.ErrDecode, result.Msg)
	}
	return msg, nil
}

// ExecuteAction executes a named SimConnect action with packed parameters.
func (g *General) ExecuteAction(ctx context.Context, actionID string, params ...ActionParam) (core.ActionCallbackMessage, error) {
	var zero core.ActionCallbackMessage
	packed := PackActionParams(params...)
	requestID := g.session.NextRequestID()
	waiter, err := g.session.AddWaiter(requestID)
	if err != nil {
		return zero, err
	}
	var ptr unsafe.Pointer
	if len(packed) > 0 {
		ptr = unsafe.Pointer(&packed[0])
	}
	if err := g.session.Bindings().ExecuteAction(requestID, actionID, uint32(len(packed)), ptr); err != nil {
		g.session.RemoveWaiter(requestID)
		return zero, err
	}
	g.session.TrackSend(requestID)

	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		g.session.RemoveWaiter(requestID)
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	msg, ok := result.Msg.(core.ActionCallbackMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected action callback, got %T", core.ErrDecode, result.Msg)
	}
	return msg, nil
}

// EnumerateSimObjectsAndLiveries lists all sim objects of the given type with their liveries.
// Placed here per apidocs General category.
func (g *General) EnumerateSimObjectsAndLiveries(ctx context.Context, objType core.SimObjectType) (core.SimObjectLiveryListMessage, error) {
	requestID := g.session.NextRequestID()
	waiter, err := g.session.AddWaiter(requestID)
	if err != nil {
		return core.SimObjectLiveryListMessage{}, err
	}
	if err := g.session.Bindings().EnumerateSimObjectsAndLiveries(bindings.SIMCONNECT_DATA_REQUEST_ID(requestID), bindings.SIMCONNECT_SIMOBJECT_TYPE(objType)); err != nil {
		g.session.RemoveWaiter(requestID)
		return core.SimObjectLiveryListMessage{}, err
	}
	g.session.TrackSend(requestID)

	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		g.session.RemoveWaiter(requestID)
		return core.SimObjectLiveryListMessage{}, ctx.Err()
	}
	if result.Err != nil {
		return core.SimObjectLiveryListMessage{}, result.Err
	}
	msg, ok := result.Msg.(core.SimObjectLiveryListMessage)
	if !ok {
		return core.SimObjectLiveryListMessage{}, fmt.Errorf("%w: expected simobject livery list, got %T", core.ErrDecode, result.Msg)
	}
	return msg, nil
}

func (g *General) removeEventSub(eventID uint32, target chan core.ClientEvent) {
	g.mu.Lock()
	defer g.mu.Unlock()
	subs := g.eventSubs[eventID]
	for i, ch := range subs {
		if ch == target {
			close(ch)
			subs = append(subs[:i], subs[i+1:]...)
			break
		}
	}
	if len(subs) == 0 {
		delete(g.eventSubs, eventID)
		return
	}
	g.eventSubs[eventID] = subs
}

func (g *General) sendEvent(ch chan core.ClientEvent, event core.ClientEvent) {
	defer func() { _ = recover() }()
	select {
	case ch <- event:
	case <-g.session.Context().Done():
	}
}
