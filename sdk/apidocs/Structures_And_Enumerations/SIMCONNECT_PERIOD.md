SIMCONNECT\_PERIOD

## SIMCONNECT\_PERIOD

The **SIMCONNECT\_PERIOD** enumeration type is used with the [SimConnect\_RequestDataOnSimObject](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm) call to specify how often data is to be sent to the client.

##### Syntax

```cpp
enum SIMCONNECT_PERIOD {
    SIMCONNECT_PERIOD_NEVER,
    SIMCONNECT_PERIOD_ONCE,
    SIMCONNECT_PERIOD_VISUAL_FRAME,
    SIMCONNECT_PERIOD_SIM_FRAME,
    SIMCONNECT_PERIOD_SECOND,
    }
```

##### Members

| Member | Description |
| `SIMCONNECT_PERIOD_NEVER` | Specifies that the data is not to be sent. |
| `SIMCONNECT_PERIOD_ONCE` | Specifies that the data should be sent once only. Note that this is not an efficient way of receiving data frequently, use one of the other periods if there is a regular frequency to the data request. |
| `SIMCONNECT_PERIOD_VISUAL_FRAME` | Specifies that the data should be sent every visual (rendered) frame. |
| `SIMCONNECT_PERIOD_SIM_FRAME` | Specifies that the data should be sent every simulated frame, whether that frame is rendered or not. |
| `SIMCONNECT_PERIOD_SECOND` | Specifies that the data should be sent once every second. |

##### Remarks

Although the period definitions are specific, data is always transmitted at the end of a frame, so even if you have specified that data should be sent every second, the data will actually be transmitted at the end of the frame that comes on or after one second has elapsed.

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