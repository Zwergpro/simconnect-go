SimConnect\_RequestSystemState

## SimConnect\_RequestSystemState

The **SimConnect\_RequestSystemState** function is used to request information from a number of Microsoft Flight Simulator system components.

##### Syntax

```cpp
HRESULT SimConnect_RequestSystemState(
    HANDLE  hSimConnect,
    SIMCONNECT_DATA_REQUEST_ID  RequestID,
    const char*  szState,
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _RequestID_ | The client defined request ID. | Integer |
| _szState_ | A null-terminated string identifying the system function. One from the table shown below. | String |

The following table shows the values available for the _szState_:

| String | Description |
| --- | --- |
| `AircraftLoaded` | Requests the full path name of the last loaded aircraft flight dynamics file. These files have a .AIR extension. |
| `DialogMode` | Requests whether the simulation is in Dialog mode or not. |
| `FlightLoaded` | Requests the full path name of the last loaded flight. Flight files have the extension .FLT. |
| `FlightPlan` | Requests the full path name of the active flight plan. An empty string will be returned if there is no active flight plan. |
| `Sim` | Requests the state of the simulation. If 1 is returned, the user is in control of the aircraft, if 0 is returned, the user is navigating the UI. This is the same state that notifications can be subscribed to with the "SimStart" and "SimStop" string with the `SimConnect_SubscribeToSystemEvent` function. |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

It is important to call this function sufficiently frequently that the queue of information received from the server is processed. If there are no messages in the queue, the **\[dwID\]** parameter will be set to [`SIMCONNECT_RECV_ID_NULL`](../Structures_And_Enumerations/SIMCONNECT_RECV_ID.htm).

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