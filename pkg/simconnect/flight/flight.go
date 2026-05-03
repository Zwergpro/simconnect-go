//go:build windows

// Package flight implements the Flights SimConnect API category:
// flight file save / load and flight-plan load.
package flight

import "github.com/Zwergpro/simconnect-go/pkg/bindings"

// Session is the subset of client.Sim methods used by this package.
type Session interface {
	Bindings() *bindings.SimConnect
}

// Flight exposes flight file management functions.
type Flight struct {
	session Session
}

// New creates a Flight wrapping the core SimConnect session.
func New(s Session) *Flight {
	return &Flight{session: s}
}

type SaveOption func(*saveConfig)

type saveConfig struct {
	title       *string
	description string
	flags       uint32
}

func WithTitle(title string) SaveOption {
	return func(cfg *saveConfig) {
		cfg.title = &title
	}
}

func WithDescription(description string) SaveOption {
	return func(cfg *saveConfig) {
		cfg.description = description
	}
}

func WithSaveFlags(flags uint32) SaveOption {
	return func(cfg *saveConfig) {
		cfg.flags = flags
	}
}

// Load loads an existing saved flight file.
func (f *Flight) Load(fileName string) error {
	return f.session.Bindings().FlightLoad(fileName)
}

// Save saves the current state of the flight to a flight file.
func (f *Flight) Save(fileName string, opts ...SaveOption) error {
	cfg := saveConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}
	return f.session.Bindings().FlightSaveWithOptionalTitle(fileName, cfg.title, cfg.description, cfg.flags)
}

// PlanLoad loads an existing flight plan file.
func (f *Flight) PlanLoad(fileName string) error {
	return f.session.Bindings().FlightPlanLoad(fileName)
}
