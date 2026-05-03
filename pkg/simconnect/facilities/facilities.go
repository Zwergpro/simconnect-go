//go:build windows

// Package facilities implements the Facilities SimConnect API category.
package facilities

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// Session is the subset of client.Client methods used by this package.
type Session interface {
	NextRequestID() uint32
	NextDefinitionID() uint32
	AddWaiter(uint32) (<-chan core.RequestResult, error)
	RemoveWaiter(uint32)
	AddDataSub(uint32, func(core.Message)) error
	RemoveDataSub(uint32)
	TrackSend(uint32)
	Bindings() *bindings.SimConnect
	Context() context.Context
	ChannelBuffer() int
}

// Facilities exposes facility listing and data request functions.
type Facilities struct {
	c Session
}

func New(c Session) *Facilities {
	return &Facilities{c: c}
}

func (s *Facilities) NearbyAirports(ctx context.Context) (core.AirportListMessage, error) {
	return s.RequestNearbyAirports(ctx)
}

type FacilityDefinition struct {
	id core.DataDefinitionID
}

func (d FacilityDefinition) ID() core.DataDefinitionID { return d.id }

type FacilityDataFilter struct {
	Path string
	Data []byte
}

func (s *Facilities) RequestNearbyAirports(ctx context.Context) (core.AirportListMessage, error) {
	messages, err := s.RequestFacilitiesList(ctx, core.FacilityListTypeAirport)
	if err != nil {
		return core.AirportListMessage{}, err
	}
	var out core.AirportListMessage
	for _, msg := range messages {
		airportMsg, ok := msg.(core.AirportListMessage)
		if !ok {
			return core.AirportListMessage{}, fmt.Errorf("%w: expected airport list, got %T", core.ErrDecode, msg)
		}
		if out.OutOf == 0 || airportMsg.EntryNumber < out.EntryNumber {
			out.FacilityListMeta = airportMsg.FacilityListMeta
		}
		out.Airports = append(out.Airports, airportMsg.Airports...)
		out.ArraySize = uint32(len(out.Airports))
	}
	return out, nil
}

func (s *Facilities) RequestFacilitiesList(ctx context.Context, listType core.FacilityListType) ([]core.Message, error) {
	requestID := s.c.NextRequestID()
	return s.collectList(ctx, requestID, func() error {
		return s.c.Bindings().RequestFacilitiesList_EX1(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID))
	})
}

func (s *Facilities) RequestAllFacilities(ctx context.Context, listType core.FacilityListType) ([]core.Message, error) {
	requestID := s.c.NextRequestID()
	return s.collectList(ctx, requestID, func() error {
		return s.c.Bindings().RequestAllFacilities(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID))
	})
}

func (s *Facilities) SubscribeFacilities(ctx context.Context, listType core.FacilityListType) (<-chan core.Message, error) {
	requestID := s.c.NextRequestID()
	ch, err := s.subscribeList(requestID)
	if err != nil {
		return nil, err
	}
	if err := s.c.Bindings().SubscribeToFacilities(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveDataSub(requestID)
		close(ch)
		return nil, err
	}
	go func() {
		<-ctx.Done()
		_ = s.c.Bindings().UnsubscribeToFacilities(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType))
		s.c.RemoveDataSub(requestID)
		close(ch)
	}()
	return ch, nil
}

