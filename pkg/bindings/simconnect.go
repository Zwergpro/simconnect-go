//go:build windows

// Package bindings provides Go bindings for the Microsoft Flight Simulator
// SimConnect SDK (SimConnect.dll).  All functions map 1-to-1 to the C API;
// see the MSFS SimConnect SDK documentation for parameter semantics.
package bindings

import (
	"bytes"
	"fmt"
	"syscall"
	"unsafe"
)

// DispatchProc is the callback signature passed to CallDispatch.
// pData points to the received packet; cbData is its size in bytes;
// pContext is the user-supplied context value (opaque uintptr).
type DispatchProc func(pData *SIMCONNECT_RECV, cbData uint32, pContext uintptr)

// SimConnect wraps a SimConnect session handle.
type SimConnect struct {
	handle HANDLE
}

// HResultError wraps an HRESULT value as a Go error.
type HResultError uint32

func (e HResultError) Error() string {
	return fmt.Sprintf("SimConnect HRESULT 0x%08X", uint32(e))
}

const hresultEFail = uintptr(0x80004005)

func hresultErr(r uintptr) error {
	if r == 0 {
		return nil
	}
	return HResultError(r)
}

// ---- DLL and procedure handles ----

var (
	dll = &syscall.LazyDLL{} // Name set by LoadDLL() before first proc.Call.

	procOpen                                       = dll.NewProc("SimConnect_Open")
	procClose                                      = dll.NewProc("SimConnect_Close")
	procCallDispatch                               = dll.NewProc("SimConnect_CallDispatch")
	procGetNextDispatch                            = dll.NewProc("SimConnect_GetNextDispatch")
	procGetLastSentPacketID                        = dll.NewProc("SimConnect_GetLastSentPacketID")
	procMapClientEventToSimEvent                   = dll.NewProc("SimConnect_MapClientEventToSimEvent")
	procTransmitClientEvent                        = dll.NewProc("SimConnect_TransmitClientEvent")
	procTransmitClientEvent_EX1                    = dll.NewProc("SimConnect_TransmitClientEvent_EX1")
	procSetSystemEventState                        = dll.NewProc("SimConnect_SetSystemEventState")
	procAddClientEventToNotificationGroup          = dll.NewProc("SimConnect_AddClientEventToNotificationGroup")
	procRemoveClientEvent                          = dll.NewProc("SimConnect_RemoveClientEvent")
	procSetNotificationGroupPriority               = dll.NewProc("SimConnect_SetNotificationGroupPriority")
	procClearNotificationGroup                     = dll.NewProc("SimConnect_ClearNotificationGroup")
	procRequestNotificationGroup                   = dll.NewProc("SimConnect_RequestNotificationGroup")
	procAddToDataDefinition                        = dll.NewProc("SimConnect_AddToDataDefinition")
	procClearDataDefinition                        = dll.NewProc("SimConnect_ClearDataDefinition")
	procRequestDataOnSimObject                     = dll.NewProc("SimConnect_RequestDataOnSimObject")
	procRequestDataOnSimObjectType                 = dll.NewProc("SimConnect_RequestDataOnSimObjectType")
	procSetDataOnSimObject                         = dll.NewProc("SimConnect_SetDataOnSimObject")
	procMapInputEventToClientEvent                 = dll.NewProc("SimConnect_MapInputEventToClientEvent")
	procMapInputEventToClientEvent_EX1             = dll.NewProc("SimConnect_MapInputEventToClientEvent_EX1")
	procSetInputGroupPriority                      = dll.NewProc("SimConnect_SetInputGroupPriority")
	procRemoveInputEvent                           = dll.NewProc("SimConnect_RemoveInputEvent")
	procClearInputGroup                            = dll.NewProc("SimConnect_ClearInputGroup")
	procSetInputGroupState                         = dll.NewProc("SimConnect_SetInputGroupState")
	procRequestReservedKey                         = dll.NewProc("SimConnect_RequestReservedKey")
	procSubscribeToSystemEvent                     = dll.NewProc("SimConnect_SubscribeToSystemEvent")
	procUnsubscribeFromSystemEvent                 = dll.NewProc("SimConnect_UnsubscribeFromSystemEvent")
	procWeatherRequestInterpolatedObservation      = dll.NewProc("SimConnect_WeatherRequestInterpolatedObservation")
	procWeatherRequestObservationAtStation         = dll.NewProc("SimConnect_WeatherRequestObservationAtStation")
	procWeatherRequestObservationAtNearestStation  = dll.NewProc("SimConnect_WeatherRequestObservationAtNearestStation")
	procWeatherCreateStation                       = dll.NewProc("SimConnect_WeatherCreateStation")
	procWeatherRemoveStation                       = dll.NewProc("SimConnect_WeatherRemoveStation")
	procWeatherSetObservation                      = dll.NewProc("SimConnect_WeatherSetObservation")
	procWeatherSetModeServer                       = dll.NewProc("SimConnect_WeatherSetModeServer")
	procWeatherSetModeTheme                        = dll.NewProc("SimConnect_WeatherSetModeTheme")
	procWeatherSetModeGlobal                       = dll.NewProc("SimConnect_WeatherSetModeGlobal")
	procWeatherSetModeCustom                       = dll.NewProc("SimConnect_WeatherSetModeCustom")
	procWeatherSetDynamicUpdateRate                = dll.NewProc("SimConnect_WeatherSetDynamicUpdateRate")
	procWeatherRequestCloudState                   = dll.NewProc("SimConnect_WeatherRequestCloudState")
	procWeatherCreateThermal                       = dll.NewProc("SimConnect_WeatherCreateThermal")
	procWeatherRemoveThermal                       = dll.NewProc("SimConnect_WeatherRemoveThermal")
	procAICreateParkedATCAircraft                  = dll.NewProc("SimConnect_AICreateParkedATCAircraft")
	procAICreateParkedATCAircraft_EX1              = dll.NewProc("SimConnect_AICreateParkedATCAircraft_EX1")
	procAICreateEnrouteATCAircraft                 = dll.NewProc("SimConnect_AICreateEnrouteATCAircraft")
	procAICreateEnrouteATCAircraft_EX1             = dll.NewProc("SimConnect_AICreateEnrouteATCAircraft_EX1")
	procAICreateNonATCAircraft                     = dll.NewProc("SimConnect_AICreateNonATCAircraft")
	procAICreateNonATCAircraft_EX1                 = dll.NewProc("SimConnect_AICreateNonATCAircraft_EX1")
	procAICreateSimulatedObject                    = dll.NewProc("SimConnect_AICreateSimulatedObject")
	procAICreateSimulatedObject_EX1                = dll.NewProc("SimConnect_AICreateSimulatedObject_EX1")
	procAIReleaseControl                           = dll.NewProc("SimConnect_AIReleaseControl")
	procAIRemoveObject                             = dll.NewProc("SimConnect_AIRemoveObject")
	procAISetAircraftFlightPlan                    = dll.NewProc("SimConnect_AISetAircraftFlightPlan")
	procExecuteMissionAction                       = dll.NewProc("SimConnect_ExecuteMissionAction")
	procCompleteCustomMissionAction                = dll.NewProc("SimConnect_CompleteCustomMissionAction")
	procRetrieveString                             = dll.NewProc("SimConnect_RetrieveString")
	procRequestResponseTimes                       = dll.NewProc("SimConnect_RequestResponseTimes")
	procInsertString                               = dll.NewProc("SimConnect_InsertString")
	procCameraSetRelative6DOF                      = dll.NewProc("SimConnect_CameraSetRelative6DOF")
	procMenuAddItem                                = dll.NewProc("SimConnect_MenuAddItem")
	procMenuDeleteItem                             = dll.NewProc("SimConnect_MenuDeleteItem")
	procMenuAddSubItem                             = dll.NewProc("SimConnect_MenuAddSubItem")
	procMenuDeleteSubItem                          = dll.NewProc("SimConnect_MenuDeleteSubItem")
	procRequestSystemState                         = dll.NewProc("SimConnect_RequestSystemState")
	procSetSystemState                             = dll.NewProc("SimConnect_SetSystemState")
	procMapClientDataNameToID                      = dll.NewProc("SimConnect_MapClientDataNameToID")
	procCreateClientData                           = dll.NewProc("SimConnect_CreateClientData")
	procAddToClientDataDefinition                  = dll.NewProc("SimConnect_AddToClientDataDefinition")
	procClearClientDataDefinition                  = dll.NewProc("SimConnect_ClearClientDataDefinition")
	procRequestClientData                          = dll.NewProc("SimConnect_RequestClientData")
	procSetClientData                              = dll.NewProc("SimConnect_SetClientData")
	procFlightLoad                                 = dll.NewProc("SimConnect_FlightLoad")
	procFlightSave                                 = dll.NewProc("SimConnect_FlightSave")
	procFlightPlanLoad                             = dll.NewProc("SimConnect_FlightPlanLoad")
	procText                                       = dll.NewProc("SimConnect_Text")
	procSubscribeToFacilities                      = dll.NewProc("SimConnect_SubscribeToFacilities")
	procUnsubscribeToFacilities                    = dll.NewProc("SimConnect_UnsubscribeToFacilities")
	procRequestFacilitiesList                      = dll.NewProc("SimConnect_RequestFacilitiesList")
	procSubscribeToFacilities_EX1                  = dll.NewProc("SimConnect_SubscribeToFacilities_EX1")
	procUnsubscribeToFacilities_EX1                = dll.NewProc("SimConnect_UnsubscribeToFacilities_EX1")
	procRequestFacilitiesList_EX1                  = dll.NewProc("SimConnect_RequestFacilitiesList_EX1")
	procAddToFacilityDefinition                    = dll.NewProc("SimConnect_AddToFacilityDefinition")
	procRequestFacilityData                        = dll.NewProc("SimConnect_RequestFacilityData")
	procRequestFacilityData_EX1                    = dll.NewProc("SimConnect_RequestFacilityData_EX1")
	procRequestAllFacilities                       = dll.NewProc("SimConnect_RequestAllFacilities")
	procAddFacilityDataDefinitionFilter            = dll.NewProc("SimConnect_AddFacilityDataDefinitionFilter")
	procClearAllFacilityDataDefinitionFilters      = dll.NewProc("SimConnect_ClearAllFacilityDataDefinitionFilters")
	procRequestJetwayData                          = dll.NewProc("SimConnect_RequestJetwayData")
	procEnumerateControllers                       = dll.NewProc("SimConnect_EnumerateControllers")
	procExecuteAction                              = dll.NewProc("SimConnect_ExecuteAction")
	procEnumerateInputEvents                       = dll.NewProc("SimConnect_EnumerateInputEvents")
	procGetInputEvent                              = dll.NewProc("SimConnect_GetInputEvent")
	procSetInputEvent                              = dll.NewProc("SimConnect_SetInputEvent")
	procSubscribeInputEvent                        = dll.NewProc("SimConnect_SubscribeInputEvent")
	procUnsubscribeInputEvent                      = dll.NewProc("SimConnect_UnsubscribeInputEvent")
	procEnumerateInputEventParams                  = dll.NewProc("SimConnect_EnumerateInputEventParams")
	procEnumerateSimObjectsAndLiveries             = dll.NewProc("SimConnect_EnumerateSimObjectsAndLiveries")
	procSubscribeToFlowEvent                       = dll.NewProc("SimConnect_SubscribeToFlowEvent")
	procUnsubscribeToFlowEvent                     = dll.NewProc("SimConnect_UnsubscribeToFlowEvent")
	procCameraAcquire                              = dll.NewProc("SimConnect_CameraAcquire")
	procCameraRelease                              = dll.NewProc("SimConnect_CameraRelease")
	procCameraGetStatus                            = dll.NewProc("SimConnect_CameraGetStatus")
	procCameraSet                                  = dll.NewProc("SimConnect_CameraSet")
	procCameraGet                                  = dll.NewProc("SimConnect_CameraGet")
	procCameraEnableFlag                           = dll.NewProc("SimConnect_CameraEnableFlag")
	procCameraDisableFlag                          = dll.NewProc("SimConnect_CameraDisableFlag")
	procSubscribeToCameraStatusUpdate              = dll.NewProc("SimConnect_SubscribeToCameraStatusUpdate")
	procUnsubscribeToCameraStatusUpdate            = dll.NewProc("SimConnect_UnsubscribeToCameraStatusUpdate")
	procEnumerateCameraDefinitions                 = dll.NewProc("SimConnect_EnumerateCameraDefinitions")
	procCameraSetUsingCameraDefinition             = dll.NewProc("SimConnect_CameraSetUsingCameraDefinition")
	procSubscribeToCommBusEvent                    = dll.NewProc("SimConnect_SubscribeToCommBusEvent")
	procUnsubscribeToCommBusEvent                  = dll.NewProc("SimConnect_UnsubscribeToCommBusEvent")
	procCallCommBusEvent                           = dll.NewProc("SimConnect_CallCommBusEvent")
	procSubscribeToCameraWorldLockerStatusUpdate   = dll.NewProc("SimConnect_SubscribeToCameraWorldLockerStatusUpdate")
	procUnsubscribeToCameraWorldLockerStatusUpdate = dll.NewProc("SimConnect_UnsubscribeToCameraWorldLockerStatusUpdate")
	procRequestCameraWorldLocker                   = dll.NewProc("SimConnect_RequestCameraWorldLocker")
	procDeleteCameraWorldLocker                    = dll.NewProc("SimConnect_DeleteCameraWorldLocker")
)

