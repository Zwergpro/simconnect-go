//go:build windows

package bindings

// All structs below match the C layout from SimConnect.h with #pragma pack(1).
// C++ inheritance is expressed as embedding; static const members are omitted
// (they are class-level constants, not instance fields).
// Variable-length tail fields (SIMCONNECT_DATAV / SIMCONNECT_STRINGV /
// SIMCONNECT_FIXEDTYPE_DATAV) are represented as a single-element array so
// callers can take the address of the first element and cast to the real type.

// SIMCONNECT_RECV is the base of every received packet.
type SIMCONNECT_RECV struct {
	DwSize    uint32 // record size in bytes
	DwVersion uint32 // interface version
	DwID      uint32 // SIMCONNECT_RECV_ID
}

// SIMCONNECT_RECV_EXCEPTION carries error information (dwID == SIMCONNECT_RECV_ID_EXCEPTION).
type SIMCONNECT_RECV_EXCEPTION struct {
	SIMCONNECT_RECV
	DwException uint32 // SIMCONNECT_EXCEPTION
	DwSendID    uint32 // packet ID that caused the error (0 = unknown)
	DwIndex     uint32 // parameter index that caused the error
}

// SIMCONNECT_RECV_OPEN is sent on a successful SimConnect_Open (dwID == SIMCONNECT_RECV_ID_OPEN).
type SIMCONNECT_RECV_OPEN struct {
	SIMCONNECT_RECV
	SzApplicationName         [256]byte
	DwApplicationVersionMajor uint32
	DwApplicationVersionMinor uint32
	DwApplicationBuildMajor   uint32
	DwApplicationBuildMinor   uint32
	DwSimConnectVersionMajor  uint32
	DwSimConnectVersionMinor  uint32
	DwSimConnectBuildMajor    uint32
	DwSimConnectBuildMinor    uint32
	DwReserved1               uint32
	DwReserved2               uint32
}

// SIMCONNECT_RECV_QUIT is sent when the simulator exits (dwID == SIMCONNECT_RECV_ID_QUIT).
type SIMCONNECT_RECV_QUIT struct {
	SIMCONNECT_RECV
}

// SIMCONNECT_RECV_EVENT carries a client event (dwID == SIMCONNECT_RECV_ID_EVENT).
type SIMCONNECT_RECV_EVENT struct {
	SIMCONNECT_RECV
	UGroupID uint32 // notification group ID (DWORD_MAX = unknown)
	UEventID uint32 // client event ID
	DwData   uint32 // event-dependent data
}

// SIMCONNECT_RECV_EVENT_FILENAME extends SIMCONNECT_RECV_EVENT with a filename.
type SIMCONNECT_RECV_EVENT_FILENAME struct {
	SIMCONNECT_RECV_EVENT
	SzFileName [MAX_PATH]byte
	DwFlags    uint32
}

// SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE is sent when an AI object is added or removed.
type SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE struct {
	SIMCONNECT_RECV_EVENT
	EObjType SIMCONNECT_SIMOBJECT_TYPE
}

// SIMCONNECT_RECV_EVENT_FRAME carries frame-rate information (dwID == SIMCONNECT_RECV_ID_EVENT_FRAME).
type SIMCONNECT_RECV_EVENT_FRAME struct {
	SIMCONNECT_RECV_EVENT
	FFrameRate float32
	FSimSpeed  float32
}

// SIMCONNECT_RECV_EVENT_MULTIPLAYER_SERVER_STARTED signals MP server start.
type SIMCONNECT_RECV_EVENT_MULTIPLAYER_SERVER_STARTED struct {
	SIMCONNECT_RECV_EVENT
}

// SIMCONNECT_RECV_EVENT_MULTIPLAYER_CLIENT_STARTED signals MP client start.
type SIMCONNECT_RECV_EVENT_MULTIPLAYER_CLIENT_STARTED struct {
	SIMCONNECT_RECV_EVENT
}

// SIMCONNECT_RECV_EVENT_MULTIPLAYER_SESSION_ENDED signals MP session end.
type SIMCONNECT_RECV_EVENT_MULTIPLAYER_SESSION_ENDED struct {
	SIMCONNECT_RECV_EVENT
}

// SIMCONNECT_RECV_EVENT_EX1 carries an event with up to 5 DWORD data values.
type SIMCONNECT_RECV_EVENT_EX1 struct {
	SIMCONNECT_RECV
	UGroupID uint32
	UEventID uint32
	DwData0  uint32
	DwData1  uint32
	DwData2  uint32
	DwData3  uint32
	DwData4  uint32
}

