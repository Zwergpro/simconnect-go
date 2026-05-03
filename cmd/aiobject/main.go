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
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/simvar"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/system"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type aircraftSnapshot struct {
	Latitude  float64 `sim:"PLANE LATITUDE,degrees"`
	Longitude float64 `sim:"PLANE LONGITUDE,degrees"`
	Altitude  float64 `sim:"PLANE ALTITUDE,feet"`
	Pitch     float64 `sim:"PLANE PITCH DEGREES,degrees"`
	Bank      float64 `sim:"PLANE BANK DEGREES,degrees"`
	Heading   float64 `sim:"PLANE HEADING DEGREES TRUE,degrees"`
	OnGround  int32   `sim:"SIM ON GROUND,bool"`
	Airspeed  float64 `sim:"AIRSPEED TRUE,knots"`
}

func main() {
	timeout := flag.Duration("timeout", 25*time.Second, "maximum time to run the AI Object API checker")
	title := flag.String("title", "", "container title to use for AI creation; empty chooses the first enumerated aircraft")
	livery := flag.String("livery", "", "livery name for EX1 creation calls; empty uses the enumerated livery or simulator default")
	tail := flag.String("tail", "SCGOAI", "tail number for aircraft creation calls")
	airport := flag.String("airport", "KSEA", "airport ICAO for parked aircraft creation")
	flightPlan := flag.String("flight-plan", "", "flight plan path for enroute creation or AISetAircraftFlightPlan")
	createParked := flag.Bool("create-parked", false, "exercise AICreateParkedATCAircraft and AICreateParkedATCAircraft_EX1")
	createNonATC := flag.Bool("create-non-atc", false, "exercise AICreateNonATCAircraft and AICreateNonATCAircraft_EX1 near the user aircraft")
	createEnroute := flag.Bool("create-enroute", false, "exercise AICreateEnrouteATCAircraft and AICreateEnrouteATCAircraft_EX1 using -flight-plan")
	createObject := flag.Bool("create-object", false, "exercise AICreateSimulatedObject and AICreateSimulatedObject_EX1 using -object-title")
	objectTitle := flag.String("object-title", "", "container title for simulated-object creation")
	setPlan := flag.Bool("set-plan", false, "after a parked aircraft is created, call AISetAircraftFlightPlan with -flight-plan")
	releaseControl := flag.Bool("release-control", false, "after creating a non-ATC aircraft or simulated object, call AIReleaseControl before cleanup")
	keepObjects := flag.Bool("keep-objects", false, "do not call AIRemoveObject for created objects")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	log.SetFlags(0)

	fmt.Println("SimConnect AI Object API checker")
	fmt.Printf("started: %s\n", time.Now().Format(time.RFC3339))

	client, err := simconnect.Open(ctx, "simconnect-go ai object checker", simconnect.WithPollInterval(25*time.Millisecond))
	if err != nil {
		log.Fatalf("SimConnect_Open: %v", err)
	}
	fmt.Println("\nSimConnect_Open")
	fmt.Println("  connected")
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("\nSimConnect_Close: %s\n", formatError(err))
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
	gen := client.System()

	snapshot, hasSnapshot := checkUserAircraftSnapshot(ctx, client)
	aircraft := checkEnumerateAircraft(ctx, gen)
	currentPlan := requestSystemState(ctx, gen, "FlightPlan")
	printConfiguration(*title, *livery, *tail, *airport, *flightPlan, *objectTitle, currentPlan.String)

	selectedTitle, selectedLivery := chooseAircraft(*title, *livery, aircraft)
	if selectedTitle != "" {
		fmt.Printf("  selected aircraft: title=%q livery=%q\n", selectedTitle, selectedLivery)
	}

	created := []createdObject{}
	if *createParked {
		created = append(created, checkCreateParked(ctx, aiClient, selectedTitle, selectedLivery, *tail, *airport)...)
	} else {
		fmt.Println("\nParked ATC Aircraft")
		fmt.Println("  skipped: pass -create-parked to exercise AICreateParkedATCAircraft and AICreateParkedATCAircraft_EX1")
	}

	if *createNonATC {
		created = append(created, checkCreateNonATC(ctx, aiClient, selectedTitle, selectedLivery, *tail, snapshot, hasSnapshot)...)
	} else {
		fmt.Println("\nNon-ATC Aircraft")
		fmt.Println("  skipped: pass -create-non-atc to exercise AICreateNonATCAircraft and AICreateNonATCAircraft_EX1")
	}

	enroutePlan := firstNonBlank(*flightPlan, currentPlan.String)
	if *createEnroute {
		created = append(created, checkCreateEnroute(ctx, aiClient, selectedTitle, selectedLivery, *tail, enroutePlan)...)
	} else {
		fmt.Println("\nEnroute ATC Aircraft")
		fmt.Println("  skipped: pass -create-enroute and -flight-plan to exercise AICreateEnrouteATCAircraft and AICreateEnrouteATCAircraft_EX1")
	}

	if *createObject {
		created = append(created, checkCreateSimulatedObject(ctx, aiClient, *objectTitle, *livery, snapshot, hasSnapshot)...)
	} else {
		fmt.Println("\nSimulated Object")
		fmt.Println("  skipped: pass -create-object and -object-title to exercise AICreateSimulatedObject and AICreateSimulatedObject_EX1")
	}

	if *setPlan {
		checkAISetAircraftFlightPlan(aiClient, created, enroutePlan)
	} else {
		fmt.Println("\nSimConnect_AISetAircraftFlightPlan")
		fmt.Println("  skipped: pass -set-plan with -flight-plan after creating a parked aircraft")
	}

	if *releaseControl {
		checkAIReleaseControl(aiClient, created)
	} else {
		fmt.Println("\nSimConnect_AIReleaseControl")
		fmt.Println("  skipped: pass -release-control to release AI control on created non-ATC/object entries")
	}

	checkAIRemoveObject(aiClient, created, *keepObjects)
}

