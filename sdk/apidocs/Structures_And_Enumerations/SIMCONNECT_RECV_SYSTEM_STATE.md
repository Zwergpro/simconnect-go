SIMCONNECT\_RECV\_SYSTEM\_STATE

## SIMCONNECT\_RECV\_SYSTEM\_STATE

The **SIMCONNECT\_RECV\_SYSTEM\_STATE** structure is used with the [SimConnect\_RequestSystemState](../General/SimConnect_RequestSystemState.htm) function to retrieve specific Microsoft Flight Simulator 2024 systems states and information.

##### Syntax

```cpp
typedef struct SIMCONNECT_RECV_SYSTEM_STATE : public SIMCONNECT_RECV {
    DWORD  dwRequestID;
    DWORD  dwInteger;
    float  fFloat;
    char   szString[MAX_PATH];
    };
```

##### Members

| Member | Description |
| `dwRequestID` | Double word containing the client defined request ID. |
| `dwInteger` | Double word containing an integer, or boolean, value. |
| `fFloat` | A float value. |
| `szString` | Null-terminated string. |

##### Remarks

This structure inherits the `SIMCONNECT_RECV` structure and is returned when the `dwID` parameter of `SIMCONNECT_RECV` is set to `SIMCONNECT_RECV_ID_SYSTEM_STATE`.

Typically only one of the received integer, float or string will contain information, which one will depend on the request and can be identified by the request ID. Refer to the descriptions of the `SimConnect_SetSystemState` (Not Supported) and `SimConnect_RequestSystemState` functions.

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