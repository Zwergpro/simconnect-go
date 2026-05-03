SIMCONNECT\_RECV\_CAMERA\_WORLD\_LOCKER

## SIMCONNECT\_RECV\_CAMERA\_WORLD\_LOCKER

The **SIMCONNECT\_RECV\_CAMERA\_WORLD\_LOCKER** structure is a response with the different data related to the camera locker.

##### Syntax

```cpp
struct SIMCONNECT_RECV_CAMERA_WORLD_LOCKER: public SIMCONNECT_RECV
{
    SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS Status;
};
```

##### Members

| Member | Description |
| `Status` | Information about the camera locker status. See [SIMCONNECT\_CAMERA\_WORLD\_LOCKER\_STATUS](SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS.htm) |

##### Remarks

This struct will be the response to a call to `SimConnect_SubscribeToCameraWorldLockerStatusUpdate`.

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