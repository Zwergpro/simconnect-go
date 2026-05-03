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
func (c *Sim) AI() *ai.AI {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.ai == nil {
		c.facetCache.ai = ai.New(c)
	}
	return c.facetCache.ai
}

// Camera returns the camera sub-API, creating it on first call.
func (c *Sim) Camera() *camera.Camera {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.camera == nil {
		c.facetCache.camera = camera.New(c)
	}
	return c.facetCache.camera
}

// Events returns the events/client-data sub-API, creating it on first call.
func (c *Sim) Events() *events.Events {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.events == nil {
		c.facetCache.events = events.New(c)
	}
	return c.facetCache.events
}

// Comm returns the CommBus sub-API, creating it on first call.
func (c *Sim) Comm() *comm.Comm {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.comm == nil {
		c.facetCache.comm = comm.New(c)
	}
	return c.facetCache.comm
}

// System returns the system-state/event/action sub-API, creating it on first call.
func (c *Sim) System() *system.System {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.sys == nil {
		c.facetCache.sys = system.New(c)
	}
	return c.facetCache.sys
}

// Flight returns the flight-file/plan sub-API, creating it on first call.
func (c *Sim) Flight() *flight.Flight {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.flight == nil {
		c.facetCache.flight = flight.New(c)
	}
	return c.facetCache.flight
}

// Facilities returns the facilities sub-API, creating it on first call.
func (c *Sim) Facilities() *facilities.Facilities {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.facilities == nil {
		c.facetCache.facilities = facilities.New(c)
	}
	return c.facetCache.facilities
}

// Input returns the input-event sub-API, creating it on first call.
func (c *Sim) Input() *input.Input {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.input == nil {
		c.facetCache.input = input.New(c)
	}
	return c.facetCache.input
}

// Debug returns the diagnostic sub-API, creating it on first call.
func (c *Sim) Debug() *dbgpkg.Debug {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.debug == nil {
		c.facetCache.debug = dbgpkg.New(c)
	}
	return c.facetCache.debug
}

// Simvar returns the typed sim-variable sub-API, creating it on first call.
func (c *Sim) Simvar() *simvar.SimVar {
	c.facetMu.Lock()
	defer c.facetMu.Unlock()
	if c.facetCache.simvar == nil {
		c.facetCache.simvar = simvar.New(c)
	}
	return c.facetCache.simvar
}
