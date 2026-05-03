//go:build windows

package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"unsafe"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// ─── rawMessage (internal only) ─────────────────────────────────────────────

type rawMessage struct {
	id   core.RecvID
	raw  *bindings.SIMCONNECT_RECV
	size uint32
}

func (m rawMessage) RecvID() core.RecvID { return m.id }

// ─── Decode functions ────────────────────────────────────────────────────────

func decodeMessage(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	if raw == nil {
		return nil, fmt.Errorf("%w: nil packet", core.ErrDecode)
	}

	switch bindings.SIMCONNECT_RECV_ID(raw.DwID) {
	case bindings.SIMCONNECT_RECV_ID_OPEN:
		msg := (*bindings.SIMCONNECT_RECV_OPEN)(unsafe.Pointer(raw))
		return core.OpenMessage{
			ApplicationName:         fixedString(msg.SzApplicationName[:]),
			ApplicationVersionMajor: msg.DwApplicationVersionMajor,
			ApplicationVersionMinor: msg.DwApplicationVersionMinor,
			ApplicationBuildMajor:   msg.DwApplicationBuildMajor,
			ApplicationBuildMinor:   msg.DwApplicationBuildMinor,
			SimConnectVersionMajor:  msg.DwSimConnectVersionMajor,
			SimConnectVersionMinor:  msg.DwSimConnectVersionMinor,
			SimConnectBuildMajor:    msg.DwSimConnectBuildMajor,
			SimConnectBuildMinor:    msg.DwSimConnectBuildMinor,
		}, nil
	case bindings.SIMCONNECT_RECV_ID_QUIT:
		return core.QuitMessage{}, nil
	case bindings.SIMCONNECT_RECV_ID_EVENT:
		msg := (*bindings.SIMCONNECT_RECV_EVENT)(unsafe.Pointer(raw))
		return core.ClientEvent{GroupID: msg.UGroupID, EventID: msg.UEventID, Data: msg.DwData, DataValues: [5]uint32{msg.DwData}}, nil
	case bindings.SIMCONNECT_RECV_ID_EVENT_EX1:
		msg := (*bindings.SIMCONNECT_RECV_EVENT_EX1)(unsafe.Pointer(raw))
		values := [5]uint32{msg.DwData0, msg.DwData1, msg.DwData2, msg.DwData3, msg.DwData4}
		return core.ClientEventEX1{ClientEvent: core.ClientEvent{GroupID: msg.UGroupID, EventID: msg.UEventID, Data: msg.DwData0, DataValues: values}}, nil
	case bindings.SIMCONNECT_RECV_ID_EVENT_FILENAME:
		msg := (*bindings.SIMCONNECT_RECV_EVENT_FILENAME)(unsafe.Pointer(raw))
		return core.FilenameEvent{
			ClientEvent: core.ClientEvent{
				GroupID:    msg.UGroupID,
				EventID:    msg.UEventID,
				Data:       msg.DwData,
				DataValues: [5]uint32{msg.DwData},
				FileName:   fixedString(msg.SzFileName[:]),
				Flags:      msg.DwFlags,
			},
			FileName: fixedString(msg.SzFileName[:]),
			Flags:    msg.DwFlags,
		}, nil
	case bindings.SIMCONNECT_RECV_ID_EVENT_OBJECT_ADDREMOVE:
		msg := (*bindings.SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE)(unsafe.Pointer(raw))
		return core.ObjectAddRemoveEvent{
			ClientEvent: core.ClientEvent{
				GroupID:    msg.UGroupID,
				EventID:    msg.UEventID,
				Data:       msg.DwData,
				DataValues: [5]uint32{msg.DwData},
				ObjectType: core.SimObjectType(msg.EObjType),
			},
			ObjectType: core.SimObjectType(msg.EObjType),
		}, nil
	case bindings.SIMCONNECT_RECV_ID_EVENT_FRAME:
		msg := (*bindings.SIMCONNECT_RECV_EVENT_FRAME)(unsafe.Pointer(raw))
		return core.FrameEvent{
			ClientEvent: core.ClientEvent{
				GroupID:    msg.UGroupID,
				EventID:    msg.UEventID,
				Data:       msg.DwData,
				DataValues: [5]uint32{msg.DwData},
				FrameRate:  msg.FFrameRate,
				SimSpeed:   msg.FSimSpeed,
			},
			FrameRate: msg.FFrameRate,
			SimSpeed:  msg.FSimSpeed,
		}, nil
	case bindings.SIMCONNECT_RECV_ID_EXCEPTION:
		msg := (*bindings.SIMCONNECT_RECV_EXCEPTION)(unsafe.Pointer(raw))
		return core.ExceptionError{Exception: core.Exception(msg.DwException), SendID: msg.DwSendID, Index: msg.DwIndex}, nil
	case bindings.SIMCONNECT_RECV_ID_SIMOBJECT_DATA:
		return decodeSimObjectData(raw, size), nil
	case bindings.SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE:
		return core.SimObjectDataByTypeMessage{SimObjectDataMessage: decodeSimObjectData(raw, size)}, nil
	case bindings.SIMCONNECT_RECV_ID_CLIENT_DATA:
		return core.ClientDataMessage{SimObjectDataMessage: decodeSimObjectData(raw, size)}, nil
	case bindings.SIMCONNECT_RECV_ID_RESERVED_KEY:
		msg := (*bindings.SIMCONNECT_RECV_RESERVED_KEY)(unsafe.Pointer(raw))
		return core.ReservedKeyMessage{ChoiceReserved: fixedString(msg.SzChoiceReserved[:]), ReservedKey: fixedString(msg.SzReservedKey[:])}, nil
	case bindings.SIMCONNECT_RECV_ID_SYSTEM_STATE:
		msg := (*bindings.SIMCONNECT_RECV_SYSTEM_STATE)(unsafe.Pointer(raw))
		return core.SystemStateMessage{
			RequestID: msg.DwRequestID,
			Integer:   msg.DwInteger,
			Float:     msg.FFloat,
			String:    fixedString(msg.SzString[:]),
		}, nil
	case bindings.SIMCONNECT_RECV_ID_ASSIGNED_OBJECT_ID:
		msg := (*bindings.SIMCONNECT_RECV_ASSIGNED_OBJECT_ID)(unsafe.Pointer(raw))
		return core.AssignedObjectIDMessage{
			RequestID: msg.DwRequestID,
			ObjectID:  core.ObjectID(msg.DwObjectID),
		}, nil
	case bindings.SIMCONNECT_RECV_ID_ACTION_CALLBACK:
		msg := (*bindings.SIMCONNECT_RECV_ACTION_CALLBACK)(unsafe.Pointer(raw))
		return core.ActionCallbackMessage{RequestID: msg.CbRequestId, ActionID: fixedString(msg.SzActionID[:])}, nil
	case bindings.SIMCONNECT_RECV_ID_AIRPORT_LIST:
		return decodeAirportList(raw, size)
	case bindings.SIMCONNECT_RECV_ID_WAYPOINT_LIST:
		return decodeWaypointList(raw, size)
	case bindings.SIMCONNECT_RECV_ID_NDB_LIST:
		return decodeNDBList(raw, size)
	case bindings.SIMCONNECT_RECV_ID_VOR_LIST:
		return decodeVORList(raw, size)
	case bindings.SIMCONNECT_RECV_ID_FACILITY_DATA:
		return decodeFacilityData(raw, size), nil
	case bindings.SIMCONNECT_RECV_ID_FACILITY_DATA_END:
		msg := (*bindings.SIMCONNECT_RECV_FACILITY_DATA_END)(unsafe.Pointer(raw))
		return core.FacilityDataEndMessage{RequestID: msg.RequestId}, nil
	case bindings.SIMCONNECT_RECV_ID_FACILITY_MINIMAL_LIST:
		return decodeFacilityMinimalList(raw, size)
	case bindings.SIMCONNECT_RECV_ID_JETWAY_DATA:
		return decodeJetwayData(raw, size)
	case bindings.SIMCONNECT_RECV_ID_ENUMERATE_INPUT_EVENTS:
		return decodeInputEventList(raw, size)
	case bindings.SIMCONNECT_RECV_ID_GET_INPUT_EVENT:
		return decodeInputEventValue(raw, size), nil
	case bindings.SIMCONNECT_RECV_ID_SUBSCRIBE_INPUT_EVENT:
		return decodeInputEventSubscription(raw, size), nil
	case bindings.SIMCONNECT_RECV_ID_ENUMERATE_INPUT_EVENT_PARAMS:
		return decodeInputEventParams(raw), nil
	case bindings.SIMCONNECT_RECV_ID_CONTROLLERS_LIST:
		return decodeControllersList(raw, size)
	case bindings.SIMCONNECT_RECV_ID_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST:
		return decodeSimObjectLiveryList(raw, size)
	case bindings.SIMCONNECT_RECV_ID_FLOW_EVENT:
		msg := (*bindings.SIMCONNECT_RECV_FLOW_EVENT)(unsafe.Pointer(raw))
		return core.FlowEventMessage{Event: core.FlowEvent(msg.FlowEvent), FLTPath: fixedString(msg.FltPath[:])}, nil
	case bindings.SIMCONNECT_RECV_ID_COMM_BUS:
		return decodeCommBus(raw, size)
	case bindings.SIMCONNECT_RECV_ID_CAMERA_DATA:
		msg := (*bindings.SIMCONNECT_RECV_CAMERA_DATA)(unsafe.Pointer(raw))
		return core.CameraDataMessage{CameraData: cameraDataFromBinding(msg.CameraData)}, nil
	case bindings.SIMCONNECT_RECV_ID_CAMERA_STATUS:
		msg := (*bindings.SIMCONNECT_RECV_CAMERA_STATUS)(unsafe.Pointer(raw))
		return core.CameraStatusMessage{
			AcquiredState:  core.CameraAvailability(msg.AcquiredState),
			GameControlled: msg.BGameControlled != 0,
		}, nil
	case bindings.SIMCONNECT_RECV_ID_CAMERA_DEFINITION_LIST:
		return decodeCameraDefinitionList(raw, size)
	case bindings.SIMCONNECT_RECV_ID_CAMERA_WORLD_LOCKER:
		msg := (*bindings.SIMCONNECT_RECV_CAMERA_WORLD_LOCKER)(unsafe.Pointer(raw))
		return core.CameraWorldLockerMessage{Status: core.CameraWorldLockerStatus(msg.Status)}, nil
	default:
		return rawMessage{id: core.RecvID(raw.DwID), raw: raw, size: size}, nil
	}
}

