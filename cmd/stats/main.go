//go:build windows

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/debug"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/facilities"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/input"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/simvar"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/system"
)

type positionStats struct {
	Latitude        float64 `sim:"PLANE LATITUDE,degrees"`
	Longitude       float64 `sim:"PLANE LONGITUDE,degrees"`
	Altitude        float64 `sim:"PLANE ALTITUDE,feet"`
	AltitudeAGL     float64 `sim:"PLANE ALT ABOVE GROUND,feet"`
	GroundAltitude  float64 `sim:"GROUND ALTITUDE,feet"`
	Pitch           float64 `sim:"PLANE PITCH DEGREES,degrees"`
	Bank            float64 `sim:"PLANE BANK DEGREES,degrees"`
	HeadingTrue     float64 `sim:"PLANE HEADING DEGREES TRUE,degrees"`
	HeadingMagnetic float64 `sim:"PLANE HEADING DEGREES MAGNETIC,degrees"`
	OnGround        bool    `sim:"SIM ON GROUND,bool"`
}

type motionStats struct {
	IndicatedAirspeed float64 `sim:"AIRSPEED INDICATED,knots"`
	TrueAirspeed      float64 `sim:"AIRSPEED TRUE,knots"`
	GroundSpeed       float64 `sim:"GROUND VELOCITY,knots"`
	VerticalSpeed     float64 `sim:"VERTICAL SPEED,feet per minute"`
	Mach              float64 `sim:"AIRSPEED MACH,mach"`
	GForce            float64 `sim:"G FORCE,gforce"`
	VelocityBodyX     float64 `sim:"VELOCITY BODY X,feet per second"`
	VelocityBodyY     float64 `sim:"VELOCITY BODY Y,feet per second"`
	VelocityBodyZ     float64 `sim:"VELOCITY BODY Z,feet per second"`
	RotationVelocityX float64 `sim:"ROTATION VELOCITY BODY X,radians per second"`
	RotationVelocityY float64 `sim:"ROTATION VELOCITY BODY Y,radians per second"`
	RotationVelocityZ float64 `sim:"ROTATION VELOCITY BODY Z,radians per second"`
}

type simulatorStats struct {
	SlewActive bool    `sim:"IS SLEW ACTIVE,bool"`
	SimRate    float64 `sim:"SIMULATION RATE,number"`
}

type timeStats struct {
	ZuluTime       float64 `sim:"ZULU TIME,seconds"`
	ZuluDayOfYear  int32   `sim:"ZULU DAY OF YEAR,number"`
	LocalTime      float64 `sim:"LOCAL TIME,seconds"`
	LocalDayOfYear int32   `sim:"LOCAL DAY OF YEAR,number"`
	AbsoluteTime   float64 `sim:"ABSOLUTE TIME,seconds"`
}

type environmentStats struct {
	TotalWeight      float64 `sim:"TOTAL WEIGHT,pounds"`
	EmptyWeight      float64 `sim:"EMPTY WEIGHT,pounds"`
	AmbientDensity   float64 `sim:"AMBIENT DENSITY,slugs per cubic feet"`
	AmbientPressure  float64 `sim:"AMBIENT PRESSURE,inHg"`
	AmbientTemp      float64 `sim:"AMBIENT TEMPERATURE,celsius"`
	AmbientWindDir   float64 `sim:"AMBIENT WIND DIRECTION,degrees"`
	AmbientWindSpeed float64 `sim:"AMBIENT WIND VELOCITY,knots"`
}

type controlStats struct {
	ElevatorPosition  float64 `sim:"ELEVATOR POSITION,position"`
	AileronPosition   float64 `sim:"AILERON POSITION,position"`
	RudderPosition    float64 `sim:"RUDDER POSITION,position"`
	FlapsHandleIndex  int32   `sim:"FLAPS HANDLE INDEX,number"`
	FlapsHandlePct    float64 `sim:"FLAPS HANDLE PERCENT,percent"`
	SpoilersHandlePct float64 `sim:"SPOILERS HANDLE POSITION,percent"`
	GearHandle        bool    `sim:"GEAR HANDLE POSITION,bool"`
	GearCenter        float64 `sim:"GEAR CENTER POSITION,percent"`
	GearLeft          float64 `sim:"GEAR LEFT POSITION,percent"`
	GearRight         float64 `sim:"GEAR RIGHT POSITION,percent"`
	BrakeParking      bool    `sim:"BRAKE PARKING INDICATOR,bool"`
	BrakeLeft         float64 `sim:"BRAKE LEFT POSITION,position"`
	BrakeRight        float64 `sim:"BRAKE RIGHT POSITION,position"`
}

