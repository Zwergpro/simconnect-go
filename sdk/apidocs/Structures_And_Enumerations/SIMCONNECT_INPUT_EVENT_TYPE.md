SIMCONNECT\_INPUT\_EVENT\_TYPE

## SIMCONNECT\_INPUT\_EVENT\_TYPE

The **SIMCONNECT\_INPUT\_EVENT\_TYPE** enumeration type is used with the [SimConnect\_GetInputEvent](../InputEvents/SimConnect_GetInputEvent.htm) call to specify the data type used and help you cast the return value correctly.

##### Syntax

```cpp
enum SIMCONNECT_INPUT_EVENT_TYPE : DWORD{
    SIMCONNECT_INPUT_EVENT_TYPE_NONE,
    SIMCONNECT_INPUT_EVENT_TYPE_DOUBLE,
    SIMCONNECT_INPUT_EVENT_TYPE_STRING
};
```

##### Members

| Member | Description |
| `SIMCONNECT_INPUT_EVENT_TYPE_NONE` | No data type specification required. |
| `SIMCONNECT_INPUT_EVENT_TYPE_DOUBLE` | Specifies a double. |
| `SIMCONNECT_INPUT_EVENT_TYPE_STRING` | Specifies a string. |

##### Remarks

N/A

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_EnumerateInputEvents](../InputEvents/SimConnect_EnumerateInputEvents.htm)
4. [SimConnect\_GetInputEvent](../InputEvents/SimConnect_GetInputEvent.htm)
5. [SimConnect\_SubscribeInputEvent](../InputEvents/SimConnect_SubscribeInputEvent.htm)
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