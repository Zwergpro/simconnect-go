SIMCONNECT\_RECV\_FLOW\_EVENT

## SIMCONNECT\_RECV\_FLOW\_EVENT

The **SIMCONNECT\_RECV\_FLOW\_EVENT** structure is received if a flow event is triggered by the sim and if the client has subscribed to those events using [SimConnect\_SubscribeToFlowEvent](../Events_And_Data/SimConnect_SubscribeToFlowEvent.htm).

##### Syntax

```cpp
SIMCONNECT_REFSTRUCT SIMCONNECT_RECV_FLOW_EVENT : public SIMCONNECT_RECV
{
    SIMCONNECT_FLOW_EVENT FlowEvent;
    SIMCONNECT_STRING(FltPath, 256);
};
```

##### Members

| Member | Description |
| `FlowEvent` | Type of the event which has been triggered by the sim. Event is part of the enum [SIMCONNECT\_FLOW\_EVENT](SIMCONNECT_FLOW_EVENT.htm). |
| `FltPath` | VFS path of a potential FLT file that is loaded. |

##### Remarks

This struct inherits members from the `SIMCONNECT_RECV` struct.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_RequestJetwayData](../Facilities/SimConnect_RequestJetwayData.htm)
4. [SIMCONNECT\_JETWAY\_DATA](SIMCONNECT_JETWAY_DATA.htm)
5. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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