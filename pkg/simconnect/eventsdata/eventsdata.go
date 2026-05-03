//go:build windows

// Package eventsdata implements the Events_And_Data SimConnect API category.
// It covers custom/system events, sim-object data definitions and subscriptions,
// client data areas, and flow events.
package eventsdata

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// Session is the subset of client.Sim methods used by this package.
type Session interface {
	NextEventID() uint32
	NextRequestID() uint32
	NextDefinitionID() uint32
	AddWaiter(uint32) (<-chan core.RequestResult, error)
	RemoveWaiter(uint32)
	AddDataSub(uint32, func(core.Message)) error
	RemoveDataSub(uint32)
	TrackSend(uint32)
	ReportError(error)
	Bindings() *bindings.SimConnect
	RegisterHandler(core.RecvID, func(core.Message))
	RegisterCloseHook(func())
	ChannelBuffer() int
	Context() context.Context
}

// ─── Event types ────────────────────────────────────────────────────────────

// Event is a handle to a mapped SimConnect client event.
type Event struct {
	id core.EventID
}

func (e Event) ID() core.EventID { return e.id }

type TransmitOption func(*transmitConfig)

type transmitConfig struct {
	objectID core.ObjectID
	groupID  core.NotificationGroupID
	flags    core.EventFlag
}

func defaultTransmitConfig() transmitConfig {
	return transmitConfig{
		objectID: core.UserAircraft,
		groupID:  core.NotificationGroupID(core.GroupPriorityDefault),
		flags:    core.EventFlagDefault,
	}
}

func WithTransmitObject(objectID core.ObjectID) TransmitOption {
	return func(cfg *transmitConfig) { cfg.objectID = objectID }
}

func WithTransmitGroup(groupID core.NotificationGroupID) TransmitOption {
	return func(cfg *transmitConfig) { cfg.groupID = groupID }
}

func WithTransmitFlags(flags core.EventFlag) TransmitOption {
	return func(cfg *transmitConfig) { cfg.flags = flags }
}

// ─── Data definition types ───────────────────────────────────────────────────

type Definition[T any] struct {
	core *definitionCore
}

type definitionCore struct {
	fields []definitionField
	typ    reflect.Type
}

type definitionField struct {
	Field
	index int
	size  int
}

type Field struct {
	Name    string
	Units   string
	Type    core.DataType
	Epsilon float32
}

type DataUpdate[T any] struct {
	ObjectID core.ObjectID
	Value    T
}

type DataOption func(*dataConfig)

type dataConfig struct {
	flags    core.DataRequestFlag
	origin   uint32
	interval uint32
	limit    uint32
}

func ChangedOnly() DataOption {
	return func(cfg *dataConfig) {
		cfg.flags |= core.DataRequestChanged
	}
}

func WithDataRequestTiming(origin, interval, limit uint32) DataOption {
	return func(cfg *dataConfig) {
		cfg.origin = origin
		cfg.interval = interval
		cfg.limit = limit
	}
}

// ─── ClientData types ────────────────────────────────────────────────────────

type ClientDataDefinition struct {
	id core.ClientDataDefinitionID
}

func (d ClientDataDefinition) ID() core.ClientDataDefinitionID { return d.id }

type ClientDataDefinitionItem struct {
	Offset     uint32
	SizeOrType uint32
	Epsilon    float32
	DatumID    uint32
}

// ─── Client ──────────────────────────────────────────────────────────────────

// EventsData exposes events, data definitions/requests, client data areas, and flow events.
type EventsData struct {
	c Session

	mu          sync.Mutex
	definitions map[*definitionCore]uint32
	flowSubs    []chan core.FlowEventMessage
}

// New creates an EventsData and registers flow-event handler.
func New(c Session) *EventsData {
	s := &EventsData{
		c:           c,
		definitions: make(map[*definitionCore]uint32),
	}
	c.RegisterHandler(core.RecvIDFlowEvent, func(msg core.Message) {
		m, ok := msg.(core.FlowEventMessage)
		if !ok {
			return
		}
		s.mu.Lock()
		subs := append([]chan core.FlowEventMessage(nil), s.flowSubs...)
		s.mu.Unlock()
		for _, ch := range subs {
			s.sendFlowEvent(ch, m)
		}
	})
	c.RegisterCloseHook(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		for _, ch := range s.flowSubs {
			close(ch)
		}
		s.flowSubs = nil
	})
	return s
}

