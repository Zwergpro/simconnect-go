General

## GENERAL CALLS

The following SimConnect functions are general in nature and essential for using SimConnect. For a list of other functions please see the [SimConnect API Reference](../../SimConnect_API_Reference.htm).

| Function | Description |
| `DispatchProc` | Written by the developer of the SimConnect client, as a callback function to handle all the communications with the server |
| `SimConnect_Open` | Used to send a request to the Flight Simulator server to open up communications with a new client. |
| `SimConnect_Close` | Used to request that the communication with the server is ended. |
| `SimConnect_CallDispatch` | Used to process the next SimConnect message received through the specified callback function. |
| `SimConnect_GetNextDispatch` | Used to process the next SimConnect message received, without the use of a callback function. |
| `SimConnect_RequestSystemState` | Used to request information from a number of Flight Simulator system components. |
| `SimConnect_SetNotificationGroupPriority` | Used to set the priority of a notification group. |
| `SimConnect_ExecuteAction` | Used to call an XML action with a variable number of parameters. Available actions are listed here: [SimConnect Actions](../SimConnect_Actions.htm) |

### System Events

| Function | Description |
| --- | --- |
| `SimConnect_SubscribeToSystemEvent` | Used to request that a specific system event is notified to the client. |
| `SimConnect_SetSystemEventState` | Used to turn requests for event information from the server on and off. |
| `SimConnect_UnsubscribeFromSystemEvent` | Used to request that notifications are no longer received for the specified system event. |

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