// ---- Helper ----

// cstrPins keeps the last N *byte allocations reachable so the GC cannot
// collect them between cstr() returning and the proc.Call that uses the
// uintptr value.  SimConnect is single-threaded; 32 slots covers every
// call that has multiple string arguments.
var (
	cstrPins [32]*byte
	cstrPin  int
)

func cstr(s string) uintptr {
	p, _ := syscall.BytePtrFromString(s)
	cstrPins[cstrPin%len(cstrPins)] = p
	cstrPin++
	return uintptr(unsafe.Pointer(p))
}

func cstrPtr(s *string) uintptr {
	if s == nil {
		return 0
	}
	return cstr(*s)
}

func f32(f float32) uintptr {
	return uintptr(*(*uint32)(unsafe.Pointer(&f)))
}

func f64(f float64) uintptr {
	return uintptr(*(*uint64)(unsafe.Pointer(&f)))
}

// ---- Session management ----

// Open connects to MSFS SimConnect.
// hWnd and userEventWin32 may be 0; configIndex should be 0 for local.
func Open(name string, hWnd HWND, userEventWin32 uint32, hEventHandle HANDLE, configIndex uint32) (*SimConnect, error) {
	var h HANDLE
	r, _, _ := procOpen.Call(
		uintptr(unsafe.Pointer(&h)),
		cstr(name),
		hWnd,
		uintptr(userEventWin32),
		hEventHandle,
		uintptr(configIndex),
	)
	if err := hresultErr(r); err != nil {
		return nil, err
	}
	return &SimConnect{handle: h}, nil
}

// Close terminates the SimConnect session.
func (sc *SimConnect) Close() error {
	r, _, _ := procClose.Call(sc.handle)
	return hresultErr(r)
}

// GetNextDispatch retrieves the next packet without blocking.
// Returns nil, 0, nil when no packet is available.
func (sc *SimConnect) GetNextDispatch() (*SIMCONNECT_RECV, uint32, error) {
	var pData *SIMCONNECT_RECV
	var cbData uint32
	r, _, _ := procGetNextDispatch.Call(
		sc.handle,
		uintptr(unsafe.Pointer(&pData)),
		uintptr(unsafe.Pointer(&cbData)),
	)
	// SimConnect_GetNextDispatch returns E_FAIL when the receive queue is empty.
	if r == hresultEFail {
		return nil, 0, nil
	}
	if err := hresultErr(r); err != nil {
		return nil, 0, err
	}
	return pData, cbData, nil
}

