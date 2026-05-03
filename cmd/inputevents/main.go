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
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/input"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 25*time.Second, "maximum time to run the Input Events API checker")
	setSample := flag.Bool("set-sample", false, "write the currently-read sample input event value back with SimConnect_SetInputEvent")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	log.SetFlags(0)

	fmt.Println("SimConnect Input Events API checker")
	fmt.Printf("started: %s\n", time.Now().Format(time.RFC3339))

	client, err := simconnect.Open(ctx, "simconnect-go input events checker", simconnect.WithPollInterval(25*time.Millisecond))
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
	ed := client.Events()
	inp := client.Input()

	checkControllers(ctx, inp)
	events := checkEnumerateInputEvents(ctx, inp)
	sample, ok := chooseSampleInputEvent(events)
	if ok {
		value, hasValue := checkGetInputEvent(ctx, inp, sample)
		checkEnumerateInputEventParams(ctx, inp, sample)
		checkSubscribeInputEvent(ctx, inp, sample)
		checkSetInputEvent(inp, sample, value, hasValue, *setSample)
	} else {
		fmt.Println("\nInput Event Detail")
		fmt.Println("  skipped: no input event with a non-zero hash returned")
	}
	checkInputGroupMapping(client, ed)
}

func checkControllers(ctx context.Context, inp *input.Input) {
	fmt.Println("\nSimConnect_EnumerateControllers")

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	controllers, err := inp.EnumerateControllers(requestCtx)
	cancel()
	if err != nil {
		fmt.Printf("  error: %s\n", formatError(err))
		return
	}

	fmt.Printf("  page=%d/%d entries=%d\n", controllers.EntryNumber+1, controllers.OutOf, len(controllers.Controllers))
	for i, controller := range controllers.Controllers {
		if i >= 8 {
			fmt.Printf("  ... %d more\n", len(controllers.Controllers)-i)
			break
		}
		fmt.Printf(
			"  [%d] name=%q device=%d product=%d composite=%d hw=%d.%d.%d.%d\n",
			i,
			controller.DeviceName,
			controller.DeviceID,
			controller.ProductID,
			controller.CompositeID,
			controller.HardwareVersion.Major,
			controller.HardwareVersion.Minor,
			controller.HardwareVersion.Revision,
			controller.HardwareVersion.Build,
		)
	}
}

func checkEnumerateInputEvents(ctx context.Context, inp *input.Input) core.InputEventListMessage {
	fmt.Println("\nSimConnect_EnumerateInputEvents")

	requestCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
	events, err := inp.EnumerateInputEvents(requestCtx)
	cancel()
	if err != nil {
		fmt.Printf("  error: %s\n", formatError(err))
		return core.InputEventListMessage{}
	}

	fmt.Printf("  page=%d/%d entries=%d\n", events.EntryNumber+1, events.OutOf, len(events.Events))
	for i, event := range events.Events {
		if i >= 12 {
			fmt.Printf("  ... %d more\n", len(events.Events)-i)
			break
		}
		fmt.Printf("  [%d] hash=%d type=%s name=%q\n", i, event.Hash, inputEventTypeName(event.Type), event.Name)
	}
	return events
}

func chooseSampleInputEvent(events core.InputEventListMessage) (core.InputEventDescriptor, bool) {
	for _, event := range events.Events {
		if event.Hash != 0 && event.Type == core.InputEventTypeDouble {
			return event, true
		}
	}
	for _, event := range events.Events {
		if event.Hash != 0 {
			return event, true
		}
	}
	return core.InputEventDescriptor{}, false
}

func checkGetInputEvent(ctx context.Context, inp *input.Input, event core.InputEventDescriptor) (core.InputEventValueMessage, bool) {
	fmt.Println("\nSimConnect_GetInputEvent")
	fmt.Printf("  sample: hash=%d type=%s name=%q\n", event.Hash, inputEventTypeName(event.Type), event.Name)

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	value, err := inp.GetInputEvent(requestCtx, event.Hash)
	cancel()
	if err != nil {
		fmt.Printf("  error: %s\n", formatError(err))
		return core.InputEventValueMessage{}, false
	}
	fmt.Printf("  value: requestID=%d type=%s %s payload=%d\n", value.RequestID, inputEventTypeName(value.Type), formatInputValue(value), len(value.Payload))
	return value, true
}

func checkEnumerateInputEventParams(ctx context.Context, inp *input.Input, event core.InputEventDescriptor) {
	fmt.Println("\nSimConnect_EnumerateInputEventParams")

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	params, err := inp.EnumerateInputEventParams(requestCtx, event.Hash)
	cancel()
	if err != nil {
		fmt.Printf("  hash=%d error: %s\n", event.Hash, formatError(err))
		return
	}
	fmt.Printf("  hash=%d raw=%q params=%d\n", params.Hash, params.Value, len(params.Params))
	for i, param := range params.Params {
		if i >= 8 {
			fmt.Printf("  ... %d more\n", len(params.Params)-i)
			break
		}
		fmt.Printf("  [%d] %s\n", i, param)
	}
}

