SIMCONNECT\_FACILITY\_DATA\_TYPE

## SIMCONNECT\_FACILITY\_DATA\_TYPE

The **SIMCONNECT\_FACILITY\_DATA\_TYPE** enumeration type is used within the [SIMCONNECT\_RECV\_FACILITY\_DATA](SIMCONNECT_RECV_FACILITY_DATA.htm) return to give the type of data that is being received.

##### Syntax

```cpp
enum SIMCONNECT_FACILITY_DATA_TYPE{
    SIMCONNECT_FACILITY_DATA_AIRPORT,
    SIMCONNECT_FACILITY_DATA_RUNWAY,
    SIMCONNECT_FACILITY_DATA_START,
    SIMCONNECT_FACILITY_DATA_FREQUENCY,
    SIMCONNECT_FACILITY_DATA_HELIPAD,
    SIMCONNECT_FACILITY_DATA_APPROACH,
    SIMCONNECT_FACILITY_DATA_APPROACH_TRANSITION,
    SIMCONNECT_FACILITY_DATA_APPROACH_LEG,
    SIMCONNECT_FACILITY_DATA_FINAL_APPROACH_LEG,
    SIMCONNECT_FACILITY_DATA_MISSED_APPROACH_LEG,
    SIMCONNECT_FACILITY_DATA_DEPARTURE,
    SIMCONNECT_FACILITY_DATA_ARRIVAL,
    SIMCONNECT_FACILITY_DATA_RUNWAY_TRANSITION,
    SIMCONNECT_FACILITY_DATA_ENROUTE_TRANSITION,
    SIMCONNECT_FACILITY_DATA_TAXI_POINT,
    SIMCONNECT_FACILITY_DATA_TAXI_PARKING,
    SIMCONNECT_FACILITY_DATA_TAXI_PATH,
    SIMCONNECT_FACILITY_DATA_TAXI_NAME,
    SIMCONNECT_FACILITY_DATA_JETWAY,
    SIMCONNECT_FACILITY_DATA_VOR,
    SIMCONNECT_FACILITY_DATA_NDB,
    SIMCONNECT_FACILITY_DATA_WAYPOINT,
    SIMCONNECT_FACILITY_DATA_ROUTE,
    SIMCONNECT_FACILITY_DATA_PAVEMENT,
    SIMCONNECT_FACILITY_DATA_APPROACH_LIGHTS,
    SIMCONNECT_FACILITY_DATA_VASI
    };
```

##### Members

