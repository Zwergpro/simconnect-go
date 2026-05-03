//go:build windows

package general

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestPackActionParams(t *testing.T) {
	data := PackActionParams(ActionBool(true), ActionFloat32(1.5), ActionString256("ABC"))
	if len(data) != 1+4+256 {
		t.Fatalf("packed size = %d", len(data))
	}
	if data[0] != 1 {
		t.Fatalf("bool byte = %d", data[0])
	}
	if got := math.Float32frombits(binary.LittleEndian.Uint32(data[1:5])); got != 1.5 {
		t.Fatalf("float = %f", got)
	}
	if string(data[5:8]) != "ABC" || data[8] != 0 {
		t.Fatalf("string bytes = %q next=%d", string(data[5:8]), data[8])
	}
}
