SIMCONNECT\_RECV\_EVENT\_RACE\_LAP

## SIMCONNECT\_RECV\_EVENT\_RACE\_LAP

The **SIMCONNECT\_RECV\_EVENT\_RACE\_LAP** structure is used in multi-player racing to hold the results for one player at the end of a lap.

##### Syntax

```cpp
struct SIMCONNECT_RECV_EVENT_RACE_LAP : public SIMCONNECT_RECV_EVENT {
    DWORD  dwLapIndex;
    SIMCONNECT_DATA_RACE_RESULT  RacerData;
};
```

##### Members

| Member | Description |
| `dwLapIndex` | The index of the lap the results are for. Laps are indexed from 0. |
| `RacerData` | A [SIMCONNECT\_DATA\_RACE\_RESULT](SIMCONNECT_DATA_RACE_RESULT.htm) structure. |

##### Remarks

N/A

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