// CallDispatch calls pfn for every queued packet.
// The callback runs synchronously inside this call.
// Pass any context value as pContext; it will be forwarded to pfn unchanged.
func (sc *SimConnect) CallDispatch(pfn DispatchProc, pContext uintptr) error {
	cb := syscall.NewCallback(func(pData *SIMCONNECT_RECV, cbData uint32, ctx uintptr) uintptr {
		pfn(pData, cbData, ctx)
		return 0
	})
	r, _, _ := procCallDispatch.Call(sc.handle, cb, pContext)
	return hresultErr(r)
}

// GetLastSentPacketID retrieves the packet ID of the last sent packet.
func (sc *SimConnect) GetLastSentPacketID() (uint32, error) {
	var id uint32
	r, _, _ := procGetLastSentPacketID.Call(sc.handle, uintptr(unsafe.Pointer(&id)))
	return id, hresultErr(r)
}

// ---- Events ----

// MapClientEventToSimEvent associates a client event ID with a named simulator event.
// Pass an empty string to create a private client event.
func (sc *SimConnect) MapClientEventToSimEvent(eventID SIMCONNECT_CLIENT_EVENT_ID, eventName string) error {
	r, _, _ := procMapClientEventToSimEvent.Call(sc.handle, uintptr(eventID), cstr(eventName))
	return hresultErr(r)
}

// TransmitClientEvent sends an event to the simulator.
func (sc *SimConnect) TransmitClientEvent(objectID SIMCONNECT_OBJECT_ID, eventID SIMCONNECT_CLIENT_EVENT_ID, dwData uint32, groupID SIMCONNECT_NOTIFICATION_GROUP_ID, flags SIMCONNECT_EVENT_FLAG) error {
	r, _, _ := procTransmitClientEvent.Call(
		sc.handle,
		uintptr(objectID),
		uintptr(eventID),
		uintptr(dwData),
		uintptr(groupID),
		uintptr(flags),
	)
	return hresultErr(r)
}

// TransmitClientEvent_EX1 sends an event with up to 5 DWORD data values.
func (sc *SimConnect) TransmitClientEvent_EX1(objectID SIMCONNECT_OBJECT_ID, eventID SIMCONNECT_CLIENT_EVENT_ID, groupID SIMCONNECT_NOTIFICATION_GROUP_ID, flags SIMCONNECT_EVENT_FLAG, dwData0, dwData1, dwData2, dwData3, dwData4 uint32) error {
	r, _, _ := procTransmitClientEvent_EX1.Call(
		sc.handle,
		uintptr(objectID),
		uintptr(eventID),
		uintptr(groupID),
		uintptr(flags),
		uintptr(dwData0),
		uintptr(dwData1),
		uintptr(dwData2),
		uintptr(dwData3),
		uintptr(dwData4),
	)
	return hresultErr(r)
}

// SetSystemEventState enables or disables a previously subscribed system event.
func (sc *SimConnect) SetSystemEventState(eventID SIMCONNECT_CLIENT_EVENT_ID, state SIMCONNECT_STATE) error {
	r, _, _ := procSetSystemEventState.Call(sc.handle, uintptr(eventID), uintptr(state))
	return hresultErr(r)
}

// SubscribeToSystemEvent subscribes to a named simulator system event.
func (sc *SimConnect) SubscribeToSystemEvent(eventID SIMCONNECT_CLIENT_EVENT_ID, systemEventName string) error {
	r, _, _ := procSubscribeToSystemEvent.Call(sc.handle, uintptr(eventID), cstr(systemEventName))
	return hresultErr(r)
}

// UnsubscribeFromSystemEvent cancels a system event subscription.
func (sc *SimConnect) UnsubscribeFromSystemEvent(eventID SIMCONNECT_CLIENT_EVENT_ID) error {
	r, _, _ := procUnsubscribeFromSystemEvent.Call(sc.handle, uintptr(eventID))
	return hresultErr(r)
}

// ---- Notification groups ----

// AddClientEventToNotificationGroup adds a client event to a notification group.
func (sc *SimConnect) AddClientEventToNotificationGroup(groupID SIMCONNECT_NOTIFICATION_GROUP_ID, eventID SIMCONNECT_CLIENT_EVENT_ID, maskable bool) error {
	var m uintptr
	if maskable {
		m = 1
	}
	r, _, _ := procAddClientEventToNotificationGroup.Call(sc.handle, uintptr(groupID), uintptr(eventID), m)
	return hresultErr(r)
}

// RemoveClientEvent removes a client event from a notification group.
func (sc *SimConnect) RemoveClientEvent(groupID SIMCONNECT_NOTIFICATION_GROUP_ID, eventID SIMCONNECT_CLIENT_EVENT_ID) error {
	r, _, _ := procRemoveClientEvent.Call(sc.handle, uintptr(groupID), uintptr(eventID))
	return hresultErr(r)
}

// SetNotificationGroupPriority sets the priority of a notification group.
func (sc *SimConnect) SetNotificationGroupPriority(groupID SIMCONNECT_NOTIFICATION_GROUP_ID, priority uint32) error {
	r, _, _ := procSetNotificationGroupPriority.Call(sc.handle, uintptr(groupID), uintptr(priority))
	return hresultErr(r)
}

// ClearNotificationGroup removes all events from a notification group.
func (sc *SimConnect) ClearNotificationGroup(groupID SIMCONNECT_NOTIFICATION_GROUP_ID) error {
	r, _, _ := procClearNotificationGroup.Call(sc.handle, uintptr(groupID))
	return hresultErr(r)
}

// RequestNotificationGroup requests notifications to be re-sent.
func (sc *SimConnect) RequestNotificationGroup(groupID SIMCONNECT_NOTIFICATION_GROUP_ID, dwReserved, flags uint32) error {
	r, _, _ := procRequestNotificationGroup.Call(sc.handle, uintptr(groupID), uintptr(dwReserved), uintptr(flags))
	return hresultErr(r)
}

// ---- Data definitions ----

// AddToDataDefinition adds a simulation variable to a data definition.
// Pass SIMCONNECT_UNUSED for datumID to auto-assign.
func (sc *SimConnect) AddToDataDefinition(defineID SIMCONNECT_DATA_DEFINITION_ID, datumName, unitsName string, datumType SIMCONNECT_DATATYPE, epsilon float32, datumID uint32) error {
	r, _, _ := procAddToDataDefinition.Call(
		sc.handle,
		uintptr(defineID),
		cstr(datumName),
		cstr(unitsName),
		uintptr(datumType),
		f32(epsilon),
		uintptr(datumID),
	)
	return hresultErr(r)
}

// ClearDataDefinition removes all variables from a data definition.
func (sc *SimConnect) ClearDataDefinition(defineID SIMCONNECT_DATA_DEFINITION_ID) error {
	r, _, _ := procClearDataDefinition.Call(sc.handle, uintptr(defineID))
	return hresultErr(r)
}

// RequestDataOnSimObject requests data for a specific sim object.
func (sc *SimConnect) RequestDataOnSimObject(requestID SIMCONNECT_DATA_REQUEST_ID, defineID SIMCONNECT_DATA_DEFINITION_ID, objectID SIMCONNECT_OBJECT_ID, period SIMCONNECT_PERIOD, flags SIMCONNECT_DATA_REQUEST_FLAG, origin, interval, limit uint32) error {
	r, _, _ := procRequestDataOnSimObject.Call(
		sc.handle,
		uintptr(requestID),
		uintptr(defineID),
		uintptr(objectID),
		uintptr(period),
		uintptr(flags),
		uintptr(origin),
		uintptr(interval),
		uintptr(limit),
	)
	return hresultErr(r)
}

// RequestDataOnSimObjectType requests data for objects of a given type within a radius.
func (sc *SimConnect) RequestDataOnSimObjectType(requestID SIMCONNECT_DATA_REQUEST_ID, defineID SIMCONNECT_DATA_DEFINITION_ID, radiusMeters uint32, objType SIMCONNECT_SIMOBJECT_TYPE) error {
	r, _, _ := procRequestDataOnSimObjectType.Call(
		sc.handle,
		uintptr(requestID),
		uintptr(defineID),
		uintptr(radiusMeters),
		uintptr(objType),
	)
	return hresultErr(r)
}

