SIMCONNECT\_RECV\_NDB\_LIST

## SIMCONNECT\_RECV\_NDB\_LIST

The **SIMCONNECT\_RECV\_NDB\_LIST** structure is used to return a list of [SIMCONNECT\_DATA\_FACILITY\_NDB](SIMCONNECT_DATA_FACILITY_NDB.htm) structures.

##### Syntax

```cpp
struct SIMCONNECT_RECV_NDB_LIST : public SIMCONNECT_RECV_FACILITES_LIST{
    SIMCONNECT_DATA_FACILITY_NDB  rgData[dwArraySize];
    };
```

##### Members

| Member | Description |
| `rgData[dwArraySize]` | Array of SIMCONNECT\_DATA\_FACILITY\_NDB structures. |

##### Remarks

This structure inherits the [`SIMCONNECT_RECV_FACILITIES_LIST`](SIMCONNECT_RECV_FACILITIES_LIST.htm) structure, which identifies the number of elements in the list, and the number of packets needed to transmit all the data.

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