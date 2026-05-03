//go:build windows

// Package ai implements the AI_Object SimConnect API category.
package ai

import (
	"context"
	"fmt"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// Session is the subset of client.Client methods used by this package.
type Session interface {
	NextRequestID() uint32
	AddWaiter(uint32) (<-chan core.RequestResult, error)
	RemoveWaiter(uint32)
	TrackSend(uint32)
	Bindings() *bindings.SimConnect
}

// AI exposes AI object creation and control functions.
type AI struct {
	c Session
}

// New creates an AI client wrapping the core SimConnect session.
func New(c Session) *AI {
	return &AI{c: c}
}

func (s *AI) CreateParkedATC(ctx context.Context, containerTitle, tailNumber, airportID string) (core.ObjectID, error) {
	requestID, waiter, err := s.startRequest()
	if err != nil {
		return 0, err
	}
	if err := s.c.Bindings().AICreateParkedATCAircraft(containerTitle, tailNumber, airportID, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveWaiter(requestID)
		return 0, err
	}
	s.c.TrackSend(requestID)
	return s.waitAssignedObjectID(ctx, requestID, waiter)
}

func (s *AI) CreateParkedATCEX1(ctx context.Context, containerTitle, livery, tailNumber, airportID string) (core.ObjectID, error) {
	requestID, waiter, err := s.startRequest()
	if err != nil {
		return 0, err
	}
	if err := s.c.Bindings().AICreateParkedATCAircraft_EX1(containerTitle, livery, tailNumber, airportID, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveWaiter(requestID)
		return 0, err
	}
	s.c.TrackSend(requestID)
	return s.waitAssignedObjectID(ctx, requestID, waiter)
}

func (s *AI) CreateEnrouteATC(ctx context.Context, containerTitle, tailNumber string, flightNumber int32, flightPlanPath string, flightPlanPosition float64, touchAndGo bool) (core.ObjectID, error) {
	requestID, waiter, err := s.startRequest()
	if err != nil {
		return 0, err
	}
	if err := s.c.Bindings().AICreateEnrouteATCAircraft(containerTitle, tailNumber, flightNumber, flightPlanPath, flightPlanPosition, touchAndGo, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveWaiter(requestID)
		return 0, err
	}
	s.c.TrackSend(requestID)
	return s.waitAssignedObjectID(ctx, requestID, waiter)
}

func (s *AI) CreateEnrouteATCEX1(ctx context.Context, containerTitle, livery, tailNumber string, flightNumber int32, flightPlanPath string, flightPlanPosition float64, touchAndGo bool) (core.ObjectID, error) {
	requestID, waiter, err := s.startRequest()
	if err != nil {
		return 0, err
	}
	if err := s.c.Bindings().AICreateEnrouteATCAircraft_EX1(containerTitle, livery, tailNumber, flightNumber, flightPlanPath, flightPlanPosition, touchAndGo, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveWaiter(requestID)
		return 0, err
	}
	s.c.TrackSend(requestID)
	return s.waitAssignedObjectID(ctx, requestID, waiter)
}

func (s *AI) CreateNonATC(ctx context.Context, containerTitle, tailNumber string, initPos core.InitPosition) (core.ObjectID, error) {
	requestID, waiter, err := s.startRequest()
	if err != nil {
		return 0, err
	}
	if err := s.c.Bindings().AICreateNonATCAircraft(containerTitle, tailNumber, core.InitPositionToBinding(initPos), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveWaiter(requestID)
		return 0, err
	}
	s.c.TrackSend(requestID)
	return s.waitAssignedObjectID(ctx, requestID, waiter)
}

func (s *AI) CreateNonATCEX1(ctx context.Context, containerTitle, livery, tailNumber string, initPos core.InitPosition) (core.ObjectID, error) {
	requestID, waiter, err := s.startRequest()
	if err != nil {
		return 0, err
	}
	if err := s.c.Bindings().AICreateNonATCAircraft_EX1(containerTitle, livery, tailNumber, core.InitPositionToBinding(initPos), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveWaiter(requestID)
		return 0, err
	}
	s.c.TrackSend(requestID)
	return s.waitAssignedObjectID(ctx, requestID, waiter)
}

func (s *AI) CreateSimulatedObject(ctx context.Context, containerTitle string, initPos core.InitPosition) (core.ObjectID, error) {
	requestID, waiter, err := s.startRequest()
	if err != nil {
		return 0, err
	}
	if err := s.c.Bindings().AICreateSimulatedObject(containerTitle, core.InitPositionToBinding(initPos), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveWaiter(requestID)
		return 0, err
	}
	s.c.TrackSend(requestID)
	return s.waitAssignedObjectID(ctx, requestID, waiter)
}

func (s *AI) CreateSimulatedObjectEX1(ctx context.Context, containerTitle, livery string, initPos core.InitPosition) (core.ObjectID, error) {
	requestID, waiter, err := s.startRequest()
	if err != nil {
		return 0, err
	}
	if err := s.c.Bindings().AICreateSimulatedObject_EX1(containerTitle, livery, core.InitPositionToBinding(initPos), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		s.c.RemoveWaiter(requestID)
		return 0, err
	}
	s.c.TrackSend(requestID)
	return s.waitAssignedObjectID(ctx, requestID, waiter)
}

func (s *AI) ReleaseControl(objectID core.ObjectID) error {
	requestID := s.c.NextRequestID()
	if err := s.c.Bindings().AIReleaseControl(bindings.SIMCONNECT_OBJECT_ID(objectID), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		return err
	}
	s.c.TrackSend(requestID)
	return nil
}

func (s *AI) Remove(objectID core.ObjectID) error {
	requestID := s.c.NextRequestID()
	if err := s.c.Bindings().AIRemoveObject(bindings.SIMCONNECT_OBJECT_ID(objectID), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		return err
	}
	s.c.TrackSend(requestID)
	return nil
}

// SetFlightPlan assigns a flight plan to an AI aircraft.
// Placed here per apidocs AI_Object category.
func (s *AI) SetFlightPlan(objectID core.ObjectID, flightPlanPath string) error {
	requestID := s.c.NextRequestID()
	if err := s.c.Bindings().AISetAircraftFlightPlan(bindings.SIMCONNECT_OBJECT_ID(objectID), flightPlanPath, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		return err
	}
	s.c.TrackSend(requestID)
	return nil
}

func (s *AI) startRequest() (uint32, <-chan core.RequestResult, error) {
	requestID := s.c.NextRequestID()
	waiter, err := s.c.AddWaiter(requestID)
	if err != nil {
		return 0, nil, err
	}
	return requestID, waiter, nil
}

func (s *AI) waitAssignedObjectID(ctx context.Context, requestID uint32, waiter <-chan core.RequestResult) (core.ObjectID, error) {
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		s.c.RemoveWaiter(requestID)
		return 0, ctx.Err()
	}
	if result.Err != nil {
		return 0, result.Err
	}
	msg, ok := result.Msg.(core.AssignedObjectIDMessage)
	if !ok {
		return 0, fmt.Errorf("%w: expected assigned object id, got %T", core.ErrDecode, result.Msg)
	}
	return msg.ObjectID, nil
}
