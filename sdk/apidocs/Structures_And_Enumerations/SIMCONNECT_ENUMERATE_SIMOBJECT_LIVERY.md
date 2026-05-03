SIMCONNECT\_ENUMERATE\_SIMOBJECT\_LIVERY

## SIMCONNECT\_ENUMERATE\_SIMOBJECT\_LIVERY

The **SIMCONNECT\_ENUMERATE\_SIMOBJECT\_LIVERY** struct contains data related to a single combination of a SimObject and its livery (if any).

##### Syntax

```cpp
struct SIMCONNECT_ENUMERATE_SIMOBJECT_LIVERY
{
    SIMCONNECT_STRING(AircraftTitle, 256);
    SIMCONNECT_STRING(LiveryName, 256);
};
```

##### Members

| Member | Description |
| `AircraftTitle` | A string that gives the container title of a SimObject, as found in the [title](../../../../5_Content_Configuration/CFG_Files/aircraft_cfg.htm#title) parameter of the `aircraft.cfg` file. |
| `LiveryName` | A string that gives the name of a livery, as defined in the `name` parameter of the `livery.cfg` file. |

##### Remarks

See `SimConnect_EnumerateSimObjectsAndLiveries` for more information.

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