Events And Data

## EVENTS AND DATA

The following SimConnect functions are used for setting, retrieving and generally manipulating different data sets. For a list of other functions please see the [SimConnect API Reference](../../SimConnect_API_Reference.htm).

### Flow API

The functions here are designed to warn add-on packages about the different states of the simulation flow and the potential [FLT files](../../../../5_Content_Configuration/FLT_Files/FLT_Properties.htm) that are loaded. The Flow API is a "global" API that is mainly used for dealing with [Back On Track](../../../../5_Content_Configuration/FLT_Files/Back_On_Track.htm) events, and it is available across these platforms:

- [WebAssembly](../../../WASM/WebAssembly.htm)
- [JavaScript](../../../JavaScript/JavaScript.htm)
- [SimConnect SDK](../../SimConnect_API_Reference.htm#flow)

| Function | Description |
| `SimConnect_SubscribeToFlowEvent` | Used to subscribe to the flow events so you can receive messages when the simulation triggers certain of these events. |
| `SimConnect_UnsubscribeToFlowEvent` | Used to request that notifications are no longer received for subscribed flow events. |

### Client Events

| Function | Description |
| --- | --- |
| `SimConnect_AddToDataDefinition` | Used to add a Flight Simulator simulation variable name to a client defined object definition. |
| `SimConnect_ClearDataDefinition` | Used to remove all simulation variables from a client defined object. |
| `SimConnect_AddClientEventToNotificationGroup` | Used to add an individual client defined event to a notification group. |
| `SimConnect_RemoveClientEvent` | Used to remove a client defined event from a notification group. |
| `SimConnect_TransmitClientEvent` | Used to request that the Flight Simulator server transmit to all SimConnect clients the specified client event along with a single event parameter. |
| `SimConnect_TransmitClientEvent_EX1` | Used to request that the Flight Simulator server transmit to all SimConnect clients the specified client event with up to five event parameters. |
| `SimConnect_MapClientDataNameToID` | Used to associate an ID with a named client date area. |
| `SimConnect_MapClientEventToSimEvent` | Used to associate a client defined event ID with a Flight Simulator event name. |
| `SimConnect_RequestClientData` | Used to request that the data in an area created by another client be sent to this client. |
| `SimConnect_CreateClientData` | Used to request the creation of a reserved data area for this client. |
| `SimConnect_AddToClientDataDefinition` | Used to add an offset and a size in bytes, or a type, to a client data definition. |
| `SimConnect_SetClientData` | Used to write one or more units of data to a client data area. |
| `SimConnect_ClearClientDataDefinition` | Used to clear the definition of the specified client data. |

### SimObjects

| Function | Description |
| --- | --- |
| `SimConnect_EnumerateSimObjectsAndLiveries` | Used to retrieve a list of spawnable SimObjects, and - if applicable - the name of their liveries. |
| `SimConnect_RequestDataOnSimObject` | Used to request when the SimConnect client is to receive data values for a specific object. |
| `SimConnect_RequestDataOnSimObjectType` | Used to retrieve information about simulation objects of a given type that are within a specified radius of the user's aircraft. |
| `SimConnect_SetDataOnSimObject` | Used to make changes to the data properties of an object. |

### Misc

| Function | Description |
| --- | --- |
| `SimConnect_RequestNotificationGroup` | Used to request events from a notification group when the simulation is in Dialog Mode. |
| `SimConnect_ClearNotificationGroup` | Used to remove all the client defined events from a notification group. |
| `SimConnect_RequestReservedKey` | Used to request a specific keyboard TAB-key combination applies only to this client. |

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