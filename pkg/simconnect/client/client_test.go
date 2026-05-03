//go:build windows

package client

import (
	"context"
	"encoding/binary"
	"math"
	"testing"
	"time"
	"unsafe"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

func TestIDAllocator(t *testing.T) {
	ids := newIDAllocator(10)
	if got := ids.Next(); got != 10 {
		t.Fatalf("first id = %d, want 10", got)
	}
	if got := ids.Next(); got != 11 {
		t.Fatalf("second id = %d, want 11", got)
	}
}

func TestOptions(t *testing.T) {
	cfg := defaultClientConfig()
	WithPollInterval(time.Second)(&cfg)
	WithChannelBuffer(32)(&cfg)
	WithManualDispatch()(&cfg)
	WithWindowHandle(1)(&cfg)
	WithEventID(2)(&cfg)
	WithEventHandle(3)(&cfg)
	WithConfigIndex(4)(&cfg)

	if cfg.pollInterval != time.Second || cfg.channelBuffer != 32 || !cfg.manualDispatch {
		t.Fatalf("unexpected option config: %+v", cfg)
	}
	if cfg.hwnd != 1 || cfg.eventID != 2 || cfg.eventHandle != 3 || cfg.configIndex != 4 {
		t.Fatalf("unexpected raw open config: %+v", cfg)
	}
}

func TestFixedString(t *testing.T) {
	if got := fixedString([]byte{'M', 'S', 'F', 'S', 0, 'x'}); got != "MSFS" {
		t.Fatalf("fixedString = %q", got)
	}
}

func TestDecodeMessages(t *testing.T) {
	open := bindings.SIMCONNECT_RECV_OPEN{}
	open.DwID = uint32(bindings.SIMCONNECT_RECV_ID_OPEN)
	copy(open.SzApplicationName[:], "Flight Simulator")

	msg, err := decodeMessage(&open.SIMCONNECT_RECV, uint32(unsafe.Sizeof(open)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.OpenMessage).ApplicationName; got != "Flight Simulator" {
		t.Fatalf("application name = %q", got)
	}

	event := bindings.SIMCONNECT_RECV_EVENT{}
	event.DwID = uint32(bindings.SIMCONNECT_RECV_ID_EVENT)
	event.UEventID = 42
	event.DwData = 7
	msg, err = decodeMessage(&event.SIMCONNECT_RECV, uint32(unsafe.Sizeof(event)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.ClientEvent); got.EventID != 42 || got.Data != 7 {
		t.Fatalf("event = %+v", got)
	}
	if got := msg.(core.ClientEvent).DataValues[0]; got != 7 {
		t.Fatalf("event data values[0] = %d", got)
	}

	eventEX1 := bindings.SIMCONNECT_RECV_EVENT_EX1{}
	eventEX1.DwID = uint32(bindings.SIMCONNECT_RECV_ID_EVENT_EX1)
	eventEX1.UEventID = 43
	eventEX1.DwData0 = 1
	eventEX1.DwData4 = 5
	msg, err = decodeMessage(&eventEX1.SIMCONNECT_RECV, uint32(unsafe.Sizeof(eventEX1)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.ClientEventEX1); got.EventID != 43 || got.Data != 1 || got.DataValues[4] != 5 {
		t.Fatalf("event ex1 = %+v", got)
	}

	exception := bindings.SIMCONNECT_RECV_EXCEPTION{}
	exception.DwID = uint32(bindings.SIMCONNECT_RECV_ID_EXCEPTION)
	exception.DwException = uint32(bindings.SIMCONNECT_EXCEPTION_NAME_UNRECOGNIZED)
	exception.DwSendID = 99
	msg, err = decodeMessage(&exception.SIMCONNECT_RECV, uint32(unsafe.Sizeof(exception)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.ExceptionError); got.SendID != 99 || got.Exception != core.ExceptionNameUnrecognized {
		t.Fatalf("exception = %+v", got)
	}
}

func TestDecodeGeneralMessages(t *testing.T) {
	filename := bindings.SIMCONNECT_RECV_EVENT_FILENAME{}
	filename.DwID = uint32(bindings.SIMCONNECT_RECV_ID_EVENT_FILENAME)
	filename.UEventID = 10
	filename.DwData = 1
	filename.DwFlags = 2
	copy(filename.SzFileName[:], `C:\Flights\Test.FLT`)

	msg, err := decodeMessage(&filename.SIMCONNECT_RECV, uint32(unsafe.Sizeof(filename)))
	if err != nil {
		t.Fatal(err)
	}
	fileEvent := msg.(core.FilenameEvent)
	if fileEvent.EventID != 10 || fileEvent.FileName != `C:\Flights\Test.FLT` || fileEvent.ClientEvent.FileName != fileEvent.FileName {
		t.Fatalf("filename event = %+v", fileEvent)
	}

	frame := bindings.SIMCONNECT_RECV_EVENT_FRAME{}
	frame.DwID = uint32(bindings.SIMCONNECT_RECV_ID_EVENT_FRAME)
	frame.UEventID = 11
	frame.FFrameRate = 60
	frame.FSimSpeed = 2
	msg, err = decodeMessage(&frame.SIMCONNECT_RECV, uint32(unsafe.Sizeof(frame)))
	if err != nil {
		t.Fatal(err)
	}
	frameEvent := msg.(core.FrameEvent)
	if frameEvent.EventID != 11 || frameEvent.FrameRate != 60 || frameEvent.ClientEvent.SimSpeed != 2 {
		t.Fatalf("frame event = %+v", frameEvent)
	}

	systemState := bindings.SIMCONNECT_RECV_SYSTEM_STATE{}
	systemState.DwID = uint32(bindings.SIMCONNECT_RECV_ID_SYSTEM_STATE)
	systemState.DwRequestID = 12
	systemState.DwInteger = 1
	copy(systemState.SzString[:], "running")
	msg, err = decodeMessage(&systemState.SIMCONNECT_RECV, uint32(unsafe.Sizeof(systemState)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.SystemStateMessage); got.RequestID != 12 || got.Integer != 1 || got.String != "running" {
		t.Fatalf("system state = %+v", got)
	}

	assigned := bindings.SIMCONNECT_RECV_ASSIGNED_OBJECT_ID{}
	assigned.DwID = uint32(bindings.SIMCONNECT_RECV_ID_ASSIGNED_OBJECT_ID)
	assigned.DwRequestID = 14
	assigned.DwObjectID = 99
	msg, err = decodeMessage(&assigned.SIMCONNECT_RECV, uint32(unsafe.Sizeof(assigned)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.AssignedObjectIDMessage); got.RequestID != 14 || got.ObjectID != 99 {
		t.Fatalf("assigned object id = %+v", got)
	}

	action := bindings.SIMCONNECT_RECV_ACTION_CALLBACK{}
	action.DwID = uint32(bindings.SIMCONNECT_RECV_ID_ACTION_CALLBACK)
	action.CbRequestId = 13
	copy(action.SzActionID[:], "MyAction")
	msg, err = decodeMessage(&action.SIMCONNECT_RECV, uint32(unsafe.Sizeof(action)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.ActionCallbackMessage); got.RequestID != 13 || got.ActionID != "MyAction" {
		t.Fatalf("action callback = %+v", got)
	}
}

func TestDecodeSimObjectDataPayload(t *testing.T) {
	payload := make([]byte, 8)
	binary.LittleEndian.PutUint64(payload, math.Float64bits(123.5))
	raw, size := simObjectPacket(12, payload)

	msg, err := decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	data := msg.(core.SimObjectDataMessage)
	if data.RequestID != 12 {
		t.Fatalf("request id = %d", data.RequestID)
	}
	if got := math.Float64frombits(binary.LittleEndian.Uint64(data.Payload)); got != 123.5 {
		t.Fatalf("payload value = %f", got)
	}

	raw, size = simObjectDataByTypePacket(13, 99, payload)
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	byType := msg.(core.SimObjectDataByTypeMessage)
	if byType.RequestID != 13 || byType.ObjectID != 99 {
		t.Fatalf("by-type data = %+v", byType)
	}

	raw, size = clientDataPacket(14, payload)
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	clientData := msg.(core.ClientDataMessage)
	if clientData.RequestID != 14 || len(clientData.Payload) != 8 {
		t.Fatalf("client data = %+v", clientData)
	}
}

func TestDecodeEventsAndDataMessages(t *testing.T) {
	reserved := bindings.SIMCONNECT_RECV_RESERVED_KEY{}
	reserved.DwID = uint32(bindings.SIMCONNECT_RECV_ID_RESERVED_KEY)
	copy(reserved.SzChoiceReserved[:], "Ctrl+Shift+A")
	copy(reserved.SzReservedKey[:], "Ctrl+Shift+A")
	msg, err := decodeMessage(&reserved.SIMCONNECT_RECV, uint32(unsafe.Sizeof(reserved)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.ReservedKeyMessage); got.ChoiceReserved != "Ctrl+Shift+A" || got.ReservedKey != "Ctrl+Shift+A" {
		t.Fatalf("reserved key = %+v", got)
	}

	raw, size := simObjectLiveryListPacket(41, []core.SimObjectLivery{{AircraftTitle: "Cessna 172", LiveryName: "Classic"}})
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	liveries := msg.(core.SimObjectLiveryListMessage)
	if liveries.RequestID != 41 || len(liveries.Liveries) != 1 || liveries.Liveries[0].LiveryName != "Classic" {
		t.Fatalf("liveries = %+v", liveries)
	}

	flow := bindings.SIMCONNECT_RECV_FLOW_EVENT{}
	flow.DwID = uint32(bindings.SIMCONNECT_RECV_ID_FLOW_EVENT)
	flow.FlowEvent = bindings.SIMCONNECT_FLOW_EVENT_FLT_LOADED
	copy(flow.FltPath[:], `C:\Flights\Test.FLT`)
	msg, err = decodeMessage(&flow.SIMCONNECT_RECV, uint32(unsafe.Sizeof(flow)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.FlowEventMessage); got.Event != core.FlowEventFLTLoaded || got.FLTPath != `C:\Flights\Test.FLT` {
		t.Fatalf("flow event = %+v", got)
	}

	raw, size = commBusPacket(66, []byte("hello\x00"))
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	commBus := msg.(core.CommBusMessage)
	if commBus.EventID != 66 || commBus.Data != "hello" || string(commBus.Payload) != "hello\x00" {
		t.Fatalf("comm bus = %+v", commBus)
	}
}

func TestDecodeFacilityMessages(t *testing.T) {
	raw, size := airportListPacket(21, []string{"KSEA", "KLAX"})

	msg, err := decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	airports := msg.(core.AirportListMessage)
	if airports.RequestID != 21 || airports.ArraySize != 2 || len(airports.Airports) != 2 {
		t.Fatalf("airport list = %+v", airports)
	}
	if got := airports.Airports[1].Ident; got != "KLAX" {
		t.Fatalf("second airport = %q", got)
	}

	facilityPayload := []byte{1, 2, 3, 4, 5}
	raw, size = facilityDataPacket(22, bindings.SIMCONNECT_FACILITY_DATA_RUNWAY, facilityPayload)
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	facilityData := msg.(core.FacilityDataMessage)
	if facilityData.UserRequestID != 22 || facilityData.Type != core.FacilityDataRunway {
		t.Fatalf("facility data header = %+v", facilityData)
	}
	if string(facilityData.Payload) != string(facilityPayload) {
		t.Fatalf("facility payload = %v", facilityData.Payload)
	}
}

func TestDecodeInputEventMessages(t *testing.T) {
	raw, size := inputEventListPacket(31, []string{"ENGINE_START"})

	msg, err := decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	events := msg.(core.InputEventListMessage)
	if events.RequestID != 31 || len(events.Events) != 1 {
		t.Fatalf("input events = %+v", events)
	}
	if events.Events[0].Name != "ENGINE_START" || events.Events[0].Hash != 12345 {
		t.Fatalf("input event descriptor = %+v", events.Events[0])
	}

	raw, size = inputEventDoublePacket(32, 42.25)
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	value := msg.(core.InputEventValueMessage)
	if value.RequestID != 32 || value.Type != core.InputEventTypeDouble || value.Double != 42.25 {
		t.Fatalf("input event value = %+v", value)
	}

	raw, size = inputEventStringPacket(33, "ARMED")
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	stringValue := msg.(core.InputEventValueMessage)
	if stringValue.RequestID != 33 || stringValue.String != "ARMED" {
		t.Fatalf("input event string = %+v", stringValue)
	}

	raw, size = inputEventSubscriptionPacket(12345, 7.5)
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	sub := msg.(core.InputEventSubscriptionMessage)
	if sub.Hash != 12345 || sub.Double != 7.5 {
		t.Fatalf("input event subscription = %+v", sub)
	}

	raw, size = inputEventParamsPacket(12345, ";FLOAT64;STRING")
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	params := msg.(core.InputEventParamsMessage)
	if params.Hash != 12345 || params.Value != ";FLOAT64;STRING" || len(params.Params) != 2 || params.Params[0] != "FLOAT64" || params.Params[1] != "STRING" {
		t.Fatalf("input event params = %+v", params)
	}

	raw, size = controllersListPacket([]string{"Keyboard", "Throttle"})
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	controllers := msg.(core.ControllersListMessage)
	if len(controllers.Controllers) != 2 || controllers.Controllers[1].DeviceName != "Throttle" || controllers.Controllers[1].DeviceID != 101 {
		t.Fatalf("controllers = %+v", controllers)
	}
}

func TestDecodeCameraMessages(t *testing.T) {
	camera := bindings.SIMCONNECT_RECV_CAMERA_DATA{}
	camera.DwID = uint32(bindings.SIMCONNECT_RECV_ID_CAMERA_DATA)
	camera.CameraData.Position.X.SetFloat64(1)
	camera.CameraData.PositionReferential = bindings.SIMCONNECT_POSITION_REFERENTIAL_WORLD
	camera.CameraData.Pbh.Heading = 270
	camera.CameraData.Fov.SetFloat64(65.5)
	msg, err := decodeMessage(&camera.SIMCONNECT_RECV, uint32(unsafe.Sizeof(camera)))
	if err != nil {
		t.Fatal(err)
	}
	data := msg.(core.CameraDataMessage)
	if data.CameraData.FOV != 65.5 || data.CameraData.Position.X != 1 || data.CameraData.PositionReferential != core.PositionReferentialWorld || data.CameraData.PBH.Heading != 270 {
		t.Fatalf("camera data = %+v", data)
	}

	status := bindings.SIMCONNECT_RECV_CAMERA_STATUS{}
	status.DwID = uint32(bindings.SIMCONNECT_RECV_ID_CAMERA_STATUS)
	status.AcquiredState = uint32(bindings.SIMCONNECT_CAMERA_ACQUIRED)
	status.BGameControlled = 1
	msg, err = decodeMessage(&status.SIMCONNECT_RECV, uint32(unsafe.Sizeof(status)))
	if err != nil {
		t.Fatal(err)
	}
	cameraStatus := msg.(core.CameraStatusMessage)
	if cameraStatus.AcquiredState != core.CameraAcquired || !cameraStatus.GameControlled {
		t.Fatalf("camera status = %+v", cameraStatus)
	}

	raw, size := cameraDefinitionListPacket([]string{"Pilot", "Drone"})
	msg, err = decodeMessage(raw, size)
	if err != nil {
		t.Fatal(err)
	}
	definitions := msg.(core.CameraDefinitionListMessage)
	if len(definitions.Definitions) != 2 || definitions.Definitions[1] != "Drone" {
		t.Fatalf("camera definitions = %+v", definitions)
	}

	locker := bindings.SIMCONNECT_RECV_CAMERA_WORLD_LOCKER{}
	locker.DwID = uint32(bindings.SIMCONNECT_RECV_ID_CAMERA_WORLD_LOCKER)
	locker.Status = bindings.SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_SUCCESS
	msg, err = decodeMessage(&locker.SIMCONNECT_RECV, uint32(unsafe.Sizeof(locker)))
	if err != nil {
		t.Fatal(err)
	}
	if got := msg.(core.CameraWorldLockerMessage); got.Status != core.CameraWorldLockerStatusSuccess {
		t.Fatalf("camera world locker = %+v", got)
	}
}

func TestDispatchRoutesRequestAndHandlers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := &Sim{
		ctx:       ctx,
		errs:      make(chan error, 1),
		waiters:   map[uint32]chan core.RequestResult{},
		sendToReq: map[uint32]uint32{},
		dataSubs:  map[uint32][]func(core.Message){},
		handlers:  map[core.RecvID][]func(core.Message){},
	}

	// Request-keyed messages wake up registered waiters.
	waiter := make(chan core.RequestResult, 1)
	c.waiters[77] = waiter
	c.dispatch(core.SimObjectDataMessage{RequestID: 77})
	if got := <-waiter; got.Err != nil {
		t.Fatalf("sim object waiter err = %v", got.Err)
	}

	waiter = make(chan core.RequestResult, 1)
	c.waiters[78] = waiter
	c.dispatch(core.ClientDataMessage{SimObjectDataMessage: core.SimObjectDataMessage{RequestID: 78}})
	if got := <-waiter; got.Err != nil {
		t.Fatalf("client data waiter err = %v", got.Err)
	}

	waiter = make(chan core.RequestResult, 1)
	c.waiters[79] = waiter
	c.dispatch(core.SimObjectLiveryListMessage{FacilityListMeta: core.FacilityListMeta{RequestID: 79}})
	if got := <-waiter; got.Err != nil {
		t.Fatalf("livery waiter err = %v", got.Err)
	}

	waiter = make(chan core.RequestResult, 1)
	c.waiters[80] = waiter
	c.dispatch(core.AssignedObjectIDMessage{RequestID: 80, ObjectID: 777})
	if got := <-waiter; got.Err != nil || got.Msg.(core.AssignedObjectIDMessage).ObjectID != 777 {
		t.Fatalf("assigned object waiter = %+v", got)
	}

	waiter = make(chan core.RequestResult, 1)
	c.waiters[88] = waiter
	c.dispatch(core.AirportListMessage{FacilityListMeta: core.FacilityListMeta{RequestID: 88}})
	if got := <-waiter; got.Err != nil {
		t.Fatalf("facility waiter err = %v", got.Err)
	}

	waiter = make(chan core.RequestResult, 1)
	c.waiters[99] = waiter
	c.dispatch(core.FacilityDataMessage{UserRequestID: 99})
	if got := <-waiter; got.Err != nil {
		t.Fatalf("facility data waiter err = %v", got.Err)
	}

	waiter = make(chan core.RequestResult, 1)
	c.waiters[100] = waiter
	c.dispatch(core.InputEventValueMessage{RequestID: 100})
	if got := <-waiter; got.Err != nil {
		t.Fatalf("input event waiter err = %v", got.Err)
	}

	// Non-request messages go to registered handlers.
	eventCh := make(chan core.ClientEvent, 1)
	c.RegisterHandler(core.RecvIDEvent, func(msg core.Message) {
		if m, ok := msg.(core.ClientEvent); ok {
			select {
			case eventCh <- m:
			default:
			}
		}
	})
	c.dispatch(core.ClientEvent{EventID: 55, Data: 9})
	if got := <-eventCh; got.Data != 9 {
		t.Fatalf("event data = %d", got.Data)
	}

	flowCh := make(chan core.FlowEventMessage, 1)
	c.RegisterHandler(core.RecvIDFlowEvent, func(msg core.Message) {
		if m, ok := msg.(core.FlowEventMessage); ok {
			select {
			case flowCh <- m:
			default:
			}
		}
	})
	c.dispatch(core.FlowEventMessage{Event: core.FlowEventFlightStart})
	if got := <-flowCh; got.Event != core.FlowEventFlightStart {
		t.Fatalf("flow event = %+v", got)
	}

	commBusCh := make(chan core.CommBusMessage, 1)
	c.RegisterHandler(core.RecvIDCommBus, func(msg core.Message) {
		if m, ok := msg.(core.CommBusMessage); ok {
			select {
			case commBusCh <- m:
			default:
			}
		}
	})
	c.dispatch(core.CommBusMessage{EventID: 66, Data: "hello"})
	if got := <-commBusCh; got.Data != "hello" {
		t.Fatalf("comm bus event = %+v", got)
	}
}

// --- packet builder helpers ---

func simObjectPacket(requestID uint32, payload []byte) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_SIMOBJECT_DATA{}
	size := int(unsafe.Offsetof(header.DwData)) + len(payload)
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_SIMOBJECT_DATA)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_SIMOBJECT_DATA)
	msg.DwRequestID = requestID
	msg.DwDefineCount = 1
	copy(buf[unsafe.Offsetof(header.DwData):], payload)
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func simObjectDataByTypePacket(requestID, objectID uint32, payload []byte) (*bindings.SIMCONNECT_RECV, uint32) {
	raw, size := simObjectPacket(requestID, payload)
	raw.DwID = uint32(bindings.SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE)
	msg := (*bindings.SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE)(unsafe.Pointer(raw))
	msg.DwObjectID = objectID
	return raw, size
}

func clientDataPacket(requestID uint32, payload []byte) (*bindings.SIMCONNECT_RECV, uint32) {
	raw, size := simObjectPacket(requestID, payload)
	raw.DwID = uint32(bindings.SIMCONNECT_RECV_ID_CLIENT_DATA)
	return raw, size
}

func airportListPacket(requestID uint32, idents []string) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_AIRPORT_LIST{}
	size := int(unsafe.Offsetof(header.RgData)) + len(idents)*int(unsafe.Sizeof(bindings.SIMCONNECT_DATA_FACILITY_AIRPORT{}))
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_AIRPORT_LIST)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_AIRPORT_LIST)
	msg.DwRequestID = requestID
	msg.DwArraySize = uint32(len(idents))
	msg.DwOutOf = 1

	airports := unsafe.Slice((*bindings.SIMCONNECT_DATA_FACILITY_AIRPORT)(unsafe.Pointer(&msg.RgData[0])), len(idents))
	for i, ident := range idents {
		copy(airports[i].Ident[:], ident)
	}
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func facilityDataPacket(requestID uint32, dataType bindings.SIMCONNECT_FACILITY_DATA_TYPE, payload []byte) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_FACILITY_DATA{}
	size := int(unsafe.Offsetof(header.Data)) + len(payload)
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_FACILITY_DATA)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_FACILITY_DATA)
	msg.UserRequestId = requestID
	msg.Type = uint32(dataType)
	copy(buf[unsafe.Offsetof(header.Data):], payload)
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func inputEventListPacket(requestID uint32, names []string) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS{}
	size := int(unsafe.Offsetof(header.RgData)) + len(names)*int(unsafe.Sizeof(bindings.SIMCONNECT_INPUT_EVENT_DESCRIPTOR{}))
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_ENUMERATE_INPUT_EVENTS)
	msg.DwRequestID = requestID
	msg.DwArraySize = uint32(len(names))
	msg.DwOutOf = 1

	events := unsafe.Slice((*bindings.SIMCONNECT_INPUT_EVENT_DESCRIPTOR)(unsafe.Pointer(&msg.RgData[0])), len(names))
	for i, name := range names {
		copy(events[i].Name[:], name)
		events[i].Hash.SetUint64(12345 + uint64(i))
		events[i].EType = bindings.SIMCONNECT_INPUT_EVENT_TYPE_DOUBLE
	}
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func inputEventDoublePacket(requestID uint32, value float64) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_GET_INPUT_EVENT{}
	size := int(unsafe.Offsetof(header.Value)) + 8
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_GET_INPUT_EVENT)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_GET_INPUT_EVENT)
	msg.DwRequestID = requestID
	msg.EType = bindings.SIMCONNECT_INPUT_EVENT_TYPE_DOUBLE
	binary.LittleEndian.PutUint64(buf[unsafe.Offsetof(header.Value):], math.Float64bits(value))
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func inputEventStringPacket(requestID uint32, value string) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_GET_INPUT_EVENT{}
	payload := append([]byte(value), 0)
	size := int(unsafe.Offsetof(header.Value)) + len(payload)
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_GET_INPUT_EVENT)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_GET_INPUT_EVENT)
	msg.DwRequestID = requestID
	msg.EType = bindings.SIMCONNECT_INPUT_EVENT_TYPE_STRING
	copy(buf[unsafe.Offsetof(header.Value):], payload)
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func inputEventSubscriptionPacket(hash uint64, value float64) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT{}
	size := int(unsafe.Offsetof(header.Value)) + 8
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_SUBSCRIBE_INPUT_EVENT)
	msg.Hash.SetUint64(hash)
	msg.EType = bindings.SIMCONNECT_INPUT_EVENT_TYPE_DOUBLE
	binary.LittleEndian.PutUint64(buf[unsafe.Offsetof(header.Value):], math.Float64bits(value))
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func inputEventParamsPacket(hash uint64, value string) (*bindings.SIMCONNECT_RECV, uint32) {
	msg := bindings.SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS{}
	msg.DwSize = uint32(unsafe.Sizeof(msg))
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_ENUMERATE_INPUT_EVENT_PARAMS)
	msg.Hash.SetUint64(hash)
	copy(msg.Value[:], value)
	return &msg.SIMCONNECT_RECV, uint32(unsafe.Sizeof(msg))
}

