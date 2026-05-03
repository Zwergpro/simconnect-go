SIMCONNECT\_RECV\_EVENT\_MULTIPLAYER\_CLIENT\_STARTED

## SIMCONNECT\_RECV\_EVENT\_MULTIPLAYER\_CLIENT\_STARTED

The **SIMCONNECT\_RECV\_EVENT\_MULTIPLAYER\_CLIENT\_STARTED** structure is sent to a client when they have successfully joined a multi-player race.

##### Syntax

```cpp
struct SIMCONNECT_RECV_EVENT_MULTIPLAYER_CLIENT_STARTED : public SIMCONNECT_RECV_EVENT {};
```

##### Members

This structure takes no parameters in addition to those inherited from the [SIMCONNECT\_RECV\_EVENT](SIMCONNECT_RECV_EVENT.htm) structure.

##### Remarks

This event is not transmitted to the host of the session, only to the client that has joined in. To receive these events, refer to the [SimConnect\_SubscribeToSystemEvent](../General/SimConnect_SubscribeToSystemEvent.htm) function.

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