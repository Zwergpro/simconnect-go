//go:build windows

package eventsdata

import (
	"encoding/binary"
	"math"
	"testing"

	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

func TestDefineFromTagsAndDecodeData(t *testing.T) {
	type AircraftState struct {
		Latitude float64 `sim:"PLANE LATITUDE,degree"`
		Airspeed float32 `sim:"AIRSPEED INDICATED,knots"`
		Gear     int32   `sim:"GEAR HANDLE POSITION,bool"`
	}

	def, err := Define[AircraftState]()
	if err != nil {
		t.Fatal(err)
	}
	if len(def.core.fields) != 3 {
		t.Fatalf("fields = %d, want 3", len(def.core.fields))
	}
	if def.core.fields[0].Name != "PLANE LATITUDE" || def.core.fields[0].Units != "degree" {
		t.Fatalf("first field = %+v", def.core.fields[0])
	}

	payload := make([]byte, 16)
	binary.LittleEndian.PutUint64(payload[0:8], math.Float64bits(47.5))
	binary.LittleEndian.PutUint32(payload[8:12], math.Float32bits(132.25))
	binary.LittleEndian.PutUint32(payload[12:16], uint32(1))

	got, err := decodeData[AircraftState](def.core, payload)
	if err != nil {
		t.Fatal(err)
	}
	if got.Latitude != 47.5 || got.Airspeed != 132.25 || got.Gear != 1 {
		t.Fatalf("decoded = %+v", got)
	}
}

func TestDefineFieldsScalar(t *testing.T) {
	def, err := DefineFields[float64](Field{Name: "PLANE ALTITUDE", Units: "feet"})
	if err != nil {
		t.Fatal(err)
	}
	payload := make([]byte, 8)
	binary.LittleEndian.PutUint64(payload, math.Float64bits(12000))
	got, err := decodeData[float64](def.core, payload)
	if err != nil {
		t.Fatal(err)
	}
	if got != 12000 {
		t.Fatalf("decoded scalar = %f", got)
	}
}

func TestDefineStringFieldDecodeAndEncode(t *testing.T) {
	type AircraftInfo struct {
		Title string `sim:"TITLE,string"`
	}

	def, err := Define[AircraftInfo]()
	if err != nil {
		t.Fatal(err)
	}
	if def.core.fields[0].Type != core.DataTypeString256 {
		t.Fatalf("string datatype = %d", def.core.fields[0].Type)
	}

	payload := make([]byte, 256)
	copy(payload, "Cessna 172")
	got, err := decodeData[AircraftInfo](def.core, payload)
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "Cessna 172" {
		t.Fatalf("decoded title = %q", got.Title)
	}

	encoded, err := encodeData(def.core, AircraftInfo{Title: "Bonanza"})
	if err != nil {
		t.Fatal(err)
	}
	if len(encoded) != 256 || fixedString(encoded) != "Bonanza" {
		t.Fatalf("encoded string len=%d value=%q", len(encoded), fixedString(encoded))
	}
}

func TestDefineBoolFieldDecodeAndEncode(t *testing.T) {
	type AircraftState struct {
		OnGround bool `sim:"SIM ON GROUND,bool"`
	}

	def, err := Define[AircraftState]()
	if err != nil {
		t.Fatal(err)
	}
	if def.core.fields[0].Type != core.DataTypeInt32 {
		t.Fatalf("bool datatype = %d", def.core.fields[0].Type)
	}

	payload := make([]byte, 4)
	binary.LittleEndian.PutUint32(payload, 1)
	got, err := decodeData[AircraftState](def.core, payload)
	if err != nil {
		t.Fatal(err)
	}
	if !got.OnGround {
		t.Fatalf("decoded bool = false")
	}

	encoded, err := encodeData(def.core, AircraftState{OnGround: true})
	if err != nil {
		t.Fatal(err)
	}
	if binary.LittleEndian.Uint32(encoded) != 1 {
		t.Fatalf("encoded bool = %d", binary.LittleEndian.Uint32(encoded))
	}
}
