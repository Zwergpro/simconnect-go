SimConnect\_RequestFacilitesList

## SimConnect\_RequestFacilitesList

The **SimConnect\_RequestFacilitesList** function is used to request a list of all the facilities of a given type currently in the world, _or_ within a radius of the aircraft depending on the requested type.

##### Syntax

```cpp
HRESULT SimConnect_RequestFacilitesList(
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

The simulation keeps a facilities cache of all the airports within the world, and all the waypoints, NDB and VOR stations within a certain radius of the aircraft. This radius varies depending on where the aircraft is in the world, but is at least large enough to encompass the whole of the reality bubble for waypoints, and can be over 200 miles for VOR and NDB stations. As the user aircraft moves, waypoints, NDB and VOR stations will be added to and removed from the cache, however no changes will be made to airports as they are all already in the list. In the interests pf performance, hysteresis is built into the system.

To receive event notifications when a facility is added, use the [`SimConnect_SubscribeToFacilities`](SimConnect_SubscribeToFacilities.htm) function. When this function is first called, a full list from the world will be sent, thereafter just the additions will be transmitted. To terminate these notifications use the [`SimConnect_UnsubscribeToFacilities`](SimConnect_UnsubscribeToFacilities.htm) function.

When requesting types of facility information, one function call has to be made for each of the four types of data. The data will be returned in one of the four structures:

- [`SIMCONNECT_RECV_AIRPORT_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_AIRPORT_LIST.htm), which will contain a list of [`SIMCONNECT_DATA_FACILITY_AIRPORT`](../Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_AIRPORT.htm) structures, for the _entire_ world, approximately 40,000 items. To get only those airports present within the reality bubble cache, use the [`SimConnect_RequestFacilitiesList_EX1`](SimConnect_RequestFacilitiesList_EX1.htm) function.
- [`SIMCONNECT_RECV_NDB_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_NDB_LIST.htm), which will contain a list of [`SIMCONNECT_DATA_FACILITY_NDB`](../Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_NDB.htm) structures.
- [`SIMCONNECT_RECV_VOR_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_VOR_LIST.htm), which will contain a list of [`SIMCONNECT_DATA_FACILITY_VOR`](../Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_VOR.htm) structures.
- [`SIMCONNECT_RECV_WAYPOINT_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_WAYPOINT_LIST.htm), which will contain a list of [`SIMCONNECT_DATA_FACILITY_WAYPOINT`](../Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_WAYPOINT.htm) structures.

The four list structures inherit the data from the [`SIMCONNECT_RECV_FACILITIES_LIST`](../Structures_And_Enumerations/SIMCONNECT_RECV_FACILITIES_LIST.htm) structure. Given that the list of returned facilities could be large, it may be split across several packets, and each packet must be interpreted separately by the client.

##### See Also

- [SimConnect API Reference](../../SimConnect_API_Reference.htm)
- [SimConnect\_SubscribeToFacilities](SimConnect_SubscribeToFacilities.htm)
- [SimConnect\_UnsubscribeToFacilities](SimConnect_UnsubscribeToFacilities.htm)

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