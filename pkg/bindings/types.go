//go:build windows

package bindings

// Windows primitive types used by SimConnect
type (
	DWORD   = uint32
	HANDLE  = uintptr
	HWND    = uintptr
	BOOL    = int32
	HRESULT = int32
	UINT64  = uint64
	BYTE    = byte
)

const MAX_PATH = 260

// GUID matches the Windows GUID / UUID layout
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}
