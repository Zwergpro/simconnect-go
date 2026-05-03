SimConnect\_SetNotificationGroupPriority

## SimConnect\_SetNotificationGroupPriority

The `SimConnect_SetNotificationGroupPriority` function is used to set the priority for a notification group.

##### Syntax

```cpp
HRESULT SimConnect_SetNotificationGroupPriority(
  HANDLE  hSimConnect,
  SIMCONNECT_NOTIFICATION_GROUP_ID  GroupID,
  DWORD  uPriority
);
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined group. | Integer |
| _uPriority_ | Requests the group's priority. See the explanation of [SimConnect Priorities](../../SimConnect_API_Reference.htm#simconnect-priorities). | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| `S_OK` | The function succeeded. |
| `E_FAIL` | The function failed. |

##### Example

```cpp
hr=SimConnect_SetNotificationGroupPriority(hSimConnect, GROUP0, SIMCONNECT_GROUP_PRIORITY_HIGHEST);
```

##### Remarks

See [SimConnect Priorities](../../SimConnect_API_Reference.htm#simconnect-priorities) for additional information.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AddClientEventToNotificationGroup](../Events_And_Data/SimConnect_AddClientEventToNotificationGroup.htm)
4. [SimConnect\_RemoveClientEvent](../Events_And_Data/SimConnect_RemoveClientEvent.htm)
5. [SimConnect\_ClearNotificationGroup](../Events_And_Data/SimConnect_ClearNotificationGroup.htm)
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