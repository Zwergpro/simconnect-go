//go:build windows

// Package communication implements the Communication SimConnect API category (CommBus).
package communication

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

// Communication exposes CommBus event subscribe/call functions.
type Communication struct {
	c Session

	mu          sync.Mutex
	commBusSubs map[uint32][]chan core.CommBusMessage
}

// New creates a Communication and registers the CommBus message handler.
func New(c Session) *Communication {
	s := &Communication{
		c:           c,
		commBusSubs: make(map[uint32][]chan core.CommBusMessage),
	}
	c.RegisterHandler(core.RecvIDCommBus, func(msg core.Message) {
		m, ok := msg.(core.CommBusMessage)
		if !ok {
			return
		}
		s.mu.Lock()
		subs := append([]chan core.CommBusMessage(nil), s.commBusSubs[m.EventID]...)
		s.mu.Unlock()
		for _, ch := range subs {
			s.send(ch, m)
		}
	})
	c.RegisterCloseHook(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		for id, subs := range s.commBusSubs {
			for _, ch := range subs {
				close(ch)
			}
			delete(s.commBusSubs, id)
		}
	})
	return s
}

func (s *Communication) SubscribeCommBusEvent(ctx context.Context, name string) (<-chan core.CommBusMessage, error) {
	eventID := core.EventID(s.c.NextEventID())
	if err := s.c.Bindings().SubscribeToCommBusEvent(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID), name); err != nil {
		return nil, err
	}

	ch := make(chan core.CommBusMessage, s.c.ChannelBuffer())
	s.mu.Lock()
	s.commBusSubs[uint32(eventID)] = append(s.commBusSubs[uint32(eventID)], ch)
	s.mu.Unlock()

	go func() {
		<-ctx.Done()
		_ = s.c.Bindings().UnsubscribeToCommBusEvent(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID))
		s.removeCommBusSub(uint32(eventID), ch)
	}()
	return ch, nil
}

func (s *Communication) CallCommBusEvent(name string, broadcastTo core.CommBusBroadcastTo, data string) error {
	return s.c.Bindings().CallCommBusEvent(name, bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO(broadcastTo), data)
}

func (s *Communication) CallCommBusEventBytes(name string, broadcastTo core.CommBusBroadcastTo, data []byte) error {
	return s.c.Bindings().CallCommBusEventBytes(name, bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO(broadcastTo), data)
}

func (s *Communication) removeCommBusSub(eventID uint32, target chan core.CommBusMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	subs := s.commBusSubs[eventID]
	for i, ch := range subs {
		if ch == target {
			close(ch)
			subs = append(subs[:i], subs[i+1:]...)
			break
		}
	}
	if len(subs) == 0 {
		delete(s.commBusSubs, eventID)
		return
	}
	s.commBusSubs[eventID] = subs
}

func (s *Communication) send(ch chan core.CommBusMessage, msg core.CommBusMessage) {
	defer func() { _ = recover() }()
	select {
	case ch <- msg:
	case <-s.c.Context().Done():
	}
}