type engineStats struct {
	EngineType         int32   `sim:"ENGINE TYPE,enum"`
	EngineCount        int32   `sim:"NUMBER OF ENGINES,number"`
	Eng1Combustion     bool    `sim:"GENERAL ENG COMBUSTION:1,bool"`
	Eng2Combustion     bool    `sim:"GENERAL ENG COMBUSTION:2,bool"`
	Eng1ThrottleLever  float64 `sim:"GENERAL ENG THROTTLE LEVER POSITION:1,percent"`
	Eng2ThrottleLever  float64 `sim:"GENERAL ENG THROTTLE LEVER POSITION:2,percent"`
	Eng1RPM            float64 `sim:"GENERAL ENG RPM:1,rpm"`
	Eng2RPM            float64 `sim:"GENERAL ENG RPM:2,rpm"`
	Eng1N1             float64 `sim:"TURB ENG N1:1,percent"`
	Eng2N1             float64 `sim:"TURB ENG N1:2,percent"`
	Eng1N2             float64 `sim:"TURB ENG N2:1,percent"`
	Eng2N2             float64 `sim:"TURB ENG N2:2,percent"`
	Eng1ITT            float64 `sim:"TURB ENG ITT:1,celsius"`
	Eng2ITT            float64 `sim:"TURB ENG ITT:2,celsius"`
	Eng1OilPressure    float64 `sim:"ENG OIL PRESSURE:1,psi"`
	Eng2OilPressure    float64 `sim:"ENG OIL PRESSURE:2,psi"`
	Eng1OilTemperature float64 `sim:"ENG OIL TEMPERATURE:1,celsius"`
	Eng2OilTemperature float64 `sim:"ENG OIL TEMPERATURE:2,celsius"`
	Eng1FuelFlow       float64 `sim:"ENG FUEL FLOW GPH:1,gallons per hour"`
	Eng2FuelFlow       float64 `sim:"ENG FUEL FLOW GPH:2,gallons per hour"`
	Prop1RPM           float64 `sim:"PROP RPM:1,rpm"`
	Prop2RPM           float64 `sim:"PROP RPM:2,rpm"`
	Prop1Beta          float64 `sim:"PROP BETA:1,degrees"`
	Prop2Beta          float64 `sim:"PROP BETA:2,degrees"`
}

type fuelElectricalStats struct {
	FuelTotalQuantity  float64 `sim:"FUEL TOTAL QUANTITY,gallons"`
	FuelTotalCapacity  float64 `sim:"FUEL TOTAL CAPACITY,gallons"`
	FuelTotalWeight    float64 `sim:"FUEL TOTAL QUANTITY WEIGHT,pounds"`
	FuelLeftQuantity   float64 `sim:"FUEL LEFT QUANTITY,gallons"`
	FuelRightQuantity  float64 `sim:"FUEL RIGHT QUANTITY,gallons"`
	FuelCenterQuantity float64 `sim:"FUEL TANK CENTER QUANTITY,gallons"`
}

type electricalStats struct {
	ElectricalVoltage float64 `sim:"ELECTRICAL MAIN BUS VOLTAGE,volts"`
	ElectricalAmps    float64 `sim:"ELECTRICAL MAIN BUS AMPS,amperes"`
	AvionicsMaster    bool    `sim:"AVIONICS MASTER SWITCH,bool"`
	Alternator1       bool    `sim:"GENERAL ENG MASTER ALTERNATOR:1,bool"`
	Alternator2       bool    `sim:"GENERAL ENG MASTER ALTERNATOR:2,bool"`
}

