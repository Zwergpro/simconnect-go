SimConnect\_TransmitClientEvent\_EX1

## SimConnect\_TransmitClientEvent\_EX1

The **SimConnect\_TransmitClientEvent\_EX1** function is used to request that the Microsoft Flight Simulator server transmit to all SimConnect clients the specified client event. This function is specifically designed to permit the sending of multiple parameters for the key event (up to five), unlike the `SimConnect_TransmitClientEvent` which only permits one.

##### Syntax

```cpp
HRESULT SimConnect_TransmitClientEvent_EX1(
    HANDLE  hSimConnect,
    SIMCONNECT_OBJECT_ID  ObjectID,
    SIMCONNECT_CLIENT_EVENT_ID  EventID,
    SIMCONNECT_NOTIFICATION_GROUP_ID GroupID,
    SIMCONNECT_EVENT_FLAG Flags
    DWORD dwData0,
    DWORD dwData1,
    DWORD dwData2,
    DWORD dwData3,
    DWORD dwData4
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _ObjectID_ | Specifies the ID of the server defined object. If this parameter is set to `SIMCONNECT_OBJECT_ID_USER`, then the transmitted event will be sent to the other clients in priority order. If this parameters contains another object ID, then the event will be sent direct to that sim-object, and no other clients will receive it. | Integer |
| _EventID_ | Specifies the ID of the client event. | Integer |
| _GroupID_ | The default behavior is that this specifies the GroupID of the event. The SimConnect server will use the priority of this group to send the message to all clients with a lower priority. To receive the event notification other SimConnect clients must have subscribed to receive the event. See the explanation of [SimConnect Priorities](../../SimConnect_API_Reference.htm#simconnect-priorities). The exception to the default behavior is set by the `SIMCONNECT_EVENT_FLAG_GROUPID_IS_PRIORITY` flag, described below. | Integer |
| _Flags_ | One or more of the flags shown in the table below. | Integer |
| _dwData0_ to _dwData4_ | Double word containing any additional number(s) required by the event. If the event is a Microsoft Flight Simulator 2024 event, then refer to the [Event IDs document](../../../Key_Events/Key_Events.htm) for information on this additional value(s). If the event is a custom event, then any values put in these parameters will be available to the clients that receive the event. | Integer |

The following table shows the different _flags_ that can be used:

| Flag | Description |
| --- | --- |
| 0 | Do nothing. |
| `SIMCONNECT_EVENT_FLAG_SLOW_REPEAT_TIMER` | The flag will effectively reset the repeat timer to simulate slow repeat. Use this flag to reset the repeat timer that works with various keyboard events and mouse clicks. |
| `SIMCONNECT_EVENT_FLAG_FAST_REPEAT_TIMER` | The flag will effectively reset the repeat timer to simulate fast repeat. |
| `SIMCONNECT_EVENT_FLAG_GROUPID_IS_PRIORITY` | Indicates to the SimConnect server to treat the GroupID as a priority value. If this flag is set, and the GroupID parameter is set to `SIMCONNECT_GROUP_PRIORITY_HIGHEST` then all client notification groups that have subscribed to the event will receive the notification (unless one of them masks it). The event will be transmitted to clients starting at the priority specified in the GroupID parameter, though this can result in the client that transmitted the event, receiving it again. |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
static enum EVENT_ID {
    EVENT_MY_EVENT
    EVENT_DME
    };
hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_MY_EVENT, "Custom.Event");
SimConnect_TransmitClientEvent_EX1(hSimConnect, 0, EVENT_MY_EVENT, SIMCONNECT_GROUP_PRIORITY_HIGHEST, SIMCONNECT_EVENT_FLAG_GROUPID_IS_PRIORITY, 0);

hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_DME, "DME_SELECT");
SimConnect_TransmitClientEvent_EX1(hSimConnect, 0, EVENT_DME, SIMCONNECT_GROUP_PRIORITY_DEFAULT, SIMCONNECT_EVENT_FLAG_GROUPID_IS_PRIORITY, 2);
```

##### Remarks

Typically use this function to transmit an event to other SimConnect clients, including the simulation engine, although the client that transmits the event can also receive it. The order in which client notification groups are informed of the event is determined by the priority of each group. The higher the priority of the group, the earlier it will receive the event notification. Refer to the explanation of the **_maskable_** parameter for the [SimConnect\_AddClientEventToNotificationGroup](SimConnect_AddClientEventToNotificationGroup.htm) call, which describes when the event may be masked and not transmitted to lower priority groups. Also see the explanation of [SimConnect Priorities](../../SimConnect_API_Reference.htm#simconnect-priorities).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_MapClientEventToSimEvent](SimConnect_MapClientEventToSimEvent.htm)
4. [SIMCONNECT\_RECV\_ID](../Structures_And_Enumerations/SIMCONNECT_RECV_ID.htm)
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