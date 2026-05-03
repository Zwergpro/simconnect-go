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

// Session is the subset of client.Sim methods used by this package.
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
	session Session
}

func New(s Session) *Facilities {
	return &Facilities{session: s}
}

func (f *Facilities) NearbyAirports(ctx context.Context) (core.AirportListMessage, error) {
	return f.RequestNearbyAirports(ctx)
}

type FacilityDefinition struct {
	id core.DataDefinitionID
}

func (d FacilityDefinition) ID() core.DataDefinitionID { return d.id }

type FacilityDataFilter struct {
	Path string
	Data []byte
}

func (f *Facilities) RequestNearbyAirports(ctx context.Context) (core.AirportListMessage, error) {
	messages, err := f.RequestFacilitiesList(ctx, core.FacilityListTypeAirport)
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

func (f *Facilities) RequestFacilitiesList(ctx context.Context, listType core.FacilityListType) ([]core.Message, error) {
	requestID := f.session.NextRequestID()
	return f.collectList(ctx, requestID, func() error {
		return f.session.Bindings().RequestFacilitiesList_EX1(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID))
	})
}

func (f *Facilities) RequestAllFacilities(ctx context.Context, listType core.FacilityListType) ([]core.Message, error) {
	requestID := f.session.NextRequestID()
	return f.collectList(ctx, requestID, func() error {
		return f.session.Bindings().RequestAllFacilities(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID))
	})
}

func (f *Facilities) SubscribeFacilities(ctx context.Context, listType core.FacilityListType) (<-chan core.Message, error) {
	requestID := f.session.NextRequestID()
	ch, err := f.subscribeList(requestID)
	if err != nil {
		return nil, err
	}
	if err := f.session.Bindings().SubscribeToFacilities(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		f.session.RemoveDataSub(requestID)
		close(ch)
		return nil, err
	}
	go func() {
		<-ctx.Done()
		_ = f.session.Bindings().UnsubscribeToFacilities(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType))
		f.session.RemoveDataSub(requestID)
		close(ch)
	}()
	return ch, nil
}

func (f *Facilities) SubscribeFacilitiesEX1(ctx context.Context, listType core.FacilityListType) (<-chan core.Message, <-chan core.Message, error) {
	newRequestID := f.session.NextRequestID()
	oldRequestID := f.session.NextRequestID()
	newCh, err := f.subscribeList(newRequestID)
	if err != nil {
		return nil, nil, err
	}
	oldCh, err := f.subscribeList(oldRequestID)
	if err != nil {
		f.session.RemoveDataSub(newRequestID)
		close(newCh)
		return nil, nil, err
	}
	if err := f.session.Bindings().SubscribeToFacilities_EX1(
		bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType),
		bindings.SIMCONNECT_DATA_REQUEST_ID(newRequestID),
		bindings.SIMCONNECT_DATA_REQUEST_ID(oldRequestID),
	); err != nil {
		f.session.RemoveDataSub(newRequestID)
		f.session.RemoveDataSub(oldRequestID)
		close(newCh)
		close(oldCh)
		return nil, nil, err
	}
	go func() {
		<-ctx.Done()
		_ = f.session.Bindings().UnsubscribeToFacilities_EX1(bindings.SIMCONNECT_FACILITY_LIST_TYPE(listType), true, true)
		f.session.RemoveDataSub(newRequestID)
		f.session.RemoveDataSub(oldRequestID)
		close(newCh)
		close(oldCh)
	}()
	return newCh, oldCh, nil
}

func (f *Facilities) NewFacilityDefinition(fields ...string) (FacilityDefinition, error) {
	def := FacilityDefinition{id: core.DataDefinitionID(f.session.NextDefinitionID())}
	for _, field := range fields {
		if err := f.AddToFacilityDefinition(def, field); err != nil {
			return FacilityDefinition{}, err
		}
	}
	return def, nil
}

func (f *Facilities) AddToFacilityDefinition(def FacilityDefinition, field string) error {
	return f.session.Bindings().AddToFacilityDefinition(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id), field)
}

