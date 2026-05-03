SIMCONNECT\_CAMERA\_AVAILABILITY

## SIMCONNECT\_CAMERA\_AVAILABILITY

The **SIMCONNECT\_CAMERA\_AVAILABILITY** enum is used to define the availability of the currently acquired camera.

##### Syntax

```cpp
enum SIMCONNECT_CAMERA_AVAILABILITY
{
    SIMCONNECT_CAMERA_NOT_ACQUIRED,
    SIMCONNECT_CAMERA_ACQUIRED,
    SIMCONNECT_CAMERA_ACQUIRED_BY_OTHER,
    SIMCONNECT_CAMERA_USER_DISABLED
};
```

##### Members

| Member | Description |
| `SIMCONNECT_CAMERA_NOT_ACQUIRED` | There is no add-on camera acquired. |
| `SIMCONNECT_CAMERA_ACQUIRED` | The add-on camera has been acquired. |
| `SIMCONNECT_CAMERA_ACQUIRED_BY_OTHER` | Some other client has acquired the add-on camera. |
| `SIMCONNECT_CAMERA_USER_DISABLED` | The add-on camera has been disabled by the user in the simulation options. |

##### Remarks

This enum will be used by the `SimConnect_CameraGetStatus` and other camera functions.

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