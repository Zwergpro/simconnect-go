SimConnect\_EnumerateControllers

## SimConnect\_EnumerateControllers

The **SimConnect\_EnumerateControllers** function is used to retrieve a list of every device that is currently plugged into the simulation.

##### Syntax

```cpp
SIMCONNECTAPI SimConnect_EnumerateControllers(
    HANDLE hSimConnect
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

Might throw the following exception:

- `SIMCONNECT_EXCEPTION_ERROR` \- an internal error occurs or if no devices are plugged

##### Remarks

This function will generate a `SIMCONNECT_RECV_CONTROLLERS_LIST` response (with the [ID](../Structures_And_Enumerations/SIMCONNECT_RECV_ID.htm)`SIMCONNECT_RECV_ID_CONTROLLERS_LIST`), which will contain a list `SIMCONNECT_CONTROLLER_ITEM` structs.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_RECV\_GET\_INPUT\_EVENT](../Structures_And_Enumerations/SIMCONNECT_RECV_GET_INPUT_EVENT.htm)
4. [SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENTS](../Structures_And_Enumerations/SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS.htm)
5. [SIMCONNECT\_INPUT\_EVENT\_DESCRIPTOR](../Structures_And_Enumerations/SIMCONNECT_INPUT_EVENT_DESCRIPTOR.htm)
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