type createdObject struct {
	ID     core.ObjectID
	Kind   string
	Title  string
	Livery string
}

func checkUserAircraftSnapshot(ctx context.Context, client *simconnect.Sim) (aircraftSnapshot, bool) {
	fmt.Println("\nUser Aircraft Snapshot")

	def, err := simvar.Define[aircraftSnapshot]()
	if err != nil {
		fmt.Printf("  define: %v\n", err)
		return aircraftSnapshot{}, false
	}

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	snapshot, err := simvar.GetOnce(requestCtx, client, def, core.UserAircraft)
	cancel()
	if err != nil {
		fmt.Printf("  request: %s\n", formatError(err))
		return aircraftSnapshot{}, false
	}

	fmt.Printf(
		"  lat=%.6f lon=%.6f alt=%.1fft pitch=%.2f bank=%.2f heading=%.2f onGround=%t airspeed=%.1fkt\n",
		snapshot.Latitude,
		snapshot.Longitude,
		snapshot.Altitude,
		snapshot.Pitch,
		snapshot.Bank,
		snapshot.Heading,
		snapshot.OnGround != 0,
		snapshot.Airspeed,
	)
	return snapshot, true
}

func checkEnumerateAircraft(ctx context.Context, gen *system.System) core.SimObjectLiveryListMessage {
	fmt.Println("\nSimConnect_EnumerateSimObjectsAndLiveries")

	requestCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
	aircraft, err := gen.EnumerateSimObjectsAndLiveries(requestCtx, core.SimObjectTypeAircraft)
	cancel()
	if err != nil {
		fmt.Printf("  aircraft error: %s\n", formatError(err))
		return core.SimObjectLiveryListMessage{}
	}

	fmt.Printf("  aircraft page=%d/%d entries=%d\n", aircraft.EntryNumber+1, aircraft.OutOf, len(aircraft.Liveries))
	for i, item := range aircraft.Liveries {
		if i >= 12 {
			fmt.Printf("  ... %d more\n", len(aircraft.Liveries)-i)
			break
		}
		fmt.Printf("  [%d] title=%q livery=%q\n", i, item.AircraftTitle, item.LiveryName)
	}
	return aircraft
}

func printConfiguration(title, livery, tail, airport, flightPlan, objectTitle, currentPlan string) {
	fmt.Println("\nConfiguration")
	fmt.Printf("  title=%q livery=%q tail=%q airport=%q\n", title, livery, tail, airport)
	fmt.Printf("  flightPlan=%q currentFlightPlan=%q\n", flightPlan, currentPlan)
	fmt.Printf("  objectTitle=%q\n", objectTitle)
}

