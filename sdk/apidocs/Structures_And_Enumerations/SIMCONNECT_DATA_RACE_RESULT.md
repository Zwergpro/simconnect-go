SIMCONNECT\_DATA\_RACE\_RESULT

## SIMCONNECT\_DATA\_RACE\_RESULT

The **SIMCONNECT\_DATA\_RACE\_RESULT** structure is used to hold multiplayer racing results.

##### Syntax

```cpp
struct SIMCONNECT_DATA_RACE_RESULT{
    DWORD  dwNumberOfRacers;
    GUID  MissionGUID;
    char  szPlayerName[MAX_PATH];
    char  szSessionType[MAX_PATH];
    char  szAircraft[MAX_PATH];
    char  szPlayerRole[MAX_PATH];
    double  fTotalTime;
    double  fPenaltyTime;
    DWORD  dwIsDisqualified;
};
```

##### Members

| Member | Description |
| `dwNumberOfRacers` | The total number of racers. |
| `MissionGUID` | The GUID of the mission that has been selected by the host. |
| `szPlayerName[MAX_PATH]` | Null terminated string containing the name of the player. |
| `szSessionType[MAX_PATH]` | Null terminated string containing the type of the multiplayer session, currently one of: "LAN" or "GAMESPY". |
| `szAircraft[MAX_PATH]` | Null terminated string containing the aircraft type. |
| `szPlayerRole[MAX_PATH]` | Null terminated string containing the player's role (or name) in the mission. This string wil be filled from the Name property of the Player object in the mission file. |
| `fTotalTime` | If this structure is a member of a [SIMCONNECT\_RECV\_EVENT\_RACE\_END](SIMCONNECT_RECV_EVENT_RACE_END.htm) structure, then this contains the final race time in seconds, or 0 for DNF (Did Not Finish). If this structure is a member of a [SIMCONNECT\_RECV\_EVENT\_RACE\_LAP](SIMCONNECT_RECV_EVENT_RACE_LAP.htm) structure, then this contains the lap time in seconds. |
| `fPenaltyTime` | If this structure is a member of a [SIMCONNECT\_RECV\_EVENT\_RACE\_END](SIMCONNECT_RECV_EVENT_RACE_END.htm) structure, then this contains the final penalty time in seconds. If this structure is a member of a [SIMCONNECT\_RECV\_EVENT\_RACE\_LAP](SIMCONNECT_RECV_EVENT_RACE_LAP.htm) structure, then this contains the total penalty time in seconds received so far (not just for this lap). |
| `dwIsDisqualified` | A boolean value, 0 indicating the player has not been disqualified, non-zero indicating they have been disqualified. |

##### Remarks

This structure is never sent on its own, but is always a member of either a [SIMCONNECT\_RECV\_EVENT\_RACE\_END](SIMCONNECT_RECV_EVENT_RACE_END.htm) structure or a [SIMCONNECT\_RECV\_EVENT\_RACE\_LAP](SIMCONNECT_RECV_EVENT_RACE_LAP.htm) structure.

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