// ─── Custom / notification-group events ──────────────────────────────────────

func (s *EventsData) MapEvent(name string) (Event, error) {
	eventID := core.EventID(s.c.NextEventID())
	if err := s.c.Bindings().MapClientEventToSimEvent(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID), name); err != nil {
		return Event{}, err
	}
	return Event{id: eventID}, nil
}

func (s *EventsData) Transmit(event Event, data uint32, opts ...TransmitOption) error {
	cfg := defaultTransmitConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	return s.c.Bindings().TransmitClientEvent(
		bindings.SIMCONNECT_OBJECT_ID(cfg.objectID),
		bindings.SIMCONNECT_CLIENT_EVENT_ID(event.id),
		data,
		bindings.SIMCONNECT_NOTIFICATION_GROUP_ID(cfg.groupID),
		bindings.SIMCONNECT_EVENT_FLAG(cfg.flags),
	)
}

func (s *EventsData) TransmitEX1(event Event, data0, data1, data2, data3, data4 uint32, opts ...TransmitOption) error {
	cfg := defaultTransmitConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	return s.c.Bindings().TransmitClientEvent_EX1(
		bindings.SIMCONNECT_OBJECT_ID(cfg.objectID),
		bindings.SIMCONNECT_CLIENT_EVENT_ID(event.id),
		bindings.SIMCONNECT_NOTIFICATION_GROUP_ID(cfg.groupID),
		bindings.SIMCONNECT_EVENT_FLAG(cfg.flags),
		data0, data1, data2, data3, data4,
	)
}

func (s *EventsData) AddClientEventToNotificationGroup(groupID core.NotificationGroupID, event Event, maskable bool) error {
	return s.c.Bindings().AddClientEventToNotificationGroup(bindings.SIMCONNECT_NOTIFICATION_GROUP_ID(groupID), bindings.SIMCONNECT_CLIENT_EVENT_ID(event.id), maskable)
}

func (s *EventsData) RemoveClientEvent(groupID core.NotificationGroupID, event Event) error {
	return s.c.Bindings().RemoveClientEvent(bindings.SIMCONNECT_NOTIFICATION_GROUP_ID(groupID), bindings.SIMCONNECT_CLIENT_EVENT_ID(event.id))
}

func (s *EventsData) ClearNotificationGroup(groupID core.NotificationGroupID) error {
	return s.c.Bindings().ClearNotificationGroup(bindings.SIMCONNECT_NOTIFICATION_GROUP_ID(groupID))
}

func (s *EventsData) RequestNotificationGroup(groupID core.NotificationGroupID) error {
	return s.c.Bindings().RequestNotificationGroup(bindings.SIMCONNECT_NOTIFICATION_GROUP_ID(groupID), 0, 0)
}

func (s *EventsData) RequestReservedKey(event Event, keyChoice1, keyChoice2, keyChoice3 string) error {
	return s.c.Bindings().RequestReservedKey(bindings.SIMCONNECT_CLIENT_EVENT_ID(event.id), keyChoice1, keyChoice2, keyChoice3)
}

// ─── Flow events ──────────────────────────────────────────────────────────────

func (s *EventsData) SubscribeFlowEvents(ctx context.Context) (<-chan core.FlowEventMessage, error) {
	if err := s.c.Bindings().SubscribeToFlowEvent(); err != nil {
		return nil, err
	}
	ch := make(chan core.FlowEventMessage, s.c.ChannelBuffer())
	s.mu.Lock()
	s.flowSubs = append(s.flowSubs, ch)
	s.mu.Unlock()

	go func() {
		<-ctx.Done()
		s.removeFlowSub(ch)
	}()
	return ch, nil
}

func (s *EventsData) UnsubscribeFlowEvents() error {
	return s.c.Bindings().UnsubscribeToFlowEvent()
}

func (s *EventsData) removeFlowSub(target chan core.FlowEventMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, ch := range s.flowSubs {
		if ch == target {
			close(ch)
			s.flowSubs = append(s.flowSubs[:i], s.flowSubs[i+1:]...)
			if len(s.flowSubs) == 0 {
				_ = s.c.Bindings().UnsubscribeToFlowEvent()
			}
			return
		}
	}
}