type lightStats struct {
	LightBeacon  bool `sim:"LIGHT BEACON,bool"`
	LightLanding bool `sim:"LIGHT LANDING,bool"`
	LightTaxi    bool `sim:"LIGHT TAXI,bool"`
	LightNav     bool `sim:"LIGHT NAV,bool"`
	LightStrobe  bool `sim:"LIGHT STROBE,bool"`
}

type avionicsStats struct {
	AutopilotMaster       bool    `sim:"AUTOPILOT MASTER,bool"`
	AutopilotWingLeveler  bool    `sim:"AUTOPILOT WING LEVELER,bool"`
	AutopilotHeadingLock  bool    `sim:"AUTOPILOT HEADING LOCK,bool"`
	AutopilotHeadingValue float64 `sim:"AUTOPILOT HEADING LOCK DIR,degrees"`
	AutopilotAltitudeLock bool    `sim:"AUTOPILOT ALTITUDE LOCK,bool"`
	AutopilotAltitude     float64 `sim:"AUTOPILOT ALTITUDE LOCK VAR,feet"`
	AutopilotAirspeedHold bool    `sim:"AUTOPILOT AIRSPEED HOLD,bool"`
	AutopilotAirspeed     float64 `sim:"AUTOPILOT AIRSPEED HOLD VAR,knots"`
	AutopilotVerticalHold bool    `sim:"AUTOPILOT VERTICAL HOLD,bool"`
	AutopilotVerticalRate float64 `sim:"AUTOPILOT VERTICAL HOLD VAR,feet per minute"`
	AutopilotApproachHold bool    `sim:"AUTOPILOT APPROACH HOLD,bool"`
	Com1Frequency         float64 `sim:"COM ACTIVE FREQUENCY:1,megahertz"`
	Com2Frequency         float64 `sim:"COM ACTIVE FREQUENCY:2,megahertz"`
	Nav1Frequency         float64 `sim:"NAV ACTIVE FREQUENCY:1,megahertz"`
	Nav2Frequency         float64 `sim:"NAV ACTIVE FREQUENCY:2,megahertz"`
	TransponderCode       int32   `sim:"TRANSPONDER CODE:1,number"`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	client, err := simconnect.Open(ctx, "simconnect-go stats", simconnect.WithPollInterval(200*time.Millisecond))
	if err != nil {
		log.Fatalf("connect to MSFS 2024 SimConnect: %v", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("close SimConnect: %v", err)
		}
	}()

	go func() {
		for err := range client.Errors() {
			log.Printf("simconnect error: %v", err)
		}
	}()
	dbg := client.Debug()
	fac := client.Facilities()
	gen := client.System()
	inp := client.Input()

	positionDef := mustDefine[positionStats]("position")
	motionDef := mustDefine[motionStats]("motion")
	simulatorDef := mustDefine[simulatorStats]("simulator")
	timeDef := mustDefine[timeStats]("time")
	environmentDef := mustDefine[environmentStats]("environment")
	controlDef := mustDefine[controlStats]("controls")
	engineDef := mustDefine[engineStats]("engines")
	fuelElectricalDef := mustDefine[fuelElectricalStats]("fuel/electrical")
	electricalDef := mustDefine[electricalStats]("electrical")
	lightDef := mustDefine[lightStats]("lights")
	avionicsDef := mustDefine[avionicsStats]("avionics")

	fmt.Println("Connected to MSFS 2024. Printing simulator stats every 10 seconds.")
	printStats(ctx, client, fac, gen, inp, dbg, positionDef, motionDef, simulatorDef, timeDef, environmentDef, controlDef, engineDef, fuelElectricalDef, electricalDef, lightDef, avionicsDef)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stats stopped.")
			return
		case <-ticker.C:
			printStats(ctx, client, fac, gen, inp, dbg, positionDef, motionDef, simulatorDef, timeDef, environmentDef, controlDef, engineDef, fuelElectricalDef, electricalDef, lightDef, avionicsDef)
		}
	}
}

func mustDefine[T any](name string) *simvar.Definition[T] {
	def, err := simvar.Define[T]()
	if err != nil {
		log.Fatalf("define %s data: %v", name, err)
	}
	return def
}

