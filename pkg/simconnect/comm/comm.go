//go:build windows

// Package comm contains CommBus APIs.
package comm

import (
	"context"

	"github.com/Zwergpro/simconnect-go/pkg/simconnect/communication"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

type Comm struct {
	*communication.Communication
}

// Session is the interface required by New; re-exported from communication.
type Session = communication.Session

func New(c Session) *Comm { return &Comm{Communication: communication.New(c)} }

func (c *Comm) Subscribe(ctx context.Context, name string) (<-chan core.CommBusMessage, error) {
	return c.SubscribeCommBusEvent(ctx, name)
}

func (c *Comm) Call(name string, broadcastTo core.CommBusBroadcastTo, data string) error {
	return c.CallCommBusEvent(name, broadcastTo, data)
}

func (c *Comm) CallBytes(name string, broadcastTo core.CommBusBroadcastTo, data []byte) error {
	return c.CallCommBusEventBytes(name, broadcastTo, data)
}
