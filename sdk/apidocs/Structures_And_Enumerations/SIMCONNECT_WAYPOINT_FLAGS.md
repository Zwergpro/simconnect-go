SIMCONNECT\_WAYPOINT\_FLAGS

## SIMCONNECT\_WAYPOINT\_FLAGS

The **SIMCONNECT\_WAYPOINT\_FLAGS** enumeration type is used with the [SIMCONNECT\_DATA\_WAYPOINT](SIMCONNECT_DATA_WAYPOINT.htm) structure to define waypoints.

##### Syntax

```cpp
enum SIMCONNECT_WAYPOINT_FLAGS{
    SIMCONNECT_WAYPOINT_SPEED_REQUESTED = 0x04,
    SIMCONNECT_WAYPOINT_THROTTLE_REQUESTED = 0x08,
    SIMCONNECT_WAYPOINT_COMPUTE_VERTICAL_SPEED = 0x10,
    SIMCONNECT_WAYPOINT_ALTITUDE_IS_AGL = 0x20,
    SIMCONNECT_WAYPOINT_ON_GROUND = 0x00100000,
    SIMCONNECT_WAYPOINT_REVERSE = 0x00200000,
    SIMCONNECT_WAYPOINT_WRAP_TO_FIRST = 0x00400000,
    };
```

##### Members

| Member | Description |
| `SIMCONNECT_WAYPOINT_SPEED_REQUESTED` | Specifies requested speed is valid. |
| `SIMCONNECT_WAYPOINT_THROTTLE_REQUESTED` | Specifies requested throttle percentage is valid. |
| `SIMCONNECT_WAYPOINT_COMPUTE_VERTICAL_SPEED` | Specifies that the vertical should be calculated to reach the required speed when crossing the waypoint. |
| `SIMCONNECT_WAYPOINT_ALTITUDE_IS_AGL` | Specifies the altitude specified is AGL. |
| `SIMCONNECT_WAYPOINT_ON_GROUND` | Specifies the waypoint should be on the ground. Make sure this flag is set if the aircraft is to taxi to this point. |
| `SIMCONNECT_WAYPOINT_REVERSE` | Specifies that the aircraft should back up to this waypoint. This is only valid on the first waypoint. |
| `SIMCONNECT_WAYPOINT_WRAP_TO_FIRST` | Specifies that the next waypoint is the first waypoint. This is only valid on the last waypoint. |

##### Remarks

To set multiple waypoint flags simply OR them together. See the remarks for the `SIMCONNECT_DATA_WAYPOINT` structure.

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