func (f *Facilities) AddFacilityDataDefinitionFilter(def FacilityDefinition, filter FacilityDataFilter) error {
	var pData *byte
	if len(filter.Data) > 0 {
		pData = &filter.Data[0]
	}
	return f.session.Bindings().AddFacilityDataDefinitionFilter(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id), filter.Path, uint32(len(filter.Data)), unsafe.Pointer(pData))
}

func (f *Facilities) ClearAllFacilityDataDefinitionFilters(def FacilityDefinition) error {
	return f.session.Bindings().ClearAllFacilityDataDefinitionFilters(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id))
}

func (f *Facilities) RequestFacilityData(ctx context.Context, def FacilityDefinition, icao, region string) ([]core.FacilityDataMessage, error) {
	requestID := f.session.NextRequestID()
	return f.collectFacilityData(ctx, requestID, func() error {
		return f.session.Bindings().RequestFacilityData(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID), icao, region)
	})
}

func (f *Facilities) RequestFacilityDataEX1(ctx context.Context, def FacilityDefinition, icao, region string, facilityType byte) ([]core.FacilityDataMessage, error) {
	requestID := f.session.NextRequestID()
	return f.collectFacilityData(ctx, requestID, func() error {
		return f.session.Bindings().RequestFacilityData_EX1(bindings.SIMCONNECT_DATA_DEFINITION_ID(def.id), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID), icao, region, facilityType)
	})
}

func (f *Facilities) RequestJetwayData(ctx context.Context, airportICAO string, indexes []int32) (core.JetwayDataMessage, error) {
	ch := make(chan core.JetwayDataMessage, 1)
	const jetwayRequestID uint32 = 0
	if err := f.session.AddDataSub(jetwayRequestID, func(msg core.Message) {
		m, ok := msg.(core.JetwayDataMessage)
		if !ok {
			return
		}
		select {
		case ch <- m:
		case <-f.session.Context().Done():
		}
	}); err != nil {
		return core.JetwayDataMessage{}, err
	}
	defer f.session.RemoveDataSub(jetwayRequestID)
	if err := f.session.Bindings().RequestJetwayData(airportICAO, indexes); err != nil {
		return core.JetwayDataMessage{}, err
	}
	select {
	case msg := <-ch:
		return msg, nil
	case <-ctx.Done():
		return core.JetwayDataMessage{}, ctx.Err()
	}
}

func (f *Facilities) collectList(ctx context.Context, requestID uint32, call func() error) ([]core.Message, error) {
	waiter, err := f.session.AddWaiter(requestID)
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
		case <-f.session.Context().Done():
		}
	}
	if err := f.session.AddDataSub(requestID, handler); err != nil {
		f.session.RemoveWaiter(requestID)
		return nil, err
	}
	defer f.session.RemoveDataSub(requestID)
	if err := call(); err != nil {
		f.session.RemoveWaiter(requestID)
		return nil, err
	}
	f.session.TrackSend(requestID)

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
			f.session.RemoveWaiter(requestID)
			return nil, ctx.Err()
		}
	}
}

func (f *Facilities) collectFacilityData(ctx context.Context, requestID uint32, call func() error) ([]core.FacilityDataMessage, error) {
	waiter, err := f.session.AddWaiter(requestID)
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
		case <-f.session.Context().Done():
		}
	}
	if err := f.session.AddDataSub(requestID, handler); err != nil {
		f.session.RemoveWaiter(requestID)
		return nil, err
	}
	defer f.session.RemoveDataSub(requestID)
	if err := call(); err != nil {
		f.session.RemoveWaiter(requestID)
		return nil, err
	}
	f.session.TrackSend(requestID)

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
			f.session.RemoveWaiter(requestID)
			return nil, ctx.Err()
		}
	}
}

func (f *Facilities) subscribeList(requestID uint32) (chan core.Message, error) {
	ch := make(chan core.Message, f.session.ChannelBuffer())
	handler := func(msg core.Message) {
		if !isListMessageForRequest(msg, requestID) {
			return
		}
		select {
		case ch <- msg:
		case <-f.session.Context().Done():
		}
	}
	if err := f.session.AddDataSub(requestID, handler); err != nil {
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
