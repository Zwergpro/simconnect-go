//go:build windows

package bindings

import (
	"strings"
	"testing"
	"unsafe"
)

func TestPackedLayoutFacilityRecords(t *testing.T) {
	checkSize(t, "SIMCONNECT_DATA_FACILITY_AIRPORT", unsafe.Sizeof(SIMCONNECT_DATA_FACILITY_AIRPORT{}), 36)
	checkOffset(t, "SIMCONNECT_DATA_FACILITY_AIRPORT.Latitude", unsafe.Offsetof(SIMCONNECT_DATA_FACILITY_AIRPORT{}.Latitude), 12)

	checkSize(t, "SIMCONNECT_DATA_FACILITY_WAYPOINT", unsafe.Sizeof(SIMCONNECT_DATA_FACILITY_WAYPOINT{}), 40)
	checkOffset(t, "SIMCONNECT_DATA_FACILITY_WAYPOINT.FMagVar", unsafe.Offsetof(SIMCONNECT_DATA_FACILITY_WAYPOINT{}.FMagVar), 36)

	checkSize(t, "SIMCONNECT_DATA_FACILITY_NDB", unsafe.Sizeof(SIMCONNECT_DATA_FACILITY_NDB{}), 44)
	checkOffset(t, "SIMCONNECT_DATA_FACILITY_NDB.FFrequency", unsafe.Offsetof(SIMCONNECT_DATA_FACILITY_NDB{}.FFrequency), 40)

	checkSize(t, "SIMCONNECT_DATA_FACILITY_VOR", unsafe.Sizeof(SIMCONNECT_DATA_FACILITY_VOR{}), 80)
	checkOffset(t, "SIMCONNECT_DATA_FACILITY_VOR.Flags", unsafe.Offsetof(SIMCONNECT_DATA_FACILITY_VOR{}.Flags), 44)
	checkOffset(t, "SIMCONNECT_DATA_FACILITY_VOR.GlideLat", unsafe.Offsetof(SIMCONNECT_DATA_FACILITY_VOR{}.GlideLat), 52)
	checkOffset(t, "SIMCONNECT_DATA_FACILITY_VOR.FGlideSlopeAngle", unsafe.Offsetof(SIMCONNECT_DATA_FACILITY_VOR{}.FGlideSlopeAngle), 76)

	checkSize(t, "SIMCONNECT_DATA_WAYPOINT", unsafe.Sizeof(SIMCONNECT_DATA_WAYPOINT{}), 44)
	checkOffset(t, "SIMCONNECT_DATA_WAYPOINT.Flags", unsafe.Offsetof(SIMCONNECT_DATA_WAYPOINT{}.Flags), 24)
	checkOffset(t, "SIMCONNECT_DATA_WAYPOINT.KtsSpeed", unsafe.Offsetof(SIMCONNECT_DATA_WAYPOINT{}.KtsSpeed), 28)
	checkOffset(t, "SIMCONNECT_DATA_WAYPOINT.PercentThrottle", unsafe.Offsetof(SIMCONNECT_DATA_WAYPOINT{}.PercentThrottle), 36)
}

