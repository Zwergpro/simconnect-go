# simconnect-go

Go bindings for the **Microsoft Flight Simulator 2024 SimConnect SDK**.

`simconnect-go` is a CGo-free, pure-`syscall` wrapper over `SimConnect.dll`, plus an idiomatic high-level client with context cancellation, typed SimVar definitions via reflection, channel-based request/response correlation, and per-domain facets (AI, camera, events, simvars, facilities, flights, input, system, …).

```
import simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
```

## Features

- **No CGo, no C toolchain.** `SimConnect.dll` is embedded into the binary and loaded at runtime via `syscall.NewLazyDLL` — `go build` produces a self-contained `.exe` that works on any Windows host. Override with `simconnect.WithDLLPath(...)` or the `SIMCONNECT_DLL` env var; missing override paths fall back to the embedded copy.
- **Two layers, your choice.** Use the raw [`pkg/bindings`](pkg/bindings/) shim (one Go function per native API, names mirror the SDK verbatim) or the high-level [`pkg/simconnect`](pkg/simconnect/) client.
- **Typed SimVar definitions.** Tag a Go struct with `sim:"NAME,units"` and `simvar.Define[T]()` reflects it into a SimConnect data definition.
- **Context-cancellable session.** `simconnect.Open(ctx, ...)` runs a background dispatch goroutine; `Close` is idempotent and fans `ErrClosed` to outstanding waiters.
- **Domain facets.** `sim.AI()`, `sim.Camera()`, `sim.Events()`, `sim.Simvar()`, `sim.System()`, `sim.Facilities()`, `sim.Flight()`, `sim.Input()`, `sim.Comm()`, `sim.Debug()` — each lazily constructed and cached on the parent client.
- **Three dispatch patterns.** One-shot waiters (`simvar.GetOnce`, `system.RequestState`), recurring subscriptions (`SIMCONNECT_PERIOD_SIM_FRAME`), and type-keyed listeners (`RECV_ID_QUIT`, `RECV_ID_EXCEPTION`, …).

## Requirements

- Windows host with **Microsoft Flight Simulator 2024** running for runtime exercise.
- No separate DLL deployment needed — the MSFS 2024 SimConnect redistributable is embedded. Use `simconnect.WithDLLPath(path)` or `SIMCONNECT_DLL=path` to load a different DLL (newer SDK, system-managed install). If the override path doesn't exist, the loader falls back to the embedded copy and logs one notice.
- Go 1.26+.

The whole `pkg/` tree carries `//go:build windows`. Tests cross-compile and run on any host.

## Install

```bash
go get github.com/Zwergpro/simconnect-go
```

## Quick start

This is [cmd/monitor/main.go](cmd/monitor/main.go) — the canonical "open → define → request → close" recipe. It prints user-aircraft position every 10 seconds.

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

    go func() {
        for err := range sim.Errors() {
            log.Printf("simconnect error: %v", err)
        }
    }()

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

### `sim:"…"` tag format

`simvar.Define[T]()` walks exported fields of `T` and reads `sim:"NAME,units"`:

- `NAME` — the SimVar name (`PLANE LATITUDE`, `INDICATED ALTITUDE`, …) verbatim from the MSFS SDK SimVar list.
- `units` — measurement units exactly as the SDK expects (`degrees`, `feet`, `knots`, `radians`, `bool`, …).

`bool` fields are encoded as SimConnect `INT32`. For shapes reflection can't express (fixed-width strings, unexported fields), use `simvar.DefineFields[T](fields ...Field)`.

## Architecture

Two layers stacked on the SDK:

- **[`pkg/bindings/`](pkg/bindings/)** — raw, hand-written `syscall` shim over `SimConnect.dll`. One Go function per native API, names mirror the SDK verbatim. No CGo, no `.lib` linkage.
- **[`pkg/simconnect/`](pkg/simconnect/)** — idiomatic Go API on top: context-cancellable session, ID allocators, channel-based request/response correlation, reflection-based SimVar definitions, per-domain facets.

For the full architectural tour with end-to-end traces, read [docs/OVERVIEW.md](docs/OVERVIEW.md).

### Facet → SDK family

| Facet | Surface |
|---|---|
| [`ai`](pkg/simconnect/ai/) | Spawn ATC / non-ATC / parked / simulated objects; release control; remove; upload flight plans. |
| [`camera`](pkg/simconnect/camera/) | Acquire / release; set & get camera; enumerate definitions; status and world-locker subscriptions. |
| [`comm`](pkg/simconnect/comm/) | CommBus subscribe / call (inter-client and panel-to-app messaging). |
| [`debug`](pkg/simconnect/debug/) | `LastSentPacketID`, `RequestResponseTimes` for diagnosing timing. |
| [`events`](pkg/simconnect/events/) | Client events, notification groups, client-data areas, flow events. |
| [`facilities`](pkg/simconnect/facilities/) | Airport / waypoint / NDB / VOR list and detail queries with paginated aggregation. |
| [`flight`](pkg/simconnect/flight/) | Flight file save / load; flight-plan load. |
| [`input`](pkg/simconnect/input/) | Input event enumeration, get / set / subscribe (hash-keyed dispatch). |
| [`simvar`](pkg/simconnect/simvar/) | Typed SimVar definitions, one-shot reads, writes, and subscriptions. |
| [`system`](pkg/simconnect/system/) | System state queries, `ExecuteAction`, system-event subscription, notification-group management. |

