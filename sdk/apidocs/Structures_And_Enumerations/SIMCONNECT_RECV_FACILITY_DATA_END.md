SIMCONNECT\_RECV\_FACILITY\_DATA\_END

## SIMCONNECT\_RECV\_FACILITY\_DATA\_END

The **SIMCONNECT\_RECV\_FACILITY\_DATA\_END** structure is used to signify the end of a data stream from the server after a call to `SimConnect_RequestFacilityData`.

##### Syntax

```cpp
struct SIMCONNECT_RECV_FACILITY_DATA_END : public SIMCONNECT_RECV{
    DWORD RequestId
    };
```

##### Members

| Member | Description |
| `RequestId` | Double word containing the client defined request ID. |

##### Remarks

This structure inherits the [`SIMCONNECT_RECV`](SIMCONNECT_RECV.htm) structure, so use the [`SIMCONNECT_RECV_ID`](SIMCONNECT_RECV_ID.htm) enumeration to determine which list structure has been received.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_RECV\_ID](SIMCONNECT_RECV_ID.htm)
4. [SimConnect\_RequestFacilityData](../Facilities/SimConnect_RequestFacilityData.htm)
5. [SIMCONNECT\_RECV\_FACILITY\_DATA](SIMCONNECT_RECV_FACILITY_DATA.htm)
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