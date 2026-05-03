SIMCONNECT\_DATA\_INITPOSITION

## SIMCONNECT\_DATA\_INITPOSITION

The **SIMCONNECT\_DATA\_INITPOSITION** structure is used to initialize the position of the user aircraft, AI controlled aircraft, or other simulation object.

##### Syntax

```cpp
struct SIMCONNECT_DATA_INITPOSITION{
    double  Latitude;
    double  Longitude;
    double  Altitude;
    double  Pitch;
    double  Bank;
    double  Heading;
    DWORD  OnGround;
    DWORD  Airspeed;
    };
```

##### Members

| Member | Description |
| `Latitude` | Latitude in degrees. |
| `Longitude` | Longitude in degrees. |
| `Altitude` | Altitude in feet. |
| `Pitch` | Pitch in degrees. |
| `Bank` | Bank in degrees. |
| `Heading` | Heading in degrees. |
| `OnGround` | Set this to 1 to place the object on the ground, or 0 if the object is to be airborne. |
| `Airspeed` | The airspeed in knots, or one of the following special values:<br>- `INITPOSITION_AIRSPEED_CRUISE` (-1): The aircraft's design cruising speed.<br>- `INITPOSITION_AIRSPEED_KEEP` (-2): Maintain the current airspeed. |

##### Remarks

The primary use of this structure is to initialize the positioning of the user aircraft, because it also optimizes some of the terrain systems and other Microsoft Flight Simulator 2024 systems. Simply setting parameters such as latitude, longitude and altitude does not perform this kind of optimization. This structure should not be used to incrementally move the user aircraft (as this will unnecessarily initiate the reloading of scenery), in this case change the latitude, longitude, altitude and other parameters of the aircraft appropriately (using the variables described in the [Simulation Variables document](../../../SimVars/Simulation_Variables.htm)).

This structure can be used to incrementally move or reposition an AI controlled aircraft, or any other aircraft not controlled by the user, as the terrain system optimizations are not performed in this case.

This structure is used by the functions: [SimConnect\_AICreateNonATCAircraft](../AI_Object/SimConnect_AICreateNonATCAircraft.htm), [SimConnect\_AICreateSimulatedObject](../AI_Object/SimConnect_AICreateSimulatedObject.htm) and [SimConnect\_AddToDataDefinition](../Events_And_Data/SimConnect_AddToDataDefinition.htm).

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