func (s *Facilities) SubscribeFacilitiesEX1(ctx context.Context, listType core.FacilityListType) (<-chan core.Message, <-chan core.Message, error) {
	newRequestID := s.c.NextRequestID()
	oldRequestID := s.c.NextRequestID()
	newCh, err := s.subscribeList(newRequestID)
	if err != nil {
		return nil, nil, err
	}
	oldCh, err := s.subscribeList(oldRequestID)
	if err != nil {
		s.c.RemoveDataSub(newRequestID)
		close(newCh)
		return nil, nil, err
	}
	if err := s.c.Bindings().SubscribeToFacilities_EX1(
		bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType),
		bindings.SIMCONNECT_DATA_REQUEST_ID(newRequestID),
		bindings.SIMCONNECT_DATA_REQUEST_ID(oldRequestID),
	); err != nil {
		s.c.RemoveDataSub(newRequestID)
		s.c.RemoveDataSub(oldRequestID)
		close(newCh)
		close(oldCh)
		return nil, nil, err
	}
	go func() {
		<-ctx.Done()
		_ = s.c.Bindings().UnsubscribeToFacilities_EX1(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType), true, true)
		s.c.RemoveDataSub(newRequestID)
		s.c.RemoveDataSub(oldRequestID)
		close(newCh)
		close(oldCh)
	}()
	return newCh, oldCh, nil
}

func (s *Facilities) NewFacilityDefinition(fields ...string) (FacilityDefinition, error) {
	def := FacilityDefinition{id: core.DataDefinitionID(s.c.NextDefinitionID())}
	for _, field := range fields {
		if err := s.AddToFacilityDefinition(def, field); err != nil {
			return FacilityDefinition{}, err
		}
	}
	return def, nil
}

func (s *Facilities) AddToFacilityDefinition(def FacilityDefinition, field string) error {
	return s.c.Bindings().AddToFacilityDefinition(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id), field)
}

func (s *Facilities) AddFacilityDataDefinitionFilter(def FacilityDefinition, filter FacilityDataFilter) error {
	var pData *byte
	if len(filter.Data) > 0 {
		pData = &filter.Data[0]
	}
	return s.c.Bindings().AddFacilityDataDefinitionFilter(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id), filter.Path, uint32(len(filter.Data)), unsafe.Pointer(pData))
}

func (s *Facilities) ClearAllFacilityDataDefinitionFilters(def FacilityDefinition) error {
	return s.c.Bindings().ClearAllFacilityDataDefinitionFilters(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id))
}

func (s *Facilities) RequestFacilityData(ctx context.Context, def FacilityDefinition, icao, region string) ([]core.FacilityDataMessage, error) {
	requestID := s.c.NextRequestID()
	return s.collectFacilityData(ctx, requestID, func() error {
		return s.c.Bindings().RequestFacilityData(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID), icao, region)
	})
}

func (s *Facilities) RequestFacilityDataEX1(ctx context.Context, def FacilityDefinition, icao, region string, facilityType byte) ([]core.FacilityDataMessage, error) {
	requestID := s.c.NextRequestID()
	return s.collectFacilityData(ctx, requestID, func() error {
		return s.c.Bindings().RequestFacilityData_EX1(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID), icao, region, facilityType)
	})
}

func (s *Facilities) RequestJetwayData(ctx context.Context, airportICAO string, indexes []int32) (core.JetwayDataMessage, error) {
	ch := make(chan core.JetwayDataMessage, 1)
	const jetwayRequestID uint32 = 0
	if err := s.c.AddDataSub(jetwayRequestID, func(msg core.Message) {
		m, ok := msg.(core.JetwayDataMessage)
		if !ok {
			return
		}
		select {
		case ch <- m:
		case <-s.c.Context().Done():
		}
	}); err != nil {
		return core.JetwayDataMessage{}, err
	}
	defer s.c.RemoveDataSub(jetwayRequestID)
	if err := s.c.Bindings().RequestJetwayData(airportICAO, indexes); err != nil {
		return core.JetwayDataMessage{}, err
	}
	select {
	case msg := <-ch:
		return msg, nil
	case <-ctx.Done():
		return core.JetwayDataMessage{}, ctx.Err()
	}
}

