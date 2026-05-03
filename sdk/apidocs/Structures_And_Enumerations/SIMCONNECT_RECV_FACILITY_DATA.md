SIMCONNECT\_RECV\_FACILITY\_DATA

## SIMCONNECT\_RECV\_FACILITY\_DATA

The **SIMCONNECT\_RECV\_FACILITY\_DATA** structure is used to provide information that has been requested from the server using the `SimConnect_RequestFacilityData` function. This struct may be received multiple times before receiving `SIMCONNECT_RECV_FACILITY_DATA_END`.

##### Syntax

```cpp
struct SIMCONNECT_RECV_FACILITY_DATA : public SIMCONNECT_RECV{
    DWORD UserRequestId;
    DWORD UniqueRequestId;
    DWORD ParentUniqueRequestId;
    SIMCONNECT_FACILITY_DATA_TYPE Type;
    bool IsListItem;
    DWORD ItemIndex;
    DWORD ListSize;
    DWORD Data;
    };
```

##### Members

| Member | Description |
| `UserRequestId` | Double word containing the client defined request ID. |
| `UniqueRequestId` | The unique request ID, so the client can identify it. |
| `ParentUniqueRequestId` | If the current message is about a child object, this field will contain the parent's `UniqueRequestId`, otherwise it will be 0. |
| `Type` | Specifies the type of the object, will be a value from the [`SIMCONNECT_FACILITY_DATA_TYPE`](SIMCONNECT_FACILITY_DATA_TYPE.htm) enum. |
| `IsListItem` | If the current message is about a child object, this specifies if it is an orphan object or not. |
| `ItemIndex` | If `IsListItem` is true then this specifies the index in the list. |
| `ListSize` | If `IsListItem` is true, then this specifies the list size. |
| `Data` | Buffer of data. Have to cast it to a struct which matches the definition. |

##### Remarks

This structure inherits the [`SIMCONNECT_RECV`](SIMCONNECT_RECV.htm) structure, so use the [`SIMCONNECT_RECV_ID`](SIMCONNECT_RECV_ID.htm) enumeration to determine which list structure has been received.

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SIMCONNECT\_RECV\_ID](SIMCONNECT_RECV_ID.htm)
3. [SimConnect\_RequestFacilityData](../Facilities/SimConnect_RequestFacilityData.htm)
4. [SIMCONNECT\_FACILITY\_DATA\_TYPE](SIMCONNECT_FACILITY_DATA_TYPE.htm)
5. [SIMCONNECT\_RECV\_FACILITY\_DATA\_END](SIMCONNECT_RECV_FACILITY_DATA_END.htm)
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