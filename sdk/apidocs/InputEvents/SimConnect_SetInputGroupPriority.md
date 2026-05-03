SimConnect\_SetInputGroupPriority

## SimConnect\_SetInputGroupPriority

The **SimConnect\_SetInputGroupPriority** function is used to set the priority for a specified input group object.

##### Syntax

```cpp
RESULT SimConnect_SetInputGroupPriority(
    HANDLE  hSimConnect,
    SIMCONNECT_INPUT_GROUP_ID  GroupID,
    DWORD  uPriority
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined input group that the priority setting is to apply to. | Integer |
| _uPriority_ | Specifies the priority setting for the input group. See the explanation of [SimConnect Priorities](../../SimConnect_API_Reference.htm#simconnect-priorities). | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
hr=SimConnect_SetInputGroupPriority(hSimConnect,INPUT0,SIMCONNECT_GROUP_PRIORITY_HIGHEST);
```

##### Remarks

A priority setting must be made for all input groups, otherwise event notifications will not be sent by the SimConnect server.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_MapInputEventToClientEvent](../Events_And_Data/SimConnect_MapInputEventToClientEvent.htm)
4. [SimConnect\_RemoveInputEvent](../Events_And_Data/SimConnect_RemoveInputEvent.htm)
5. [SimConnect\_ClearInputGroup](../Events_And_Data/SimConnect_ClearInputGroup.htm)
6. [SimConnect\_SetInputGroupState](../Events_And_Data/SimConnect_SetInputGroupState.htm)
7. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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