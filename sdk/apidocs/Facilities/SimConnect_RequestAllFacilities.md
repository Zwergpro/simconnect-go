SimConnect\_RequestAllFacilities

## SimConnect\_RequestAllFacilites

The **SimConnect\_RequestAllFacilites** function is used to request a list of all the facilities of a given type currently in the world depending on the requested type.

##### Syntax

```cpp
HRESULT SimConnect_RequestAllFacilites(
    HANDLE  hSimConnect,
    SIMCONNECT_FACILITY_LIST_TYPE  type,
    SIMCONNECT_DATA_REQUEST_ID  RequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| hSimConnect | Handle to a SimConnect object. | Integer |
| type | Specifies one member of the [`SIMCONNECT_FACILITY_LIST_TYPE`](../Structures_And_Enumerations/SIMCONNECT_FACILITY_LIST_TYPE.htm) enumeration type. | Enum |
| RequestID | Specifies the client defined request ID. This will be returned along with the data. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

This function returns the entire list of facilities, no matter the distance to each elements. To retrieve information about facilities around you, use functions such as `SimConnect_RequestFacilitiesList_EX1` or `SimConnect_SubscribeToFacilities_EX1`.

When requesting types of facility information, one function call has to be made for each of the four types of data. The data will be returned in one of the four structures:

- [`SIMCONNECT_RECV_AIRPORT_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_AIRPORT_LIST.htm), which will contain a list of [`SIMCONNECT_DATA_FACILITY_AIRPORT`](../Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_AIRPORT.htm) structures.
- [`SIMCONNECT_RECV_NDB_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_NDB_LIST.htm), which will contain a list of [`SIMCONNECT_DATA_FACILITY_NDB`](../Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_NDB.htm) structures.
- [`SIMCONNECT_RECV_VOR_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_VOR_LIST.htm), which will contain a list of [`SIMCONNECT_DATA_FACILITY_VOR`](../Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_VOR.htm) structures.
- [`SIMCONNECT_RECV_WAYPOINT_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_WAYPOINT_LIST.htm), which will contain a list of [`SIMCONNECT_DATA_FACILITY_WAYPOINT`](../Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_WAYPOINT.htm) structures.

The four list structures inherit the data from the [`SIMCONNECT_RECV_FACILITIES_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_FACILITIES_LIST.htm) structure. Given that the list of returned facilities could be large, it may be split across several packets, and each packet must be interpreted separately by the client.

##### See Also

- [SimConnect API Reference](../../SimConnect_API_Reference.htm)

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