SimConnect Actions

## SIMCONNECT ACTIONS

**NOTE**: This page is currently WIP and will be improved and finalised in future updates.

On this page you can find listed all the available **actions** that can be used by the `SimConnect_ExecuteAction` function, along with the parameters that they require.

### OneShotSoundAction

Plays the specified sound file.

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `SoundFileName` | Text | Filename of the sound file. |

### DialogAction

Will display and play a dialog in the adventure window.

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `Text` | Text |  |
| `SoundFileName` | Text | Filename of the sound file |
| `PackageName` | Text | Filename of the package |
| `DelaySeconds` | Float | Number of seconds to delay |
| `TargetPlayer` | GUID | One of the specially defined GUIDs to refer to relative targets:<br>1. **Local User**: {8B7615FA-BBBF-4baa-B959-5EF4F59BB575}<br>2. **All Players**: {D09ADD85-4CA8-4be9-939C-4544BA259DA2}<br>3. **Trigger**: {9A82F89B-F271-4285-AE90-1F733945D975}<br>4. **All Players Except Trigger**: {D2764AC8-FBB3-4b9b-9721-3C3EA0CB402B} |
| `Skippable` | Bool | If true the animation is skippable |
| `Speaker` | Text | Speaker |
| `Priority` | Enum | Priority for a dialog |
| `CanBeSuspended` | Bool | A dialog can be suspended |

### AITakeControlsAction

AI takes controls.

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `AIControl` | Enum | True: AI takes controls. False: Human has controls. Default: Assistance value |
| `SimActionsLimited` | Bool | True: AI controls are limited. False: Sim has all the controls. |
| `DisableCockpitInteraction` | Bool | Disable cursor interaction in cockpit view |

### ChangeAssistanceItemAction

Change assistance item setting.

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `SetChangeable` | Enum |  |
| `SetValue` | Enum |  |
| `AssistanceItemID` | Enum | Enum of assistance item IDs |

### FailureAction

Fail a system.

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `System` | Enum |  |
| `SystemIndex` | LONG |  |
| `Behavior` | Enum |  |
| `HealthPercent` | Float |  |
| `TargetPlayer` | GUID | One of the specially defined GUIDs to refer to relative targets:<br>1. **Local User**: {8B7615FA-BBBF-4baa-B959-5EF4F59BB575}<br>2. **All Players**: {D09ADD85-4CA8-4be9-939C-4544BA259DA2}<br>3. **Trigger**: {9A82F89B-F271-4285-AE90-1F733945D975}<br>4. **All Players Except Trigger**: {D2764AC8-FBB3-4b9b-9721-3C3EA0CB402B} |

### PlaySoundAction

Play a sound

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `NodeName` | Text | Name of a node inside an scene hierarchy. |
| `WorldPosition` | LLA | Position of object in 3 Space |
| `Orientation` | PBH | Orientation of object in 3 Space |
| `WwiseEventName` | Text | Name of a Wwise event |

### StopSoundAction

Stop a sound

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `NodeName` | Text | Name of a node inside an scene hierarchy. |
| `FadeOutTime` | Float | Fade out time (in sec) |

### SoundEffectAction

Sets sound effects

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `DopplerEnabled` | Bool | Doppler enabled |

### RTPCAction

Send a RTPC

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `NodeName` | Text | Name of a node inside an scene hierarchy. |
| `WwiseRTPCValue` | Float | Value of a Wwise RTPC |

### RumbleAction

Controller rumble effect

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `Intensity` | Enum | Intensity of rumble effect |
| `Duration` | Float | Duration in seconds |

### RefillAction

Add or subtract percent of a substance

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `PercentFuel` | Float |  |
| `PercentNitrous` | Float |  |
| `PercentAntiDetonation` | Float |  |
| `SystemIndex` | LONG |  |
| `TargetPlayer` | GUID | One of the specially defined GUIDs to refer to relative targets:<br>1. **Local User**: {8B7615FA-BBBF-4baa-B959-5EF4F59BB575}<br>2. **All Players**: {D09ADD85-4CA8-4be9-939C-4544BA259DA2}<br>3. **Trigger**: {9A82F89B-F271-4285-AE90-1F733945D975}<br>4. **All Players Except Trigger**: {D2764AC8-FBB3-4b9b-9721-3C3EA0CB402B} |

### ShowLogbookAction

Show logbook

### FadeToColorAction

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `ColorRed` | ULONG | Red component between 0-255 |
| `ColorGreen` | ULONG | Green component between 0-255 |
| `ColorBlue` | ULONG | Blue component between 0-255 |
| `FadeToColor` | Bool | TRUE to fade to color. FALSE to fade from color |
| `FadeDuration` | Float | Duration of the fade |

