SIMCONNECT\_RECV\_FACILITIES\_LIST

## SIMCONNECT\_RECV\_FACILITIES\_LIST

The **SIMCONNECT\_RECV\_FACILITIES\_LIST** structure is used to provide information on the number of elements in a list of facilities returned to the client, and the number of packets that were used to transmit the data.

##### Syntax

```cpp
struct SIMCONNECT_RECV_FACILITIES_LIST : public SIMCONNECT_RECV{
    DWORD  dwRequestID;
    DWORD  dwArraySize;
    DWORD  dwEntryNumber;
    DWORD  dwOutOf;
    };
```

##### Members

| Member | Description |
| `dwRequestID` | Double word containing the client defined request ID. |
| `dwArraySize` | Double word containing the number of elements in the list that are within this packet. For example, if there are 25 airports returned in the [SIMCONNECT\_RECV\_AIRPORT\_LIST](SIMCONNECT_RECV_AIRPORT_LIST.htm) structure, then this field will contain 25, but if there are 400 airports in the list and the data is returned in two packets, then this value will contain the number of entries within each packet. |
| `dwEntryNumber` | Double word containing the index number of this list packet. This number will be from 0 to dwOutOf - 1. |
| `dwOutOf` | Double word containing the total number of packets used to transmit the list. |

##### Remarks

This structure is inherited by [`SIMCONNECT_RECV_AIRPORT_LIST`](SIMCONNECT_RECV_AIRPORT_LIST.htm), [`SIMCONNECT_RECV_NDB_LIST`](SIMCONNECT_RECV_NDB_LIST.htm), [`SIMCONNECT_RECV_VOR_LIST`](SIMCONNECT_RECV_VOR_LIST.htm) and [`SIMCONNECT_RECV_WAYPOINT_LIST`](SIMCONNECT_RECV_WAYPOINT_LIST.htm).

This structure inherits the [`SIMCONNECT_RECV`](SIMCONNECT_RECV.htm) structure, so use the [`SIMCONNECT_RECV_ID`](SIMCONNECT_RECV_ID.htm) enumeration to determine which list structure has been received.

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