// SetDataOnSimObject writes data to a sim object.
// pDataSet must point to a buffer of cbUnitSize * arrayCount bytes.
func (sc *SimConnect) SetDataOnSimObject(defineID SIMCONNECT_DATA_DEFINITION_ID, objectID SIMCONNECT_OBJECT_ID, flags SIMCONNECT_DATA_SET_FLAG, arrayCount, cbUnitSize uint32, pDataSet unsafe.Pointer) error {
	r, _, _ := procSetDataOnSimObject.Call(
		sc.handle,
		uintptr(defineID),
		uintptr(objectID),
		uintptr(flags),
		uintptr(arrayCount),
		uintptr(cbUnitSize),
		uintptr(pDataSet),
	)
	return hresultErr(r)
}

// ---- Input events ----

// MapInputEventToClientEvent maps an input definition to client events.
func (sc *SimConnect) MapInputEventToClientEvent(groupID SIMCONNECT_INPUT_GROUP_ID, szInputDefinition string, downEventID SIMCONNECT_CLIENT_EVENT_ID, downValue uint32, upEventID SIMCONNECT_CLIENT_EVENT_ID, upValue uint32, maskable bool) error {
	var m uintptr
	if maskable {
		m = 1
	}
	r, _, _ := procMapInputEventToClientEvent.Call(
		sc.handle,
		uintptr(groupID),
		cstr(szInputDefinition),
		uintptr(downEventID),
		uintptr(downValue),
		uintptr(upEventID),
		uintptr(upValue),
		m,
	)
	return hresultErr(r)
}

// MapInputEventToClientEvent_EX1 is the extended version of MapInputEventToClientEvent.
func (sc *SimConnect) MapInputEventToClientEvent_EX1(groupID SIMCONNECT_INPUT_GROUP_ID, szInputDefinition string, downEventID SIMCONNECT_CLIENT_EVENT_ID, downValue uint32, upEventID SIMCONNECT_CLIENT_EVENT_ID, upValue uint32, maskable bool) error {
	var m uintptr
	if maskable {
		m = 1
	}
	r, _, _ := procMapInputEventToClientEvent_EX1.Call(
		sc.handle,
		uintptr(groupID),
		cstr(szInputDefinition),
		uintptr(downEventID),
		uintptr(downValue),
		uintptr(upEventID),
		uintptr(upValue),
		m,
	)
	return hresultErr(r)
}

// SetInputGroupPriority sets the priority of an input group.
func (sc *SimConnect) SetInputGroupPriority(groupID SIMCONNECT_INPUT_GROUP_ID, priority uint32) error {
	r, _, _ := procSetInputGroupPriority.Call(sc.handle, uintptr(groupID), uintptr(priority))
	return hresultErr(r)
}

// RemoveInputEvent removes a single input definition from an input group.
func (sc *SimConnect) RemoveInputEvent(groupID SIMCONNECT_INPUT_GROUP_ID, szInputDefinition string) error {
	r, _, _ := procRemoveInputEvent.Call(sc.handle, uintptr(groupID), cstr(szInputDefinition))
	return hresultErr(r)
}

// ClearInputGroup removes all input definitions from an input group.
func (sc *SimConnect) ClearInputGroup(groupID SIMCONNECT_INPUT_GROUP_ID) error {
	r, _, _ := procClearInputGroup.Call(sc.handle, uintptr(groupID))
	return hresultErr(r)
}

// SetInputGroupState enables or disables an input group.
func (sc *SimConnect) SetInputGroupState(groupID SIMCONNECT_INPUT_GROUP_ID, state uint32) error {
	r, _, _ := procSetInputGroupState.Call(sc.handle, uintptr(groupID), uintptr(state))
	return hresultErr(r)
}

// RequestReservedKey requests exclusive use of up to three key combinations.
func (sc *SimConnect) RequestReservedKey(eventID SIMCONNECT_CLIENT_EVENT_ID, keyChoice1, keyChoice2, keyChoice3 string) error {
	r, _, _ := procRequestReservedKey.Call(
		sc.handle,
		uintptr(eventID),
		cstr(keyChoice1),
		cstr(keyChoice2),
		cstr(keyChoice3),
	)
	return hresultErr(r)
}

// ---- Weather ----

// WeatherRequestInterpolatedObservation requests a METAR at the given lat/lon/alt.
func (sc *SimConnect) WeatherRequestInterpolatedObservation(requestID SIMCONNECT_DATA_REQUEST_ID, lat, lon, alt float32) error {
	r, _, _ := procWeatherRequestInterpolatedObservation.Call(
		sc.handle, uintptr(requestID), f32(lat), f32(lon), f32(alt),
	)
	return hresultErr(r)
}

// WeatherRequestObservationAtStation requests a METAR at a named ICAO station.
func (sc *SimConnect) WeatherRequestObservationAtStation(requestID SIMCONNECT_DATA_REQUEST_ID, icao string) error {
	r, _, _ := procWeatherRequestObservationAtStation.Call(sc.handle, uintptr(requestID), cstr(icao))
	return hresultErr(r)
}

// WeatherRequestObservationAtNearestStation requests a METAR at the nearest station.
func (sc *SimConnect) WeatherRequestObservationAtNearestStation(requestID SIMCONNECT_DATA_REQUEST_ID, lat, lon float32) error {
	r, _, _ := procWeatherRequestObservationAtNearestStation.Call(sc.handle, uintptr(requestID), f32(lat), f32(lon))
	return hresultErr(r)
}

// WeatherCreateStation creates a weather reporting station.
func (sc *SimConnect) WeatherCreateStation(requestID SIMCONNECT_DATA_REQUEST_ID, icao, name string, lat, lon, alt float32) error {
	r, _, _ := procWeatherCreateStation.Call(
		sc.handle, uintptr(requestID), cstr(icao), cstr(name), f32(lat), f32(lon), f32(alt),
	)
	return hresultErr(r)
}

// WeatherRemoveStation removes a weather reporting station.
func (sc *SimConnect) WeatherRemoveStation(requestID SIMCONNECT_DATA_REQUEST_ID, icao string) error {
	r, _, _ := procWeatherRemoveStation.Call(sc.handle, uintptr(requestID), cstr(icao))
	return hresultErr(r)
}

// WeatherSetObservation sets a METAR observation string at a station.
func (sc *SimConnect) WeatherSetObservation(seconds uint32, metar string) error {
	r, _, _ := procWeatherSetObservation.Call(sc.handle, uintptr(seconds), cstr(metar))
	return hresultErr(r)
}

// WeatherSetModeServer sets weather to server mode on the given port.
func (sc *SimConnect) WeatherSetModeServer(port, seconds uint32) error {
	r, _, _ := procWeatherSetModeServer.Call(sc.handle, uintptr(port), uintptr(seconds))
	return hresultErr(r)
}

// WeatherSetModeTheme sets weather to a named theme.
func (sc *SimConnect) WeatherSetModeTheme(themeName string) error {
	r, _, _ := procWeatherSetModeTheme.Call(sc.handle, cstr(themeName))
	return hresultErr(r)
}

// WeatherSetModeGlobal sets weather to global (real-world) mode.
func (sc *SimConnect) WeatherSetModeGlobal() error {
	r, _, _ := procWeatherSetModeGlobal.Call(sc.handle)
	return hresultErr(r)
}

// WeatherSetModeCustom sets weather to custom mode.
func (sc *SimConnect) WeatherSetModeCustom() error {
	r, _, _ := procWeatherSetModeCustom.Call(sc.handle)
	return hresultErr(r)
}

// WeatherSetDynamicUpdateRate controls how often weather updates are applied.
func (sc *SimConnect) WeatherSetDynamicUpdateRate(rate uint32) error {
	r, _, _ := procWeatherSetDynamicUpdateRate.Call(sc.handle, uintptr(rate))
	return hresultErr(r)
}

