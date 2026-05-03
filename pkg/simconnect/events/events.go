//go:build windows

// Package events contains client event, client data, and flow-event APIs.
package events

import (
	"context"

	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/eventsdata"
)

type Events struct {
	*eventsdata.EventsData
}

// Session is the interface required by New; re-exported from eventsdata.
type Session = eventsdata.Session

type Event = eventsdata.Event
type TransmitOption = eventsdata.TransmitOption
type ClientDataDefinition = eventsdata.ClientDataDefinition
type ClientDataDefinitionItem = eventsdata.ClientDataDefinitionItem

func New(c Session) *Events { return &Events{EventsData: eventsdata.New(c)} }

func WithTransmitObject(objectID core.ObjectID) TransmitOption {
	return eventsdata.WithTransmitObject(objectID)
}

func WithTransmitGroup(groupID core.NotificationGroupID) TransmitOption {
	return eventsdata.WithTransmitGroup(groupID)
}

func WithTransmitFlags(flags core.EventFlag) TransmitOption {
	return eventsdata.WithTransmitFlags(flags)
}

func (e *Events) SubscribeFlow(ctx context.Context) (<-chan core.FlowEventMessage, error) {
	return e.SubscribeFlowEvents(ctx)
}
