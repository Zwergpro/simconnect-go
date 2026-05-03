SIMCONNECT\_DATA\_LATLONALT

## SIMCONNECT\_DATA\_LATLONALT

The **SIMCONNECT\_DATA\_LATLONALT** structure is used to hold a world position.

##### Syntax

```cpp
struct SIMCONNECT_DATA_LATLONALT{
    double  Latitude;
    double  Longitude;
    double  Altitude;
    };
```

##### Members

| Member | Description |
| `Latitude` | Latitude in degrees. |
| `Longitude` | Longitude in degrees. |
| `Altitude` | Altitude in feet. |

##### Remarks

This structure is used when one of the following simulation variables is requested (with a call to [SimConnect\_RequestDataOnSimObject](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm)):

- `ADF LATLONALT`
- `NAV DME LATLONALT`
- `NAV GS LATLONALT`
- `NAV VOR LATLONALT`
- `INNER MARKER LATLONALT`
- `MIDDLE MARKER LATLONALT`
- `OUTER MARKER LATLONALT`

It is also returned as part of the `SIMCONNECT_RECV_FACILITIES_MINIMAL` structure.

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