func decodeSimObjectData(raw *bindings.SIMCONNECT_RECV, size uint32) core.SimObjectDataMessage {
	msg := (*bindings.SIMCONNECT_RECV_SIMOBJECT_DATA)(unsafe.Pointer(raw))
	offset := unsafe.Offsetof(msg.DwData)
	var payload []byte
	if uintptr(size) > offset {
		payload = unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(msg), offset)), int(uintptr(size)-offset))
	}
	return core.SimObjectDataMessage{
		RequestID:   msg.DwRequestID,
		ObjectID:    msg.DwObjectID,
		DefineID:    msg.DwDefineID,
		Flags:       msg.DwFlags,
		EntryNumber: msg.DwEntryNumber,
		OutOf:       msg.DwOutOf,
		DefineCount: msg.DwDefineCount,
		Payload:     payload,
	}
}

func decodeAirportList(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_AIRPORT_LIST)(unsafe.Pointer(raw))
	airports, err := decodeFixedArray[bindings.SIMCONNECT_DATA_FACILITY_AIRPORT](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	return core.AirportListMessage{FacilityListMeta: facilityListMeta(msg.SIMCONNECT_RECV_FACILITIES_LIST), Airports: airportsFromBinding(airports)}, nil
}

func decodeWaypointList(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_WAYPOINT_LIST)(unsafe.Pointer(raw))
	waypoints, err := decodeFixedArray[bindings.SIMCONNECT_DATA_FACILITY_WAYPOINT](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	return core.WaypointListMessage{FacilityListMeta: facilityListMeta(msg.SIMCONNECT_RECV_FACILITIES_LIST), Waypoints: waypointsFromBinding(waypoints)}, nil
}