func printStats(
	ctx context.Context,
	client *simconnect.Sim,
	fac *facilities.Facilities,
	gen *system.System,
	inp *input.Input,
	dbg *debug.Debug,
	positionDef *simvar.Definition[positionStats],
	motionDef *simvar.Definition[motionStats],
	simulatorDef *simvar.Definition[simulatorStats],
	timeDef *simvar.Definition[timeStats],
	environmentDef *simvar.Definition[environmentStats],
	controlDef *simvar.Definition[controlStats],
	engineDef *simvar.Definition[engineStats],
	fuelElectricalDef *simvar.Definition[fuelElectricalStats],
	electricalDef *simvar.Definition[electricalStats],
	lightDef *simvar.Definition[lightStats],
	avionicsDef *simvar.Definition[avionicsStats],
) {
	fmt.Printf("\n%s\n", strings.Repeat("=", 96))
	fmt.Printf("%s user aircraft simulator snapshot\n", time.Now().Format(time.RFC3339))

	requestAndPrint(ctx, client, "Position", positionDef, printPosition)
	printAirportSection(ctx, client, fac, positionDef)
	requestAndPrint(ctx, client, "Motion", motionDef, printMotion)
	printSimulatorSection(ctx, client, gen, simulatorDef, timeDef, environmentDef)
	requestAndPrint(ctx, client, "Controls / Surfaces", controlDef, printControls)
	requestAndPrint(ctx, client, "Engines", engineDef, printEngines)
	printFuelElectricalSection(ctx, client, fuelElectricalDef, electricalDef, lightDef)
	requestAndPrint(ctx, client, "Avionics / Radios", avionicsDef, printAvionics)
	printGetFetchAPISection(ctx, inp, dbg)
}

func requestAndPrint[T any](
	ctx context.Context,
	client *simconnect.Sim,
	title string,
	def *simvar.Definition[T],
	print func(T),
) {
	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	value, err := simvar.GetOnce(requestCtx, client, def, core.UserAircraft)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		log.Printf("request %s stats: %v", strings.ToLower(title), err)
		return
	}

	fmt.Printf("\n%s\n", title)
	print(value)
}

func requestQuiet[T any](ctx context.Context, client *simconnect.Sim, def *simvar.Definition[T]) (T, bool) {
	var zero T
	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	value, err := simvar.GetOnce(requestCtx, client, def, core.UserAircraft)
	if err != nil {
		return zero, false
	}
	return value, true
}

func printPosition(s positionStats) {
	fmt.Printf("  lat/lon:        %.6f, %.6f\n", s.Latitude, s.Longitude)
	fmt.Printf("  altitude:       %.1f ft MSL | %.1f ft AGL | ground %.1f ft\n", s.Altitude, s.AltitudeAGL, s.GroundAltitude)
	fmt.Printf("  attitude:       pitch %.2f deg | bank %.2f deg\n", s.Pitch, s.Bank)
	fmt.Printf("  heading:        true %.2f deg | magnetic %.2f deg | on ground %s\n", s.HeadingTrue, s.HeadingMagnetic, yesNo(s.OnGround))
}

func printMotion(s motionStats) {
	fmt.Printf("  airspeed:       IAS %.1f kt | TAS %.1f kt | Mach %.3f\n", s.IndicatedAirspeed, s.TrueAirspeed, s.Mach)
	fmt.Printf("  movement:       ground %.1f kt | vertical %.1f fpm | g %.2f\n", s.GroundSpeed, s.VerticalSpeed, s.GForce)
	fmt.Printf("  body velocity:  X %.2f | Y %.2f | Z %.2f ft/s\n", s.VelocityBodyX, s.VelocityBodyY, s.VelocityBodyZ)
	fmt.Printf("  rotation:       X %.4f | Y %.4f | Z %.4f rad/s\n", s.RotationVelocityX, s.RotationVelocityY, s.RotationVelocityZ)
}

