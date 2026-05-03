SIMCONNECT\_RECV\_CAMERA\_DEFINITION\_LIST

## SIMCONNECT\_RECV\_CAMERA\_DEFINITION\_LIST

The **SIMCONNECT\_RECV\_CAMERA\_DEFINITION\_LIST** struct is used to return an array containing the various names of the available cameras.

##### Syntax

```cpp
struct SIMCONNECT_RECV_CAMERA_DEFINITION_LIST : public SIMCONNECT_RECV_LIST_TEMPLATE
{
    SIMCONNECT_CAMERA_DEFINITION_ITEM[] rgData;
};
```

##### Members

| Member | Description |
| `rgData` | This is an array that will contain a number of `SIMCONNECT_CAMERA_DEFINITION_ITEM` structs. Each of these structs will hold a camare ID string. |

##### Remarks

This structure is sent when the `SimConnect_EnumerateCameraDefinitions` function is called.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [Camera](../Camera/Camera.htm)
4. [SIMCONNECT\_RECV](SIMCONNECT_RECV.htm)
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