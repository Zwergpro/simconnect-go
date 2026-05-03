SimConnect\_RequestDataOnSimObjectType

## SimConnect\_RequestDataOnSimObjectType

The **SimConnect\_RequestDataOnSimObjectType** function is used to retrieve information about simulation objects of a given type that are within a specified radius of the user's aircraft.

##### Syntax

```cpp
HRESULT SimConnect_RequestDataOnSimObjectType(
    HANDLE  hSimConnect,
    SIMCONNECT_DATA_REQUEST_ID  RequestID,
    SIMCONNECT_DATA_DEFINITION_ID  DefineID,
    DWORD  dwRadiusMeters,
    SIMCONNECT_SIMOBJECT_TYPE  type
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _RequestID_ | Specifies the ID of the client defined group. | Integer |
| _DefineID_ | Specifies the ID of the client defined [data definition](SimConnect_AddToDataDefinition.htm). | Integer |
| _dwRadiusMeters_ | Double word containing the radius in meters. If this is set to zero only information on the user aircraft will be returned, although this value is ignored if type is set to [`SIMCONNECT_SIMOBJECT_TYPE_USER`](../Structures_And_Enumerations/SIMCONNECT_SIMOBJECT_TYPE.htm). The error [`SIMCONNECT_EXCEPTION_OUT_OF_BOUNDS`](../Structures_And_Enumerations/SIMCONNECT_EXCEPTION.htm) will be returned if a radius is given and it exceeds the maximum allowed (200,000 meters, or 200 Km). |  |
| _type_ | Specifies the type of object to receive information on. One member of the `SIMCONNECT_SIMOBJECT_TYPE` enumeration type |  |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
void CALLBACK dispatchEvents(SIMCONNECT_RECV*  pData, DWORD  cbData, void *  pContext)
{
HRESULT hr;
switch (pData->dwID)
    {
    case SIMCONNECT_RECV_ID_EVENT:
        {
        SIMCONNECT_RECV_EVENT *evt = (SIMCONNECT_RECV_EVENT*)pData;
        switch (evt->uEventID)
            {
            case STARTUP:
                stats.lastDeleted = 0;
                stats.nrOfDeletions = 0;
                stats.nrOfRequests = 0;
                // Now the sim is running, request information on the user aircraft
                hr = SimConnect_RequestDataOnSimObjectType(hSimConnect, R1, AIPARKDATA, conf.radius, SIMCONNECT_SIMOBJECT_TYPE_AIRCRAFT);
                ++stats.nrOfRequests;
                break;
            case R2:
                fsecCnt++;
                if (fsecCnt >= (conf.requestEveryXSec/4))
                    {
                    fsecCnt = 0;
                    if (SUCCEEDED(SimConnect_RequestDataOnSimObjectType(hSimConnect, R1, AIPARKDATA, conf.radius, SIMCONNECT_SIMOBJECT_TYPE_AIRCRAFT)))
                        {
                        ++stats.nrOfRequests;
                        }
                    }
                break;
            }
        break;
    // further cases here
    default:
        printf("\nReceived:%d", pData->dwID);
        break;
    }
}
```

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AddToDataDefinition](SimConnect_AddToDataDefinition.htm)
4. [SimConnect\_ClearDataDefinition](SimConnect_ClearDataDefinition.htm)
5. [SimConnect\_RequestDataOnSimObject](SimConnect_RequestDataOnSimObject.htm)
6. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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