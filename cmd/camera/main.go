//go:build windows

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/camera"
	simconnect "github.com/Zwergpro/simconnect-go/pkg/simconnect/client"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 20*time.Second, "maximum time to run the Camera API checker")
	moveCamera := flag.Bool("move-camera", false, "exercise CameraSet and CameraSetUsingDefinition; this visibly moves the add-on camera")
	lockWorld := flag.Bool("world-locker", false, "exercise RequestCameraWorldLocker at the current camera position, then delete it")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	log.SetFlags(0)

	fmt.Println("SimConnect Camera API checker")
	fmt.Printf("started: %s\n", time.Now().Format(time.RFC3339))

	client, err := simconnect.Open(ctx, "simconnect-go camera checker", simconnect.WithPollInterval(25*time.Millisecond))
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
	cam := client.Camera()

	statusCh, stopStatus := checkSubscribeCameraStatus(ctx, cam)
	lockerCh, stopLocker := checkSubscribeWorldLockerStatus(ctx, cam)

	printCameraStatus(ctx, cam, "initial")
	initialData, hasInitialData := checkCameraGet(ctx, cam)
	definitions := checkEnumerateCameraDefinitions(ctx, cam)
	checkCameraAcquire(ctx, cam, statusCh)
	printCameraStatus(ctx, cam, "after acquire")
	currentData, hasCurrentData := checkCameraGet(ctx, cam)
	checkCameraFlags(cam)

	if *moveCamera {
		if hasCurrentData {
			checkCameraSet(cam, currentData.CameraData)
		} else if hasInitialData {
			checkCameraSet(cam, initialData.CameraData)
		} else {
			fmt.Println("\nSimConnect_CameraSet")
			fmt.Println("  skipped: no camera data was available to echo back")
		}
		checkCameraSetUsingDefinition(cam, definitions)
	} else {
		fmt.Println("\nCamera Mutations")
		fmt.Println("  skipped CameraSet and CameraSetUsingDefinition; pass -move-camera to exercise them")
	}

	if *lockWorld {
		checkWorldLocker(ctx, cam, lockerCh, currentCameraData(initialData, hasInitialData, currentData, hasCurrentData))
	} else {
		fmt.Println("\nWorld Locker")
		fmt.Println("  skipped RequestCameraWorldLocker and DeleteCameraWorldLocker; pass -world-locker to exercise them")
	}

	checkCameraRelease(cam, definitions)

	stopStatus()
	readCameraStatusMessage(ctx, statusCh, "after status unsubscribe", 750*time.Millisecond)
	stopLocker()
	readWorldLockerMessage(ctx, lockerCh, "after world-locker unsubscribe", 750*time.Millisecond)
}

func checkSubscribeCameraStatus(ctx context.Context, cam *camera.Camera) (<-chan core.CameraStatusMessage, context.CancelFunc) {
	fmt.Println("\nSimConnect_SubscribeToCameraStatusUpdate")

	subCtx, cancel := context.WithCancel(ctx)
	ch, err := cam.SubscribeStatus(subCtx)
	if err != nil {
		cancel()
		fmt.Printf("  error: %s\n", formatError(err))
		return nil, func() {}
	}
	fmt.Println("  subscribed")
	return ch, func() {
		cancel()
		if err := cam.UnsubscribeStatus(); err != nil {
			fmt.Printf("  SimConnect_UnsubscribeToCameraStatusUpdate: %s\n", formatError(err))
		} else {
			fmt.Println("\nSimConnect_UnsubscribeToCameraStatusUpdate")
			fmt.Println("  unsubscribed")
		}
	}
}

func checkSubscribeWorldLockerStatus(ctx context.Context, cam *camera.Camera) (<-chan core.CameraWorldLockerMessage, context.CancelFunc) {
	fmt.Println("\nSimConnect_SubscribeToCameraWorldLockerStatusUpdate")

	subCtx, cancel := context.WithCancel(ctx)
	ch, err := cam.SubscribeWorldLockerStatus(subCtx)
	if err != nil {
		cancel()
		fmt.Printf("  error: %s\n", formatError(err))
		return nil, func() {}
	}
	fmt.Println("  subscribed")
	return ch, func() {
		cancel()
		if err := cam.UnsubscribeWorldLockerStatus(); err != nil {
			fmt.Printf("  SimConnect_UnsubscribeToCameraWorldLockerStatusUpdate: %s\n", formatError(err))
		} else {
			fmt.Println("\nSimConnect_UnsubscribeToCameraWorldLockerStatusUpdate")
			fmt.Println("  unsubscribed")
		}
	}
}

