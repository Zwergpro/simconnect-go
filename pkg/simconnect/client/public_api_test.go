//go:build windows

package client_test

import (
	"testing"

	"github.com/Zwergpro/simconnect-go/pkg/simconnect/ai"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/camera"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/comm"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/debug"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/events"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/facilities"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/flight"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/input"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/simvar"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/system"
)

func TestPublicAPIUsesHighLevelTypes(t *testing.T) {
	var sim *client.Sim

	_ = client.WithWindowHandle(uintptr(0))
	_ = client.WithEventID(1)
	_ = client.WithEventHandle(uintptr(0))
	_ = client.WithConfigIndex(0)

	_ = core.ExceptionError{
		Exception: core.ExceptionNameUnrecognized,
		SendID:    1,
		Index:     2,
	}
	_ = core.InitPosition{Latitude: 47.0, Longitude: -122.0, Altitude: 500, OnGround: 1}
	_ = core.XYZ{X: 1, Y: 2, Z: 3}
	_ = core.PBH{Pitch: 1, Bank: 2, Heading: 3}
	_ = core.Airport{Ident: "KSEA", Latitude: 47.4489, Longitude: -122.3094}
	_ = core.CameraData{Position: core.XYZ{}, PBH: core.PBH{}, FOV: 50}
	_ = core.FacilityMinimal{ICAO: core.ICAO{Ident: "KSEA"}}
	_ = core.JetwayData{ParkingIndex: 1}
	_ = core.Version{Major: 1, Minor: 2}

	var msg core.Message
	if msg != nil && msg.RecvID() == core.RecvIDNull {
		t.Fatal("unexpected null message")
	}
	var handler func(core.Message) = func(core.Message) {}
	_ = handler

	// *client.Sim satisfies every facet Session interface (duck typing).
	// These are compile-time assertions: if *Sim ever stops satisfying a
	// Session interface the build fails here rather than at the call site.
	var _ ai.Session = sim
	var _ camera.Session = sim
	var _ comm.Session = sim
	var _ debug.Session = sim
	var _ events.Session = sim
	var _ facilities.Session = sim
	var _ flight.Session = sim
	var _ system.Session = sim
	var _ input.Session = sim
	var _ simvar.Session = sim

	_ = simvar.Field{Name: "PLANE ALTITUDE", Units: "feet", Type: core.DataTypeFloat64}
	_ = facilities.FacilityDataFilter{Path: "RUNWAYS/INDEX", Data: []byte{1, 0, 0, 0}}
	_ = flight.WithTitle("API check")
	_ = system.PackActionParams(system.ActionBool(true), system.ActionFloat32(1.5))

	_ = sim
}
