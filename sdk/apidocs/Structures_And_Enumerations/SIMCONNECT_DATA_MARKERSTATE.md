SIMCONNECT\_DATA\_MARKERSTATE

## SIMCONNECT\_DATA\_MARKERSTATE

The **SIMCONNECT\_DATA\_MARKERSTATE** structure is used to help graphically link flight model data with the graphics model.

##### Syntax

```cpp
struct SIMCONNECT_DATA_MARKERSTATE{
    char  szMarkerName[64];
    DWORD  dwMarkerState;
    };
```

##### Members

| Member | Description |
| `szMarkerName[64]` | Null-terminated string containing the marker name. One from the following list:<br>`Cg`<br>`ModelCenter`<br>`Wheel`<br>`Skid`<br>`Ski`<br>`Float`<br>`Scrape`<br>`Engine`<br>`Prop`<br>`Eyepoint`<br>`LongScale`<br>`LatScale`<br>`VertScale`<br>`AeroCenter`<br>`WingApex`<br>`RefChord`<br>`Datum`<br>`WingTip`<br>`FuelTank`<br>`Forces` |
| `dwMarkerState` | Double word containing the marker state, set to 1 for on and 0 for off. |

##### Remarks

The `SimConnect_AddToDataDefinition` call can be used to add a `SIMCONNECT_DATA_MARKERSTATE` structure to a data definition. Use of this call and structure is to help determine that points specified in the flight model of an aircraft match the graphics model for that aircraft, by turning on the specified marker lights. A SimConnect client created to do this becomes a tool to aid to the accurate development of aircraft models, rather than an add-on that an end user might run.

This structure can only be used as input, it cannot be used as part of a data request.

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