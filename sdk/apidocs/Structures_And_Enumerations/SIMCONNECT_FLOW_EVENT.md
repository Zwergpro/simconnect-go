SIMCONNECT\_FLOW\_EVENT

## SIMCONNECT\_FLOW\_EVENT

The **SIMCONNECT\_FLOW\_EVENT** enumeration type is used with the [SimConnect\_SubscribeToFlowEvent](../Events_And_Data/SimConnect_SubscribeToFlowEvent.htm) call and the [SIMCONNECT\_RECV\_FLOW\_EVENT](SIMCONNECT_RECV_FLOW_EVENT.htm) message.

##### Syntax

```cpp
enum SIMCONNECT_FLOW_EVENT : DWORD
{
    SIMCONNECT_FLOW_EVENT_NONE,
    SIMCONNECT_FLOW_EVENT_FLT_LOAD,
    SIMCONNECT_FLOW_EVENT_FLT_LOADED,
    SIMCONNECT_FLOW_EVENT_TELEPORT_START,
    SIMCONNECT_FLOW_EVENT_TELEPORT_DONE,
    SIMCONNECT_FLOW_EVENT_BACK_ON_TRACK_START,
    SIMCONNECT_FLOW_EVENT_BACK_ON_TRACK_DONE,
    SIMCONNECT_FLOW_EVENT_SKIP_START,
    SIMCONNECT_FLOW_EVENT_SKIP_DONE,
    SIMCONNECT_FLOW_EVENT_BACK_TO_MAIN_MENU,
    SIMCONNECT_FLOW_EVENT_RTC_START,
    SIMCONNECT_FLOW_EVENT_RTC_END,
    SIMCONNECT_FLOW_EVENT_REPLAY_START,
    SIMCONNECT_FLOW_EVENT_REPLAY_END,
    SIMCONNECT_FLOW_EVENT_FLIGHT_START,
    SIMCONNECT_FLOW_EVENT_FLIGHT_END,
    SIMCONNECT_FLOW_EVENT_PLANE_CRASH,
};
```

##### Members

| Member | Description |
| `SIMCONNECT_FLOW_EVENT_NONE` | Should never be received. |
| `SIMCONNECT_FLOW_EVENT_FLT_LOAD` | Generic event that indicates that a flt file will be loaded. |
| `SIMCONNECT_FLOW_EVENT_FLT_LOADED` | Generic event that indicated that a flt file has been loaded which means that the aircraft is in a new state and that some internal variable should be re-initialized using the ones from the sim. |
| `SIMCONNECT_FLOW_EVENT_TELEPORT_START` | Generic event that indicates that a teleport action will be executed. During a teleportation, some systems, physics... might need to be paused to avoid weird behavior. |
| `SIMCONNECT_FLOW_EVENT_TELEPORT_DONE` | Generic event that indicates that a teleport action has been executed. |
| `SIMCONNECT_FLOW_EVENT_BACK_ON_TRACK_START` | Event that indicated that a back on track will be executed. Such as the teleportation, some systems might need to be paused. |
| `SIMCONNECT_FLOW_EVENT_BACK_ON_TRACK_DONE` | Event that indicates that a back on track has been executed. |
| `SIMCONNECT_FLOW_EVENT_SKIP_START` | Event that indicates that a part of an activity will be skipped (for instance in career mode, taxi part can be skipped to start on runway which will load the corresponding flt) |
| `SIMCONNECT_FLOW_EVENT_SKIP_DONE` | Event that indicates that a part of an activity has been skipped. |
| `SIMCONNECT_FLOW_EVENT_BACK_TO_MAIN_MENU` | Event that indicates that the sim is going back to the menu. |
| `SIMCONNECT_FLOW_EVENT_RTC_START` | Event that indicates that a RTC has started. Such as the teleportation, some systems might need to be paused. |
| `SIMCONNECT_FLOW_EVENT_RTC_END` | Event that indicates that a RTC has ended. |
| `SIMCONNECT_FLOW_EVENT_REPLAY_START` | Event that indicates that a replay has started. Such as the teleportation, some systems might need to be paused. |
| `SIMCONNECT_FLOW_EVENT_REPLAY_END` | Event that indicates that a replay has ended. |
| `SIMCONNECT_FLOW_EVENT_FLIGHT_START` | Event that indicates that a flight has started. |
| `SIMCONNECT_FLOW_EVENT_FLIGHT_END` | Event that indicates that a flight has ended. |
| `SIMCONNECT_FLOW_EVENT_PLANE_CRASH` | Event that indicates that an aircraft has crashed. |

##### Remarks

N/A

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_EnumerateInputEvents](../InputEvents/SimConnect_EnumerateInputEvents.htm)
4. [SimConnect\_GetInputEvent](../InputEvents/SimConnect_GetInputEvent.htm)
5. [SimConnect\_SubscribeInputEvent](../InputEvents/SimConnect_SubscribeInputEvent.htm)
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