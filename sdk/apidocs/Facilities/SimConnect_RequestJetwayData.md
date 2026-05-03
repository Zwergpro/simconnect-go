SimConnect\_RequestJetwayData

## SimConnect\_RequestJetwayData

The **SimConnect\_RequestJetwayData** function is used to request data from one or more jetways.

##### Syntax

```cpp
HRESULT SimConnect_RequestJetwayData(
    HANDLE hSimConnect,
    const char* szAirportIcao,
    DWORD dwArrayCount,
    int* pIndexes
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `szAirportIcao` | The airport ICAO to check. | Integer |
| `dwArrayCount` | This is the number of elements in `pIndexes`. Can be 0. | Integer |
| `pIndexes` | An array of parking indices. Can be `nullptr`. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| `S_OK` | The function succeeded. |
| `E_FAIL` | The function failed. |
| `SIMCONNECT_EXCEPTION_JETWAY_DATA` | The function has thrown an exception where:<br>1. dwIndex = 1:<br>   1. ICAO is wrong<br>   2. Airport is not spawned<br>2. dwIndex = 2:<br>   1. At least one of the given indices is wrong<br>3. dwIndex = 99:<br>   1. Internal error |

##### Remarks

If the Request works, the client will receive the [`SIMCONNECT_RECV_JETWAY_DATA`](../Structures_And_Enumerations/SIMCONNECT_RECV_JETWAY_DATA.htm) structure with an array of [`SIMCONNECT_JETWAY_DATA`](../Structures_And_Enumerations/SIMCONNECT_JETWAY_DATA.htm) structs containing the jetway data.

##### See Also

- [SimConnect API Reference](../../SimConnect_API_Reference.htm)
- [SIMCONNECT\_RECV\_JETWAY\_DATA](../Structures_And_Enumerations/SIMCONNECT_RECV_JETWAY_DATA.htm)

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