func (s *EventsData) sendFlowEvent(ch chan core.FlowEventMessage, msg core.FlowEventMessage) {
	defer func() { _ = recover() }()
	select {
	case ch <- msg:
	case <-s.c.Context().Done():
	}
}

// ─── Data definitions ─────────────────────────────────────────────────────────

func Define[T any]() (*Definition[T], error) {
	var zero T
	dc, err := buildDefinition(reflect.TypeOf(zero), nil)
	if err != nil {
		return nil, err
	}
	return &Definition[T]{core: dc}, nil
}

func DefineFields[T any](fields ...Field) (*Definition[T], error) {
	var zero T
	dc, err := buildDefinition(reflect.TypeOf(zero), fields)
	if err != nil {
		return nil, err
	}
	return &Definition[T]{core: dc}, nil
}

func RequestDataOnce[T any](ctx context.Context, s *EventsData, def *Definition[T], object core.ObjectID) (T, error) {
	var zero T
	if def == nil {
		return zero, fmt.Errorf("%w: nil definition", core.ErrDecode)
	}
	defineID, err := s.ensureDefinition(def.core)
	if err != nil {
		return zero, err
	}
	requestID := s.c.NextRequestID()
	waiter, err := s.c.AddWaiter(requestID)
	if err != nil {
		return zero, err
	}
	err = s.c.Bindings().RequestDataOnSimObject(
		bindings.SIMCONNECT_DATA_REQUEST_ID(requestID),
		bindings.SIMCONNECT_DATA_DEFINITION_ID(defineID),
		bindings.SIMCONNECT_OBJECT_ID(object),
		bindings.SIMCONNECT_PERIOD(core.PeriodOnce),
		bindings.SIMCONNECT_DATA_REQUEST_FLAG(core.DataRequestDefault),
		0, 0, 0,
	)
	if err != nil {
		s.c.RemoveWaiter(requestID)
		return zero, err
	}
	s.c.TrackSend(requestID)

	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.c.RemoveWaiter(requestID)
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	data, ok := result.Msg.(core.SimObjectDataMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected simobject data, got %T", core.ErrDecode, result.Msg)
	}
	return decodeData[T](def.core, data.Payload)
}

func RequestDataByTypeOnce[T any](ctx context.Context, s *EventsData, def *Definition[T], radiusMeters uint32, objType core.SimObjectType) (DataUpdate[T], error) {
	var zero DataUpdate[T]
	if def == nil {
		return zero, fmt.Errorf("%w: nil definition", core.ErrDecode)
	}
	defineID, err := s.ensureDefinition(def.core)
	if err != nil {
		return zero, err
	}
	requestID := s.c.NextRequestID()
	waiter, err := s.c.AddWaiter(requestID)
	if err != nil {
		return zero, err
	}
	err = s.c.Bindings().RequestDataOnSimObjectType(
		bindings.SIMCONNECT_DATA_REQUEST_ID(requestID),
		bindings.SIMCONNECT_DATA_DEFINITION_ID(defineID),
		radiusMeters,
		bindings.SIMCONNECT_SIMOBJECT_TYPE(objType),
	)
	if err != nil {
		s.c.RemoveWaiter(requestID)
		return zero, err
	}
	s.c.TrackSend(requestID)

	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.c.RemoveWaiter(requestID)
		return zero, ctx.Err()
	}
	if result.Err != nil {
		return zero, result.Err
	}
	data, ok := result.Msg.(core.SimObjectDataByTypeMessage)
	if !ok {
		return zero, fmt.Errorf("%w: expected simobject data by type, got %T", core.ErrDecode, result.Msg)
	}
	value, err := decodeData[T](def.core, data.Payload)
	if err != nil {
		return zero, err
	}
	return DataUpdate[T]{ObjectID: core.ObjectID(data.ObjectID), Value: value}, nil
}