// WeatherRequestCloudState requests cloud-cover state for a geographic box.
func (sc *SimConnect) WeatherRequestCloudState(requestID SIMCONNECT_DATA_REQUEST_ID, minLat, minLon, minAlt, maxLat, maxLon, maxAlt float32, flags uint32) error {
	r, _, _ := procWeatherRequestCloudState.Call(
		sc.handle, uintptr(requestID),
		f32(minLat), f32(minLon), f32(minAlt),
		f32(maxLat), f32(maxLon), f32(maxAlt),
		uintptr(flags),
	)
	return hresultErr(r)
}

// WeatherCreateThermal creates a thermal column.
func (sc *SimConnect) WeatherCreateThermal(requestID SIMCONNECT_DATA_REQUEST_ID, lat, lon, alt, radius, height, coreRate, coreTurbulence, sinkRate, sinkTurbulence, coreSize, coreTransitionSize, sinkLayerSize, sinkTransitionSize float32) error {
	r, _, _ := procWeatherCreateThermal.Call(
		sc.handle, uintptr(requestID),
		f32(lat), f32(lon), f32(alt),
		f32(radius), f32(height),
		f32(coreRate), f32(coreTurbulence),
		f32(sinkRate), f32(sinkTurbulence),
		f32(coreSize), f32(coreTransitionSize),
		f32(sinkLayerSize), f32(sinkTransitionSize),
	)
	return hresultErr(r)
}

// WeatherRemoveThermal removes a thermal previously created with WeatherCreateThermal.
func (sc *SimConnect) WeatherRemoveThermal(objectID SIMCONNECT_OBJECT_ID) error {
	r, _, _ := procWeatherRemoveThermal.Call(sc.handle, uintptr(objectID))
	return hresultErr(r)
}

// ---- AI objects ----

// AICreateParkedATCAircraft creates a parked AI ATC aircraft.
func (sc *SimConnect) AICreateParkedATCAircraft(containerTitle, tailNumber, airportID string, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procAICreateParkedATCAircraft.Call(
		sc.handle, cstr(containerTitle), cstr(tailNumber), cstr(airportID), uintptr(requestID),
	)
	return hresultErr(r)
}

// AICreateParkedATCAircraft_EX1 is the extended version with livery support.
func (sc *SimConnect) AICreateParkedATCAircraft_EX1(containerTitle, livery, tailNumber, airportID string, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procAICreateParkedATCAircraft_EX1.Call(
		sc.handle, cstr(containerTitle), cstr(livery), cstr(tailNumber), cstr(airportID), uintptr(requestID),
	)
	return hresultErr(r)
}

// AICreateEnrouteATCAircraft creates an en-route AI ATC aircraft.
func (sc *SimConnect) AICreateEnrouteATCAircraft(containerTitle, tailNumber string, flightNumber int32, flightPlanPath string, flightPlanPosition float64, touchAndGo bool, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	var tg uintptr
	if touchAndGo {
		tg = 1
	}
	r, _, _ := procAICreateEnrouteATCAircraft.Call(
		sc.handle,
		cstr(containerTitle),
		cstr(tailNumber),
		uintptr(flightNumber),
		cstr(flightPlanPath),
		f64(flightPlanPosition),
		tg,
		uintptr(requestID),
	)
	return hresultErr(r)
}

// AICreateEnrouteATCAircraft_EX1 is the extended version with livery support.
func (sc *SimConnect) AICreateEnrouteATCAircraft_EX1(containerTitle, livery, tailNumber string, flightNumber int32, flightPlanPath string, flightPlanPosition float64, touchAndGo bool, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	var tg uintptr
	if touchAndGo {
		tg = 1
	}
	r, _, _ := procAICreateEnrouteATCAircraft_EX1.Call(
		sc.handle,
		cstr(containerTitle),
		cstr(livery),
		cstr(tailNumber),
		uintptr(flightNumber),
		cstr(flightPlanPath),
		f64(flightPlanPosition),
		tg,
		uintptr(requestID),
	)
	return hresultErr(r)
}

// AICreateNonATCAircraft creates a non-ATC AI aircraft at an initial position.
func (sc *SimConnect) AICreateNonATCAircraft(containerTitle, tailNumber string, initPos SIMCONNECT_DATA_INITPOSITION, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procAICreateNonATCAircraft.Call(
		sc.handle,
		cstr(containerTitle),
		cstr(tailNumber),
		uintptr(unsafe.Pointer(&initPos)),
		uintptr(requestID),
	)
	return hresultErr(r)
}

// AICreateNonATCAircraft_EX1 is the extended version with livery support.
func (sc *SimConnect) AICreateNonATCAircraft_EX1(containerTitle, livery, tailNumber string, initPos SIMCONNECT_DATA_INITPOSITION, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procAICreateNonATCAircraft_EX1.Call(
		sc.handle,
		cstr(containerTitle),
		cstr(livery),
		cstr(tailNumber),
		uintptr(unsafe.Pointer(&initPos)),
		uintptr(requestID),
	)
	return hresultErr(r)
}

// AICreateSimulatedObject creates a non-aircraft AI object.
func (sc *SimConnect) AICreateSimulatedObject(containerTitle string, initPos SIMCONNECT_DATA_INITPOSITION, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procAICreateSimulatedObject.Call(
		sc.handle, cstr(containerTitle), uintptr(unsafe.Pointer(&initPos)), uintptr(requestID),
	)
	return hresultErr(r)
}

// AICreateSimulatedObject_EX1 is the extended version with livery support.
func (sc *SimConnect) AICreateSimulatedObject_EX1(containerTitle, livery string, initPos SIMCONNECT_DATA_INITPOSITION, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procAICreateSimulatedObject_EX1.Call(
		sc.handle, cstr(containerTitle), cstr(livery), uintptr(unsafe.Pointer(&initPos)), uintptr(requestID),
	)
	return hresultErr(r)
}

// AIReleaseControl releases AI control of an object.
func (sc *SimConnect) AIReleaseControl(objectID SIMCONNECT_OBJECT_ID, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procAIReleaseControl.Call(sc.handle, uintptr(objectID), uintptr(requestID))
	return hresultErr(r)
}

// AIRemoveObject removes an AI object from the simulation.
func (sc *SimConnect) AIRemoveObject(objectID SIMCONNECT_OBJECT_ID, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procAIRemoveObject.Call(sc.handle, uintptr(objectID), uintptr(requestID))
	return hresultErr(r)
}

// AISetAircraftFlightPlan assigns a flight plan to an AI aircraft.
func (sc *SimConnect) AISetAircraftFlightPlan(objectID SIMCONNECT_OBJECT_ID, flightPlanPath string, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procAISetAircraftFlightPlan.Call(
		sc.handle, uintptr(objectID), cstr(flightPlanPath), uintptr(requestID),
	)
	return hresultErr(r)
}

// ---- Mission actions ----

// ExecuteMissionAction executes a named mission action by GUID.
func (sc *SimConnect) ExecuteMissionAction(guidInstanceID GUID) error {
	r, _, _ := procExecuteMissionAction.Call(sc.handle, uintptr(unsafe.Pointer(&guidInstanceID)))
	return hresultErr(r)
}

// CompleteCustomMissionAction completes a custom mission action.
func (sc *SimConnect) CompleteCustomMissionAction(guidInstanceID GUID) error {
	r, _, _ := procCompleteCustomMissionAction.Call(sc.handle, uintptr(unsafe.Pointer(&guidInstanceID)))
	return hresultErr(r)
}

// ---- Utilities ----

// RetrieveString extracts a variable-length string from a received packet.
func RetrieveString(pData *SIMCONNECT_RECV, cbData uint32, pStringV unsafe.Pointer) (string, error) {
	var pszString *byte
	var cbString uint32
	r, _, _ := procRetrieveString.Call(
		uintptr(unsafe.Pointer(pData)),
		uintptr(cbData),
		uintptr(pStringV),
		uintptr(unsafe.Pointer(&pszString)),
		uintptr(unsafe.Pointer(&cbString)),
	)
	if err := hresultErr(r); err != nil {
		return "", err
	}
	if pszString == nil {
		return "", nil
	}
	b := unsafe.Slice(pszString, cbString)
	return cStringBytes(b), nil
}

