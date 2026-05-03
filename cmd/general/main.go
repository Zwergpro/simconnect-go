//go:build windows

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/events"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/system"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 20*time.Second, "maximum time to run the General API checker")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	log.SetFlags(0)

	fmt.Println("SimConnect General API checker")
	fmt.Printf("started: %s\n", time.Now().Format(time.RFC3339))

	checkGetNextDispatch(ctx)
	checkCallDispatch(ctx)

	client, err := simconnect.Open(ctx, "simconnect-go general checker", simconnect.WithPollInterval(25*time.Millisecond))
	if err != nil {
		log.Fatalf("SimConnect_Open: %v", err)
	}
	fmt.Println("\nSimConnect_Open")
	fmt.Println("  connected with automatic polling")
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("  SimConnect_Close: %v\n", err)
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
	ed := client.Events()
	gen := client.System()

	checkSystemStates(ctx, gen)
	checkNotificationGroupPriority(ed, gen)
	checkSystemEvents(ctx, gen)
	checkExecuteAction(ctx, gen)
}

func checkGetNextDispatch(ctx context.Context) {
	fmt.Println("\nSimConnect_GetNextDispatch")

	client, err := simconnect.Open(ctx, "simconnect-go get-next-dispatch checker", simconnect.WithManualDispatch())
	if err != nil {
		fmt.Printf("  SimConnect_Open: %v\n", err)
		return
	}
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("  SimConnect_Close: %v\n", err)
		}
	}()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		msg, ok, err := client.GetNextMessage()
		if err != nil {
			fmt.Printf("  get next dispatch: %v\n", err)
			return
		}
		if !ok {
			time.Sleep(25 * time.Millisecond)
			continue
		}
		printDecodedPacket("  packet", msg)
		return
	}

	fmt.Println("  no packet available in 2s")
}

func checkCallDispatch(ctx context.Context) {
	fmt.Println("\nSimConnect_CallDispatch / DispatchProc")

	client, err := simconnect.Open(ctx, "simconnect-go call-dispatch checker", simconnect.WithManualDispatch())
	if err != nil {
		fmt.Printf("  SimConnect_Open: %v\n", err)
		return
	}
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("  SimConnect_Close: %v\n", err)
		}
	}()

	var callbacks int
	var lastID core.RecvID
	callback := func(msg core.Message, pContext uintptr) {
		callbacks++
		lastID = msg.RecvID()
		fmt.Printf("  callback #%d context=%d\n", callbacks, pContext)
		printDecodedPacket("    packet", msg)
	}

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) && callbacks == 0 {
		err := client.CallDispatch(callback, 42)
		if err != nil {
			fmt.Printf("  call dispatch: %v\n", err)
			time.Sleep(25 * time.Millisecond)
			continue
		}
		time.Sleep(25 * time.Millisecond)
	}
	if callbacks == 0 {
		fmt.Println("  no callback in 2s")
		return
	}
	fmt.Printf("  callbacks=%d lastRecvID=%s\n", callbacks, recvIDName(lastID))
}

func checkSystemStates(ctx context.Context, gen *system.System) {
	fmt.Println("\nSimConnect_RequestSystemState")

	for _, state := range []string{"Sim", "DialogMode", "AircraftLoaded", "FlightLoaded", "FlightPlan"} {
		requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
		msg, err := gen.RequestSystemState(requestCtx, state)
		cancel()
		if err != nil {
			fmt.Printf("  %-14s error: %v\n", state, err)
			continue
		}
		fmt.Printf("  %-14s int=%d float=%.3f string=%q\n", state, msg.Integer, msg.Float, msg.String)
	}
}

func checkNotificationGroupPriority(ed *events.Events, gen *system.System) {
	fmt.Println("\nSimConnect_SetNotificationGroupPriority")

	const groupID core.NotificationGroupID = 0
	event, err := ed.MapEvent("")
	if err != nil {
		fmt.Printf("  create private event for group=%d error: %v\n", groupID, err)
		return
	}
	if err := ed.AddClientEventToNotificationGroup(groupID, event, false); err != nil {
		fmt.Printf("  add private event=%d to group=%d error: %v\n", event.ID(), groupID, err)
		return
	}
	defer func() {
		if err := ed.ClearNotificationGroup(groupID); err != nil {
			fmt.Printf("  clear group=%d error: %v\n", groupID, err)
		}
	}()

	if err := gen.SetNotificationGroupPriority(groupID, core.GroupPriorityDefault); err != nil {
		fmt.Printf("  group=%d priority=%d error: %v\n", groupID, core.GroupPriorityDefault, err)
		return
	}
	fmt.Printf("  group=%d event=%d priority=%d ok\n", groupID, event.ID(), core.GroupPriorityDefault)
}

