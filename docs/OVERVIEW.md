# SimConnect-Go: Detailed Overview

A Go binding for the **Microsoft Flight Simulator 2024 SimConnect SDK**. This document explains how SimConnect works conceptually, what the two layers of this library do, and how to use it end-to-end.

Module path: `github.com/Zwergpro/simconnect-go`.

---

## 1. What is SimConnect?

SimConnect is the official client API that Microsoft Flight Simulator (2020 / 2024) ships for third-party applications: cockpit panels, ATC tools, hardware bridges, telemetry dashboards, mission/scenario scripts, multi-monitor camera rigs. The SDK artifacts in this repo:

- `sdk/SimConnect.h` — C/C++ header (~1700 lines, MSFS 2024 schema), kept as reference material
- `sdk/SimConnect.lib` — static import library (unused by this project; we load the DLL dynamically)
- `pkg/bindings/SimConnect.dll` — runtime library that talks to the simulator; embedded into the binary via `//go:embed` and extracted at runtime

**Connection model.** A client process loads `SimConnect.dll`, calls `SimConnect_Open`, and is given an opaque `HANDLE`. Behind the scenes the DLL connects to the running simulator over a named pipe (local) or TCP (remote). The simulator is the server; multiple clients can attach simultaneously.

**Four pillars of the protocol.** Every SimConnect interaction belongs to one of these:

1. **Request data** — read SimVars (e.g. `PLANE LATITUDE`) from the user aircraft, an AI object, or a typed "client data" area. Either one-shot or recurring (every visual frame, sim frame, second, etc.).
2. **Set data** — write back to writable SimVars or to client data areas.
3. **Transmit / receive events** — fire client events that map to sim controls (`PARKING_BRAKES`, `THROTTLE_FULL`, …) or receive events the sim raises (gear up, crash, frame).
4. **Subscribe to system state** — flight loaded, paused, sim running, view changed, weather changed, AI radar contact, etc.

**Asynchronous, ID-keyed.** Requests don't return values directly. The client posts a request with a numeric ID, then pumps a queue of typed response packets — `SIMCONNECT_RECV_OPEN`, `SIMCONNECT_RECV_EVENT`, `SIMCONNECT_RECV_SIMOBJECT_DATA`, `SIMCONNECT_RECV_EXCEPTION`, etc. The pump is either:

- `SimConnect_GetNextDispatch` — poll one packet, returns `E_FAIL` when the queue is empty.
- `SimConnect_CallDispatch` — drain the queue into a callback (`DispatchProc`).

Every packet starts with a `SIMCONNECT_RECV` header containing `dwSize`, `dwVersion`, and `dwID`. The client switches on `dwID` and casts to the matching `SIMCONNECT_RECV_*` struct.

---

## 2. SDK surface

The official documentation under [sdk/apidocs/](../sdk/apidocs/) groups the ~120 native functions into these families:

| Family | Representative functions | apidocs folder |
|---|---|---|
| General / session | `SimConnect_Open`, `SimConnect_Close`, `SimConnect_GetNextDispatch`, `SimConnect_CallDispatch`, `SimConnect_RequestSystemState`, `SimConnect_ExecuteAction` | `General/` |
| Events & data | `SimConnect_AddToDataDefinition`, `SimConnect_RequestDataOnSimObject`, `SimConnect_SetDataOnSimObject`, `SimConnect_MapClientEventToSimEvent`, `SimConnect_TransmitClientEvent_EX1`, `SimConnect_SubscribeToSystemEvent`, plus client data and flow events | `Events_And_Data/` |
| AI objects | `SimConnect_AICreateParkedATCAircraft`, `SimConnect_AICreateEnrouteATCAircraft`, `SimConnect_AICreateNonATCAircraft`, `SimConnect_AICreateSimulatedObject`, `SimConnect_AIRemoveObject`, `SimConnect_AIReleaseControl`, `SimConnect_AISetAircraftFlightPlan` | `AI_Object/` |
| Facilities | `SimConnect_RequestFacilitiesList_EX1`, `SimConnect_SubscribeToFacilities_EX1`, `SimConnect_RequestFacilityData_EX1`, `SimConnect_RequestJetwayData` | `Facilities/` |
| Input events | `SimConnect_MapInputEventToClientEvent_EX1`, `SimConnect_EnumerateInputEvents`, `SimConnect_GetInputEvent`, `SimConnect_SetInputEvent`, `SimConnect_SubscribeInputEvent`, `SimConnect_EnumerateControllers` | `InputEvents/` |
| Camera | `SimConnect_CameraAcquire`, `SimConnect_CameraRelease`, `SimConnect_CameraSet`, `SimConnect_CameraGet`, `SimConnect_EnumerateCameraDefinitions`, world-locker calls | `Camera/` |
| Communication (CommBus) | `SimConnect_CallCommBusEvent`, `SimConnect_SubscribeToCommBusEvent`, `SimConnect_UnsubscribeToCommBusEvent` | `Communication/` |
| Flights | `SimConnect_FlightLoad`, `SimConnect_FlightSave`, `SimConnect_FlightPlanLoad` | `Flights/` |
| Debug | `SimConnect_GetLastSentPacketID`, `SimConnect_RequestResponseTimes` | `Debug/` |

