//go:build windows

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/debug"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/system"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "maximum time to run the Debug API checker")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	log.SetFlags(0)

	fmt.Println("SimConnect Debug API checker")
	fmt.Printf("started: %s\n", time.Now().Format(time.RFC3339))

	client, err := simconnect.Open(ctx, "simconnect-go debug checker", simconnect.WithPollInterval(25*time.Millisecond))
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
	dbg := client.Debug()
	gen := client.System()

	checkLastSentPacketID(ctx, dbg, gen)
	checkResponseTimes(dbg)
}

func checkLastSentPacketID(ctx context.Context, dbg *debug.Debug, gen *system.System) {
	fmt.Println("\nSimConnect_GetLastSentPacketID")

	before, err := dbg.LastSentPacketID()
	if err != nil {
		fmt.Printf("  before request: %v\n", err)
	} else {
		fmt.Printf("  before request: %d\n", before)
	}

	requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	state, err := gen.RequestSystemState(requestCtx, "Sim")
	cancel()
	if err != nil {
		fmt.Printf("  request Sim state: %s\n", formatError(err))
	} else {
		fmt.Printf("  request Sim state: int=%d float=%.3f string=%q\n", state.Integer, state.Float, state.String)
	}

	after, err := dbg.LastSentPacketID()
	if err != nil {
		fmt.Printf("  after request: %v\n", err)
		return
	}
	fmt.Printf("  after request:  %d\n", after)
}

func checkResponseTimes(dbg *debug.Debug) {
	fmt.Println("\nSimConnect_RequestResponseTimes")

	times, err := dbg.RequestResponseTimes(5)
	if err != nil {
		fmt.Printf("  request 5 timings: %s\n", formatError(err))
		return
	}
	labels := []string{
		"round trip",
		"request to client send",
		"request to server receive",
		"request to server response",
		"server response to client receive",
	}
	for i, elapsed := range times {
		label := fmt.Sprintf("timing[%d]", i)
		if i < len(labels) {
			label = labels[i]
		}
		fmt.Printf("  %-34s %.6fs\n", label+":", elapsed)
	}

	short, err := dbg.RequestResponseTimes(3)
	if err != nil {
		fmt.Printf("  request 3 timings: %s\n", formatError(err))
		return
	}
	fmt.Printf("  first 3 only: %s\n", formatFloat32s(short))
}

func formatFloat32s(values []float32) string {
	if len(values) == 0 {
		return "none"
	}
	parts := make([]string, len(values))
	for i, value := range values {
		parts[i] = fmt.Sprintf("%.6fs", value)
	}
	return strings.Join(parts, ", ")
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
