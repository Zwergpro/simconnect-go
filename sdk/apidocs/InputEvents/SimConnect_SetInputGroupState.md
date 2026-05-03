SimConnect\_SetInputGroupState

## SimConnect\_SetInputGroupState

The **SimConnect\_SetInputGroupState** function is used to turn requests for input event information from the server on and off.

##### Syntax

```cpp
HRESULT SimConnect_SetInputGroupState(
      HANDLE  hSimConnect,
      SIMCONNECT_INPUT_GROUP_ID  GroupID,
      DWORD  dwState
      );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined input group that is to have its state changed. | Integer |
| _dwState_ | Double word containing the new state. One member of the [SIMCONNECT\_STATE](../Structures_And_Enumerations/SIMCONNECT_STATE.htm) enumeration type. | Integer |

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
    INPUT_2
    };
static enum EVENT_ID {
    EVENT_1,
    EVENT_2
    };
      ....
hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_1, "parking_brakes");
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "ctrl+U+Q", EVENT_1);
hr = SimConnect_SetInputGroupState(hSimConnect, INPUT_1, SIMCONNECT_STATE_ON);
```

##### Remarks

The default state for input groups is to be inactive, so make sure to call this function each time an input group is to become active.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_MapInputEventToClientEvent](../Events_And_Data/SimConnect_MapInputEventToClientEvent.htm)
4. [SimConnect\_SetInputGroupPriority](../Events_And_Data/SimConnect_SetInputGroupPriority.htm)
5. [SimConnect\_RemoveInputEvent](../Events_And_Data/SimConnect_RemoveInputEvent.htm)
6. [SimConnect\_ClearInputGroup](../Events_And_Data/SimConnect_ClearInputGroup.htm)
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