SimConnect\_AIReleaseControl

## SimConnect\_AIReleaseControl

The **SimConnect\_AIReleaseControl** function is used to clear the AI control of a simulated object, typically an aircraft, in order for it to be controlled by a SimConnect client.

##### Syntax

```cpp
HRESULT SimConnect_AIReleaseControl(
    HANDLE  hSimConnect,
    SIMCONNECT_OBJECT_ID  ObjectID,
    SIMCONNECT_DATA_REQUEST_ID  RequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _ObjectID_ | Specifies the server defined object ID. | Integer |
| _RequestID_ | Specifies the client defined request ID. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

This function should be used to transfer the control of an aircraft, or other object, from the AI system to the SimConnect client. If this is not done the AI system and client may fight each other with unpredictable results. To prevent the simulation engine from updating the latitude, longitude, altitude and attitude of an aircraft, refer to the range of `FREEZE_*` [Event ID](../../../Key_Events/Key_Events.htm) s.

The object ID can be obtained in a number of ways, refer to the [SimConnect\_RequestDataOnSimObjectType](../Events_And_Data/SimConnect_RequestDataOnSimObjectType.htm) call, and also the use of the [SIMCONNECT\_RECV\_ASSIGNED\_OBJECT\_ID](../Structures_And_Enumerations/SIMCONNECT_RECV_ASSIGNED_OBJECT_ID.htm) structure.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AICreateSimulatedObject](SimConnect_AICreateSimulatedObject.htm)
4. [SimConnect\_AIRemoveObject](SimConnect_AIRemoveObject.htm)
5. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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