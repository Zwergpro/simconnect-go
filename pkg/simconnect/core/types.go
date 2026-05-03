//go:build windows

// Package core contains the shared protocol types, errors, and message
// definitions used by both the client session package and all domain facets.
// Placing them here breaks the import cycle that would otherwise arise when
// pkg/simconnect/client imports the facet packages (for facade accessors) while
// the facets import pkg/simconnect/client for shared types.
package core

import "github.com/Zwergpro/simconnect-go/pkg/bindings"

// ─── Primitive type aliases ──────────────────────────────────────────────────

type RecvID uint32
type ObjectID uint32
type Period uint32
type NotificationGroupID uint32
type InputGroupID uint32
type DataDefinitionID uint32
type DataRequestID uint32
type EventID uint32
type State uint32
type DataType int32
type Exception uint32
type InputEventType uint32
type SimObjectType uint32
type DataRequestFlag uint32
type DataSetFlag uint32
type EventFlag uint32
type ClientDataID uint32
type ClientDataDefinitionID uint32
type ClientDataPeriod uint32
type ClientDataRequestFlag uint32
type ClientDataSetFlag uint32
type ClientDataCreateFlag uint32
type FlowEvent uint32
type FacilityListType uint32
type FacilityDataType uint32
type CommBusBroadcastTo uint32
type PositionReferential uint32
type CameraDataMask uint32
type CameraFlag uint32
type CameraAvailability uint32
type CameraWorldLockerStatus uint32

// ─── Struct types ────────────────────────────────────────────────────────────

type InitPosition struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
	Pitch     float64
	Bank      float64
	Heading   float64
	OnGround  uint32
	Airspeed  uint32
}

type XYZ struct {
	X float64
	Y float64
	Z float64
}

type PBH struct {
	Pitch   float32
	Bank    float32
	Heading float32
}

type LatLonAlt struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

type ICAO struct {
	Type    byte
	Ident   string
	Region  string
	Airport string
}

type Airport struct {
	Ident     string
	Region    string
	Latitude  float64
	Longitude float64
	Altitude  float64
}

type Waypoint struct {
	Airport
	MagVar float32
}

type NDB struct {
	Waypoint
	Frequency uint32
}

type VOR struct {
	NDB
	Flags           uint32
	Localizer       float32
	GlideLat        float64
	GlideLon        float64
	GlideAlt        float64
	GlideSlopeAngle float32
}

type FacilityMinimal struct {
	ICAO ICAO
	LLA  LatLonAlt
}

type JetwayData struct {
	AirportICAO         string
	ParkingIndex        int32
	LLA                 LatLonAlt
	PBH                 PBH
	Status              int32
	Door                int32
	ExitDoorRelativePos XYZ
	MainHandlePos       XYZ
	SecondaryHandle     XYZ
	WheelGroundLock     XYZ
	JetwayObjectID      ObjectID
	AttachedObjectID    ObjectID
}

type Version struct {
	Major    uint16
	Minor    uint16
	Revision uint16
	Build    uint16
}

// ─── RecvID constants ────────────────────────────────────────────────────────

const (
	RecvIDNull RecvID = iota
	RecvIDException
	RecvIDOpen
	RecvIDQuit
	RecvIDEvent
	RecvIDEventObjectAddRemove
	RecvIDEventFilename
	RecvIDEventFrame
	RecvIDSimObjectData
	RecvIDSimObjectDataByType
	RecvIDWeatherObservation
	RecvIDCloudState
	RecvIDAssignedObjectID
	RecvIDReservedKey
	RecvIDCustomAction
	RecvIDSystemState
	RecvIDClientData
	RecvIDEventWeatherMode
	RecvIDAirportList
	RecvIDVORList
	RecvIDNDBList
	RecvIDWaypointList
	RecvIDEventMultiplayerServerStarted
	RecvIDEventMultiplayerClientStarted
	RecvIDEventMultiplayerSessionEnded
	RecvIDEventRaceEnd
	RecvIDEventRaceLap
	RecvIDEventEX1
	RecvIDFacilityData
	RecvIDFacilityDataEnd
	RecvIDFacilityMinimalList
	RecvIDJetwayData
	RecvIDControllersList
	RecvIDActionCallback
	RecvIDEnumerateInputEvents
	RecvIDGetInputEvent
	RecvIDSubscribeInputEvent
	RecvIDEnumerateInputEventParams
	RecvIDEnumerateSimObjectAndLiveryList
	RecvIDFlowEvent
	RecvIDCameraData
	RecvIDCameraStatus
	RecvIDCameraDefinitionList
	RecvIDCommBus
	RecvIDCameraWorldLocker
)

