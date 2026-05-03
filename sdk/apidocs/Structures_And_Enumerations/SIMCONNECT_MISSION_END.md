SIMCONNECT\_MISSION\_END

## SIMCONNECT\_MISSION\_END

The **SIMCONNECT\_MISSION\_END** enumeration type is used to specify the three possible outcomes of a mission.

##### Syntax

```cpp
enum SIMCONNECT_MISSION_END {
    SIMCONNECT_MISSION_FAILED,
    SIMCONNECT_MISSION_CRASHED,
    SIMCONNECT_MISSION_SUCCEEDED
    };
```

##### Members

| Member | Description |
| `SIMCONNECT_MISSION_FAILED` | The mission failed for some reason other than a crash. |
| `SIMCONNECT_MISSION_CRASHED` | The mission failed because of a crash. |
| `SIMCONNECT_MISSION_SUCCEEDED` | The mission was completed successfully. |

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