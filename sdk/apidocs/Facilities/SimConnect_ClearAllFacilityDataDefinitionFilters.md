SimConnect\_ClearAllFacilityDataDefinitionFilters

## SimConnect\_ClearAllFacilityDataDefinitionFilters

The **SimConnect\_ClearAllFacilityDataDefinitionFilters** function is used to clear all applied facility definition filters.

##### Syntax

```cpp
HRESULT SimConnect_ClearAllFacilityDataDefinitionFilters(
    HANDLE hSimConnect,
    SIMCONNECT_DATA_DEFINITION_ID DefineID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `DefineID` | Specifies the ID of the client defined data definition. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| `S_OK` | The function succeeded. |
| `E_FAIL` | The function failed. |

This function might throw the following [exception](../Structures_And_Enumerations/SIMCONNECT_EXCEPTION.htm):

- `SIMCONNECT_EXCEPTION_UNRECOGNIZED_ID` (offset == 1) - If DefineID is not recognized

##### Remarks

Please see `SimConnect_AddFacilityDataDefinitionFilter` for details on creating filters, and see the accompanying examples for details on how to clear a single filter.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_RequestFacilityData](SimConnect_RequestFacilityData.htm)
4. [SimConnect\_AddFacilityDataDefinitionFilter](SimConnect_AddFacilityDataDefinitionFilter.htm)
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