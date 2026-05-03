SIMCONNECT\_DATA\_FACILITY\_WAYPOINT

## SIMCONNECT\_DATA\_FACILITY\_WAYPOINT

The **SIMCONNECT\_DATA\_FACILITY\_WAYPOINT** structure used to return information on a single waypoint in the facilities cache.

##### Syntax

```cpp
struct SIMCONNECT_DATA_FACILITY_WAYPOINT : public SIMCONNECT_DATA_FACILITY_AIRPORT {
    float  fMagVar;
    };
```

##### Members

| Member | Description |
| `fMagVar` | The magnetic variation of the waypoint in degrees. |

##### Remarks

This structure is returned as one element in the [`SIMCONNECT_RECV_WAYPOINT_LIST`](SIMCONNECT_RECV_WAYPOINT_LIST.htm) structure. It inherits all the members of the [`SIMCONNECT_DATA_FACILITY_AIRPORT`](SIMCONNECT_DATA_FACILITY_AIRPORT.htm) structure.

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