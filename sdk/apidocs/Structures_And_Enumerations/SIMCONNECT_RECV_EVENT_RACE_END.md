SIMCONNECT\_RECV\_EVENT\_RACE\_END

## SIMCONNECT\_RECV\_EVENT\_RACE\_END

The **SIMCONNECT\_RECV\_EVENT\_RACE\_END** structure is used in multi-player racing to hold the results for one player at the end of a race.

##### Syntax

```cpp
struct SIMCONNECT_RECV_EVENT_RACE_END : public SIMCONNECT_RECV_EVENT {
    DWORD  dwRacerNumber;
    SIMCONNECT_DATA_RACE_RESULT  RacerData;
};
```

##### Members

| Member | Description |
| `dwRacerNumber` | The index of the racer the results are for. Players are indexed from 0. |
| `RacerData` | A SIMCONNECT\_DATA\_RACE\_RESULT structure. |

##### Remarks

In a multi-player race players can come and go, so index numbers are not a reliable means of identifiying the players. The `szPlayerName` and `szPlayerRole` parameters of the [SIMCONNECT\_DATA\_RACE\_RESULT](SIMCONNECT_DATA_RACE_RESULT.htm) structure should be used to identify each player.

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