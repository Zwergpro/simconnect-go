SimConnect\_AddToFacilityDefinition

## SimConnect\_AddToFacilityDefinition

The **SimConnect\_AddToFacilityDefinition** function is used to create a facility data definition.

##### Syntax

```cpp
HRESULT SimConnect_AddToFacilityDefinition(
    HANDLE hSimConnect,
    SIMCONNECT_DATA_DEFINITION_ID DefineID,
    const char * FieldName
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `DefineID` | Specifies the ID of the client defined data definition | Integer |
| `FieldName` | Specifies the name of the field you want to add to the object definition. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| `S_OK` | The function succeeded. |
| `E_FAIL` | The function failed. |

It is also possible that the following [Exception IDs](../Structures_And_Enumerations/SIMCONNECT_EXCEPTION.htm) will be received:

- `SIMCONNECT_EXCEPTION_UNRECOGNIZED_ID` \- An invalid DefineID has been supplied.
- `SIMCONNECT_EXCEPTION_DATA_ERROR` \- An invalid FieldName has been supplied.

##### Remarks

This is used to add a new field to a client defined object definition in order to retrieve information. Using this function will require multiple calls and these calls will need to be formatted in a very specific way, initially using one of the four available **entry points**:

- [AIRPORT](SimConnect_AddToFacilityDefinition.htm#airport)
- [WAYPOINT](SimConnect_AddToFacilityDefinition.htm#waypoint)
- [NDB](SimConnect_AddToFacilityDefinition.htm#ndb)
- [VOR](SimConnect_AddToFacilityDefinition.htm#vor)

Each entry point will need to be prefixed with "OPEN " before any data is requested, and then prefixed with "CLOSED " to end the data retrieval, and this format works for the different child members of these entry points. For example:

```cpp
SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "OPEN AIRPORT");
// Request further airport data
SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "CLOSE AIRPORT");
SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "OPEN NDB");
// Request further NDB data
SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "CLOSE NDB");
// etc...
```

Essentially, for every member that you _open_ you must have a corresponding _close_, otherwise you will have an error. Note too that after defining the data that you want to retrieve, you will need to call `SimConnect_RequestFacilityData` to finalise and send the request.

At the bottom of this page you can find information on each of the entry points and the different members and children that they can contain, and if you expand the text below, you can find a schema showing the hierarchy of members and children within the returned data.

Struct Member HierarchyStruct Member Hierarchy

![](../../../../assets/images/6_Programming/SimConnect/simconnect_facilities_schema.png)

Note that results may be filtered using the `SimConnect_AddFacilityDataDefinitionFilter` function, ensuring that you receive less data, and that the data is relevant.

##### Example

```cpp
SIMCONNECT_DATA_DEFINITION_ID FACILITY_DATA_DEF_AIRPORT = 123;
{
    // Entry point will be an Airport
    SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "OPEN AIRPORT");

    // Add some airport members
    SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "A");
    SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "B");
    SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "C");

    // Runway is a child of Airport, so we can request them in Airport FacilityDataDefinition
    SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "OPEN RUNWAY");

    // Add some runway members
    SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "a");
    SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "b");

    // We requested every thing we want about runway, so "close" it to get back to Airport Definition
    SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "CLOSE RUNWAY");

    // We requested every thing we want about airport, so "close" it.
    SimConnect_AddToFacilityDefinition(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "CLOSE AIRPORT");

    // Request data from Airport LFPG which will follow the format defined earlier.
    SimConnect_RequestFacilityData(hSimConnect, FACILITY_DATA_DEF_AIRPORT, 123, "LFPG");
}
```

In the example given above, we have created a facility data definition which will request members "A", "B", and "C" of an airport and members "a" and "b" of each runway for the airport with ICAO "LFPG" (which has four runways). You should receive back from this the following:

- A message about Airport data \[A, B, C\]
- 4 messages about runway data \[a, b\]

So, imagine we complete an example request like the one above and each runway has a child with members "1", "2", and "3". In that case you would receive back the following data:

```
[A, B, C] [a, b] [1, 2, 3] [a, b] [1, 2, 3] [a, b] [1, 2, 3] [a, b] [1, 2, 3]
```

##### See Also

- [SimConnect API Reference](../../SimConnect_API_Reference.htm)
- [SimConnect\_RequestFacilityData](SimConnect_RequestFacilityData.htm)
- [SimConnect\_RequestFacilitesList](SimConnect_RequestFacilitesList.htm)

### AIRPORT

For the airport entry point you can request the following information directly:

| Member Name | DataType | Description |
| `LATITUDE` | FLOAT64 | The airport latitude, in degrees. |
| `LONGITUDE` | FLOAT64 | The airport longitude, in degrees. |
| `ALTITUDE` | FLOAT64 | The airport altitude, in meters. |
| `MAGVAR` | FLOAT32 | This is the magnetic variation for the airport position. |
| `NAME` | STRING32 | This is the name of the airport itself. |
| `NAME64` | STRING64 | This is the name of the airport itself using a format that permits more than 32 characters. |
| `ICAO` | STRING8 | The airport ICAO code. |
| `REGION` | STRING8 | The airport [region code](https://en.wikipedia.org/wiki/ICAO_airport_code#Prefixes "Wikipedia - Region Code Prefixes"). |
| `TOWER_LATITUDE` | FLOAT64 | The control tower latitude. |
| `TOWER_LONGITUDE` | FLOAT64 | The control tower longitude. |
| `TOWER_ALTITUDE` | FLOAT64 | The control tower altitude. |
| `TRANSITION_ALTITUDE` | FLOAT32 | The transition altitude, in meters, or 0 if not defined. |
| `TRANSITION_LEVEL` | FLOAT32 | The transition level, in meters, or 0 if not defined. |
| `IS_CLOSED` | INT8 | The airport close state. |
| `COUNTRY` | STRING256 | The name of the contry in which the airport is located. |
| `CITY_STATE` | STRING256 | The name of the city/state in which the airport is located. |
| `N_RUNWAYS` | INT32 | The number of runways available at the airport. Individual runway data can be retrieved using the `RUNWAY` child member. |
| `N_STARTS` | INT32 | The number of starting points available at the airport. Individual runway data can be retrieved using the `START` child member. |
| `N_FREQUENCIES` | INT32 | The number of frequencies available at the airport. Individual frequencies can be retrieved using the `FREQUENCY` child member. |
| `N_HELIPADS` | INT32 | The number of helipads available at the airport. Individual helipad data can be retrieved using the `HELIPAD` child member. |
| `N_APPROACHES` | INT32 | The number of approaches available at the airport. Individual approach data can be retrieved using the `APPROACH` child member. |
| `N_DEPARTURES` | INT32 | The number of departures available at the airport. Individual departure data can be retrieved using the `DEPARTURE` child member. |
| `N_ARRIVALS` | INT32 | The number of arrivals available at the airport. Individual arrivals can be retrieved using the `ARRIVAL` child member. |
| `N_TAXI_POINTS` | INT32 | The number of taxiway points available. Individual points can be retrieved using the `TAXI_POINT` child member. |
| `N_TAXI_PARKINGS` | INT32 | The number of taxiway parking spots available at the airport. Individual parking spot data can be retrieved using the `TAXI_PARKING` child member. |
| `N_TAXI_PATHS` | INT32 | The number of taxiway paths available at the airport. Individual path data can be retrieved using the `TAXI_PATH` child member. |
| `N_TAXI_NAMES` | INT32 | The number of taxiway names available at the airport. Individual taxiway name data can be retrieved using the `TAXI_NAME` child member. |
| `N_JETWAYS` | INT32 | The number of jetways available at the airport. Individual jetway data can be retrieved using the `JETWAY` child member. |
| `N_VDGS` | INT32 | The number of VDGS available at the airport. Individual VDGS data can be retrieved using the `JETWAY` child member. |
| `N_HOLDING_PATTERNS` | INT32 | The number of VDGS available at the airport. Individual holding pattern data can be retrieved using the `JETWAY` child member. |
| `TAXI_NAME` | STRUCT: `TAXI_NAME` | Contains the different taxiway names. |
| `TAXI_PATH` | STRUCT: `TAXI_PATH` | Contains taxiway path data. |
| `TAXI_PARKING` | STRUCT: `TAXI_PARKING` | Contains taxiway parking spot data. |
| `TAXI_POINT` | STRUCT: `TAXI_POINT` | Contains data on the various taxiway points. |
| `ARRIVAL` | STRUCT: `ARRIVAL` | Contains airport arrival data. |
| `DEPARTURE` | STRUCT: `DEPARTURE` | Contains airport departure data. |
| `APPROACH` | STRUCT: `APPROACH` | Contains airport approach data. |
| `HELIPAD` | STRUCT: `HELIPAD` | Contains data related to any helipads at the airport. |
| `FREQUENCY` | STRUCT: `FREQUENCY` | Contains data related to the navigation aids and frequencies for the airport. |
| `JETWAY` | STRUCT: `JETWAY` | Contains data on the various airport jetways. |
| `VDGS` | STRUCT: `VDGS` | Contains data on the various parking spot VDGS. |
| `HOLDING_PATTERN` | STRUCT: `HOLDING_PATTERN` | Contains data related to the holding patterns available at the airport. |
| `START` | STRUCT: `START` | Contains data related to the various runway start points for aircraft. |
| `RUNWAY` | STRUCT: `RUNWAY` | Contains data related to the various runways. |

##### TAXI\_NAME

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `NAME` | STRING32 | A taxiway name. |

##### TAXI\_PATH

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `TYPE` | INT32 | the type of taxiway path. Can be any one of the following values:<br>1. 0: NONE<br>2. 1: TAXI<br>3. 2: RUNWAY<br>4. 3: PARKING<br>5. 4: PATH<br>6. 5: CLOSED<br>7. 6: VEHICLE<br>8. 7: ROAD<br>9. 8: PAINTEDLINE |
| `WIDTH` | FLOAT32 | The width of the taxiway, in meters. |
| `LEFT_HALF_WIDTH` | FLOAT32 | The left-side width in case of an asymmetric taxiway. |
| `RIGHT_HALF_WIDTH` | FLOAT32 | The right-side width in case of an asymmetric taxiway. |
| `WEIGHT` | UINT32 | The weight limit in lbs. |
| `RUNWAY_NUMBER` | INT32 | The number of the runway this taxiway path is for, one of the following values:<br> <br>01. 0: NONE<br>02. 1 - 36: RUNWAY ID<br>03. 37: NORTH<br>04. 38: NORTHEAST<br>05. 39: EAST<br>06. 40: SOUTHEAST<br>07. 41: SOUTH<br>08. 42: SOUTHWEST<br>09. 43: WEST<br>10. 44: NORTHWEST<br>11. 45: LAST |
| `RUNWAY_DESIGNATOR` | INT32 | The Designator of the runway this taxiway path is for, one of the following values:<br> <br>1. 0: NONE<br>2. 1: LEFT<br>3. 2: RIGHT<br>4. 3: CENTER<br>5. 4: WATER<br>6. 5: A<br>7. 6: B<br>8. 7: LAST |
| `LEFT_EDGE` | INT32 | The left edge type, one of the following:<br>1. 0: NONE<br>2. 1: SOLID<br>3. 2: DASHED<br>4. 3: SOLID\_DASHED |
| `LEFT_EDGE_LIGHTED` | INT32 | Whether the taxiway path left edge line is lit (1, TRUE) or not (0, FALSE). The default is 0. |
| `RIGHT_EDGE` | INT32 | The right edge type, one of the following:<br>1. 0: NONE<br>2. 1: SOLID<br>3. 2: DASHED<br>4. 3: SOLID\_DASHED |
| `RIGHT_EDGE_LIGHTED` | INT32 | Whether the taxiway path right edge line is lit (1, TRUE) or not (0, FALSE). The default is 0. |
| `CENTER_LINE` | INT32 | Whether the taxiway path has a center line (1, TRUE) or not (0, FALSE). The default is 0. |
| `CENTER_LINE_LIGHTED` | INT32 | Whether the taxiway path center line is lit (1, TRUE) or not (0, FALSE). The default is 0. |
| `START` | INT32 | The index number of taxiway point or parking space the path starts from. Value from 0 to 65534. |
| `END` | INT32 | The index number of taxiway point or parking space the path ends on. Value from 0 to 65534. |
| `NAME_INDEX` | UINT32 | The name index. |

##### TAXI\_PARKING

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `TYPE` | INT32 | The type of parking spot. Can be any one of the following values:<br>01. 0: NONE<br>02. 1: RAMP\_GA<br>03. 2: RAMP\_GA\_SMALL<br>04. 3: RAMP\_GA\_MEDIUM<br>05. 4: RAMP\_GA\_LARGE<br>06. 5: RAMP\_CARGO<br>07. 6: RAMP\_MIL\_CARGO<br>08. 7: RAMP\_MIL\_COMBAT<br>09. 8: GATE\_SMALL<br>10. 9: GATE\_MEDIUM<br>11. 10: GATE\_HEAVY<br>12. 11: DOCK\_GA<br>13. 12: FUEL<br>14. 13: VEHICLE<br>15. 14: RAMP\_GA\_EXTRA<br>16. 15: GATE\_EXTRA |
| `TAXI_POINT_TYPE` | INT32 | Taxiway point type. Can be one of the following:<br>1. 0: NONE<br>2. 1: NORMAL<br>3. 2: HOLD\_SHORT<br>4. 4: ILS\_HOLD\_SHORT<br>5. 5: HOLD\_SHORT\_NO\_DRAW<br>6. 6: ILS\_HOLD\_SHORT\_NO\_DRAW |
| `NAME` | INT32 | The name of the parking spot. Can be any one of the following:<br>01. 0: NONE<br>02. 1: PARKING<br>03. 2: N\_PARKING<br>04. 3: NE\_PARKING<br>05. 4: E\_PARKING<br>06. 5: SE\_PARKING<br>07. 6: S\_PARKING<br>08. 7: SW\_PARKING<br>09. 8: W\_PARKING<br>10. 9: NW\_PARKING<br>11. 10: GATE<br>12. 11: DOCK<br>13. 12 - 37: GATE\_A - GATE\_Z |
| `SUFFIX` | INT32 | The suffix of the parking spot. Can be any one of the following:<br>01. 0: NONE<br>02. 1: PARKING<br>03. 2: N\_PARKING<br>04. 3: NE\_PARKING<br>05. 4: E\_PARKING<br>06. 5: SE\_PARKING<br>07. 6: S\_PARKING<br>08. 7: SW\_PARKING<br>09. 8: W\_PARKING<br>10. 9: NW\_PARKING<br>11. 10: GATE<br>12. 11: DOCK<br>13. 12 - 37: GATE\_A - GATE\_Z |
| `NUMBER` | UINT32 | The number of the parking place. |
| `ORIENTATION` | INT32 | The orientation when the taxi point type is hold short:<br>1. 0: FORWARD<br>2. 1: REVERSE |
| `HEADING` | FLOAT32 | The heading of the parking spot, in degrees true. |
| `RADIUS` | FLOAT32 | The size of the parking spot, in meters. |
| `BIAS_X` | FLOAT32 | Bias from airport reference along the longitudinal axis in meters. |
| `BIAS_Z` | FLOAT32 | Bias from airport reference along the latitudinal axis in meters. |
| `AIRLINE` | STRUCT: `AIRLINE` | Contains the names of the all the airlines that can use this taxi parking. |
| `N_AIRLINES` | INT32 | The number of [AIRLINE](SimConnect_AddToFacilityDefinition.htm#AIRLINE) s linked to this taxi parking. |

##### AIRLINE

This is a child member of the `TAXI_PARKING` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `NAME` | STRING8 | The name of the given airline. |

##### TAXI\_POINT

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `TYPE` | INT32 | Taxiway point type. Can be one of the following:<br>1. 0: NONE<br>2. 1: NORMAL<br>3. 2: HOLD\_SHORT<br>4. 4: ILS\_HOLD\_SHORT<br>5. 5: HOLD\_SHORT\_NO\_DRAW<br>6. 6: ILS\_HOLD\_SHORT\_NO\_DRAW |
| `ORIENTATION` | INT32 | The orientation when the type is hold short:<br>1. 0: FORWARD<br>2. 1: REVERSE |
| `BIAS_X` | FLOAT32 | Bias from airport reference along the longitudinal axis in meters. |
| `BIAS_Z` | FLOAT32 | Bias from airport reference along the latitudinal axis in meters. |

##### HELIPAD

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `LATITUDE` | FLOAT64 | The latitude of the center of the helipad, in degrees. |
| `LONGITUDE` | FLOAT64 | The longitude of the center of the helipad, in degrees. |
| `ALTITUDE` | FLOAT64 | The altitude of the center of the helipad, in meters. |
| `HEADING` | FLOAT32 | The heading of the helipad, in degrees. |
| `LENGTH` | FLOAT32 | The length of the helipad, in meters. |
| `WIDTH` | FLOAT32 | The width of the helipad, in meters. |
| `SURFACE` | INT32 | The type of pavement used by the helipad. The return value will be one of the following:<br>0: CONCRETE<br>1: GRASS<br>2: WATER FSX<br>3: GRASS BUMPY<br>4: ASPHALT<br>5: SHORT GRASS<br>6: LONG GRASS<br>7: HARD TURF<br>8: SNOW<br>9: ICE<br>10: URBAN<br>11: FOREST<br>12: DIRT<br>13: CORAL<br>14: GRAVEL<br>15: OIL TREATED<br>16: STEEL MATS<br>17: BITUMINUS<br>18: BRICK<br>19: MACADAM<br>20: PLANKS<br>21: SAND<br>22: SHALE<br>23: TARMAC<br>24: WRIGHT FLYER TRACK<br>26: OCEAN<br>27: WATER<br>28: POND<br>29: LAKE<br>30: RIVER<br>31: WASTE WATER<br>32: PAINT<br>254: UNKNOWN<br>255: UNDEFINED |
| `TYPE` | INT32 | The type of helipad, which will be one of the following values:<br>1. 0: NONE<br>2. 1: H<br>3. 2: SQUARE<br>4. 3: CIRCLE<br>5. 4: MEDICAL |
| `TOUCH_DOWN_LENGTH` | FLOAT32 | The length of TLOF area in meters. |
| `FATO_LENGTH` | FLOAT32 | The length of FATO (Final Approach and Takeoff) area in meters. |
| `FATO_WIDTH` | FLOAT32 | The width of FATO (Final Approach and Takeoff) area in meters. |

##### FREQUENCY

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `TYPE` | INT32 | The type of radio frequency. Can be one of the following values:<br>01. 0: NONE<br>02. 1: ATIS<br>03. 2: MULTICOM<br>04. 3: UNICOM<br>05. 4: CTAF<br>06. 5: GROUND<br>07. 6: TOWER<br>08. 7: CLEARANCE (Clearance Delivery)<br>09. 8: APPROACH<br>10. 9: DEPARTURE<br>11. 10: CENTER<br>12. 11: FSS<br>13. 12: AWOS<br>14. 13: ASOS<br>15. 14: CPT (Clearance Pre-Taxi)<br>16. 15: GCO (Remote Clearance Delivery) |
| `FREQUENCY` | INT32 | The actual frequency, in Hz. |
| `NAME` | STRING64 | The channel name for the frequency. |

##### JETWAY

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `PARKING_GATE` | INT32 | The name of the parking spot the jetway is assigned to. Can be any one of the following:<br>01. 0: NONE<br>02. 1: PARKING<br>03. 2: N\_PARKING<br>04. 3: NE\_PARKING<br>05. 4: E\_PARKING<br>06. 5: SE\_PARKING<br>07. 6: S\_PARKING<br>08. 7: SW\_PARKING<br>09. 8: W\_PARKING<br>10. 9: NW\_PARKING<br>11. 10: GATE<br>12. 11: DOCK<br>13. 12 - 37: GATE\_A - GATE\_Z |
| `PARKING_SUFFIX` | INT32 | The suffix of the parking spot the jetway is assigned to. Can be any one of the following:<br>01. 0: NONE<br>02. 1: PARKING<br>03. 2: N\_PARKING<br>04. 3: NE\_PARKING<br>05. 4: E\_PARKING<br>06. 5: SE\_PARKING<br>07. 6: S\_PARKING<br>08. 7: SW\_PARKING<br>09. 8: W\_PARKING<br>10. 9: NW\_PARKING<br>11. 10: GATE<br>12. 11: DOCK<br>13. 12 - 37: GATE\_A - GATE\_Z |
| `PARKING_SPOT` | INT32 | The index of taxiway point where the parking spot is located (a value between 0 and 65534.) |

##### VDGS

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `LATITUDE` | FLOAT64 | The latitude of the center of the VDGD, in degrees. |
| `LONGITUDE` | FLOAT64 | The longitude of the center of the VDGD, in degrees. |
| `ALTITUDE` | FLOAT64 | The altitude of the center of the VDGD, in degrees. |
| `PARKING_NUMBER` | INT32 | The number of the parking spot. the VDGS is assigned to. |
| `PARKING_GATE` | INT32 | The name of the parking spot the VDGS is assigned to. Can be any one of the following:<br>01. 0: NONE<br>02. 1: PARKING<br>03. 2: N\_PARKING<br>04. 3: NE\_PARKING<br>05. 4: E\_PARKING<br>06. 5: SE\_PARKING<br>07. 6: S\_PARKING<br>08. 7: SW\_PARKING<br>09. 8: W\_PARKING<br>10. 9: NW\_PARKING<br>11. 10: GATE<br>12. 11: DOCK<br>13. 12 - 37: GATE\_A - GATE\_Z |
| `PARKING_SUFFIX` | INT32 | The suffix of the parking spot the VDGS is assigned to. Can be any one of the following:<br>01. 0: NONE<br>02. 1: PARKING<br>03. 2: N\_PARKING<br>04. 3: NE\_PARKING<br>05. 4: E\_PARKING<br>06. 5: SE\_PARKING<br>07. 6: S\_PARKING<br>08. 7: SW\_PARKING<br>09. 8: W\_PARKING<br>10. 9: NW\_PARKING<br>11. 10: GATE<br>12. 11: DOCK<br>13. 12 - 37: GATE\_A - GATE\_Z |
| `PARKING_INDEX` | INT32 | The index of the parking spot. the VDGS is assigned to. |

##### HOLDING\_PATTERN

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `NAME` | STRING64 | The name of the holding pattern. |
| `FIX_ICAO` | STRING8 | ICAO of the defining point. |
| `FIX_REGION` | STRING8 | The Region of the defining point. |
| `FIX_TYPE` | INT32 | The type of defining point, one of the following:<br>1. 65 ('A') - Airport<br>2. 86 ('V') - VOR<br>3. 78 ('N') - NDB<br>4. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `INBOUND_HOLDING_COURSE` | FLOAT32 | The inbound course to the holding in degrees. |
| `TURN_RIGHT` | INT32 | Returns whether the holding turns right \[1\] or left \[0\]. |
| `LEG_LENGTH` | FLOAT32 | The distance between the point at which the aircraft rolls out on the inbound leg of the holding pattern and the fix at which the holding pattern is defined in nautical miles. |
| `LEG_TIME` | FLOAT32 | The length of the inbound leg of a holding pattern in units of time, in minutes. |
| `MIN_ALTITUDE` | FLOAT32 | The minimum altitude of the holding pattern. |
| `MAX_ALTITUDE` | FLOAT32 | The maximum altitude of the holding pattern. |
| `HOLD_SPEED` | FLOAT32 | The maximum speed in an holding pattern, in knots. |
| `REQUIRED_NAVIGATION_PERFORMANCE` | FLOAT32 | The required navigation performance for the leg in meters, or 0 if not defined. |
| `ARC_RADIUS` | FLOAT32 | The radius of the holding pattern, in nautical miles. |

#### ARRIVAL

This is a child member of the `AIRPORT` entry point, and is itself an entry point for the `RUNWAY_TRANSITION`, `ENROUTE_TRANSITION`, and `APPROACH_LEG` structs.

In addition to the member structs listed above, you can get more data using the following member parameters:

| Member Name | DataType | Description |
| `NAME` | STRING8 | The name of the arrival. |
| `IS_RNPAR` | INT32 | Returns whether the departure is RNP-AR \[1\] or not \[0\]. |
| `N_RUNWAY_TRANSITIONS` | INT32 | The number of [RUNWAY\_TRANSITION](#runway_transition) s for the arrival. |
| `N_ENROUTE_TRANSITIONS` | INT32 | The number of [ENROUTE\_TRANSITION](#enroute_transition) s for the arrival. |
| `N_APPROACH_LEGS` | INT32 | The number of [APPROACH\_LEG](#approach_leg) s for the arrival. |
| `RUNWAY_TRANSITION` | STRUCT: `RUNWAY_TRANSITION` | Contains runway transition data. |
| `ENROUTE_TRANSITION` | STRUCT: `ENROUTE_TRANSITION` | Contains enroute transition data. |
| `APPROACH_LEG` | STRUCT: `APPROACH_LEG` | Contains approach leg data. |

##### RUNWAY\_TRANSITION

This is a child member of the `DEPARTURE` and `ARRIVAL` entry points and is itself an entry point for the `APPROACH_LEG` struct.

In addition to the member struct listed above, you can get more data using the following member parameters:

| Member Name | DataType | Description |
| `RUNWAY_NUMBER` | INT32 | 01. The number of the runway this transition is for, one of the following values:<br>02. 0: NONE<br>03. 1 - 36: RUNWAY ID<br>04. 37: NORTH<br>05. 38: NORTHEAST<br>06. 39: EAST<br>07. 40: SOUTHEAST<br>08. 41: SOUTH<br>09. 42: SOUTHWEST<br>10. 43: WEST<br>11. 44: NORTHWEST<br>12. 45: LAST |
| `RUNWAY_DESIGNATOR` | INT32 | 1. The designator of the runway this transition is for, one of the following values:<br>2. 0: NONE<br>3. 1: LEFT<br>4. 2: RIGHT<br>5. 3: CENTER<br>6. 4: WATER<br>7. 5: A<br>8. 6: B<br>9. 7: LAST |
| `N_APPROACH_LEGS` | INT32 | The number of [APPROACH\_LEG](#approach_leg) s for the transition. |
| `APPROACH_LEG` | STRUCT: `APPROACH_LEG` | Contains approach leg data. |

##### ENROUTE\_TRANSITION

This is a child member of the `DEPARTURE` and `ARRIVAL` entry points and is itself an entry point for the `APPROACH_LEG` struct.

In addition to the member struct given above, you can get more data using the following member parameters:

| Member Name | DataType | Description |
| `NAME` | STRING8 | The transition name. |
| `N_APPROACH_LEGS` | INT32 | The number of [APPROACH\_LEG](#approach_leg) s for the transition. |
| `APPROACH_LEG` | STRUCT: `APPROACH_LEG` | Contains approach leg data. |

#### DEPARTURE

This is a child member of the `AIRPORT` entry point, and is itself an entry point for `RUNWAY_TRANSITION`, `ENROUTE_TRANSITION`, and `APPROACH_LEG` structs.

In addition to the member structs listed above, you can get more data using the following member parameters:

| Member Name | DataType | Description |
| `NAME` | STRING8 | The name of the departure. |
| `IS_RNPAR` | INT32 | Returns whether the departure is RNP-AR \[1\] or not \[0\]. |
| `N_RUNWAY_TRANSITIONS` | INT32 | The number of [RUNWAY\_TRANSITION](#runway_transition) s for the departure. |
| `N_ENROUTE_TRANSITIONS` | INT32 | The number of [ENROUTE\_TRANSITION](#enroute_transition) s for the departure. |
| `N_APPROACH_LEGS` | INT32 | The number of [APPROACH\_LEG](#approach_leg) s for the departure. |
| `RUNWAY_TRANSITION` | STRUCT: `RUNWAY_TRANSITION` | Contains runway transition data. |
| `ENROUTE_TRANSITION` | STRUCT: `ENROUTE_TRANSITION` | Contains enroute transition data. |
| `APPROACH_LEG` | STRUCT: `APPROACH_LEG` | Contains approach leg data. |

#### APPROACH

This is a child member of the `AIRPORT` entry point, and is itself an entry point for the `APPROACH_TRANSITION`, `FINAL_APPROACH_LEG`, and `MISSED_APPROACH_LEG` structs.

In addition to the member structs listed above, you can get more data using the following member parameters:

| Member Name | DataType | Description |
| `TYPE` | INT32 | The approach type. Can be one of the following values:<br>01. 0: UNDEFINED<br>02. 1: GPS<br>03. 2: VOR<br>04. 3: NDB<br>05. 4: ILS<br>06. 5: LOCALIZER<br>07. 6: SDF<br>08. 7: LDA<br>09. 8: VORDME<br>10. 9: NDBDME<br>11. 10: RNAV<br>12. 11: LOCALIZER\_BACK\_COURSE |
| `SUFFIX` | INT32 | The multiple indicator suffix (must be converted to Char Alphanumeric or blank). |
| `RUNWAY_NUMBER` | INT32 | The number of the runway this approach is for, one of the following values:<br> <br>01. 0: NONE<br>02. 1 - 36: RUNWAY ID<br>03. 37: NORTH<br>04. 38: NORTHEAST<br>05. 39: EAST<br>06. 40: SOUTHEAST<br>07. 41: SOUTH<br>08. 42: SOUTHWEST<br>09. 43: WEST<br>10. 44: NORTHWEST<br>11. 45: LAST |
| `RUNWAY_DESIGNATOR` | INT32 | The Designator of the runway this approach is for, one of the following values:<br> <br>1. 0: NONE<br>2. 1: LEFT<br>3. 2: RIGHT<br>4. 3: CENTER<br>5. 4: WATER<br>6. 5: A<br>7. 6: B<br>8. 7: LAST |
| `FAF_ICAO` | STRING8 | Final approach fix ICAO. |
| `FAF_REGION` | STRING8 | Final approach fix region |
| `FAF_HEADING` | FLOAT32 | Final approach fix heading, in degrees. |
| `FAF_ALTITUDE` | FLOAT32 | Final approach fix altitude, in Meters. |
| `FAF_TYPE` | INT32 | The final approach type, one of the following:<br>1. 65 ('A') - Airport<br>2. 86 ('V') - VOR<br>3. 78 ('N') - NDB<br>4. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `MISSED_ALTITUDE` | FLOAT32 | Altitude of the first leg of the Missed approach, in ft. |
| `HAS_LNAV` | INT32 | Returns whether the approach has lateral navigation (1) or not (0). |
| `HAS_LNAVVNAV` | INT32 | Returns whether the approach has lateral and vertical navigation (1) or not (0). |
| `HAS_LP` | INT32 | Returns whether the approach has localizer performance (1) or not (0). |
| `HAS_LPV` | INT32 | Returns whether the approach has localizer performance with vertical guidance (1) or not (0). |
| `IS_RNPAR` | INT32 | Returns whether the approach is RNP-AR \[1\] or not \[0\]. |
| `IS_RNPAR_MISSED` | INT32 | Returns whether the missedapproach is RNP-AR \[1\] or not \[0\]. |
| `N_TRANSITIONS` | INT32 | The number of [APPROACH\_TRANSITION](#approach_transition) s for the approach. |
| `N_FINAL_APPROACH_LEGS` | INT32 | The number of [FINAL\_APPROACH\_LEG](#final_approach_leg) for the approach. |
| `N_MISSED_APPROACH_LEGS` | INT32 | The number of [MISSED\_APPROACH\_LEG](#missed_approach_leg) the approach. |
| * * * |
| `APPROACH_TRANSITION` | STRUCT: `APPROACH_TRANSITION` | Contains approach transition data. |
| `FINAL_APPROACH_LEG` | STRUCT: `FINAL_APPROACH_LEG` | Contains the final approach leg data. |
| `MISSED_APPROACH_LEG` | STRUCT: `MISSED_APPROACH_LEG` | Contains the missed approach leg data. |

##### APPROACH\_TRANSITION

This is a child member of the `APPROACH` entry point, and is itself an entry point for the `APPROACH_LEG` struct.

In addition to the member struct given above, you can get more data using the following member parameters:

| Member Name | DataType | Description |
| `TYPE` | INT32 | The approach type. Can be one of the following values:<br>01. 0: UNDEFINED<br>02. 1: GPS<br>03. 2: VOR<br>04. 3: NDB<br>05. 4: ILS<br>06. 5: LOCALIZER<br>07. 6: SDF<br>08. 7: LDA<br>09. 8: VORDME<br>10. 9: NDBDME<br>11. 10: RNAV<br>12. 11: LOCALIZER\_BACK\_COURSE |
| `IAF_ICAO` | STRING8 | ICAO code of the NAVAID at the transition. |
| `IAF_REGION` | STRING8 | The region of the NAVAID at the transition. |
| `IAF_TYPE` | INT32 | The type of NAVAID transition, one of the following:<br>1. 65 ('A') - Airport<br>2. 86 ('V') - VOR<br>3. 78 ('N') - NDB<br>4. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `IAF_ALTITUDE` | FLOAT32 | Desired altitude at the transition, in ft. |
| `DME_ARC_ICAO` | STRING8 | ICAO to the DME station. |
| `DME_ARC_REGION` | STRING8 | The region for the DME station. |
| `DME_ARC_TYPE` | INT32 | The type of DME station, one of the following:<br>1. 65 ('A') - Airport<br>2. 86 ('V') - VOR<br>3. 78 ('N') - NDB<br>4. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `DME_ARC_RADIAL` | INT32 | The name of the radial for the arc. |
| `DME_ARC_DISTANCE` | FLOAT32 | DME distance from the radial, in meters. |
| `NAME` | STRING8 | The name of the transition. |
| `N_APPROACH_LEGS` | INT32 | The number of [APPROACH\_LEG](#approach_leg) s for the transition. |
| `APPROACH_LEG` | STRUCT: `APPROACH_LEG` | Contains approach leg data. |

##### APPROACH\_LEG

##### FINAL\_APPROACH\_LEG

##### MISSED\_APPROACH\_LEG

This is a child member of the `APPROACH`, `APPROACH_TRANSITION`, `RUNWAY_TRANSITION`, `ENROUTE_TRANSITION`, `DEPARTURE` and `ARRIVAL` entry points, and has the following available data:

| Member Name | DataType | Description |
| `TYPE` | INT32 | The approach leg type, which can be one of the following values:<br>01. 0: UNKNOWN,<br>02. 1: AF - dme arc to fix<br>03. 2: CA - course to altitude<br>04. 3: CD - course to dme distance<br>05. 4: CF - course to fix<br>06. 5: CI - course to intercept<br>07. 6: CR - course to radial<br>08. 7: DF - direct to fix<br>09. 8: FA - fix to altitude<br>10. 9: FC - track from fix<br>11. 10: FD - track from fix to dme distance<br>12. 11: FM - track from fix to manual terminator<br>13. 12: HA - racetrack course reversal to altitude<br>14. 13: HF - racetrack course reversal to fix<br>15. 14: HM - racetrack course reversal to manual terminator<br>16. 15: IF - initial fix<br>17. 16: PI - procedure turn<br>18. 17: RF - constant radius arc<br>19. 18: TF - track to fix<br>20. 19: VA - heading to altitude<br>21. 20: VD - heading to dme distance<br>22. 21: VI - heading to intercept<br>23. 22: VM - heading to manual termination<br>24. 23: VR - heading to radial |
| `FIX_ICAO` | STRING8 | ICAO of the defining point. |
| `FIX_REGION` | STRING8 | The Region of the defining point. |
| `FIX_TYPE` | INT32 | The type of defining point, one of the following:<br>1. 65 ('A') - Airport<br>2. 86 ('V') - VOR<br>3. 78 ('N') - NDB<br>4. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `FIX_LATITUDE` | FLOAT64 | The latitude of the defining point. |
| `FIX_LONGITUDE` | FLOAT64 | The longitude of the defining point. |
| `FIX_ALTITUDE` | FLOAT64 | The altitude of the defining point. |
| `FLY_OVER` | INT32 | Whether the point is fly-over (1) as opposed to fly-by (0). |
| `DISTANCE_MINUTE` | INT32 | When (1), the `ROUTE_DISTANCE` field is in minutes, when (0) it is in meters. |
| `TRUE_DEGREE` | INT32 | When (1) `THETA` and `COURSE` are actually true north degrees, when (0) they are magnetic north degrees. |
| `TURN_DIRECTION` | INT32 | The approach turn direction, which can be one of the following values:<br>1. 0: NONE<br>2. 1: LEFT<br>3. 2: RIGHT<br>4. 3: EITHER |
| `ORIGIN_ICAO` | STRING8 | ICAO of origin navaid. |
| `ORIGIN_REGION` | STRING8 | The region of origin navaid. |
| `ORIGIN_TYPE` | INT32 | The type of origin navaid, one of the following:<br>1. 65 ('A') - Airport<br>2. 86 ('V') - VOR<br>3. 78 ('N') - NDB<br>4. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `ORIGIN_LATITUDE` | FLOAT64 | The latitude of the origin navaid. |
| `ORIGIN_LONGITUDE` | FLOAT64 | The longitude of the origin navaid. |
| `ORIGIN_ALTITUDE` | FLOAT64 | The altitude of the origin navaid. |
| `THETA` | FLOAT32 | The magnetic bearing to the waypoint from origin navaid, in degrees. Note thjat this will be the true bearing if the `TRUE_DEGREE` parameter is true (1). |
| `RHO` | FLOAT32 | The distance to the waypoint from origin navaid, in meters. |
| `COURSE` | FLOAT32 | The magnetic course to the waypoint from origin navaid, in degrees. Note thjat this will be the true course if the `TRUE_DEGREE` parameter is true (1). |
| `ROUTE_DISTANCE` | FLOAT32 | The route distance measured in either meters or minutes, depending on the `DISTANCE_MINUTE` value. |
| `APPROACH_ALT_DESC` | INT32 | Alternative description for the approach leg, which can be one of the following:<br>1. 0: NOT\_USED<br>2. 1: AT<br>3. 2: AT\_OR\_ABOVE<br>4. 3: AT\_OR\_BELOW<br>5. 4: IN\_BETWEEN |
| `ALTITUDE1` | FLOAT32 | Returns the leg altitude, in meters. This is only available when the `APPROACH_ALT_DESC` value is not 0. |
| `ALTITUDE2` | FLOAT32 | Returns the second leg altitude entry, in meters. This is only available when `APPROACH_ALT_DESC` is 4. |
| `SPEED_LIMIT` | FLOAT32 | The speed limit, in kias. |
| `VERTICAL_ANGLE` | FLOAT32 | The vertical angle, in degrees. |
| `ARC_CENTER_FIX_ICAO` | STRING8 | ICAO of the arc center fix (for RF legs) |
| `ARC_CENTER_FIX_REGION` | STRING8 | Region of the arc center fix (for RF legs) |
| `ARC_CENTER_FIX_TYPE` |  | The type of of arc center fix, one of the following:<br>1. 65 ('A') - Airport<br>2. 86 ('V') - VOR<br>3. 78 ('N') - NDB<br>4. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `ARC_CENTER_FIX_LATITUDE` | FLOAT64 | The latitude of the arc center fix (for RF legs) |
| `ARC_CENTER_FIX_LONGITUDE` | FLOAT64 | The longitude of the arc center fix (for RF legs) |
| `ARC_CENTER_FIX_ALTITUDE` | FLOAT64 | The altitude of the arc center fix (for RF legs) |
| `RADIUS` | FLOAT32 | The radius of the arc in meters. |
| `IS_IAF` | INT32 | Returns whether the waypoint defined by the `FIX_ICAO` entry is the Initial Approach Fix (1) or not (0). |
| `IS_IF` | INT32 | Returns whether the waypoint defined by the `FIX_ICAO` entry is the Intermediate Approach Fix (1) or not (0). |
| `IS_FAF` | INT32 | Returns whether the waypoint defined by the `FIX_ICAO` entry is the Final Approach Fix (1) or not (0). |
| `IS_MAP` | INT32 | Returns whether the waypoint defined by the `FIX_ICAO` entry is the Missed Approach Point (1) or not (0). |
| `REQUIRED_NAVIGATION_PERFORMANCE` | FLOAT32 | The required navigation performance for the leg in meters, or 0 if not defined. |
| `APPROACH_SPEED_DESC` | INT32 | The description of the SPEED\_LIMIT value, which can be one of the following:<br>0: NONE<br>1: AT<br>2: AT\_OR\_ABOVE<br>3: AT\_OR\_BELOW |

##### START

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `LATITUDE` | FLOAT64 | Latitude of the start position, in degrees between -90.0° and 90.0°. |
| `LONGITUDE` | FLOAT64 | Longitude of the start position, in degrees between -180.0° and 180.0°. |
| `ALTITUDE` | FLOAT64 | Altitude of the start position, in meters. |
| `HEADING` | FLOAT32 | Facing angle for the start position. Value between 0° and 360°. |
| `NUMBER` | INT32 | Number of the runway to start on. |
| `DESIGNATOR` | INT32 | Designator of the runway chosen to start on. |
| `TYPE` | INT32 | The type of starting point, which can be one of the following:<br>1. 0: UNKNOWN<br>2. 1: RUNWAY<br>3. 2: WATER<br>4. 3: HELIPAD<br>5. 4: TRACK |

#### RUNWAY

This is a child member of the `AIRPORT` entry point and you can request the following data from it:

| Member Name | DataType | Description |
| `LATITUDE` | FLOAT64 | The latitude of the center of the runway, in degrees. |
| `LONGITUDE` | FLOAT64 | The longitude of the center of the runway, in degrees. |
| `ALTITUDE` | FLOAT64 | The altitude of the center of the runway, in meters. |
| `HEADING` | FLOAT32 | The runway heading, in degrees. |
| `LENGTH` | FLOAT32 | The runway length, in Meters. Note that this includes offset thresholds, but does _not_ include blast pads or overruns. |
| `WIDTH` | FLOAT32 | The runway in meters. |
| `PATTERN_ALTITUDE` | FLOAT32 | The pattern altitude, in meters. |
| `SLOPE` | FLOAT32 | The runway slope, in degrees. |
| `TRUE_SLOPE` | FLOAT32 | The runway true slope, in degrees. |
| `SURFACE` | INT32 | The type of pavement used by the runway. The return value will be one of the following: |
| `PRIMARY_ILS_ICAO` | STRING8 | ILSICAO code for the primary end of the runway. |
| `PRIMARY_ILS_REGION` | STRING8 | ILS region for the primary end of the runway. |
| `PRIMARY_ILS_TYPE` | INT32 | The primary ILS type, one of the following:<br>1. 65 ('A') - Airport<br>2. 86 ('V') - VOR<br>3. 78 ('N') - NDB<br>4. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `PRIMARY_NUMBER` | INT32 | Number of the primary end of the runway, one of the following values:<br>01. 0: NONE<br>02. 1 - 36: RUNWAY ID<br>03. 37: NORTH<br>04. 38: NORTHEAST<br>05. 39: EAST<br>06. 40: SOUTHEAST<br>07. 41: SOUTH<br>08. 42: SOUTHWEST<br>09. 43: WEST<br>10. 44: NORTHWEST<br>11. 45: LAST |
| `PRIMARY_DESIGNATOR` | INT32 | Designator of the primary end of the runway, one of the following values:<br>1. 0: NONE<br>2. 1: LEFT<br>3. 2: RIGHT<br>4. 3: CENTER<br>5. 4: WATER<br>6. 5: A<br>7. 6: B<br>8. 7: LAST |
| `SECONDARY_ILS_ICAO` | STRING8 | ILSICAO code for the secondary end of the runway. |
| `SECONDARY_ILS_REGION` | STRING8 | ILS region for the secondary end of the runway. |
| `SECONDARY_ILS_TYPE` | INT32 | The secondary ILS type, one of the following:<br>1. 65 ('A') - Airport<br>2. 86 ('V') - VOR<br>3. 78 ('N') - NDB<br>4. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `SECONDARY_NUMBER` | INT32 | Number of the secondary end of the runway, one of the following values:<br>01. 0: NONE<br>02. 1 - 36: RUNWAY ID<br>03. 37: NORTH<br>04. 38: NORTHEAST<br>05. 39: EAST<br>06. 40: SOUTHEAST<br>07. 41: SOUTH<br>08. 42: SOUTHWEST<br>09. 43: WEST<br>10. 44: NORTHWEST<br>11. 45: LAST |
| `SECONDARY_DESIGNATOR` | INT32 | Designator of the secondary end of the runway, one of the following values:<br> <br>1. 0: NONE<br>2. 1: LEFT<br>3. 2: RIGHT<br>4. 3: CENTER<br>5. 4: WATER<br>6. 5: A<br>7. 6: B<br>8. 7: LAST |
| `EDGE_LIGHTS` | INT8 | Status of edge lights for the given runway, one of the following values :<br>1. 0: NONE<br>2. 1: LOW<br>3. 2: MEDIUM<br>4. 3: HIGH INTENSITY |
| `CENTER_LIGHTS` | INT8 | Status of center lights for the given runway, one of the following values :<br>1. 0: NONE<br>2. 1: LOW<br>3. 2: MEDIUM<br>4. 3: HIGH INTENSITY |
| `PRIMARY_CLOSED` | INT8 | Is the primary opened or closed. |
| `SECONDARY_CLOSED` | INT8 | Is the secondary opened or closed. |
| `PRIMARY_TAKEOFF` | INT8 | Is takeoff allowed from primary. |
| `PRIMARY_LANDING` | INT8 | Is landing allowed from primary. |
| `SECONDARY_TAKEOFF` | INT8 | Is takeoff allowed from secondary. |
| `SECONDARY_LANDING` | INT8 | Is landing allowed from secondary. |
| `PRIMARY_THRESHOLD` | STRUCT: `PAVEMENT` | Contains data on the runway primary threshold. |
| `PRIMARY_BLASTPAD` | STRUCT: `PAVEMENT` | Contains data on the runway primary blastpad. |
| `PRIMARY_OVERRUN` | STRUCT: `PAVEMENT` | Contains data on the runway primary overrun. |
| `PRIMARY_APPROACH_LIGHTS` | STRUCT: `APPROACHLIGHTS` | Contains data on the runway primary approach lights. |
| `PRIMARY_LEFT_VASI` | STRUCT: `VASI` | Contains data on the runway primary left VASI. |
| `PRIMARY_RIGHT_VASI` | STRUCT: `VASI` | Contains data on the runway primary right VASI. |
| `SECONDARY_THRESHOLD` | STRUCT: `PAVEMENT` | Contains data on the runway secondary threshold. |
| `SECONDARY_BLASTPAD` | STRUCT: `PAVEMENT` | Contains data on the runway secondary blastpad. |
| `SECONDARY_OVERRUN` | STRUCT: `PAVEMENT` | Contains data on the runway secondary overrun. |
| `SECONDARY_APPROACH_LIGHTS` | STRUCT: `APPROACHLIGHTS` | Contains data on the runway secondary approach lights. |
| `SECONDARY_LEFT_VASI` | STRUCT: `VASI` | Contains data on the runway secondary left VASI. |
| `SECONDARY_RIGHT_VASI` | STRUCT: `VASI` | Contains data on the runway secondary right VASI. |

##### PAVEMENT

This is a child member of the `RUNWAY` entry point - specifically from the `PRIMARY_THRESHOLD`, `PRIMARY_BLASTPAD`, `PRIMARY_OVERRUN`, `SECONDARY_THRESHOLD`, `SECONDARY_BLASTPAD`, and `SECONDARY_OVERRUN` members - and you can request the following data from it:

| Member Name | DataType | Description |
| `LENGTH` | FLOAT32 | The length of the pavement area (in meters) |
| `WIDTH` | FLOAT32 | The width of the pavement area (in meters) |
| `ENABLE` | INT32 | Whether the requested pavement area is actually enabled for the runway or not. |

##### APPROACHLIGHTS

This is a child member of the `RUNWAY` entry point - specifically from the `PRIMARY_APPROACH_LIGHTS` and `SECONDARY_APPROACH_LIGHTS` members - and you can request the following data from it:

| Member Name | DataType | Description |
| `SYSTEM` | INT32 | The type of approach light system, which can be one of the following:<br>01. 0: NONE<br>02. 1: ODALS<br>03. 2: MALSF<br>04. 3: MALSR<br>05. 4: SSALF<br>06. 5: SSALR<br>07. 6: ALSF1<br>08. 7: ALSF2<br>09. 8: RAIL<br>10. 9: CALVERT<br>11. 10: CALVERT2<br>12. 11: MALS<br>13. 12: SALS<br>14. 13: SALSF<br>15. 14: SSALS |
| `STROBE_COUNT` | INT32 | The number of sequenced strobes. |
| `HAS_END_LIGHTS` | INT32 | Whether the runway has end lights or not. |
| `HAS_REIL_LIGHTS` | INT32 | Whether the runway has reil lights or not. |
| `HAS_TOUCHDOWN_LIGHTS` | INT32 | Whether the runway has touchdown lights or not. |
| `ON_GROUND` | INT32 | Indicates whether the lights are snapped to fit the ground terrain or not. |
| `ENABLE` | INT32 | Whether the approach lights are enabled or not. |
| `OFFSET` | FLOAT32 | The offset for the lights in meters. |
| `SPACING` | FLOAT32 | The spacing between the lights, in meters. |
| `SLOPE` | FLOAT32 | The slope of the approach lights expressed in degrees. |

##### VASI

This is a child member of the `RUNWAY` entry point - specifically from the `PRIMARY_LEFT_VASI`, `PRIMARY_RIGHT_VASI`, `SECONDARY_LEFT_VASI`, and `SECONDARY_RIGHT_VASI` members - and you can request the following data from it:

| Member Name | DataType | Description |
| `TYPE` | INT32 | The type of VASI being used, which can be one of the following:<br>01. 0: NONE<br>02. 1: VASI21 - 2 rows, 1 box/row<br>03. 2: VASI22 - 2 rows, 2 boxes/row<br>04. 3: VASI23 - 2 rows, 3 boxes/row<br>05. 4: VASI31 - 3 rows, 1 box/row<br>06. 5: VASI32 - 3 rows, 2 boxes/row<br>07. 6: VASI33 - 3 rows, 3 boxes/row (far row has only two boxes, 8 light system)<br>08. 7: PAPI2 - 2 light PAPI<br>09. 8: PAPI4 - 4 light PAPI<br>10. 9: TRICOLOR - Tri Color VASI<br>11. 10: PVASI - Pulsating VASI<br>12. 11: TVASI - colored VASI in a shared-"T" shape<br>13. 12: BALL - presently remapped to PVASI<br>14. 13: APAP - panels |
| `BIAS_X` | FLOAT32 | Distance from the runway center-line across the runway width to the reference point of the VASI, in meters. |
| `BIAS_Z` | FLOAT32 | Distance along the runway from the runway center point to the VASI reference point, in meters. |
| `SPACING` | FLOAT32 | Distance between light rows, in meters. Note that this is only applicable to the following types: `VASI21`, `VASI31`, `VASI22`, `VASI32`, `VASI23`, `VASI33`, and `TVASI`. |
| `ANGLE` | FLOAT32 | The approach angle, in degrees, of the VASI. |

### WAYPOINT

The waypoint entry point can have the `ROUTE` child member and you can request the following information directly:

| Member Name | DataType | Description |
| `LATITUDE` | FLOAT64 | The waypoint latitude, in degrees. |
| `LONGITUDE` | FLOAT64 | The waypoint longitude, in degrees. |
| `ALTITUDE` | FLOAT64 | The waypoint altitude, in meters. |
| `TYPE` | INT32 | The type of waypoint, which can be any one of the following values:<br>01. 0: NONE<br>02. 1: NAMED<br>03. 2: UNNAMED<br>04. 3: VOR<br>05. 4: NDB<br>06. 5: OFFROUTE<br>07. 6: IAF<br>08. 7: FAF<br>09. 8: RNAV<br>10. 9: VFR |
| `MAGVAR` | FLOAT32 | This is the magnetic variation for the waypoint position. |
| `N_ROUTES` | INT32 | The number of [ROUTE](SimConnect_AddToFacilityDefinition.htm#route) s for the waypoint. |
| `ICAO` | STRING8 | The ICAO code for the waypoint. |
| `REGION` | STRING8 | The region (if available) for the waypoint. WIll be "" if no region is specified. |
| `IS_TERMINAL_WPT` | INT32 | Will be 1 if the waypoint is a terminal waypoint, or 0 otherwise. |
| `ROUTE` | STRUCT: `ROUTE` | Contains data on the runway primary threshold. |

#### ROUTE

This is a child member of the `WAYPOINT` entry point and you can request the following data from it:

| Member Name | DataType |  |
| `NAME` | STRING32 | The name of the route. |
| `TYPE` | INT32 | The airway type, which can be any one of the following:<br>1. 0: NONE<br>2. 1: VICTOR<br>3. 2: JET<br>4. 3: BOTH |
| `NEXT_ICAO` | STRING8 | ICAO of the next waypoint of the route. |
| `NEXT_REGION` | STRING8 | Region of the next waypoint of the route. Will be "" if the region isn't specified. |
| `NEXT_TYPE` | INT32 | The next waypoint type which can be any one of the following:<br>1. 86 ('V') - VOR<br>2. 78 ('N') - NDB<br>3. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `NEXT_LATITUDE` | FLOAT64 | The latitude of the next waypoint. |
| `NEXT_LONGITUDE` | FLOAT64 | The longitude of the next waypoint. |
| `NEXT_ALTITUDE` | FLOAT32 | Minimum altitude, in meters, of the next route point. |
| `PREV_ICAO` | STRING8 | ICAO of the previous point of the route. |
| `PREV_REGION` | STRING8 | Region of the previous point of the route. Will be "" if the region isn't specified. |
| `PREV_TYPE` | INT32 | The previous waypoint type which can be any one of the following ASCii values:<br>1. 86 ('V') - VOR<br>2. 78 ('N') - NDB<br>3. 87 ('W') - Waypoint<br>(Note that this will have to be cast as a char) |
| `NEXT_LATITUDE` | FLOAT64 | The latitude of the previous waypoint. |
| `NEXT_LONGITUDE` | FLOAT64 | The longitude of the previous waypoint. |
| `PREV_ALTITUDE` | FLOAT32 | Minimum altitude, in meters, of the previous route point. |

### NDB

The NDB entry point can request the following information:

| Member Name | DataType | Description |
| `LATITUDE` | FLOAT64 | The latitude, in degrees, for the NDB transmitter. |
| `LONGITUDE` | FLOAT64 | The longitude, in degrees, for the NDB transmitter. |
| `ALTITUDE` | FLOAT64 | The altitude, in meters, for the NDB transmitter. |
| `FREQUENCY` | UINT32 | The frequency of the transmitter. |
| `TYPE` | INT32 | The type of transmitter, which can be one of the following values:<br>1. 0: Compass Locator<br>2. 1: Medium Homing<br>3. 2: Homing<br>4. 3: High Homing |
| `RANGE` | FLOAT32 | The station transmitter range, in meters. |
| `MAGVAR` | FLOAT32 | This is the magnetic variation for the station. |
| `IS_TERMINAL_NDB` | INT32 | Returns (1) if the station is a terminal NDB, or (0) otherwise. |
| `NAME` | STRING64 | The name of the NDB transmitter station. |
| `BFO_REQUIRED` | INT32 | Returns whether the NDB audio signal should be received with a beat frequency oscillator \[1\] or not \[0\]. |

### VOR

The VOR entry point can request the following information:

| Member Name | DataType | Description |
| `VOR_LATITUDE` | FLOAT64 | The latitude, in degrees, for the VOR transmitter. |
| `VOR_LONGITUDE` | FLOAT64 | The longitude, in degrees, for the VOR transmitter. |
| `VOR_ALTITUDE` | FLOAT64 | The altitude, in meters, for the VOR transmitter. |
| `DME_LATITUDE` | FLOAT64 | The latitude, in degrees, for the DME beacon. |
| `DME_LONGITUDE` | FLOAT64 | The longitude, in degrees, for the DME beacon. |
| `DME_ALTITUDE` | FLOAT64 | The altitude, in meters, for the DME beacon. |
| `GS_LATITUDE` | FLOAT64 | The latitude, in degrees, for the glideslope. |
| `GS_LONGITUDE` | FLOAT64 | The longitude, in degrees, for the glideslope. |
| `GS_ALTITUDE` | FLOAT64 | The altitude, in meters, for the glideslope. |
| `TACAN_LATITUDE` | FLOAT64 | The latitude, in degrees, for the Tacan system. |
| `TACAN_LONGITUDE` | FLOAT64 | The longitude, in degrees, for the Tacan system. |
| `TACAN_ALTITUDE` | FLOAT64 | The altitude, in meters, for the Tacan system. |
| `IS_NAV` | INT32 | Returns whether the beacon is a Nav beacon (1) or not (0). |
| `IS_DME` | INT32 | Returns whether the beacon is a DME beacon (1) or not (0). |
| `IS_TACAN` | INT32 | Returns whether the beacon is a TACAN transmitter (1) or not (0). |
| `HAS_GLIDE_SLOPE` | INT32 | Returns whether the beacon has a glideslope (1) or not (0). |
| `DME_AT_NAV` | INT32 | Returns true (1) when the DME is co-located with the VOR station, and returns false (0) when it's not. |
| `DME_AT_GLIDE_SLOPE` | INT32 | Returns true (1) when the DME is co-located with the Glideslope or Localizer, and returns false (0) when it's not. |
| `HAS_BACK_COURSE` | INT32 | Returns whether the beacon has a back course (1) or not (0). |
| `FREQUENCY` | UINT32 | The frequency for the station. |
| `TYPE` | INT32 | The VOR type, which can be one of the following values:<br>1. 0: VOR\_UNKNOWN<br>2. 1: VOR\_TERMINAL<br>3. 2: VOR\_LOW\_ALTITUDE / VOR\_LOW\_ALT<br>4. 3: VOR\_HIGH ALTITUDE / VOR\_HIGH\_ALT<br>5. 4: VOR\_ILS<br>6. 5: VOR\_VOT |
| `NAV_RANGE` | FLOAT32 | The station range, in meters. |
| `MAGVAR` | FLOAT32 | This is the magnetic variation for the station. |
| `LOCALIZER` | FLOAT32 | The localizer heading, in degrees. |
| `LOCALIZER_WIDTH` | FLOAT32 | The localizer beam width, in degrees. |
| `GLIDE_SLOPE` | FLOAT32 | The glide slope, in degrees. |
| `NAME` | STRING64 | The VOR station name. |
| `DME_BIAS` | FLOAT32 | The bias, in meters, to be subtracted from DME measurements. |
| `LS_CATEGORY` | INT32 | The category of the landing system, which can be on of:<br>01. 0: NONE<br>02. 1: CAT1<br>03. 2: CAT2<br>04. 3: CAT3<br>05. 4: LOCALIZER<br>06. 5: IGS<br>07. 6: LDA\_NO\_GS<br>08. 7: LDA\_WITH\_GS<br>09. 8: SDF\_NO\_GS<br>10. 9: SDF\_WITH\_GS |
| `IS_TRUE_REFERENCED` | INT32 | Returns whether the facility is referenced to true north \[1\] or not \[0\]. |

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