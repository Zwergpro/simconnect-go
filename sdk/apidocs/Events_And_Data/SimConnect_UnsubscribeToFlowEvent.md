SimConnect\_UnsubscribeToFlowEvent

## SimConnect\_UnsubscribeToFlowEvent

The **SimConnect\_UnsubscribeToFlowEvent** function is used to request that notifications are no longer received for the flow event.

- This API is available through Wasm, Javascript and through SimConnect. Even if SimConnect is accessible in Wasm, we strongly advise you to use the following functions than the ones provided by SimConnect.

You can find a sample project to use as a reference when using the Flow API here:

- [FlowAircraft](../../../../7_Samples_Tutorials/Samples/SimObjects_Aircraft/ModularAircraft/FlowAircraft.htm)

##### Syntax

```cpp
HRESULT SimConnect_UnsubscribeToFlowEvent(
    HANDLE  hSimConnect,
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
3. [SimConnect\_SubscribeToSystemEvent](SimConnect_SubscribeToSystemEvent.htm)
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