func checkCreateParked(ctx context.Context, aiClient *ai.AI, title, livery, tail, airport string) []createdObject {
	fmt.Println("\nParked ATC Aircraft")

	if title == "" {
		fmt.Println("  skipped: no aircraft title available; pass -title")
		return nil
	}
	if strings.TrimSpace(airport) == "" {
		fmt.Println("  skipped: pass -airport with an ICAO code")
		return nil
	}

	var created []createdObject
	runCreate(ctx, "SimConnect_AICreateParkedATCAircraft", func(requestCtx context.Context) (core.ObjectID, error) {
		return aiClient.CreateParkedATC(requestCtx, title, tail, airport)
	}, func(id core.ObjectID) {
		created = append(created, createdObject{ID: id, Kind: "parked", Title: title})
	})

	runCreate(ctx, "SimConnect_AICreateParkedATCAircraft_EX1", func(requestCtx context.Context) (core.ObjectID, error) {
		return aiClient.CreateParkedATCEX1(requestCtx, title, livery, tail+"X", airport)
	}, func(id core.ObjectID) {
		created = append(created, createdObject{ID: id, Kind: "parked_ex1", Title: title, Livery: livery})
	})
	return created
}

func checkCreateNonATC(ctx context.Context, aiClient *ai.AI, title, livery, tail string, snapshot aircraftSnapshot, hasSnapshot bool) []createdObject {
	fmt.Println("\nNon-ATC Aircraft")

	if title == "" {
		fmt.Println("  skipped: no aircraft title available; pass -title")
		return nil
	}
	if !hasSnapshot {
		fmt.Println("  skipped: no user aircraft position available")
		return nil
	}

	initPos := initPositionNear(snapshot, 0.002)
	var created []createdObject
	runCreate(ctx, "SimConnect_AICreateNonATCAircraft", func(requestCtx context.Context) (core.ObjectID, error) {
		return aiClient.CreateNonATC(requestCtx, title, tail, initPos)
	}, func(id core.ObjectID) {
		created = append(created, createdObject{ID: id, Kind: "non_atc", Title: title})
	})

	initPos = initPositionNear(snapshot, 0.004)
	runCreate(ctx, "SimConnect_AICreateNonATCAircraft_EX1", func(requestCtx context.Context) (core.ObjectID, error) {
		return aiClient.CreateNonATCEX1(requestCtx, title, livery, tail+"X", initPos)
	}, func(id core.ObjectID) {
		created = append(created, createdObject{ID: id, Kind: "non_atc_ex1", Title: title, Livery: livery})
	})
	return created
}

func checkCreateEnroute(ctx context.Context, aiClient *ai.AI, title, livery, tail, flightPlanPath string) []createdObject {
	fmt.Println("\nEnroute ATC Aircraft")

	if title == "" {
		fmt.Println("  skipped: no aircraft title available; pass -title")
		return nil
	}
	if strings.TrimSpace(flightPlanPath) == "" {
		fmt.Println("  skipped: pass -flight-plan or load a flight plan in the simulator")
		return nil
	}

	var created []createdObject
	runCreate(ctx, "SimConnect_AICreateEnrouteATCAircraft", func(requestCtx context.Context) (core.ObjectID, error) {
		return aiClient.CreateEnrouteATC(requestCtx, title, tail, 9001, flightPlanPath, 0, false)
	}, func(id core.ObjectID) {
		created = append(created, createdObject{ID: id, Kind: "enroute", Title: title})
	})

	runCreate(ctx, "SimConnect_AICreateEnrouteATCAircraft_EX1", func(requestCtx context.Context) (core.ObjectID, error) {
		return aiClient.CreateEnrouteATCEX1(requestCtx, title, livery, tail+"X", 9002, flightPlanPath, 0, false)
	}, func(id core.ObjectID) {
		created = append(created, createdObject{ID: id, Kind: "enroute_ex1", Title: title, Livery: livery})
	})
	return created
}

func checkCreateSimulatedObject(ctx context.Context, aiClient *ai.AI, title, livery string, snapshot aircraftSnapshot, hasSnapshot bool) []createdObject {
	fmt.Println("\nSimulated Object")

	if strings.TrimSpace(title) == "" {
		fmt.Println("  skipped: pass -object-title with a valid non-aircraft SimObject container title")
		return nil
	}
	if !hasSnapshot {
		fmt.Println("  skipped: no user aircraft position available")
		return nil
	}

	initPos := initPositionNear(snapshot, 0.006)
	var created []createdObject
	runCreate(ctx, "SimConnect_AICreateSimulatedObject", func(requestCtx context.Context) (core.ObjectID, error) {
		return aiClient.CreateSimulatedObject(requestCtx, title, initPos)
	}, func(id core.ObjectID) {
		created = append(created, createdObject{ID: id, Kind: "object", Title: title})
	})

	initPos = initPositionNear(snapshot, 0.008)
	runCreate(ctx, "SimConnect_AICreateSimulatedObject_EX1", func(requestCtx context.Context) (core.ObjectID, error) {
		return aiClient.CreateSimulatedObjectEX1(requestCtx, title, livery, initPos)
	}, func(id core.ObjectID) {
		created = append(created, createdObject{ID: id, Kind: "object_ex1", Title: title, Livery: livery})
	})
	return created
}