func decodeNDBList(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_NDB_LIST)(unsafe.Pointer(raw))
	ndbs, err := decodeFixedArray[bindings.SIMCONNECT_DATA_FACILITY_NDB](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	return core.NDBListMessage{FacilityListMeta: facilityListMeta(msg.SIMCONNECT_RECV_FACILITIES_LIST), NDBs: ndbsFromBinding(ndbs)}, nil
}

func decodeVORList(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_VOR_LIST)(unsafe.Pointer(raw))
	vors, err := decodeFixedArray[bindings.SIMCONNECT_DATA_FACILITY_VOR](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	return core.VORListMessage{FacilityListMeta: facilityListMeta(msg.SIMCONNECT_RECV_FACILITIES_LIST), VORs: vorsFromBinding(vors)}, nil
}

func decodeFacilityData(raw *bindings.SIMCONNECT_RECV, size uint32) core.FacilityDataMessage {
	msg := (*bindings.SIMCONNECT_RECV_FACILITY_DATA)(unsafe.Pointer(raw))
	offset := unsafe.Offsetof(msg.Data)
	var payload []byte
	if uintptr(size) > offset {
		payload = append([]byte(nil), unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(msg), offset)), int(uintptr(size)-offset))...)
	}
	return core.FacilityDataMessage{
		UserRequestID:         msg.UserRequestId,
		UniqueRequestID:       msg.UniqueRequestId,
		ParentUniqueRequestID: msg.ParentUniqueRequestId,
		Type:                  core.FacilityDataType(msg.Type),
		IsListItem:            msg.IsListItem != 0,
		ItemIndex:             msg.ItemIndex,
		ListSize:              msg.ListSize,
		Payload:               payload,
	}
}

