SimConnect\_CameraGet

## SimConnect\_CameraGet

The **SimConnect\_CameraGet** function is used to get the current add-on camera properties, regardless of whether it has been acquired or not.

##### Syntax

```cpp
HRESULT SimConnect_CameraGet(
    HANDLE hSimConnect,
    DWORD Referential
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _Referential_ | The referential to use for the camera position values (including the camera "target" position values, if used), can be one of the following `SIMCONNECT_POSITION_REFERENTIAL` enum members:<br>1. `SIMCONNECT_POSITION_REFERENTIAL_AIRCRAFT` \- The position is an offset from the aircraft [Datum Reference Point](../../../../5_Content_Configuration/Modular_SimObjects/Aircraft/Aircraft.htm#DRP) (expressed in meters).<br>2. `SIMCONNECT_POSITION_REFERENTIAL_WORLD` \- The position is expressed using latitude, longitude, and altitude.<br>3. `SIMCONNECT_POSITION_REFERENTIAL_EYEPOINT` \- The position is an offset from the aircraft pilot eyepoint (expressed in meters). | enum |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

Request camera data:

```cpp
SimConnect_CameraGet(hSimConnect, SIMCONNECT_POSITION_REFERENTIAL_AIRCRAFT); // Request CameraGet
while (!quit)
{
    SimConnect_CallDispatch(hSimConnect, MyDispatchProc, NULL);
    Sleep(1);
}
```

The user defined dispatch function:

```cpp
void CALLBACK MyDispatchProc(SIMCONNECT_RECV* pData, DWORD cbData, void* pContext)
{
    switch (pData->dwID)
    {
        case SIMCONNECT_RECV_ID_CAMERA_DATA:
        {
            // Receive the coordinate of the camera
            SIMCONNECT_RECV_CAMERA_DATA* pCam = (SIMCONNECT_RECV_CAMERA_DATA*)pData;
            printf("\n<-Get: %f %f %f %f %f %f %f", pCam->CameraData.Pos.x, pCam->CameraData.Pos.y, pCam->CameraData.Pos.z, pCam->CameraData.Pbh.Pitch, pCam->CameraData.Pbh.Bank, pCam->CameraData.Pbh.Heading, pCam->CameraData.Fov);
            break;
        }
    }
}
```

##### Remarks

This function will generate a `SIMCONNECT_DATA_CAMERA` struct which will contain all the information about the current camera, unless the camera has not been acquired, in which case a `SIMCONNECT_EXCEPTION_CAMERA_API` exception will be sent (see `SIMCONNECT_EXCEPTION` for more information).

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