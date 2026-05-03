SIMCONNECT\_DATA\_FACILITY\_NDB

## SIMCONNECT\_DATA\_FACILITY\_NDB

The **SIMCONNECT\_DATA\_FACILITY\_NDB** structure is used to return information on a single NDB station in the facilities cache.

##### Syntax

```cpp
struct SIMCONNECT_DATA_FACILITY_NDB : public SIMCONNECT_DATA_FACILITY_WAYPOINT{
    DWORD  fFrequency;
    };
```

##### Members

| Member | Description |
| `fFrequency` | Frequency of the station in Hz. |

##### Remarks

This structure is returned as one element in the [SIMCONNECT\_RECV\_NDB\_LIST](SIMCONNECT_RECV_NDB_LIST.htm) structure. It inherits all the members of the [SIMCONNECT\_DATA\_FACILITY\_WAYPOINT](SIMCONNECT_DATA_FACILITY_WAYPOINT.htm) structure.

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