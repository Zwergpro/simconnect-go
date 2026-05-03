//go:build windows

package core

import "github.com/Zwergpro/simconnect-go/pkg/bindings"

// InitPositionToBinding converts a core.InitPosition to the bindings equivalent.
func InitPositionToBinding(data InitPosition) bindings.SIMCONNECT_DATA_INITPOSITION {
	return bindings.SIMCONNECT_DATA_INITPOSITION{
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Altitude:  data.Altitude,
		Pitch:     data.Pitch,
		Bank:      data.Bank,
		Heading:   data.Heading,
		OnGround:  data.OnGround,
		Airspeed:  data.Airspeed,
	}
}

// XYZToBinding converts a core.XYZ to bindings.SIMCONNECT_DATA_XYZ (plain float64 fields).
func XYZToBinding(data XYZ) bindings.SIMCONNECT_DATA_XYZ {
	return bindings.SIMCONNECT_DATA_XYZ{X: data.X, Y: data.Y, Z: data.Z}
}

// XYZToPackedBinding converts a core.XYZ to bindings.SIMCONNECT_PACKED_DATA_XYZ
// (used for camera position / jetway fields that use packed float64).
func XYZToPackedBinding(data XYZ) bindings.SIMCONNECT_PACKED_DATA_XYZ {
	var raw bindings.SIMCONNECT_PACKED_DATA_XYZ
	raw.X.SetFloat64(data.X)
	raw.Y.SetFloat64(data.Y)
	raw.Z.SetFloat64(data.Z)
	return raw
}

// PBHToBinding converts a core.PBH to the bindings equivalent.
func PBHToBinding(data PBH) bindings.SIMCONNECT_DATA_PBH {
	return bindings.SIMCONNECT_DATA_PBH{Pitch: data.Pitch, Bank: data.Bank, Heading: data.Heading}
}

// CameraDataToBinding converts a core.CameraData to the bindings equivalent.
func CameraDataToBinding(data CameraData) bindings.SIMCONNECT_DATA_CAMERA {
	raw := bindings.SIMCONNECT_DATA_CAMERA{
		Position:                    XYZToPackedBinding(data.Position),
		PositionReferential:         bindings.SIMCONNECT_POSITION_REFERENTIAL(data.PositionReferential),
		PositionReferentialObjectId: uint32(data.PositionReferentialObjectID),
		TargetedPos:                 XYZToPackedBinding(data.TargetedPos),
		Pbh:                         PBHToBinding(data.PBH),
		RotationReferential:         bindings.SIMCONNECT_POSITION_REFERENTIAL(data.RotationReferential),
		RotationReferentialObjectId: uint32(data.RotationReferentialObjectID),
	}
	raw.Fov.SetFloat64(data.FOV)
	return raw
}
