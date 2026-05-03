SIMCONNECT\_DATA\_CAMERA

## SIMCONNECT\_DATA\_CAMERA

The **SIMCONNECT\_DATA\_CAMERA** structure is a response containing data on the current camera.

##### Syntax

```cpp
SIMCONNECT_STRUCT SIMCONNECT_DATA_CAMERA
{
    SIMCONNECT_DATA_XYZ Position;
    SIMCONNECT_POSITION_REFERENTIAL PositionReferential;
    DWORD PositionReferentialObjectId;
    SIMCONNECT_DATA_XYZ TargetedPos;
    SIMCONNECT_DATA_PBH Pbh;
    SIMCONNECT_POSITION_REFERENTIAL RotationReferential;
    DWORD RotationReferentialObjectId;
    double Fov;
};
```

##### Members

| Member | Description |
| `Position` | The camera position expressed as a `SIMCONNECT_DATA_XYZ` struct. The units and values returned will depend of the selected `Referential`. |
| `PositionReferential` | The reference used to define the `position` value. One of the following:<br>1. **SimObject**\- The position is an offset from the simobject position (expressed in meters).<br>    <br>2. **SimObject Datum**\- The position is an offset from the simobject [Datum Reference Point](../../../../5_Content_Configuration/Modular_SimObjects/Aircraft/Aircraft.htm#DRP) (expressed in meters).<br>    <br>3. **Eyepoint** \- The position is an offset from the aircraft pilot eyepoint (expressed in meters).<br>4. **World** \- The position is expressed using latitude, longitude, and altitude. |
| `PositionReferentialObjectId` | Ignored if is `positionReferential` set to World.<br>The object Id of the simobject that will be used to set position by SimObject, SimObject Datum and Eyepoint referential.<br>Set to 0 to focus user's aircraft. If the object Id is invalid, user's aircraft will be used instead. |
| `TargetedPos` | A 3 value vector (x, y, z) defining the "target" position for the camera. This is used to define a "look at" point for the camera direction/orientation. The units and values returned will depend of the selected `Referential`. Note that if you _set_ a target position, it will override any values given for `Pbh`. |
| `Pbh` | The camera pitch, bank, and heading, expressed as a `SIMCONNECT_DATA_PBH` struct. Note when setting values, if `TargetedPos` is set, it will override the values given here. |
| `RotationReferential` | The reference used to define the `pbh` and `targetedPosition` values. One of the following:<br>1. **SimObject**\- The position is an offset from the simobject position (expressed in meters).<br>    <br>2. **SimObject Datum**\- The position is an offset from the simobject [Datum Reference Point](../../../../5_Content_Configuration/Modular_SimObjects/Aircraft/Aircraft.htm#DRP) (expressed in meters).<br>    <br>3. **Eyepoint** \- The position is an offset from the aircraft pilot eyepoint (expressed in meters).<br>4. **World** \- The position is expressed using latitude, longitude, and altitude. |
| `RotationReferentialObjectId` | Ignored if is `rotationReferential` set to World.<br>The object Id of the simobject that will be used to set rotation by SimObject, SimObject Datum and Eyepoint referential.<br>Set to 0 to focus user's aircraft. If the object Id is invalid, user's aircraft will be used instead. |
| `Fov` | The camera field of view, in radians. |

##### Remarks

This struct will be included in the `SIMCONNECT_RECV_CAMERA_DATA` struct as a response to to the `SimConnect_CameraGet` function.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)
4. [Camera](../Camera/Camera.htm)

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