func SetDataOnSimObject[T any](s *EventsData, def *Definition[T], object core.ObjectID, value T, flags core.DataSetFlag) error {
	if def == nil {
		return fmt.Errorf("%w: nil definition", core.ErrDecode)
	}
	defineID, err := s.ensureDefinition(def.core)
	if err != nil {
		return err
	}
	payload, err := encodeData(def.core, value)
	if err != nil {
		return err
	}
	var pData *byte
	if len(payload) > 0 {
		pData = &payload[0]
	}
	return s.c.Bindings().SetDataOnSimObject(
		bindings.SIMCONNECT_DATA_DEFINITION_ID(defineID),
		bindings.SIMCONNECT_OBJECT_ID(object),
		bindings.SIMCONNECT_DATA_SET_FLAG(flags),
		0,
		uint32(len(payload)),
		unsafe.Pointer(pData),
	)
}

func SubscribeData[T any](ctx context.Context, s *EventsData, def *Definition[T], object core.ObjectID, period core.Period, opts ...DataOption) (<-chan DataUpdate[T], error) {
	if def == nil {
		return nil, fmt.Errorf("%w: nil definition", core.ErrDecode)
	}
	defineID, err := s.ensureDefinition(def.core)
	if err != nil {
		return nil, err
	}
	cfg := dataConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}
	requestID := s.c.NextRequestID()
	ch := make(chan DataUpdate[T], s.c.ChannelBuffer())
	handler := func(msg core.Message) {
		data, ok := msg.(core.SimObjectDataMessage)
		if !ok {
			return
		}
		value, err := decodeData[T](def.core, data.Payload)
		if err != nil {
			s.c.ReportError(err)
			return
		}
		sendDataUpdate(ctx, s.c.Context(), ch, DataUpdate[T]{ObjectID: core.ObjectID(data.ObjectID), Value: value})
	}
	if err := s.c.AddDataSub(requestID, handler); err != nil {
		close(ch)
		return nil, err
	}
	err = s.c.Bindings().RequestDataOnSimObject(
		bindings.SIMCONNECT_DATA_REQUEST_ID(requestID),
		bindings.SIMCONNECT_DATA_DEFINITION_ID(defineID),
		bindings.SIMCONNECT_OBJECT_ID(object),
		bindings.SIMCONNECT_PERIOD(period),
		bindings.SIMCONNECT_DATA_REQUEST_FLAG(cfg.flags),
		cfg.origin,
		cfg.interval,
		cfg.limit,
	)
	if err != nil {
		s.c.RemoveDataSub(requestID)
		return nil, err
	}
	s.c.TrackSend(requestID)

	go func() {
		<-ctx.Done()
		_ = s.c.Bindings().RequestDataOnSimObject(
			bindings.SIMCONNECT_DATA_REQUEST_ID(requestID),
			bindings.SIMCONNECT_DATA_DEFINITION_ID(defineID),
			bindings.SIMCONNECT_OBJECT_ID(object),
			bindings.SIMCONNECT_PERIOD(core.PeriodNever),
			bindings.SIMCONNECT_DATA_REQUEST_FLAG(core.DataRequestDefault),
			0, 0, 0,
		)
		s.c.RemoveDataSub(requestID)
		close(ch)
	}()
	return ch, nil
}

// ─── Client data areas ────────────────────────────────────────────────────────

func (s *EventsData) MapClientDataNameToID(name string, id core.ClientDataID) error {
	return s.c.Bindings().MapClientDataNameToID(name, bindings.SIMCONNECT_CLIENT_DATA_ID(id))
}

func (s *EventsData) CreateClientData(id core.ClientDataID, size uint32, flags core.ClientDataCreateFlag) error {
	return s.c.Bindings().CreateClientData(bindings.SIMCONNECT_CLIENT_DATA_ID(id), size, bindings.SIMCONNECT_CREATE_CLIENT_DATA_FLAG(flags))
}

func (s *EventsData) NewClientDataDefinition(items ...ClientDataDefinitionItem) (ClientDataDefinition, error) {
	defineID := core.ClientDataDefinitionID(s.c.NextDefinitionID())
	for _, item := range items {
		if err := s.c.Bindings().AddToClientDataDefinition(bindings.SIMCONNECT_CLIENT_DATA_DEFINITION_ID(defineID), item.Offset, item.SizeOrType, item.Epsilon, item.DatumID); err != nil {
			return ClientDataDefinition{}, err
		}
	}
	return ClientDataDefinition{id: defineID}, nil
}

