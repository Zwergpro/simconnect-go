SIMCONNECT\_COMM\_BUS\_BROADCAST\_TO

## SIMCONNECT\_COMM\_BUS\_BROADCAST\_TO

The **SIMCONNECT\_COMM\_BUS\_BROADCAST\_TO** enum is used to define where a communication (CommBus) event should be broadcast.

##### Syntax

```cpp
enum SIMCONNECT_COMM_BUS_BROADCAST_TO
{
    SIMCONNECT_COMM_BUS_BROADCAST_TO_JS = 1 << 0;
    SIMCONNECT_COMM_BUS_BROADCAST_TO_WASM = 1 << 1;
    SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT = 1 << 3;
    SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT_SELF_CALL = 1 << 4;
    SIMCONNECT_COMM_BUS_BROADCAST_TO_DEFAULT = SIMCONNECT_COMM_BUS_BROADCAST_TO_JS | SIMCONNECT_COMM_BUS_BROADCAST_TO_WASM | SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT;
    SIMCONNECT_COMM_BUS_BROADCAST_TO_ALL_SIMCONNECT = SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT | SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT_SELF_CALL;
    SIMCONNECT_COMM_BUS_BROADCAST_TO_ALL = SIMCONNECT_COMM_BUS_BROADCAST_TO_JS | SIMCONNECT_COMM_BUS_BROADCAST_TO_WASM | SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT | SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT_SELF_CALL;
};
```

##### Members

| Member | Description |
| `SIMCONNECT_COMM_BUS_BROADCAST_TO_JS` | Broadcast the event to all subscribed JS gauges. |
| `SIMCONNECT_COMM_BUS_BROADCAST_TO_WASM` | Broadcast the event to all subscribed WASM gauges. |
| `SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT` | Broadcast the event to all subscribed SimConnect clients, _except_ itself. |
| `SIMCONNECT_COMM_BUS_BROADCAST_TO_SIMCONNECT_SELF_CALL` | Broadcast the event to all SimConnect clients, _including_ itself. |
| `SIMCONNECT_COMM_BUS_BROADCAST_TO_DEFAULT` | Broadcast the event to all JS _and_ WASM gauges, as well as all SimConnect clients, _except_ itself. |
| `SIMCONNECT_COMM_BUS_BROADCAST_TO_ALL` | Broadcast the event to all JS _and_ WASM gauges, as well as all SimConnect clients, _including_ itself. |

##### Remarks

N/A

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)
4. [Communication](../Communication/Communication.htm)

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