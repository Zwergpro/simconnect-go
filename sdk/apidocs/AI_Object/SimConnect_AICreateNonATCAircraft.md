SimConnect\_AICreateNonATCAircraft

## SimConnect\_AICreateNonATCAircraft

The **SimConnect\_AICreateNonATCAircraft** function is used to create an aircraft that is not flying under ATC control (so is typically flying under VFR rules).

**NOTE**: This function is a **legacy** function and only works with **non-modular aircraft**. In Microsoft Flight Simulator 2024 we recommend you use `SimConnect_AICreateNonATCAircraft_EX1`, which can be used with legacy **and** modular aircraft.

##### Syntax

```cpp
HRESULT SimConnect_AICreateNonATCAircraft(
    HANDLE  hSimConnect,
    const char*  szContainerTitle,
    const char*  szTailNumber,
    SIMCONNECT_DATA_INITPOSITION  InitPos,
    SIMCONNECT_DATA_REQUEST_ID  RequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _szContainerTitle_ | Null-terminated string containing the container title. The container title is found in the [aircraft.cfg](../../../../5_Content_Configuration/CFG_Files/ai_cfg.htm) file.<br>**IMPORTANT!** This title must be in an aircraft [Preset](../../../../5_Content_Configuration/Modular_SimObjects/Modular_SimObject_Project_Structure.htm#presets).Alternatively, the aircraft title can be obtained via the Aircraft Selector ( _DevMode_-> _Windows_-> _Aircraft_ selector). Finally, the information can be found using the `SimConnect_EnumerateSimObjectsAndLiveries` function.<br>Examples of aircraft titles:<br>1. `title=Boeing 747-8f Asobo`<br>2. `title=DA62 Asobo`<br>3. `title=VL3 Asobo` | String |
| _szTailNumber_ | Null-terminated string containing the tail number. This should have a maximum of 12 characters. | String |
| _InitPos_ | Specifies the initial position, using a SIMCONNECT\_DATA\_INITPOSITION structure. | Integer |
| _RequestID_ | Specifies the client defined request ID. | String |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

A non-ATC aircraft can be on the ground or airborne when it is created by this function. The following exception can be caused by this function (refer to the [SIMCONNECT\_EXCEPTION](../Structures_And_Enumerations/SIMCONNECT_EXCEPTION.htm) enum for more details):

- `SIMCONNECT_EXCEPTION_CREATE_OBJECT_FAILED`

Refer to the remarks for [SimConnect\_AICreateEnrouteATCAircraft](SimConnect_AICreateEnrouteATCAircraft.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AICreateEnrouteATCAircraft](SimConnect_AICreateEnrouteATCAircraft.htm)
4. [SimConnect\_AICreateParkedATCAircraft](SimConnect_AICreateParkedATCAircraft.htm)
5. [SimConnect\_AISetAircraftFlightPlan](SimConnect_AISetAircraftFlightPlan.htm)
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