func printCameraStatus(ctx context.Context, cam *camera.Camera, label string) {
	fmt.Println("\nSimConnect_CameraGetStatus")

	requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	status, err := cam.GetStatus(requestCtx)
	cancel()
	if err != nil {
		fmt.Printf("  %s: %s\n", label, formatError(err))
		return
	}
	fmt.Printf("  %s: acquired=%s gameControlled=%t\n", label, cameraAvailabilityName(status.AcquiredState), status.GameControlled)
}

func checkCameraGet(ctx context.Context, cam *camera.Camera) (core.CameraDataMessage, bool) {
	fmt.Println("\nSimConnect_CameraGet")

	referentials := []struct {
		name string
		ref  core.PositionReferential
	}{
		{"SIMOBJECT", core.PositionReferentialSimObject},
		{"WORLD", core.PositionReferentialWorld},
		{"EYEPOINT", core.PositionReferentialEyepoint},
	}

	for _, item := range referentials {
		requestCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
		msg, err := cam.Get(requestCtx, item.ref)
		cancel()
		if err != nil {
			fmt.Printf("  %-9s error: %s\n", item.name, formatError(err))
			continue
		}
		fmt.Printf("  %-9s %s\n", item.name, formatCameraData(msg.CameraData))
		return msg, true
	}

	return core.CameraDataMessage{}, false
}

func checkEnumerateCameraDefinitions(ctx context.Context, cam *camera.Camera) core.CameraDefinitionListMessage {
	fmt.Println("\nSimConnect_EnumerateCameraDefinitions")

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	definitions, err := cam.EnumerateDefinitions(requestCtx)
	cancel()
	if err != nil {
		fmt.Printf("  error: %s\n", formatError(err))
		return core.CameraDefinitionListMessage{}
	}

	fmt.Printf("  page=%d/%d entries=%d\n", definitions.EntryNumber+1, definitions.OutOf, len(definitions.Definitions))
	for i, name := range definitions.Definitions {
		if i >= 12 {
			fmt.Printf("  ... %d more\n", len(definitions.Definitions)-i)
			break
		}
		fmt.Printf("  [%d] %q\n", i, name)
	}
	return definitions
}

func checkCameraAcquire(ctx context.Context, cam *camera.Camera, statusCh <-chan core.CameraStatusMessage) {
	fmt.Println("\nSimConnect_CameraAcquire")

	drainCameraStatusMessages(statusCh)

	clientID := fmt.Sprintf("simconnect-go.camera.%d", time.Now().UnixNano())
	if err := cam.Acquire(clientID); err != nil {
		fmt.Printf("  clientID=%q error: %s\n", clientID, formatError(err))
		return
	}
	fmt.Printf("  clientID=%q request sent\n", clientID)
	readCameraStatusUntil(ctx, statusCh, "acquire notification", 3*time.Second, core.CameraAcquired)
}

func checkCameraFlags(cam *camera.Camera) {
	fmt.Println("\nSimConnect_CameraEnableFlag / CameraDisableFlag")

	flags := []struct {
		name string
		flag core.CameraFlag
	}{
		{"INTERACTION", core.CameraFlagInteraction},
		{"ABOVE_GROUND", core.CameraFlagAboveGround},
	}

	for _, item := range flags {
		if err := cam.EnableFlag(item.flag); err != nil {
			fmt.Printf("  enable %-12s error: %s\n", item.name, formatError(err))
		} else {
			fmt.Printf("  enable %-12s ok\n", item.name)
		}
		if err := cam.DisableFlag(item.flag); err != nil {
			fmt.Printf("  disable %-11s error: %s\n", item.name, formatError(err))
		} else {
			fmt.Printf("  disable %-11s ok\n", item.name)
		}
	}
}

func checkCameraSet(cam *camera.Camera, base core.CameraData) {
	fmt.Println("\nSimConnect_CameraSet")

	data := base
	if data.PositionReferential == core.PositionReferentialNone {
		data.PositionReferential = core.PositionReferentialSimObject
	}
	if data.RotationReferential == core.PositionReferentialNone {
		data.RotationReferential = data.PositionReferential
	}
	if data.FOV <= 0 || math.IsNaN(data.FOV) || math.IsInf(data.FOV, 0) {
		data.FOV = 0.8
	}

	err := cam.Set(data, core.CameraDataMaskAllRotation)
	if err != nil {
		fmt.Printf("  echo current data: %s\n", formatError(err))
		return
	}
	fmt.Printf("  echoed current data with mask=%s %s\n", cameraDataMaskName(core.CameraDataMaskAllRotation), formatCameraData(data))
}

