SIMCONNECT\_DATA\_FACILITY\_AIRPORT

## SIMCONNECT\_DATA\_FACILITY\_AIRPORT

The **SIMCONNECT\_DATA\_FACILITY\_AIRPORT** structure is used to return information on a single airport in the facilities cache.

##### Syntax

```cpp
struct SIMCONNECT_DATA_FACILITY_AIRPORT{
    char ident[9];
    char region[3];
    double  Latitude;
    double  Longitude;
    double  Altitude;
    };
```

##### Members

| Member | Description |
| `ident` | The ICAO of the facility. |
| `region` | The ICAO of the region. |
| `Latitude` | Latitude of the airport facility. |
| `Longitude` | Longitude of the airport facility. |
| `Altitude` | Altitude of the facility in meters. |

##### Remarks

This structure is returned as one element in the [`SIMCONNECT_RECV_AIRPORT_LIST`](SIMCONNECT_RECV_AIRPORT_LIST.htm) structure. Note that this structure is inherited by [`SIMCONNECT_DATA_FACILITY_WAYPOINT`](SIMCONNECT_DATA_FACILITY_WAYPOINT.htm), [`SIMCONNECT_DATA_FACILITY_NDB`](SIMCONNECT_DATA_FACILITY_NDB.htm) and [`SIMCONNECT_DATA_FACILITY_VOR`](SIMCONNECT_DATA_FACILITY_VOR.htm), so the latitude, longitude, and altitude will apply to those facilities in that case.

See the remarks for [SimConnect\_RequestFacilitesList](../Facilities/SimConnect_RequestFacilitesList.htm).

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