func printSimulatorSection(
	ctx context.Context,
	client *simconnect.Sim,
	gen *system.System,
	simulatorDef *simvar.Definition[simulatorStats],
	timeDef *simvar.Definition[timeStats],
	environmentDef *simvar.Definition[environmentStats],
) {
	fmt.Printf("\nSimulator / Environment\n")
	if s, ok := requestQuiet(ctx, client, simulatorDef); ok {
		fmt.Printf("  state:          running %s | slew %s | rate %.2fx\n", simRunning(ctx, gen), yesNo(s.SlewActive), s.SimRate)
	} else {
		fmt.Printf("  state:          n/a\n")
	}
	if s, ok := requestQuiet(ctx, client, timeDef); ok {
		fmt.Printf("  time:           zulu %s day %d | local %s day %d\n", clock(s.ZuluTime), s.ZuluDayOfYear, clock(s.LocalTime), s.LocalDayOfYear)
		fmt.Printf("  absolute time:  %.0f s\n", s.AbsoluteTime)
	} else {
		fmt.Printf("  time:           n/a\n")
	}
	if s, ok := requestQuiet(ctx, client, environmentDef); ok {
		fmt.Printf("  weight:         total %.1f lb | empty %.1f lb\n", s.TotalWeight, s.EmptyWeight)
		fmt.Printf("  atmosphere:     %.1f C | %.2f inHg | density %.5f slug/ft3\n", s.AmbientTemp, s.AmbientPressure, s.AmbientDensity)
		fmt.Printf("  wind:           %.0f deg at %.1f kt\n", s.AmbientWindDir, s.AmbientWindSpeed)
	} else {
		fmt.Printf("  environment:    n/a\n")
	}
}

func printControls(s controlStats) {
	fmt.Printf("  primary:        elevator %.3f | aileron %.3f | rudder %.3f\n", s.ElevatorPosition, s.AileronPosition, s.RudderPosition)
	fmt.Printf("  flaps/spoilers: flaps index %d (%.1f%%) | spoilers %.1f%%\n", s.FlapsHandleIndex, s.FlapsHandlePct, s.SpoilersHandlePct)
	fmt.Printf("  gear:           handle %s | center %.1f%% | left %.1f%% | right %.1f%%\n", yesNo(s.GearHandle), s.GearCenter, s.GearLeft, s.GearRight)
	fmt.Printf("  brakes:         parking %s | left %.3f | right %.3f\n", yesNo(s.BrakeParking), s.BrakeLeft, s.BrakeRight)
}

func printEngines(s engineStats) {
	fmt.Printf("  type/count:     type %d | engines %d\n", s.EngineType, s.EngineCount)
	fmt.Printf("  combustion:     engine 1 %s | engine 2 %s\n", yesNo(s.Eng1Combustion), yesNo(s.Eng2Combustion))
	fmt.Printf("  throttle:       engine 1 %.1f%% | engine 2 %.1f%%\n", s.Eng1ThrottleLever, s.Eng2ThrottleLever)
	fmt.Printf("  rpm:            engine 1 %.0f | engine 2 %.0f | prop 1 %.0f | prop 2 %.0f\n", s.Eng1RPM, s.Eng2RPM, s.Prop1RPM, s.Prop2RPM)
	fmt.Printf("  turbine:        N1 %.1f/%.1f%% | N2 %.1f/%.1f%% | ITT %.1f/%.1f C\n", s.Eng1N1, s.Eng2N1, s.Eng1N2, s.Eng2N2, s.Eng1ITT, s.Eng2ITT)
	fmt.Printf("  oil:            pressure %.1f/%.1f psi | temp %.1f/%.1f C\n", s.Eng1OilPressure, s.Eng2OilPressure, s.Eng1OilTemperature, s.Eng2OilTemperature)
	fmt.Printf("  fuel/prop:      flow %.1f/%.1f gph | beta %.1f/%.1f deg\n", s.Eng1FuelFlow, s.Eng2FuelFlow, s.Prop1Beta, s.Prop2Beta)
}

