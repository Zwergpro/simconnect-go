SIMCONNECT\_RECV\_LIST\_TEMPLATE

## SIMCONNECT\_RECV\_LIST\_TEMPLATE

The **SIMCONNECT\_RECV\_LIST\_TEMPLATE** structure is used to provide information on the number of elements in a list returned to the client, and the number of packets that were used to transmit the data.

##### Syntax

```cpp
struct SIMCONNECT_RECV_LIST_TEMPLATE : public SIMCONNECT_RECV{
    DWORD  dwRequestID;
    DWORD  dwArraySize;
    DWORD  dwEntryNumber;
    DWORD  dwOutOf;
    };
```

##### Members

| Member | Description |
| `dwRequestID` | Double word containing the client defined request ID. |
| `dwArraySize` | Double word containing the number of elements in the list that are within this packet. For example, if there are items returned in the structure, then this field will contain 25, but if there are 400 items in the list and the data is returned in two packets, then this value will contain the number of entries within each packet, ie: 200. |
| `dwREntry` | Double word containing the index number of this list packet. This number will be from 0 to dwOutOf - 1. |
| `dwOutOf` | Double word containing the total number of packets used to transmit the list. |

##### Remarks

This structure is inherited by a number of other structs, for example: [`SIMCONNECT_RECV_FACILITY_MINIMAL_LIST`](SIMCONNECT_RECV_FACILITY_MINIMAL_LIST.htm).

This structure inherits the [`SIMCONNECT_RECV`](SIMCONNECT_RECV.htm) structure, so use the [`SIMCONNECT_RECV_ID`](SIMCONNECT_RECV_ID.htm) enumeration to determine which list structure has been received, which in turn will also determine the other members of the struct specific to the function that generated it.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_RECV\_ID](SIMCONNECT_RECV_ID.htm)
4. [SIMCONNECT\_RECV\_FACILITY\_MINIMAL\_LIST](SIMCONNECT_RECV_FACILITY_MINIMAL_LIST.htm)
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