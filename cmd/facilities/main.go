//go:build windows

package main

import (
	"context"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/facilities"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/simvar"
	"log"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type positionSnapshot struct {
	Latitude  float64 `sim:"PLANE LATITUDE,degrees"`
	Longitude float64 `sim:"PLANE LONGITUDE,degrees"`
	OnGround  int32   `sim:"SIM ON GROUND,bool"`
}

type airportCandidate struct {
	Ident          string
	Region         string
	Latitude       float64
	Longitude      float64
	AltitudeMeters float64
	DistanceMeters float64
}

func main() {
	timeout := flag.Duration("timeout", 30*time.Second, "maximum time to run the Facilities API checker")
	includeAll := flag.Bool("all", false, "also call SimConnect_RequestAllFacilities; this can return a large worldwide list")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	log.SetFlags(0)

	fmt.Println("SimConnect Facilities API checker")
	fmt.Printf("started: %s\n", time.Now().Format(time.RFC3339))

	client, err := simconnect.Open(ctx, "simconnect-go facilities checker", simconnect.WithPollInterval(25*time.Millisecond))
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

	fac := client.Facilities()
	position, _ := requestPosition(ctx)

	airports := checkFacilitiesLists(ctx, fac)
	nearest, hasNearest := printNearestAirport(position, airports)
	checkFacilityDefinition(ctx, fac, nearest, hasNearest)
	checkJetwayData(ctx, fac, nearest, hasNearest)
	checkFacilitySubscription(ctx, fac)
	if *includeAll {
		checkRequestAllFacilities(ctx, fac)
	} else {
		fmt.Println("\nSimConnect_RequestAllFacilities")
		fmt.Println("  skipped: pass -all to request a worldwide list")
	}
}

func requestPosition(ctx context.Context) (positionSnapshot, bool) {
	auto, err := simconnect.Open(ctx, "simconnect-go facilities position", simconnect.WithPollInterval(25*time.Millisecond))
	if err != nil {
		fmt.Printf("  position helper connect: %v\n", err)
		return positionSnapshot{}, false
	}
	defer auto.Close()

	def, err := simvar.Define[positionSnapshot]()
	if err != nil {
		fmt.Printf("  define position helper: %v\n", err)
		return positionSnapshot{}, false
	}

	requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	pos, err := simvar.GetOnce(requestCtx, auto, def, core.UserAircraft)
	if err != nil {
		fmt.Printf("  request position helper: %v\n", err)
		return positionSnapshot{}, false
	}
	return pos, true
}

func checkFacilitiesLists(ctx context.Context, fac *facilities.Facilities) []airportCandidate {
	fmt.Println("\nSimConnect_RequestFacilitiesList_EX1")

	requests := []struct {
		name string
		typ  core.FacilityListType
	}{
		{"AIRPORT", core.FacilityListTypeAirport},
		{"WAYPOINT", core.FacilityListTypeWaypoint},
		{"NDB", core.FacilityListTypeNDB},
		{"VOR", core.FacilityListTypeVOR},
	}

	var airports []airportCandidate
	for _, req := range requests {
		requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		packets, err := fac.RequestFacilitiesList(requestCtx, req.typ)
		cancel()
		if err != nil {
			fmt.Printf("  %-8s request error: %s\n", req.name, formatError(err))
			continue
		}
		count := 0
		for _, msg := range packets {
			switch m := msg.(type) {
			case core.AirportListMessage:
				count += len(m.Airports)
				for _, airport := range m.Airports {
					airports = append(airports, airportCandidateFromAirport(airport))
				}
			case core.WaypointListMessage:
				count += len(m.Waypoints)
			case core.NDBListMessage:
				count += len(m.NDBs)
			case core.VORListMessage:
				count += len(m.VORs)
			}
		}
		fmt.Printf("  %-8s packets=%d entries=%d\n", req.name, len(packets), count)
		printListSample(req.name, packets)
	}
	return airports
}

func printNearestAirport(pos positionSnapshot, airports []airportCandidate) (airportCandidate, bool) {
	fmt.Println("\nNearest Cached Airport")
	if len(airports) == 0 {
		fmt.Println("  none returned")
		return airportCandidate{}, false
	}

	best := airports[0]
	for _, airport := range airports {
		airport.DistanceMeters = greatCircleDistanceMeters(pos.Latitude, pos.Longitude, airport.Latitude, airport.Longitude)
		if best.DistanceMeters == 0 || airport.DistanceMeters < best.DistanceMeters {
			best = airport
		}
	}
	fmt.Printf("  aircraft lat/lon: %.6f, %.6f onGround=%t\n", pos.Latitude, pos.Longitude, pos.OnGround != 0)
	fmt.Printf("  nearest: %s%s lat=%.6f lon=%.6f elev=%.0fm distance=%.0fm\n",
		best.Ident, formatRegion(best.Region), best.Latitude, best.Longitude, best.AltitudeMeters, best.DistanceMeters)
	return best, true
}

func checkFacilityDefinition(ctx context.Context, fac *facilities.Facilities, airport airportCandidate, ok bool) {
	fmt.Println("\nFacility Definition / Data")
	if !ok || airport.Ident == "" {
		fmt.Println("  skipped: no airport candidate")
		return
	}

	fields := []string{
		"OPEN AIRPORT",
		"ICAO",
		"REGION",
		"LATITUDE",
		"LONGITUDE",
		"ALTITUDE",
		"OPEN RUNWAY",
		"PRIMARY_NUMBER",
		"SECONDARY_NUMBER",
		"CLOSE RUNWAY",
		"CLOSE AIRPORT",
	}
	def, err := fac.NewFacilityDefinition(fields...)
	if err != nil {
		fmt.Printf("  SimConnect_AddToFacilityDefinition: %s\n", formatError(err))
		return
	}
	fmt.Printf("  SimConnect_AddToFacilityDefinition: %d fields ok\n", len(fields))

	var primaryNumber uint32 = 1
	var primaryNumberData [4]byte
	binary.LittleEndian.PutUint32(primaryNumberData[:], primaryNumber)
	err = fac.AddFacilityDataDefinitionFilter(
		def,
		facilities.FacilityDataFilter{
			Path: "AIRPORT:RUNWAY:PRIMARY_NUMBER",
			Data: primaryNumberData[:],
		},
	)
	if err != nil {
		fmt.Printf("  SimConnect_AddFacilityDataDefinitionFilter: %s\n", formatError(err))
	} else {
		fmt.Printf("  SimConnect_AddFacilityDataDefinitionFilter: PRIMARY_NUMBER=%d ok\n", primaryNumber)
	}
	if err := fac.ClearAllFacilityDataDefinitionFilters(def); err != nil {
		fmt.Printf("  SimConnect_ClearAllFacilityDataDefinitionFilters: %s\n", formatError(err))
	} else {
		fmt.Println("  SimConnect_ClearAllFacilityDataDefinitionFilters: ok")
	}

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	data, err := fac.RequestFacilityData(requestCtx, def, airport.Ident, airport.Region)
	cancel()
	if err != nil {
		fmt.Printf("  SimConnect_RequestFacilityData(%s%s): %s\n", airport.Ident, formatRegion(airport.Region), formatError(err))
	} else {
		printFacilityDataMessages(data, "SimConnect_RequestFacilityData")
	}

	requestCtx, cancel = context.WithTimeout(ctx, 5*time.Second)
	data, err = fac.RequestFacilityDataEX1(requestCtx, def, airport.Ident, airport.Region, 0)
	cancel()
	if err != nil {
		fmt.Printf("  SimConnect_RequestFacilityData_EX1(%s%s,type=0): %s\n", airport.Ident, formatRegion(airport.Region), formatError(err))
	} else {
		printFacilityDataMessages(data, "SimConnect_RequestFacilityData_EX1")
	}
}

func checkJetwayData(ctx context.Context, fac *facilities.Facilities, airport airportCandidate, ok bool) {
	fmt.Println("\nSimConnect_RequestJetwayData")
	if !ok || airport.Ident == "" {
		fmt.Println("  skipped: no airport candidate")
		return
	}
	requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	m, err := fac.RequestJetwayData(requestCtx, airport.Ident, nil)
	cancel()
	if err != nil {
		fmt.Printf("  request %s: %s\n", airport.Ident, formatError(err))
		return
	}
	fmt.Printf("  packets=1 jetways=%d requestID=%d\n", len(m.Jetways), m.RequestID)
	for i, jetway := range m.Jetways {
		if i >= 3 {
			fmt.Printf("    ... %d more\n", len(m.Jetways)-i)
			break
		}
		fmt.Printf("    %s parking=%d status=%d attached=%d\n",
			jetway.AirportICAO, jetway.ParkingIndex, jetway.Status, jetway.AttachedObjectID)
	}
}

func checkFacilitySubscription(ctx context.Context, fac *facilities.Facilities) {
	fmt.Println("\nSubscribe / Unsubscribe Facilities")

	subCtx, cancel := context.WithCancel(ctx)
	newCh, _, err := fac.SubscribeFacilitiesEX1(subCtx, core.FacilityListTypeAirport)
	if err != nil {
		fmt.Printf("  SimConnect_SubscribeToFacilities_EX1: %s\n", formatError(err))
		return
	}
	fmt.Println("  SimConnect_SubscribeToFacilities_EX1(AIRPORT): ok")

	packets := collectFromChannel(ctx, newCh, 3*time.Second, 2)
	fmt.Printf("  new-in-range packets=%d\n", len(packets))
	printListSample("SUB AIRPORT", packets)
	cancel()
	fmt.Println("  SimConnect_UnsubscribeToFacilities_EX1: requested by canceling subscription context")

	subCtx, cancel = context.WithCancel(ctx)
	vorCh, err := fac.SubscribeFacilities(subCtx, core.FacilityListTypeVOR)
	if err != nil {
		cancel()
		fmt.Printf("  SimConnect_SubscribeToFacilities(VOR): %s\n", formatError(err))
		return
	}
	fmt.Println("  SimConnect_SubscribeToFacilities(VOR): ok")
	_ = collectFromChannel(ctx, vorCh, 1*time.Second, 1)
	cancel()
	fmt.Println("  SimConnect_UnsubscribeToFacilities: requested by canceling subscription context")
}

func checkRequestAllFacilities(ctx context.Context, fac *facilities.Facilities) {
	fmt.Println("\nSimConnect_RequestAllFacilities")
	requestCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
	packets, err := fac.RequestAllFacilities(requestCtx, core.FacilityListTypeVOR)
	cancel()
	if err != nil {
		fmt.Printf("  request all VOR: %s\n", formatError(err))
		return
	}
	count := 0
	for _, msg := range packets {
		if m, ok := msg.(core.VORListMessage); ok {
			count += len(m.VORs)
		}
	}
	fmt.Printf("  VOR packets=%d entries=%d\n", len(packets), count)
	printListSample("ALL VOR", packets)
}

func printFacilityDataMessages(data []core.FacilityDataMessage, title string) {
	for i, msg := range data {
		if i >= 6 {
			break
		}
		fmt.Printf("  %s data #%d type=%s list=%t index=%d/%d payload=%d %s\n",
			title, i+1, facilityDataTypeName(msg.Type), msg.IsListItem, msg.ItemIndex, msg.ListSize, len(msg.Payload), formatPayload(msg.Payload))
	}
	fmt.Printf("  %s complete=true dataMessages=%d\n", title, len(data))
}

func collectFromChannel(ctx context.Context, ch <-chan core.Message, timeout time.Duration, maxPackets int) []core.Message {
	var packets []core.Message
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	for len(packets) < maxPackets {
		select {
		case msg, ok := <-ch:
			if !ok {
				return packets
			}
			packets = append(packets, msg)
		case <-timer.C:
			return packets
		case <-ctx.Done():
			return packets
		}
	}
	return packets
}

func collectPackets(ctx context.Context, client *simconnect.Sim, requestID uint32, timeout time.Duration, maxPackets int) []core.Message {
	var packets []core.Message
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) && len(packets) < maxPackets {
		msg, ok := nextMessage(ctx, client, 250*time.Millisecond)
		if !ok {
			continue
		}
		switch m := msg.(type) {
		case core.AirportListMessage:
			if m.RequestID == requestID {
				packets = append(packets, m)
				if m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf {
					return packets
				}
			} else {
				printUnexpected(msg)
			}
		case core.WaypointListMessage:
			if m.RequestID == requestID {
				packets = append(packets, m)
				if m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf {
					return packets
				}
			} else {
				printUnexpected(msg)
			}
		case core.NDBListMessage:
			if m.RequestID == requestID {
				packets = append(packets, m)
				if m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf {
					return packets
				}
			} else {
				printUnexpected(msg)
			}
		case core.VORListMessage:
			if m.RequestID == requestID {
				packets = append(packets, m)
				if m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf {
					return packets
				}
			} else {
				printUnexpected(msg)
			}
		case core.ExceptionError:
			fmt.Printf("  exception: %s\n", formatError(m))
			return packets
		default:
			printUnexpected(msg)
		}
	}
	return packets
}

