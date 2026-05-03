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

	sim, err := simconnect.Open(ctx, "simconnect-go monitor", simconnect.WithPollInterval(5*time.Second))
	if err != nil {
		log.Fatalf("connect to MSFS 2024 SimConnect: %v", err)
	}
	defer func() {
		if err := sim.Close(); err != nil {
			log.Printf("close SimConnect: %v", err)
		}
	}()

	go func() {
		for err := range sim.Errors() {
			log.Printf("simconnect error: %v", err)
		}
	}()
	positionDef, err := simvar.Define[aircraftPosition]()
	if err != nil {
		log.Fatalf("define aircraft position data: %v", err)
	}

	fmt.Println("Connected to MSFS 2024. Printing user aircraft position every 10 seconds.")
	printPosition(ctx, sim, positionDef)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Monitor stopped.")
			return
		case <-ticker.C:
			printPosition(ctx, sim, positionDef)
		}
	}
}

func printPosition(ctx context.Context, sim *simconnect.Sim, def *simvar.Definition[aircraftPosition]) {
	requestCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	position, err := simvar.GetOnce(requestCtx, sim, def, core.UserAircraft)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		log.Printf("request aircraft position: %v", err)
		return
	}

	fmt.Printf(
		"%s lat=%.6f lon=%.6f alt=%.1f ft\n",
		time.Now().Format(time.RFC3339),
		position.Latitude,
		position.Longitude,
		position.Altitude,
	)
}
