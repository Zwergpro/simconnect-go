SimConnect\_AIRemoveObject

## SimConnect\_AIRemoveObject

The **SimConnect\_AIRemoveObject** function is used to remove any object created by the client using one of the AI creation functions.

##### Syntax

```cpp
HRESULT SimConnect_AIRemoveObject(
    HANDLE  hSimConnect,
    SIMCONNECT_OBJECT_ID  ObjectID,
    SIMCONNECT_DATA_REQUEST_ID  RequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _ObjectID_ | Specifies the server defined object ID (refer to the [SIMCONNECT\_RECV\_ASSIGNED\_OBJECT\_ID](../Structures_And_Enumerations/SIMCONNECT_RECV_ASSIGNED_OBJECT_ID.htm) structure). | Integer |
| _RequestID_ | Specifies the client defined request ID. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

A client application can only remove AI controlled objects that it created, not objects created by other clients, or by the simulation itself.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AICreateSimulatedObject](SimConnect_AICreateSimulatedObject.htm)
4. [SimConnect\_AIReleaseControl](SimConnect_AIReleaseControl.htm)
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