// RequestResponseTimes retrieves response-time measurements.
func (sc *SimConnect) RequestResponseTimes(n uint32) ([]float32, error) {
	if n == 0 {
		return nil, nil
	}
	buf := make([]float32, n)
	r, _, _ := procRequestResponseTimes.Call(
		sc.handle,
		uintptr(n),
		uintptr(unsafe.Pointer(&buf[0])),
	)
	return buf, hresultErr(r)
}

// InsertString inserts a string into a variable-length string buffer.
func InsertString(pDest []byte, pSource string) (end unsafe.Pointer, cbStringV uint32, err error) {
	if len(pDest) == 0 {
		return nil, 0, fmt.Errorf("simconnect insert string: empty destination")
	}
	src, _ := syscall.BytePtrFromString(pSource)
	r, _, _ := procInsertString.Call(
		uintptr(unsafe.Pointer(&pDest[0])),
		uintptr(len(pDest)),
		uintptr(unsafe.Pointer(&end)),
		uintptr(unsafe.Pointer(&cbStringV)),
		uintptr(unsafe.Pointer(src)),
	)
	return end, cbStringV, hresultErr(r)
}

func cStringBytes(b []byte) string {
	if n := bytes.IndexByte(b, 0); n >= 0 {
		return string(b[:n])
	}
	return string(b)
}

// ---- Camera ----

// CameraSetRelative6DOF adjusts the camera position relative to its current position.
func (sc *SimConnect) CameraSetRelative6DOF(deltaX, deltaY, deltaZ, pitchDeg, bankDeg, headingDeg float32) error {
	r, _, _ := procCameraSetRelative6DOF.Call(
		sc.handle,
		f32(deltaX), f32(deltaY), f32(deltaZ),
		f32(pitchDeg), f32(bankDeg), f32(headingDeg),
	)
	return hresultErr(r)
}

// CameraAcquire acquires camera control with the given client ID.
func (sc *SimConnect) CameraAcquire(clientID string) error {
	r, _, _ := procCameraAcquire.Call(sc.handle, cstr(clientID))
	return hresultErr(r)
}

// CameraRelease releases a previously acquired camera.
func (sc *SimConnect) CameraRelease(cameraDefName string) error {
	r, _, _ := procCameraRelease.Call(sc.handle, cstr(cameraDefName))
	return hresultErr(r)
}

// CameraGetStatus requests the current camera acquisition status.
func (sc *SimConnect) CameraGetStatus() error {
	r, _, _ := procCameraGetStatus.Call(sc.handle)
	return hresultErr(r)
}

// CameraSet sets the camera using a SIMCONNECT_DATA_CAMERA and a data mask.
func (sc *SimConnect) CameraSet(cameraData SIMCONNECT_DATA_CAMERA, dataMask uint32) error {
	r, _, _ := procCameraSet.Call(
		sc.handle,
		uintptr(unsafe.Pointer(&cameraData)),
		uintptr(dataMask),
	)
	return hresultErr(r)
}

// CameraGet requests camera data for the given referential.
func (sc *SimConnect) CameraGet(referential uint32) error {
	r, _, _ := procCameraGet.Call(sc.handle, uintptr(referential))
	return hresultErr(r)
}

// CameraEnableFlag enables a camera flag.
func (sc *SimConnect) CameraEnableFlag(flag uint32) error {
	r, _, _ := procCameraEnableFlag.Call(sc.handle, uintptr(flag))
	return hresultErr(r)
}

// CameraDisableFlag disables a camera flag.
func (sc *SimConnect) CameraDisableFlag(flag uint32) error {
	r, _, _ := procCameraDisableFlag.Call(sc.handle, uintptr(flag))
	return hresultErr(r)
}

// SubscribeToCameraStatusUpdate subscribes to camera status change events.
func (sc *SimConnect) SubscribeToCameraStatusUpdate() error {
	r, _, _ := procSubscribeToCameraStatusUpdate.Call(sc.handle)
	return hresultErr(r)
}

// UnsubscribeToCameraStatusUpdate cancels a camera status subscription.
func (sc *SimConnect) UnsubscribeToCameraStatusUpdate() error {
	r, _, _ := procUnsubscribeToCameraStatusUpdate.Call(sc.handle)
	return hresultErr(r)
}

// EnumerateCameraDefinitions requests the list of available camera definitions.
func (sc *SimConnect) EnumerateCameraDefinitions() error {
	r, _, _ := procEnumerateCameraDefinitions.Call(sc.handle)
	return hresultErr(r)
}

// CameraSetUsingCameraDefinition sets a camera using a named camera definition.
func (sc *SimConnect) CameraSetUsingCameraDefinition(cameraDefinition string) error {
	r, _, _ := procCameraSetUsingCameraDefinition.Call(sc.handle, cstr(cameraDefinition))
	return hresultErr(r)
}

// SubscribeToCameraWorldLockerStatusUpdate subscribes to world-locker status events.
func (sc *SimConnect) SubscribeToCameraWorldLockerStatusUpdate() error {
	r, _, _ := procSubscribeToCameraWorldLockerStatusUpdate.Call(sc.handle)
	return hresultErr(r)
}

// UnsubscribeToCameraWorldLockerStatusUpdate cancels a world-locker status subscription.
func (sc *SimConnect) UnsubscribeToCameraWorldLockerStatusUpdate() error {
	r, _, _ := procUnsubscribeToCameraWorldLockerStatusUpdate.Call(sc.handle)
	return hresultErr(r)
}

// RequestCameraWorldLocker requests a world-locker at the given position.
func (sc *SimConnect) RequestCameraWorldLocker(lockerPosition SIMCONNECT_DATA_XYZ, referential SIMCONNECT_POSITION_REFERENTIAL, objectID uint32) error {
	r, _, _ := procRequestCameraWorldLocker.Call(
		sc.handle,
		uintptr(unsafe.Pointer(&lockerPosition)),
		uintptr(referential),
		uintptr(objectID),
	)
	return hresultErr(r)
}

// DeleteCameraWorldLocker removes the active world locker.
func (sc *SimConnect) DeleteCameraWorldLocker() error {
	r, _, _ := procDeleteCameraWorldLocker.Call(sc.handle)
	return hresultErr(r)
}

// ---- Menus ----

// MenuAddItem adds a top-level SimConnect menu item.
func (sc *SimConnect) MenuAddItem(menuItem string, menuEventID SIMCONNECT_CLIENT_EVENT_ID, dwData uint32) error {
	r, _, _ := procMenuAddItem.Call(sc.handle, cstr(menuItem), uintptr(menuEventID), uintptr(dwData))
	return hresultErr(r)
}

// MenuDeleteItem removes a top-level menu item.
func (sc *SimConnect) MenuDeleteItem(menuEventID SIMCONNECT_CLIENT_EVENT_ID) error {
	r, _, _ := procMenuDeleteItem.Call(sc.handle, uintptr(menuEventID))
	return hresultErr(r)
}

// MenuAddSubItem adds a sub-menu item to an existing menu item.
func (sc *SimConnect) MenuAddSubItem(menuEventID SIMCONNECT_CLIENT_EVENT_ID, menuItem string, subMenuEventID SIMCONNECT_CLIENT_EVENT_ID, dwData uint32) error {
	r, _, _ := procMenuAddSubItem.Call(
		sc.handle, uintptr(menuEventID), cstr(menuItem), uintptr(subMenuEventID), uintptr(dwData),
	)
	return hresultErr(r)
}

// MenuDeleteSubItem removes a sub-menu item.
func (sc *SimConnect) MenuDeleteSubItem(menuEventID SIMCONNECT_CLIENT_EVENT_ID, subMenuEventID SIMCONNECT_CLIENT_EVENT_ID) error {
	r, _, _ := procMenuDeleteSubItem.Call(sc.handle, uintptr(menuEventID), uintptr(subMenuEventID))
	return hresultErr(r)
}

// ---- System state ----

