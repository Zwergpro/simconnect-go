//go:build windows

// Package simvar contains the typed SimVar definition and data request API.
package simvar

import (
	"context"

	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/eventsdata"
)

type Definition[T any] = eventsdata.Definition[T]
type Field = eventsdata.Field
type Update[T any] = eventsdata.DataUpdate[T]
type Option = eventsdata.DataOption

// Session is the interface required by New; re-exported from eventsdata.
type Session = eventsdata.Session

// SimVar wraps an eventsdata.EventsData to provide typed sim-variable access.
type SimVar struct {
	eventsData *eventsdata.EventsData
}

// New creates a SimVar client backed by the provided eventsdata.Session.
func New(c Session) *SimVar {
	return &SimVar{eventsData: eventsdata.New(c)}
}

func ChangedOnly() Option {
	return eventsdata.ChangedOnly()
}

func WithTiming(origin, interval, limit uint32) Option {
	return eventsdata.WithDataRequestTiming(origin, interval, limit)
}

func Define[T any]() (*Definition[T], error) {
	return eventsdata.Define[T]()
}

func DefineFields[T any](fields ...Field) (*Definition[T], error) {
	return eventsdata.DefineFields[T](fields...)
}

// GetOnce requests a sim variable once for the given object.
func (sv *SimVar) GetOnce(ctx context.Context, def *Definition[any], obj core.ObjectID) (any, error) {
	return eventsdata.RequestDataOnce(ctx, sv.eventsData, def, obj)
}

// GetByTypeOnce requests a sim variable once by object type.
func (sv *SimVar) GetByTypeOnce(ctx context.Context, def *Definition[any], radiusMeters uint32, t core.SimObjectType) (Update[any], error) {
	return eventsdata.RequestDataByTypeOnce(ctx, sv.eventsData, def, radiusMeters, t)
}

// Set sets a sim variable on the given object.
func (sv *SimVar) Set(ctx context.Context, def *Definition[any], obj core.ObjectID, value any, flags core.DataSetFlag) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	return eventsdata.SetDataOnSimObject(sv.eventsData, def, obj, value, flags)
}

// Subscribe subscribes to sim variable updates for the given object.
func (sv *SimVar) Subscribe(ctx context.Context, def *Definition[any], obj core.ObjectID, period core.Period, opts ...Option) (<-chan Update[any], error) {
	return eventsdata.SubscribeData(ctx, sv.eventsData, def, obj, period, opts...)
}

// Package-level generic helpers preserved for backward compatibility.
// These create a one-shot eventsdata.EventsData each call; prefer New() + methods.

func GetOnce[T any](ctx context.Context, s Session, def *Definition[T], obj core.ObjectID) (T, error) {
	return eventsdata.RequestDataOnce(ctx, eventsdata.New(s), def, obj)
}

func GetByTypeOnce[T any](ctx context.Context, s Session, def *Definition[T], radiusMeters uint32, t core.SimObjectType) (Update[T], error) {
	return eventsdata.RequestDataByTypeOnce(ctx, eventsdata.New(s), def, radiusMeters, t)
}

func Set[T any](ctx context.Context, s Session, def *Definition[T], obj core.ObjectID, value T, flags core.DataSetFlag) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	return eventsdata.SetDataOnSimObject(eventsdata.New(s), def, obj, value, flags)
}

func Subscribe[T any](ctx context.Context, s Session, def *Definition[T], obj core.ObjectID, period core.Period, opts ...Option) (<-chan Update[T], error) {
	return eventsdata.SubscribeData(ctx, eventsdata.New(s), def, obj, period, opts...)
}
