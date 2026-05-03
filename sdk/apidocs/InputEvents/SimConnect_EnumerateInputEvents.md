SimConnect\_EnumerateInputEvents

## SimConnect\_EnumerateInputEvents

The **SimConnect\_EnumerateInputEvents** function is used to retrieve a paginated list of all available InputEvents for the current aircraft along with their associated hash (CRC based).

##### Syntax

```cpp
HRESULT SimConnect_EnumerateInputEvents(
    HANDLE hSimConnect,
    SIMCONNECT_DATA_REQUEST_ID RequestID
)
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `RequestID` | The ID that will identify the current request in the response event | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

With this function you can request data on the input events currently available for the current SimObject. The function will trigger one or more [`SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS`](../Structures_And_Enumerations/SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS.htm) responses, and in this response struct you will find the `Data` member, which has a list of [`SIMCONNECT_INPUT_EVENT_DESCRIPTOR`](../Structures_And_Enumerations/SIMCONNECT_INPUT_EVENT_DESCRIPTOR.htm) structs, where each struct represents an input event, for example:

```
{ INPUT_EVENT_NAME_1; hash_1 }
{ INPUT_EVENT_NAME_2; hash_2 }
{ INPUT_EVENT_NAME_3; hash_3 }
// etc...
```

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENTS](../Structures_And_Enumerations/SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS.htm)
4. [SIMCONNECT\_INPUT\_EVENT\_DESCRIPTOR](../Structures_And_Enumerations/SIMCONNECT_INPUT_EVENT_DESCRIPTOR.htm)
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