func decodeFacilityMinimalList(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_FACILITY_MINIMAL_LIST)(unsafe.Pointer(raw))
	facilities, err := decodeFixedArray[bindings.SIMCONNECT_FACILITY_MINIMAL](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	return core.FacilityMinimalListMessage{FacilityListMeta: listTemplateMeta(msg.SIMCONNECT_RECV_LIST_TEMPLATE), Facilities: facilitiesFromBinding(facilities)}, nil
}

func decodeJetwayData(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_JETWAY_DATA)(unsafe.Pointer(raw))
	jetways, err := decodeFixedArray[bindings.SIMCONNECT_JETWAY_DATA](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	return core.JetwayDataMessage{FacilityListMeta: listTemplateMeta(msg.SIMCONNECT_RECV_LIST_TEMPLATE), Jetways: jetwaysFromBinding(jetways)}, nil
}

func decodeInputEventList(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS)(unsafe.Pointer(raw))
	rawEvents, err := decodeFixedArray[bindings.SIMCONNECT_INPUT_EVENT_DESCRIPTOR](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	events := make([]core.InputEventDescriptor, len(rawEvents))
	for i, event := range rawEvents {
		events[i] = core.InputEventDescriptor{
			Name: fixedString(event.Name[:]),
			Hash: event.Hash.Uint64(),
			Type: core.InputEventType(event.EType),
		}
	}
	return core.InputEventListMessage{FacilityListMeta: listTemplateMeta(msg.SIMCONNECT_RECV_LIST_TEMPLATE), Events: events}, nil
}

