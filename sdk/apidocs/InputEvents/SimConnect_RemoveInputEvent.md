SimConnect\_RemoveInputEvent

## SimConnect\_RemoveInputEvent

The **SimConnect\_RemoveInputEvent** function is used to remove an input event from a specified input group object.

##### Syntax

```cpp
HRESULT SimConnect_RemoveInputEvent(
      HANDLE  hSimConnect,
      SIMCONNECT_INPUT_GROUP_ID  GroupID,
      const char*  pszInputDefinition
      );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined input group from which the event is to be removed. | Integer |
| _pszInputDefinition_ | Pointer to a null-terminated string containing the input definition. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |
| E\_INVALIDARG | A SimConnect section in the [SimConnect.cfg](../../SimConnect_CFG_Definition.htm) file did not contain the config index requested in the parameters. |

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
hr = SimConnect_RemoveInputEvent(hSimConnect, INPUT_1, "a+B");
```

##### Remarks

The input string definitions must match exactly, before anything is removed from the group definition. For example, the string definitions "A+B" and "a+B" do not match.

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SimConnect\_MapInputEventToClientEvent](../Events_And_Data/SimConnect_MapInputEventToClientEvent.htm)
3. [SimConnect\_SetInputGroupPriority](../Events_And_Data/SimConnect_SetInputGroupPriority.htm)
4. [SimConnect\_ClearInputGroup](../Events_And_Data/SimConnect_ClearInputGroup.htm)
5. [SimConnect\_SetInputGroupState](../Events_And_Data/SimConnect_SetInputGroupState.htm)
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