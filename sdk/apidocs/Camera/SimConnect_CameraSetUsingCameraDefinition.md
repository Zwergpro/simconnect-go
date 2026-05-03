SimConnect\_CameraSetUsingCameraDefinition

## SimConnect\_CameraSetUsingCameraDefinition

The **SimConnect\_CameraSetUsingCameraDefinition** function is used to set specific parts of the camera struct based on a predefined camera definition from the cameras.cfg file. This function can only be used when the add-on camera has been acquired.

##### Syntax

```cpp
HRESULT SimConnect_CameraSetUsingCameraDefinition(
    HANDLE hSimConnect,
    HANDLE hSimConnect, const char* cameraDefinitionName
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _cameraDefinitionName_ | This is the [Title](../../../../5_Content_Configuration/CFG_Files/cameras_cfg.htm#Title) string of a camera definition to use. | String |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

N/A

##### Remarks

The data that will be used from the given camera definition is as follows:

- **Position** \- From the `InitialXyz` parameter
- **Pbh** \- From the `InitialPbh` parameter
- **FOV** \- From the `Initialzoom` parameter

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