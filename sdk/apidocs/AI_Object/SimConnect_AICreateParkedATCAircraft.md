SimConnect\_AICreateParkedATCAircraft

## SimConnect\_AICreateParkedATCAircraft

The **SimConnect\_AICreateParkedATCAircraft** function is used to create an AI controlled aircraft that is currently parked and does not have a flight plan.

**NOTE**: This function is a **legacy** function and only works with **non-modular aircraft**. In Microsoft Flight Simulator 2024 we recommend you use `SimConnect_AICreateParkedATCAircraft_EX1`, which can be used with legacy **and** modular aircraft.

##### Syntax

```cpp
HRESULT SimConnect_AICreateParkedATCAircraft(
    HANDLE  hSimConnect,
    const char*  szContainerTitle,
    const char*  szTailNumber,
    const char*  szAirportID,
    SIMCONNECT_DATA_REQUEST_ID  RequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _szContainerTitle_ | Null-terminated string containing the container title. The container title is found in the [ai.cfg](../../../../5_Content_Configuration/CFG_Files/ai_cfg.htm) file.<br>**IMPORTANT!** This title must be in an aircraft [Preset](../../../../5_Content_Configuration/Modular_SimObjects/Modular_SimObject_Project_Structure.htm#presets).Alternatively, the aircraft title can be obtained via the Aircraft Selector ( _DevMode_-> _Windows_-> _Aircraft_ selector). Finally, the information can be found using the `SimConnect_EnumerateSimObjectsAndLiveries` function.<br>Examples of aircraft titles:<br>`title=Boeing 747-8f Asobo`<br>`title=DA62 Asobo`<br>`title=VL3 Asobo` | String |
| _szTailNumber_ | Null-terminated string containing the tail number. This should have a maximum of 12 characters. | String |
| _szAirportID_ | Null-terminated string containing the airport ID. This is the ICAO code string, for example, KSEA for SeaTac International. | String |
| _RequestID_ | Specifies the client defined request ID. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

Calling this function is no guarantee that there is sufficient parking space at the specified airport. An error will be returned if there is insufficient parking space, and an aircraft will not be created. The following exceptions can be created by this function (refer to the [SIMCONNECT\_EXCEPTION](../Structures_And_Enumerations/SIMCONNECT_EXCEPTION.htm) enum for more details):

- `SIMCONNECT_EXCEPTION_OBJECT_OUTSIDE_REALITY_BUBBLE`
- `SIMCONNECT_EXCEPTION_OBJECT_CONTAINER`
- `SIMCONNECT_EXCEPTION_OBJECT_AI`
- `SIMCONNECT_EXCEPTION_OBJECT_ATC`
- `SIMCONNECT_EXCEPTION_OBJECT_SCHEDULE`
- `SIMCONNECT_EXCEPTION_CREATE_OBJECT_FAILED`

After creating an aircraft with this function, a call to [SimConnect\_AISetAircraftFlightPlan](SimConnect_AISetAircraftFlightPlan.htm) will set the aircraft in motion.

When allocating a parking space to an aircraft, the simulation uses a radius based on half the wingspan defined in the `[aircraft_geometry]` section of the [Aircraft Configuration File.](../../../../5_Content_Configuration/CFG_Files/aircraft_cfg.htm)

Refer also to the remarks for [SimConnect\_AICreateEnrouteATCAircraft](SimConnect_AICreateEnrouteATCAircraft.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AICreateEnrouteATCAircraft](SimConnect_AICreateEnrouteATCAircraft.htm)
4. [SimConnect\_AICreateNonATCAircraft](SimConnect_AICreateNonATCAircraft.htm)
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