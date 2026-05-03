Facilities

## FACILITIES

The following SimConnect functions are used to get, set and manipulate data related to the different simulation facilities. For a list of other functions please see the [SimConnect API Reference](../../SimConnect_API_Reference.htm).

| Function | Description |
| --- | --- |
| `SimConnect_AddToFacilityDefinition` | This function is used to create a facility data definition. |
| `SimConnect_AddFacilityDataDefinitionFilter` | Used add a filter on a node in the FacilityDataDefinition to block sending data according to this filter, thus reduce the amount of data received and limit it to only that which is required. |
| `SimConnect_ClearAllFacilityDataDefinitionFilters` | Used to clear all applied facility definition filters. |
| `SimConnect_RequestFacilitesList` | Request a list of all the facilities of a given type currently held in the facilities cache. |
| `SimConnect_RequestFacilitiesList_EX1` | Alias for the `SimConnect_RequestFacilitesList`. Should be used with the other `_EX1` functions. |
| `SimConnect_RequestFacilityData` | Used to request data according to a predefined object, an ICAO and a region. |
| `SimConnect_RequestFacilityData_EX1` | This function is practically identical in functionality to the `SimConnect_RequestFacilityData` function, only it has an additional return value used to identify waypoints when there is an ICAO/Region overlap with VOR or NDB. |
| `SimConnect_RequestJetwayData` | Used to request data from an airport about the jetways available. |
| `SimConnect_SubscribeToFacilities` | Used to request notifications when a facility of a certain type is added to the facilities cache. |
| `SimConnect_SubscribeToFacilities_EX1` | Used to request notifications when a facility of a certain type is added to the facilities cache, as well as provide callbacks for when an element is added/removed. |
| `SimConnect_UnsubscribeToFacilities` | Used to request that notifications of additions to the facilities cache are not longer sent. |
| `SimConnect_UnsubscribeToFacilities_EX1` | Used to request that notifications of additions or removals to the facilities cache are not longer sent. |

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