func checkCameraSetUsingDefinition(cam *camera.Camera, definitions core.CameraDefinitionListMessage) {
	fmt.Println("\nSimConnect_CameraSetUsingDefinition")

	name, ok := chooseCameraDefinition(definitions.Definitions)
	if !ok {
		fmt.Println("  skipped: no camera definitions returned")
		return
	}
	if err := cam.SetUsingDefinition(name); err != nil {
		fmt.Printf("  definition=%q error: %s\n", name, formatError(err))
		return
	}
	fmt.Printf("  definition=%q ok\n", name)
}

func checkWorldLocker(ctx context.Context, cam *camera.Camera, lockerCh <-chan core.CameraWorldLockerMessage, data core.CameraData) {
	fmt.Println("\nSimConnect_RequestCameraWorldLocker")

	position := data.Position
	referential := data.PositionReferential
	objectID := data.PositionReferentialObjectID
	if referential == core.PositionReferentialNone {
		referential = core.PositionReferentialSimObject
		objectID = core.UserAircraft
	}

	requestCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	msg, err := cam.RequestWorldLocker(requestCtx, position, referential, objectID)
	cancel()
	if err != nil {
		fmt.Printf("  position=%s referential=%s objectID=%d error: %s\n", formatXYZ(position), referentialName(referential), objectID, formatError(err))
	} else {
		fmt.Printf("  position=%s referential=%s objectID=%d status=%s\n", formatXYZ(position), referentialName(referential), objectID, worldLockerStatusName(msg.Status))
	}
	readWorldLockerMessage(ctx, lockerCh, "locker notification", 2*time.Second)

	fmt.Println("\nSimConnect_DeleteCameraWorldLocker")
	if err := cam.DeleteWorldLocker(); err != nil {
		fmt.Printf("  error: %s\n", formatError(err))
		return
	}
	fmt.Println("  delete requested")
	readWorldLockerMessage(ctx, lockerCh, "delete notification", 2*time.Second)
}

func checkCameraRelease(cam *camera.Camera, definitions core.CameraDefinitionListMessage) {
	fmt.Println("\nSimConnect_CameraRelease")

	name, ok := chooseCameraDefinition(definitions.Definitions)
	if !ok {
		name = ""
	}
	if err := cam.Release(name); err != nil {
		fmt.Printf("  definition=%q error: %s\n", name, formatError(err))
		return
	}
	fmt.Printf("  definition=%q ok\n", name)
}

func readCameraStatusMessage(ctx context.Context, ch <-chan core.CameraStatusMessage, label string, wait time.Duration) {
	if ch == nil {
		fmt.Printf("  %s: no subscription\n", label)
		return
	}

	timer := time.NewTimer(wait)
	defer timer.Stop()

	select {
	case msg, ok := <-ch:
		if !ok {
			fmt.Printf("  %s: channel closed\n", label)
			return
		}
		fmt.Printf("  %s: acquired=%s gameControlled=%t\n", label, cameraAvailabilityName(msg.AcquiredState), msg.GameControlled)
	case <-timer.C:
		fmt.Printf("  %s: no message in %s\n", label, wait)
	case <-ctx.Done():
		fmt.Printf("  %s: %v\n", label, ctx.Err())
	}
}

func readCameraStatusUntil(ctx context.Context, ch <-chan core.CameraStatusMessage, label string, wait time.Duration, want core.CameraAvailability) {
	if ch == nil {
		fmt.Printf("  %s: no subscription\n", label)
		return
	}

	timer := time.NewTimer(wait)
	defer timer.Stop()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Printf("  %s: channel closed\n", label)
				return
			}
			fmt.Printf("  %s: acquired=%s gameControlled=%t\n", label, cameraAvailabilityName(msg.AcquiredState), msg.GameControlled)
			if msg.AcquiredState == want {
				return
			}
		case <-timer.C:
			fmt.Printf("  %s: no %s message in %s\n", label, cameraAvailabilityName(want), wait)
			return
		case <-ctx.Done():
			fmt.Printf("  %s: %v\n", label, ctx.Err())
			return
		}
	}
}

func drainCameraStatusMessages(ch <-chan core.CameraStatusMessage) {
	for {
		select {
		case <-ch:
		default:
			return
		}
	}
}

