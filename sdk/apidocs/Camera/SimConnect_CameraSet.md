SimConnect\_CameraSet

## SimConnect\_CameraSet

The **SimConnect\_CameraSet** function is used to set the properties of the camera. This function can _only_ be used when the add-on camera has been **acquired**.

##### Syntax

```cpp
HRESULT SimConnect_CameraSet(
    HANDLE hSimConnect,
    SIMCONNECT_DATA_CAMERA CameraData,
    DWORD DataMask
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _CameraData_ | This is a `SIMCONNECT_DATA_CAMERA` struct containing the member parameters you wish to set. The parameters should correspond to the `DataMask` used. | Struct |
| _DataMask_ | This is a bitmask that can be created using the `SIMCONNECT_CAMERA_DATA_MASK` enum members, and defines which parts of the `CameraData` struct should be used for the camera. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

Set camera data:

```cpp
SIMCONNECT_DATA_CAMERA data;
data.PositionReferential = SIMCONNECT_POSITION_REFERENTIAL_SIMOBJECT;
data.PositionReferentialObjectId = 0;
data.Position.z =-10.0;
data.Position.x = 0;
data.Position.y = 4.0;
data.RotationReferential = SIMCONNECT_POSITION_REFERENTIAL_SIMOBJECT;
data.RotationReferentialObjectId = 0;
data.Pbh.Pitch = 0;
data.Pbh.Bank = 0;
data.Pbh.Heading = 0;
data.Fov = 0.80;
SimConnect_CameraSet(hSimConnect, data, SIMCONNECT_CAMERA_DATA_MASK_ALL_ROTATION);
```

##### Remarks

Note that if the add-on camera has not been acquired previously, then a `SIMCONNECT_EXCEPTION_CAMERA_API` exception will be sent (see `SIMCONNECT_EXCEPTION` for more information).

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