func controllersListPacket(names []string) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_CONTROLLERS_LIST{}
	size := int(unsafe.Offsetof(header.RgData)) + len(names)*int(unsafe.Sizeof(bindings.SIMCONNECT_CONTROLLER_ITEM{}))
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_CONTROLLERS_LIST)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_CONTROLLERS_LIST)
	msg.DwArraySize = uint32(len(names))
	msg.DwOutOf = 1

	controllers := unsafe.Slice((*bindings.SIMCONNECT_CONTROLLER_ITEM)(unsafe.Pointer(&msg.RgData[0])), len(names))
	for i, name := range names {
		copy(controllers[i].DeviceName[:], name)
		controllers[i].DeviceId = uint32(100 + i)
	}
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func simObjectLiveryListPacket(requestID uint32, liveries []core.SimObjectLivery) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST{}
	size := int(unsafe.Offsetof(header.RgData)) + len(liveries)*int(unsafe.Sizeof(bindings.SIMCONNECT_ENUMERATE_SIMOBJECT_LIVERY{}))
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST)
	msg.DwRequestID = requestID
	msg.DwArraySize = uint32(len(liveries))
	msg.DwOutOf = 1

	rawLiveries := unsafe.Slice((*bindings.SIMCONNECT_ENUMERATE_SIMOBJECT_LIVERY)(unsafe.Pointer(&msg.RgData[0])), len(liveries))
	for i, livery := range liveries {
		copy(rawLiveries[i].AircraftTitle[:], livery.AircraftTitle)
		copy(rawLiveries[i].LiveryName[:], livery.LiveryName)
	}
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func commBusPacket(eventID uint32, payload []byte) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_COMM_BUS{}
	size := int(unsafe.Offsetof(header.RgData)) + len(payload)
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_COMM_BUS)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_COMM_BUS)
	msg.DwArraySize = uint32(len(payload))
	msg.DwOutOf = 1
	msg.UEventID = eventID
	copy(buf[unsafe.Offsetof(header.RgData):], payload)
	return &msg.SIMCONNECT_RECV, uint32(size)
}

func cameraDefinitionListPacket(definitions []string) (*bindings.SIMCONNECT_RECV, uint32) {
	header := bindings.SIMCONNECT_RECV_CAMERA_DEFINITION_LIST{}
	size := int(unsafe.Offsetof(header.RgData)) + len(definitions)*int(unsafe.Sizeof(bindings.SIMCONNECT_CAMERA_DEFINITION_ITEM{}))
	buf := make([]byte, size)
	msg := (*bindings.SIMCONNECT_RECV_CAMERA_DEFINITION_LIST)(unsafe.Pointer(&buf[0]))
	msg.DwSize = uint32(size)
	msg.DwID = uint32(bindings.SIMCONNECT_RECV_ID_CAMERA_DEFINITION_LIST)
	msg.DwArraySize = uint32(len(definitions))
	msg.DwOutOf = 1

	rawDefinitions := unsafe.Slice((*bindings.SIMCONNECT_CAMERA_DEFINITION_ITEM)(unsafe.Pointer(&msg.RgData[0])), len(definitions))
	for i, definition := range definitions {
		copy(rawDefinitions[i].Str[:], definition)
	}
	return &msg.SIMCONNECT_RECV, uint32(size)
}
