SimConnect\_RemoveClientEvent

## SimConnect\_RemoveClientEvent

The **SimConnect\_RemoveClientEvent** function is used to remove a client defined event from a notification group.

##### Syntax

```cpp
HRESULT SimConnect_RemoveClientEvent(
      HANDLE  hSimConnect,
      SIMCONNECT_NOTIFICATION_GROUP_ID  GroupID,
      SIMCONNECT_CLIENT_EVENT_ID  EventID
      );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined group. | Integer |
| _EventID_ | Specifies the ID of the client defined event ID that is to be removed from the group. | Integer |

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
hr = SimConnect_RemoveClientEvent(hSimConnect, GROUP_1, EVENT_2);
```

##### Remarks

Use this function to permanently remove the client event. There is no reliable procedure to temporarily turn off a client event.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AddClientEventToNotificationGroup](SimConnect_AddClientEventToNotificationGroup.htm)
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