SIMCONNECT\_RECV\_CAMERA\_DATA

## SIMCONNECT\_RECV\_CAMERA\_DATA

The **SIMCONNECT\_RECV\_CAMERA\_DATA** structure is a response with the different data related to the current camera.

##### Syntax

```cpp
struct SIMCONNECT_RECV_CAMERA_DATA: public SIMCONNECT_RECV
{
    SIMCONNECT_DATA_CAMERA CameraData;
};
```

##### Members

| Member | Description |
| `CameraData` | Information about the camera, contained in the `SIMCONNECT_DATA_CAMERA` struct. |

##### Remarks

This struct will be the response to a call to `SimConnect_CameraGet`.

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