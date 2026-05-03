SIMCONNECT\_RECV\_CONTROLLERS\_LIST

## SIMCONNECT\_RECV\_CONTROLLERS\_LIST

The **SIMCONNECT\_RECV\_CONTROLLERS\_LIST** structure is used to return an array of data related to available controllers.

##### Syntax

```cpp
struct SIMCONNECT_RECV_CONTROLLERS_LIST : public SIMCONNECT_RECV {
    SIMCONNECT_CONTROLLER_ITEM rgData[dwArraySize];
};
```

##### Members

| Member | Description |
| `rgData[dwArraySize]` | Array of `SIMCONNECT_CONTROLLER_ITEM` structures. |

##### Remarks

This structure is retrieved using the `SimConnect_EnumerateControllers` function and will contain an array of `SIMCONNECT_CONTROLLER_ITEM` structures, where each entry corresponds to a detected controller that is currently connected to the simulation.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_EnumerateControllers](../InputEvents/SimConnect_EnumerateControllers.htm)
4. [SIMCONNECT\_CONTROLLER\_ITEM](SIMCONNECT_CONTROLLER_ITEM.htm)
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