func decodeInputEventValue(raw *bindings.SIMCONNECT_RECV, size uint32) core.InputEventValueMessage {
	msg := (*bindings.SIMCONNECT_RECV_GET_INPUT_EVENT)(unsafe.Pointer(raw))
	offset := unsafe.Offsetof(msg.Value)
	var payload []byte
	if uintptr(size) > offset {
		payload = append([]byte(nil), unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(msg), offset)), int(uintptr(size)-offset))...)
	}
	value := core.InputEventValueMessage{
		RequestID: msg.DwRequestID,
		Type:      core.InputEventType(msg.EType),
		Payload:   payload,
	}
	switch msg.EType {
	case bindings.SIMCONNECT_INPUT_EVENT_TYPE_DOUBLE:
		if len(payload) >= 8 {
			value.Double = math.Float64frombits(binary.LittleEndian.Uint64(payload[:8]))
		}
	case bindings.SIMCONNECT_INPUT_EVENT_TYPE_STRING:
		value.String = fixedString(payload)
	}
	return value
}

func decodeInputEventSubscription(raw *bindings.SIMCONNECT_RECV, size uint32) core.InputEventSubscriptionMessage {
	msg := (*bindings.SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT)(unsafe.Pointer(raw))
	offset := unsafe.Offsetof(msg.Value)
	var payload []byte
	if uintptr(size) > offset {
		payload = append([]byte(nil), unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(msg), offset)), int(uintptr(size)-offset))...)
	}
	value := core.InputEventSubscriptionMessage{
		Hash:    msg.Hash.Uint64(),
		Type:    core.InputEventType(msg.EType),
		Payload: payload,
	}
	switch msg.EType {
	case bindings.SIMCONNECT_INPUT_EVENT_TYPE_DOUBLE:
		if len(payload) >= 8 {
			value.Double = math.Float64frombits(binary.LittleEndian.Uint64(payload[:8]))
		}
	case bindings.SIMCONNECT_INPUT_EVENT_TYPE_STRING:
		value.String = fixedString(payload)
	}
	return value
}

func decodeInputEventParams(raw *bindings.SIMCONNECT_RECV) core.InputEventParamsMessage {
	msg := (*bindings.SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS)(unsafe.Pointer(raw))
	value := fixedString(msg.Value[:])
	params := splitInputEventParams(value)
	return core.InputEventParamsMessage{Hash: msg.Hash.Uint64(), Value: value, Params: params}
}

func decodeControllersList(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_CONTROLLERS_LIST)(unsafe.Pointer(raw))
	rawControllers, err := decodeFixedArray[bindings.SIMCONNECT_CONTROLLER_ITEM](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	controllers := make([]core.ControllerItem, len(rawControllers))
	for i, controller := range rawControllers {
		controllers[i] = core.ControllerItem{
			DeviceName:      fixedString(controller.DeviceName[:]),
			DeviceID:        controller.DeviceId,
			ProductID:       controller.ProductId,
			CompositeID:     controller.CompositeID,
			HardwareVersion: versionFromBinding(controller.HardwareVersion),
		}
	}
	return core.ControllersListMessage{FacilityListMeta: listTemplateMeta(msg.SIMCONNECT_RECV_LIST_TEMPLATE), Controllers: controllers}, nil
}

func decodeSimObjectLiveryList(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST)(unsafe.Pointer(raw))
	rawLiveries, err := decodeFixedArray[bindings.SIMCONNECT_ENUMERATE_SIMOBJECT_LIVERY](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	liveries := make([]core.SimObjectLivery, len(rawLiveries))
	for i, livery := range rawLiveries {
		liveries[i] = core.SimObjectLivery{
			AircraftTitle: fixedString(livery.AircraftTitle[:]),
			LiveryName:    fixedString(livery.LiveryName[:]),
		}
	}
	return core.SimObjectLiveryListMessage{FacilityListMeta: listTemplateMeta(msg.SIMCONNECT_RECV_LIST_TEMPLATE), Liveries: liveries}, nil
}

