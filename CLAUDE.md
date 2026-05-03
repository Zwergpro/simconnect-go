# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Go bindings for Microsoft Flight Simulator's **SimConnect SDK** (`SimConnect.dll`).
The SDK is Windows-only and ships as a C++ header (`sdk/SimConnect.h`) plus DLL/import-lib (`sdk/SimConnect.dll`, `sdk/SimConnect.lib`).

## Build & Verify

Every `.go` file in `pkg/` carries `//go:build windows`. On Linux/macOS hosts, **always** cross-compile:

```bash
GOOS=windows GOARCH=amd64 go build ./pkg/...
GOOS=windows GOARCH=amd64 go vet  ./pkg/...
GOOS=windows GOARCH=amd64 go test ./...
```

A bare `go build ./...` on Linux silently produces nothing because the build tag excludes the whole package — that's not a passing build, it's an empty one. Use the `GOOS=windows` form to actually exercise the code.

Tests run on any host via cross-compile; they don't need MSFS:
- `pkg/bindings/layout_test.go` — `unsafe.Sizeof` / `unsafe.Offsetof` assertions guarding C struct layout. Run this whenever you touch `structs.go` or `packed.go`.
- `pkg/simconnect/client/{client,public_api}_test.go` — exercises the dispatch loop, ID allocators, and facet wiring.

Runtime exercise (anything in `cmd/`) requires a Windows host with MSFS 2024 running, plus `SimConnect.dll` discoverable at startup — easiest is to copy `sdk/SimConnect.dll` next to the built `.exe`.

For the full architectural tour with end-to-end traces, read [docs/OVERVIEW.md](docs/OVERVIEW.md). This file is the cliff-notes; OVERVIEW is the manual.

## Architecture

The codebase is two layers stacked on the SDK:

- **`pkg/bindings/`** — raw, hand-written `syscall` shim over `SimConnect.dll`. One Go function per native API. No CGo, no `.lib` linkage. Names mirror the SDK verbatim.
- **`pkg/simconnect/`** — idiomatic Go API on top: context-cancellable session, ID allocators, channel-based request/response correlation, reflection-based SimVar definitions, per-domain facets.

`cmd/<area>/main.go` examples each demonstrate one facet; [cmd/monitor](cmd/monitor/main.go) is the simplest "open → define → request → close" recipe and the right starting point.

### Layer 2 — `pkg/simconnect/`

`pkg/simconnect/client.Sim` is the spine. `client.Open(ctx, appName, opts...)` connects and (unless `WithManualDispatch()`) starts a goroutine that ticks `cfg.pollInterval`, calls `GetNextDispatch`, decodes packets, and routes them through three maps:

- `waiters[reqID]` — one-shot requests (e.g. `simvar.GetOnce`, `system.RequestState`).
- `dataSubs[reqID]` — recurring subscriptions (e.g. `SIMCONNECT_PERIOD_SIM_FRAME`).
- `handlers[recvID]` — type-keyed listeners (`RECV_ID_QUIT`, `RECV_ID_EXCEPTION`, …).

Errors with no specific request fan out to `sim.Errors() <-chan error`.

Domain functionality lives in subpackages (`ai`, `camera`, `comm`, `debug`, `events`, `facilities`, `flight`, `input`, `simvar`, `system`), each accessed lazily through a cached accessor on the parent: `sim.AI()`, `sim.Camera()`, `sim.Simvar()`, etc. They share the parent's ID allocators so request/definition/event IDs are globally unique across the session. Direct `events.New(sim)` etc. construction stays available as an escape hatch for tests.

`simvar.Define[T]()` reflects struct tags into a SimConnect data definition:

```go
type pos struct {
    Lat float64 `sim:"PLANE LATITUDE,degrees"`
    Lon float64 `sim:"PLANE LONGITUDE,degrees"`
    OnGround bool `sim:"SIM ON GROUND,bool"` // bool → INT32 under the hood
}
```

`NAME` and `units` are taken verbatim from the MSFS SDK SimVar list. For shapes reflection can't express (fixed-width strings, unexported fields), use `simvar.DefineFields[T](fields ...Field)`.

**Dispatch handlers must not block.** They run inline on the polling goroutine. Long work belongs on a worker goroutine fed by a buffered channel.

### Layer 1 — No CGo, pure `syscall`
Bindings are implemented with `syscall.NewLazyDLL("SimConnect.dll")` + `proc.Call(...)`. There is no `import "C"`, no compiler toolchain dependency, no `.lib` linkage. The DLL is resolved at runtime.

