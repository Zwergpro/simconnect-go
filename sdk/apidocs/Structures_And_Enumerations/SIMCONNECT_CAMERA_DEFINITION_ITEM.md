SIMCONNECT\_CAMERA\_DEFINITION\_ITEM

## SIMCONNECT\_CAMERA\_DEFINITION\_ITEM

The **SIMCONNECT\_CAMERA\_DEFINITION\_ITEM** structure is used to hold a single string which is the unique ID string identifying a camera.

##### Syntax

```cpp
struct SIMCONNECT_CAMERA_DEFINITION_ITEM
{
    SIMCONNECT_STRING(Str, 256);
};
```

##### Members

| Member | Description |
| `str` | A camera ID string. |

##### Remarks

This struct is used as part of the `SIMCONNECT_RECV_CAMERA_DEFINITION_LIST`.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)
4. [Camera](../Camera/Camera.htm)

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