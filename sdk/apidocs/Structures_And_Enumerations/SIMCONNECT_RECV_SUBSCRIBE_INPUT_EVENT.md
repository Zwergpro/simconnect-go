SIMCONNECT\_RECV\_SUBSCRIBE\_INPUT\_EVENT

## SIMCONNECT\_RECV\_SUBSCRIBE\_INPUT\_EVENT

The **SIMCONNECT\_RECV\_SUBSCRIBE\_INPUT\_EVENT** structure is used to return the value of a subscribed input event.

##### Syntax

```cpp
struct SIMCONNECT_RECV_SUBSCRIBE_INPUT_EVENT: public SIMCONNECT_RECV
{
    UINT64 Hash;
    SIMCONNECT_INPUT_EVENT_TYPE Type;
    PVOID Value;
};
```

##### Members

| Member | Description |
| `Hash` | Hash ID that will identify the subscribed input event. |
| `Type` | One member of the [`SIMCONNECT_INPUT_EVENT_TYPE`](SIMCONNECT_INPUT_EVENT_TYPE.htm) enumeration type. This is used to cast the `Value` to the correct type. |
| `Value` | The value of the subscribed input event, which should be cast to the correct format (float / string). |

##### Remarks

This struct will be the response to a call to [`SimConnect_SubscribeInputEvent`](../InputEvents/SimConnect_SubscribeInputEvent.htm). The value stored in this struct must be cast on the client side in the correct format (as a float or a string, only). If it is a string, it can have a maximum length of 256 characters (including \\0).

**NOTE**: In C#, the Value is an object array containing one `uint` value. For details on how this should be handled please see this [Note for C#](../InputEvents/SimConnect_GetInputEvent.htm#note).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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