func TestPackedLayoutUnalignedRecords(t *testing.T) {
	checkSize(t, "SIMCONNECT_RECV_ASSIGNED_OBJECT_ID", unsafe.Sizeof(SIMCONNECT_RECV_ASSIGNED_OBJECT_ID{}), 20)
	checkOffset(t, "SIMCONNECT_RECV_ASSIGNED_OBJECT_ID.DwObjectID", unsafe.Offsetof(SIMCONNECT_RECV_ASSIGNED_OBJECT_ID{}.DwObjectID), 16)

	checkSize(t, "SIMCONNECT_DATA_INITPOSITION", unsafe.Sizeof(SIMCONNECT_DATA_INITPOSITION{}), 56)
	checkOffset(t, "SIMCONNECT_DATA_INITPOSITION.OnGround", unsafe.Offsetof(SIMCONNECT_DATA_INITPOSITION{}.OnGround), 48)
	checkOffset(t, "SIMCONNECT_DATA_INITPOSITION.Airspeed", unsafe.Offsetof(SIMCONNECT_DATA_INITPOSITION{}.Airspeed), 52)

	checkSize(t, "SIMCONNECT_DATA_RACE_RESULT", unsafe.Sizeof(SIMCONNECT_DATA_RACE_RESULT{}), 1080)
	checkOffset(t, "SIMCONNECT_DATA_RACE_RESULT.FTotalTime", unsafe.Offsetof(SIMCONNECT_DATA_RACE_RESULT{}.FTotalTime), 1060)
	checkSize(t, "SIMCONNECT_RECV_EVENT_RACE_END", unsafe.Sizeof(SIMCONNECT_RECV_EVENT_RACE_END{}), 1108)
	checkOffset(t, "SIMCONNECT_RECV_EVENT_RACE_END.RacerData", unsafe.Offsetof(SIMCONNECT_RECV_EVENT_RACE_END{}.RacerData), 28)

	checkSize(t, "SIMCONNECT_FACILITY_MINIMAL", unsafe.Sizeof(SIMCONNECT_FACILITY_MINIMAL{}), 42)
	checkOffset(t, "SIMCONNECT_FACILITY_MINIMAL.Lla", unsafe.Offsetof(SIMCONNECT_FACILITY_MINIMAL{}.Lla), 18)

	checkSize(t, "SIMCONNECT_JETWAY_DATA", unsafe.Sizeof(SIMCONNECT_JETWAY_DATA{}), 160)
	checkOffset(t, "SIMCONNECT_JETWAY_DATA.Lla", unsafe.Offsetof(SIMCONNECT_JETWAY_DATA{}.Lla), 12)

	checkSize(t, "SIMCONNECT_INPUT_EVENT_DESCRIPTOR", unsafe.Sizeof(SIMCONNECT_INPUT_EVENT_DESCRIPTOR{}), 76)
	checkOffset(t, "SIMCONNECT_INPUT_EVENT_DESCRIPTOR.Hash", unsafe.Offsetof(SIMCONNECT_INPUT_EVENT_DESCRIPTOR{}.Hash), 64)
	checkOffset(t, "SIMCONNECT_INPUT_EVENT_DESCRIPTOR.EType", unsafe.Offsetof(SIMCONNECT_INPUT_EVENT_DESCRIPTOR{}.EType), 72)

	checkOffset(t, "SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT.Hash", unsafe.Offsetof(SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT{}.Hash), 12)
	checkOffset(t, "SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT.EType", unsafe.Offsetof(SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT{}.EType), 20)
	checkOffset(t, "SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT.Value", unsafe.Offsetof(SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT{}.Value), 24)

	checkOffset(t, "SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS.Hash", unsafe.Offsetof(SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS{}.Hash), 12)
	checkOffset(t, "SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS.Value", unsafe.Offsetof(SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS{}.Value), 20)

	checkSize(t, "SIMCONNECT_DATA_CAMERA", unsafe.Sizeof(SIMCONNECT_DATA_CAMERA{}), 84)
	checkOffset(t, "SIMCONNECT_DATA_CAMERA.Fov", unsafe.Offsetof(SIMCONNECT_DATA_CAMERA{}.Fov), 76)
}

func TestCStringBytes(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want string
	}{
		{name: "without terminator", in: []byte("Title"), want: "Title"},
		{name: "with terminator", in: []byte{'T', 'i', 't', 'l', 'e', 0}, want: "Title"},
		{name: "empty string", in: []byte{0}, want: ""},
		{name: "nil", in: nil, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cStringBytes(tt.in); got != tt.want {
				t.Fatalf("cStringBytes(%v) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestInsertStringRejectsEmptyDestination(t *testing.T) {
	end, size, err := InsertString(nil, "Title")
	if err == nil {
		t.Fatal("InsertString(nil, ...) error = nil, want error")
	}
	if !strings.Contains(err.Error(), "empty destination") {
		t.Fatalf("InsertString(nil, ...) error = %q, want empty destination", err)
	}
	if end != nil || size != 0 {
		t.Fatalf("InsertString(nil, ...) end=%v size=%d, want nil/0", end, size)
	}
}

func checkSize(t *testing.T, name string, got, want uintptr) {
	t.Helper()
	if got != want {
		t.Fatalf("%s size = %d, want %d", name, got, want)
	}
}

func checkOffset(t *testing.T, name string, got, want uintptr) {
	t.Helper()
	if got != want {
		t.Fatalf("%s offset = %d, want %d", name, got, want)
	}
}