func checkSystemEvents(ctx context.Context, gen *system.System) {
	fmt.Println("\nSystem Events")

	simCtx, simCancel := context.WithCancel(ctx)
	defer simCancel()
	simSub, err := gen.SubscribeSystemEventWithID(simCtx, "Sim")
	if err != nil {
		fmt.Printf("  SimConnect_SubscribeToSystemEvent(\"Sim\"): %v\n", err)
	} else {
		fmt.Printf("  subscribed Sim eventID=%d\n", simSub.EventID)
		printEventOrTimeout(ctx, "  Sim current", simSub.Events, 2*time.Second)
	}

	secCtx, secCancel := context.WithCancel(ctx)
	defer secCancel()
	secSub, err := gen.SubscribeSystemEventWithID(secCtx, "1sec")
	if err != nil {
		fmt.Printf("  SimConnect_SubscribeToSystemEvent(\"1sec\"): %v\n", err)
		return
	}
	fmt.Printf("  subscribed 1sec eventID=%d\n", secSub.EventID)

	if err := gen.SetSystemEventState(secSub.EventID, core.StateOff); err != nil {
		fmt.Printf("  SimConnect_SetSystemEventState(OFF): %v\n", err)
	} else {
		fmt.Println("  1sec state off")
	}
	printEventOrTimeout(ctx, "  1sec while off", secSub.Events, 1200*time.Millisecond)

	if err := gen.SetSystemEventState(secSub.EventID, core.StateOn); err != nil {
		fmt.Printf("  SimConnect_SetSystemEventState(ON): %v\n", err)
	} else {
		fmt.Println("  1sec state on")
	}
	printEventOrTimeout(ctx, "  1sec while on", secSub.Events, 2500*time.Millisecond)

	secCancel()
	fmt.Println("  SimConnect_UnsubscribeFromSystemEvent(\"1sec\") requested by canceling subscription context")
}

func checkExecuteAction(ctx context.Context, gen *system.System) {
	fmt.Println("\nSimConnect_ExecuteAction")

	params := []system.ActionParam{
		system.ActionBool(true),
		system.ActionFloat32(0),
		system.ActionBool(false),
		system.ActionString256(""),
		system.ActionString256(""),
	}

	requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	msg, err := gen.ExecuteAction(requestCtx, "simconnect_go_general_checker_noop", params...)
	cancel()
	if err != nil {
		fmt.Printf("  sent intentionally unknown action with %d packed bytes; response error: %s\n", len(system.PackActionParams(params...)), formatError(err))
		return
	}
	fmt.Printf("  callback requestID=%d actionID=%q\n", msg.RequestID, msg.ActionID)
}

func printEventOrTimeout(ctx context.Context, label string, events <-chan core.ClientEvent, timeout time.Duration) {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case event, ok := <-events:
		if !ok {
			fmt.Printf("%s: channel closed\n", label)
			return
		}
		fmt.Printf("%s: group=%d event=%d data=%d values=%v\n", label, event.GroupID, event.EventID, event.Data, event.DataValues)
	case <-timer.C:
		fmt.Printf("%s: no event in %s\n", label, timeout)
	case <-ctx.Done():
		fmt.Printf("%s: %v\n", label, ctx.Err())
	}
}

func printDecodedPacket(prefix string, msg core.Message) {
	switch m := msg.(type) {
	case core.OpenMessage:
		fmt.Printf("%s %s app=%q appVersion=%d.%d.%d.%d simconnect=%d.%d.%d.%d\n",
			prefix,
			recvIDName(msg.RecvID()),
			m.ApplicationName,
			m.ApplicationVersionMajor,
			m.ApplicationVersionMinor,
			m.ApplicationBuildMajor,
			m.ApplicationBuildMinor,
			m.SimConnectVersionMajor,
			m.SimConnectVersionMinor,
			m.SimConnectBuildMajor,
			m.SimConnectBuildMinor,
		)
	case core.ClientEvent:
		fmt.Printf("%s %s group=%d event=%d data=%d\n", prefix, recvIDName(msg.RecvID()), m.GroupID, m.EventID, m.Data)
	case core.SystemStateMessage:
		fmt.Printf("%s %s request=%d int=%d float=%.3f string=%q\n", prefix, recvIDName(msg.RecvID()), m.RequestID, m.Integer, m.Float, m.String)
	default:
		fmt.Printf("%s %s %T\n", prefix, recvIDName(msg.RecvID()), msg)
	}
}

func recvIDName(id core.RecvID) string {
	return id.String()
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
