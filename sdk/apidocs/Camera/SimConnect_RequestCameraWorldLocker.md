SimConnect\_RequestCameraWorldLocker

## SimConnect\_RequestCameraWorldLocker

The **SimConnect\_RequestCameraWorldLocker** function is used to set a locker in the world to load the terrain, scenery and objects around it, but _only if you have acquired the camera first_. This is useful if you wish to move the camera far from the user aircraft where the world data may have been unloaded from memory.

**IMPORTANT!** World lockers are **performance heavy**, and as such should be used with care and **released as soon as possible**.

##### Syntax

```cpp
HRESULT SimConnect_RequestCameraWorldLocker(
    HANDLE  hSimConnect,
    SIMCONNECT_DATA_XYZ lockerPosition,
    SIMCONNECT_POSITION_REFERENTIAL referential,
    DWORD objectId
    );
```

##### Parameters

| Parameter | Description | Type |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `lockerPosition` | Coordinates of the locker, depending of the selected `referential` | `SIMCONNECT_DATA_XYZ` |
| `referential` | The reference used to define the `lockerPosition` value. One of the following:<br>1. **SimObject**\- The position is an offset from the simobject position (expressed in meters).<br>    <br>2. **SimObject Datum**\- The position is an offset from the simobject [Datum Reference Point](../../../../5_Content_Configuration/Modular_SimObjects/Aircraft/Aircraft.htm#DRP) (expressed in meters).<br>    <br>3. **Eyepoint** \- The position is an offset from the aircraft pilot eyepoint (expressed in meters).<br>4. **World** \- The position is expressed using latitude, longitude, and altitude. | `SIMCONNECT_POSITION_REFERENTIAL` |
| `objectId` | Ignored if is `referential` set to World.<br>The object Id of the simobject that will be used to set position by SimObject, SimObject Datum and Eyepoint referential.<br>Set to 0 to focus user's aircraft. If the object Id is invalid, user's aircraft will be used instead. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. If the camera locker failed to be setup, SimConnect will raise a `SIMCONNECT_EXCEPTION_CAMERA_API` exception to let you know that the operation failed (refer to the `SIMCONNECT_EXCEPTION` enum for more details). |

##### Example

```cpp
HANDLE  hSimConnect = NULL;
SIMCONNECT_DATA_XYZ LFPG_LLA{ 49.0083899664, 2.53844117956, 320 };
SimConnect_RequestCameraWorldLocker(hSimConnect, LFPG_LLA, SIMCONNECT_POSITION_REFERENTIAL_WORLD, 0);
```

##### Remarks

The sim will send a `SIMCONNECT_RECV_CAMERA_WORLD_LOCKER` to the client at the end of acquisition which you can use to check the camera has been properly acquired. It's handled through `SimConnect_CallDispatch`.

Related Topics

1. [Camera API](Camera_API.htm)
2. [SimConnect SDK](../../SimConnect_SDK.htm)
3. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
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