// ─── Object ID constants ─────────────────────────────────────────────────────

const (
	UserAircraft ObjectID = ObjectID(bindings.SIMCONNECT_OBJECT_ID_USER_AIRCRAFT)
	UserAvatar   ObjectID = ObjectID(bindings.SIMCONNECT_OBJECT_ID_USER_AVATAR)
	UserCurrent  ObjectID = ObjectID(bindings.SIMCONNECT_OBJECT_ID_USER_CURRENT)

	Unused               = bindings.SIMCONNECT_UNUSED
	ConfigIndexLocal     = bindings.SIMCONNECT_OPEN_CONFIGINDEX_LOCAL
	ClientDataMaxSize    = bindings.SIMCONNECT_CLIENTDATA_MAX_SIZE
	ClientDataOffsetAuto = bindings.SIMCONNECT_CLIENTDATAOFFSET_AUTO

	InitPositionAirspeedCruise = bindings.INITPOSITION_AIRSPEED_CRUISE
	InitPositionAirspeedKeep   = bindings.INITPOSITION_AIRSPEED_KEEP
)

// ─── Period constants ────────────────────────────────────────────────────────

const (
	PeriodOnce   Period = Period(bindings.SIMCONNECT_PERIOD_ONCE)
	PeriodNever  Period = Period(bindings.SIMCONNECT_PERIOD_NEVER)
	PeriodSecond Period = Period(bindings.SIMCONNECT_PERIOD_SECOND)
	PeriodFrame  Period = Period(bindings.SIMCONNECT_PERIOD_SIM_FRAME)
	PeriodVisual Period = Period(bindings.SIMCONNECT_PERIOD_VISUAL_FRAME)
)

// ─── State constants ─────────────────────────────────────────────────────────

const (
	StateOff State = State(bindings.SIMCONNECT_STATE_OFF)
	StateOn  State = State(bindings.SIMCONNECT_STATE_ON)
)

// ─── DataType constants ──────────────────────────────────────────────────────

const (
	DataTypeInvalid      DataType = DataType(bindings.SIMCONNECT_DATATYPE_INVALID)
	DataTypeInt32        DataType = DataType(bindings.SIMCONNECT_DATATYPE_INT32)
	DataTypeInt64        DataType = DataType(bindings.SIMCONNECT_DATATYPE_INT64)
	DataTypeFloat32      DataType = DataType(bindings.SIMCONNECT_DATATYPE_FLOAT32)
	DataTypeFloat64      DataType = DataType(bindings.SIMCONNECT_DATATYPE_FLOAT64)
	DataTypeString8      DataType = DataType(bindings.SIMCONNECT_DATATYPE_STRING8)
	DataTypeString32     DataType = DataType(bindings.SIMCONNECT_DATATYPE_STRING32)
	DataTypeString64     DataType = DataType(bindings.SIMCONNECT_DATATYPE_STRING64)
	DataTypeString128    DataType = DataType(bindings.SIMCONNECT_DATATYPE_STRING128)
	DataTypeString256    DataType = DataType(bindings.SIMCONNECT_DATATYPE_STRING256)
	DataTypeString260    DataType = DataType(bindings.SIMCONNECT_DATATYPE_STRING260)
	DataTypeStringV      DataType = DataType(bindings.SIMCONNECT_DATATYPE_STRINGV)
	DataTypeInitPosition DataType = DataType(bindings.SIMCONNECT_DATATYPE_INITPOSITION)
	DataTypeMarkerState  DataType = DataType(bindings.SIMCONNECT_DATATYPE_MARKERSTATE)
	DataTypeWaypoint     DataType = DataType(bindings.SIMCONNECT_DATATYPE_WAYPOINT)
	DataTypeLatLonAlt    DataType = DataType(bindings.SIMCONNECT_DATATYPE_LATLONALT)
	DataTypeXYZ          DataType = DataType(bindings.SIMCONNECT_DATATYPE_XYZ)
	DataTypeInt8         DataType = DataType(bindings.SIMCONNECT_DATATYPE_INT8)
)