**Actions catalogue.** [sdk/apidocs/SimConnect_Actions.md](../sdk/apidocs/SimConnect_Actions.md) lists the high-level XML actions invokable through `SimConnect_ExecuteAction` — audio (`PlaySoundAction`, `PlayMusicAction`), UI (`DialogAction`, `ShowLogbookAction`), camera (`RequestTeleportCameraAction`, `FocusInstrumentAction`), gameplay (`AITakeControlsAction`, `RefillAction`, `FailureAction`), Wwise audio, and instructor notifications. The upstream document is marked WIP — treat the action list as best-effort.

**Key enums to know.** They are defined in `sdk/SimConnect.h` and re-exported in `pkg/bindings/enums.go`:

- `SIMCONNECT_RECV_ID` — ~54 message types, the discriminator for dispatched packets.
- `SIMCONNECT_DATATYPE` — `INT32/64`, `FLOAT32/64`, fixed-width strings (`STRING8` … `STRING260`), `STRINGV`, `INITPOSITION`, `LATLONALT`, `XYZ`.
- `SIMCONNECT_PERIOD` — `NEVER`, `ONCE`, `VISUAL_FRAME`, `SIM_FRAME`, `SECOND` for recurring data requests.
- `SIMCONNECT_DATA_REQUEST_FLAG` — `CHANGED`, `TAGGED` modifiers.
- `SIMCONNECT_EXCEPTION` — ~42 numeric exception codes returned via `SIMCONNECT_RECV_EXCEPTION`.
- `SIMCONNECT_SIMOBJECT_TYPE` — `USER`, `ALL`, `AIRCRAFT`, `HELICOPTER`, `BOAT`, `GROUND`, etc.
- `SIMCONNECT_FACILITY_DATA_TYPE` — `AIRPORT`, `RUNWAY`, `APPROACH`, `VOR`, `NDB`, `WAYPOINT`, `TAXI_PARKING`, `JETWAY`, …

---

## 3. Layer 1 — `pkg/bindings` (raw syscall wrapper)

The bindings package is a thin, hand-maintained, **CGo-free** shim over `SimConnect.dll`. The DLL ships embedded in the package (`//go:embed SimConnect.dll` in [pkg/bindings/dll_embed.go](../pkg/bindings/dll_embed.go)) and is extracted to `%TEMP%\simconnect-go-<sha256_16>.dll` on first use; `syscall.NewLazyDLL` loads from that path and every entry point is invoked via `proc.Call(...)`. There is no C compiler dependency. Callers can override the DLL path with `client.WithDLLPath(path)` / `bindings.SetDLLPath(path)` or the `SIMCONNECT_DLL` env var — paths that don't exist transparently fall back to the embedded copy.

Every file in `pkg/bindings/` carries `//go:build windows`.

### File map