func (s *EventsData) AddToClientDataDefinition(def ClientDataDefinition, item ClientDataDefinitionItem) error {
	return s.c.Bindings().AddToClientDataDefinition(bindings.SIMCONNECT_CLIENT_DATA_DEFINITION_ID(def.id), item.Offset, item.SizeOrType, item.Epsilon, item.DatumID)
}

func (s *EventsData) ClearClientDataDefinition(def ClientDataDefinition) error {
	return s.c.Bindings().ClearClientDataDefinition(bindings.SIMCONNECT_CLIENT_DATA_DEFINITION_ID(def.id))
}

func (s *EventsData) RequestClientData(ctx context.Context, clientDataID core.ClientDataID, def ClientDataDefinition, period core.ClientDataPeriod, flags core.ClientDataRequestFlag, origin, interval, limit uint32) (core.ClientDataMessage, error) {
	requestID := s.c.NextRequestID()
	waiter, err := s.c.AddWaiter(requestID)
	if err != nil {
		return core.ClientDataMessage{}, err
	}
	err = s.c.Bindings().RequestClientData(
		bindings.SIMCONNECT_CLIENT_DATA_ID(clientDataID),
		bindings.SIMCONNECT_DATA_REQUEST_ID(requestID),
		bindings.SIMCONNECT_CLIENT_DATA_DEFINITION_ID(def.id),
		bindings.SIMCONNECT_CLIENT_DATA_PERIOD(period),
		bindings.SIMCONNECT_CLIENT_DATA_REQUEST_FLAG(flags),
		origin, interval, limit,
	)
	if err != nil {
		s.c.RemoveWaiter(requestID)
		return core.ClientDataMessage{}, err
	}
	s.c.TrackSend(requestID)

	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.c.RemoveWaiter(requestID)
		return core.ClientDataMessage{}, ctx.Err()
	}
	if result.Err != nil {
		return core.ClientDataMessage{}, result.Err
	}
	data, ok := result.Msg.(core.ClientDataMessage)
	if !ok {
		return core.ClientDataMessage{}, fmt.Errorf("%w: expected client data, got %T", core.ErrDecode, result.Msg)
	}
	return data, nil
}

func (s *EventsData) SetClientData(clientDataID core.ClientDataID, def ClientDataDefinition, flags core.ClientDataSetFlag, data []byte) error {
	var pData *byte
	if len(data) > 0 {
		pData = &data[0]
	}
	return s.c.Bindings().SetClientData(
		bindings.SIMCONNECT_CLIENT_DATA_ID(clientDataID),
		bindings.SIMCONNECT_CLIENT_DATA_DEFINITION_ID(def.id),
		bindings.SIMCONNECT_CLIENT_DATA_SET_FLAG(flags),
		0,
		uint32(len(data)),
		unsafe.Pointer(pData),
	)
}

// ─── Internal: definition management ─────────────────────────────────────────

func (s *EventsData) ensureDefinition(dc *definitionCore) (uint32, error) {
	if dc == nil {
		return 0, fmt.Errorf("%w: nil definition", core.ErrDecode)
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if id, ok := s.definitions[dc]; ok {
		return id, nil
	}
	defineID := s.c.NextDefinitionID()
	for _, field := range dc.fields {
		if err := s.c.Bindings().AddToDataDefinition(
			bindings.SIMCONNECT_DATA_DEFINITION_ID(defineID),
			field.Name, field.Units, bindings.SIMCONNECT_DATATYPE(field.Type), field.Epsilon,
			core.Unused,
		); err != nil {
			return 0, err
		}
	}
	s.definitions[dc] = defineID
	return defineID, nil
}

// ─── Internal: definition building ───────────────────────────────────────────

func buildDefinition(typ reflect.Type, fields []Field) (*definitionCore, error) {
	if typ == nil {
		return nil, fmt.Errorf("%w: nil type", core.ErrUnsupportedType)
	}
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if len(fields) > 0 {
		return buildExplicitDefinition(typ, fields)
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%w: Define requires a struct with sim tags", core.ErrUnsupportedType)
	}
	def := &definitionCore{typ: typ}
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		tag := sf.Tag.Get("sim")
		if tag == "" || tag == "-" {
			continue
		}
		field, err := parseSimTag(tag)
		if err != nil {
			return nil, err
		}
		defField, err := completeField(field, sf.Type, i)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", sf.Name, err)
		}
		def.fields = append(def.fields, defField)
	}
	if len(def.fields) == 0 {
		return nil, fmt.Errorf("%w: no sim fields found", core.ErrUnsupportedType)
	}
	return def, nil
}

