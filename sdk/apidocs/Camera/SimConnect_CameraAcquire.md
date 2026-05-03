SimConnect\_CameraAcquire

## SimConnect\_CameraAcquire

The **SimConnect\_CameraAcquire** function is used to acquire the add-on camera. Note that the camera may not be immediately useable on acquisition, as it may be under simulation control. Please see the main [Camera API](Camera_API.htm) page for more information.

##### Syntax

```cpp
HRESULT SimConnect_CameraAcquire(
    HANDLE  hSimConnect
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. If the camera failed to be acquired, SimConnect will raise a `SIMCONNECT_EXCEPTION_CAMERA_API` exception to let you know that the operation failed (refer to the `SIMCONNECT_EXCEPTION` enum for more details). |

##### Example

```cpp
HANDLE  hSimConnect = NULL;
SimConnect_CameraAcquire(hSimConnect)
```

##### Remarks

The sim will send a [SIMCONNECT\_RECV\_CAMERA\_STATUS](../Structures_And_Enumerations/SIMCONNECT_RECV_CAMERA_STATUS.htm) to the client at the end of acquisition which you can use to check the camera has been properly acquired. It's handled through [SimConnect\_CallDispatch](../General/SimConnect_CallDispatch.htm).

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