SimConnect\_ClearNotificationGroup

## SimConnect\_ClearNotificationGroup

The **SimConnect\_ClearNotificationGroup** function is used to remove all the client defined events from a notification group.

##### Syntax

```cpp
HRESULT SimConnect_ClearNotificationGroup(
    HANDLE  hSimConnect,
    SIMCONNECT_NOTIFICATION_GROUP_ID  GroupID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined group that is to have all its events removed. | Integer |

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
hr = SimConnect_ClearNotificationGroup(hSimConnect, GROUP_1);
```

##### Remarks

There is a maximum of 20 notification groups in any SimConnect client. Use this function if the maximum has been reached, but one or more are not longer required.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AddClientEventToNotificationGroup](SimConnect_AddClientEventToNotificationGroup.htm)
4. [SimConnect\_RemoveClientEvent](SimConnect_RemoveClientEvent.htm)
5. [SimConnect\_SetNotificationGroupPriority](../General/SimConnect_SetNotificationGroupPriority.htm)
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