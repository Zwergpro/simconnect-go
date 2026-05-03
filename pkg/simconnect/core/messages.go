//go:build windows

package core

// Message is the common interface for all decoded SimConnect packets.
type Message interface {
	RecvID() RecvID
}

// ─── Session-level messages ──────────────────────────────────────────────────

type OpenMessage struct {
	ApplicationName         string
	ApplicationVersionMajor uint32
	ApplicationVersionMinor uint32
	ApplicationBuildMajor   uint32
	ApplicationBuildMinor   uint32
	SimConnectVersionMajor  uint32
	SimConnectVersionMinor  uint32
	SimConnectBuildMajor    uint32
	SimConnectBuildMinor    uint32
}

func (OpenMessage) RecvID() RecvID { return RecvIDOpen }

type QuitMessage struct{}

func (QuitMessage) RecvID() RecvID { return RecvIDQuit }

// ─── Event messages ──────────────────────────────────────────────────────────

type ClientEvent struct {
	GroupID    uint32
	EventID    uint32
	Data       uint32
	DataValues [5]uint32
	FileName   string
	Flags      uint32
	ObjectType SimObjectType
	FrameRate  float32
	SimSpeed   float32
}

func (ClientEvent) RecvID() RecvID { return RecvIDEvent }

type FilenameEvent struct {
	ClientEvent
	FileName string
	Flags    uint32
}

func (FilenameEvent) RecvID() RecvID { return RecvIDEventFilename }

type ObjectAddRemoveEvent struct {
	ClientEvent
	ObjectType SimObjectType
}

func (ObjectAddRemoveEvent) RecvID() RecvID { return RecvIDEventObjectAddRemove }

type FrameEvent struct {
	ClientEvent
	FrameRate float32
	SimSpeed  float32
}

func (FrameEvent) RecvID() RecvID { return RecvIDEventFrame }

type ClientEventEX1 struct {
	ClientEvent
}

func (ClientEventEX1) RecvID() RecvID { return RecvIDEventEX1 }

type ReservedKeyMessage struct {
	ChoiceReserved string
	ReservedKey    string
}

func (ReservedKeyMessage) RecvID() RecvID { return RecvIDReservedKey }

// ─── Request-correlated messages ─────────────────────────────────────────────
//
// DispatchRequestID returns the request ID used by the client dispatch loop
// to route a response to the correct waiter channel.  Separate from the field
// name to avoid a Go "method conflicts with field" error on types that expose
// RequestID as a plain uint32 field.

type SystemStateMessage struct {
	RequestID uint32
	Integer   uint32
	Float     float32
	String    string
}

func (SystemStateMessage) RecvID() RecvID                { return RecvIDSystemState }
func (m SystemStateMessage) DispatchRequestID() uint32   { return m.RequestID }

type AssignedObjectIDMessage struct {
	RequestID uint32
	ObjectID  ObjectID
}

func (AssignedObjectIDMessage) RecvID() RecvID                 { return RecvIDAssignedObjectID }
func (m AssignedObjectIDMessage) DispatchRequestID() uint32    { return m.RequestID }

type ActionCallbackMessage struct {
	RequestID uint32
	ActionID  string
}

func (ActionCallbackMessage) RecvID() RecvID                { return RecvIDActionCallback }
func (m ActionCallbackMessage) DispatchRequestID() uint32   { return m.RequestID }

// ─── SimObject data messages ──────────────────────────────────────────────────

type SimObjectDataMessage struct {
	RequestID   uint32
	ObjectID    uint32
	DefineID    uint32
	Flags       uint32
	EntryNumber uint32
	OutOf       uint32
	DefineCount uint32
	Payload     []byte
}

func (SimObjectDataMessage) RecvID() RecvID                { return RecvIDSimObjectData }
func (m SimObjectDataMessage) DispatchRequestID() uint32   { return m.RequestID }

type SimObjectDataByTypeMessage struct {
	SimObjectDataMessage
}

func (SimObjectDataByTypeMessage) RecvID() RecvID                { return RecvIDSimObjectDataByType }
func (m SimObjectDataByTypeMessage) DispatchRequestID() uint32   { return m.RequestID }

type ClientDataMessage struct {
	SimObjectDataMessage
}

func (ClientDataMessage) RecvID() RecvID                { return RecvIDClientData }
func (m ClientDataMessage) DispatchRequestID() uint32   { return m.RequestID }

// ─── Facility messages ────────────────────────────────────────────────────────

type FacilityListMeta struct {
	RequestID   uint32
	ArraySize   uint32
	EntryNumber uint32
	OutOf       uint32
}

type AirportListMessage struct {
	FacilityListMeta
	Airports []Airport
}

func (AirportListMessage) RecvID() RecvID                { return RecvIDAirportList }
func (m AirportListMessage) DispatchRequestID() uint32   { return m.RequestID }

type WaypointListMessage struct {
	FacilityListMeta
	Waypoints []Waypoint
}

func (WaypointListMessage) RecvID() RecvID                 { return RecvIDWaypointList }
func (m WaypointListMessage) DispatchRequestID() uint32    { return m.RequestID }

type NDBListMessage struct {
	FacilityListMeta
	NDBs []NDB
}

