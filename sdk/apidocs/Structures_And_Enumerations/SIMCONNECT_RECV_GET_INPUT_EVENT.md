SIMCONNECT\_RECV\_GET\_INPUT\_EVENT

## SIMCONNECT\_RECV\_GET\_INPUT\_EVENT

The **SIMCONNECT\_RECV\_GET\_INPUT\_EVENT** structure is used to return the value of a specific input event.

##### Syntax

```cpp
struct SIMCONNECT_RECV_GET_INPUT_EVENT : public SIMCONNECT_RECV
{
    DWORD   RequestID;
    SIMCONNECT_INPUT_EVENT_TYPE Type;
    PVOID   Value;
};
```

##### Members

| Member | Description |
| `RequestID` | Specifies the client defined request ID. |
| `Type` | One member of the `` enumeration type. This is used to cast the `Value` to the correct type. |
| `Value` | The value of the requested input event, which should be cast to the correct format (float / string). |

##### Remarks

This struct will be the response to a call to [`SimConnect_GetInputEvent`](../InputEvents/SimConnect_GetInputEvent.htm). The value stored in this struct must be cast on the client side in the correct format (as a float or a string, only).

**NOTE**: In C#, `Value` is an object array containing one uint value. This value must be converted to a byte array before being converted to the correct type.

If it is a string, it can have a maximum length of only 32 characters (including \\0). Note that if the unit is different to the expected in-simulation unit, it must be converted to the appropriate unit on the clients side.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_GetInputEvent](../InputEvents/SimConnect_GetInputEvent.htm)
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