| File | Holds |
|---|---|
| [pkg/bindings/types.go](../pkg/bindings/types.go) | Windows primitive aliases (`DWORD`, `HANDLE`, `HWND`, `BOOL`, `HRESULT`, `UINT64`, `BYTE`), the `GUID` struct, `MAX_PATH`. |
| [pkg/bindings/constants.go](../pkg/bindings/constants.go) | Every `static const DWORD` from the header (~50 entries: object IDs, priorities, request/event/data flags, comm-bus targets), with sign-flipped `^uint32(N)` for negative special values such as `SIMCONNECT_UNUSED`. |
| [pkg/bindings/enums.go](../pkg/bindings/enums.go) | All `SIMCONNECT_ENUM` types as Go typed-`uint32` (or `int32`) constants — ~20 enums covering `RECV_ID`, `DATATYPE`, `EXCEPTION`, `PERIOD`, `SIMOBJECT_TYPE`, facility/camera/weather/flow/input families. |
| [pkg/bindings/structs.go](../pkg/bindings/structs.go) | All `SIMCONNECT_RECV_*` and `SIMCONNECT_DATA_*` structs (~50 types). C++ inheritance becomes Go embedding; `#pragma pack(1)` is preserved by field order. |
| [pkg/bindings/packed.go](../pkg/bindings/packed.go) | `SIMCONNECT_PACKED_FLOAT64`, `SIMCONNECT_PACKED_UINT64`, plus `SIMCONNECT_PACKED_DATA_LATLONALT` and `SIMCONNECT_PACKED_DATA_XYZ` wrappers. Methods like `.Float64()` / `.SetFloat64()` use `binary.LittleEndian` + `math.Float64frombits` for safe access on misaligned offsets. |
| [pkg/bindings/simconnect.go](../pkg/bindings/simconnect.go) | DLL handle, ~120 `*LazyProc` vars, the `SimConnect` session type, marshalling helpers, and 116 method wrappers — one per native API. |
| [pkg/bindings/layout_test.go](../pkg/bindings/layout_test.go) | `unsafe.Sizeof` / `unsafe.Offsetof` assertions guarding struct layout against accidental reordering. |

### The `SimConnect` session type

```go
type SimConnect struct {
    handle HANDLE
}

func Open(name string, hWnd HWND, userEventWin32 uint32, hEventHandle HANDLE, configIndex uint32) (*SimConnect, error)
func (sc *SimConnect) Close() error
func (sc *SimConnect) GetNextDispatch() (*SIMCONNECT_RECV, uint32, error)
func (sc *SimConnect) CallDispatch(pfn DispatchProc, pContext uintptr) error
func (sc *SimConnect) GetLastSentPacketID() (uint32, error)
```

`GetNextDispatch` returns `(nil, 0, nil)` when the queue is empty (the underlying `E_FAIL` is folded into a "no packet" signal). `CallDispatch` accepts a `DispatchProc` (`func(*SIMCONNECT_RECV, uint32, uintptr)`) wrapped via `syscall.NewCallback`; the context parameter is intentionally a plain `uintptr` to dodge the `unsafe.Pointer(uintptr)` `go vet` warning that would otherwise fire on every dispatch.

### CGo-free conversion helpers

Each call crosses the syscall boundary as a list of `uintptr` arguments. The non-trivial conversions live in [pkg/bindings/simconnect.go](../pkg/bindings/simconnect.go):

| Helper | Purpose |
|---|---|
| `cstr(s string) *byte` | Encodes a Go string to a NUL-terminated byte slice and pins it through the call so the GC can't move it. |
| `f32(v float32) uintptr` / `f64(v float64) uintptr` | Reinterpret-cast float bits into `uintptr`; on amd64 a `uintptr` is 64 bits, so `float64` arguments fit in a single register slot. |
| `hresultErr(r1 uintptr) error` | Returns `nil` for `S_OK == 0`, else an `HResultError{HRESULT: …}` with hex formatting. |
| `RetrieveString` / `InsertString` | Read and write the variable-length strings that some SimConnect calls return inside packed packet bodies. |
| `SIMCONNECT_PACKED_*` accessors | Safely read/write 64-bit values inside `#pragma pack(1)` structs without unaligned-load undefined behaviour. |

`uint64` arguments such as the input-event `Hash` are passed as a single `uintptr(hash)` — no hi/lo split is needed on amd64.

### Why bindings are hand-written, not generated

`simconnect.yml` (a c-for-go config left over from an early experiment) **does not work and should not be re-run**. The SDK header is C++ — `enum Foo : int`, public-inheritance structs, `static const` class members, default function parameters — and c-for-go's libclang front-end fails on the very first line (`typedef DWORD SIMCONNECT_OBJECT_ID;`) because Windows types aren't visible. Extend the hand-written bindings instead.

