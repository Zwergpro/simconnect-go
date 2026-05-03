SimConnect\_AISetAircraftFlightPlan

## SimConnect\_AISetAircraftFlightPlan

The **SimConnect\_AISetAircraftFlightPlan** function is used to set or change the flight plan of an AI controlled aircraft.

##### Syntax

```cpp
HRESULT SimConnect_AISetAircraftFlightPlan(
        HANDLE  hSimConnect,
        SIMCONNECT_OBJECT_ID  ObjectID,
        const char*  szFlightPlanPath,
        SIMCONNECT_DATA_REQUEST_ID  RequestID
      );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _ObjectID_ | Specifies the server defined object ID. | Integer |
| _szFlightPlanPath_ | Null-terminated string containing the path to the flight plan file. Flight plans have the extension .pln, but no need to enter an extension here. The easiest way to create flight plans is to create them from within the simulation, and then save them off for use with the AI controlled aircraft. | String |
| _RequestID_ | Specifies the client defined request ID. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

The following errors may apply to AI objects (refer to the [SIMCONNECT\_EXCEPTION](../Structures_And_Enumerations/SIMCONNECT_EXCEPTION.htm) enum for more details):

- `SIMCONNECT_EXCEPTION_LOAD_FLIGHTPLAN_FAILED`
- `SIMCONNECT_EXCEPTION_OPERATION_INVALID_FOR_OBJECT_TYPE`
- `SIMCONNECT_EXCEPTION_ERROR`

Typically this function would be used some time after the aircraft was created using the [SimConnect\_AICreateParkedATCAircraft](SimConnect_AICreateParkedATCAircraft.htm) call.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AICreateEnrouteATCAircraft](SimConnect_AICreateEnrouteATCAircraft.htm)
4. [SimConnect\_AICreateNonATCAircraft](SimConnect_AICreateNonATCAircraft.htm)
5. [SimConnect\_AICreateParkedATCAircraft](SimConnect_AICreateParkedATCAircraft.htm)
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