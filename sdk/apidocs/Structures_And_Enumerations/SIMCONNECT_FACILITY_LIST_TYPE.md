SIMCONNECT\_FACILITY\_LIST\_TYPE

## SIMCONNECT\_FACILITY\_LIST\_TYPE

The **SIMCONNECT\_FACILITY\_LIST\_TYPE** enumeration type is used to determine which type of facilities data is being requested or returned.

##### Syntax

```cpp
enum SIMCONNECT_FACILITY_LIST_TYPE{
    SIMCONNECT_FACILITY_LIST_TYPE_AIRPORT,
    SIMCONNECT_FACILITY_LIST_TYPE_WAYPOINT,
    SIMCONNECT_FACILITY_LIST_TYPE_NDB,
    SIMCONNECT_FACILITY_LIST_TYPE_VOR,
    SIMCONNECT_FACILITY_LIST_TYPE_COUNT
    };
```

##### Members

| Member | Description |
| `SIMCONNECT_FACILITY_LIST_TYPE_AIRPORT` | Specifies that the type of information is for an airport, see [`SIMCONNECT_DATA_FACILITY_AIRPORT`](SIMCONNECT_DATA_FACILITY_AIRPORT.htm). |
| `SIMCONNECT_FACILITY_LIST_TYPE_WAYPOINT` | Specifies that the type of information is for a waypoint, see [`SIMCONNECT_DATA_FACILITY_WAYPOINT`](SIMCONNECT_DATA_FACILITY_WAYPOINT.htm). |
| `SIMCONNECT_FACILITY_LIST_TYPE_NDB` | Specifies that the type of information is for an NDB, see [`SIMCONNECT_DATA_FACILITY_NDB`](SIMCONNECT_DATA_FACILITY_NDB.htm). |
| `SIMCONNECT_FACILITY_LIST_TYPE_VOR` | Specifies that the type of information is for a VOR, see [`SIMCONNECT_DATA_FACILITY_VOR`](SIMCONNECT_DATA_FACILITY_VOR.htm). |
| `SIMCONNECT_FACILITY_LIST_TYPE_COUNT` | Not valid as a list type, but simply the number of list types. |

##### Remarks

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