// ─── InputEventType constants ────────────────────────────────────────────────

const (
	InputEventTypeDouble InputEventType = InputEventType(bindings.SIMCONNECT_INPUT_EVENT_TYPE_DOUBLE)
	InputEventTypeString InputEventType = InputEventType(bindings.SIMCONNECT_INPUT_EVENT_TYPE_STRING)
)

// ─── Priority/group constants ────────────────────────────────────────────────

const (
	GroupPriorityHighest         = bindings.SIMCONNECT_GROUP_PRIORITY_HIGHEST
	GroupPriorityHighestMaskable = bindings.SIMCONNECT_GROUP_PRIORITY_HIGHEST_MASKABLE
	GroupPriorityStandard        = bindings.SIMCONNECT_GROUP_PRIORITY_STANDARD
	GroupPriorityDefault         = bindings.SIMCONNECT_GROUP_PRIORITY_DEFAULT
	GroupPriorityLowest          = bindings.SIMCONNECT_GROUP_PRIORITY_LOWEST
)

// ─── SimObjectType constants ─────────────────────────────────────────────────

const (
	SimObjectTypeUser          SimObjectType = SimObjectType(bindings.SIMCONNECT_SIMOBJECT_TYPE_USER)
	SimObjectTypeAll           SimObjectType = SimObjectType(bindings.SIMCONNECT_SIMOBJECT_TYPE_ALL)
	SimObjectTypeAircraft      SimObjectType = SimObjectType(bindings.SIMCONNECT_SIMOBJECT_TYPE_AIRCRAFT)
	SimObjectTypeHelicopter    SimObjectType = SimObjectType(bindings.SIMCONNECT_SIMOBJECT_TYPE_HELICOPTER)
	SimObjectTypeBoat          SimObjectType = SimObjectType(bindings.SIMCONNECT_SIMOBJECT_TYPE_BOAT)
	SimObjectTypeGround        SimObjectType = SimObjectType(bindings.SIMCONNECT_SIMOBJECT_TYPE_GROUND)
	SimObjectTypeHotAirBalloon SimObjectType = SimObjectType(bindings.SIMCONNECT_SIMOBJECT_TYPE_HOT_AIR_BALLOON)
	SimObjectTypeUserAvatar    SimObjectType = SimObjectType(bindings.SIMCONNECT_SIMOBJECT_TYPE_USER_AVATAR)
	SimObjectTypeUserCurrent   SimObjectType = SimObjectType(bindings.SIMCONNECT_SIMOBJECT_TYPE_USER_CURRENT)
)

// ─── DataRequestFlag / DataSetFlag / EventFlag constants ─────────────────────

const (
	DataRequestDefault DataRequestFlag = DataRequestFlag(bindings.SIMCONNECT_DATA_REQUEST_FLAG_DEFAULT)
	DataRequestChanged DataRequestFlag = DataRequestFlag(bindings.SIMCONNECT_DATA_REQUEST_FLAG_CHANGED)
	DataRequestTagged  DataRequestFlag = DataRequestFlag(bindings.SIMCONNECT_DATA_REQUEST_FLAG_TAGGED)

	DataSetDefault DataSetFlag = DataSetFlag(bindings.SIMCONNECT_DATA_SET_FLAG_DEFAULT)
	DataSetTagged  DataSetFlag = DataSetFlag(bindings.SIMCONNECT_DATA_SET_FLAG_TAGGED)

	EventFlagDefault           EventFlag = EventFlag(bindings.SIMCONNECT_EVENT_FLAG_DEFAULT)
	EventFlagFastRepeatTimer   EventFlag = EventFlag(bindings.SIMCONNECT_EVENT_FLAG_FAST_REPEAT_TIMER)
	EventFlagSlowRepeatTimer   EventFlag = EventFlag(bindings.SIMCONNECT_EVENT_FLAG_SLOW_REPEAT_TIMER)
	EventFlagGroupIDIsPriority EventFlag = EventFlag(bindings.SIMCONNECT_EVENT_FLAG_GROUPID_IS_PRIORITY)
)

// ─── ClientData constants ────────────────────────────────────────────────────