func buildExplicitDefinition(typ reflect.Type, fields []Field) (*definitionCore, error) {
	def := &definitionCore{typ: typ}
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		if len(fields) > typ.NumField() {
			return nil, fmt.Errorf("%w: more fields than struct members", core.ErrDecode)
		}
		for i, field := range fields {
			defField, err := completeField(field, typ.Field(i).Type, i)
			if err != nil {
				return nil, err
			}
			def.fields = append(def.fields, defField)
		}
		return def, nil
	}
	if len(fields) != 1 {
		return nil, fmt.Errorf("%w: scalar definitions require exactly one field", core.ErrUnsupportedType)
	}
	defField, err := completeField(fields[0], typ, -1)
	if err != nil {
		return nil, err
	}
	def.fields = append(def.fields, defField)
	return def, nil
}

func parseSimTag(tag string) (Field, error) {
	parts := strings.Split(tag, ",")
	if len(parts) < 2 || strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
		return Field{}, fmt.Errorf("%w: sim tag must be \"name,units\"", core.ErrDecode)
	}
	return Field{Name: strings.TrimSpace(parts[0]), Units: strings.TrimSpace(parts[1])}, nil
}

func completeField(field Field, typ reflect.Type, index int) (definitionField, error) {
	if typ.Kind() == reflect.Pointer {
		return definitionField{}, fmt.Errorf("%w: pointer fields are not supported", core.ErrUnsupportedType)
	}
	if field.Name == "" {
		return definitionField{}, fmt.Errorf("%w: field name is required", core.ErrDecode)
	}
	if field.Type == core.DataTypeInvalid {
		field.Type = datatypeForKind(typ.Kind())
	}
	if !datatypeMatchesKind(field.Type, typ.Kind()) {
		return definitionField{}, fmt.Errorf("%w: %s cannot receive %d", core.ErrUnsupportedType, typ.Kind(), field.Type)
	}
	size := datatypeSize(field.Type)
	if size == 0 {
		return definitionField{}, fmt.Errorf("%w: %s", core.ErrUnsupportedType, typ.Kind())
	}
	return definitionField{Field: field, index: index, size: size}, nil
}

func datatypeMatchesKind(datatype core.DataType, kind reflect.Kind) bool {
	switch datatype {
	case core.DataTypeInt32:
		return kind == reflect.Int32 || kind == reflect.Bool
	case core.DataTypeInt64:
		return kind == reflect.Int64
	case core.DataTypeFloat32:
		return kind == reflect.Float32
	case core.DataTypeFloat64:
		return kind == reflect.Float64
	case core.DataTypeString8,
		core.DataTypeString32,
		core.DataTypeString64,
		core.DataTypeString128,
		core.DataTypeString256,
		core.DataTypeString260:
		return kind == reflect.String
	default:
		return false
	}
}

func datatypeForKind(kind reflect.Kind) core.DataType {
	switch kind {
	case reflect.Bool:
		return core.DataTypeInt32
	case reflect.Int32:
		return core.DataTypeInt32
	case reflect.Int64:
		return core.DataTypeInt64
	case reflect.Float32:
		return core.DataTypeFloat32
	case reflect.Float64:
		return core.DataTypeFloat64
	case reflect.String:
		return core.DataTypeString256
	default:
		return core.DataTypeInvalid
	}
}

func datatypeSize(t core.DataType) int {
	switch t {
	case core.DataTypeInt32, core.DataTypeFloat32:
		return 4
	case core.DataTypeInt64, core.DataTypeFloat64:
		return 8
	case core.DataTypeString8:
		return 8
	case core.DataTypeString32:
		return 32
	case core.DataTypeString64:
		return 64
	case core.DataTypeString128:
		return 128
	case core.DataTypeString256:
		return 256
	case core.DataTypeString260:
		return 260
	default:
		return 0
	}
}

