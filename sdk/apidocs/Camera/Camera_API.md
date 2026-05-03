Camera API

## CAMERA API

**IMPORTANT!** This API is currently in **beta** and may be subject to change.

The **Camera API** is used for manipulating and controlling the camera using [SimConnect SDK](../../SimConnect_SDK.htm). There is also a corresponding API for WASM, which you can find here:

- [Camera API](../../../WASM/Camera_API/Camera_API.htm)

| Function | Description |
| `SimConnect_CameraAcquire` | Used to acquire the add-on camara, if possible. |
| `SimConnect_CameraDisableFlag` | Used to disable camera specific features, but only if the camera has been correctly acquired. |
| `SimConnect_CameraEnableFlag` | Used to enable camera specific features, but only if the camera has been correctly acquired. |
| `SimConnect_CameraGet` | Used to get the add-on camera settings, regardless of whether it has been acquired or not. |
| `SimConnect_CameraGetStatus` | Used to get the add-on camera status. |
| `SimConnect_CameraRelease` | Release a previously acquired camera when no longer required or for other systems to acquire it. |
| `SimConnect_CameraSet` | Used to set the add-on camera settings, if it has been correctly acquired. |
| `SimConnect_CameraSetUsingCameraDefinition` | Used to set the add-on camera settings using parameters taken from a predefined camera. |
| `SimConnect_DeleteCameraWorldLocker` | Used to delete a world locker previously set in the world. |
| `SimConnect_EnumerateCameraDefinitions` | Used to retrieve an array of all the defined camera names. |
| `SimConnect_RequestCameraWorldLocker` | Used to set a locker to load the world around it. |
| `SimConnect_SubscribeToCameraWorldLockerStatusUpdate` | Used to subscribe to camera world locker status update messages. |
| `SimConnect_SubscribeToCameraStatusUpdate` | Used to subscribe to camera status update messages. |
| `SimConnect_UnsubscribeToCameraWorldLockerStatusUpdate` | Used to unsubscribe from camera world locker update messages. |
| `SimConnect_UnsubscribeToCameraStatusUpdate` | Used to unsubscribe from camera update messages. |

The camera these functions access is a unique camera instance that exists _only_ to be used by add-ons, and it can be acquired using the functions listed below. Camera acquisition is based on a **queue**, which means that if a client _A_ has acquired the camera, another client _B_ can later "steal" the camera, in which case client _A_ will receive a message informing them that another client possesses the camera. Once client _B_ releases the camera, it will be given back to client _A_, unless that client has also released it.

It is important to note that just because a client has acquired the camera, it does not mean that the client as _all_ rights on the simulation's cameras and the following circumstances can revoke that control, either temporarily or permanently:

- **The simulation can temporarily take back control of the cameras** during specific parts of the flow (during menu interactions, when the simulation is paused, during RTC's etc…). Outside of these specific situations, control of cameras will be given back to the add-on which had acquired the dedicated camera previously.
- **The user has all rights over the cameras**. Through the simulation camera UI panel, the user can revoke the acquisition of the add-on camera and/or forbid its future acquisition.

Under either of the above circumstances, appropriate information will be sent to all clients which have subscribed to the camera event, even if they haven't acquired the camera, so that measures can be taken.

#### World Lockers

The camera API also permits the placement of a **world locker** to "lock" the world at a location before the camera is moved there. This locker will essentially "lock" the scenery and SimObjects at the location in memory so the user doesn't experience loading issues and "popping" when the camera moves to the location. It is important to note the following caveats:

- A world locker cannot be placed in the world unless the _camera has first been acquired by the add-on_.
- Lockers are _very resource intensive_ and should only be used sparingly, and they should be removed the moment they are no longer required.

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