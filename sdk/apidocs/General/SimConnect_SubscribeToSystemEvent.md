SimConnect\_SubscribeToSystemEvent

## SimConnect\_SubscribeToSystemEvent

The **SimConnect\_SubscribeToSystemEvent** function is used to request that a specific system event is notified to the client.

##### Syntax

```cpp
HRESULT SimConnect_SubscribeToSystemEvent(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_EVENT_ID  EventID,
    const char*  SystemEventName
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _EventID_ | Specifies the ID of the client event. | Integer |
| _SystemEventName_ | The string name for the requested system event, which should be one from the table below (note that the event names are not case-sensitive). Unless otherwise stated in the Description, notifications of the event are returned in a **[SIMCONNECT\_RECV\_EVENT](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT.htm)** structure (identify the event from the _EventID_ given with this function). | String |

The following table shows the available _SystemEventName_ variables:

| System Event Name | Description |
| --- | --- |
| `1sec` | Request a notification every second. |
| `4sec` | Request a notification every four seconds. |
| `6Hz` | Request notifications six times per second. This is the same rate that joystick movement events are transmitted. |
| `AircraftLoaded` | Request a notification when the aircraft flight dynamics file is changed. These files have a .AIR extension. The filename is returned in a [SIMCONNECT\_RECV\_EVENT\_FILENAME](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT_FILENAME.htm) structure. |
| `Crashed` | Request a notification if the user aircraft crashes. |
| `CrashReset` | Request a notification when the crash cut-scene has completed. |
| `CustomMissionActionExecuted` | Request a notification when a mission action has been executed. |
| `FlightLoaded` | Request a notification when a flight is loaded. Note that when a flight is ended, a default flight is typically loaded, so these events will occur when flights and missions are started and finished. The filename of the flight loaded is returned in a [SIMCONNECT\_RECV\_EVENT\_FILENAME](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT_FILENAME.htm) structure. |
| `FlightSaved` | Request a notification when a flight is saved correctly. The filename of the flight saved is returned in a [SIMCONNECT\_RECV\_EVENT\_FILENAME](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT_FILENAME.htm) structure. |
| `FlightPlanActivated` | Request a notification when a new flight plan is activated. The filename of the activated flight plan is returned in a [SIMCONNECT\_RECV\_EVENT\_FILENAME](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT_FILENAME.htm) structure. |
| `FlightPlanDeactivated` | Request a notification when the active flight plan is de-activated. |
| `Frame` | Request notifications every visual frame. Information is returned in a [SIMCONNECT\_RECV\_EVENT](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT.htm) structure. |
| `ObjectAdded` | Request a notification when an AI object is added to the simulation. Refer also to the [SIMCONNECT\_RECV\_EVENT\_OBJECT\_ADDREMOVE](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE.htm) structure. |
| `ObjectRemoved` | Request a notification when an AI object is removed from the simulation. Refer also to the [SIMCONNECT\_RECV\_EVENT\_OBJECT\_ADDREMOVE](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE.htm) structure. |
| `Pause` | Request notifications when the flight is paused or unpaused, and also immediately returns the current pause state (1 = paused or 0 = unpaused). The state is returned in the dwData parameter. |
| `Pause_EX1` | Request notifications when the flight is paused or unpaused, and also immediately returns the current pause state with more detail than the regular `Pause` system event. The state is returned in the `dwData` parameter, and will be a combination (OR) of the following defines:<br>```cpp<br>#define PAUSE_STATE_FLAG_OFF                0    // No Pause<br>#define PAUSE_STATE_FLAG_PAUSE              1    // "full" Pause (sim + traffic + etc...)<br>#define PAUSE_STATE_FLAG_PAUSE_WITH_SOUND   2    // FSX Legacy Pause (not used anymore)<br>#define PAUSE_STATE_FLAG_ACTIVE_PAUSE       4    // Pause was activated using the "Active Pause" Button<br>#define PAUSE_STATE_FLAG_SIM_PAUSE          8    // Pause the player sim but traffic, multi, etc... will still run<br>``` |
| `Paused` | Request a notification when the flight is paused. |
| `PauseFrame` | Request notifications for every visual frame that the simulation is paused. Information is returned in a [SIMCONNECT\_RECV\_EVENT](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT.htm) structure. |
| `PositionChanged` | Request a notification when the user changes the position of their aircraft through a dialog. |
| `Sim` | Request notifications when the flight is running or not, and also immediately returns the current state (1 = running or 0 = not running). The state is returned in the dwData parameter. |
| `SimStart` | The simulator is running. Typically the user is actively controlling the aircraft on the ground or in the air. However, in some cases additional pairs of `SimStart`/`SimStop` events are sent. For example, when a flight is reset the events that are sent are `SimStop`, `SimStart`, `SimStop`, `SimStart`. Also when a flight is started with the `SHOW_OPENING_SCREEN` value set to zero, then an additional `SimStart`/`SimStop` pair are sent before a second `SimStart` event is sent when the scenery is fully loaded. The opening screen provides the options to change aircraft, departure airport, and so on. |
| `SimStop` | The simulator is not running. Typically the user is loading a flight, navigating the shell or in a dialog. |
| `Sound` | Requests a notification when the master sound switch is changed. This request will also return the current state of the master sound switch immediately. A flag is returned in the `dwData` parameter, 0 if the switch is off, `SIMCONNECT_SOUND_SYSTEM_EVENT_DATA_MASTER` (0x1) if the switch is on. |
| `Unpaused` | Request a notification when the flight is un-paused. |
| `View` | Requests a notification when the user aircraft view is changed. This request will also return the current view immediately. A flag is returned in the dwData parameter, one of: `SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_2D``SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_VIRTUAL``SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_ORTHOGONAL` (the map view). |
| `WeatherModeChanged` | Request a notification when the weather mode is changed. |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
static enum EVENT_ID {
    EVENT_FLIGHT_LOAD,
    EVENT_RECUR_1SEC,
    EVENT_RECUR_FRAME,
    };
hr = SimConnect_SubscribeToSystemEvent(hSimConnect, EVENT_FLIGHT_LOAD, "FlightLoaded");
hr = SimConnect_SubscribeToSystemEvent(hSimConnect, EVENT_RECUR_1SEC,  "1sec");
hr = SimConnect_SubscribeToSystemEvent(hSimConnect, EVENT_RECUR_FRAME, "frame");
hr = SimConnect_SetSystemEventState(hSimConnect, EVENT_RECUR_FRAME, SIMCONNECT_STATE_OFF);
Working Samples
```

##### Remarks

A single call to this function is all that is necessary to receive the notifications. For greatest efficiency use `SimConnect_SetSystemEventState` to turn these requests on and off temporarily, and call [SimConnect\_UnsubscribeFromSystemEvent](SimConnect_UnsubscribeFromSystemEvent.htm) once only to permanently terminate the notifications of these events.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_RequestSystemState](../Events_And_Data/General/SimConnect_RequestSystemState.htm)
4. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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