### Struct layout rules

- C++ inheritance → Go embedding. `SIMCONNECT_RECV_EXCEPTION : public SIMCONNECT_RECV` becomes a Go struct that embeds `SIMCONNECT_RECV` as its **first** field.
- C++ `static const` members are class-level constants, not instance fields — they are omitted from the Go layout.
- Variable-length tail fields (`SIMCONNECT_DATAV`, `SIMCONNECT_STRINGV`, `SIMCONNECT_FIXEDTYPE_DATAV`) are declared as `[1]T`. Callers take `&s.Field[0]` and re-slice using the count carried in a sibling field (`dwArraySize`, `dwDefineCount`, …).
- Field names are Go-exported PascalCase but mirror the C field order one-for-one. **Don't reorder fields for "Go style"** — the struct is overlaid onto C memory, and `layout_test.go` will catch a mistake on the next test run.

---

## 4. Layer 2 — `pkg/simconnect` (high-level client)

Layer 2 wraps `pkg/bindings` in an idiomatic Go API: a context-cancellable session, ID allocators, channel-based request/response correlation, typed data definitions via Go reflection, and per-domain subpackages.

### `pkg/simconnect/client` — the spine

Defined in [pkg/simconnect/client/client.go](../pkg/simconnect/client/client.go). The primary type is `Sim`:

```go
type Sim struct {
    raw *bindings.SimConnect
    cfg clientConfig

    ctx    context.Context
    cancel context.CancelFunc
    done   chan struct{}

    closeOnce sync.Once
    closeErr  error

    errs chan error

    reqIDs   *idAllocator
    defIDs   *idAllocator
    eventIDs *idAllocator

    mu         sync.Mutex
    closed     bool
    waiters    map[uint32]chan RequestResult       // one-shot requests
    sendToReq  map[uint32]uint32                   // sent-packet → request ID, for exception correlation
    dataSubs   map[uint32][]func(Message)          // continuous subscriptions
    handlers   map[RecvID][]func(Message)          // type-keyed callbacks
    closeHooks []func()

    facetMu    sync.Mutex // guards lazily-built facet accessors
    facetCache facets
}

func Open(ctx context.Context, appName string, opts ...Option) (*Sim, error)
func (c *Sim) Close() error
func (c *Sim) Errors() <-chan error
func (c *Sim) Poll() error
```

`Open` calls `bindings.Open`, builds the dispatch state, and (unless `WithManualDispatch()` is passed) starts a background goroutine that ticks at `cfg.pollInterval` and drains every queued packet via `Poll → GetNextDispatch → decodeMessage → dispatch`. `Close` is idempotent (`sync.Once`): it cancels the context, waits for the dispatch goroutine, fans `ErrClosed` to all outstanding waiters, runs registered close hooks, calls the underlying `SimConnect_Close`, and closes the error channel.

### `pkg/simconnect/core` — shared protocol types

`core` is the single home for the protocol types (`ObjectID`, `Period`, `RecvID`, `Exception`, …), the `Message` interface and every concrete `*Message` decoded packet type, the typed errors (`HResultError`, `ErrClosed`, `ExceptionError`, `RequestResult`), and the small `bindings_conv.go` adapters. It exists so that facet packages (`ai`, `camera`, `events`, …) can share types with `client` without forming an import cycle: `client` imports the facets (for the `sim.AI()`-style accessors); the facets only import `core`.

These names are **not** re-exported from `client`. Code that needs a protocol value imports `core` directly: `core.UserAircraft`, `core.PeriodSecond`, `core.ExceptionError`, `core.OpenMessage`, etc. `client` exports only what it owns — the session itself (`Sim`), constructors (`Open`, `Dial`), options (`WithPollInterval`, `WithManualDispatch`), and the dispatch primitives (`AddWaiter`, `RegisterHandler`, …).

### Domain subpackages (`pkg/simconnect/<area>`)

Each subpackage wraps a `*client.Sim` and exposes a typed surface for one SDK family. There are two layers here:

- **Facet packages** named after the SDK area users actually think about. These are what callers normally import.
- **Implementation packages** grouped by `apidocs/` family — `eventsdata`, `general`. The facet packages are thin re-exports over these, so the heavy lifting lives in one place per apidocs group instead of being duplicated. Calling code can ignore the implementation packages unless it needs an API the facet hasn't lifted yet.

