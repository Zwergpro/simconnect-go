SimConnect\_SetInputEvent

## SimConnect\_SetInputEvent

The **SimConnect\_SetInputEvent** function is used to set the value of a specific input event (identified by its hash). See [SimConnect\_EnumerateInputEventParams](SimConnect_EnumerateInputEventParams.htm) for an example of use.

##### Syntax

```cpp
HRESULT SimConnect_SetInputEvent(
    HANDLE hSimConnect,
    DWORD Hash,
    DWORD cbUnitSize,
    PVOID Value
)
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `Hash` | Hash ID that will identify the desired inputEvent. | Integer |
| `cbUnitSize` | Specifies the size of the value in bytes. | Integer |
| `Value` | New value of the specified inputEvent. | Float/String |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

This function might throw one of the following `SIMCONNECT_EXCEPTION`:

- `SIMCONNECT_EXCEPTION_ERROR` in case of internal errors.
- `isSIMCONNECT_EXCEPTION_SET_INPUT_EVENT_FAILED` if the given hash in wrong.

##### Remarks

This function will set the value of an input event and generates no response event. You can get the hashes for the available input events using [`SimConnect_EnumerateInputEvents`](SimConnect_EnumerateInputEvents.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_EnumerateInputEvents](SimConnect_EnumerateInputEvents.htm)
4. [SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENTS](../Structures_And_Enumerations/SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS.htm)
5. [SIMCONNECT\_INPUT\_EVENT\_DESCRIPTOR](../Structures_And_Enumerations/SIMCONNECT_INPUT_EVENT_DESCRIPTOR.htm)
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