func printFuelElectricalSection(
	ctx context.Context,
	client *simconnect.Sim,
	fuelDef *simvar.Definition[fuelElectricalStats],
	electricalDef *simvar.Definition[electricalStats],
	lightDef *simvar.Definition[lightStats],
) {
	fmt.Printf("\nFuel / Electrical / Lights\n")
	if s, ok := requestQuiet(ctx, client, fuelDef); ok {
		fmt.Printf("  fuel:           %.1f / %.1f gal | %.1f lb\n", s.FuelTotalQuantity, s.FuelTotalCapacity, s.FuelTotalWeight)
		fmt.Printf("  tanks:          left %.1f gal | right %.1f gal | center %.1f gal\n", s.FuelLeftQuantity, s.FuelRightQuantity, s.FuelCenterQuantity)
	} else {
		fmt.Printf("  fuel:           n/a\n")
	}
	if s, ok := requestQuiet(ctx, client, electricalDef); ok {
		fmt.Printf("  electrical:     main bus %.1f V | %.1f A | avionics %s | alternators %s/%s\n", s.ElectricalVoltage, s.ElectricalAmps, yesNo(s.AvionicsMaster), yesNo(s.Alternator1), yesNo(s.Alternator2))
	} else {
		fmt.Printf("  electrical:     n/a\n")
	}
	if s, ok := requestQuiet(ctx, client, lightDef); ok {
		fmt.Printf("  lights:         beacon %s | landing %s | taxi %s | nav %s | strobe %s\n", yesNo(s.LightBeacon), yesNo(s.LightLanding), yesNo(s.LightTaxi), yesNo(s.LightNav), yesNo(s.LightStrobe))
	} else {
		fmt.Printf("  lights:         n/a\n")
	}
}

func printAvionics(s avionicsStats) {
	fmt.Printf("  autopilot:      master %s | wing leveler %s | approach %s\n", yesNo(s.AutopilotMaster), yesNo(s.AutopilotWingLeveler), yesNo(s.AutopilotApproachHold))
	fmt.Printf("  ap modes:       hdg %s %.0f deg | alt %s %.0f ft | ias %s %.0f kt | vs %s %.0f fpm\n", yesNo(s.AutopilotHeadingLock), s.AutopilotHeadingValue, yesNo(s.AutopilotAltitudeLock), s.AutopilotAltitude, yesNo(s.AutopilotAirspeedHold), s.AutopilotAirspeed, yesNo(s.AutopilotVerticalHold), s.AutopilotVerticalRate)
	fmt.Printf("  radios:         COM %.3f / %.3f MHz | NAV %.3f / %.3f MHz | XPDR %04d\n", s.Com1Frequency, s.Com2Frequency, s.Nav1Frequency, s.Nav2Frequency, s.TransponderCode)
}

type airportCandidate struct {
	Ident          string
	Region         string
	Latitude       float64
	Longitude      float64
	AltitudeMeters float64
	DistanceMeters float64
}

func (a airportCandidate) DisplayIdent() string {
	if len(a.Ident) != 3 {
		return a.Ident
	}
	if a.Region != "" {
		return a.Region + a.Ident
	}
	if isContiguousUS(a.Latitude, a.Longitude) {
		return "K" + a.Ident
	}
	return a.Ident
}

func printAirportSection(ctx context.Context, client *simconnect.Sim, fac *facilities.Facilities, positionDef *simvar.Definition[positionStats]) {
	fmt.Printf("\nAirport\n")

	pos, ok := requestQuiet(ctx, client, positionDef)
	if !ok {
		fmt.Printf("  current:        n/a (aircraft position unavailable)\n")
		return
	}

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	airports, err := fac.RequestNearbyAirports(requestCtx)
	cancel()
	if err != nil {
		fmt.Printf("  current:        n/a (airport list unavailable: %v)\n", err)
		return
	}

	nearest, ok := nearestAirport(pos, airports)
	if !ok {
		fmt.Printf("  current:        n/a (no nearby airports returned)\n")
		return
	}

	status := "nearest"
	if pos.OnGround && nearest.DistanceMeters <= 5000 {
		status = "current"
	}
	fmt.Printf("  %s:        %s%s | %.0f m away | airport elev %.0f m\n",
		status,
		nearest.DisplayIdent(),
		formatAirportRegion(nearest.Region),
		nearest.DistanceMeters,
		nearest.AltitudeMeters,
	)
	if status != "current" {
		fmt.Printf("  current:        unknown (not on ground or nearest airport is over 5 km away)\n")
	}
}

