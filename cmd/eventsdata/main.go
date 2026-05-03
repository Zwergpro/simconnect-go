//go:build windows

package main

import (
	"context"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/events"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/simvar"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/system"
)

type aircraftSnapshot struct {
	Latitude    float64 `sim:"PLANE LATITUDE,degrees"`
	Longitude   float64 `sim:"PLANE LONGITUDE,degrees"`
	Altitude    float64 `sim:"PLANE ALTITUDE,feet"`
	GroundSpeed float64 `sim:"GROUND VELOCITY,knots"`
	OnGround    bool    `sim:"SIM ON GROUND,bool"`
}

type simRateSnapshot struct {
	Rate float64 `sim:"SIMULATION RATE,number"`
}

func main() {
	timeout := flag.Duration("timeout", 25*time.Second, "maximum time to run the Events and Data API checker")
	setSameSimRate := flag.Bool("set-same-sim-rate", false, "exercise SimConnect_SetDataOnSimObject by writing the current SIMULATION RATE value back")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	log.SetFlags(0)

	fmt.Println("SimConnect Events and Data API checker")
	fmt.Printf("started: %s\n", time.Now().Format(time.RFC3339))

	client, err := simconnect.Open(ctx, "simconnect-go events-data checker", simconnect.WithPollInterval(25*time.Millisecond))
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
	gen := client.System()

	checkFlowEvents(ctx, ed)
	checkClientEvents(ed)
	checkSimObjectData(ctx, client, *setSameSimRate)
	checkClientData(ctx, ed)
	checkSimObjectsAndLiveries(ctx, gen)
}

