//go:build windows

// Package ai implements the AI_Object SimConnect API category.
package ai

import (
	"context"
	"fmt"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// Session is the subset of client.Sim methods used by this package.
type Session interface {
	NextRequestID() uint32
	AddWaiter(uint32) (<-chan core.RequestResult, error)
	RemoveWaiter(uint32)
	TrackSend(uint32)
	Bindings() *bindings.SimConnect
}

// AI exposes AI object creation and control functions.
type AI struct {
	session Session
}

// New creates an AI client wrapping the core SimConnect session.
func New(s Session) *AI {
	return &AI{session: s}
}

func (ai *AI) CreateParkedATC(ctx context.Context, containerTitle, tailNumber, airportID string) (core.ObjectID, error) {
	requestID, waiter, err := ai.startRequest()
	if err != nil {
		return 0, err
	}
	if err := ai.session.Bindings().AICreateParkedATCAircraft(containerTitle, tailNumber, airportID, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		ai.session.RemoveWaiter(requestID)
		return 0, err
	}
	ai.session.TrackSend(requestID)
	return ai.waitAssignedObjectID(ctx, requestID, waiter)
}

func (ai *AI) CreateParkedATCEX1(ctx context.Context, containerTitle, livery, tailNumber, airportID string) (core.ObjectID, error) {
	requestID, waiter, err := ai.startRequest()
	if err != nil {
		return 0, err
	}
	if err := ai.session.Bindings().AICreateParkedATCAircraft_EX1(containerTitle, livery, tailNumber, airportID, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		ai.session.RemoveWaiter(requestID)
		return 0, err
	}
	ai.session.TrackSend(requestID)
	return ai.waitAssignedObjectID(ctx, requestID, waiter)
}

func (ai *AI) CreateEnrouteATC(ctx context.Context, containerTitle, tailNumber string, flightNumber int32, flightPlanPath string, flightPlanPosition float64, touchAndGo bool) (core.ObjectID, error) {
	requestID, waiter, err := ai.startRequest()
	if err != nil {
		return 0, err
	}
	if err := ai.session.Bindings().AICreateEnrouteATCAircraft(containerTitle, tailNumber, flightNumber, flightPlanPath, flightPlanPosition, touchAndGo, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		ai.session.RemoveWaiter(requestID)
		return 0, err
	}
	ai.session.TrackSend(requestID)
	return ai.waitAssignedObjectID(ctx, requestID, waiter)
}

func (ai *AI) CreateEnrouteATCEX1(ctx context.Context, containerTitle, livery, tailNumber string, flightNumber int32, flightPlanPath string, flightPlanPosition float64, touchAndGo bool) (core.ObjectID, error) {
	requestID, waiter, err := ai.startRequest()
	if err != nil {
		return 0, err
	}
	if err := ai.session.Bindings().AICreateEnrouteATCAircraft_EX1(containerTitle, livery, tailNumber, flightNumber, flightPlanPath, flightPlanPosition, touchAndGo, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		ai.session.RemoveWaiter(requestID)
		return 0, err
	}
	ai.session.TrackSend(requestID)
	return ai.waitAssignedObjectID(ctx, requestID, waiter)
}

func (ai *AI) CreateNonATC(ctx context.Context, containerTitle, tailNumber string, initPos core.InitPosition) (core.ObjectID, error) {
	requestID, waiter, err := ai.startRequest()
	if err != nil {
		return 0, err
	}
	if err := ai.session.Bindings().AICreateNonATCAircraft(containerTitle, tailNumber, core.InitPositionToBinding(initPos), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		ai.session.RemoveWaiter(requestID)
		return 0, err
	}
	ai.session.TrackSend(requestID)
	return ai.waitAssignedObjectID(ctx, requestID, waiter)
}

func (ai *AI) CreateNonATCEX1(ctx context.Context, containerTitle, livery, tailNumber string, initPos core.InitPosition) (core.ObjectID, error) {
	requestID, waiter, err := ai.startRequest()
	if err != nil {
		return 0, err
	}
	if err := ai.session.Bindings().AICreateNonATCAircraft_EX1(containerTitle, livery, tailNumber, core.InitPositionToBinding(initPos), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		ai.session.RemoveWaiter(requestID)
		return 0, err
	}
	ai.session.TrackSend(requestID)
	return ai.waitAssignedObjectID(ctx, requestID, waiter)
}

func (ai *AI) CreateSimulatedObject(ctx context.Context, containerTitle string, initPos core.InitPosition) (core.ObjectID, error) {
	requestID, waiter, err := ai.startRequest()
	if err != nil {
		return 0, err
	}
	if err := ai.session.Bindings().AICreateSimulatedObject(containerTitle, core.InitPositionToBinding(initPos), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		ai.session.RemoveWaiter(requestID)
		return 0, err
	}
	ai.session.TrackSend(requestID)
	return ai.waitAssignedObjectID(ctx, requestID, waiter)
}

func (ai *AI) CreateSimulatedObjectEX1(ctx context.Context, containerTitle, livery string, initPos core.InitPosition) (core.ObjectID, error) {
	requestID, waiter, err := ai.startRequest()
	if err != nil {
		return 0, err
	}
	if err := ai.session.Bindings().AICreateSimulatedObject_EX1(containerTitle, livery, core.InitPositionToBinding(initPos), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		ai.session.RemoveWaiter(requestID)
		return 0, err
	}
	ai.session.TrackSend(requestID)
	return ai.waitAssignedObjectID(ctx, requestID, waiter)
}

func (ai *AI) ReleaseControl(objectID core.ObjectID) error {
	requestID := ai.session.NextRequestID()
	if err := ai.session.Bindings().AIReleaseControl(bindings.SIMCONNECT_OBJECT_ID(objectID), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		return err
	}
	ai.session.TrackSend(requestID)
	return nil
}

func (ai *AI) Remove(objectID core.ObjectID) error {
	requestID := ai.session.NextRequestID()
	if err := ai.session.Bindings().AIRemoveObject(bindings.SIMCONNECT_OBJECT_ID(objectID), bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		return err
	}
	ai.session.TrackSend(requestID)
	return nil
}

// SetFlightPlan assigns a flight plan to an AI aircraft.
// Placed here per apidocs AI_Object category.
func (ai *AI) SetFlightPlan(objectID core.ObjectID, flightPlanPath string) error {
	requestID := ai.session.NextRequestID()
	if err := ai.session.Bindings().AISetAircraftFlightPlan(bindings.SIMCONNECT_OBJECT_ID(objectID), flightPlanPath, bindings.SIMCONNECT_DATA_REQUEST_ID(requestID)); err != nil {
		return err
	}
	ai.session.TrackSend(requestID)
	return nil
}

func (ai *AI) startRequest() (uint32, <-chan core.RequestResult, error) {
	requestID := ai.session.NextRequestID()
	waiter, err := ai.session.AddWaiter(requestID)
	if err != nil {
		return 0, nil, err
	}
	return requestID, waiter, nil
}

func (ai *AI) waitAssignedObjectID(ctx context.Context, requestID uint32, waiter <-chan core.RequestResult) (core.ObjectID, error) {
	var result core.RequestResult
	select {
	case result = <-waiter:
	case <-ctx.Done():
		ai.session.RemoveWaiter(requestID)
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
