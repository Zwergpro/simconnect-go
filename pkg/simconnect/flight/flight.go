//go:build windows

// Package flight contains flight file and flight plan APIs.
package flight

import (
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/flights"
)

type Flight struct {
	*flights.Flights
}

// Session is the interface required by New; re-exported from flights.
type Session = flights.Session

type SaveOption = flights.FlightSaveOption

func New(c Session) *Flight { return &Flight{Flights: flights.New(c)} }

func WithTitle(title string) SaveOption { return flights.WithFlightTitle(title) }
func WithDescription(description string) SaveOption {
	return flights.WithFlightDescription(description)
}
func WithSaveFlags(flags uint32) SaveOption { return flights.WithFlightSaveFlags(flags) }

func (f *Flight) Load(fileName string) error { return f.FlightLoad(fileName) }
func (f *Flight) Save(fileName string, opts ...SaveOption) error {
	return f.FlightSave(fileName, opts...)
}
func (f *Flight) PlanLoad(fileName string) error { return f.FlightPlanLoad(fileName) }