func nextMessage(ctx context.Context, client *simconnect.Sim, timeout time.Duration) (core.Message, bool) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		msg, ok, err := client.GetNextMessage()
		if err != nil {
			fmt.Printf("  GetNextDispatch: %s\n", formatError(err))
			return nil, false
		}
		if !ok {
			select {
			case <-ctx.Done():
				return nil, false
			case <-time.After(25 * time.Millisecond):
				continue
			}
		}
		return msg, true
	}
	return nil, false
}

func drainOpen(ctx context.Context, client *simconnect.Sim) {
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		msg, ok := nextMessage(ctx, client, 250*time.Millisecond)
		if !ok {
			continue
		}
		if open, ok := msg.(core.OpenMessage); ok {
			fmt.Printf("  server app=%q simconnect=%d.%d.%d.%d\n",
				open.ApplicationName,
				open.SimConnectVersionMajor,
				open.SimConnectVersionMinor,
				open.SimConnectBuildMajor,
				open.SimConnectBuildMinor,
			)
			return
		}
		printUnexpected(msg)
	}
}

func printListSample(name string, packets []core.Message) {
	for _, msg := range packets {
		switch m := msg.(type) {
		case core.AirportListMessage:
			if len(m.Airports) > 0 {
				a := airportCandidateFromAirport(m.Airports[0])
				fmt.Printf("    %s sample airport: %s%s lat=%.5f lon=%.5f\n", name, a.Ident, formatRegion(a.Region), a.Latitude, a.Longitude)
			}
			return
		case core.WaypointListMessage:
			if len(m.Waypoints) > 0 {
				w := airportCandidateFromAirport(m.Waypoints[0].Airport)
				fmt.Printf("    %s sample waypoint: %s%s lat=%.5f lon=%.5f\n", name, w.Ident, formatRegion(w.Region), w.Latitude, w.Longitude)
			}
			return
		case core.NDBListMessage:
			if len(m.NDBs) > 0 {
				n := airportCandidateFromAirport(m.NDBs[0].Airport)
				fmt.Printf("    %s sample NDB: %s%s freq=%d\n", name, n.Ident, formatRegion(n.Region), m.NDBs[0].Frequency)
			}
			return
		case core.VORListMessage:
			if len(m.VORs) > 0 {
				v := airportCandidateFromAirport(m.VORs[0].Airport)
				fmt.Printf("    %s sample VOR: %s%s freq=%d flags=%d\n", name, v.Ident, formatRegion(v.Region), m.VORs[0].Frequency, m.VORs[0].Flags)
			}
			return
		}
	}
}