func decodeCommBus(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_COMM_BUS)(unsafe.Pointer(raw))
	offset := unsafe.Offsetof(msg.RgData)
	if uintptr(size) < offset {
		return nil, fmt.Errorf("%w: packet too small for comm bus data", core.ErrDecode)
	}
	payload := append([]byte(nil), unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(msg), offset)), int(uintptr(size)-offset))...)
	return core.CommBusMessage{
		FacilityListMeta: listTemplateMeta(msg.SIMCONNECT_RECV_LIST_TEMPLATE),
		EventID:          msg.UEventID,
		Data:             fixedString(payload),
		Payload:          payload,
	}, nil
}

func decodeCameraDefinitionList(raw *bindings.SIMCONNECT_RECV, size uint32) (core.Message, error) {
	msg := (*bindings.SIMCONNECT_RECV_CAMERA_DEFINITION_LIST)(unsafe.Pointer(raw))
	rawDefinitions, err := decodeFixedArray[bindings.SIMCONNECT_CAMERA_DEFINITION_ITEM](
		unsafe.Pointer(&msg.RgData[0]), msg.DwArraySize, unsafe.Offsetof(msg.RgData), size,
	)
	if err != nil {
		return nil, err
	}
	definitions := make([]string, len(rawDefinitions))
	for i, definition := range rawDefinitions {
		definitions[i] = fixedString(definition.Str[:])
	}
	return core.CameraDefinitionListMessage{FacilityListMeta: listTemplateMeta(msg.SIMCONNECT_RECV_LIST_TEMPLATE), Definitions: definitions}, nil
}

func cameraDataFromBinding(data bindings.SIMCONNECT_DATA_CAMERA) core.CameraData {
	return core.CameraData{
		Position:                    xyzFromBinding(data.Position),
		PositionReferential:         core.PositionReferential(data.PositionReferential),
		PositionReferentialObjectID: core.ObjectID(data.PositionReferentialObjectId),
		TargetedPos:                 xyzFromBinding(data.TargetedPos),
		PBH:                         pbhFromBinding(data.Pbh),
		RotationReferential:         core.PositionReferential(data.RotationReferential),
		RotationReferentialObjectID: core.ObjectID(data.RotationReferentialObjectId),
		FOV:                         data.Fov.Float64(),
	}
}

func cameraDataToBinding(data core.CameraData) bindings.SIMCONNECT_DATA_CAMERA {
	raw := bindings.SIMCONNECT_DATA_CAMERA{
		Position:                    xyzToBinding(data.Position),
		PositionReferential:         bindings.SIMCONNECT_POSITION_REFERENTIAL(data.PositionReferential),
		PositionReferentialObjectId: uint32(data.PositionReferentialObjectID),
		TargetedPos:                 xyzToBinding(data.TargetedPos),
		Pbh:                         pbhToBinding(data.PBH),
		RotationReferential:         bindings.SIMCONNECT_POSITION_REFERENTIAL(data.RotationReferential),
		RotationReferentialObjectId: uint32(data.RotationReferentialObjectID),
	}
	raw.Fov.SetFloat64(data.FOV)
	return raw
}

func xyzFromBinding(data bindings.SIMCONNECT_PACKED_DATA_XYZ) core.XYZ {
	return core.XYZ{X: data.X.Float64(), Y: data.Y.Float64(), Z: data.Z.Float64()}
}

func dataXYZFromBinding(data bindings.SIMCONNECT_DATA_XYZ) core.XYZ {
	return core.XYZ{X: data.X, Y: data.Y, Z: data.Z}
}