func (s *Facilities) collectList(ctx context.Context, requestID uint32, call func() error) ([]core.Message, error) {
	waiter, err := s.c.AddWaiter(requestID)
	if err != nil {
		return nil, err
	}
	packets := make(chan core.Message, 16)
	handler := func(msg core.Message) {
		if !isListMessageForRequest(msg, requestID) {
			return
		}
		select {
		case packets <- msg:
		case <-s.c.Context().Done():
		}
	}
	if err := s.c.AddDataSub(requestID, handler); err != nil {
		s.c.RemoveWaiter(requestID)
		return nil, err
	}
	defer s.c.RemoveDataSub(requestID)
	if err := call(); err != nil {
		s.c.RemoveWaiter(requestID)
		return nil, err
	}
	s.c.TrackSend(requestID)

	var out []core.Message
	for {
		select {
		case result, ok := <-waiter:
			if ok && result.Err != nil {
				return nil, result.Err
			}
			waiter = nil
		case msg := <-packets:
			out = append(out, msg)
			if listComplete(msg) {
				return out, nil
			}
		case <-ctx.Done():
			s.c.RemoveWaiter(requestID)
			return nil, ctx.Err()
		}
	}
}

func (s *Facilities) collectFacilityData(ctx context.Context, requestID uint32, call func() error) ([]core.FacilityDataMessage, error) {
	waiter, err := s.c.AddWaiter(requestID)
	if err != nil {
		return nil, err
	}
	packets := make(chan core.Message, 16)
	handler := func(msg core.Message) {
		switch m := msg.(type) {
		case core.FacilityDataMessage:
			if m.UserRequestID != requestID {
				return
			}
		case core.FacilityDataEndMessage:
			if m.RequestID != requestID {
				return
			}
		default:
			return
		}
		select {
		case packets <- msg:
		case <-s.c.Context().Done():
		}
	}
	if err := s.c.AddDataSub(requestID, handler); err != nil {
		s.c.RemoveWaiter(requestID)
		return nil, err
	}
	defer s.c.RemoveDataSub(requestID)
	if err := call(); err != nil {
		s.c.RemoveWaiter(requestID)
		return nil, err
	}
	s.c.TrackSend(requestID)

	var out []core.FacilityDataMessage
	for {
		select {
		case result, ok := <-waiter:
			if ok && result.Err != nil {
				return nil, result.Err
			}
			waiter = nil
		case msg := <-packets:
			switch m := msg.(type) {
			case core.FacilityDataMessage:
				out = append(out, m)
			case core.FacilityDataEndMessage:
				return out, nil
			}
		case <-ctx.Done():
			s.c.RemoveWaiter(requestID)
			return nil, ctx.Err()
		}
	}
}

func (s *Facilities) subscribeList(requestID uint32) (chan core.Message, error) {
	ch := make(chan core.Message, s.c.ChannelBuffer())
	handler := func(msg core.Message) {
		if !isListMessageForRequest(msg, requestID) {
			return
		}
		select {
		case ch <- msg:
		case <-s.c.Context().Done():
		}
	}
	if err := s.c.AddDataSub(requestID, handler); err != nil {
		close(ch)
		return nil, err
	}
	return ch, nil
}

func isListMessageForRequest(msg core.Message, requestID uint32) bool {
	switch m := msg.(type) {
	case core.AirportListMessage:
		return m.RequestID == requestID
	case core.WaypointListMessage:
		return m.RequestID == requestID
	case core.NDBListMessage:
		return m.RequestID == requestID
	case core.VORListMessage:
		return m.RequestID == requestID
	case core.FacilityMinimalListMessage:
		return m.RequestID == requestID
	case core.JetwayDataMessage:
		return m.RequestID == requestID
	default:
		return false
	}
}

func listComplete(msg core.Message) bool {
	switch m := msg.(type) {
	case core.AirportListMessage:
		return m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf
	case core.WaypointListMessage:
		return m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf
	case core.NDBListMessage:
		return m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf
	case core.VORListMessage:
		return m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf
	case core.FacilityMinimalListMessage:
		return m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf
	case core.JetwayDataMessage:
		return m.OutOf == 0 || m.EntryNumber+1 >= m.OutOf
	default:
		return true
	}
}