// SIMCONNECT_DATA_RACE_RESULT holds race/lap result data.
type SIMCONNECT_DATA_RACE_RESULT struct {
	DwNumberOfRacers uint32
	MissionGUID      GUID
	SzPlayerName     [MAX_PATH]byte
	SzSessionType    [MAX_PATH]byte
	SzAircraft       [MAX_PATH]byte
	SzPlayerRole     [MAX_PATH]byte
	FTotalTime       SIMCONNECT_PACKED_FLOAT64
	FPenaltyTime     SIMCONNECT_PACKED_FLOAT64
	DwIsDisqualified uint32
}

// SIMCONNECT_RECV_EVENT_RACE_END carries end-of-race data.
type SIMCONNECT_RECV_EVENT_RACE_END struct {
	SIMCONNECT_RECV_EVENT
	DwRacerNumber uint32
	RacerData     SIMCONNECT_DATA_RACE_RESULT
}

// SIMCONNECT_RECV_EVENT_RACE_LAP carries lap data.
type SIMCONNECT_RECV_EVENT_RACE_LAP struct {
	SIMCONNECT_RECV_EVENT
	DwLapIndex uint32
	RacerData  SIMCONNECT_DATA_RACE_RESULT
}

// SIMCONNECT_RECV_SIMOBJECT_DATA carries sim-object data (dwID == SIMCONNECT_RECV_ID_SIMOBJECT_DATA).
// DwData is a one-element placeholder; the actual data follows it in memory.
type SIMCONNECT_RECV_SIMOBJECT_DATA struct {
	SIMCONNECT_RECV
	DwRequestID   uint32
	DwObjectID    uint32
	DwDefineID    uint32
	DwFlags       uint32
	DwEntryNumber uint32
	DwOutOf       uint32
	DwDefineCount uint32
	DwData        [1]uint32 // variable-length data starts here
}

// SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE is the by-type variant of the above.
type SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE struct {
	SIMCONNECT_RECV_SIMOBJECT_DATA
}

// SIMCONNECT_RECV_CLIENT_DATA carries client-data (dwID == SIMCONNECT_RECV_ID_CLIENT_DATA).
type SIMCONNECT_RECV_CLIENT_DATA struct {
	SIMCONNECT_RECV_SIMOBJECT_DATA
}

// SIMCONNECT_RECV_WEATHER_OBSERVATION carries METAR data.
// SzMetar is a variable-length string; only the first byte is declared here.
type SIMCONNECT_RECV_WEATHER_OBSERVATION struct {
	SIMCONNECT_RECV
	DwRequestID uint32
	SzMetar     [1]byte // variable-length; max MAX_METAR_LENGTH
}

// SIMCONNECT_RECV_CLOUD_STATE carries cloud-state data.
type SIMCONNECT_RECV_CLOUD_STATE struct {
	SIMCONNECT_RECV
	DwRequestID uint32
	DwArraySize uint32
	RgbData     [1]byte // variable-length; dwArraySize bytes
}

// SIMCONNECT_RECV_ASSIGNED_OBJECT_ID carries an assigned object ID.
type SIMCONNECT_RECV_ASSIGNED_OBJECT_ID struct {
	SIMCONNECT_RECV
	DwRequestID uint32
	DwObjectID  uint32
}

// SIMCONNECT_RECV_RESERVED_KEY carries reserved key information.
type SIMCONNECT_RECV_RESERVED_KEY struct {
	SIMCONNECT_RECV
	SzChoiceReserved [30]byte
	SzReservedKey    [50]byte
}

// SIMCONNECT_RECV_SYSTEM_STATE carries system state data.
type SIMCONNECT_RECV_SYSTEM_STATE struct {
	SIMCONNECT_RECV
	DwRequestID uint32
	DwInteger   uint32
	FFloat      float32
	SzString    [MAX_PATH]byte
}

// SIMCONNECT_RECV_CUSTOM_ACTION carries custom mission-action data.
// SzPayLoad is variable-length; only the first byte is declared.
type SIMCONNECT_RECV_CUSTOM_ACTION struct {
	SIMCONNECT_RECV_EVENT
	GuidInstanceId      GUID
	DwWaitForCompletion uint32
	SzPayLoad           [1]byte // variable-length string
}

// SIMCONNECT_RECV_EVENT_WEATHER_MODE carries the new weather mode in DwData.
type SIMCONNECT_RECV_EVENT_WEATHER_MODE struct {
	SIMCONNECT_RECV_EVENT
}

