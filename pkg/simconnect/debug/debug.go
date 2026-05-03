//go:build windows

// Package debug implements the Debug SimConnect API category.
package debug

import "github.com/Zwergpro/simconnect-go/pkg/bindings"

// Session is the subset of client.Client methods used by this package.
type Session interface {
	Bindings() *bindings.SimConnect
}

// Debug exposes diagnostic functions for timing and packet tracking.
type Debug struct {
	c Session
}

// New creates a Debug client wrapping the core SimConnect session.
func New(c Session) *Debug {
	return &Debug{c: c}
}

// LastSentPacketID returns the packet ID of the most recently sent SimConnect call.
func (s *Debug) LastSentPacketID() (uint32, error) {
	return s.c.Bindings().GetLastSentPacketID()
}

// RequestResponseTimes returns response-time measurements for the last count packets.
func (s *Debug) RequestResponseTimes(count uint32) ([]float32, error) {
	return s.c.Bindings().RequestResponseTimes(count)
}