const (
	ClientDataTypeInt8    = bindings.SIMCONNECT_CLIENTDATATYPE_INT8
	ClientDataTypeInt16   = bindings.SIMCONNECT_CLIENTDATATYPE_INT16
	ClientDataTypeInt32   = bindings.SIMCONNECT_CLIENTDATATYPE_INT32
	ClientDataTypeInt64   = bindings.SIMCONNECT_CLIENTDATATYPE_INT64
	ClientDataTypeFloat32 = bindings.SIMCONNECT_CLIENTDATATYPE_FLOAT32
	ClientDataTypeFloat64 = bindings.SIMCONNECT_CLIENTDATATYPE_FLOAT64

	ClientDataCreateDefault  ClientDataCreateFlag = ClientDataCreateFlag(bindings.SIMCONNECT_CREATE_CLIENT_DATA_FLAG_DEFAULT)
	ClientDataCreateReadOnly ClientDataCreateFlag = ClientDataCreateFlag(bindings.SIMCONNECT_CREATE_CLIENT_DATA_FLAG_READ_ONLY)

	ClientDataPeriodNever       ClientDataPeriod = ClientDataPeriod(bindings.SIMCONNECT_CLIENT_DATA_PERIOD_NEVER)
	ClientDataPeriodOnce        ClientDataPeriod = ClientDataPeriod(bindings.SIMCONNECT_CLIENT_DATA_PERIOD_ONCE)
	ClientDataPeriodVisualFrame ClientDataPeriod = ClientDataPeriod(bindings.SIMCONNECT_CLIENT_DATA_PERIOD_VISUAL_FRAME)
	ClientDataPeriodOnSet       ClientDataPeriod = ClientDataPeriod(bindings.SIMCONNECT_CLIENT_DATA_PERIOD_ON_SET)
	ClientDataPeriodSecond      ClientDataPeriod = ClientDataPeriod(bindings.SIMCONNECT_CLIENT_DATA_PERIOD_SECOND)

	ClientDataRequestDefault ClientDataRequestFlag = ClientDataRequestFlag(bindings.SIMCONNECT_CLIENT_DATA_REQUEST_FLAG_DEFAULT)
	ClientDataRequestChanged ClientDataRequestFlag = ClientDataRequestFlag(bindings.SIMCONNECT_CLIENT_DATA_REQUEST_FLAG_CHANGED)
	ClientDataRequestTagged  ClientDataRequestFlag = ClientDataRequestFlag(bindings.SIMCONNECT_CLIENT_DATA_REQUEST_FLAG_TAGGED)

	ClientDataSetDefault ClientDataSetFlag = ClientDataSetFlag(bindings.SIMCONNECT_CLIENT_DATA_SET_FLAG_DEFAULT)
	ClientDataSetTagged  ClientDataSetFlag = ClientDataSetFlag(bindings.SIMCONNECT_CLIENT_DATA_SET_FLAG_TAGGED)
)

// ─── CommBus constants ───────────────────────────────────────────────────────

const (
	CommBusBroadcastToJS                 CommBusBroadcastTo = CommBusBroadcastTo(bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO_JS)
	CommBusBroadcastToWASM               CommBusBroadcastTo = CommBusBroadcastTo(bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO_WASM)
	CommBusBroadcastToSimConnect         CommBusBroadcastTo = CommBusBroadcastTo(bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT)
	CommBusBroadcastToSimConnectSelfCall CommBusBroadcastTo = CommBusBroadcastTo(bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT_SELF_CALL)
	CommBusBroadcastToDefault            CommBusBroadcastTo = CommBusBroadcastTo(bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO_DEFAULT)
	CommBusBroadcastToAllSimConnect      CommBusBroadcastTo = CommBusBroadcastTo(bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO_ALL_SIMCONNECT)
	CommBusBroadcastToAll                CommBusBroadcastTo = CommBusBroadcastTo(bindings.SIMCONNECT_COMM_BUS_BROADCAST_TO_ALL)
)

// ─── Facility constants ──────────────────────────────────────────────────────