// SIMCONNECT_RECV_FACILITIES_LIST is the base for facility list responses.
type SIMCONNECT_RECV_FACILITIES_LIST struct {
	SIMCONNECT_RECV
	DwRequestID   uint32
	DwArraySize   uint32
	DwEntryNumber uint32
	DwOutOf       uint32
}

// SIMCONNECT_RECV_LIST_TEMPLATE is the generic list base used by newer APIs.
type SIMCONNECT_RECV_LIST_TEMPLATE struct {
	SIMCONNECT_RECV
	DwRequestID   uint32
	DwArraySize   uint32
	DwEntryNumber uint32
	DwOutOf       uint32
}

// SIMCONNECT_DATA_FACILITY_AIRPORT holds airport facility data.
type SIMCONNECT_DATA_FACILITY_AIRPORT struct {
	Ident     [9]byte
	Region    [3]byte
	Latitude  SIMCONNECT_PACKED_FLOAT64
	Longitude SIMCONNECT_PACKED_FLOAT64
	Altitude  SIMCONNECT_PACKED_FLOAT64
}

// SIMCONNECT_RECV_AIRPORT_LIST carries an array of SIMCONNECT_DATA_FACILITY_AIRPORT.
type SIMCONNECT_RECV_AIRPORT_LIST struct {
	SIMCONNECT_RECV_FACILITIES_LIST
	RgData [1]SIMCONNECT_DATA_FACILITY_AIRPORT // variable-length
}

// SIMCONNECT_DATA_FACILITY_WAYPOINT extends airport data with magnetic variation.
type SIMCONNECT_DATA_FACILITY_WAYPOINT struct {
	SIMCONNECT_DATA_FACILITY_AIRPORT
	FMagVar float32
}

// SIMCONNECT_RECV_WAYPOINT_LIST carries an array of waypoints.
type SIMCONNECT_RECV_WAYPOINT_LIST struct {
	SIMCONNECT_RECV_FACILITIES_LIST
	RgData [1]SIMCONNECT_DATA_FACILITY_WAYPOINT // variable-length
}

// SIMCONNECT_DATA_FACILITY_NDB extends waypoint data with frequency.
type SIMCONNECT_DATA_FACILITY_NDB struct {
	SIMCONNECT_DATA_FACILITY_WAYPOINT
	FFrequency uint32
}

// SIMCONNECT_RECV_NDB_LIST carries an array of NDB facilities.
type SIMCONNECT_RECV_NDB_LIST struct {
	SIMCONNECT_RECV_FACILITIES_LIST
	RgData [1]SIMCONNECT_DATA_FACILITY_NDB // variable-length
}

// SIMCONNECT_DATA_FACILITY_VOR extends NDB data with VOR-specific fields.
type SIMCONNECT_DATA_FACILITY_VOR struct {
	SIMCONNECT_DATA_FACILITY_NDB
	Flags            uint32
	FLocalizer       float32
	GlideLat         SIMCONNECT_PACKED_FLOAT64
	GlideLon         SIMCONNECT_PACKED_FLOAT64
	GlideAlt         SIMCONNECT_PACKED_FLOAT64
	FGlideSlopeAngle float32
}

// SIMCONNECT_RECV_VOR_LIST carries an array of VOR facilities.
type SIMCONNECT_RECV_VOR_LIST struct {
	SIMCONNECT_RECV_FACILITIES_LIST
	RgData [1]SIMCONNECT_DATA_FACILITY_VOR // variable-length
}

// SIMCONNECT_RECV_FACILITY_DATA carries a single facility data element.
// Data is variable-length; only the first uint32 is declared.
type SIMCONNECT_RECV_FACILITY_DATA struct {
	SIMCONNECT_RECV
	UserRequestId         uint32
	UniqueRequestId       uint32
	ParentUniqueRequestId uint32
	Type                  uint32
	IsListItem            uint32
	ItemIndex             uint32
	ListSize              uint32
	Data                  [1]uint32 // variable-length
}

// SIMCONNECT_RECV_FACILITY_DATA_END signals the end of a facility data stream.
type SIMCONNECT_RECV_FACILITY_DATA_END struct {
	SIMCONNECT_RECV
	RequestId uint32
}

// SIMCONNECT_ICAO holds an ICAO identifier.
type SIMCONNECT_ICAO struct {
	Type    byte
	Ident   [9]byte
	Region  [3]byte
	Airport [5]byte
}