func xyzToBinding(data core.XYZ) bindings.SIMCONNECT_PACKED_DATA_XYZ {
	var raw bindings.SIMCONNECT_PACKED_DATA_XYZ
	raw.X.SetFloat64(data.X)
	raw.Y.SetFloat64(data.Y)
	raw.Z.SetFloat64(data.Z)
	return raw
}

func dataXYZToBinding(data core.XYZ) bindings.SIMCONNECT_DATA_XYZ {
	return bindings.SIMCONNECT_DATA_XYZ{X: data.X, Y: data.Y, Z: data.Z}
}

func pbhFromBinding(data bindings.SIMCONNECT_DATA_PBH) core.PBH {
	return core.PBH{Pitch: data.Pitch, Bank: data.Bank, Heading: data.Heading}
}

func pbhToBinding(data core.PBH) bindings.SIMCONNECT_DATA_PBH {
	return bindings.SIMCONNECT_DATA_PBH{Pitch: data.Pitch, Bank: data.Bank, Heading: data.Heading}
}

func initPositionToBinding(data core.InitPosition) bindings.SIMCONNECT_DATA_INITPOSITION {
	return bindings.SIMCONNECT_DATA_INITPOSITION{
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Altitude:  data.Altitude,
		Pitch:     data.Pitch,
		Bank:      data.Bank,
		Heading:   data.Heading,
		OnGround:  data.OnGround,
		Airspeed:  data.Airspeed,
	}
}

func airportsFromBinding(in []bindings.SIMCONNECT_DATA_FACILITY_AIRPORT) []core.Airport {
	out := make([]core.Airport, len(in))
	for i, item := range in {
		out[i] = airportFromBinding(item)
	}
	return out
}

func airportFromBinding(item bindings.SIMCONNECT_DATA_FACILITY_AIRPORT) core.Airport {
	return core.Airport{
		Ident:     fixedString(item.Ident[:]),
		Region:    fixedString(item.Region[:]),
		Latitude:  item.Latitude.Float64(),
		Longitude: item.Longitude.Float64(),
		Altitude:  item.Altitude.Float64(),
	}
}

func waypointsFromBinding(in []bindings.SIMCONNECT_DATA_FACILITY_WAYPOINT) []core.Waypoint {
	out := make([]core.Waypoint, len(in))
	for i, item := range in {
		out[i] = waypointFromBinding(item)
	}
	return out
}

func waypointFromBinding(item bindings.SIMCONNECT_DATA_FACILITY_WAYPOINT) core.Waypoint {
	return core.Waypoint{Airport: airportFromBinding(item.SIMCONNECT_DATA_FACILITY_AIRPORT), MagVar: item.FMagVar}
}

func ndbsFromBinding(in []bindings.SIMCONNECT_DATA_FACILITY_NDB) []core.NDB {
	out := make([]core.NDB, len(in))
	for i, item := range in {
		out[i] = ndbFromBinding(item)
	}
	return out
}

func ndbFromBinding(item bindings.SIMCONNECT_DATA_FACILITY_NDB) core.NDB {
	return core.NDB{Waypoint: waypointFromBinding(item.SIMCONNECT_DATA_FACILITY_WAYPOINT), Frequency: item.FFrequency}
}

func vorsFromBinding(in []bindings.SIMCONNECT_DATA_FACILITY_VOR) []core.VOR {
	out := make([]core.VOR, len(in))
	for i, item := range in {
		out[i] = core.VOR{
			NDB:             ndbFromBinding(item.SIMCONNECT_DATA_FACILITY_NDB),
			Flags:           item.Flags,
			Localizer:       item.FLocalizer,
			GlideLat:        item.GlideLat.Float64(),
			GlideLon:        item.GlideLon.Float64(),
			GlideAlt:        item.GlideAlt.Float64(),
			GlideSlopeAngle: item.FGlideSlopeAngle,
		}
	}
	return out
}

