SimConnect\_UnsubscribeInputEvent

## SimConnect\_UnsubscribeInputEvent

The **SimConnect\_UnsubscribeInputEvent** function is used to unsubscribe from an input event that has previously been subscribed to.

##### Syntax

```cpp
HRESULT SimConnect_UnsubscribeInputEvent(
    HANDLE hSimConnect,
    UINT64 Hash
)
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `Hash` | Hash ID that will identify the desired input event to unsubscribe from. You can use 0 here to unsubscribe from _all_ input events. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

This function will unsubscribe from a specific input event based on the given hash (you can retrieve the input event hashes using [`SimConnect_EnumerateInputEvents`](SimConnect_EnumerateInputEvents.htm)). You may also supply 0 (null) as the hash value, in which case you will be unsubscribing from _all_ input events. The function generates no response event.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_SubscribeInputEvent](SimConnect_SubscribeInputEvent.htm)
4. [SimConnect\_EnumerateInputEvents](SimConnect_EnumerateInputEvents.htm)
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