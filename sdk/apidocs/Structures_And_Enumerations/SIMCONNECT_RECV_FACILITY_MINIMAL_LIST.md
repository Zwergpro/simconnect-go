SIMCONNECT\_RECV\_FACILITY\_MINIMAL\_LIST

## SIMCONNECT\_RECV\_FACILITY\_MINIMAL\_LIST

The **SIMCONNECT\_RECV\_FACILITY\_MINIMAL\_LIST** structure is used to provide minimal information on the number of elements in a list of facilities returned to the client, and the number of packets that were used to transmit the data.

##### Syntax

```cpp
struct SIMCONNECT_RECV_FACILITY_MINIMAL_LIST : public SIMCONNECT_RECV {
    DWORD   RequestID;
    DWORD   ArraySize;
    DWORD   EntryNumber;
    DWORD   OutOf;
    SIMCONNECT_FACILITY_MINIMAL rgData[dwArraySize]
    };
```

##### Members

| Member | Description |
| `RequestID` | ID of the request. |
| `ArraySize` | Number of elements in the `Data` struct. |
| `EntryNumber` | Current page index, from 0 to `OutOf`-1\. Mutltiple pages may be required when the `Data` member has too much information for a single return. |
| `OutOf` | The total number of pages that this data forms a part of. |
| `rgData[dwArraySize]` | Array of [`SIMCONNECT_FACILITY_MINIMAL`](SIMCONNECT_FACILITY_MINIMAL.htm) structures. |

##### Remarks

This structure inherits the [`SIMCONNECT_RECV_FACILITIES_LIST`](SIMCONNECT_RECV_LIST_TEMPLATE.htm) structure, which identifies the number of elements in the list, and the number of packets needed to transmit all the data.

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