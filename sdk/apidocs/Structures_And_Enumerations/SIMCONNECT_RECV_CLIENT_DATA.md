SIMCONNECT\_RECV\_CLIENT\_DATA

## SIMCONNECT\_RECV\_CLIENT\_DATA

The **SIMCONNECT\_RECV\_CLIENT\_DATA** structure will be received by the client after a successful call to [SimConnect\_RequestClientData](../Events_And_Data/SimConnect_RequestClientData.htm). It is an identical structure to SIMCONNECT\_RECV\_SIMOBJECT\_DATA.

##### Syntax

```cpp
struct SIMCONNECT_RECV_CLIENT_DATA : public SIMCONNECT_RECV_SIMBOBJECT_DATA{
    };
```

##### Remarks

This structure inherits the [`SIMCONNECT_RECV_SIMOBJECT_DATA`](SIMCONNECT_RECV_SIMOBJECT_DATA.htm) structure and is returned when the `dwID` parameter of [`SIMCONNECT_RECV`](SIMCONNECT_RECV.htm) is set to [`SIMCONNECT_RECV_ID_CLIENT_DATA`](SIMCONNECT_RECV_ID.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_RECV](SIMCONNECT_RECV.htm)
4. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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