| Facet | Wraps | Surface |
|---|---|---|
| [`ai`](../pkg/simconnect/ai/) | (self) | Spawn ATC / non-ATC / parked / simulated objects; release control; remove; upload flight plans. |
| [`camera`](../pkg/simconnect/camera/) | (self) | Acquire / release; set & get camera; enumerate definitions; status and world-locker subscriptions via channels. |
| [`comm`](../pkg/simconnect/comm/) | (self) | CommBus subscribe / call (inter-client and panel-to-app messaging). |
| [`debug`](../pkg/simconnect/debug/) | (self) | `LastSentPacketID`, `RequestResponseTimes` for diagnosing request/response timing. |
| [`events`](../pkg/simconnect/events/) | [`eventsdata`](../pkg/simconnect/eventsdata/) | Client events, notification groups, client-data areas, flow events. |
| [`simvar`](../pkg/simconnect/simvar/) | [`eventsdata`](../pkg/simconnect/eventsdata/) | Typed SimVar definitions, one-shot reads, writes, and subscriptions. |
| [`facilities`](../pkg/simconnect/facilities/) | (self) | Airport / waypoint / NDB / VOR list and detail queries with paginated aggregation. |
| [`flight`](../pkg/simconnect/flight/) | (self) | Flight file save / load; flight-plan load. |
| [`input`](../pkg/simconnect/input/) | (self) | Input event enumeration, get / set / subscribe (hash-keyed dispatch). |
| [`system`](../pkg/simconnect/system/) | [`general`](../pkg/simconnect/general/) | System state queries, `ExecuteAction` with packed parameters, system-event subscription, notification-group management. |

[`client`](../pkg/simconnect/client/) is the core session itself — dispatch loop, ID allocation, handler registry, error channel, message decoding, and the `Option` builders. Protocol values like `UserAircraft` and `ObjectID` live in `core`, not here; import `core` directly when you need them.

Domain packages share the parent client's ID allocators, so request, definition, and event IDs are globally unique across the session.

`*client.Sim` exposes each facet as a cached accessor: `sim.AI()`, `sim.Camera()`, `sim.Events()`, `sim.Comm()`, `sim.System()`, `sim.Flight()`, `sim.Facilities()`, `sim.Input()`, `sim.Debug()`, `sim.Simvar()`. Each is constructed once on first call (guarded by `facetMu`) and reused for the lifetime of the session. Direct construction (`events.New(sim)`, `camera.New(sim)`, …) remains available as an escape hatch for tests and advanced composition.

---

## 5. How it all fits — the dispatch pump

Tracing `simvar.GetOnce(ctx, sim, def, core.UserAircraft)` end-to-end:

```
┌──────────────────┐  Define[T]() reflects struct tags into a list of (NAME, units, type) fields
│ user code        │  simvar.GetOnce(...) is the entry point
└────────┬─────────┘
         ▼
┌──────────────────┐  • allocates definitionID (defIDs.Next) the first time the def is used
│ simvar          │  • allocates requestID  (reqIDs.Next)
│  (uses Sim)     │  • registers waiter:  ch, _ := client.AddWaiter(requestID)
└────────┬─────────┘  • calls bindings: client.raw.AddToDataDefinition(...) per field
         │                         then client.raw.RequestDataOnSimObject(reqID, defID, objectID, ONCE, …)
         ▼
┌──────────────────┐  syscall through SimConnect.dll → MSFS server
│ pkg/bindings     │
└────────┬─────────┘
         ▼
┌──────────────────┐  ticks every cfg.pollInterval
│ Sim.run()        │  → Poll() → raw.GetNextDispatch() → SIMCONNECT_RECV_SIMOBJECT_DATA
│  goroutine       │  → decodeMessage(raw, size) → dispatch(msg)
└────────┬─────────┘
         ▼
┌──────────────────┐  matches msg.DispatchRequestID() against c.waiters[reqID]
│ dispatch.go      │  sends RequestResult{Msg: msg} down the channel
└────────┬─────────┘
         ▼
┌──────────────────┐  reads from ch in GetOnce, copies the packet body
│ simvar           │  into a fresh T using the same field reflection, returns (T, error)
└──────────────────┘
```

