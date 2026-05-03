SIMCONNECT\_RECV\_EVENT\_MULTIPLAYER\_SESSION\_ENDED

## SIMCONNECT\_RECV\_EVENT\_MULTIPLAYER\_SESSION\_ENDED

The **SIMCONNECT\_RECV\_EVENT\_MULTIPLAYER\_SESSION\_ENDED** structure is sent to a client when they have requested to leave a race, or to all players when the session is terminated by the host.

##### Syntax

```cpp
struct SIMCONNECT_RECV_EVENT_MULTIPLAYER_SESSION_ENDED : public SIMCONNECT_RECV_EVENT {};
```

##### Members

This structure takes no parameters in addition to those inherited from the [SIMCONNECT\_RECV\_EVENT](SIMCONNECT_RECV_EVENT.htm) structure.

##### Remarks

This is the only event that is broadcast to all the players in a multi-player race, in the situation where the host terminates, or simply leaves, the race. If a client ends their own participation in the race, they will be the only one to receive the event.

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