SIMCONNECT\_RECV\_CAMERA\_STATUS

## SIMCONNECT\_RECV\_CAMERA\_STATUS

The **SIMCONNECT\_RECV\_CAMERA\_STATUS** structure is a response with the status of the add-on camera.

##### Syntax

```cpp
struct SIMCONNECT_RECV_CAMERA_STATUS: public SIMCONNECT_RECV
{
    DWORD acquiredState;
    BOOL  bGameControlled;
};
```

##### Members

| Member | Description |
| `acquiredState` | The current acquired state of the camera, which will be one of the following:<br>1. `NOT_ACQUIRED` \- The add-on camera is free to be acquired.<br>2. `ACQUIRED` \- The add-on camera is owned by the client and can be set.<br>3. `ACQUIRED_BY_OTHER` \- The add-on camera is owned by another client.<br>4. `USER_DISABLED` \- The camera API is disabled by the user in the simulation options, and cameras cannot be acquired. |
| `bGameControlled` | Returns whether the camera is currently under control by the simulation (TRUE) or not (FALSE). |

##### Remarks

This struct will be the response to a call to `SimConnect_CameraGetStatus`.

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