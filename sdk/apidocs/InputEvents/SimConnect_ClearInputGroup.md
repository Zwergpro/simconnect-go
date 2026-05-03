SimConnect\_ClearInputGroup

## SimConnect\_ClearInputGroup

The **SimConnect\_ClearInputGroup** function is used to remove all the input events from a specified input group object.

##### Syntax

```cpp
HRESULT SimConnect_ClearInputGroup(
    HANDLE  hSimConnect,
    SIMCONNECT_INPUT_GROUP_ID  GroupID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined input group that is to have all its events removed. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
static enum INPUT_ID {
    INPUT_1,
    };
static enum EVENT_ID {
    EVENT_1,
    };
hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_1, "parking_brakes");
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "a+B", EVENT_1);
      ....
hr = SimConnect_ClearInputGroup(hSimConnect, INPUT_1);
```

##### Remarks

Use this function to permanently delete an input group. Use the [SimConnect\_SetInputGroupState](SimConnect_SetInputGroupState.htm) function to temporarily suspend input group notifications.

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SimConnect\_MapInputEventToClientEvent](../Events_And_Data/SimConnect_MapInputEventToClientEvent.htm)
3. [SimConnect\_SetInputGroupPriority](../Events_And_Data/SimConnect_SetInputGroupPriority.htm)
4. [SimConnect\_RemoveInputEvent](../Events_And_Data/SimConnect_RemoveInputEvent.htm)
5. [SimConnect\_SetInputGroupState](SimConnect_SetInputGroupState.htm)
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