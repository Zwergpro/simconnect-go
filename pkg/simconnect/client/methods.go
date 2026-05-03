//go:build windows

package client

import (
	"unsafe"

	"github.com/Zwergpro/simconnect-go/pkg/bindings"
	"github.com/Zwergpro/simconnect-go/pkg/simconnect/core"
)

// ClearDataDefinition clears all data definitions for the given define ID.
func (s *Sim) ClearDataDefinition(defineID core.DataDefinitionID) error {
	return s.raw.ClearDataDefinition(bindings.SIMCONNECT_DATA_DEFINITION_ID(defineID))
}

// MapInputEventToClientEvent maps an input event to a client event.
func (s *Sim) MapInputEventToClientEvent(groupID core.InputGroupID, inputDefinition string, downEventID core.EventID, downValue uint32, upEventID core.EventID, upValue uint32, maskable bool) error {
	return s.raw.MapInputEventToClientEvent(
		bindings.SIMCONNECT_INPUT_GROUP_ID(groupID),
		inputDefinition,
		bindings.SIMCONNECT_CLIENT_EVENT_ID(downEventID),
		downValue,
		bindings.SIMCONNECT_CLIENT_EVENT_ID(upEventID),
		upValue,
		maskable,
	)
}

// MapInputEventToClientEventEX1 maps an input event to a client event (EX1 variant).
func (s *Sim) MapInputEventToClientEventEX1(groupID core.InputGroupID, inputDefinition string, downEventID core.EventID, downValue uint32, upEventID core.EventID, upValue uint32, maskable bool) error {
	return s.raw.MapInputEventToClientEvent_EX1(
		bindings.SIMCONNECT_INPUT_GROUP_ID(groupID),
		inputDefinition,
		bindings.SIMCONNECT_CLIENT_EVENT_ID(downEventID),
		downValue,
		bindings.SIMCONNECT_CLIENT_EVENT_ID(upEventID),
		upValue,
		maskable,
	)
}

// SetInputGroupPriority sets the priority for an input group.
func (s *Sim) SetInputGroupPriority(groupID core.InputGroupID, priority uint32) error {
	return s.raw.SetInputGroupPriority(bindings.SIMCONNECT_INPUT_GROUP_ID(groupID), priority)
}

// RemoveInputEvent removes an input event from an input group.
func (s *Sim) RemoveInputEvent(groupID core.InputGroupID, inputDefinition string) error {
	return s.raw.RemoveInputEvent(bindings.SIMCONNECT_INPUT_GROUP_ID(groupID), inputDefinition)
}

// ClearInputGroup clears all input events from an input group.
func (s *Sim) ClearInputGroup(groupID core.InputGroupID) error {
	return s.raw.ClearInputGroup(bindings.SIMCONNECT_INPUT_GROUP_ID(groupID))
}

// SetInputGroupState sets the state of an input group.
func (s *Sim) SetInputGroupState(groupID core.InputGroupID, state core.State) error {
	return s.raw.SetInputGroupState(bindings.SIMCONNECT_INPUT_GROUP_ID(groupID), uint32(state))
}

// SubscribeToCommBusEvent subscribes to a CommBus event.
func (s *Sim) SubscribeToCommBusEvent(eventID core.EventID, eventName string) error {
	return s.raw.SubscribeToCommBusEvent(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID), eventName)
}

// UnsubscribeToCommBusEvent unsubscribes from a CommBus event.
func (s *Sim) UnsubscribeToCommBusEvent(eventID core.EventID) error {
	return s.raw.UnsubscribeToCommBusEvent(bindings.SIMCONNECT_CLIENT_EVENT_ID(eventID))
}

// SetDataOnSimObject sets data on a sim object.
func (s *Sim) SetDataOnSimObject(defineID core.DataDefinitionID, objectID core.ObjectID, flags core.DataSetFlag, data []byte) error {
	var pData *byte
	if len(data) > 0 {
		pData = &data[0]
	}
	return s.raw.SetDataOnSimObject(
		bindings.SIMCONNECT_DATA_DEFINITION_ID(defineID),
		bindings.SIMCONNECT_OBJECT_ID(objectID),
		bindings.SIMCONNECT_DATA_SET_FLAG(flags),
		0,
		uint32(len(data)),
		unsafe.Pointer(pData),
	)
}
