SimConnect\_AICreateParkedATCAircraft\_EX1

## SimConnect\_AICreateParkedATCAircraft\_EX1

The **SimConnect\_AICreateParkedATCAircraft** function is used to create an AI controlled aircraft that is currently parked and does not have a flight plan. This function can be used with both _legacy_ SimObjects, and **modular** SimObjects.

##### Syntax

```cpp
HRESULT SimConnect_AICreateParkedATCAircraft_EX1(
    HANDLE  hSimConnect,
    const char*  szContainerTitle,
    const char*  szLivery,
    const char*  szTailNumber,
    const char*  szAirportID,
    SIMCONNECT_DATA_REQUEST_ID  RequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _szContainerTitle_ | Null-terminated string containing the container title. The container title is found in the [title](../../../../5_Content_Configuration/CFG_Files/aircraft_cfg.htm#title) parameter of the `aircraft.cfg` file, or alternatively, you can find it from the list found in the [Aircraft Capture Tool](../../../../2_DevMode/Menus/Tools.htm#aircraft_capture).<br>**IMPORTANT!** This title must be in an aircraft [Preset](../../../../5_Content_Configuration/Modular_SimObjects/Modular_SimObject_Project_Structure.htm#presets).You may also retrieve the information with SimConnect using the `SimConnect_EnumerateSimObjectsAndLiveries` function.<br>Examples of aircraft titles:<br>1. `title=Boeing 747-8f Asobo`<br>2. `title=DA62 Asobo`<br>3. `title=VL3 Asobo` | String |
| _szLivery_ | Null-terminated string containing the livery name. This field is only required when checking **modular** SimObjects, since their liveries are dynamically handled and do not have a unique [title](../../../../5_Content_Configuration/CFG_Files/aircraft_cfg.htm#title) parameter like legacy aircraft.<br>This parameter can be set to either the _name_ of the livery as defined in the `name` parameter of the `livery.cfg`, or to the _folder name_ of the livery in the package. Note that you can also get the livery name using the `SimConnect_EnumerateSimObjectsAndLiveries` function.<br>The function will check for the **folder** first, then the **parameter name** second, and if the given name is not found in either location, then the default livery will be used. | String |
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

When allocating a parking space to an aircraft, the simulation uses a radius based on half the wingspan defined in the `[AIRPLANE_GEOMETRY]` section of the `flight_model.cfg`.

Refer also to the remarks for [SimConnect\_AICreateEnrouteATCAircraft\_EX1](SimConnect_AICreateEnrouteATCAircraft_EX1.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AICreateEnrouteATCAircraft\_EX1](SimConnect_AICreateEnrouteATCAircraft_EX1.htm)
4. [SimConnect\_AICreateNonATCAircraft\_EX1](SimConnect_AICreateNonATCAircraft_EX1.htm)
5. [SimConnect\_AICreateSimulatedObject\_EX1](SimConnect_AICreateSimulatedObject_EX1.htm)
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