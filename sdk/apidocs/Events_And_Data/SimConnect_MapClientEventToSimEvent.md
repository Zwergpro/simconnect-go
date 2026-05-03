SimConnect\_MapClientEventToSimEvent

## SimConnect\_MapClientEventToSimEvent

The **SimConnect\_MapClientEventToSimEvent** function associates a client defined event ID with a Microsoft Flight Simulator 2024 event name.

##### Syntax

```cpp
HRESULT SimConnect_MapClientEventToSimEvent(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_EVENT_ID  EventID,
    const char*  EventName
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _EventID_ | Specifies the ID of the client event. | Integer |
| _EventName_ | Specifies the Microsoft Flight Simulator 2024 event name. Refer to the [Event IDs](../../../Key_Events/Key_Events.htm) document for a list of event names (listed under String Name). If the event name includes one or more periods (such as "Custom.Event" in the example below) then they are custom events specified by the client, and will only be recognized by another client (and not Microsoft Flight Simulator 2024) that has been coded to receive such events. No Microsoft Flight Simulator 2024 events include periods. If no entry is made for this parameter, the event is private to the client.<br>Alternatively enter a decimal number in the format "#nnnn" or a hex number in the format "#0xnnnn", where these numbers are in the range THIRD\_PARTY\_EVENT\_ID\_MIN and THIRD\_PARTY\_EVENT\_ID\_MAX, in order to receive events from third-party add-ons to Flight Simulator X. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
static enum EVENT_ID {
    EVENT_PAUSE,
    EVENT_BRAKES,
    EVENT_CUSTOM,
    EVENT_PRIVATE,
    };
hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_BRAKES, "brakes");
hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_PAUSE, "pause_toggle");
hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_CUSTOM, "Custom.Event");
hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_PRIVATE);
```

##### Remarks

It is important to call this function sufficiently frequently that the queue of information received from the server is processed. If there are no messages in the queue, the **\[dwID\]** parameter will be set to [`SIMCONNECT_RECV_ID_NULL`](../Structures_And_Enumerations/SIMCONNECT_RECV_ID.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_TransmitClientEvent](SimConnect_TransmitClientEvent.htm)
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