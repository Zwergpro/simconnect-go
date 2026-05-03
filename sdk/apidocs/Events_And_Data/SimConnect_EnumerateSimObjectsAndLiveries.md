SimConnect\_EnumerateSimObjectsAndLiveries

## SimConnect\_EnumerateSimObjectsAndLiveries

The **SimConnect\_EnumerateSimObjectsAndLiveries** function is used to retrieve the list of spawnable SimObjects (and the name of their livery if applicable) that can be used with `SimConnect_AICreate_*` functions.

##### Syntax

```cpp
HRESULT SimConnect_EnumerateSimObjectsAndLiveries(
    HANDLE  hSimConnect,
    SIMCONNECT_DATA_REQUEST_ID  RequestID,
    SIMCONNECT_SIMOBJECT_TYPE Type
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _RequestId_ | Specifies the client defined request ID. | Integer |
| _Type_ | Specifies the type of SimObjects that will be retrieved by the function:<br>- `SIMCONNECT_SIMOBJECT_TYPE_USER`: SimObject types selectable by the user (airplane, helicopter and hot air balloon)<br>- `SIMCONNECT_SIMOBJECT_TYPE_ALL`: All types listed in this note.<br>- `SIMCONNECT_SIMOBJECT_TYPE_AIRCRAFT`: Retrieve only airplanes.<br>- `SIMCONNECT_SIMOBJECT_TYPE_HELICOPTER`: Retrieve only helicopters.<br>- `SIMCONNECT_SIMOBJECT_TYPE_BOAT`: Retrieve only boats.<br>- `SIMCONNECT_SIMOBJECT_TYPE_GROUND`: Retrieve only ground vehicles.<br>- `SIMCONNECT_SIMOBJECT_TYPE_HOT_AIR_BALLOON`: Retrieve only hot air balloons.<br>- `SIMCONNECT_SIMOBJECT_TYPE_ANIMAL`: Retrieve only animals.<br>- `SIMCONNECT_SIMOBJECT_TYPE_USER_AVATAR`: Retrieve the user avatar. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

This function should be used to retrieve the list of spawnable SimObjects (and the name of their liveries if applicable) that are mandatory information used by `SimConnect_AICreate*` functions. This function can send two kinds of messages :

- A `SIMCONNECT_RECV_ENUMERATE_SIMOBJECT_AND_LIVERY_LIST` message which contains the list (or a part of the list) of all possible combinations.
- A `SIMCONNECT_EXCEPTION` message with the `SIMCONNECT_EXCPETION_INTERNAL` ID, which implies that something went wrong while creating the list to retrieve to the client (which should never happen).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AICreateEnrouteATCAircraft\_EX1](../AI_Object/SimConnect_AICreateEnrouteATCAircraft_EX1.htm)
4. [SimConnect\_AICreateNonATCAircraft\_EX1](../AI_Object/SimConnect_AICreateNonATCAircraft_EX1.htm)
5. [SimConnect\_AICreateParkedATCAircraft\_EX1](../AI_Object/SimConnect_AICreateParkedATCAircraft_EX1.htm)
6. [SimConnect\_AICreateSimulatedObject\_EX1](../AI_Object/SimConnect_AICreateSimulatedObject_EX1.htm)

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