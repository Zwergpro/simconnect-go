SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENTS

## SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENTS

The **SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENTS** structure is used to return a single page of data about an input event.

##### Syntax

```cpp
struct SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS : public SIMCONNECT_RECV_LIST_TEMPLATE {
    SIMCONNECT_INPUT_EVENT_DESCRIPTOR rgData[dwArraySize];
};
```

##### Members

| Member | Description |
| `rgData[dwArraySize]` | Array of [`SIMCONNECT_INPUT_EVENT_DESCRIPTOR`](SIMCONNECT_INPUT_EVENT_DESCRIPTOR.htm) structures. |

This function might throw :

- SIMCONNECT\_EXCEPTION\_ERROR in the case of internal errors (see [SIMCONNECT\_EXCEPTION](SIMCONNECT_EXCEPTION.htm))

##### Remarks

One or more of these structs is returned as the response when you call the [`SimConnect_EnumerateInputEvents`](../InputEvents/SimConnect_EnumerateInputEvents.htm) function, where each struct returned represents a page of data. As part of the struct for a page, there will be an rgData member, which is an array of data comprised of 1 or more [`SIMCONNECT_INPUT_EVENT_DESCRIPTOR`](SIMCONNECT_INPUT_EVENT_DESCRIPTOR.htm) items. Note that this struct inherits members from the [`SIMCONNECT_RECV_LIST_TEMPLATE`](SIMCONNECT_RECV_LIST_TEMPLATE.htm) struct.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_EnumerateInputEvents](../InputEvents/SimConnect_EnumerateInputEvents.htm)
4. [SIMCONNECT\_INPUT\_EVENT\_DESCRIPTOR](SIMCONNECT_INPUT_EVENT_DESCRIPTOR.htm)
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