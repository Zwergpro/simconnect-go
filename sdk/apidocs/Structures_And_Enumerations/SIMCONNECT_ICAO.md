SIMCONNECT\_ICAO

## SIMCONNECT\_ICAO

The **SIMCONNECT\_ICAO** enumeration type is returned as part of the contents of the [`SIMCONNECT_RECV_FACILITIES_MINIMAL`](SIMCONNECT_FACILITY_MINIMAL.htm) struct to get data about a specific ICAO.

##### Syntax

```cpp
struct SIMCONNECT_ICAO {
    char Type;
    char Ident(8 + 1);
    char Region(2 + 1);
    char Airport(4 + 1);
    };
```

##### Members

| Member | Description |
| `Ident` | The identity string. |
| `Region` | The region string. |
| `Airport` | The airport string. |

##### Remarks

N/A

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_FACILITY\_MINIMAL](SIMCONNECT_FACILITY_MINIMAL.htm)
4. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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