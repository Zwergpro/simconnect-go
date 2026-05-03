SimConnect\_RequestNotificationGroup

## SimConnect\_RequestNotificationGroup

The **SimConnect\_RequestNotificationGroup** function is used to request events are transmitted from a notification group, when the simulation is in Dialog Mode.

##### Syntax

```cpp
HRESULT SimConnect_RequestNotificationGroup(
    HANDLE  hSimConnect,
    SIMCONNECT_NOTIFICATION_GROUP_ID  GroupID,
    DWORD  dwReserved = 0,
    DWORD  Flags = 0
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined input group that is to have all its events removed. | Integer |
| _dwReserved_ | Reserved for future use. | N/A |
| _Flags_ | Reserved for future use. | N/A |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

In this version this function has the specific purpose of enabling the sending of events, particularly joystick events, when the simulation is in Dialog Mode.

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