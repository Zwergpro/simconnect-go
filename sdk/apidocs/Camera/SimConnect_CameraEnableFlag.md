SimConnect\_CameraEnableFlag

## SimConnect\_CameraEnableFlag

The **SimConnect\_CameraEnableFlag** function can be used to enable camera specific features, but only if the camera has been correctly acquired.

##### Syntax

```cpp
SimConnect_CameraEnableFlag(
    HANDLE hSimConnect,
    DWORD Flag
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _Flag_ | This is one or more of the `SIMCONNECT_CAMERA_FLAG` enum members, merged using bitwise "or" `|` to create a bitmask which flags specific camera options as enabled. | Bitmask |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
SimConnect_CameraEnableFlag(hSimConnect, SIMCONNECT_CAMERA_FLAG_INTERACTION);
```

##### Remarks

This function

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