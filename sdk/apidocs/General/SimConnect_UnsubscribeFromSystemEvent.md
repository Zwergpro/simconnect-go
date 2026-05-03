SimConnect\_UnsubscribeFromSystemEvent

## SimConnect\_UnsubscribeFromSystemEvent

The **SimConnect\_UnsubscribeFromSystemEvent** function is used to request that notifications are no longer received for the specified system event.

##### Syntax

```cpp
HRESULT SimConnect_UnsubscribeFromSystemEvent(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_EVENT_ID  EventID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _EventID_ | Specifies the ID of the client event. | Integer |

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
hr = SimConnect_SubscribeToSystemEvent(hSimConnect, EVENT_RECUR_1SEC);
hr = SimConnect_UnsubscribeFromSystemEvent(hSimConnect, EVENT_RECUR_1SEC);
```

##### Remarks

There is no limit to the number of system events that can be subscribed to, but use this function to improve performance when a system event notification is no longer needed.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_SubscribeToSystemEvent](../Events_And_Data/SimConnect_SubscribeToSystemEvent.htm)
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