func (NDBListMessage) RecvID() RecvID                { return RecvIDNDBList }
func (m NDBListMessage) DispatchRequestID() uint32   { return m.RequestID }

type VORListMessage struct {
	FacilityListMeta
	VORs []VOR
}

func (VORListMessage) RecvID() RecvID                { return RecvIDVORList }
func (m VORListMessage) DispatchRequestID() uint32   { return m.RequestID }

type FacilityDataMessage struct {
	UserRequestID         uint32
	UniqueRequestID       uint32
	ParentUniqueRequestID uint32
	Type                  FacilityDataType
	IsListItem            bool
	ItemIndex             uint32
	ListSize              uint32
	Payload               []byte
}

func (FacilityDataMessage) RecvID() RecvID                { return RecvIDFacilityData }
func (m FacilityDataMessage) DispatchRequestID() uint32   { return m.UserRequestID }

type FacilityDataEndMessage struct {
	RequestID uint32
}

func (FacilityDataEndMessage) RecvID() RecvID                { return RecvIDFacilityDataEnd }
func (m FacilityDataEndMessage) DispatchRequestID() uint32   { return m.RequestID }

type FacilityMinimalListMessage struct {
	FacilityListMeta
	Facilities []FacilityMinimal
}

func (FacilityMinimalListMessage) RecvID() RecvID                 { return RecvIDFacilityMinimalList }
func (m FacilityMinimalListMessage) DispatchRequestID() uint32    { return m.RequestID }

type JetwayDataMessage struct {
	FacilityListMeta
	Jetways []JetwayData
}

func (JetwayDataMessage) RecvID() RecvID                { return RecvIDJetwayData }
func (m JetwayDataMessage) DispatchRequestID() uint32   { return m.RequestID }

// ─── Input event messages ────────────────────────────────────────────────────

type InputEventDescriptor struct {
	Name string
	Hash uint64
	Type InputEventType
}

type InputEventListMessage struct {
	FacilityListMeta
	Events []InputEventDescriptor
}

func (InputEventListMessage) RecvID() RecvID                { return RecvIDEnumerateInputEvents }
func (m InputEventListMessage) DispatchRequestID() uint32   { return m.RequestID }

type InputEventValueMessage struct {
	RequestID uint32
	Type      InputEventType
	Double    float64
	String    string
	Payload   []byte
}

func (InputEventValueMessage) RecvID() RecvID                { return RecvIDGetInputEvent }
func (m InputEventValueMessage) DispatchRequestID() uint32   { return m.RequestID }

type InputEventSubscriptionMessage struct {
	Hash    uint64
	Type    InputEventType
	Double  float64
	String  string
	Payload []byte
}

func (InputEventSubscriptionMessage) RecvID() RecvID { return RecvIDSubscribeInputEvent }

type InputEventParamsMessage struct {
	Hash   uint64
	Value  string
	Params []string
}

func (InputEventParamsMessage) RecvID() RecvID { return RecvIDEnumerateInputEventParams }

// ─── Controller / livery messages ─────────────────────────────────────────────

type ControllerItem struct {
	DeviceName      string
	DeviceID        uint32
	ProductID       uint32
	CompositeID     uint32
	HardwareVersion Version
}

type ControllersListMessage struct {
	FacilityListMeta
	Controllers []ControllerItem
}

func (ControllersListMessage) RecvID() RecvID { return RecvIDControllersList }

type SimObjectLivery struct {
	AircraftTitle string
	LiveryName    string
}

type SimObjectLiveryListMessage struct {
	FacilityListMeta
	Liveries []SimObjectLivery
}

func (SimObjectLiveryListMessage) RecvID() RecvID                 { return RecvIDEnumerateSimObjectAndLiveryList }
func (m SimObjectLiveryListMessage) DispatchRequestID() uint32    { return m.RequestID }

// ─── Misc messages ────────────────────────────────────────────────────────────

type FlowEventMessage struct {
	Event   FlowEvent
	FLTPath string
}

func (FlowEventMessage) RecvID() RecvID { return RecvIDFlowEvent }

type CommBusMessage struct {
	FacilityListMeta
	EventID uint32
	Data    string
	Payload []byte
}

func (CommBusMessage) RecvID() RecvID { return RecvIDCommBus }

// ─── Camera messages ─────────────────────────────────────────────────────────

type CameraData struct {
	Position                    XYZ
	PositionReferential         PositionReferential
	PositionReferentialObjectID ObjectID
	TargetedPos                 XYZ
	PBH                         PBH
	RotationReferential         PositionReferential
	RotationReferentialObjectID ObjectID
	FOV                         float64
}

type CameraDataMessage struct {
	CameraData CameraData
}

func (CameraDataMessage) RecvID() RecvID { return RecvIDCameraData }

type CameraStatusMessage struct {
	AcquiredState  CameraAvailability
	GameControlled bool
}

func (CameraStatusMessage) RecvID() RecvID { return RecvIDCameraStatus }

type CameraDefinitionListMessage struct {
	FacilityListMeta
	Definitions []string
}

func (CameraDefinitionListMessage) RecvID() RecvID { return RecvIDCameraDefinitionList }

type CameraWorldLockerMessage struct {
	Status CameraWorldLockerStatus
}

func (CameraWorldLockerMessage) RecvID() RecvID { return RecvIDCameraWorldLocker }