// SIMCONNECT_DATA_LATLONALT holds a geographic position.
type SIMCONNECT_DATA_LATLONALT struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

// SIMCONNECT_DATA_PBH holds pitch, bank, heading.
type SIMCONNECT_DATA_PBH struct {
	Pitch   float32
	Bank    float32
	Heading float32
}

// SIMCONNECT_FACILITY_MINIMAL is a minimal facility record.
type SIMCONNECT_FACILITY_MINIMAL struct {
	Icao SIMCONNECT_ICAO
	Lla  SIMCONNECT_PACKED_DATA_LATLONALT
}

// SIMCONNECT_RECV_FACILITY_MINIMAL_LIST carries an array of minimal facility records.
type SIMCONNECT_RECV_FACILITY_MINIMAL_LIST struct {
	SIMCONNECT_RECV_LIST_TEMPLATE
	RgData [1]SIMCONNECT_FACILITY_MINIMAL // variable-length
}

// SIMCONNECT_DATA_INITPOSITION specifies an initial object position.
type SIMCONNECT_DATA_INITPOSITION struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
	Pitch     float64
	Bank      float64
	Heading   float64
	OnGround  uint32
	Airspeed  uint32
}

// SIMCONNECT_DATA_MARKERSTATE holds a marker name and state.
type SIMCONNECT_DATA_MARKERSTATE struct {
	SzMarkerName  [64]byte
	DwMarkerState uint32
}

// SIMCONNECT_DATA_WAYPOINT holds a waypoint definition.
type SIMCONNECT_DATA_WAYPOINT struct {
	Latitude        SIMCONNECT_PACKED_FLOAT64
	Longitude       SIMCONNECT_PACKED_FLOAT64
	Altitude        SIMCONNECT_PACKED_FLOAT64
	Flags           uint32
	KtsSpeed        SIMCONNECT_PACKED_FLOAT64
	PercentThrottle SIMCONNECT_PACKED_FLOAT64
}

// SIMCONNECT_DATA_XYZ holds a 3-D position or vector.
type SIMCONNECT_DATA_XYZ struct {
	X float64
	Y float64
	Z float64
}

// SIMCONNECT_JETWAY_DATA holds jetway state information.
type SIMCONNECT_JETWAY_DATA struct {
	AirportIcao         [8]byte
	ParkingIndex        int32
	Lla                 SIMCONNECT_PACKED_DATA_LATLONALT
	Pbh                 SIMCONNECT_DATA_PBH
	Status              int32
	Door                int32
	ExitDoorRelativePos SIMCONNECT_PACKED_DATA_XYZ
	MainHandlePos       SIMCONNECT_PACKED_DATA_XYZ
	SecondaryHandle     SIMCONNECT_PACKED_DATA_XYZ
	WheelGroundLock     SIMCONNECT_PACKED_DATA_XYZ
	JetwayObjectId      uint32
	AttachedObjectId    uint32
}

// SIMCONNECT_RECV_JETWAY_DATA carries an array of jetway records.
type SIMCONNECT_RECV_JETWAY_DATA struct {
	SIMCONNECT_RECV_LIST_TEMPLATE
	RgData [1]SIMCONNECT_JETWAY_DATA // variable-length
}

// SIMCONNECT_RECV_ACTION_CALLBACK is sent when an action completes.
type SIMCONNECT_RECV_ACTION_CALLBACK struct {
	SIMCONNECT_RECV
	SzActionID  [MAX_PATH]byte
	CbRequestId uint32
}

// SIMCONNECT_INPUT_EVENT_DESCRIPTOR describes a single input event.
type SIMCONNECT_INPUT_EVENT_DESCRIPTOR struct {
	Name  [64]byte
	Hash  SIMCONNECT_PACKED_UINT64
	EType SIMCONNECT_INPUT_EVENT_TYPE
}

// SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS carries an array of input event descriptors.
type SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS struct {
	SIMCONNECT_RECV_LIST_TEMPLATE
	RgData [1]SIMCONNECT_INPUT_EVENT_DESCRIPTOR // variable-length
}

// SIMCONNECT_RECV_GET_INPUT_EVENT carries the current value of an input event.
// Value is variable-length (type depends on EType).
type SIMCONNECT_RECV_GET_INPUT_EVENT struct {
	SIMCONNECT_RECV
	DwRequestID uint32
	EType       SIMCONNECT_INPUT_EVENT_TYPE
	Value       [1]uint32 // variable-length
}

// SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT carries a subscribed input event value.
type SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT struct {
	SIMCONNECT_RECV
	Hash  SIMCONNECT_PACKED_UINT64
	EType SIMCONNECT_INPUT_EVENT_TYPE
	Value [1]uint32 // variable-length
}

// SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS carries an input event parameter.
type SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS struct {
	SIMCONNECT_RECV
	Hash  SIMCONNECT_PACKED_UINT64
	Value [MAX_PATH]byte
}

// SIMCONNECT_DATA_CAMERA holds camera position, orientation and field-of-view.
type SIMCONNECT_DATA_CAMERA struct {
	Position                    SIMCONNECT_PACKED_DATA_XYZ
	PositionReferential         SIMCONNECT_POSITION_REFERENTIAL
	PositionReferentialObjectId uint32
	TargetedPos                 SIMCONNECT_PACKED_DATA_XYZ
	Pbh                         SIMCONNECT_DATA_PBH
	RotationReferential         SIMCONNECT_POSITION_REFERENTIAL
	RotationReferentialObjectId uint32
	Fov                         SIMCONNECT_PACKED_FLOAT64
}

// SIMCONNECT_RECV_CAMERA_DATA carries camera data in response to CameraGet.
type SIMCONNECT_RECV_CAMERA_DATA struct {
	SIMCONNECT_RECV
	CameraData SIMCONNECT_DATA_CAMERA
}

// SIMCONNECT_RECV_CAMERA_WORLD_LOCKER carries the world-locker status.
type SIMCONNECT_RECV_CAMERA_WORLD_LOCKER struct {
	SIMCONNECT_RECV
	Status SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS
}

// SIMCONNECT_RECV_CAMERA_STATUS carries the camera acquisition state.
type SIMCONNECT_RECV_CAMERA_STATUS struct {
	SIMCONNECT_RECV
	AcquiredState   uint32 // SIMCONNECT_CAMERA_AVAILABILITY
	BGameControlled int32  // BOOL
}

// SIMCONNECT_VERSION_BASE_TYPE holds a four-part version number.
type SIMCONNECT_VERSION_BASE_TYPE struct {
	Major    uint16
	Minor    uint16
	Revision uint16
	Build    uint16
}

// SIMCONNECT_CONTROLLER_ITEM describes a connected controller.
type SIMCONNECT_CONTROLLER_ITEM struct {
	DeviceName      [256]byte
	DeviceId        uint32
	ProductId       uint32
	CompositeID     uint32
	HardwareVersion SIMCONNECT_VERSION_BASE_TYPE
}

// SIMCONNECT_RECV_CONTROLLERS_LIST carries an array of controller items.
type SIMCONNECT_RECV_CONTROLLERS_LIST struct {
	SIMCONNECT_RECV_LIST_TEMPLATE
	RgData [1]SIMCONNECT_CONTROLLER_ITEM // variable-length
}

// SIMCONNECT_ENUMERATE_SIMOBJECT_LIVERY describes an aircraft livery.
type SIMCONNECT_ENUMERATE_SIMOBJECT_LIVERY struct {
	AircraftTitle [256]byte
	LiveryName    [256]byte
}

// SIMCONNECT_RECV_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST carries livery records.
type SIMCONNECT_RECV_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST struct {
	SIMCONNECT_RECV_LIST_TEMPLATE
	RgData [1]SIMCONNECT_ENUMERATE_SIMOBJECT_LIVERY // variable-length
}

// SIMCONNECT_RECV_FLOW_EVENT carries a flight flow event with optional FLT path.
type SIMCONNECT_RECV_FLOW_EVENT struct {
	SIMCONNECT_RECV
	FlowEvent SIMCONNECT_FLOW_EVENT
	FltPath   [256]byte
}

// SIMCONNECT_CAMERA_DEFINITION_ITEM is a single camera definition name.
type SIMCONNECT_CAMERA_DEFINITION_ITEM struct {
	Str [256]byte
}

// SIMCONNECT_RECV_CAMERA_DEFINITION_LIST carries camera definition names.
type SIMCONNECT_RECV_CAMERA_DEFINITION_LIST struct {
	SIMCONNECT_RECV_LIST_TEMPLATE
	RgData [1]SIMCONNECT_CAMERA_DEFINITION_ITEM // variable-length
}

// SIMCONNECT_RECV_COMM_BUS carries a comm-bus event with variable-length payload.
type SIMCONNECT_RECV_COMM_BUS struct {
	SIMCONNECT_RECV_LIST_TEMPLATE
	UEventID uint32
	RgData   [1]byte // variable-length string payload
}
