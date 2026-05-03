Input Events

## INPUT EVENTS

The following SimConnect functions are used when dealing with inputs and input events. For a list of other functions please see the [SimConnect API Reference](../../SimConnect_API_Reference.htm).

| Function | Description |
| --- | --- |
| `SimConnect_EnumerateControllers` | This is used to retrieve a list of every device that is currently plugged into the simulation. |
| `SimConnect_EnumerateInputEvents` | Used to retrieve a paginated list of all available InputEvents for the current aircraft along with their associated hash (CRC based). |
| `SimConnect_EnumerateInputEventParams` | Used to retrieve a list of all parameters from an input event. |
| `SimConnect_GetInputEvent` | Used to retrieve the value of a specific input event (identified by its hash). |
| `SimConnect_MapInputEventToClientEvent` | Used to connect input events (such as keystrokes, joystick or mouse movements) with the sending of appropriate event notifications.<br>IMPORTANT! This function is deprecated due to a bug and you should always use `SimConnect_MapInputEventToClientEvent_EX1`. |
| `SimConnect_MapInputEventToClientEvent_EX1` | Used to connect input events (such as keystrokes, joystick or mouse movements) with the sending of appropriate event notifications. |
| `SimConnect_SetInputEvent` | Used to set the value of a specific input event (identified by its hash). |
| `SimConnect_SubscribeInputEvent` | Used to subscribe an input event and generate when the value changes. |
| `SimConnect_UnsubscribeInputEvent` | Used to unsubscribe from an input event that has previously been subscribed to. |
| `SimConnect_ClearInputGroup` | Used to remove all the input events from a specified input group object. |
| `SimConnect_SetInputGroupPriority` | Used to set the priority for a specified input group object. |
| `SimConnect_SetInputGroupState` | Used to turn requests for input event information from the server on and off. |
| `SimConnect_RemoveInputEvent` | Used to remove an input event from a specified input group object. |

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