func checkFlowEvents(ctx context.Context, ed *events.Events) {
	fmt.Println("\nFlow API")

	flowCtx, cancel := context.WithCancel(ctx)
	ch, err := ed.SubscribeFlowEvents(flowCtx)
	if err != nil {
		cancel()
		fmt.Printf("  SimConnect_SubscribeToFlowEvent: %s\n", formatError(err))
		return
	}
	fmt.Println("  SimConnect_SubscribeToFlowEvent: ok")

	timer := time.NewTimer(1500 * time.Millisecond)
	select {
	case msg, ok := <-ch:
		if !ok {
			fmt.Println("  flow channel closed")
		} else {
			fmt.Printf("  flow event: event=%d flt=%q\n", msg.Event, msg.FLTPath)
		}
	case <-timer.C:
		fmt.Println("  no flow event in 1.5s")
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

	if err := ed.UnsubscribeFlowEvents(); err != nil {
		fmt.Printf("  SimConnect_UnsubscribeToFlowEvent: %s\n", formatError(err))
		return
	}
	fmt.Println("  SimConnect_UnsubscribeToFlowEvent: ok")
}

func checkClientEvents(ed *events.Events) {
	fmt.Println("\nClient Events / Notification Groups")

	const groupID core.NotificationGroupID = 42
	privateEvent, err := ed.MapEvent("")
	if err != nil {
		fmt.Printf("  SimConnect_MapClientEventToSimEvent(private): %s\n", formatError(err))
		return
	}
	fmt.Printf("  SimConnect_MapClientEventToSimEvent(private): event=%d ok\n", privateEvent.ID())

	if err := ed.AddClientEventToNotificationGroup(groupID, privateEvent, false); err != nil {
		fmt.Printf("  SimConnect_AddClientEventToNotificationGroup: %s\n", formatError(err))
		return
	}
	fmt.Printf("  SimConnect_AddClientEventToNotificationGroup: group=%d event=%d ok\n", groupID, privateEvent.ID())

	if err := ed.RequestNotificationGroup(groupID); err != nil {
		fmt.Printf("  SimConnect_RequestNotificationGroup: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_RequestNotificationGroup: group=%d ok\n", groupID)
	}

	if err := ed.Transmit(privateEvent, 12345, events.WithTransmitGroup(groupID)); err != nil {
		fmt.Printf("  SimConnect_TransmitClientEvent: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_TransmitClientEvent: event=%d data=%d ok\n", privateEvent.ID(), 12345)
	}

	if err := ed.TransmitEX1(privateEvent, 1, 2, 3, 4, 5, events.WithTransmitGroup(groupID)); err != nil {
		fmt.Printf("  SimConnect_TransmitClientEvent_EX1: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_TransmitClientEvent_EX1: event=%d data=[1 2 3 4 5] ok\n", privateEvent.ID())
	}

	if err := ed.RemoveClientEvent(groupID, privateEvent); err != nil {
		fmt.Printf("  SimConnect_RemoveClientEvent: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_RemoveClientEvent: group=%d event=%d ok\n", groupID, privateEvent.ID())
	}

	if err := ed.ClearNotificationGroup(groupID); err != nil {
		fmt.Printf("  SimConnect_ClearNotificationGroup: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_ClearNotificationGroup: group=%d ok\n", groupID)
	}

	keyEvent, err := ed.MapEvent("")
	if err != nil {
		fmt.Printf("  private event for reserved key: %s\n", formatError(err))
		return
	}
	if err := ed.RequestReservedKey(keyEvent, "VK_F24", "VK_F23", "VK_F22"); err != nil {
		fmt.Printf("  SimConnect_RequestReservedKey: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_RequestReservedKey: event=%d choices=[VK_F24 VK_F23 VK_F22] request ok\n", keyEvent.ID())
	}
}

func checkSimObjectData(ctx context.Context, client *simconnect.Sim, setSameSimRate bool) {
	fmt.Println("\nSimObject Data")

	def, err := simvar.Define[aircraftSnapshot]()
	if err != nil {
		fmt.Printf("  SimConnect_AddToDataDefinition: %v\n", err)
		return
	}
	fmt.Println("  SimConnect_AddToDataDefinition: aircraft snapshot definition ok")

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	snapshot, err := simvar.GetOnce(requestCtx, client, def, core.UserAircraft)
	cancel()
	if err != nil {
		fmt.Printf("  SimConnect_RequestDataOnSimObject: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_RequestDataOnSimObject: lat=%.6f lon=%.6f alt=%.1fft gs=%.1fkt onGround=%t\n",
			snapshot.Latitude, snapshot.Longitude, snapshot.Altitude, snapshot.GroundSpeed, snapshot.OnGround)
	}

	requestCtx, cancel = context.WithTimeout(ctx, 5*time.Second)
	byType, err := simvar.GetByTypeOnce(requestCtx, client, def, 1000, core.SimObjectTypeUser)
	cancel()
	if err != nil {
		fmt.Printf("  SimConnect_RequestDataOnSimObjectType: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_RequestDataOnSimObjectType: object=%d lat=%.6f lon=%.6f alt=%.1fft\n",
			byType.ObjectID, byType.Value.Latitude, byType.Value.Longitude, byType.Value.Altitude)
	}

	subCtx, subCancel := context.WithTimeout(ctx, 2500*time.Millisecond)
	ch, err := simvar.Subscribe(subCtx, client, def, core.UserAircraft, core.PeriodSecond, simvar.WithTiming(0, 0, 2))
	if err != nil {
		fmt.Printf("  subscribe RequestDataOnSimObject: %s\n", formatError(err))
	} else {
		count := 0
		for update := range ch {
			count++
			fmt.Printf("  subscription update #%d: object=%d alt=%.1fft gs=%.1fkt\n", count, update.ObjectID, update.Value.Altitude, update.Value.GroundSpeed)
		}
		fmt.Printf("  subscription stopped after %d update(s)\n", count)
	}
	subCancel()

	if !setSameSimRate {
		fmt.Println("  SimConnect_SetDataOnSimObject: skipped; pass -set-same-sim-rate to write current SIMULATION RATE back")
	} else {
		checkSetDataOnSimObject(ctx, client)
	}

	if err := client.ClearDataDefinition(core.DataDefinitionID(1)); err != nil {
		fmt.Printf("  SimConnect_ClearDataDefinition: %s\n", formatError(err))
	} else {
		fmt.Println("  SimConnect_ClearDataDefinition: definition 1 cleared ok")
	}
}

func checkSetDataOnSimObject(ctx context.Context, client *simconnect.Sim) {
	def, err := simvar.Define[simRateSnapshot]()
	if err != nil {
		fmt.Printf("  define sim rate: %v\n", err)
		return
	}
	requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	current, err := simvar.GetOnce(requestCtx, client, def, core.UserAircraft)
	cancel()
	if err != nil {
		fmt.Printf("  read current SIMULATION RATE: %s\n", formatError(err))
		return
	}
	if err := simvar.Set(ctx, client, def, core.UserAircraft, current, core.DataSetDefault); err != nil {
		fmt.Printf("  SimConnect_SetDataOnSimObject: %s\n", formatError(err))
		return
	}
	fmt.Printf("  SimConnect_SetDataOnSimObject: wrote current SIMULATION RATE %.2f back ok\n", current.Rate)
}

func checkClientData(ctx context.Context, ed *events.Events) {
	fmt.Println("\nClient Data")

	name := fmt.Sprintf("simconnect-go-events-data-%d", time.Now().UnixNano())
	const clientDataID core.ClientDataID = 61001
	if err := ed.MapClientDataNameToID(name, clientDataID); err != nil {
		fmt.Printf("  SimConnect_MapClientDataNameToID: %s\n", formatError(err))
		return
	}
	fmt.Printf("  SimConnect_MapClientDataNameToID: %q -> %d ok\n", name, clientDataID)

	if err := ed.CreateClientData(clientDataID, 16, core.ClientDataCreateDefault); err != nil {
		fmt.Printf("  SimConnect_CreateClientData: %s\n", formatError(err))
		return
	}
	fmt.Println("  SimConnect_CreateClientData: size=16 ok")

	def, err := ed.NewClientDataDefinition(events.ClientDataDefinitionItem{
		Offset:     0,
		SizeOrType: 16,
		Epsilon:    0,
		DatumID:    core.Unused,
	})
	if err != nil {
		fmt.Printf("  SimConnect_AddToClientDataDefinition: %s\n", formatError(err))
		return
	}
	fmt.Printf("  SimConnect_AddToClientDataDefinition: def=%d offset=0 size=16 ok\n", def.ID())

	data := make([]byte, 16)
	binary.LittleEndian.PutUint32(data[0:4], 0xC0DEC0DE)
	binary.LittleEndian.PutUint64(data[8:16], math.Float64bits(123.456))
	if err := ed.SetClientData(clientDataID, def, core.ClientDataSetDefault, data); err != nil {
		fmt.Printf("  SimConnect_SetClientData: %s\n", formatError(err))
		return
	}
	fmt.Printf("  SimConnect_SetClientData: bytes=%s ok\n", formatBytes(data))

	requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	msg, err := ed.RequestClientData(requestCtx, clientDataID, def, core.ClientDataPeriodOnce, core.ClientDataRequestDefault, 0, 0, 0)
	cancel()
	if err != nil {
		fmt.Printf("  SimConnect_RequestClientData: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_RequestClientData: request=%d object=%d define=%d flags=%d entries=%d/%d\n",
			msg.RequestID, msg.ObjectID, msg.DefineID, msg.Flags, msg.EntryNumber, msg.OutOf)
	}

	if err := ed.ClearClientDataDefinition(def); err != nil {
		fmt.Printf("  SimConnect_ClearClientDataDefinition: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_ClearClientDataDefinition: def=%d ok\n", def.ID())
	}
}

func checkSimObjectsAndLiveries(ctx context.Context, gen *system.System) {
	fmt.Println("\nSimObjects and Liveries")

	requestCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
	liveries, err := gen.EnumerateSimObjectsAndLiveries(requestCtx, core.SimObjectTypeAircraft)
	cancel()
	if err != nil {
		fmt.Printf("  SimConnect_EnumerateSimObjectsAndLiveries: %s\n", formatError(err))
		return
	}
	fmt.Printf("  SimConnect_EnumerateSimObjectsAndLiveries: page=%d/%d entries=%d\n",
		liveries.EntryNumber+1, liveries.OutOf, len(liveries.Liveries))
	for i, item := range liveries.Liveries {
		if i >= 8 {
			fmt.Printf("  ... %d more\n", len(liveries.Liveries)-i)
			break
		}
		fmt.Printf("  [%d] aircraft=%q livery=%q\n", i, item.AircraftTitle, item.LiveryName)
	}
}

func formatBytes(data []byte) string {
	if len(data) == 0 {
		return "[]"
	}
	parts := make([]byte, 0, len(data)*3+2)
	parts = append(parts, '[')
	for i, b := range data {
		if i > 0 {
			parts = append(parts, ' ')
		}
		parts = fmt.Appendf(parts, "%02X", b)
	}
	parts = append(parts, ']')
	return string(parts)
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