### FocusInstrumentAction

Change marker template action

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `SetHighlight` | Bool | If true, Highlight the part is set. If false, stop Highlighting the part |
| `HighlightColor` | FLOAT4 | Highlight color. |
| `FocusDuration` | Float | Focus duration in seconds |
| `SetCamera` | Bool | If true, camera is set. See the [SetCamera Additional Information](SimConnect_Actions.htm#setcamera-info) section for more information. |
| `SetPulse` | Bool | If true, highlight pulses. |
| `SetEyeIcon` | Bool | If true, show an Eye Icon on focused part. |
| `InstrumentPartId` | Text | PartId of a cockpit instrument |
| `InstrumentHtmlId` | Text | Id of a cockpit html instrument |

### LookAtInstrumentAction

Indicate user to look at an instrument

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `SetHighlight` | Bool | If true, Highlight the part is set. If false, stop Highlighting the part |
| `HighlightColor` | FLOAT4 | Highlight color. |
| `FocusDuration` | Float | Focus duration in seconds |
| `InstrumentPartId` | Text | PartId of a cockpit instrument |
| `InstrumentHtmlId` | Text | Id of a cockpit html instrument |

### InstructorDialogAction

Display and play dialog in the adventure window.

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `Text` | Text | Text to use with the action. |
| `SoundFileName` | Text | Filename of the sound file |
| `DelaySeconds` | Float | Number of seconds to delay |
| `TargetPlayer` | GUID | One of the specially defined GUIDs to refer to relative targets:<br>1. **Local User**: {8B7615FA-BBBF-4baa-B959-5EF4F59BB575}<br>2. **All Players**: {D09ADD85-4CA8-4be9-939C-4544BA259DA2}<br>3. **Trigger**: {9A82F89B-F271-4285-AE90-1F733945D975}<br>4. **All Players Except Trigger**: {D2764AC8-FBB3-4b9b-9721-3C3EA0CB402B} |
| `Skippable` | Bool | If true the animation is skippable |
| `Speaker` | Text | Speaker |
| `Priority` | Enum | Priority for a dialog |
| `CanBeSuspended` | Bool | A dialog can be suspended |

### RequestTeleportCameraAction

Teleport current camera to a position given by THE FIRST camera that validates ONE of the arguments

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `CameraNameCFG` | Text | Name of the camera in the CFG file to extract position |
| `TeleportCamera` | Bool | If false, the transition to the new camera is smooth |

### ForceMusicAction

Force Music Color

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `MusicSelection` | Text | Name of the music to play (as a string) |
| `MusicSelectionTT` | Text |  |

### StopMusicAction

Stop music action

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `MusicSelection` | Text | Name of the music to stop (as a string) |
| `FadeOutTime` | Float | Fade out time (in sec). |

### PlayMusicAction

Action to play some music.

This action has the following parameters:

| Name | Type | Description |
| --- | --- | --- |
| `MusicSelection` | Text | Name of the music to play (as a string) |
| `IsLooping` | Bool | Whether the music should loop or not. |
| `FadeOutTime` | Float | Fade out time (in sec). |

### TimeAction

action to execute in time

### SetCamera Additional Information

When using the [FocusInstrumentAction](SimConnect_Actions.htm#FocusInstrumentAction) method along with the `SetCamera` parameter, there a few things that will need to be taken into account before it will work correctly for all components. The steps are as follows:

- The first step is to add a new `[CAMERADEFINITION.N]` in the `cameras.cfg` file and place the camera where you want it to be when you use the `FocusInstrumentAction`. If there is an existing camera that also fits the same role then that camera can be used instead.

- Next, in the `<behavior>.xml` for the aircraft (or its cockpit counterpart), find the `<Component>` element which corresponds to the one that you will want to focus the camera on.

- Inside of this `<Component>` element, you will then need to add a `<CameraTitle>` sub-element with the title of the camera you want to use, as defined previously in the `cameras.cfg` file.

- The final XML should look something like this:


```cpp
<Component ID="EXT_18_NACELLE" Node="EXT_18_NACELLE">
<PartID>EXT_18_NACELLE</PartID>
<CameraTitle>EXT_18_NACELLE_CAMERA</CameraTitle>
<!-- EXT_18_NACELLE_CAMERA comes from cameras.cfg -->
</Component>
```


Related Topics

1. [SimConnect\_SDK](../../SimConnect_SDK.htm)
2. [SimConnect\_API\_Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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