The same wiring serves three patterns:

- **One-shot** (`waiters` map) — `simvar.GetOnce`, `system.RequestState`.
- **Continuous** (`dataSubs` map) — recurring `SIMCONNECT_PERIOD_SIM_FRAME` data requests, input-event subscriptions.
- **Type-keyed** (`handlers` map) — generic listeners for `RECV_ID_QUIT`, `RECV_ID_EXCEPTION`, etc.

Errors that fire outside any specific request are pushed to `client.Errors() <-chan error`.

---

## 6. Quick start

This is the [cmd/monitor/main.go](../cmd/monitor/main.go) example — the canonical "open, define, request, loop, close" recipe.

```go
//go:build windows

package main

import (
    "context"
    "errors"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
    "github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
    "github.com/Zwergpro/simconnect-go/pkg/simconnect/simvar"
)

type aircraftPosition struct {
    Latitude  float64 `sim:"PLANE LATITUDE,degrees"`
    Longitude float64 `sim:"PLANE LONGITUDE,degrees"`
    Altitude  float64 `sim:"PLANE ALTITUDE,feet"`
}

func main() {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    sim, err := simconnect.Open(ctx, "simconnect-go monitor",
        simconnect.WithPollInterval(5*time.Second))
    if err != nil {
        log.Fatalf("connect to MSFS 2024 SimConnect: %v", err)
    }
    defer sim.Close()

    // Drain async errors in the background.
    go func() {
        for err := range sim.Errors() {
            log.Printf("simconnect error: %v", err)
        }
    }()

    // Reflect aircraftPosition into a SimConnect data definition.
    positionDef, err := simvar.Define[aircraftPosition]()
    if err != nil {
        log.Fatalf("define aircraft position data: %v", err)
    }

    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
            pos, err := simvar.GetOnce(reqCtx, sim, positionDef, core.UserAircraft)
            cancel()
            if err != nil {
                if !errors.Is(err, context.Canceled) {
                    log.Printf("request: %v", err)
                }
                continue
            }
            fmt.Printf("lat=%.6f lon=%.6f alt=%.1f ft\n",
                pos.Latitude, pos.Longitude, pos.Altitude)
        }
    }
}
```

The usual recipe is **`Open` → `simvar.Define`/configure → `simvar.GetOnce` for sim-variable reads, or `sim.Events()` / `sim.Camera()` / `sim.System()` / etc. for facet-specific operations → `Close`.** Each `sim.X()` accessor lazily constructs and caches the facet on first call; direct `events.New(sim)` / `camera.New(sim)` construction remains available as a low-level escape hatch for tests and advanced composition.

### `sim:"…"` field tag format

`simvar.Define[T]()` walks the exported fields of `T` and reads `sim:"NAME,units"`:

- `NAME` — the SimVar name (`PLANE LATITUDE`, `INDICATED ALTITUDE`, `THROTTLE LOWER LIMIT`). Names are taken verbatim from the MSFS SDK SimVar list.
- `units` — measurement units exactly as the SDK expects (`degrees`, `feet`, `knots`, `radians`, …).

Field types must match the SimConnect datatype; `float64` is the safe default for numeric SimVars.

If you need a definition you can't express with reflection, for example fixed-width strings or non-exported field names, use `simvar.DefineFields[T](fields ...Field)` instead. `bool` fields are encoded as SimConnect `INT32`, so tags such as `sim:"SIM ON GROUND,bool"` can use Go `bool`.

---

## 7. Building and running

The whole `pkg/` tree is `//go:build windows`. On a non-Windows host, **always cross-compile**:

```bash
GOOS=windows GOARCH=amd64 go build ./pkg/...
GOOS=windows GOARCH=amd64 go vet  ./pkg/...
GOOS=windows GOARCH=amd64 go test ./...
```

A bare `go build ./...` on Linux/macOS *silently produces nothing* because the build tag excludes every file in the package. That isn't a passing build — it's an empty one. Use the `GOOS=windows` form to actually exercise the code.

To build an example:

```bash
GOOS=windows GOARCH=amd64 go build -o monitor.exe ./cmd/monitor
```

