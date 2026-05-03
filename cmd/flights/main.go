//go:build windows

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/ai"
	simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/debug"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/flight"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/system"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 20*time.Second, "maximum time to run the Flights API checker")
	saveDir := flag.String("save-dir", filepath.Join(os.TempDir(), "simconnect-go-flights"), "directory for the saved flight probe")
	loadSaved := flag.Bool("load-saved", false, "load the just-saved flight after saving it; this changes the active simulator session")
	loadCurrentPlan := flag.Bool("load-current-plan", false, "load the current flight plan path reported by SimConnect; this can change the active simulator session")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	log.SetFlags(0)

	fmt.Println("SimConnect Flights API checker")
	fmt.Printf("started: %s\n", time.Now().Format(time.RFC3339))

	client, err := simconnect.Open(ctx, "simconnect-go flights checker", simconnect.WithPollInterval(25*time.Millisecond))
	if err != nil {
		log.Fatalf("SimConnect_Open: %v", err)
	}
	fmt.Println("\nSimConnect_Open")
	fmt.Println("  connected")
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("\nSimConnect_Close: %v\n", err)
			return
		}
		fmt.Println("\nSimConnect_Close")
		fmt.Println("  closed")
	}()

	go func() {
		for err := range client.Errors() {
			fmt.Printf("  async error: %s\n", formatError(err))
		}
	}()
	aiClient := client.AI()
	dbg := client.Debug()
	flt := client.Flight()
	gen := client.System()

	currentFlight := requestSystemState(ctx, gen, "FlightLoaded")
	currentPlan := requestSystemState(ctx, gen, "FlightPlan")
	printCurrentFiles(currentFlight, currentPlan)

	savedPath := checkFlightSave(ctx, flt, *saveDir)
	checkFlightLoad(flt, savedPath, *loadSaved)
	checkFlightPlanLoad(flt, currentPlan.String, *loadCurrentPlan)
	checkAISetAircraftFlightPlan(aiClient, dbg, currentPlan.String)
}

func printCurrentFiles(currentFlight, currentPlan core.SystemStateMessage) {
	fmt.Println("\nCurrent Flight Files")
	fmt.Printf("  FlightLoaded: int=%d float=%.3f path=%q\n", currentFlight.Integer, currentFlight.Float, currentFlight.String)
	fmt.Printf("  FlightPlan:   int=%d float=%.3f path=%q\n", currentPlan.Integer, currentPlan.Float, currentPlan.String)
}

func checkFlightSave(ctx context.Context, flt *flight.Flight, saveDir string) string {
	fmt.Println("\nSimConnect_FlightSave")

	if err := os.MkdirAll(saveDir, 0755); err != nil {
		fmt.Printf("  create save dir %q: %v\n", saveDir, err)
		return ""
	}

	base := filepath.Join(saveDir, "simconnect-go-flight-check-"+time.Now().Format("20060102-150405"))
	title := "simconnect-go flight save checker"
	description := "Saved by cmd/flights to verify SimConnect_FlightSave."

	if err := flt.Save(
		base,
		flight.WithTitle(title),
		flight.WithDescription(description),
	); err != nil {
		fmt.Printf("  save %q: %s\n", base, formatError(err))
		return ""
	}
	fmt.Printf("  save requested path=%q title=%q\n", base, title)

	fltPath, ok := waitForSavedFlight(ctx, base, 5*time.Second)
	if !ok {
		fmt.Println("  saved .FLT file not observed within 5s")
		return base
	}

	info, err := os.Stat(fltPath)
	if err != nil {
		fmt.Printf("  stat saved file %q: %v\n", fltPath, err)
		return base
	}
	fmt.Printf("  saved file: %q size=%d modified=%s\n", fltPath, info.Size(), info.ModTime().Format(time.RFC3339))

	if preview, err := previewFile(fltPath, 8); err == nil {
		fmt.Println("  preview:")
		for _, line := range preview {
			fmt.Printf("    %s\n", line)
		}
	}

	return base
}

func checkFlightLoad(flt *flight.Flight, savedPath string, enabled bool) {
	fmt.Println("\nSimConnect_FlightLoad")

	if savedPath == "" {
		fmt.Println("  skipped: no saved flight path available")
		return
	}
	if !enabled {
		fmt.Printf("  skipped: pass -load-saved to load %q\n", savedPath)
		return
	}

	if err := flt.Load(savedPath); err != nil {
		fmt.Printf("  load %q: %s\n", savedPath, formatError(err))
		return
	}
	fmt.Printf("  load requested path=%q\n", savedPath)
}

func checkFlightPlanLoad(flt *flight.Flight, currentPlanPath string, enabled bool) {
	fmt.Println("\nSimConnect_FlightPlanLoad")

	if strings.TrimSpace(currentPlanPath) == "" {
		fmt.Println("  skipped: SimConnect reported no active flight plan")
		return
	}
	if !enabled {
		fmt.Printf("  skipped: pass -load-current-plan to load %q\n", currentPlanPath)
		return
	}

	if err := flt.PlanLoad(currentPlanPath); err != nil {
		fmt.Printf("  load plan %q: %s\n", currentPlanPath, formatError(err))
		return
	}
	fmt.Printf("  load plan requested path=%q\n", currentPlanPath)
}

func checkAISetAircraftFlightPlan(aiClient *ai.AI, dbg *debug.Debug, currentPlanPath string) {
	fmt.Println("\nSimConnect_AISetAircraftFlightPlan")

	if strings.TrimSpace(currentPlanPath) == "" {
		fmt.Println("  skipped: no active flight plan to pass to an AI aircraft")
		return
	}
	fmt.Println("  skipped: requires an AI aircraft object ID created or owned by this client")

	if id, err := dbg.LastSentPacketID(); err == nil {
		fmt.Printf("  last sent packet id: %d\n", id)
	}
}

func requestSystemState(ctx context.Context, gen *system.System, name string) core.SystemStateMessage {
	requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	msg, err := gen.RequestSystemState(requestCtx, name)
	if err != nil {
		fmt.Printf("  request system state %q: %s\n", name, formatError(err))
		return core.SystemStateMessage{}
	}
	return msg
}

func waitForSavedFlight(ctx context.Context, base string, timeout time.Duration) (string, bool) {
	candidates := []string{base, base + ".FLT", base + ".flt"}
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		for _, candidate := range candidates {
			if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
				return candidate, true
			}
		}

		select {
		case <-ctx.Done():
			return "", false
		case <-deadline.C:
			return "", false
		case <-ticker.C:
		}
	}
}

func previewFile(path string, lines int) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	all := strings.Split(string(data), "\n")
	if len(all) > lines {
		all = all[:lines]
	}
	for i := range all {
		all[i] = strings.TrimRight(all[i], "\r")
	}
	return all, nil
}

func formatError(err error) string {
	var ex core.ExceptionError
	if errors.As(err, &ex) {
		return fmt.Sprintf("%s send_id=%d index=%d", exceptionName(ex.Exception), ex.SendID, ex.Index)
	}
	return err.Error()
}

func exceptionName(exception core.Exception) string {
	return exception.String()
}