func readWorldLockerMessage(ctx context.Context, ch <-chan core.CameraWorldLockerMessage, label string, wait time.Duration) {
	if ch == nil {
		fmt.Printf("  %s: no subscription\n", label)
		return
	}

	timer := time.NewTimer(wait)
	defer timer.Stop()

	select {
	case msg, ok := <-ch:
		if !ok {
			fmt.Printf("  %s: channel closed\n", label)
			return
		}
		fmt.Printf("  %s: status=%s\n", label, worldLockerStatusName(msg.Status))
	case <-timer.C:
		fmt.Printf("  %s: no message in %s\n", label, wait)
	case <-ctx.Done():
		fmt.Printf("  %s: %v\n", label, ctx.Err())
	}
}

func currentCameraData(initial core.CameraDataMessage, hasInitial bool, current core.CameraDataMessage, hasCurrent bool) core.CameraData {
	if hasCurrent {
		return current.CameraData
	}
	if hasInitial {
		return initial.CameraData
	}
	return core.CameraData{
		PositionReferential:         core.PositionReferentialSimObject,
		PositionReferentialObjectID: core.UserAircraft,
		RotationReferential:         core.PositionReferentialSimObject,
		RotationReferentialObjectID: core.UserAircraft,
		FOV:                         0.8,
	}
}

func chooseCameraDefinition(definitions []string) (string, bool) {
	for _, name := range definitions {
		if name != "" {
			return name, true
		}
	}
	return "", false
}

func formatCameraData(data core.CameraData) string {
	return fmt.Sprintf(
		"pos=%s posRef=%s posObj=%d target=%s pbh={pitch=%.3f bank=%.3f heading=%.3f} rotRef=%s rotObj=%d fov=%.6f",
		formatXYZ(data.Position),
		referentialName(data.PositionReferential),
		data.PositionReferentialObjectID,
		formatXYZ(data.TargetedPos),
		data.PBH.Pitch,
		data.PBH.Bank,
		data.PBH.Heading,
		referentialName(data.RotationReferential),
		data.RotationReferentialObjectID,
		data.FOV,
	)
}

func formatXYZ(v core.XYZ) string {
	return fmt.Sprintf("{x=%.6f y=%.6f z=%.6f}", v.X, v.Y, v.Z)
}

func referentialName(ref core.PositionReferential) string {
	switch ref {
	case core.PositionReferentialNone:
		return "NONE"
	case core.PositionReferentialSimObject:
		return "SIMOBJECT"
	case core.PositionReferentialWorld:
		return "WORLD"
	case core.PositionReferentialEyepoint:
		return "EYEPOINT"
	case core.PositionReferentialSimObjectDatum:
		return "SIMOBJECT_DATUM"
	default:
		return fmt.Sprintf("PositionReferential(%d)", ref)
	}
}

func cameraAvailabilityName(state core.CameraAvailability) string {
	switch state {
	case core.CameraNotAcquired:
		return "NOT_ACQUIRED"
	case core.CameraAcquired:
		return "ACQUIRED"
	case core.CameraAcquiredByOther:
		return "ACQUIRED_BY_OTHER"
	case core.CameraUserDisabled:
		return "USER_DISABLED"
	default:
		return fmt.Sprintf("CameraAvailability(%d)", state)
	}
}

func worldLockerStatusName(status core.CameraWorldLockerStatus) string {
	switch status {
	case core.CameraWorldLockerStatusNone:
		return "NONE"
	case core.CameraWorldLockerStatusStart:
		return "START"
	case core.CameraWorldLockerStatusSuccess:
		return "SUCCESS"
	case core.CameraWorldLockerStatusCancel:
		return "CANCEL"
	case core.CameraWorldLockerStatusFail:
		return "FAIL"
	default:
		return fmt.Sprintf("CameraWorldLockerStatus(%d)", status)
	}
}

func cameraDataMaskName(mask core.CameraDataMask) string {
	switch mask {
	case core.CameraDataMaskNone:
		return "NONE"
	case core.CameraDataMaskPosition:
		return "POSITION"
	case core.CameraDataMaskRotation:
		return "ROTATION"
	case core.CameraDataMaskTargeted:
		return "TARGETED"
	case core.CameraDataMaskFOV:
		return "FOV"
	case core.CameraDataMaskAllRotation:
		return "ALL_ROTATION"
	case core.CameraDataMaskAllTargeted:
		return "ALL_TARGETED"
	default:
		return fmt.Sprintf("CameraDataMask(0x%X)", uint32(mask))
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
