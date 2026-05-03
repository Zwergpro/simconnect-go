SIMCONNECT\_RECV\_ID

## SIMCONNECT\_RECV\_ID

The **SIMCONNECT\_RECV\_ID** enumeration type is used within the SIMCONNECT\_RECV structure to indicate which type of structure has been returned.

##### Syntax

```cpp
enum SIMCONNECT_RECV_ID{
    SIMCONNECT_RECV_ID_NULL,
    SIMCONNECT_RECV_ID_EXCEPTION,
    SIMCONNECT_RECV_ID_OPEN,
    SIMCONNECT_RECV_ID_QUIT,
    SIMCONNECT_RECV_ID_EVENT,
    SIMCONNECT_RECV_ID_EVENT_OBJECT_ADDREMOVE,
    SIMCONNECT_RECV_ID_EVENT_FILENAME,
    SIMCONNECT_RECV_ID_EVENT_FRAME,
    SIMCONNECT_RECV_ID_SIMOBJECT_DATA,
    SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE,
    SIMCONNECT_RECV_ID_WEATHER_OBSERVATION,
    SIMCONNECT_RECV_ID_CLOUD_STATE,
    SIMCONNECT_RECV_ID_ASSIGNED_OBJECT_ID,
    SIMCONNECT_RECV_ID_RESERVED_KEY,
    SIMCONNECT_RECV_ID_CUSTOM_ACTION,
    SIMCONNECT_RECV_ID_SYSTEM_STATE,
    SIMCONNECT_RECV_ID_CLIENT_DATA,
    SIMCONNECT_RECV_ID_EVENT_WEATHER_MODE,
    SIMCONNECT_RECV_ID_AIRPORT_LIST,
    SIMCONNECT_RECV_ID_VOR_LIST,
    SIMCONNECT_RECV_ID_NDB_LIST,
    SIMCONNECT_RECV_ID_WAYPOINT_LIST,
    SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_SERVER_STARTED,
    SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_CLIENT_STARTED,
    SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_SESSION_ENDED,
    SIMCONNECT_RECV_ID_EVENT_RACE_END,
    SIMCONNECT_RECV_ID_EVENT_RACE_LAP,
    SIMCONNECT_RECV_ID_PICK,
    SIMCONNECT_RECV_ID_EVENT_EX1,
    SIMCONNECT_RECV_ID_FACILITY_DATA,
    SIMCONNECT_RECV_ID_FACILITY_DATA_END,
    SIMCONNECT_RECV_ID_FACILITY_MINIMAL_LIST,
    SIMCONNECT_RECV_ID_JETWAY_DATA,
    SIMCONNECT_RECV_ID_CONTROLLERS_LIST,
    SIMCONNECT_RECV_ID_ACTION_CALLBACK,
    SIMCONNECT_RECV_ID_ENUMERATE_INPUT_EVENTS,
    SIMCONNECT_RECV_ID_GET_INPUT_EVENT,
    SIMCONNECT_RECV_ID_SUBSCRIBE_INPUT_EVENT,
    SIMCONNECT_RECV_ID_ENUMERATE_INPUT_EVENT_PARAMS,
    SIMCONNECT_RECV_ID_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST,
    SIMCONNECT_RECV_ID_FLOW_EVENT,
    SIMCONNECT_RECV_ID_CAMERA_DATA,
    SIMCONNECT_RECV_ID_CAMERA_STATUS,
    SIMCONNECT_RECV_ID_CAMERA_DEFINITION_LIST,
    SIMCONNECT_RECV_ID_COMM_BUS,
    SIMCONNECT_RECV_ID_CAMERA_WORLD_LOCKER,
    };
```

##### Members