func runCreate(ctx context.Context, name string, create func(context.Context) (core.ObjectID, error), record func(core.ObjectID)) {
	requestCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
	id, err := create(requestCtx)
	cancel()
	if err != nil {
		fmt.Printf("  %-42s error: %s\n", name, formatError(err))
		return
	}
	fmt.Printf("  %-42s objectID=%d\n", name, id)
	record(id)
}

func checkAISetAircraftFlightPlan(aiClient *ai.AI, created []createdObject, flightPlanPath string) {
	fmt.Println("\nSimConnect_AISetAircraftFlightPlan")

	if strings.TrimSpace(flightPlanPath) == "" {
		fmt.Println("  skipped: no flight plan path available")
		return
	}
	for _, obj := range created {
		if obj.Kind != "parked" && obj.Kind != "parked_ex1" {
			continue
		}
		if err := aiClient.SetFlightPlan(obj.ID, flightPlanPath); err != nil {
			fmt.Printf("  objectID=%d kind=%s path=%q error: %s\n", obj.ID, obj.Kind, flightPlanPath, formatError(err))
		} else {
			fmt.Printf("  objectID=%d kind=%s path=%q request sent\n", obj.ID, obj.Kind, flightPlanPath)
		}
		return
	}
	fmt.Println("  skipped: no parked AI aircraft was created in this run")
}

func checkAIReleaseControl(aiClient *ai.AI, created []createdObject) {
	fmt.Println("\nSimConnect_AIReleaseControl")

	ran := false
	for _, obj := range created {
		if obj.Kind == "parked" || obj.Kind == "parked_ex1" || obj.Kind == "enroute" || obj.Kind == "enroute_ex1" {
			continue
		}
		ran = true
		if err := aiClient.ReleaseControl(obj.ID); err != nil {
			fmt.Printf("  objectID=%d kind=%s error: %s\n", obj.ID, obj.Kind, formatError(err))
		} else {
			fmt.Printf("  objectID=%d kind=%s request sent\n", obj.ID, obj.Kind)
		}
	}
	if !ran {
		fmt.Println("  skipped: no created non-ATC aircraft or simulated object")
	}
}

func checkAIRemoveObject(aiClient *ai.AI, created []createdObject, keep bool) {
	fmt.Println("\nSimConnect_AIRemoveObject")

	if len(created) == 0 {
		fmt.Println("  skipped: no objects were created")
		return
	}
	if keep {
		fmt.Printf("  skipped cleanup for %d object(s): -keep-objects was set\n", len(created))
		for _, obj := range created {
			fmt.Printf("  kept objectID=%d kind=%s title=%q livery=%q\n", obj.ID, obj.Kind, obj.Title, obj.Livery)
		}
		return
	}

	for _, obj := range created {
		if err := aiClient.Remove(obj.ID); err != nil {
			fmt.Printf("  objectID=%d kind=%s error: %s\n", obj.ID, obj.Kind, formatError(err))
		} else {
			fmt.Printf("  objectID=%d kind=%s remove requested\n", obj.ID, obj.Kind)
		}
	}
}

func initPositionNear(snapshot aircraftSnapshot, longitudeOffset float64) core.InitPosition {
	altitude := snapshot.Altitude
	if snapshot.OnGround != 0 {
		altitude += 8
	}
	airspeed := uint32(snapshot.Airspeed)
	if airspeed < 40 {
		airspeed = core.InitPositionAirspeedKeep
	}
	return core.InitPosition{
		Latitude:  snapshot.Latitude,
		Longitude: snapshot.Longitude + longitudeOffset,
		Altitude:  altitude,
		Pitch:     snapshot.Pitch,
		Bank:      snapshot.Bank,
		Heading:   snapshot.Heading,
		OnGround:  uint32(snapshot.OnGround),
		Airspeed:  airspeed,
	}
}

func chooseAircraft(title, livery string, aircraft core.SimObjectLiveryListMessage) (string, string) {
	if title != "" {
		return title, livery
	}
	for _, item := range aircraft.Liveries {
		if item.AircraftTitle == "" {
			continue
		}
		selectedLivery := livery
		if selectedLivery == "" {
			selectedLivery = item.LiveryName
		}
		return item.AircraftTitle, selectedLivery
	}
	return "", livery
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

func firstNonBlank(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
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
