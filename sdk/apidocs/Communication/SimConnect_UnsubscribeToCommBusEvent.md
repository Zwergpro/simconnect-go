SimConnect\_UnsubscribeToCommBusEvent

## SimConnect\_UnsubscribeToCommBusEvent

The **SimConnect\_UnsubscribeToCommBusEvent** function is used to unsubscribe the client from a communication (CommBus) event.

**C++****C#**

##### Syntax

```cpp
HRESULT SimConnect_UnsubscribeToCommBusEvent(
    HANDLE hSimConnect,
    SIMCONNECT_CLIENT_EVENT_ID EventID
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

##### Syntax

```cs
void SimConnect::UnsubscribeToCommBusEvent(
    Enum EventID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _EventID_ | The ID with which the event has been registered on the client. | Enum |

##### Return Values

N/A (use a try/catch to test for errors).


##### Remarks

N/A

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)
4. [Communication](Communication.htm)

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