func printGetFetchAPISection(ctx context.Context, inp *input.Input, dbg *debug.Debug) {
	fmt.Printf("\nGET / FETCH API Probes\n")

	if id, err := dbg.LastSentPacketID(); err == nil {
		fmt.Printf("  last packet id: %d\n", id)
	} else {
		fmt.Printf("  last packet id: n/a (%v)\n", err)
	}

	if times, err := dbg.RequestResponseTimes(5); err == nil {
		fmt.Printf("  response times: %s\n", formatResponseTimes(times))
	} else {
		fmt.Printf("  response times: n/a (%v)\n", err)
	}

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	events, err := inp.EnumerateInputEvents(requestCtx)
	cancel()
	if err != nil {
		fmt.Printf("  input events:   n/a (%v)\n", err)
		return
	}
	fmt.Printf("  input events:   page %d/%d | %d entries\n", events.EntryNumber+1, events.OutOf, len(events.Events))

	event, ok := firstInputEvent(events.Events)
	if !ok {
		fmt.Printf("  input get:      n/a (no events in first page)\n")
		return
	}
	fmt.Printf("  input sample:   %s | hash %d | type %s\n", event.Name, event.Hash, inputEventTypeName(event.Type))

	requestCtx, cancel = context.WithTimeout(ctx, 5*time.Second)
	value, err := inp.GetInputEvent(requestCtx, event.Hash)
	cancel()
	if err != nil {
		fmt.Printf("  input get:      n/a (%v)\n", err)
		return
	}
	fmt.Printf("  input get:      %s\n", formatInputEventValue(value))
}

func firstInputEvent(events []core.InputEventDescriptor) (core.InputEventDescriptor, bool) {
	for _, event := range events {
		if event.Hash != 0 {
			return event, true
		}
	}
	return core.InputEventDescriptor{}, false
}

func formatResponseTimes(times []float32) string {
	if len(times) == 0 {
		return "none"
	}
	parts := make([]string, len(times))
	for i, value := range times {
		parts[i] = fmt.Sprintf("%.4fs", value)
	}
	return strings.Join(parts, ", ")
}

func inputEventTypeName(t core.InputEventType) string {
	switch t {
	case core.InputEventTypeDouble:
		return "double"
	case core.InputEventTypeString:
		return "string"
	default:
		return fmt.Sprintf("type %d", t)
	}
}

func formatInputEventValue(value core.InputEventValueMessage) string {
	switch value.Type {
	case core.InputEventTypeDouble:
		return fmt.Sprintf("%.6f", value.Double)
	case core.InputEventTypeString:
		return fmt.Sprintf("%q", value.String)
	default:
		return fmt.Sprintf("%d raw bytes", len(value.Payload))
	}
}

func nearestAirport(pos positionStats, airports core.AirportListMessage) (airportCandidate, bool) {
	var best airportCandidate
	for _, airport := range airports.Airports {
		candidate := airportCandidate{
			Ident:          airport.Ident,
			Region:         airport.Region,
			Latitude:       airport.Latitude,
			Longitude:      airport.Longitude,
			AltitudeMeters: airport.Altitude,
		}
		if candidate.Ident == "" {
			continue
		}
		candidate.DistanceMeters = greatCircleDistanceMeters(pos.Latitude, pos.Longitude, candidate.Latitude, candidate.Longitude)
		if best.Ident == "" || candidate.DistanceMeters < best.DistanceMeters {
			best = candidate
		}
	}
	return best, best.Ident != ""
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

func fixedBytesString(b []byte) string {
	if n := strings.IndexByte(string(b), 0); n >= 0 {
		return string(b[:n])
	}
	return string(b)
}

func formatAirportRegion(region string) string {
	if region == "" {
		return ""
	}
	return " / " + region
}

func isContiguousUS(lat, lon float64) bool {
	return lat >= 24 && lat <= 50 && lon >= -125 && lon <= -66
}

func yesNo(v bool) string {
	if !v {
		return "no"
	}
	return "yes"
}

func simRunning(ctx context.Context, gen *system.System) string {
	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	state, err := gen.RequestSystemState(requestCtx, "Sim")
	if err != nil {
		return "n/a"
	}
	if state.Integer == 0 {
		return "no"
	}
	return "yes"
}

func clock(seconds float64) string {
	if seconds < 0 {
		seconds = 0
	}
	seconds = float64(int64(seconds) % 86400)
	h := int(seconds) / 3600
	m := int(seconds) % 3600 / 60
	s := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
