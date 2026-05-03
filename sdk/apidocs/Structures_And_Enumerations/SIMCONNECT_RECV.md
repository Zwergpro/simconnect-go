SIMCONNECT\_RECV

## SIMCONNECT\_RECV

The **SIMCONNECT\_RECV** structure is used with the [SIMCONNECT\_RECV\_ID](SIMCONNECT_RECV_ID.htm) enumeration to indicate which type of structure has been returned.

##### Syntax

```cpp
struct SIMCONNECT_RECV{
    DWORD  dwSize;
    DWORD  dwVersion;
    DWORD  dwID;
    };
```

##### Members

| Member | Description |
| `dwSize` | The total size of the returned structure in bytes (that is, not usually the size of the SIMCONNECT\_RECV structure, but of the structure that inherits it). |
| `dwVersion` | The version number of the SimConnect server. |
| `dwID` | The ID of the returned structure. One member of SIMCONNECT\_RECV\_ID. |

##### Remarks

This structure is inherited directly by:

- [SIMCONNECT\_RECV\_OPEN](SIMCONNECT_RECV_OPEN.htm)
- [SIMCONNECT\_RECV\_EVENT](SIMCONNECT_RECV_EVENT.htm)
- [SIMCONNECT\_RECV\_EVENT\_EX1](SIMCONNECT_RECV_EVENT_EX1.htm)
- [SIMCONNECT\_RECV\_EXCEPTION](SIMCONNECT_RECV_EXCEPTION.htm)
- [SIMCONNECT\_RECV\_SIMOBJECT\_DATA](SIMCONNECT_RECV_SIMOBJECT_DATA.htm)
- [SIMCONNECT\_RECV\_ASSIGNED\_OBJECT\_ID](SIMCONNECT_RECV_ASSIGNED_OBJECT_ID.htm)
- [SIMCONNECT\_RECV\_RESERVED\_KEY](SIMCONNECT_RECV_RESERVED_KEY.htm)
- [SIMCONNECT\_RECV\_SYSTEM\_STATE](SIMCONNECT_RECV_SYSTEM_STATE.htm)
- [SIMCONNECT\_RECV\_FACILITIES\_LIST](SIMCONNECT_RECV_FACILITIES_LIST.htm)
- [SIMCONNECT\_RECV\_FACILITY\_DATA](SIMCONNECT_RECV_FACILITY_DATA.htm)
- [SIMCONNECT\_RECV\_FACILITY\_DATA\_END](SIMCONNECT_RECV_FACILITY_DATA_END.htm)
- [SIMCONNECT\_RECV\_SYSTEM\_STATE](SIMCONNECT_RECV_SYSTEM_STATE.htm)
- [SIMCONNECT\_RECV\_RESERVED\_KEY](SIMCONNECT_RECV_RESERVED_KEY.htm)
- [SIMCONNECT\_RECV\_LIST\_TEMPLATE](SIMCONNECT_RECV_LIST_TEMPLATE.htm)

The structure is also inherited by [`SIMCONNECT_RECV_QUIT`](SIMCONNECT_RECV_QUIT.htm), which does not add any new members. This structure is received when the user quits Microsoft Flight Simulator.

This structure is inherited by the [`SIMCONNECT_RECV_EVENT`](SIMCONNECT_RECV_EVENT.htm) structure, which is itself inherited by several other structures:

- [SIMCONNECT\_RECV\_EVENT\_FILENAME](SIMCONNECT_RECV_EVENT_FILENAME.htm)
- [SIMCONNECT\_RECV\_EVENT\_FRAME](SIMCONNECT_RECV_EVENT_FRAME.htm)
- [SIMCONENCT\_RECV\_EVENT\_OBJECT\_ADDREMOVE](SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE.htm)

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