| Member | Description |
| `SIMCONNECT_FACILITY_DATA_AIRPORT` | Contents of the parent struct are for an airport. See [AIRPORT](../Facilities/SimConnect_AddToFacilityDefinition.htm#airport) for more information. |
| `SIMCONNECT_FACILITY_DATA_RUNWAY` | Contents of the parent struct are for a runway. See [RUNWAY](../Facilities/SimConnect_AddToFacilityDefinition.htm#runway) for more information. |
| `SIMCONNECT_FACILITY_DATA_START` | Contents of the parent struct are for defining an airport start position. See [START](../Facilities/SimConnect_AddToFacilityDefinition.htm#start) for more information. |
| `SIMCONNECT_FACILITY_DATA_FREQUENCY` | Contents of the parent struct are for frequencies. See [FREQUENCY](../Facilities/SimConnect_AddToFacilityDefinition.htm#frequency) for more information. |
| `SIMCONNECT_FACILITY_DATA_HELIPAD` | Contents of the parent struct are for a helipad. See [HELIPAD](../Facilities/SimConnect_AddToFacilityDefinition.htm#helipad) for more information. |
| `SIMCONNECT_FACILITY_DATA_APPROACH` | Contents of the parent struct are for an approach. See [APPROACH](../Facilities/SimConnect_AddToFacilityDefinition.htm#approach) for more information. |
| `SIMCONNECT_FACILITY_DATA_APPROACH_TRANSITION` | Contents of the parent struct are for an approach transition. See [APPROACH\_TRANSITION](../Facilities/SimConnect_AddToFacilityDefinition.htm#approach_transition) for more information. |
| `SIMCONNECT_FACILITY_DATA_APPROACH_LEG` | Contents of the parent struct are for an approach leg. See [APPROACH\_LEG](../Facilities/SimConnect_AddToFacilityDefinition.htm#approach_leg) for more information. |
| `SIMCONNECT_FACILITY_DATA_FINAL_APPROACH_LEG` | Contents of the parent struct are for a final approach leg. See [FINAL\_APPROACH\_LEG](../Facilities/SimConnect_AddToFacilityDefinition.htm#final_approach_leg) for more information. |
| `SIMCONNECT_FACILITY_DATA_MISSED_APPROACH_LEG` | Contents of the parent struct are for a missed approach leg. See [MISSED\_APPROACH\_LEG](../Facilities/SimConnect_AddToFacilityDefinition.htm#missed_approach_leg) for more information. |
| `SIMCONNECT_FACILITY_DATA_DEPARTURE` | Contents of the parent struct are for a departure. See [DEPARTURE](../Facilities/SimConnect_AddToFacilityDefinition.htm#departure) for more information. |
| `SIMCONNECT_FACILITY_DATA_ARRIVAL` | Contents of the parent struct are for an arrival. See [ARRIVAL](../Facilities/SimConnect_AddToFacilityDefinition.htm#arrival) for more information. |
| `SIMCONNECT_FACILITY_DATA_RUNWAY_TRANSITION` | Contents of the parent struct are for a runway transition. See [RUNWAY\_TRANSITION](../Facilities/SimConnect_AddToFacilityDefinition.htm#runway_transition) for more information. |
| `SIMCONNECT_FACILITY_DATA_ENROUTE_TRANSITION` | Contents of the parent struct are for a route transition. See [ENROUTE\_TRANSITION](../Facilities/SimConnect_AddToFacilityDefinition.htm#enroute_transition) for more information. |
| `SIMCONNECT_FACILITY_DATA_TAXI_POINT` | Contents of the parent struct are for a taxiway point. See [TAXI\_POINT](../Facilities/SimConnect_AddToFacilityDefinition.htm#taxi_point) for more information. |
| `SIMCONNECT_FACILITY_DATA_TAXI_PARKING` | Contents of the parent struct are for a taxiway parking spot. See [TAXI\_PARKING](../Facilities/SimConnect_AddToFacilityDefinition.htm#taxi_parking) for more information. |
| `SIMCONNECT_FACILITY_DATA_TAXI_PATH` | Contents of the parent struct are for a taxiway path. See [TAXI\_PATH](../Facilities/SimConnect_AddToFacilityDefinition.htm#taxi_path) for more information. |
| `SIMCONNECT_FACILITY_DATA_TAXI_NAME` | Contents of the parent struct are for a taxi name. See [TAXI\_NAME](../Facilities/SimConnect_AddToFacilityDefinition.htm#taxi_name) for more information. |
| `SIMCONNECT_FACILITY_DATA_JETWAY` | Contents of the parent struct are for a jetway. See [JETWAY](../Facilities/SimConnect_AddToFacilityDefinition.htm#jetway) for more information. |
| `SIMCONNECT_FACILITY_DATA_VOR` | Contents of the parent struct are for a VOR station. See [VOR](../Facilities/SimConnect_AddToFacilityDefinition.htm#vor) for more information. |
| `SIMCONNECT_FACILITY_DATA_NDB` | Contents of the parent struct are for an NDB station. See [NDB](../Facilities/SimConnect_AddToFacilityDefinition.htm#ndb) for more information. |
| `SIMCONNECT_FACILITY_DATA_WAYPOINT` | Contents of the parent struct are for a waypoint. See [WAYPOINT](../Facilities/SimConnect_AddToFacilityDefinition.htm#waypoint) for more information. |
| `SIMCONNECT_FACILITY_DATA_ROUTE` | Contents of the parent struct are for a route. See [ROUTE](../Facilities/SimConnect_AddToFacilityDefinition.htm#route) for more information. |
| `SIMCONNECT_FACILITY_DATA_PAVEMENT` | Contents of the parent struct are for a pavement element. See [PAVEMENT](../Facilities/SimConnect_AddToFacilityDefinition.htm#PAVEMENT) for more information. |
| `SIMCONNECT_FACILITY_DATA_APPROACH_LIGHTS` | Contents of the parent struct are for the runway approach lights. See [APPROACHLIGHTS](../Facilities/SimConnect_AddToFacilityDefinition.htm#APPROACHLIGHTS) for more information. |
| `SIMCONNECT_FACILITY_DATA_VASI` | Contents of the parent struct are for VASI information. See [VASI](../Facilities/SimConnect_AddToFacilityDefinition.htm#VASI) for more information. |

##### Remarks

See the remarks and examples for [SimConnect\_RequestFacilityData](../Facilities/SimConnect_RequestFacilityData.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_RequestFacilityData](../Facilities/SimConnect_RequestFacilityData.htm)
4. [SIMCONNECT\_RECV\_FACILITIES\_LIST](SIMCONNECT_RECV_FACILITIES_LIST.htm)
5. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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