The resulting binary is self-contained: `SimConnect.dll` is embedded and extracted to `%TEMP%\simconnect-go-<hash>.dll` on first use. Override with `client.WithDLLPath(...)` or `SIMCONNECT_DLL=<path>` to point at a different DLL; missing override paths fall back to the embedded copy. Runtime exercise also needs MSFS 2024 actually running and accepting connections.

There are no Go tests against the simulator; `pkg/bindings/layout_test.go` covers struct layout in isolation and runs anywhere via `GOOS=windows GOARCH=amd64 go test ./pkg/bindings/...`.

---

## 8. Example tour (`cmd/`)

| Directory | Demonstrates |
|---|---|
| [cmd/aiobject](../cmd/aiobject/) | Spawning AI ATC and parked aircraft via `pkg/simconnect/ai`. |
| [cmd/camera](../cmd/camera/) | Camera acquire/set/get and definition enumeration. |
| [cmd/communication](../cmd/communication/) | CommBus subscribe and call round-trip. |
| [cmd/debug](../cmd/debug/) | `LastSentPacketID` and `RequestResponseTimes`. |
| [cmd/eventsdata](../cmd/eventsdata/) | Defining custom data structs, mapping events, streaming SimObject data. |
| [cmd/facilities](../cmd/facilities/) | Airport / waypoint / NDB / VOR list queries with pagination. |
| [cmd/flights](../cmd/flights/) | Flight-file load / save with metadata. |
| [cmd/general](../cmd/general/) | Comprehensive system-state, system-event, action, and notification-group walkthrough — toggles between automatic and manual dispatch. |
| [cmd/inputevents](../cmd/inputevents/) | Enumerate input events, subscribe, get/set parameters. |
| [cmd/monitor](../cmd/monitor/) | Simplest example — prints user aircraft position every 10 s. Start here. |
| [cmd/stats](../cmd/stats/) | Continuous airspeed / altitude / telemetry monitor. |

---

## 9. Pointers and gotchas

- **Field order is load-bearing.** Every `SIMCONNECT_*` struct is overlaid on raw memory the DLL hands back. Don't reorder fields, change types, or insert padding for "Go style". `pkg/bindings/layout_test.go` exists to catch breakage.
- **Use packed accessors.** Any field declared `SIMCONNECT_PACKED_FLOAT64`, `SIMCONNECT_PACKED_UINT64`, `SIMCONNECT_PACKED_DATA_LATLONALT`, or `SIMCONNECT_PACKED_DATA_XYZ` requires the `.Float64()` / `.SetFloat64()` (or equivalent) methods. Reading the underlying byte array directly works on amd64 by accident but is a portability hazard and signals a misuse on review.
- **Names mirror the SDK verbatim.** `SIMCONNECT_RECV_EVENT_FRAME`, `SIMCONNECT_PERIOD_SIM_FRAME`, etc. — the official MSFS SDK documentation and `sdk/apidocs/` are the only meaningful semantic reference. Don't Go-ify them.
- **Don't regenerate `simconnect.yml`.** It's a leftover from an unsuccessful c-for-go attempt — the C++ header trips libclang. Extend the hand-written bindings.
- **`SimConnect_Actions.md` is upstream-WIP.** Treat the action catalogue as best-effort; some actions are documented before they ship, others ship before they're documented.
- **Cross-compile or get nothing.** `go build ./...` on Linux/macOS produces no artifact and no error because of the Windows build tag. Always pass `GOOS=windows GOARCH=amd64`.
- **Dispatch handlers must not block.** `Sim.dispatch` runs every registered handler and waiter callback inline on the polling goroutine. Long work belongs on a worker goroutine fed by a buffered channel.
- **Protocol values live in `core`, not `client`.** `core.UserAircraft`, `core.PeriodSecond`, `core.DataSetDefault`, `core.ExceptionError`, `core.OpenMessage`, every `RecvID*`/`Period*`/`*Flag`/`Exception*` constant, and the `Message` interface itself are imported from `pkg/simconnect/core`. `client` exports only the session (`Sim`, `Open`, `Dial`, options, dispatch primitives). If you see `client.X` for a protocol value in old code, it's a leftover from before re-exports were removed — switch the import to `core`.
- **Don't import the implementation packages directly.** `pkg/simconnect/eventsdata` and `general` are visible in the source tree but are wrapped by `events`/`simvar` and `system` respectively. Reach for the facet name first; only drop down when the facet hasn't lifted what you need.
