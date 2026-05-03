SimConnect\_SubscribeToFlowEvent

## SimConnect\_SubscribeToFlowEvent

The **SimConnect\_SubscribeToFlowEvent** function is used to receive a [SIMCONNECT\_RECV\_FLOW\_EVENT](../Structures_And_Enumerations/SIMCONNECT_RECV_FLOW_EVENT.htm) message when the simulation triggers some flow events, for example during a [Back On Track](../../../../5_Content_Configuration/FLT_Files/Back_On_Track.htm) event.

- This API is available through [Wasm](../../../WASM/Flow_API/Flow_API.htm), [Javascript](../../../JavaScript/Flow_API/Flow_API.htm) and through [SimConnect](SimConnect_SubscribeToFlowEvent.htm).

You can find a sample project to use as a reference when using the Flow API here:

- [FlowAircraft](../../../../7_Samples_Tutorials/Samples/SimObjects_Aircraft/ModularAircraft/FlowAircraft.htm)

##### Syntax

```cpp
HRESULT SimConnect_SubscribeToFlowEvent(
    HANDLE  hSimConnect
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_RequestSystemState](General/SimConnect_RequestSystemState.htm)
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