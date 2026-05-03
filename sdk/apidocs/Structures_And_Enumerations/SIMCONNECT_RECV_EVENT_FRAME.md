SIMCONNECT\_RECV\_EVENT\_FRAME

## SIMCONNECT\_RECV\_EVENT\_FRAME

The **SIMCONNECT\_RECV\_EVENT\_FRAME** structure is used with the [SimConnect\_SubscribeToSystemEvent](../General/SimConnect_SubscribeToSystemEvent.htm) call to return the frame rate and simulation speed to the client.

##### Syntax

```cpp
struct SIMCONNECT_RECV_EVENT_FRAME : SIMCONNECT_RECV_EVENT {
        float  fFrameRate;
        float  fSimSpeed;
      };
```

##### Members

| Member | Description |
| `fFrameRate` | The visual frame rate in frames per second. |
| `fSimSpeed` | The simulation rate. For example if the simulation is running at four times normal speed - 4X - then 4.0 will be returned. |

##### Remarks

This structure inherits the `SIMCONNECT_RECV_EVENT` structure, which inherits the [`SIMCONNECT_RECV`](SIMCONNECT_RECV.htm) structure, and is returned when the `dwID` parameter of [`SIMCONNECT_RECV`](SIMCONNECT_RECV.htm) is set to [`SIMCONNECT_RECV_ID_EVENT_FRAME`](SIMCONNECT_RECV_ID.htm). Set the requested system event to "Frame" or "PauseFrame" with the [SimConnect\_SubscribeToSystemEvent](../General/SimConnect_SubscribeToSystemEvent.htm) function to receive this data.

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