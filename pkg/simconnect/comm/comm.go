//go:build windows

// Package comm implements the Communication SimConnect API category (CommBus).
package comm

import (
	"context"
	"sync"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// Session is the subset of client.Sim methods used by this package.
type Session interface {
	NextEventID() uint32
	Bindings() *bindings.SimConnect
	RegisterHandler(core.RecvID, func(core.Message))
	RegisterCloseHook(func())
	ChannelBuffer() int
	Context() context.Context
}

// Comm exposes CommBus event subscribe/call functions.
type Comm struct {
	session Session

	mu   sync.Mutex
	subs map[uint32][]chan core.CommBusMessage
}

// New creates a Comm and registers the CommBus message handler.
func New(s Session) *Comm {
	c := &Comm{
		session: s,
		subs:    make(map[uint32][]chan core.CommBusMessage),
	}
	s.RegisterHandler(core.RecvIDCommBus, func(msg core.Message) {
		m, ok := msg.(core.CommBusMessage)
		if !ok {
			return
		}
		c.mu.Lock()
		subs := append([]chan core.CommBusMessage(nil), c.subs[m.EventID]...)
		c.mu.Unlock()
		for _, ch := range subs {
			c.send(ch, m)
		}
	})
	s.RegisterCloseHook(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		for id, subs := range c.subs {
			for _, ch := range subs {
				close(ch)
			}
			delete(c.subs, id)
		}
	})
	return c
}

// Subscribe subscribes to a CommBus event by name and returns a channel
// that receives messages until ctx is cancelled.
func (c *Comm) Subscribe(ctx context.Context, name string) (<-chan core.CommBusMessage, error) {
	eventID := core.EventID(c.session.NextEventID())
	if err := c.session.Bindings().SubscribeToCommBusEvent(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID), name); err != nil {
		return nil, err
	}

	ch := make(chan core.CommBusMessage, c.session.ChannelBuffer())
	c.mu.Lock()
	c.subs[uint32(eventID)] = append(c.subs[uint32(eventID)], ch)
	c.mu.Unlock()

	go func() {
		<-ctx.Done()
		_ = c.session.Bindings().UnsubscribeToCommBusEvent(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID))
		c.removeSub(uint32(eventID), ch)
	}()
	return ch, nil
}

// Call broadcasts a string payload over a CommBus event.
func (c *Comm) Call(name string, broadcastTo core.CommBusBroadcastTo, data string) error {
	return c.session.Bindings().CallCommBusEvent(name, bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO(broadcastTo), data)
}

// CallBytes broadcasts a raw byte payload over a CommBus event.
func (c *Comm) CallBytes(name string, broadcastTo core.CommBusBroadcastTo, data []byte) error {
	return c.session.Bindings().CallCommBusEventBytes(name, bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO(broadcastTo), data)
}

func (c *Comm) removeSub(eventID uint32, target chan core.CommBusMessage) {
	c.mu.Lock()
	defer c.mu.Unlock()
	subs := c.subs[eventID]
	for i, ch := range subs {
		if ch == target {
			close(ch)
			subs = append(subs[:i], subs[i+1:]...)
			break
		}
	}
	if len(subs) == 0 {
		delete(c.subs, eventID)
		return
	}
	c.subs[eventID] = subs
}

func (c *Comm) send(ch chan core.CommBusMessage, msg core.CommBusMessage) {
	defer func() { _ = recover() }()
	select {
	case ch <- msg:
	case <-c.session.Context().Done():
	}
}
