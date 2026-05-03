SimConnect\_RequestFacilityData\_EX1

## SimConnect\_RequestFacilityData\_EX1

The **SimConnect\_RequestFacilityData\_EX1** function is used to request data according to a predefined object, an ICAO and a region. This function is practically identical in functionality to the [SimConnect\_RequestFacilityData](SimConnect_RequestFacilityData.htm) function, only it has an additional return value used to identify waypoints when there is an ICAO/Region overlap with VOR or NDB.

##### Syntax

```cpp
HRESULT SimConnect_RequestFacilityData(
    HANDLE hSimConnect,
    SIMCONNECT_DATA_DEFINITION_ID DefineID,
    SIMCONNECT_DATA_REQUEST_ID RequestID,
    const char * ICAO,
    const char * Region = "",
    char Type = 0
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `DefineID` | Specifies the ID of the client defined data definition | Integer |
| `RequestID` | The client defined request ID | Integer |
| `ICAO` | Used to identify an airport, a VOR, an NDB or a waypoint. | Char |
| `Region` | Additional identifier for an airport, a VOR, an NDB or a waypoint. For airports, this can be omitted without issue, however for VOR / NDB / Waypoints this should be supplied if possible, although there are workarounds provided if it's not (see remarks, below). | Char |
| `Type` | Additional identifier for when requesting data to differentciate between waypoint/VOR/NDB when there are overlapping ICAO/Region identifiers. Should be one of the following ASCii values:<br>1. 86 ('V') - VOR<br>2. 78 ('N') - NDB<br>3. 87 ('W') - Waypoint | Char |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| `S_OK` | The function succeeded. |
| `E_FAIL` | The function failed. |
| `SIMCONNECT_EXCEPTION_UNRECOGNIZED_ID` | The function has an invalid `DefineId`. |
| `SIMCONNECT_EXCEPTION_ERROR` | The function was supplied an invalid `ICAO` and/or `Region` parameter. |

##### Remarks

If the Request works, the client will receive as many messages ( [`SIMCONNECT_RECV_FACILITY_DATA`](../Structures_And_Enumerations/SIMCONNECT_RECV_FACILITY_DATA.htm)) as necessary to match the defined architecture. At the end, the server will send a [`SIMCONNECT_RECV_FACILITY_DATA_END`](../Structures_And_Enumerations/SIMCONNECT_RECV_FACILITY_DATA_END.htm) message to let the client know that everything is finished. For an example of use and more information, please see [`SimConnect_AddToFacilityDefinition`](SimConnect_AddToFacilityDefinition.htm).

Note that it is possible to request a Waypoint/NDB/VOR using only the ICAO and no region. When doing so, if there is _no_ duplicated ICAO, then the appropriate data will be returned. However, if there is a conflict between multiple ICAO's, the function will return only those that correspond to the given "Type".

##### See Also

- [SimConnect API Reference](../../SimConnect_API_Reference.htm)
- [SimConnect\_AddToFacilityDefinition](SimConnect_AddToFacilityDefinition.htm)
- [SimConnect\_RequestFacilitiesList](SimConnect_RequestFacilitesList.htm)
- [SimConnect\_RequestFacilityData](SimConnect_RequestFacilityData.htm)

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