const (
	FacilityListTypeAirport  FacilityListType = FacilityListType(bindings.SIMCONNECT_FACILITY_LIST_TYPE_AIRPORT)
	FacilityListTypeWaypoint FacilityListType = FacilityListType(bindings.SIMCONNECT_FACILITY_LIST_TYPE_WAYPOINT)
	FacilityListTypeNDB      FacilityListType = FacilityListType(bindings.SIMCONNECT_FACILITY_LIST_TYPE_NDB)
	FacilityListTypeVOR      FacilityListType = FacilityListType(bindings.SIMCONNECT_FACILITY_LIST_TYPE_VOR)

	FacilityDataAirport            FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_AIRPORT)
	FacilityDataRunway             FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_RUNWAY)
	FacilityDataStart              FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_START)
	FacilityDataFrequency          FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_FREQUENCY)
	FacilityDataHelipad            FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_HELIPAD)
	FacilityDataApproach           FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_APPROACH)
	FacilityDataApproachTransition FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_APPROACH_TRANSITION)
	FacilityDataApproachLeg        FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_APPROACH_LEG)
	FacilityDataFinalApproachLeg   FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_FINAL_APPROACH_LEG)
	FacilityDataMissedApproachLeg  FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_MISSED_APPROACH_LEG)
	FacilityDataDeparture          FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_DEPARTURE)
	FacilityDataArrival            FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_ARRIVAL)
	FacilityDataRunwayTransition   FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_RUNWAY_TRANSITION)
	FacilityDataEnrouteTransition  FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_ENROUTE_TRANSITION)
	FacilityDataTaxiPoint          FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_TAXI_POINT)
	FacilityDataTaxiParking        FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_TAXI_PARKING)
	FacilityDataTaxiPath           FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_TAXI_PATH)
	FacilityDataTaxiName           FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_TAXI_NAME)
	FacilityDataJetway             FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_JETWAY)
	FacilityDataVOR                FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_VOR)
	FacilityDataNDB                FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_NDB)
	FacilityDataWaypoint           FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_WAYPOINT)
	FacilityDataRoute              FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_ROUTE)
	FacilityDataPavement           FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_PAVEMENT)
	FacilityDataApproachLights     FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_APPROACH_LIGHTS)
	FacilityDataVASI               FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_VASI)
	FacilityDataVDGS               FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_VDGS)
	FacilityDataHoldingPattern     FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_HOLDING_PATTERN)
	FacilityDataTaxiParkingAirline FacilityDataType = FacilityDataType(bindings.SIMCONNECT_FACILITY_DATA_TAXI_PARKING_AIRLINE)
)

// ─── FlowEvent constants ─────────────────────────────────────────────────────

const (
	FlowEventNone             FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_NONE)
	FlowEventFLTLoad          FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_FLT_LOAD)
	FlowEventFLTLoaded        FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_FLT_LOADED)
	FlowEventTeleportStart    FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_TELEPORT_START)
	FlowEventTeleportDone     FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_TELEPORT_DONE)
	FlowEventBackOnTrackStart FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_BACK_ON_TRACK_START)
	FlowEventBackOnTrackDone  FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_BACK_ON_TRACK_DONE)
	FlowEventSkipStart        FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_SKIP_START)
	FlowEventSkipDone         FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_SKIP_DONE)
	FlowEventBackToMainMenu   FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_BACK_TO_MAIN_MENU)
	FlowEventRTCStart         FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_RTC_START)
	FlowEventRTCEnd           FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_RTC_END)
	FlowEventReplayStart      FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_REPLAY_START)
	FlowEventReplayEnd        FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_REPLAY_END)
	FlowEventFlightStart      FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_FLIGHT_START)
	FlowEventFlightEnd        FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_FLIGHT_END)
	FlowEventPlaneCrash       FlowEvent = FlowEvent(bindings.SIMCONNECT_FLOW_EVENT_PLANE_CRASH)
)

// ─── Camera constants ────────────────────────────────────────────────────────

