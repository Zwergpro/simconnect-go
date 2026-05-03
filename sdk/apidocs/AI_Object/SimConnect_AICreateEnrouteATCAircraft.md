SimConnect\_AICreateEnrouteATCAircraft

## SimConnect\_AICreateEnrouteATCAircraft

The **SimConnect\_AICreateEnrouteATCAircraft** function is used to create an AI controlled aircraft that is about to start or is already underway on its flight plan.

**NOTE**: This function is a **legacy** function and only works with **non-modular aircraft**. In Microsoft Flight Simulator 2024 we recommend you use `SimConnect_AICreateEnrouteATCAircraft_EX1`, which can be used with legacy **and** modular aircraft.

##### Syntax

```cpp
HRESULT SimConnect_AICreateEnrouteATCAircraft(
    HANDLE  hSimConnect,
    const char*  szContainerTitle,
    const char*  szTailNumber,
    int  iFlightNumber,
    const char*  szFlightPlanPath,
    double  dFlightPlanPosition,
    BOOL  bTouchAndGo,
    SIMCONNECT_DATA_REQUEST_ID  RequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _szContainerTitle_ | Null-terminated string containing the container title. The container title is found in the [aircraft.cfg](../../../../5_Content_Configuration/CFG_Files/ai_cfg.htm) file. Alternatively, the aircraft title can be obtained via the Aircraft Selector ( _DevMode_-> _Windows_-> _Aircraft_ selector). Finally, the information can be found using the `SimConnect_EnumerateSimObjectsAndLiveries` function.<br>Examples of aircraft titles:<br>`title=Boeing 747-8f Asobo`<br>`title=DA62 Asobo`<br>`title=VL3 Asobo` | String |
| _szTailNumber_ | Null-terminated string containing the tail number. This should have a maximum of 12 characters. | String |
| _iFlightNumber_ | Integer containing the flight number. There is no specific maximum length of this number. Any negative number indicates that there is no flight number.<br>_szFlightPlanPath_. | Integer |
| _szFlightPlanPath_ | Null-terminated string containing the path to the flight plan file. Flight plans have the extension .pln, but no need to enter an extension here. The easiest way to create flight plans is to create them from within Microsoft Flight Simulator 2024 itself, and then save them off for use with the AI controlled aircraft. | String |
| _dFlightPlanPosition_ | Double floating point number containing the flight plan position. The number before the point contains the waypoint index, and the number afterwards how far along the route to the next waypoint the aircraft is to be positioned. The first waypoint index is 0. For example, 0.0 indicates that the aircraft has not started on the flight plan, 2.5 would indicate the aircraft is to be initialized halfway between the third and fourth waypoints (which would have indexes 2 and 3). The waypoints are those recorded in the flight plan, which may just be two airports, and do not include any taxiway points on the ground. Also there is a threshold that will ignore requests to have an aircraft taxiing or taking off, or landing. So set the value after the point to ensure the aircraft will be in level flight. | Float |
| _bTouchAndGo_ | Set to True to indicate that landings should be touch and go, and not full stop landings. | Bool |
| _RequestID_ | Specifies the client defined request ID. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

An enroute aircraft can be on the ground or airborne when it is created by this function. Typically this will be an aircraft flying under IFR rules, and in constant radio contact with ATC. The following exceptions can be created by this function (refer to the [SIMCONNECT\_EXCEPTION](../Structures_And_Enumerations/SIMCONNECT_EXCEPTION.htm) enum for more details):

- `SIMCONNECT_EXCEPTION_LOAD_FLIGHTPLAN_FAILED`
- `SIMCONNECT_EXCEPTION_CREATE_OBJECT_FAILED`

A [SIMCONNECT\_RECV\_ID\_EVENT\_OBJECT\_ADDREMOVE](../Structures_And_Enumerations/SIMCONNECT_RECV_ID.htm) event notification can be subscribed to (see the [SimConnect\_SubscribeToSystemEvent](../General/SimConnect_SubscribeToSystemEvent.htm) function), which will return a [SIMCONNECT\_RECV\_EVENT\_OBJECT\_ADDREMOVE](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE.htm) structure whenever any client, including the one making the change, successfully adds or removes an AI controlled object.

For float-planes the recommended procedure is to control them using waypoints, and not the ATC system, as there is no concept of a "parking space" after a water landing. So, the waypoints of the route of the float-plane should include the route that it should follow before take off and after landing. For all these cases of controlling aircraft using the client, or using waypoints, set up the object using the [SimConnect\_AICreateNonATCAircraft](SimConnect_AICreateNonATCAircraft.htm) call.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AICreateNonATCAircraft](SimConnect_AICreateNonATCAircraft.htm)
4. [SimConnect\_AICreateParkedATCAircraft](SimConnect_AICreateParkedATCAircraft.htm)
5. [SimConnect\_AISetAircraftFlightPlan](SimConnect_AISetAircraftFlightPlan.htm)
6. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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