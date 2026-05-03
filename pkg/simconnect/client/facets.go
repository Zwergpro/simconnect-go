//go:build windows

package client

// facets holds lazily-initialised sub-API clients, each created once on first
// access.  Guarded by Sim.facetMu.
//
// Imports here complete the import graph:
//
//	client ─► ai, camera, events, comm, system, flight, facilities, input,
//	          debug, simvar
//
// None of those packages import client (they use Session interfaces and import
// core for shared types), so there is no cycle.

import (
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/ai"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/camera"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/comm"
	dbgpkg "github.com/Zwergpro/simconnect-go/pkg/simconnect/debug"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/events"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/facilities"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/flight"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/input"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/simvar"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/system"
)

// facets holds lazily-initialized facet clients for the Sim facade.
type facets struct {
	ai         *ai.AI
	camera     *camera.Camera
	events     *events.Events
	comm       *comm.Comm
	sys        *system.System
	flight     *flight.Flight
	facilities *facilities.Facilities
	input      *input.Input
	debug      *dbgpkg.Debug
	simvar     *simvar.SimVar
}

// AI returns the AI-object sub-API, creating it on first call.
func (s *Sim) AI() *ai.AI {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.ai == nil {
		s.facetCache.ai = ai.New(s)
	}
	return s.facetCache.ai
}

// Camera returns the camera sub-API, creating it on first call.
func (s *Sim) Camera() *camera.Camera {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.camera == nil {
		s.facetCache.camera = camera.New(s)
	}
	return s.facetCache.camera
}

// Events returns the events/client-data sub-API, creating it on first call.
func (s *Sim) Events() *events.Events {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.events == nil {
		s.facetCache.events = events.New(s)
	}
	return s.facetCache.events
}

// Comm returns the CommBus sub-API, creating it on first call.
func (s *Sim) Comm() *comm.Comm {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.comm == nil {
		s.facetCache.comm = comm.New(s)
	}
	return s.facetCache.comm
}

// System returns the system-state/event/action sub-API, creating it on first call.
func (s *Sim) System() *system.System {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.sys == nil {
		s.facetCache.sys = system.New(s)
	}
	return s.facetCache.sys
}

// Flight returns the flight-file/plan sub-API, creating it on first call.
func (s *Sim) Flight() *flight.Flight {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.flight == nil {
		s.facetCache.flight = flight.New(s)
	}
	return s.facetCache.flight
}

// Facilities returns the facilities sub-API, creating it on first call.
func (s *Sim) Facilities() *facilities.Facilities {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.facilities == nil {
		s.facetCache.facilities = facilities.New(s)
	}
	return s.facetCache.facilities
}

// Input returns the input-event sub-API, creating it on first call.
func (s *Sim) Input() *input.Input {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.input == nil {
		s.facetCache.input = input.New(s)
	}
	return s.facetCache.input
}

// Debug returns the diagnostic sub-API, creating it on first call.
func (s *Sim) Debug() *dbgpkg.Debug {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.debug == nil {
		s.facetCache.debug = dbgpkg.New(s)
	}
	return s.facetCache.debug
}

// Simvar returns the typed sim-variable sub-API, creating it on first call.
func (s *Sim) Simvar() *simvar.SimVar {
	s.facetMu.Lock()
	defer s.facetMu.Unlock()
	if s.facetCache.simvar == nil {
		s.facetCache.simvar = simvar.New(s)
	}
	return s.facetCache.simvar
}