func facilitiesFromBinding(in []bindings.SIMCONNECT_FACILITY_MINIMAL) []core.FacilityMinimal {
	out := make([]core.FacilityMinimal, len(in))
	for i, item := range in {
		out[i] = core.FacilityMinimal{
			ICAO: icaoFromBinding(item.Icao),
			LLA:  latLonAltFromPacked(item.Lla),
		}
	}
	return out
}

func jetwaysFromBinding(in []bindings.SIMCONNECT_JETWAY_DATA) []core.JetwayData {
	out := make([]core.JetwayData, len(in))
	for i, item := range in {
		out[i] = core.JetwayData{
			AirportICAO:         fixedString(item.AirportIcao[:]),
			ParkingIndex:        item.ParkingIndex,
			LLA:                 latLonAltFromPacked(item.Lla),
			PBH:                 pbhFromBinding(item.Pbh),
			Status:              item.Status,
			Door:                item.Door,
			ExitDoorRelativePos: xyzFromBinding(item.ExitDoorRelativePos),
			MainHandlePos:       xyzFromBinding(item.MainHandlePos),
			SecondaryHandle:     xyzFromBinding(item.SecondaryHandle),
			WheelGroundLock:     xyzFromBinding(item.WheelGroundLock),
			JetwayObjectID:      core.ObjectID(item.JetwayObjectId),
			AttachedObjectID:    core.ObjectID(item.AttachedObjectId),
		}
	}
	return out
}

func icaoFromBinding(item bindings.SIMCONNECT_ICAO) core.ICAO {
	return core.ICAO{
		Type:    item.Type,
		Ident:   fixedString(item.Ident[:]),
		Region:  fixedString(item.Region[:]),
		Airport: fixedString(item.Airport[:]),
	}
}

func latLonAltFromPacked(item bindings.SIMCONNECT_PACKED_DATA_LATLONALT) core.LatLonAlt {
	return core.LatLonAlt{
		Latitude:  item.Latitude.Float64(),
		Longitude: item.Longitude.Float64(),
		Altitude:  item.Altitude.Float64(),
	}
}

func versionFromBinding(item bindings.SIMCONNECT_VERSION_BASE_TYPE) core.Version {
	return core.Version{
		Major:    item.Major,
		Minor:    item.Minor,
		Revision: item.Revision,
		Build:    item.Build,
	}
}

func splitInputEventParams(value string) []string {
	if value == "" {
		return nil
	}
	parts := bytes.Split([]byte(value), []byte(";"))
	params := make([]string, 0, len(parts))
	for _, part := range parts {
		if len(part) > 0 {
			params = append(params, string(part))
		}
	}
	return params
}

func facilityListMeta(msg bindings.SIMCONNECT_RECV_FACILITIES_LIST) core.FacilityListMeta {
	return core.FacilityListMeta{
		RequestID:   msg.DwRequestID,
		ArraySize:   msg.DwArraySize,
		EntryNumber: msg.DwEntryNumber,
		OutOf:       msg.DwOutOf,
	}
}

func listTemplateMeta(msg bindings.SIMCONNECT_RECV_LIST_TEMPLATE) core.FacilityListMeta {
	return core.FacilityListMeta{
		RequestID:   msg.DwRequestID,
		ArraySize:   msg.DwArraySize,
		EntryNumber: msg.DwEntryNumber,
		OutOf:       msg.DwOutOf,
	}
}

func decodeFixedArray[T any](data unsafe.Pointer, count uint32, offset uintptr, size uint32) ([]T, error) {
	if count == 0 {
		return nil, nil
	}
	recordSize := unsafe.Sizeof(*new(T))
	needed := offset + uintptr(count)*recordSize
	if uintptr(size) < needed {
		return nil, fmt.Errorf("%w: packet too small for facility array", core.ErrDecode)
	}
	src := unsafe.Slice((*T)(data), int(count))
	dst := make([]T, len(src))
	copy(dst, src)
	return dst, nil
}

func fixedString(b []byte) string {
	before, _, _ := bytes.Cut(b, []byte{0})
	return string(before)
}
