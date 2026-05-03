//go:build windows

// Package debug implements the Debug SimConnect API category.
package debug

import "github.com/Zwergpro/simconnect-go/pkg/bindings"

// Session is the subset of client.Sim methods used by this package.
type Session interface {
	Bindings() *bindings.SimConnect
}

// Debug exposes diagnostic functions for timing and packet tracking.
type Debug struct {
	session Session
}

// New creates a Debug client wrapping the core SimConnect session.
func New(s Session) *Debug {
	return &Debug{session: s}
}

// LastSentPacketID returns the packet ID of the most recently sent SimConnect call.
func (d *Debug) LastSentPacketID() (uint32, error) {
	return d.session.Bindings().GetLastSentPacketID()
}

// RequestResponseTimes returns response-time measurements for the last count packets.
func (d *Debug) RequestResponseTimes(count uint32) ([]float32, error) {
	return d.session.Bindings().RequestResponseTimes(count)
}