Consequences worth knowing before you edit:
- Every argument crosses the boundary as a `uintptr`. Helpers in `pkg/bindings/simconnect.go` handle non-trivial conversions: `cstr` (string → null-terminated `*byte`), `f32` (float32 reinterpret-cast), `f64` (float64 reinterpret-cast to uint64 bits).
- `uint64` parameters (e.g. input-event `Hash`) are passed as a single `uintptr(hash)` — on amd64, `uintptr` is 64-bit, so no hi/lo split is needed.
- The `DispatchProc` callback is wrapped via `syscall.NewCallback`. Its context is a plain `uintptr` to avoid the `unsafe.Pointer(uintptr)` `go vet` warning that would otherwise fire on every dispatch.
- All API functions return `HRESULT`; wrappers funnel through `hresultErr` which yields `nil` on `S_OK` (0) and an `HResultError` otherwise.

### Why bindings are hand-written, not generated
`simconnect.yml` is a leftover from a c-for-go attempt that does **not** work and should not be re-run. The SDK header is C++ (enum base types like `enum Foo:int`, public-inheritance structs, `static const` class members, default function parameters); c-for-go's libclang front-end fails on the very first line (`typedef DWORD SIMCONNECT_OBJECT_ID;`) because Windows types aren't defined. Don't re-attempt generation — extend the hand-written bindings instead.

### File layout in `pkg/bindings/`
The raw bindings package is split by *kind of declaration*, not by SDK feature. When adding to bindings, append to the matching file:

| File | Holds |
|---|---|
| `types.go` | Windows primitive aliases (`DWORD`, `HANDLE`, `HWND`, `BOOL`, `HRESULT`, `UINT64`, `BYTE`) and the `GUID` struct |
| `constants.go` | Every `static const DWORD` from the header, with sign-flipped `^uint32(N)` for the negative special values (e.g. `SIMCONNECT_UNUSED`) |
| `enums.go` | All `SIMCONNECT_ENUM` types as Go typed-`uint32` (or `int32`) constants, plus the `SIMCONNECT_USER_ENUM` and `SIMCONNECT_ENUM_FLAGS` typedefs (which are all `= uint32` aliases) |
| `structs.go` | All `SIMCONNECT_RECV_*` and `SIMCONNECT_DATA_*` structs |
| `packed.go` | `SIMCONNECT_PACKED_FLOAT64`/`UINT64` and `SIMCONNECT_PACKED_DATA_LATLONALT`/`XYZ` wrappers. Their `.Float64()` / `.SetFloat64()` accessors use `binary.LittleEndian` + `math.Float64frombits` so `#pragma pack(1)` fields are read safely on misaligned offsets — **always use these accessors** rather than peeking at the underlying byte array. |
| `simconnect.go` | DLL handle, all `proc*` `*LazyProc` vars, the `SimConnect` session type, helpers, and every function wrapper |
| `layout_test.go` | `unsafe.Sizeof` / `unsafe.Offsetof` assertions. Run after any change in `structs.go` / `packed.go`. |

### Struct conventions (matters for layout, not just style)
The C structs use `#pragma pack(1)` and the DLL hands back raw memory that we re-interpret via these Go types. Field order and types are load-bearing.

- C++ inheritance → Go embedding. `SIMCONNECT_RECV_EXCEPTION : public SIMCONNECT_RECV` becomes a Go struct that embeds `SIMCONNECT_RECV` as its first field.
- C++ `static const` members inside structs are class-level constants, **not** instance fields — they are omitted entirely from the Go layout.
- Variable-length tail fields (`SIMCONNECT_DATAV`, `SIMCONNECT_STRINGV`, `SIMCONNECT_FIXEDTYPE_DATAV`) are represented as `[1]T`. Callers take `&s.Field[0]` and re-slice using the count carried in a sibling field (`dwArraySize`, `dwDefineCount`, etc.).
- Field names are Go-exported PascalCase but mirror the C field order one-for-one. Don't reorder fields for "Go style"; the struct is overlaid onto C memory.

### Naming
Type and constant names are kept verbatim from the SDK (`SIMCONNECT_RECV_EVENT_FRAME`, `SIMCONNECT_PERIOD_SIM_FRAME`, etc.) so they map 1:1 to MSFS SDK documentation. Resist the urge to Go-ify them — the docs are the only meaningful reference for what each value does.