// ─── Internal: encode/decode ──────────────────────────────────────────────────

func decodeData[T any](dc *definitionCore, payload []byte) (T, error) {
	var out T
	value := reflect.ValueOf(&out).Elem()
	target := value
	if target.Kind() == reflect.Pointer {
		target.Set(reflect.New(target.Type().Elem()))
		target = target.Elem()
	}
	offset := 0
	for _, field := range dc.fields {
		if len(payload) < offset+field.size {
			return out, fmt.Errorf("%w: payload too small", core.ErrDecode)
		}
		fieldValue := target
		if field.index >= 0 {
			fieldValue = target.Field(field.index)
		}
		if !fieldValue.CanSet() {
			return out, fmt.Errorf("%w: field cannot be set", core.ErrDecode)
		}
		setFieldValue(fieldValue, field.Type, payload[offset:offset+field.size])
		offset += field.size
	}
	return out, nil
}

func encodeData[T any](dc *definitionCore, value T) ([]byte, error) {
	if dc == nil {
		return nil, fmt.Errorf("%w: nil definition", core.ErrDecode)
	}
	source := reflect.ValueOf(value)
	if source.Kind() == reflect.Pointer {
		if source.IsNil() {
			return nil, fmt.Errorf("%w: nil data pointer", core.ErrDecode)
		}
		source = source.Elem()
	}
	size := 0
	for _, field := range dc.fields {
		size += field.size
	}
	payload := make([]byte, size)
	offset := 0
	for _, field := range dc.fields {
		fieldValue := source
		if field.index >= 0 {
			fieldValue = source.Field(field.index)
		}
		if err := putFieldValue(payload[offset:offset+field.size], fieldValue, field.Type); err != nil {
			return nil, err
		}
		offset += field.size
	}
	return payload, nil
}

func setFieldValue(v reflect.Value, datatype core.DataType, b []byte) {
	switch datatype {
	case core.DataTypeInt32:
		if v.Kind() == reflect.Bool {
			v.SetBool(binary.LittleEndian.Uint32(b) != 0)
			return
		}
		v.SetInt(int64(int32(binary.LittleEndian.Uint32(b))))
	case core.DataTypeInt64:
		v.SetInt(int64(binary.LittleEndian.Uint64(b)))
	case core.DataTypeFloat32:
		v.SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(b))))
	case core.DataTypeFloat64:
		v.SetFloat(math.Float64frombits(binary.LittleEndian.Uint64(b)))
	case core.DataTypeString8,
		core.DataTypeString32,
		core.DataTypeString64,
		core.DataTypeString128,
		core.DataTypeString256,
		core.DataTypeString260:
		v.SetString(fixedString(b))
	}
}

func putFieldValue(b []byte, v reflect.Value, datatype core.DataType) error {
	switch datatype {
	case core.DataTypeInt32:
		if v.Kind() == reflect.Bool {
			if v.Bool() {
				binary.LittleEndian.PutUint32(b, 1)
			}
			return nil
		}
		binary.LittleEndian.PutUint32(b, uint32(v.Int()))
	case core.DataTypeInt64:
		binary.LittleEndian.PutUint64(b, uint64(v.Int()))
	case core.DataTypeFloat32:
		binary.LittleEndian.PutUint32(b, math.Float32bits(float32(v.Float())))
	case core.DataTypeFloat64:
		binary.LittleEndian.PutUint64(b, math.Float64bits(v.Float()))
	case core.DataTypeString8,
		core.DataTypeString32,
		core.DataTypeString64,
		core.DataTypeString128,
		core.DataTypeString256,
		core.DataTypeString260:
		copy(b, v.String())
	default:
		return fmt.Errorf("%w: cannot encode datatype %d", core.ErrUnsupportedType, datatype)
	}
	return nil
}

func sendDataUpdate[T any](ctx context.Context, clientCtx context.Context, ch chan DataUpdate[T], update DataUpdate[T]) {
	defer func() { _ = recover() }()
	select {
	case ch <- update:
	case <-ctx.Done():
	case <-clientCtx.Done():
	}
}

func fixedString(b []byte) string {
	if n := bytes.IndexByte(b, 0); n >= 0 {
		return string(b[:n])
	}
	return string(b)
}
