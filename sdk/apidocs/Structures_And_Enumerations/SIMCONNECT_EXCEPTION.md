SIMCONNECT\_EXCEPTION

## SIMCONNECT\_EXCEPTION

The **SIMCONNECT\_EXCEPTION** enumeration type is used with the SIMCONNECT\_RECV\_EXCEPTION structure to return information on an error that has occurred.

##### Syntax

```cpp
enum SIMCONNECT_EXCEPTION{
    SIMCONNECT_EXCEPTION_NONE,
    SIMCONNECT_EXCEPTION_ERROR,
    SIMCONNECT_EXCEPTION_SIZE_MISMATCH,
    SIMCONNECT_EXCEPTION_UNRECOGNIZED_ID,
    SIMCONNECT_EXCEPTION_UNOPENED,
    SIMCONNECT_EXCEPTION_VERSION_MISMATCH,
    SIMCONNECT_EXCEPTION_TOO_MANY_GROUPS,
    SIMCONNECT_EXCEPTION_NAME_UNRECOGNIZED,
    SIMCONNECT_EXCEPTION_TOO_MANY_EVENT_NAMES,
    SIMCONNECT_EXCEPTION_EVENT_ID_DUPLICATE,
    SIMCONNECT_EXCEPTION_TOO_MANY_MAPS,
    SIMCONNECT_EXCEPTION_TOO_MANY_OBJECTS,
    SIMCONNECT_EXCEPTION_TOO_MANY_REQUESTS,
    SIMCONNECT_EXCEPTION_WEATHER_INVALID_PORT,
    SIMCONNECT_EXCEPTION_WEATHER_INVALID_METAR,
    SIMCONNECT_EXCEPTION_WEATHER_UNABLE_TO_GET_OBSERVATION,
    SIMCONNECT_EXCEPTION_WEATHER_UNABLE_TO_CREATE_STATION,
    SIMCONNECT_EXCEPTION_WEATHER_UNABLE_TO_REMOVE_STATION,
    SIMCONNECT_EXCEPTION_INVALID_DATA_TYPE,
    SIMCONNECT_EXCEPTION_INVALID_DATA_SIZE,
    SIMCONNECT_EXCEPTION_DATA_ERROR,
    SIMCONNECT_EXCEPTION_INVALID_ARRAY,
    SIMCONNECT_EXCEPTION_CREATE_OBJECT_FAILED,
    SIMCONNECT_EXCEPTION_LOAD_FLIGHTPLAN_FAILED,
    SIMCONNECT_EXCEPTION_OPERATION_INVALID_FOR_OBJECT_TYPE,
    SIMCONNECT_EXCEPTION_ILLEGAL_OPERATION,
    SIMCONNECT_EXCEPTION_ALREADY_SUBSCRIBED,
    SIMCONNECT_EXCEPTION_INVALID_ENUM,
    SIMCONNECT_EXCEPTION_DEFINITION_ERROR,
    SIMCONNECT_EXCEPTION_DUPLICATE_ID,
    SIMCONNECT_EXCEPTION_DATUM_ID,
    SIMCONNECT_EXCEPTION_OUT_OF_BOUNDS,
    SIMCONNECT_EXCEPTION_ALREADY_CREATED,
    SIMCONNECT_EXCEPTION_OBJECT_OUTSIDE_REALITY_BUBBLE,
    SIMCONNECT_EXCEPTION_OBJECT_CONTAINER,
    SIMCONNECT_EXCEPTION_OBJECT_AI,
    SIMCONNECT_EXCEPTION_OBJECT_ATC,
    SIMCONNECT_EXCEPTION_OBJECT_SCHEDULE,
    SIMCONNECT_EXCEPTION_JETWAY_DATA,
    SIMCONNECT_EXCEPTION_ACTION_NOT_FOUND,
    SIMCONNECT_EXCEPTION_NOT_AN_ACTION,
    SIMCONNECT_EXCEPTION_INCORRECT_ACTION_PARAMS,
    SIMCONNECT_EXCEPTION_GET_INPUT_EVENT_FAILED,
    SIMCONNECT_EXCEPTION_SET_INPUT_EVENT_FAILED,
    SIMCONNECT_EXCEPTION_EVENT_NAME_RESERVED,
    SIMCONNECT_EXCEPTION_INTERNAL,
    SIMCONNECT_EXCEPTION_CAMERA_API
    };
```