// RequestSystemState requests the current value of a named system state variable.
func (sc *SimConnect) RequestSystemState(requestID SIMCONNECT_DATA_REQUEST_ID, state string) error {
	r, _, _ := procRequestSystemState.Call(sc.handle, uintptr(requestID), cstr(state))
	return hresultErr(r)
}

// SetSystemState sets a named system state variable.
func (sc *SimConnect) SetSystemState(state string, dwInteger uint32, fFloat float32, szString string) error {
	r, _, _ := procSetSystemState.Call(
		sc.handle, cstr(state), uintptr(dwInteger), f32(fFloat), cstr(szString),
	)
	return hresultErr(r)
}

// ---- Client data ----

// MapClientDataNameToID maps a client data area name to a client data ID.
func (sc *SimConnect) MapClientDataNameToID(clientDataName string, clientDataID SIMCONNECT_CLIENT_DATA_ID) error {
	r, _, _ := procMapClientDataNameToID.Call(sc.handle, cstr(clientDataName), uintptr(clientDataID))
	return hresultErr(r)
}

// CreateClientData creates a client data area of the given size.
func (sc *SimConnect) CreateClientData(clientDataID SIMCONNECT_CLIENT_DATA_ID, dwSize uint32, flags SIMCONNECT_CREATE_CLIENT_DATA_FLAG) error {
	r, _, _ := procCreateClientData.Call(sc.handle, uintptr(clientDataID), uintptr(dwSize), uintptr(flags))
	return hresultErr(r)
}

// AddToClientDataDefinition adds a datum to a client data definition.
func (sc *SimConnect) AddToClientDataDefinition(defineID SIMCONNECT_CLIENT_DATA_DEFINITION_ID, dwOffset, dwSizeOrType uint32, epsilon float32, datumID uint32) error {
	r, _, _ := procAddToClientDataDefinition.Call(
		sc.handle,
		uintptr(defineID),
		uintptr(dwOffset),
		uintptr(dwSizeOrType),
		f32(epsilon),
		uintptr(datumID),
	)
	return hresultErr(r)
}

// ClearClientDataDefinition removes all datums from a client data definition.
func (sc *SimConnect) ClearClientDataDefinition(defineID SIMCONNECT_CLIENT_DATA_DEFINITION_ID) error {
	r, _, _ := procClearClientDataDefinition.Call(sc.handle, uintptr(defineID))
	return hresultErr(r)
}

// RequestClientData requests data from a client data area.
func (sc *SimConnect) RequestClientData(clientDataID SIMCONNECT_CLIENT_DATA_ID, requestID SIMCONNECT_DATA_REQUEST_ID, defineID SIMCONNECT_CLIENT_DATA_DEFINITION_ID, period SIMCONNECT_CLIENT_DATA_PERIOD, flags SIMCONNECT_CLIENT_DATA_REQUEST_FLAG, origin, interval, limit uint32) error {
	r, _, _ := procRequestClientData.Call(
		sc.handle,
		uintptr(clientDataID),
		uintptr(requestID),
		uintptr(defineID),
		uintptr(period),
		uintptr(flags),
		uintptr(origin),
		uintptr(interval),
		uintptr(limit),
	)
	return hresultErr(r)
}

// SetClientData writes data to a client data area.
func (sc *SimConnect) SetClientData(clientDataID SIMCONNECT_CLIENT_DATA_ID, defineID SIMCONNECT_CLIENT_DATA_DEFINITION_ID, flags SIMCONNECT_CLIENT_DATA_SET_FLAG, dwReserved, cbUnitSize uint32, pDataSet unsafe.Pointer) error {
	r, _, _ := procSetClientData.Call(
		sc.handle,
		uintptr(clientDataID),
		uintptr(defineID),
		uintptr(flags),
		uintptr(dwReserved),
		uintptr(cbUnitSize),
		uintptr(pDataSet),
	)
	return hresultErr(r)
}

// ---- Flight files ----

// FlightLoad loads a saved flight (.FLT file).
func (sc *SimConnect) FlightLoad(fileName string) error {
	r, _, _ := procFlightLoad.Call(sc.handle, cstr(fileName))
	return hresultErr(r)
}

// FlightSave saves the current flight to a .FLT file.
func (sc *SimConnect) FlightSave(fileName, title, description string, flags uint32) error {
	r, _, _ := procFlightSave.Call(sc.handle, cstr(fileName), cstr(title), cstr(description), uintptr(flags))
	return hresultErr(r)
}

// FlightSaveWithOptionalTitle saves the current flight to a .FLT file.
// If title is nil, SimConnect uses fileName as the saved flight title.
func (sc *SimConnect) FlightSaveWithOptionalTitle(fileName string, title *string, description string, flags uint32) error {
	r, _, _ := procFlightSave.Call(sc.handle, cstr(fileName), cstrPtr(title), cstr(description), uintptr(flags))
	return hresultErr(r)
}

// FlightPlanLoad loads a flight plan (.PLN file).
func (sc *SimConnect) FlightPlanLoad(fileName string) error {
	r, _, _ := procFlightPlanLoad.Call(sc.handle, cstr(fileName))
	return hresultErr(r)
}

// ---- Text display ----

// Text displays text in the simulator.
func (sc *SimConnect) Text(textType SIMCONNECT_TEXT_TYPE, timeSeconds float32, eventID SIMCONNECT_CLIENT_EVENT_ID, cbUnitSize uint32, pDataSet unsafe.Pointer) error {
	r, _, _ := procText.Call(
		sc.handle,
		uintptr(textType),
		f32(timeSeconds),
		uintptr(eventID),
		uintptr(cbUnitSize),
		uintptr(pDataSet),
	)
	return hresultErr(r)
}

// ---- Facilities ----

// SubscribeToFacilities subscribes to facility additions in range.
func (sc *SimConnect) SubscribeToFacilities(listType SIMCONNECT_FACILITY_LIST_TYPE, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procSubscribeToFacilities.Call(sc.handle, uintptr(listType), uintptr(requestID))
	return hresultErr(r)
}

// UnsubscribeToFacilities cancels a facilities subscription.
func (sc *SimConnect) UnsubscribeToFacilities(listType SIMCONNECT_FACILITY_LIST_TYPE) error {
	r, _, _ := procUnsubscribeToFacilities.Call(sc.handle, uintptr(listType))
	return hresultErr(r)
}

// RequestFacilitiesList requests a one-time list of all in-range facilities.
func (sc *SimConnect) RequestFacilitiesList(listType SIMCONNECT_FACILITY_LIST_TYPE, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procRequestFacilitiesList.Call(sc.handle, uintptr(listType), uintptr(requestID))
	return hresultErr(r)
}

// SubscribeToFacilities_EX1 subscribes with separate new/old request IDs.
func (sc *SimConnect) SubscribeToFacilities_EX1(listType SIMCONNECT_FACILITY_LIST_TYPE, newElemRequestID, oldElemRequestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procSubscribeToFacilities_EX1.Call(sc.handle, uintptr(listType), uintptr(newElemRequestID), uintptr(oldElemRequestID))
	return hresultErr(r)
}

// UnsubscribeToFacilities_EX1 cancels new/old range subscriptions individually.
func (sc *SimConnect) UnsubscribeToFacilities_EX1(listType SIMCONNECT_FACILITY_LIST_TYPE, unsubscribeNewInRange, unsubscribeOldOutRange bool) error {
	var ni, oo uintptr
	if unsubscribeNewInRange {
		ni = 1
	}
	if unsubscribeOldOutRange {
		oo = 1
	}
	r, _, _ := procUnsubscribeToFacilities_EX1.Call(sc.handle, uintptr(listType), ni, oo)
	return hresultErr(r)
}

// RequestFacilitiesList_EX1 requests a one-time facility list (extended).
func (sc *SimConnect) RequestFacilitiesList_EX1(listType SIMCONNECT_FACILITY_LIST_TYPE, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procRequestFacilitiesList_EX1.Call(sc.handle, uintptr(listType), uintptr(requestID))
	return hresultErr(r)
}

// AddToFacilityDefinition adds a named field to a facility data definition.
func (sc *SimConnect) AddToFacilityDefinition(defineID SIMCONNECT_DATA_DEFINITION_ID, fieldName string) error {
	r, _, _ := procAddToFacilityDefinition.Call(sc.handle, uintptr(defineID), cstr(fieldName))
	return hresultErr(r)
}

