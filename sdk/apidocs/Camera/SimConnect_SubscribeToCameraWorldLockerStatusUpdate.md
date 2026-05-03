SimConnect\_SubscribeToCameraWorldLockerStatusUpdate

## SimConnect\_SubscribeToCameraWorldLockerStatusUpdate

The **SimConnect\_SubscribeToCameraWorldLockerStatusUpdate** function is used to receive a `SIMCONNECT_RECV_CAMERA_WORLD_LOCKER_STATUS` message whenever the simulation updates the camera status.

##### Syntax

```cpp
HRESULT SimConnect_SubscribeToCameraWorldLockerStatusUpdate(
    HANDLE hSimConnect
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
SimConnect_SubscribeToCameraStatusUpdate(hSimConnect);
```

##### Remarks

N/A

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