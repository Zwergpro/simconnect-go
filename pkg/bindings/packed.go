//go:build windows

package bindings

import (
	"encoding/binary"
	"math"
)

// SIMCONNECT_PACKED_FLOAT64 stores a C double that may be unaligned because
// SimConnect structs are declared under #pragma pack(1).
type SIMCONNECT_PACKED_FLOAT64 [8]byte

func (v SIMCONNECT_PACKED_FLOAT64) Float64() float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(v[:]))
}

func (v *SIMCONNECT_PACKED_FLOAT64) SetFloat64(f float64) {
	binary.LittleEndian.PutUint64(v[:], math.Float64bits(f))
}

// SIMCONNECT_PACKED_UINT64 stores an unaligned 64-bit integer.
type SIMCONNECT_PACKED_UINT64 [8]byte

func (v SIMCONNECT_PACKED_UINT64) Uint64() uint64 {
	return binary.LittleEndian.Uint64(v[:])
}

func (v *SIMCONNECT_PACKED_UINT64) SetUint64(n uint64) {
	binary.LittleEndian.PutUint64(v[:], n)
}

// SIMCONNECT_PACKED_DATA_LATLONALT is the packed representation of
// SIMCONNECT_DATA_LATLONALT when it appears at an unaligned offset.
type SIMCONNECT_PACKED_DATA_LATLONALT struct {
	Latitude  SIMCONNECT_PACKED_FLOAT64
	Longitude SIMCONNECT_PACKED_FLOAT64
	Altitude  SIMCONNECT_PACKED_FLOAT64
}

// SIMCONNECT_PACKED_DATA_XYZ is the packed representation of
// SIMCONNECT_DATA_XYZ when it appears in packed SimConnect records.
type SIMCONNECT_PACKED_DATA_XYZ struct {
	X SIMCONNECT_PACKED_FLOAT64
	Y SIMCONNECT_PACKED_FLOAT64
	Z SIMCONNECT_PACKED_FLOAT64
}