// RequestFacilityData requests facility data by ICAO and region.
func (sc *SimConnect) RequestFacilityData(defineID SIMCONNECT_DATA_DEFINITION_ID, requestID SIMCONNECT_DATA_REQUEST_ID, icao, region string) error {
	r, _, _ := procRequestFacilityData.Call(
		sc.handle, uintptr(defineID), uintptr(requestID), cstr(icao), cstr(region),
	)
	return hresultErr(r)
}

// RequestFacilityData_EX1 is the extended version with a type discriminator.
func (sc *SimConnect) RequestFacilityData_EX1(defineID SIMCONNECT_DATA_DEFINITION_ID, requestID SIMCONNECT_DATA_REQUEST_ID, icao, region string, facilityType byte) error {
	r, _, _ := procRequestFacilityData_EX1.Call(
		sc.handle, uintptr(defineID), uintptr(requestID), cstr(icao), cstr(region), uintptr(facilityType),
	)
	return hresultErr(r)
}

// RequestAllFacilities requests all facilities of a given type worldwide.
func (sc *SimConnect) RequestAllFacilities(listType SIMCONNECT_FACILITY_LIST_TYPE, requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procRequestAllFacilities.Call(sc.handle, uintptr(listType), uintptr(requestID))
	return hresultErr(r)
}

// AddFacilityDataDefinitionFilter adds a filter to a facility data definition.
func (sc *SimConnect) AddFacilityDataDefinitionFilter(defineID SIMCONNECT_DATA_DEFINITION_ID, filterPath string, cbUnitSize uint32, pFilterData unsafe.Pointer) error {
	r, _, _ := procAddFacilityDataDefinitionFilter.Call(
		sc.handle, uintptr(defineID), cstr(filterPath), uintptr(cbUnitSize), uintptr(pFilterData),
	)
	return hresultErr(r)
}

// ClearAllFacilityDataDefinitionFilters removes all filters from a facility data definition.
func (sc *SimConnect) ClearAllFacilityDataDefinitionFilters(defineID SIMCONNECT_DATA_DEFINITION_ID) error {
	r, _, _ := procClearAllFacilityDataDefinitionFilters.Call(sc.handle, uintptr(defineID))
	return hresultErr(r)
}

// RequestJetwayData requests jetway data for the given airport and parking indexes.
func (sc *SimConnect) RequestJetwayData(airportIcao string, indexes []int32) error {
	var pIdx *int32
	if len(indexes) > 0 {
		pIdx = &indexes[0]
	}
	r, _, _ := procRequestJetwayData.Call(
		sc.handle,
		cstr(airportIcao),
		uintptr(len(indexes)),
		uintptr(unsafe.Pointer(pIdx)),
	)
	return hresultErr(r)
}

// ---- Controllers ----

// EnumerateControllers requests a list of connected input controllers.
func (sc *SimConnect) EnumerateControllers() error {
	r, _, _ := procEnumerateControllers.Call(sc.handle)
	return hresultErr(r)
}

// ---- Actions & input events ----

// ExecuteAction executes a named WASM action.
func (sc *SimConnect) ExecuteAction(cbRequestID uint32, actionID string, cbUnitSize uint32, pParamValues unsafe.Pointer) error {
	r, _, _ := procExecuteAction.Call(
		sc.handle,
		uintptr(cbRequestID),
		cstr(actionID),
		uintptr(cbUnitSize),
		uintptr(pParamValues),
	)
	return hresultErr(r)
}

// EnumerateInputEvents requests a list of all available input events.
func (sc *SimConnect) EnumerateInputEvents(requestID SIMCONNECT_DATA_REQUEST_ID) error {
	r, _, _ := procEnumerateInputEvents.Call(sc.handle, uintptr(requestID))
	return hresultErr(r)
}

// GetInputEvent requests the current value of an input event by hash.
func (sc *SimConnect) GetInputEvent(requestID SIMCONNECT_DATA_REQUEST_ID, hash uint64) error {
	r, _, _ := procGetInputEvent.Call(sc.handle, uintptr(requestID), uintptr(hash))
	return hresultErr(r)
}

// SetInputEvent sets the value of an input event by hash.
func (sc *SimConnect) SetInputEvent(hash uint64, cbUnitSize uint32, value unsafe.Pointer) error {
	r, _, _ := procSetInputEvent.Call(
		sc.handle,
		uintptr(hash),
		uintptr(cbUnitSize),
		uintptr(value),
	)
	return hresultErr(r)
}

// SubscribeInputEvent subscribes to value changes for an input event.
func (sc *SimConnect) SubscribeInputEvent(hash uint64) error {
	r, _, _ := procSubscribeInputEvent.Call(sc.handle, uintptr(hash))
	return hresultErr(r)
}

// UnsubscribeInputEvent cancels a subscription to an input event.
func (sc *SimConnect) UnsubscribeInputEvent(hash uint64) error {
	r, _, _ := procUnsubscribeInputEvent.Call(sc.handle, uintptr(hash))
	return hresultErr(r)
}

// EnumerateInputEventParams requests the parameters of an input event.
func (sc *SimConnect) EnumerateInputEventParams(hash uint64) error {
	r, _, _ := procEnumerateInputEventParams.Call(sc.handle, uintptr(hash))
	return hresultErr(r)
}

// EnumerateSimObjectsAndLiveries requests a list of sim objects with their liveries.
func (sc *SimConnect) EnumerateSimObjectsAndLiveries(requestID SIMCONNECT_DATA_REQUEST_ID, objType SIMCONNECT_SIMOBJECT_TYPE) error {
	r, _, _ := procEnumerateSimObjectsAndLiveries.Call(sc.handle, uintptr(requestID), uintptr(objType))
	return hresultErr(r)
}

// ---- Flow events ----

// SubscribeToFlowEvent subscribes to flight flow events (load, teleport, etc.).
func (sc *SimConnect) SubscribeToFlowEvent() error {
	r, _, _ := procSubscribeToFlowEvent.Call(sc.handle)
	return hresultErr(r)
}

// UnsubscribeToFlowEvent cancels a flow event subscription.
func (sc *SimConnect) UnsubscribeToFlowEvent() error {
	r, _, _ := procUnsubscribeToFlowEvent.Call(sc.handle)
	return hresultErr(r)
}

// ---- CommBus ----

// SubscribeToCommBusEvent subscribes to a named comm-bus event.
func (sc *SimConnect) SubscribeToCommBusEvent(eventID SIMCONNECT_CLIENT_EVENT_ID, eventName string) error {
	r, _, _ := procSubscribeToCommBusEvent.Call(sc.handle, uintptr(eventID), cstr(eventName))
	return hresultErr(r)
}

// UnsubscribeToCommBusEvent cancels a comm-bus event subscription.
func (sc *SimConnect) UnsubscribeToCommBusEvent(eventID SIMCONNECT_CLIENT_EVENT_ID) error {
	r, _, _ := procUnsubscribeToCommBusEvent.Call(sc.handle, uintptr(eventID))
	return hresultErr(r)
}

// CallCommBusEvent broadcasts a comm-bus event to the given targets.
func (sc *SimConnect) CallCommBusEvent(eventName string, broadcastTo SIMCONNECT_COMM_BUS_BROADCAST_TO, data string) error {
	dataBytes := append([]byte(data), 0)
	return sc.CallCommBusEventBytes(eventName, broadcastTo, dataBytes)
}

// CallCommBusEventBytes broadcasts a comm-bus event with an exact byte payload.
func (sc *SimConnect) CallCommBusEventBytes(eventName string, broadcastTo SIMCONNECT_COMM_BUS_BROADCAST_TO, data []byte) error {
	var pData *byte
	if len(data) > 0 {
		pData = &data[0]
	}
	r, _, _ := procCallCommBusEvent.Call(
		sc.handle,
		cstr(eventName),
		uintptr(broadcastTo),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(pData)),
	)
	return hresultErr(r)
}
