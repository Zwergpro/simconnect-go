SIMCONNECT\_RECV\_ASSIGNED\_OBJECT\_ID

## SIMCONNECT\_RECV\_ASSIGNED\_OBJECT\_ID

The **SIMCONNECT\_RECV\_ASSIGNED\_OBJECT\_ID** structure is used to return an object ID that matches a request ID.

##### Syntax

```cpp
struct SIMCONNECT_RECV_ASSIGNED_OBJECT_ID : public SIMCONNECT_RECV {
    DWORD  dwRequestID;
    DWORD  dwObjectID;
    };
```

##### Members

| Member | Description |
| `dwRequestID` | Double word containing the client defined request ID. |
| `dwObjectID` | Double word containing the server defined object ID. |

##### Remarks

This structure inherits the `SIMCONNECT_RECV` structure and is returned when the dwID parameter of `SIMCONNECT_RECV` is set to [`SIMCONNECT_RECV_ID_ASSIGNED_OBJECT_ID`](SIMCONNECT_RECV_ID.htm).

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