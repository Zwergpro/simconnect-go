SimConnect\_CameraRelease

## SimConnect\_CameraRelease

The **SimConnect\_CameraRelease** function is used to release a previously acquired camera when no longer required or for other systems to acquire it.

##### Syntax

```cpp
HRESULT SimConnect_CameraRelease(
    HANDLE hSimConnect,
    const char * CameraDefName
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _CameraDefName_ | The name of the defined camera to which the simulation should switch once the add-on camera is released. If not specified, the simulation will go back to the camera being used before the add-on camera was acquired. | String |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
SimConnect_CameraRelease(hSimConnect, "Pilot"); // Will switch to camera Pilot
SimConnect_CameraRelease(hSimConnect, ""); // Will switch to last selected camera
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