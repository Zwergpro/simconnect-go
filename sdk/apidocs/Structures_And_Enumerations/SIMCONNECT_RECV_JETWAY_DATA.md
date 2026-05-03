SIMCONNECT\_RECV\_JETWAY\_DATA

## SIMCONNECT\_RECV\_JETWAY\_DATA

The **SIMCONNECT\_RECV\_JETWAY\_DATA** structure is used to return a list of [SIMCONNECT\_JETWAY\_DATA](SIMCONNECT_JETWAY_DATA.htm) structures.

##### Syntax

```cpp
struct SIMCONNECT_RECV_JETWAY_DATA : public SIMCONNECT_RECV_LIST_TEMPLATE {
    SIMCONNECT_JETWAY_DATA rgData[dwArraySize]
    };
```

##### Members

| Member | Description |
| `rgData[dwArraySize]` | Array of [`SIMCONNECT_JETWAY_DATA`](SIMCONNECT_JETWAY_DATA.htm) structures. |

##### Remarks

This struct inherits members from the [`SIMCONNECT_RECV_LIST_TEMPLATE`](SIMCONNECT_RECV_LIST_TEMPLATE.htm) struct.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_RequestJetwayData](../Facilities/SimConnect_RequestJetwayData.htm)
4. [SIMCONNECT\_JETWAY\_DATA](SIMCONNECT_JETWAY_DATA.htm)
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