const (
	PositionReferentialNone           PositionReferential = PositionReferential(bindings.SIMCONNECT_POSITION_REFERENTIAL_NONE)
	PositionReferentialSimObject      PositionReferential = PositionReferential(bindings.SIMCONNECT_POSITION_REFERENTIAL_SIMOBJECT)
	PositionReferentialWorld          PositionReferential = PositionReferential(bindings.SIMCONNECT_POSITION_REFERENTIAL_WORLD)
	PositionReferentialEyepoint       PositionReferential = PositionReferential(bindings.SIMCONNECT_POSITION_REFERENTIAL_EYEPOINT)
	PositionReferentialSimObjectDatum PositionReferential = PositionReferential(bindings.SIMCONNECT_POSITION_REFERENTIAL_SIMOBJECT_DATUM)

	CameraDataMaskNone        CameraDataMask = CameraDataMask(bindings.SIMCONNECT_CAMERA_DATA_MASK_NONE)
	CameraDataMaskPosition    CameraDataMask = CameraDataMask(bindings.SIMCONNECT_CAMERA_DATA_MASK_POSITION)
	CameraDataMaskRotation    CameraDataMask = CameraDataMask(bindings.SIMCONNECT_CAMERA_DATA_MASK_ROTATION)
	CameraDataMaskTargeted    CameraDataMask = CameraDataMask(bindings.SIMCONNECT_CAMERA_DATA_MASK_TARGETED)
	CameraDataMaskFOV         CameraDataMask = CameraDataMask(bindings.SIMCONNECT_CAMERA_DATA_MASK_FOV)
	CameraDataMaskAllRotation CameraDataMask = CameraDataMask(bindings.SIMCONNECT_CAMERA_DATA_MASK_ALL_ROTATION)
	CameraDataMaskAllTargeted CameraDataMask = CameraDataMask(bindings.SIMCONNECT_CAMERA_DATA_MASK_ALL_TARGETED)

	CameraFlagInteraction CameraFlag = CameraFlag(bindings.SIMCONNECT_CAMERA_FLAG_INTERACTION)
	CameraFlagAboveGround CameraFlag = CameraFlag(bindings.SIMCONNECT_CAMERA_FLAG_ABOVE_GROUND)
	CameraIgnoreField                = bindings.SIMCONNECT_CAMERA_IGNORE_FIELD

	CameraNotAcquired     CameraAvailability = CameraAvailability(bindings.SIMCONNECT_CAMERA_NOT_ACQUIRED)
	CameraAcquired        CameraAvailability = CameraAvailability(bindings.SIMCONNECT_CAMERA_ACQUIRED)
	CameraAcquiredByOther CameraAvailability = CameraAvailability(bindings.SIMCONNECT_CAMERA_ACQUIRED_BY_OTHER)
	CameraUserDisabled    CameraAvailability = CameraAvailability(bindings.SIMCONNECT_CAMERA_USER_DISABLED)

	CameraWorldLockerStatusNone    CameraWorldLockerStatus = CameraWorldLockerStatus(bindings.SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_NONE)
	CameraWorldLockerStatusStart   CameraWorldLockerStatus = CameraWorldLockerStatus(bindings.SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_START)
	CameraWorldLockerStatusSuccess CameraWorldLockerStatus = CameraWorldLockerStatus(bindings.SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_SUCCESS)
	CameraWorldLockerStatusCancel  CameraWorldLockerStatus = CameraWorldLockerStatus(bindings.SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_CANCEL)
	CameraWorldLockerStatusFail    CameraWorldLockerStatus = CameraWorldLockerStatus(bindings.SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_FAIL)
)

// ─── Exception constants ─────────────────────────────────────────────────────

const (
	ExceptionNone Exception = iota
	ExceptionErrorCode
	ExceptionSizeMismatch
	ExceptionUnrecognizedID
	ExceptionUnopened
	ExceptionVersionMismatch
	ExceptionTooManyGroups
	ExceptionNameUnrecognized
	ExceptionTooManyEventNames
	ExceptionEventIDDuplicate
	ExceptionTooManyMaps
	ExceptionTooManyObjects
	ExceptionTooManyRequests
	ExceptionWeatherInvalidPort
	ExceptionWeatherInvalidMETAR
	ExceptionWeatherUnableToGetObservation
	ExceptionWeatherUnableToCreateStation
	ExceptionWeatherUnableToRemoveStation
	ExceptionInvalidDataType
	ExceptionInvalidDataSize
	ExceptionDataError
	ExceptionInvalidArray
	ExceptionCreateObjectFailed
	ExceptionLoadFlightplanFailed
	ExceptionOperationInvalidForObjectType
	ExceptionIllegalOperation
	ExceptionAlreadySubscribed
	ExceptionInvalidEnum
	ExceptionDefinitionError
	ExceptionDuplicateID
	ExceptionDatumID
	ExceptionOutOfBounds
	ExceptionAlreadyCreated
	ExceptionObjectOutsideRealityBubble
	ExceptionObjectContainer
	ExceptionObjectAI
	ExceptionObjectATC
	ExceptionObjectSchedule
	ExceptionJetwayData
	ExceptionActionNotFound
	ExceptionNotAnAction
	ExceptionIncorrectActionParams
	ExceptionGetInputEventFailed
	ExceptionSetInputEventFailed
	ExceptionEventNameReserved
	ExceptionInternal
	ExceptionCameraAPI
)

