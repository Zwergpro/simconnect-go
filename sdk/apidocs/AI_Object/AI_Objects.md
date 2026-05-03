AI Objects

## AI OBJECTS

The following SimConnect functions are used for setting, retrieving and generally manipulating data related to the different AI objects in the simulation. For a list of other functions please see the [SimConnect API Reference](../../SimConnect_API_Reference.htm).

| Function | Description |
| --- | --- |
| `SimConnect_AICreateEnrouteATCAircraft` | Used to create an AI controlled aircraft that is about to start or is already underway on its flight plan. |
| `SimConnect_AICreateEnrouteATCAircraft_EX1` |
| `SimConnect_AICreateNonATCAircraft` | Used to create an aircraft that is not flying under ATC control (so is typically flying under VFR rules). |
| `SimConnect_AICreateNonATCAircraft_EX1` |
| `SimConnect_AICreateParkedATCAircraft` | Used to create an AI controlled aircraft that is currently parked and does not have a flight plan. |
| `SimConnect_AICreateParkedATCAircraft_EX1` |
| `SimConnect_AICreateSimulatedObject` | Used to create AI controlled objects other than aircraft. |
| `SimConnect_AICreateSimulatedObject_EX1` |
| `SimConnect_AIReleaseControl` | Used to clear the AI control of a simulated object, typically an aircraft, in order for it to be controlled by a SimConnect client. |
| `SimConnect_AIRemoveObject` | Used to remove any object created by the client using one of the AI creation functions. |
| `SimConnect_AISetAircraftFlightPlan` | Used to set or change the flight plan of an AI controlled aircraft. |

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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