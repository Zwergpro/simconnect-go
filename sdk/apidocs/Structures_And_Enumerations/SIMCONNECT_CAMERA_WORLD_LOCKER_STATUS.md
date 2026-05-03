SIMCONNECT\_CAMERA\_WORLD\_LOCKER\_STATUS

## SIMCONNECT\_CAMERA\_WORLD\_LOCKER\_STATUS

The **SIMCONNECT\_CAMERA\_WORLD\_LOCKER\_STATUS** enumeration type is used to describe the different possible camera locker status.

##### Syntax

```cpp
enum SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS
{
    SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_NONE,
    SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_START,
    SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_SUCCESS,
    SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_CANCEL,
    SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_FAIL
    };
```

##### Members

| Member | Description |
| `SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_NONE` | Default value, not used by the simulation. |
| `SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_START` | The camera locker has been setup and start loading the terrain and scenery around it. |
| `SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_SUCCESS` | The loading around the locker is finished. |
| `SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_CANCEL` | The camera locker has been deleted by the addon or by the sim itself (in case of low memory for example) |
| `SIMCONNECT_CAMERA_WORLD_LOCKER_STATUS_FAIL` | Internal error. |

##### Remarks

N/A

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