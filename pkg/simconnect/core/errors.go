//go:build windows

package core

import (
	"errors"
	"fmt"
)

// Sentinel errors shared across all domain packages.
var (
	ErrClosed          = errors.New("simconnect api client closed")
	ErrDecode          = errors.New("simconnect api decode error")
	ErrUnsupportedType = errors.New("simconnect api unsupported type")
	ErrBusy            = errors.New("simconnect operation already pending")
)

// ExceptionError is delivered as an error when the SimConnect server reports
// a SIMCONNECT_RECV_EXCEPTION packet.  It satisfies the Message interface so
// it can be routed through the standard dispatch path.
type ExceptionError struct {
	Exception Exception
	SendID    uint32
	Index     uint32
}

func (ExceptionError) RecvID() RecvID {
	return RecvIDException
}

func (e ExceptionError) Error() string {
	return fmt.Sprintf("simconnect exception %s send_id=%d index=%d", e.Exception, e.SendID, e.Index)
}

func (e Exception) String() string {
	if name, ok := exceptionNames[e]; ok {
		return name
	}
	return "UNKNOWN"
}

var exceptionNames = map[Exception]string{
	ExceptionNone:                          "NONE",
	ExceptionErrorCode:                     "ERROR",
	ExceptionSizeMismatch:                  "SIZE_MISMATCH",
	ExceptionUnrecognizedID:                "UNRECOGNIZED_ID",
	ExceptionUnopened:                      "UNOPENED",
	ExceptionVersionMismatch:               "VERSION_MISMATCH",
	ExceptionTooManyGroups:                 "TOO_MANY_GROUPS",
	ExceptionNameUnrecognized:              "NAME_UNRECOGNIZED",
	ExceptionTooManyEventNames:             "TOO_MANY_EVENT_NAMES",
	ExceptionEventIDDuplicate:              "EVENT_ID_DUPLICATE",
	ExceptionTooManyMaps:                   "TOO_MANY_MAPS",
	ExceptionTooManyObjects:                "TOO_MANY_OBJECTS",
	ExceptionTooManyRequests:               "TOO_MANY_REQUESTS",
	ExceptionWeatherInvalidPort:            "WEATHER_INVALID_PORT",
	ExceptionWeatherInvalidMETAR:           "WEATHER_INVALID_METAR",
	ExceptionWeatherUnableToGetObservation: "WEATHER_UNABLE_TO_GET_OBSERVATION",
	ExceptionWeatherUnableToCreateStation:  "WEATHER_UNABLE_TO_CREATE_STATION",
	ExceptionWeatherUnableToRemoveStation:  "WEATHER_UNABLE_TO_REMOVE_STATION",
	ExceptionInvalidDataType:               "INVALID_DATA_TYPE",
	ExceptionInvalidDataSize:               "INVALID_DATA_SIZE",
	ExceptionDataError:                     "DATA_ERROR",
	ExceptionInvalidArray:                  "INVALID_ARRAY",
	ExceptionCreateObjectFailed:            "CREATE_OBJECT_FAILED",
	ExceptionLoadFlightplanFailed:          "LOAD_FLIGHTPLAN_FAILED",
	ExceptionOperationInvalidForObjectType: "OPERATION_INVALID_FOR_OBJECT_TYPE",
	ExceptionIllegalOperation:              "ILLEGAL_OPERATION",
	ExceptionAlreadySubscribed:             "ALREADY_SUBSCRIBED",
	ExceptionInvalidEnum:                   "INVALID_ENUM",
	ExceptionDefinitionError:               "DEFINITION_ERROR",
	ExceptionDuplicateID:                   "DUPLICATE_ID",
	ExceptionDatumID:                       "DATUM_ID",
	ExceptionOutOfBounds:                   "OUT_OF_BOUNDS",
	ExceptionAlreadyCreated:                "ALREADY_CREATED",
	ExceptionObjectOutsideRealityBubble:    "OBJECT_OUTSIDE_REALITY_BUBBLE",
	ExceptionObjectContainer:               "OBJECT_CONTAINER",
	ExceptionObjectAI:                      "OBJECT_AI",
	ExceptionObjectATC:                     "OBJECT_ATC",
	ExceptionObjectSchedule:                "OBJECT_SCHEDULE",
	ExceptionJetwayData:                    "JETWAY_DATA",
	ExceptionActionNotFound:                "ACTION_NOT_FOUND",
	ExceptionNotAnAction:                   "NOT_AN_ACTION",
	ExceptionIncorrectActionParams:         "INCORRECT_ACTION_PARAMS",
	ExceptionGetInputEventFailed:           "GET_INPUT_EVENT_FAILED",
	ExceptionSetInputEventFailed:           "SET_INPUT_EVENT_FAILED",
	ExceptionEventNameReserved:             "EVENT_NAME_RESERVED",
	ExceptionInternal:                      "INTERNAL",
	ExceptionCameraAPI:                     "CAMERA_API",
}

// RequestResult carries the outcome of a one-shot asynchronous SimConnect
// request.  Either Msg or Err is non-nil.
type RequestResult struct {
	Msg Message
	Err error
}
