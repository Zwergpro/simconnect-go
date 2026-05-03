SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENT\_PARAMS

## SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENT\_PARAMS

The **SIMCONNECT\_RECV\_SUBSCRIBE\_INPUT\_EVENT** structure is a response with the available parameters for an input event.

##### Syntax

```cpp
struct SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS: public SIMCONNECT_RECV
{
    UINT64 Hash;
    STRING Value;
};
```

##### Members

| Member | Description |
| `Hash` | Hash ID that identifies the input event. |
| `Value` | String that contains the values, separated by ;. Values can be:<br>- char\[260\]<br>- FLOAT64 |

##### Remarks

This struct will be the response to a call to [`SimConnect_EnumerateInputEventParams`](../InputEvents/SimConnect_EnumerateInputEventParams.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_EnumerateInputEventParams](../InputEvents/SimConnect_EnumerateInputEventParams.htm)
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