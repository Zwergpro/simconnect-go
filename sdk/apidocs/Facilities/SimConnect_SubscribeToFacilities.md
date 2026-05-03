SimConnect\_SubscribeToFacilities

## SimConnect\_SubscribeToFacilities

The **SimConnect\_SubscribeToFacilities** function is used to request notifications when a facility of a certain type is added to the facilities cache.

##### Syntax

```cpp
HRESULT SimConnect_SubscribeToFacilities(
    HANDLE  hSimConnect,
    SIMCONNECT_FACILITY_LIST_TYPE  type,
    SIMCONNECT_DATA_REQUEST_ID  RequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `type` | Specifies one member of the `SIMCONNECT_FACILITY_LIST_TYPE` enumeration type. | Enum |
| `RequestID` | Specifies the client defined request ID. This will be returned along with the data. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| `S_OK` | The function succeeded. |
| `E_FAIL` | The function failed. |

##### Remarks

The simulation keeps a facilities cache of all the waypoints, NDB and VOR stations within a certain radius of the user aircraft, and - in the case of airports - in the entire world.

**NOTE**: If you wish to find only airports that are cached as being in the reality bubble, you should use the `SimConnect_SubscribeToFacilities_EX1` function. That function can also be used should you wish to set a callback for subscribed elementsd entering/leaving the reality bubble.

This radius varies slightly depending on where the aircraft is in the world, but is at least large enough to encompass the whole of the reality bubble for waypoints, and can be over 200 miles for VOR and NDB stations. As the user aircraft moves facilities will be added to, and removed from, the cache. However, in the interests pf performance, hysteresis is built into the system.

**NOTE**: The "reality bubble" is approximately 200,000 meters around the aircraft.

To receive event notifications when a facility is added, use the **SimConnect\_SubscribeToFacilities** function. When this function is first called, a full list from the cache will be sent, thereafter just the additions will be transmitted. No notification is given when a facility is removed from the cache. the Obviously to terminate these notifications use the `SimConnect_UnsubscribeToFacilities` function.

When requesting types of facility information, one function call has to be made for each of the four types of data. The data will be returned in one of the four structures:

- `SIMCONNECT_RECV_AIRPORT_LIST`, which will contain a list of `SIMCONNECT_DATA_FACILITY_AIRPORT` structures.
- `SIMCONNECT_RECV_NDB_LIST`, which will contain a list of `SIMCONNECT_DATA_FACILITY_NDB` structures.
- `SIMCONNECT_RECV_VOR_LIST`, which will contain a list of `SIMCONNECT_DATA_FACILITY_VOR` structures.
- `SIMCONNECT_RECV_WAYPOINT_LIST`, which will contain a list of `SIMCONNECT_DATA_FACILITY_WAYPOINT` structures.

The four list structures inherit the data from the `SIMCONNECT_RECV_FACILITIES_LIST` structure. Given that the list of returned facilities could be large, it may be split across several packets, and each packet must be interpreted separately by the client.

##### See Also

- [SimConnect API Reference](../../SimConnect_API_Reference.htm)
- [SimConnect\_UnsubscribeToFacilities](SimConnect_UnsubscribeToFacilities.htm)
- [SimConnect\_RequestFacilitiesList](SimConnect_SubscribeToFacilities.htm)
- [SimConnect\_SubscribeToFacilities\_EX1](SimConnect_SubscribeToFacilities_EX1.htm)

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