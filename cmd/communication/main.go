//go:build windows

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/comm"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 15*time.Second, "maximum time to run the Communication API checker")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	log.SetFlags(0)

	fmt.Println("SimConnect Communication API checker")
	fmt.Printf("started: %s\n", time.Now().Format(time.RFC3339))

	client, err := simconnect.Open(ctx, "simconnect-go communication checker", simconnect.WithPollInterval(25*time.Millisecond))
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
	bus := client.Comm()

	eventName := fmt.Sprintf("simconnect-go.communication.%d", time.Now().UnixNano())
	checkCommBusRoundTrip(ctx, bus, eventName)
	checkRawSubscribeUnsubscribe(client, eventName+".raw")
	checkBroadcastTargets(bus, eventName+".broadcast-only")
}

func checkCommBusRoundTrip(ctx context.Context, bus *comm.Comm, eventName string) {
	fmt.Println("\nCommBus Round Trip")

	publisher, err := simconnect.Open(ctx, "simconnect-go communication publisher", simconnect.WithPollInterval(25*time.Millisecond))
	if err != nil {
		fmt.Printf("  publisher SimConnect_Open: %s\n", formatError(err))
		return
	}
	defer func() {
		if err := publisher.Close(); err != nil {
			fmt.Printf("  publisher SimConnect_Close: %s\n", formatError(err))
		}
	}()
	go func() {
		for err := range publisher.Errors() {
			fmt.Printf("  publisher async error: %s\n", formatError(err))
		}
	}()
	publisherComm := publisher.Comm()
	fmt.Println("  publisher connection opened")

	subCtx, cancel := context.WithCancel(ctx)
	ch, err := bus.SubscribeCommBusEvent(subCtx, eventName)
	if err != nil {
		cancel()
		fmt.Printf("  SimConnect_SubscribeToCommBusEvent(%q): %s\n", eventName, formatError(err))
		return
	}
	fmt.Printf("  SimConnect_SubscribeToCommBusEvent: event=%q ok\n", eventName)

	stringPayload := fmt.Sprintf(`{"source":"cmd/communication","kind":"string","time":%q}`, time.Now().Format(time.RFC3339))
	if err := publisherComm.CallCommBusEvent(eventName, core.CommBusBroadcastToSimConnect, stringPayload); err != nil {
		fmt.Printf("  SimConnect_CallCommBusEvent string: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_CallCommBusEvent string: broadcast=SIMCONNECT bytes=%d ok\n", len(stringPayload)+1)
		readCommBusMessage(ctx, ch, "string")
	}

	bytePayload := []byte{0x7B, 0x22, 0x6B, 0x69, 0x6E, 0x64, 0x22, 0x3A, 0x22, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x7D, 0x00}
	if err := publisherComm.CallCommBusEventBytes(eventName, core.CommBusBroadcastToSimConnect, bytePayload); err != nil {
		fmt.Printf("  SimConnect_CallCommBusEvent bytes: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_CallCommBusEvent bytes: broadcast=SIMCONNECT bytes=%d ok\n", len(bytePayload))
		readCommBusMessage(ctx, ch, "bytes")
	}

	cancel()
	select {
	case _, ok := <-ch:
		if !ok {
			fmt.Println("  SimConnect_UnsubscribeToCommBusEvent: subscription context canceled, channel closed")
		}
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  SimConnect_UnsubscribeToCommBusEvent: subscription context canceled")
	}
}

func readCommBusMessage(ctx context.Context, ch <-chan core.CommBusMessage, label string) {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()

	select {
	case msg, ok := <-ch:
		if !ok {
			fmt.Printf("  receive %s: channel closed\n", label)
			return
		}
		fmt.Printf(
			"  receive %s: eventID=%d packet=%d/%d bytes=%d data=%q raw=%s\n",
			label,
			msg.EventID,
			msg.EntryNumber+1,
			msg.OutOf,
			len(msg.Payload),
			msg.Data,
			formatBytes(msg.Payload),
		)
	case <-timer.C:
		fmt.Printf("  receive %s: no message in 3s\n", label)
	case <-ctx.Done():
		fmt.Printf("  receive %s: %v\n", label, ctx.Err())
	}
}

func checkRawSubscribeUnsubscribe(client *simconnect.Sim, eventName string) {
	fmt.Println("\nRaw Subscribe / Unsubscribe")

	const eventID core.EventID = 9301
	if err := client.SubscribeToCommBusEvent(eventID, eventName); err != nil {
		fmt.Printf("  raw SubscribeToCommBusEvent(eventID=%d, event=%q): %s\n", eventID, eventName, formatError(err))
		return
	}
	fmt.Printf("  raw SubscribeToCommBusEvent: eventID=%d event=%q ok\n", eventID, eventName)

	if err := client.UnsubscribeToCommBusEvent(eventID); err != nil {
		fmt.Printf("  raw UnsubscribeToCommBusEvent(eventID=%d): %s\n", eventID, formatError(err))
		return
	}
	fmt.Printf("  raw UnsubscribeToCommBusEvent: eventID=%d ok\n", eventID)
}

func checkBroadcastTargets(bus *comm.Comm, eventName string) {
	fmt.Println("\nBroadcast Targets")

	targets := []struct {
		name   string
		target core.CommBusBroadcastTo
	}{
		{"SIMCONNECT", core.CommBusBroadcastToSimConnect},
		{"ALL_SIMCONNECT", core.CommBusBroadcastToAllSimConnect},
		{"DEFAULT", core.CommBusBroadcastToDefault},
	}

	for _, target := range targets {
		payload := fmt.Sprintf(`{"target":%q}`, target.name)
		err := bus.CallCommBusEvent(eventName, target.target, payload)
		if err != nil {
			fmt.Printf("  %-15s error: %s\n", target.name, formatError(err))
			continue
		}
		fmt.Printf("  %-15s call ok (no local subscription for this event)\n", target.name)
	}
}

func formatBytes(data []byte) string {
	if len(data) == 0 {
		return "[]"
	}
	limit := len(data)
	if limit > 32 {
		limit = 32
	}
	parts := make([]string, limit)
	for i := 0; i < limit; i++ {
		parts[i] = fmt.Sprintf("%02X", data[i])
	}
	if len(data) > limit {
		return "[" + strings.Join(parts, " ") + " ...]"
	}
	return "[" + strings.Join(parts, " ") + "]"
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