// ─── String methods ──────────────────────────────────────────────────────────

func (id RecvID) String() string {
	if name, ok := recvIDNames[id]; ok {
		return name
	}
	return "UNKNOWN"
}

var recvIDNames = map[RecvID]string{
	RecvIDNull:                            "NULL",
	RecvIDException:                       "EXCEPTION",
	RecvIDOpen:                            "OPEN",
	RecvIDQuit:                            "QUIT",
	RecvIDEvent:                           "EVENT",
	RecvIDEventObjectAddRemove:            "EVENT_OBJECT_ADDREMOVE",
	RecvIDEventFilename:                   "EVENT_FILENAME",
	RecvIDEventFrame:                      "EVENT_FRAME",
	RecvIDSimObjectData:                   "SIMOBJECT_DATA",
	RecvIDSimObjectDataByType:             "SIMOBJECT_DATA_BYTYPE",
	RecvIDWeatherObservation:              "WEATHER_OBSERVATION",
	RecvIDCloudState:                      "CLOUD_STATE",
	RecvIDAssignedObjectID:                "ASSIGNED_OBJECT_ID",
	RecvIDReservedKey:                     "RESERVED_KEY",
	RecvIDCustomAction:                    "CUSTOM_ACTION",
	RecvIDSystemState:                     "SYSTEM_STATE",
	RecvIDClientData:                      "CLIENT_DATA",
	RecvIDEventWeatherMode:                "EVENT_WEATHER_MODE",
	RecvIDAirportList:                     "AIRPORT_LIST",
	RecvIDVORList:                         "VOR_LIST",
	RecvIDNDBList:                         "NDB_LIST",
	RecvIDWaypointList:                    "WAYPOINT_LIST",
	RecvIDEventMultiplayerServerStarted:   "EVENT_MULTIPLAYER_SERVER_STARTED",
	RecvIDEventMultiplayerClientStarted:   "EVENT_MULTIPLAYER_CLIENT_STARTED",
	RecvIDEventMultiplayerSessionEnded:    "EVENT_MULTIPLAYER_SESSION_ENDED",
	RecvIDEventRaceEnd:                    "EVENT_RACE_END",
	RecvIDEventRaceLap:                    "EVENT_RACE_LAP",
	RecvIDEventEX1:                        "EVENT_EX1",
	RecvIDFacilityData:                    "FACILITY_DATA",
	RecvIDFacilityDataEnd:                 "FACILITY_DATA_END",
	RecvIDFacilityMinimalList:             "FACILITY_MINIMAL_LIST",
	RecvIDJetwayData:                      "JETWAY_DATA",
	RecvIDControllersList:                 "CONTROLLERS_LIST",
	RecvIDActionCallback:                  "ACTION_CALLBACK",
	RecvIDEnumerateInputEvents:            "ENUMERATE_INPUT_EVENTS",
	RecvIDGetInputEvent:                   "GET_INPUT_EVENT",
	RecvIDSubscribeInputEvent:             "SUBSCRIBE_INPUT_EVENT",
	RecvIDEnumerateInputEventParams:       "ENUMERATE_INPUT_EVENT_PARAMS",
	RecvIDEnumerateSimObjectAndLiveryList: "ENUMERATE_SIMOBJECT_AND_LIVERY_LIST",
	RecvIDFlowEvent:                       "FLOW_EVENT",
	RecvIDCameraData:                      "CAMERA_DATA",
	RecvIDCameraStatus:                    "CAMERA_STATUS",
	RecvIDCameraDefinitionList:            "CAMERA_DEFINITION_LIST",
	RecvIDCommBus:                         "COMM_BUS",
	RecvIDCameraWorldLocker:               "CAMERA_WORLD_LOCKER",
}