func airportCandidateFromAirport(airport core.Airport) airportCandidate {
	return airportCandidate{
		Ident:          airport.Ident,
		Region:         airport.Region,
		Latitude:       airport.Latitude,
		Longitude:      airport.Longitude,
		AltitudeMeters: airport.Altitude,
	}
}

func formatRegion(region string) string {
	if region == "" {
		return ""
	}
	return "/" + region
}

func greatCircleDistanceMeters(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadiusMeters = 6371000.0
	phi1 := degreesToRadians(lat1)
	phi2 := degreesToRadians(lat2)
	dPhi := degreesToRadians(lat2 - lat1)
	dLambda := degreesToRadians(lon2 - lon1)
	a := math.Sin(dPhi/2)*math.Sin(dPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*math.Sin(dLambda/2)*math.Sin(dLambda/2)
	return 2 * earthRadiusMeters * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func degreesToRadians(v float64) float64 {
	return v * math.Pi / 180
}

func printUnexpected(msg core.Message) {
	switch m := msg.(type) {
	case core.OpenMessage:
	case core.QuitMessage:
		fmt.Println("  simulator quit message received")
	case core.ExceptionError:
		fmt.Printf("  async exception: %s\n", formatError(m))
	default:
		fmt.Printf("  unexpected message: %T recvID=%d\n", msg, msg.RecvID())
	}
}

func formatPayload(payload []byte) string {
	if len(payload) == 0 {
		return "[]"
	}
	n := len(payload)
	if n > 24 {
		n = 24
	}
	parts := make([]string, n)
	for i := 0; i < n; i++ {
		parts[i] = fmt.Sprintf("%02X", payload[i])
	}
	if len(payload) > n {
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

func facilityDataTypeName(t core.FacilityDataType) string {
	names := map[core.FacilityDataType]string{
		core.FacilityDataAirport:            "AIRPORT",
		core.FacilityDataRunway:             "RUNWAY",
		core.FacilityDataStart:              "START",
		core.FacilityDataFrequency:          "FREQUENCY",
		core.FacilityDataHelipad:            "HELIPAD",
		core.FacilityDataApproach:           "APPROACH",
		core.FacilityDataTaxiPoint:          "TAXI_POINT",
		core.FacilityDataTaxiParking:        "TAXI_PARKING",
		core.FacilityDataTaxiPath:           "TAXI_PATH",
		core.FacilityDataTaxiName:           "TAXI_NAME",
		core.FacilityDataJetway:             "JETWAY",
		core.FacilityDataVOR:                "VOR",
		core.FacilityDataNDB:                "NDB",
		core.FacilityDataWaypoint:           "WAYPOINT",
		core.FacilityDataRoute:              "ROUTE",
		core.FacilityDataPavement:           "PAVEMENT",
		core.FacilityDataApproachLights:     "APPROACH_LIGHTS",
		core.FacilityDataVASI:               "VASI",
		core.FacilityDataVDGS:               "VDGS",
		core.FacilityDataHoldingPattern:     "HOLDING_PATTERN",
		core.FacilityDataTaxiParkingAirline: "TAXI_PARKING_AIRLINE",
	}
	if name, ok := names[t]; ok {
		return name
	}
	return fmt.Sprintf("FACILITY_DATA_%d", t)
}

func exceptionName(exception core.Exception) string {
	return exception.String()
}