func checkSubscribeInputEvent(ctx context.Context, inp *input.Input, event core.InputEventDescriptor) {
	fmt.Println("\nSimConnect_SubscribeInputEvent / UnsubscribeInputEvent")

	subCtx, cancel := context.WithCancel(ctx)
	ch, err := inp.SubscribeInputEvent(subCtx, event.Hash)
	if err != nil {
		cancel()
		fmt.Printf("  subscribe hash=%d: %s\n", event.Hash, formatError(err))
		return
	}
	fmt.Printf("  subscribed hash=%d; waiting briefly for a value-change notification\n", event.Hash)

	timer := time.NewTimer(2 * time.Second)
	select {
	case msg, ok := <-ch:
		if !ok {
			fmt.Println("  subscription channel closed")
		} else {
			fmt.Printf("  notification: hash=%d type=%s %s payload=%d\n", msg.Hash, inputEventTypeName(msg.Type), formatSubscriptionValue(msg), len(msg.Payload))
		}
	case <-timer.C:
		fmt.Println("  no notification in 2s")
	case <-ctx.Done():
		fmt.Printf("  wait canceled: %v\n", ctx.Err())
	}
	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
		}
	}

	cancel()
	if err := inp.UnsubscribeInputEvent(event.Hash); err != nil {
		fmt.Printf("  explicit unsubscribe hash=%d: %s\n", event.Hash, formatError(err))
		return
	}
	fmt.Printf("  unsubscribed hash=%d\n", event.Hash)
}

func checkSetInputEvent(inp *input.Input, event core.InputEventDescriptor, value core.InputEventValueMessage, hasValue bool, enabled bool) {
	fmt.Println("\nSimConnect_SetInputEvent")

	if !enabled {
		fmt.Println("  skipped: pass -set-sample to write the currently-read sample value back")
		return
	}
	if !hasValue {
		fmt.Println("  skipped: no current sample value")
		return
	}

	switch value.Type {
	case core.InputEventTypeDouble:
		if err := inp.SetInputEventDouble(event.Hash, value.Double); err != nil {
			fmt.Printf("  set hash=%d double=%f: %s\n", event.Hash, value.Double, formatError(err))
			return
		}
		fmt.Printf("  set hash=%d double=%f ok\n", event.Hash, value.Double)
	case core.InputEventTypeString:
		if err := inp.SetInputEventString(event.Hash, value.String); err != nil {
			fmt.Printf("  set hash=%d string=%q: %s\n", event.Hash, value.String, formatError(err))
			return
		}
		fmt.Printf("  set hash=%d string=%q ok\n", event.Hash, value.String)
	default:
		fmt.Printf("  skipped: unsupported sample value type %s\n", inputEventTypeName(value.Type))
	}
}

func checkInputGroupMapping(client *simconnect.Sim, ed *events.Events) {
	fmt.Println("\nInput Group Mapping APIs")

	const inputDefinition = "VK_LSHIFT+VK_LCONTROL+VK_F24"
	const ex1Group core.InputGroupID = 9001
	const legacyGroup core.InputGroupID = 9002

	downEvent, err := ed.MapEvent("")
	if err != nil {
		fmt.Printf("  create private down event: %s\n", formatError(err))
		return
	}
	upEvent, err := ed.MapEvent("")
	if err != nil {
		fmt.Printf("  create private up event: %s\n", formatError(err))
		return
	}

	if err := client.MapInputEventToClientEventEX1(ex1Group, inputDefinition, downEvent.ID(), 101, upEvent.ID(), 202, false); err != nil {
		fmt.Printf("  SimConnect_MapInputEventToClientEvent_EX1: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_MapInputEventToClientEvent_EX1: group=%d input=%q down=%d up=%d ok\n", ex1Group, inputDefinition, downEvent.ID(), upEvent.ID())
	}
	if err := client.SetInputGroupPriority(ex1Group, core.GroupPriorityDefault); err != nil {
		fmt.Printf("  SimConnect_SetInputGroupPriority: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_SetInputGroupPriority: group=%d priority=%d ok\n", ex1Group, core.GroupPriorityDefault)
	}
	if err := client.SetInputGroupState(ex1Group, core.StateOn); err != nil {
		fmt.Printf("  SimConnect_SetInputGroupState(ON): %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_SetInputGroupState(ON): group=%d ok\n", ex1Group)
	}
	if err := client.SetInputGroupState(ex1Group, core.StateOff); err != nil {
		fmt.Printf("  SimConnect_SetInputGroupState(OFF): %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_SetInputGroupState(OFF): group=%d ok\n", ex1Group)
	}
	if err := client.RemoveInputEvent(ex1Group, inputDefinition); err != nil {
		fmt.Printf("  SimConnect_RemoveInputEvent: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_RemoveInputEvent: group=%d input=%q ok\n", ex1Group, inputDefinition)
	}
	if err := client.ClearInputGroup(ex1Group); err != nil {
		fmt.Printf("  SimConnect_ClearInputGroup: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_ClearInputGroup: group=%d ok\n", ex1Group)
	}

	if err := client.MapInputEventToClientEvent(legacyGroup, inputDefinition, downEvent.ID(), 303, upEvent.ID(), 404, false); err != nil {
		fmt.Printf("  SimConnect_MapInputEventToClientEvent (deprecated): %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_MapInputEventToClientEvent (deprecated): group=%d input=%q ok\n", legacyGroup, inputDefinition)
		_ = client.ClearInputGroup(legacyGroup)
	}
}

func inputEventTypeName(t core.InputEventType) string {
	switch t {
	case core.InputEventTypeDouble:
		return "double"
	case core.InputEventTypeString:
		return "string"
	default:
		return fmt.Sprintf("type_%d", t)
	}
}

func formatInputValue(value core.InputEventValueMessage) string {
	switch value.Type {
	case core.InputEventTypeDouble:
		return fmt.Sprintf("double=%f", value.Double)
	case core.InputEventTypeString:
		return fmt.Sprintf("string=%q", value.String)
	default:
		return "value=n/a"
	}
}

func formatSubscriptionValue(value core.InputEventSubscriptionMessage) string {
	switch value.Type {
	case core.InputEventTypeDouble:
		return fmt.Sprintf("double=%f", value.Double)
	case core.InputEventTypeString:
		return fmt.Sprintf("string=%q", value.String)
	default:
		return "value=n/a"
	}
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
