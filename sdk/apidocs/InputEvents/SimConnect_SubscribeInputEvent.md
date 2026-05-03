SimConnect\_SubscribeInputEvent

## SimConnect\_SubscribeInputEvent

The **SimConnect\_SubscribeInputEvent** function is used to subscribe an input event and generate when the value changes.

##### Syntax

```cpp
HRESULT SimConnect_SubscribeInputEvent(
    HANDLE hSimConnect,
    UINT64 Hash,
)
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `Hash` | Hash ID that will identify the desired input event to subscribe to. You can use 0 here to subscribe to _all_ input events. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

This function might throw one of the following `SIMCONNECT_EXCEPTION`:

- `SIMCONNECT_EXCEPTION_ERROR` in case of internal errors
- `SIMCONNECT_EXCEPTION_GET_INPUT_EVENT_FAILED` if the given hash in wrong

##### Remarks

This function will subscribe to a specific input event based on the given hash (you can retrieve the input event hashes using [`SimConnect_EnumerateInputEvents`](SimConnect_EnumerateInputEvents.htm)). You may also supply 0 (null) as the hash value, in which case you will be subscribing to _all_ input events. The function will generate one or more [`SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT`](../Structures_And_Enumerations/SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT.htm) structs.

**NOTE**: If you are using C# please see the [Note for C#](SimConnect_GetInputEvent.htm#note).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_UnsubscribeInputEvent](SimConnect_UnsubscribeInputEvent.htm)
4. [SimConnect\_EnumerateInputEvents](SimConnect_EnumerateInputEvents.htm)
5. [SIMCONNECT\_RECV\_SUBSCRIBE\_INPUT\_EVENT](../Structures_And_Enumerations/SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT.htm)
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