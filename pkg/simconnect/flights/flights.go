//go:build windows

// Package flights implements the Flights SimConnect API category.
package flights

import "github.com/Zwergpro/simconnect-go/pkg/bindings"

// Session is the subset of client.Sim methods used by this package.
type Session interface {
	Bindings() *bindings.SimConnect
}

// Flights exposes flight file management functions.
type Flights struct {
	c Session
}

// New creates a Flights wrapping the core SimConnect session.
func New(c Session) *Flights {
	return &Flights{c: c}
}

type FlightSaveOption func(*flightSaveConfig)

type flightSaveConfig struct {
	title       *string
	description string
	flags       uint32
}

func WithFlightTitle(title string) FlightSaveOption {
	return func(cfg *flightSaveConfig) {
		cfg.title = &title
	}
}

func WithFlightDescription(description string) FlightSaveOption {
	return func(cfg *flightSaveConfig) {
		cfg.description = description
	}
}

func WithFlightSaveFlags(flags uint32) FlightSaveOption {
	return func(cfg *flightSaveConfig) {
		cfg.flags = flags
	}
}

// FlightLoad loads an existing saved flight file.
func (s *Flights) FlightLoad(fileName string) error {
	return s.c.Bindings().FlightLoad(fileName)
}

// FlightSave saves the current state of the flight to a flight file.
func (s *Flights) FlightSave(fileName string, opts ...FlightSaveOption) error {
	cfg := flightSaveConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}
	return s.c.Bindings().FlightSaveWithOptionalTitle(fileName, cfg.title, cfg.description, cfg.flags)
}

// FlightPlanLoad loads an existing flight plan file.
func (s *Flights) FlightPlanLoad(fileName string) error {
	return s.c.Bindings().FlightPlanLoad(fileName)
}
