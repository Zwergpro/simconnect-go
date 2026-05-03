SimConnect\_CameraGetStatus

## SimConnect\_CameraGetStatus

The **SimConnect\_CameraGetStatus** function is used to get the add-on camera status, regardless of whether it has been acquired or not.

##### Syntax

```cpp
HRESULT SimConnect_CameraGetStatus(
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

Request the camera status:

```cpp
SimConnect_CameraGetStatus(hSimConnect);
```

User defined dispatch function:

```cpp
void CALLBACK MyDispatchProc(SIMCONNECT_RECV* pData, DWORD cbData, void* pContext)
{
    switch (pData->dwID)
    {
        case SIMCONNECT_RECV_ID_CAMERA_STATUS:
            {
                // Receive an update on the camera status (the structure is subject to change)
                SIMCONNECT_RECV_CAMERA_STATUS* pRes = (SIMCONNECT_RECV_CAMERA_STATUS*)pData;

                // Is the game controlling the camera ? (RTC, menu, ...)
                bIsGameControlled = pRes->bGameControlled;

                // Is the camera acquired by the current user
                bIsAquired = pRes->acquiredState == SIMCONNECT_CAMERA_ACQUIRED;

                printf("\nIsSettable: %s", (bIsAquired && !bIsGameControlled) ? "true" : "false");
                printf("\nIsAcquired: %s - IsGameControlled: %s", bIsAquired ? "true" : "false", bIsGameControlled ? "true" : "false");

                break;
            }
    }
}
```

##### Remarks

The simulation will send a `SIMCONNECT_RECV_CAMERA_STATUS` to the client requesting the camera status.

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