SIMCONNECT\_RECV\_COMM\_BUS

## SIMCONNECT\_RECV\_COMM\_BUS

The **SIMCONNECT\_RECV\_COMM\_BUS** structure is used to hold data from an event subscribed to using [SimConnect\_SubscribeToCommBusEvent](../Communication/SimConnect_SubscribeToCommBusEvent.htm).

##### Syntax

```cpp
struct SIMCONNECT_RECV_COMM_BUS : public SIMCONNECT_RECV_LIST_TEMPLATE
{
    DWORD   uEventID;
    SIMCONNECT_STRINGV(rgData);
};
```

##### Members

| Member | Description |
| `uEventID` | This is the ID of the event when the event has been subscribed to. |
| `rgData` | This is the date received when the event has been called. |

##### Remarks

This struct may be sent several times for a single event in those cases where the event data is larger than the maximum packate size.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [Communication](../Communication/Communication.htm)

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