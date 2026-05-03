SIMCONNECT\_INPUT\_EVENT\_DESCRIPTOR

## SIMCONNECT\_INPUT\_EVENT\_DESCRIPTOR

The **SIMCONNECT\_INPUT\_EVENT\_DESCRIPTOR** structure is used to return an item of data for a specific input event.

##### Syntax

```cpp
struct SIMCONNECT_INPUT_EVENT_DESCRIPTOR
{
    SIMCONNECT_STRING(Name, 64);
    DWORD Hash;
    SIMCONNECT_DATATYPE Type;
    SIMCONNECT_STRING(NodeNames, 1024);
};
```

##### Members

| Member | Description |
| `Name` | The name of the Input Event. |
| `Hash` | The hash ID for the event. |
| `Type` | The expected datatype (from the `SIMCONNECT_DATATYPE` enum). Usually a `FLOAT32` or `STRING128`. |
| `NodeNames` | A list of the names of the nodes linked to this InputEvent, where each node name is separated by a `;`. |

##### Remarks

This struct will be part of the data included with the [`SIMCONNECT_RECV_GET_INPUT_EVENT`](SIMCONNECT_RECV_GET_INPUT_EVENT.htm) struct, which is returned by calls to the [`SimConnect_EnumerateInputEvents`](../InputEvents/SimConnect_EnumerateInputEvents.htm) function. There may be multiple structs of this type returned as part of the requested input event data.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_EnumerateInputEvents](../InputEvents/SimConnect_EnumerateInputEvents.htm)
4. [SIMCONNECT\_RECV\_GET\_INPUT\_EVENT](SIMCONNECT_RECV_GET_INPUT_EVENT.htm)
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