## Examples

Each subdirectory under [`cmd/`](cmd/) demonstrates one facet end-to-end.

| Example | Demonstrates |
|---|---|
| [cmd/monitor](cmd/monitor/) | Simplest example — prints user aircraft position every 10 s. **Start here.** |
| [cmd/stats](cmd/stats/) | Continuous airspeed / altitude / telemetry monitor. |
| [cmd/aiobject](cmd/aiobject/) | Spawning AI ATC and parked aircraft. |
| [cmd/camera](cmd/camera/) | Camera acquire / set / get and definition enumeration. |
| [cmd/communication](cmd/communication/) | CommBus subscribe and call round-trip. |
| [cmd/debug](cmd/debug/) | `LastSentPacketID` and `RequestResponseTimes`. |
| [cmd/eventsdata](cmd/eventsdata/) | Custom data structs, mapping events, streaming SimObject data. |
| [cmd/facilities](cmd/facilities/) | Airport / waypoint / NDB / VOR list queries with pagination. |
| [cmd/flights](cmd/flights/) | Flight-file load / save with metadata. |
| [cmd/general](cmd/general/) | Comprehensive system-state, system-event, action, notification-group walkthrough. |
| [cmd/inputevents](cmd/inputevents/) | Enumerate input events, subscribe, get / set parameters. |

## Build

The whole `pkg/` tree is `//go:build windows`. On a non-Windows host, **always cross-compile**:

```bash
GOOS=windows GOARCH=amd64 go build ./pkg/...
GOOS=windows GOARCH=amd64 go vet  ./pkg/...
GOOS=windows GOARCH=amd64 go test ./...
```

A bare `go build ./...` on Linux/macOS *silently produces nothing* because the build tag excludes every file in the package. Use the `GOOS=windows` form to actually exercise the code.

Build an example:

```bash
GOOS=windows GOARCH=amd64 go build -o monitor.exe ./cmd/monitor
```

The resulting `.exe` is self-contained: `SimConnect.dll` is embedded and extracted to `%TEMP%\simconnect-go-<hash>.dll` on first use. To load a different DLL, pass `simconnect.WithDLLPath("C:\\path\\to\\SimConnect.dll")` to `Open` or set `SIMCONNECT_DLL` in the environment; missing paths fall back to the embedded copy. Runtime exercise also needs MSFS 2024 actually running.

## Gotchas

- **Field order in `SIMCONNECT_*` structs is load-bearing.** They're overlaid on raw memory the DLL hands back. Don't reorder fields or insert padding. [`pkg/bindings/layout_test.go`](pkg/bindings/layout_test.go) catches breakage.
- **Use packed accessors.** `SIMCONNECT_PACKED_FLOAT64`, `SIMCONNECT_PACKED_UINT64`, `SIMCONNECT_PACKED_DATA_LATLONALT`, `SIMCONNECT_PACKED_DATA_XYZ` need their `.Float64()` / `.SetFloat64()` methods — reading the underlying bytes directly is a portability hazard.
- **Dispatch handlers must not block.** They run inline on the polling goroutine. Long work belongs on a worker goroutine fed by a buffered channel.
- **Protocol values live in `core`, not `client`.** `core.UserAircraft`, `core.PeriodSecond`, `core.ExceptionError`, every `RecvID*`/`Period*`/`*Flag`/`Exception*` constant, and the `Message` interface itself are in [`pkg/simconnect/core`](pkg/simconnect/core/). `client` exports only the session.
- **Names mirror the SDK verbatim.** `SIMCONNECT_RECV_EVENT_FRAME`, `SIMCONNECT_PERIOD_SIM_FRAME`, etc. — the official MSFS SDK docs are the only meaningful semantic reference.

## Documentation

- [docs/OVERVIEW.md](docs/OVERVIEW.md) — full architectural tour, end-to-end traces, dispatch pump diagram, SDK surface map.
- [sdk/apidocs/](sdk/apidocs/) — vendored upstream MSFS SDK API documentation.

## Status

This library is under active development against the MSFS 2024 SimConnect SDK. The bindings cover the full ~120-function native surface; the high-level client lifts the most common workflows. Issues and PRs are welcome.

## License

See [LICENSE](LICENSE) for details.
