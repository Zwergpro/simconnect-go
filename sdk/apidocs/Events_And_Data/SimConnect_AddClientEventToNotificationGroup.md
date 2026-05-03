SimConnect\_AddClientEventToNotificationGroup

## SimConnect\_AddClientEventToNotificationGroup

The **SimConnect\_AddClientEventToNotificationGroup** function is used to add an individual client defined event to a notification group.

##### Syntax

```cpp
HRESULT SimConnect_AddClientEventToNotificationGroup(
      HANDLE  hSimConnect,
      SIMCONNECT_NOTIFICATION_GROUP_ID  GroupID,
      SIMCONNECT_CLIENT_EVENT_ID  EventID,
      BOOL  bMaskable = FALSE
      );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined group. | Integer |
| _EventID_ | Specifies the ID of the client defined event. | Integer |
| _bMaskable_ | **True** indicates that the event will be masked by this client and will not be transmitted to any more clients, possibly including _Microsoft Flight Simulator_ itself (if the priority of the client exceeds that of _Flight Simulator_). **False** is the default. See the explanation of [SimConnect Priorities](../../SimConnect_API_Reference.htm#simconnect-priorities). | Boolean<br>(OPTIONAL) |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
static enum EVENT_ID {
    EVENT_1,
    EVENT_2
    EVENT_3
    };
static enum GROUP_ID {
    GROUP_1,
    };
hr = SimConnect_AddClientEventToNotificationGroup(hSimConnect, GROUP_1, EVENT_1);
hr = SimConnect_AddClientEventToNotificationGroup(hSimConnect, GROUP_1, EVENT_2);
hr = SimConnect_AddClientEventToNotificationGroup(hSimConnect, GROUP_1, EVENT_3, TRUE);
hr = SimConnect_SetNotificationGroupPriority(hSimConnect, GROUP_1, SIMCONNECT_GROUP_PRIORITY_HIGHEST);
```

##### Remarks

The maximum number of events that can be added to a notification group is 1000. A notification group is simply a convenient way of setting the appropriate priority for a range of events, and all client events (such as EVENT\_1, EVENT\_2, EVENT\_3 in the example above) must be assigned to a notification group before any event notifications will be received from the SimConnect server.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_RemoveClientEvent](SimConnect_RemoveClientEvent.htm)
4. [SimConnect\_SetNotificationGroupPriority](../General/SimConnect_SetNotificationGroupPriority.htm)
5. [SimConnect\_ClearNotificationGroup](SimConnect_ClearNotificationGroup.htm)
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