##### Members

| Member | Description |
| `SIMCONNECT_EXCEPTION_NONE` | Specifies that there has not been an error. This value is not currently used. |
| `SIMCONNECT_EXCEPTION_ERROR` | An unspecific error has occurred. This can be from incorrect flag settings, null or incorrect parameters, the need to have at least one up or down event with an input event, failed calls from the SimConnect server to the operating system, among other reasons. |
| `SIMCONNECT_EXCEPTION_SIZE_MISMATCH` | Specifies the size of the data provided does not match the size required. This typically occurs when the wrong string length, fixed or variable, is involved. |
| `SIMCONNECT_EXCEPTION_UNRECOGNIZED_ID` | Specifies that the client event, request ID, data definition ID, or object ID was not recognized. |
| `SIMCONNECT_EXCEPTION_UNOPENED` | Specifies that communication with the SimConnect server has not been opened. This error is not currently used. |
| `SIMCONNECT_EXCEPTION_VERSION_MISMATCH` | Specifies a versioning error has occurred. Typically this will occur when a client built on a newer version of the SimConnect client dll attempts to work with an older version of the SimConnect server. |
| `SIMCONNECT_EXCEPTION_TOO_MANY_GROUPS` | Specifies that the maximum number of groups allowed has been reached. The maximum is 20. |
| `SIMCONNECT_EXCEPTION_NAME_UNRECOGNIZED` | Specifies that the simulation event name (such as "brakes") is not recognized. |
| `SIMCONNECT_EXCEPTION_TOO_MANY_EVENT_NAMES` | Specifies that the maximum number of event names allowed has been reached. The maximum is 1000. |
| `SIMCONNECT_EXCEPTION_EVENT_ID_DUPLICATE` | Specifies that the event ID has been used already. This can occur with calls to [SimConnect\_MapClientEventToSimEvent](../Events_And_Data/SimConnect_MapClientEventToSimEvent.htm), or [SimConnect\_SubscribeToSystemEvent](../General/SimConnect_SubscribeToSystemEvent.htm). |
| `SIMCONNECT_EXCEPTION_TOO_MANY_MAPS` | Specifies that the maximum number of mappings allowed has been reached. The maximum is 20. |
| `SIMCONNECT_EXCEPTION_TOO_MANY_OBJECTS` | Specifies that the maximum number of objects allowed has been reached. The maximum is 1000. |
| `SIMCONNECT_EXCEPTION_TOO_MANY_REQUESTS` | Specifies that the maximum number of requests allowed has been reached. The maximum is 1000. |
| `SIMCONNECT_EXCEPTION_WEATHER_INVALID_PORT` | Specifies an invalid port number was requested.<br>**NOTE**: This is a legacy exception, and no longer used in the simulation. |
| `SIMCONNECT_EXCEPTION_WEATHER_INVALID_METAR` | Specifies that the metar data supplied did not match the required format.<br>**NOTE**: This is a legacy exception, and no longer used in the simulation. |
| `SIMCONNECT_EXCEPTION_WEATHER_UNABLE_TO_GET_OBSERVATION` | Specifies that the weather observation requested was not available.<br>**NOTE**: This is a legacy exception, and no longer used in the simulation. |
| `SIMCONNECT_EXCEPTION_WEATHER_UNABLE_TO_CREATE_STATION` | Specifies that the weather station could not be created.<br>**NOTE**: This is a legacy exception, and no longer used in the simulation. |
| `SIMCONNECT_EXCEPTION_WEATHER_UNABLE_TO_REMOVE_STATION` | Specifies that the weather station could not be removed.<br>**NOTE**: This is a legacy exception, and no longer used in the simulation. |
| `SIMCONNECT_EXCEPTION_INVALID_DATA_TYPE` | Specifies that the data type requested does not apply to the type of data requested. Typically this occurs with a fixed length string of the wrong length. |
| `SIMCONNECT_EXCEPTION_INVALID_DATA_SIZE` | Specifies that the size of the data provided is not what is expected. This can occur when the size of a structure provided does not match the size given, or a null string entry is made for a menu or sub-menu entry text, or data with a size of zero is added to a data definition. It can also occur with an invalid request to [SimConnect\_CreateClientData](../Events_And_Data/SimConnect_CreateClientData.htm). |
| `SIMCONNECT_EXCEPTION_DATA_ERROR` | Specifies a generic data error. This error is used by the SimConnect\_WeatherCreateThermal function to report incorrect parameters, such as negative radii or values greater than the maximum allowed. It is also used by the [SimConnect\_FlightSave](../Flights/SimConnect_FlightSave.htm) and [SimConnect\_FlightLoad](../Flights/SimConnect_FlightLoad.htm) functions to report incorrect file types. It is also used by other functions to report that flags or reserved parameters have not been set to zero. |
| `SIMCONNECT_EXCEPTION_INVALID_ARRAY` | Specifies an invalid array has been sent to the [SimConnect\_SetDataOnSimObject](../Events_And_Data/SimConnect_SetDataOnSimObject.htm) function. |
| `SIMCONNECT_EXCEPTION_CREATE_OBJECT_FAILED` | Specifies that the attempt to create an AI object failed. |
| `SIMCONNECT_EXCEPTION_LOAD_FLIGHTPLAN_FAILED` | Specifies that the specified flight plan could not be found, or did not load correctly. |
| `SIMCONNECT_EXCEPTION_OPERATION_INVALID_FOR_OJBECT_TYPE` | Specifies that the operation requested does not apply to the object type, for example trying to set a flight plan on an object that is not an aircraft will result in this error. |
| `SIMCONNECT_EXCEPTION_ILLEGAL_OPERATION` | Specifies that the AI operation requested cannot be completed, such as requesting that an object be removed when the client did not create that object. This error also applies to the Weather system. |
| `SIMCONNECT_EXCEPTION_ALREADY_SUBSCRIBED` | Specifies that the client has already subscribed to that event. |
| `SIMCONNECT_EXCEPTION_INVALID_ENUM` | Specifies that the member of the enumeration provided was not valid. Currently this is only used if an unknown type is provided to [SimConnect\_RequestDataOnSimObjectType](../Events_And_Data/SimConnect_RequestDataOnSimObjectType.htm). |
| `SIMCONNECT_EXCEPTION_DEFINITION_ERROR` | Specifies that there is a problem with a data definition. Currently this is only used if a variable length definition is sent with [SimConnect\_RequestDataOnSimObject](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm). |
| `SIMCONNECT_EXCEPTION_DUPLICATE_ID` | Specifies that the ID has already been used. This can occur with menu IDs, or with the IDs provided to [SimConnect\_AddToDataDefinition](../Events_And_Data/SimConnect_AddToDataDefinition.htm), [SimConnect\_AddClientEventToNotificationGroup](../Events_And_Data/SimConnect_AddClientEventToNotificationGroup.htm) or [SimConnect\_MapClientDataNameToID](../Events_And_Data/SimConnect_MapClientDataNameToID.htm). |
| `SIMCONNECT_EXCEPTION_DATUM_ID` | Specifies that the datum ID is not recognized. This currently occurs with a call to the [SimConnect\_SetDataOnSimObject](../Events_And_Data/SimConnect_SetDataOnSimObject.htm) function. |
| `SIMCONNECT_EXCEPTION_OUT_OF_BOUNDS` | Specifies that the radius given in the [SimConnect\_RequestDataOnSimObjectType](../Events_And_Data/SimConnect_RequestDataOnSimObjectType.htm) was outside the acceptable range, or with an invalid request to [SimConnect\_CreateClientData](../Events_And_Data/SimConnect_CreateClientData.htm). |
| `SIMCONNECT_EXCEPTION_ALREADY_CREATED` | Specifies that a client data area with the name requested by a call to [SimConnect\_MapClientDataNameToID](../Events_And_Data/SimConnect_MapClientDataNameToID.htm) has already been created by another addon. Try again with a different name. |
| `SIMCONNECT_EXCEPTION_OBJECT_OUTSIDE_REALITY_BUBBLE` | Specifies that an attempt to create an ATC controlled AI object failed because the location of the object is outside the reality bubble. |
| `SIMCONNECT_EXCEPTION_OBJECT_CONTAINER` | Specifies that an attempt to create an AI object failed because of an error with the container system for the object. |
| `SIMCONNECT_EXCEPTION_OBJECT_AI` | Specifies that an attempt to create an AI object failed because of an error with the AI system for the object. |
| `SIMCONNECT_EXCEPTION_OBJECT_ATC` | Specifies that an attempt to create an AI object failed because of an error with the ATC system for the object. |
| `SIMCONNECT_EXCEPTION_OBJECT_SCHEDULE` | Specifies that an attempt to create an AI object failed because of a scheduling problem. |
| `SIMCONNECT_EXCEPTION_JETWAY_DATA` | Specifies that an attempt to retrieve jetway data using [SimConnect\_RequestJetwayData](../Facilities/SimConnect_RequestJetwayData.htm) has caused an exception. |
| `SIMCONNECT_EXCEPTION_ACTION_NOT_FOUND` | Specifies that the given action cannot be found when using the [SimConnect\_ExecuteAction](../General/SimConnect_ExecuteAction.htm) function. |
| `SIMCONNECT_EXCEPTION_NOT_AN_ACTION` | Specifies that the given action does not exist when using the [SimConnect\_ExecuteAction](../General/SimConnect_ExecuteAction.htm) function. |
| `SIMCONNECT_EXCEPTION_INCORRECT_ACTION_PARAMS` | Specifies that the wrong parameters have been given to the function [SimConnect\_ExecuteAction](../General/SimConnect_ExecuteAction.htm). |
| `SIMCONNECT_EXCEPTION_GET_INPUT_EVENT_FAILED` | This means that the wrong name/hash has been passed to the [SimConnect\_GetInputEvent](../InputEvents/SimConnect_GetInputEvent.htm) function. |
| `SIMCONNECT_EXCEPTION_SET_INPUT_EVENT_FAILED` | This means that the wrong name/hash has been passed to the [SimConnect\_SetInputEvent](../InputEvents/SimConnect_SetInputEvent.htm) function. |
| `SIMCONNECT_EXCEPTION_EVENT_NAME_RESERVED` | This means that you tried to register a commBus event using [SimConnect\_SubscribeToCommBusEvent](../Communication/SimConnect_SubscribeToCommBusEvent.htm) , however the name used was one of the ones reserved by the simulation for internal use. |
| `SIMCONNECT_EXCEPTION_INTERNAL` | This means that an internal error has occurred while using the SimConnect API. |
| `SIMCONNECT_EXCEPTION_CAMERA_API` | This means that the camera could not be acquired due to an error with the function [SimConnect\_CameraAcquire](../Camera/SimConnect_CameraAcquire.htm). In addition, this exception will be raised for every SimConnect [Camera](../Camera/Camera_API.htm) function if the camera has not been acquired previously. |

##### Remarks

In the context of SimConnect, exceptions are error codes, and should not be confused with the C# or system concepts of exceptions.

Refer to the remarks for [SimConnect\_GetLastSentPacketID](../Debug/SimConnect_GetLastSentPacketID.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_EXCEPTION](SIMCONNECT_EXCEPTION.htm)
4. [SIMCONNECT\_RECV\_ID](SIMCONNECT_RECV_ID.htm)
5. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

Report An Issue

Please explain the issue:

0/255

SendCancel

Docs

[©2026 Microsoft](https://www.microsoft.com/)

[Privacy Policy](https://privacy.microsoft.com/en-us/privacystatement)

[SDK Dev Support](https://devsupport.flightsimulator.com/)

[MSFS Forums](https://forums.flightsimulator.com/)

[MSFS2020 SDK Documentation](https://docs.flightsimulator.com/html/Introduction/Introduction.htm)