| Member | Description |
| `SIMCONNECT_RECV_ID_NULL` | Specifies that nothing useful has been returned. |
| `SIMCONNECT_RECV_ID_EXCEPTION` | Specifies that a `SIMCONNECT_RECV_EXCEPTION` structure has been received. |
| `SIMCONNECT_RECV_ID_OPEN` | Specifies that a `SIMCONNECT_RECV_OPEN` structure has been received. |
| `SIMCONNECT_RECV_ID_QUIT` | Specifies that the user has exited from Microsoft Flight Simulator 2024. |
| `SIMCONNECT_RECV_ID_EVENT` | Specifies that a `SIMCONNECT_RECV_EVENT` structure has been received. |
| `SIMCONNECT_RECV_ID_EVENT_EX1` | Specifies that the `SIMCONNECT_RECV_EVENT_EX1` structure has been received. Can be triggered by various functions, eg: `trigger_key_event_EX1` or [`SimConnect_TransmitClientEvent_EX1`](../Events_And_Data/SimConnect_TransmitClientEvent_EX1.htm). |
| `SIMCONNECT_RECV_ID_EVENT_OBJECT_ADDREMOVE` | Specifies that a `SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE` structure has been received. |
| `SIMCONNECT_RECV_ID_EVENT_FILENAME` | Specifies that a `SIMCONNECT_RECV_EVENT_FILENAME` structure has been received. |
| `SIMCONNECT_RECV_ID_EVENT_FRAME` | Specifies that a `SIMCONNECT_RECV_EVENT_FRAME` structure has been received. |
| `SIMCONNECT_RECV_ID_SIMOBJECT_DATA` | Specifies that a `SIMCONNECT_RECV_SIMOBJECT_DATA` structure has been received. |
| `SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE` | Specifies that a `SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE` structure has been received. |
| `SIMCONNECT_RECV_ID_CLOUD_STATE` | Specifies that a `SIMCONNECT_RECV_CLOUD_STATE` structure has been received. |
| `SIMCONNECT_RECV_ID_WEATHER_OBSERVATION` | Specifies that a `SIMCONNECT_RECV_WEATHER_OBSERVATION` structure has been received. |
| `SIMCONNECT_RECV_ID_EVENT_WEATHER_MODE` | Specifies that the `dwData` parameter will contain one value of the `SIMCONNECT_WEATHER_MODE` enumeration. |
| `SIMCONNECT_RECV_ID_ASSIGNED_OJBECT_ID` | Specifies that a `SIMCONNECT_RECV_ASSIGNED_OBJECT_ID` structure has been received. |
| `SIMCONNECT_RECV_ID_RESERVED_KEY` | Specifies that a `SIMCONNECT_RECV_RESERVED_KEY` structure has been received. |
| `SIMCONNECT_RECV_ID_CUSTOM_ACTION` | Specifies that a `SIMCONNECT_RECV_CUSTOM_ACTION` structure has been received. |
| `SIMCONNECT_RECV_ID_SYSTEM_STATE` | Specifies that a `SIMCONNECT_RECV_SYSTEM_STATE` structure has been received. |
| `SIMCONNECT_RECV_ID_CLIENT_DATA` | Specifies that a `SIMCONNECT_RECV_CLIENT_DATA` structure has been received. |
| `SIMCONNECT_RECV_ID_AIRPORT_LIST` | Specifies that a `SIMCONNECT_RECV_AIRPORT_LIST` structure has been received. |
| `SIMCONNECT_RECV_ID_VOR_LIST` | Specifies that a `SIMCONNECT_RECV_VOR_LIST` structure has been received. |
| `SIMCONNECT_RECV_ID_NDB_LIST` | Specifies that a `SIMCONNECT_RECV_NDB_LIST` structure has been received. |
| `SIMCONNECT_RECV_ID_WAYPOINT_LIST` | Specifies that a `SIMCONNECT_RECV_WAYPOINT_LIST` structure has been received. |
| `SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_SERVER_STARTED` | Specifies that a `SIMCONNECT_RECV_EVENT_MULTIPLAYER_SERVER_STARTED` structure has been received. |
| `SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_CLIENT_STARTED` | Specifies that a `SIMCONNECT_RECV_EVENT_MULTIPLAYER_CLIENT_STARTED` structure has been received. |
| `SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_SESSION_ENDED` | Specifies that a `SIMCONNECT_RECV_EVENT_MULTIPLAYER_SESSION_ENDED` structure has been received. |
| `SIMCONNECT_RECV_ID_EVENT_RACE_END` | Specifies that a `SIMCONNECT_RECV_EVENT_RACE_END` structure has been received. |
| `SIMCONNECT_RECV_ID_EVENT_RACE_LAP` | Specifies that a `SIMCONNECT_RECV_EVENT_RACE_LAP` structure has been received. |
| `SIMCONNECT_RECV_ID_FACILITY_DATA` | Specifies that a `SIMCONNECT_RECV_FACILITY_DATA` structure has been received |
| `SIMCONNECT_RECV_ID_FACILITY_DATA_END` | Specifies that a `SIMCONNECT_RECV_FACILITY_DATA_END` structure has been received. |
| `SIMCONNECT_RECV_ID_FACILITY_MINIMAL_LIST` | Specifies that a `SIMCONNECT_RECV_FACILITIES_MINIMAL` structure has been received. |
| `SIMCONNECT_RECV_ID_JETWAY_DATA` | Specifies that a `SIMCONNECT_RECV_JETWAY_DATA` structure has been received. |
| `SIMCONNECT_RECV_ID_CONTROLLERS_LIST` | Specifies that the callback has been created by the `SimConnect_EnumerateControllers` function. |
| `SIMCONNECT_RECV_ID_ACTION_CALLBACK` | Specifies that the callback has been created by the `SimConnect_ExecuteAction` function. |
| `SIMCONNECT_RECV_ID_ENUMERATE_INPUT_EVENTS` | Specifies that a `SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS` structure has been recieved. |
| `SIMCONNECT_RECV_ID_GET_INPUT_EVENT` | Specifies that a `SIMCONNECT_RECV_GET_INPUT_EVENT` structure has been recieved. |
| `SIMCONNECT_RECV_ID_SUBSCRIBE_INPUT_EVENT` | Specifies that an input event has been subscribed to using the `SimConnect_SubscribeInputEvent` function. |
| `SIMCONNECT_RECV_ID_ENUMERATE_INPUT_EVENT_PARAMS` | Specifies that a `SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS` structure has been recieved. |
| `SIMCONNECT_RECV_ID_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST` | Specifies that a `SIMCONNECT_RECV_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST` structure has been recieved. |
| `SIMCONNECT_RECV_ID_FLOW_EVENT` | Specifies that an input event has been subscribed to using the `SimConnect_SubscribeInputEvent` function. |
| `SIMCONNECT_RECV_ID_CAMERA_DATA` | Specifies that a `SIMCONNECT_RECV_CAMERA_DATA` structure has been received. |
| `SIMCONNECT_RECV_ID_CAMERA_STATUS` | Specifies that a `SIMCONNECT_RECV_CAMERA_STATUS` structure has been received. |
| `SIMCONNECT_RECV_ID_CAMERA_DEFINITION_LIST` | Specifies that a `SIMCONNECT_RECV_CAMERA_DEFINITION_LIST` structure has been received. |
| `SIMCONNECT_RECV_ID_COMM_BUS` | Specifies that a communications event which has been subscribed to by `SimConnect_SubscribeToCommBusEvent` has been recieved. Details of the communication are contained in the `SIMCONNECT_RECV_COMM_BUS` struct. |
| `SIMCONNECT_RECV_ID_CAMERA_WORLD_LOCKER` | Specifies that a `SIMCONNECT_RECV_CAMERA_WORLD_LOCKER` structure has been received. |

##### Remarks

In the context of SimConnect, exceptions are error codes, and should not be confused with the C# or system concepts of exceptions.

Refer to the remarks for [SimConnect\_GetLastSentPacketID](../Debug/SimConnect_GetLastSentPacketID.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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