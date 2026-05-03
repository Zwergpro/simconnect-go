SIMCONNECT\_FACILITY\_MINIMAL

## SIMCONNECT\_FACILITY\_MINIMAL

The **SIMCONNECT\_FACILITY\_MINIMAL** structure is used to provide information on the number of elements in a list of facilities returned to the client, and the number of packets that were used to transmit the data.

##### Syntax

```cpp
struct SIMCONNECT_FACILITY_MINIMAL{
    SIMCONNECT_ICAO icao;
    SIMCONNECT_DATA_LATLONALT lla;
    };
```

##### Members

| Member | Description |
| `icao` | The [`SIMCONNECT_ICAO`](SIMCONNECT_ICAO.htm) struct with information about the facility. |
| `lla` | The latitude, longitude and altitude of the facility, returned as a [`SIMCONNECT_DATA_LATLONALT`](SIMCONNECT_DATA_LATLONALT.htm) struct. |

##### Remarks

This structure is returned as part of the array of minimal entries within the [`SIMCONNECT_RECV_FACILITY_MINIMAL_LIST`](SIMCONNECT_RECV_FACILITY_MINIMAL_LIST.htm) structure.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_RECV\_ID](SIMCONNECT_RECV_ID.htm)
4. [SIMCONNECT\_ICAO](SIMCONNECT_ICAO.htm)
5. [SIMCONNECT\_DATA\_LATLONALT](SIMCONNECT_DATA_LATLONALT.htm)
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