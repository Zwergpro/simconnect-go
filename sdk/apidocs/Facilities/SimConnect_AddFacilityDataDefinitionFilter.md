SimConnect\_AddFacilityDataDefinitionFilter

## SimConnect\_AddFacilityDataDefinitionFilter

The **SimConnect\_AddFacilityDataDefinitionFilter** function is used add a filter on a node in the FacilityDataDefinition to block sending data according to this filter, thus reduce the amount of data received and limit it to only that which is required.

##### Syntax

```cpp
HRESULT SimConnect_AddFacilityDataDefinitionFilter(
    HANDLE hSimConnect,
    SIMCONNECT_DATA_DEFINITION_ID DefineID,
    const char * szFilterPath,
    DWORD cbFilterDataSize,
    void * pFilterData
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `DefineID` | Specifies the ID of the client defined data definition. | Integer |
| `szFilterPath` | Defines the node and member that you wish to apply the filter to. | Integer |
| `cbFilterDataSize` | The size of the pFilterData. | Integer |
| `pFilterData` | Filter data as bytes (will be cast to the right type later). | Pointer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| `S_OK` | The function succeeded. |
| `E_FAIL` | The function failed. |

This function might throw the following [exceptions](../Structures_And_Enumerations/SIMCONNECT_EXCEPTION.htm):

- `SIMCONNECT_EXCEPTION_UNRECOGNIZED_ID` (offset == 1) - if the DefineID is not recognized
- `SIMCONNECT_EXCEPTION_UNRECOGNIZED_ID` (offset == 2) - if the DefineID is valid but refers to another facility data definition type (Airport, Waypoint, Vor…)
- `SIMCONNECT_EXCEPTION_DATA_ERROR` (offset == 2) - if the given path is incorrect

##### Remarks

Please see `SimConnect_AddToFacilityDefinition` for a list of all the possible nodes that can be filtered.

##### Examples

This first example applies a filter on the member `PRIMARY_NUMBER` of the runway found in `AIRPORT` (id = `FACILITY_DATA_DEF_AIRPORT`). `PRIMARY_NUMBER` is a `SIMCONNECT_DATATYPE_INT32` variable, so `pFilterData` must be an `INT32` variable too:

```cpp
unsigned var = 27;
SimConnect_AddFacilityDataDefinitionFilter(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "AIRPORT:RUNWAY:PRIMARY_NUMBER", sizeof(unsigned), &var);
```

The following example applies a filter on the member `SECONDARY_ILS_ICAO` of the runway contained in `AIRPORT` (id = `FACILITY_DATA_DEF_AIRPORT`). `SECONDARY_ILS_ICAO` is a `SIMCONNECT_DATATYPE_STRING8` variable, so `pFilterData` must be a `char[8]` variable too:

```cpp
char icao[8] = "ToT1";
SimConnect_AddFacilityDataDefinitionFilter(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "AIRPORT:RUNWAY:SECONDARY_ILS_ICAO", sizeof(icao), &icao);
```

This next example applies a filter on the member `LATITUDE` of the runway contained in `AIRPORT` (id = `FACILITY_DATA_DEF_AIRPORT`. Note that this is potentially dangerous because of double/float approximation, but it can work:

```cpp
double d = 123.456;
SimConnect_AddFacilityDataDefinitionFilter(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "AIRPORT:RUNWAY:LATITUDE", sizeof(double), &d);
```

Finally, this example shows how you can remove a specific filter by giving a `nullptr` value. It removes the filter on the `PRIMARY_NUMBER` member:

```cpp
SimConnect_AddFacilityDataDefinitionFilter(hSimConnect, FACILITY_DATA_DEF_AIRPORT, "AIRPORT:RUNWAY:PRIMARY_NUMBER", 0, nullptr);
```

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_ClearAllFacilityDataDefinitionFilters](SimConnect_ClearAllFacilityDataDefinitionFilters.htm)
4. [SimConnect\_AddToFacilityDefinition](SimConnect_AddToFacilityDefinition.htm)
5. [SimConnect\_RequestFacilityData](SimConnect_RequestFacilityData.htm)
6. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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