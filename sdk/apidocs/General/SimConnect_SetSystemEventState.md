SimConnect\_SetSystemEventState

## SimConnect\_SetSystemEventState

The **SimConnect\_SetSystemEventState** function is used to turn requests for event information from the server on and off.

##### Syntax

```cpp
HRESULT SimConnect_SetSystemEventState(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_EVENT_ID  EventID,
    SIMCONNECT_STATE  dwState
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _EventID_ | Specifies the ID of the client event that is to have its state changed. | Integer |
| _dwState_ | Double word containing the state (one member of [SIMCONNECT\_STATE](../Structures_And_Enumerations/SIMCONNECT_STATE.htm)). | Integer |

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
```

##### Remarks

If this function is not called, the default is for the state to be on. This is different from input events, which have a default state of off.

Use this function to turn system events temporarily on and off, rather than make multiple calls to [SimConnect\_SubscribeToSystemEvent](SimConnect_SubscribeToSystemEvent.htm) and [SimConnect\_UnsubscribeFromSystemEvent](SimConnect_UnsubscribeFromSystemEvent.htm), which is less efficient.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_MapClientEventToSimEvent](../Events_And_Data/SimConnect_MapClientEventToSimEvent.htm)
4. [SimConnect\_MapInputEventToClientEvent](../Events_And_Data/SimConnect_MapInputEventToClientEvent.htm)
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