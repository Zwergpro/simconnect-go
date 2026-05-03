SIMCONNECT\_DATA\_WAYPOINT

## SIMCONNECT\_DATA\_WAYPOINT

The **SIMCONNECT\_DATA\_WAYPOINT** structure is used to define a single waypoint.

##### Syntax

```cpp
struct SIMCONNECT_DATA_WAYPOINT{
    double  Latitude;
    double  Longitude;
    double  Altitude;
    unsigned long  Flags;
    double  ktsSpeed;
    double  percentThrottle;
    };
```

##### Members

| Member | Description |
| `Latitude` | Latitude in degrees. |
| `Longitude` | Longitude in degrees. |
| `Altitude` | Altitude in feet. |
| `Flags` | Specifies the flags set for this waypoint, see [SIMCONNECT\_WAYPOINT\_FLAGS](SIMCONNECT_WAYPOINT_FLAGS.htm). These flags can be OR'ed together, for example:<br>`Flags = SIMCONNECT_WAYPOINT_ON_GROUND | SIMCONNECT_WAYPOINT_REVERSE;` |
| `ktsSpeed` | Specifies the required speed in knots. If a specific speed is required, then the `SIMCONNECT_WAYPOINT_SPEED_REQUESTED` flag must be set to True. |
| `percentThrottle` | Specifies the required throttle as a percentage. If a specific throttle percentage is required, then the `SIMCONNECT_WAYPOINT_THROTTLE_REQUESTED` flag must be set to True. |

##### Remarks

The [SimConnect\_AddToDataDefinition](../Events_And_Data/SimConnect_AddToDataDefinition.htm) call can be used to add a [SIMCONNECT\_DATA\_WAYPOINT](SIMCONNECT_DATA_WAYPOINT.htm) structure to a data definition. A list of waypoints is sent to an AI object using the [SimConnect\_SetDataOnSimObject](../Events_And_Data/SimConnect_SetDataOnSimObject.htm) function. There is no limit to the number of waypoints that can be sent to an object. If just one waypoint is set, the [SIMCONNECT\_WAYPOINT\_WRAP\_TO\_FIRST](SIMCONNECT_WAYPOINT_FLAGS.htm) flag should not be used.

If a speed is requested at a waypoint, the slower that speed is the closer the object will approach the exact point of the waypoint, requests for high speeds can result in the AI system turning the object some way off of the waypoint. The pitch, bank and heading of objects controlled by the waypoint system are determined by the AI pilot, and cannot be set from a client.

This structure can only be used to set data, it cannot be used as part of a data request.

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