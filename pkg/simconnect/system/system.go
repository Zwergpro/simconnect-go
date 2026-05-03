//go:build windows

// Package system contains SimConnect system state, system event, and action APIs.
package system

import (
	"context"

	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/general"
)

type System struct {
	*general.General
}

// Session is the interface required by New; re-exported from general.
type Session = general.Session

type ActionParam = general.ActionParam
type SystemEventSubscription = general.SystemEventSubscription

func New(c Session) *System { return &System{General: general.New(c)} }

func PackActionParams(params ...ActionParam) []byte { return general.PackActionParams(params...) }
func ActionBool(value bool) ActionParam             { return general.ActionBool(value) }
func ActionFloat32(value float32) ActionParam       { return general.ActionFloat32(value) }
func ActionString256(value string) ActionParam      { return general.ActionString256(value) }

func (s *System) SubscribeEvent(ctx context.Context, name string) (<-chan core.ClientEvent, error) {
	return s.SubscribeSystemEvent(ctx, name)
}

func (s *System) SetEventState(eventID uint32, state core.State) error {
	return s.SetSystemEventState(eventID, state)
}

func (s *System) RequestState(ctx context.Context, state string) (core.SystemStateMessage, error) {
	return s.RequestSystemState(ctx, state)
}
