DispatchProc

## DispatchProc

The **DispatchProc** function is written by the developer of the SimConnect client, as a callback function to handle all the communications with the server.

##### Syntax

```cpp
void CALLBACK DispatchProc(
    SIMCONNECT_RECV*  pData,
    DWORD  cbData,
    void *  pContext
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _pData_ | Pointer to a data buffer, to be treated initially as a [SIMCONNECT\_RECV](../Structures_And_Enumerations/SIMCONNECT_RECV.htm) structure. If you are going to make a copy of the data buffer (which is maintained by the SimConnect client library) make sure that the defined buffer is large enough (the size of the returned data structure is one member of the [SIMCONNECT\_RECV](../Structures_And_Enumerations/SIMCONNECT_RECV.htm) structure. | Integer |
| _cbData_ | The size of the data buffer, in bytes. | Integer |
| _pContext_ | Contains the pointer specified by the client in the [SimConnect\_CallDispatch](SimConnect_CallDispatch.htm) function call. | Integer |

##### Return Values

This function does not return a value.

##### Example

```cpp
void CALLBACK MyDispatchProc(SIMCONNECT_RECV* pData, DWORD cbData)
{
    case SIMCONNECT_RECV_ID_OPEN:
        SIMCONNECT_RECV_OPEN *openData = (SIMCONNECT_RECV_OPEN*) pData;
        break;
    case SIMCONNECT_RECV_ID_EVENT:
        SIMCONNECT_RECV_EVENT *evt =   (SIMCONNECT_RECV_EVENT*) pData;
        break;
    case SIMCONNECT_RECV_ID_EVENT_FILENAME:
        SIMCONNECT_RECV_EVENT_FILENAME *evt =   (SIMCONNECT_RECV_EVENT_FILENAME)*) pData;
        break;
    case SIMCONNECT_RECV_ID_EVENT_OBJECT_ADDREMOVE:
        SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE *evt =   (SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE*) pData;
        break;
    case   SIMCONNECT_RECV_ID_EVENT_FRAME:
        SIMCONNECT_RECV_EVENT_FRAME *evt =   (SIMCONNECT_RECV_EVENT_FRAME*) pData;
        break;
    case   SIMCONNECT_RECV_ID_SIMOBJECT_DATA:
        SIMCONNECT_RECV_SIMOBJECT_DATA   *pObjData = (SIMCONNECT_RECV_SIMOBJECT_DATA*) pData;
        break;
    case   SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE:
        SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE   *pObjData = (SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE*) pData;
        break;
    case   SIMCONNECT_RECV_ID_QUIT:
    case   SIMCONNECT_RECV_ID_EXCEPTION:
        SIMCONNECT_RECV_EXCEPTION   *except =   (SIMCONNECT_RECV_EXCEPTION*) pData;
        break;
    case   SIMCONNECT_RECV_ID_WEATHER_OBSERVATION:
        SIMCONNECT_RECV_WEATHER_OBSERVATION* pWxData =   (SIMCONNECT_RECV_WEATHER_OBSERVATION*) pData;
        const char* pszMETAR = (const     char*) (pWxData+1);
        break;
    default:
        break;
}
```

##### Remarks

This function can be named appropriately by the client developer. The name of the function is passed to the client-side library with the [SimConnect\_CallDispatch](SimConnect_CallDispatch.htm) function call. Handle all the callback events in this function. If you do not wish to implement a callback function use [SimConnect\_GetNextDispatch](SimConnect_GetNextDispatch.htm).

To receive time based notifications, see the [SimConnect\_SubscribeToSystemEvent](SimConnect_SubscribeToSystemEvent.htm) function. To receive event based notifications see the [SimConnect\_AddClientEventToNotificationGroup](../Events_And_Data/SimConnect_AddClientEventToNotificationGroup.htm) function. To send an event to be received by other clients, see the [SimConnect\_TransmitClientEvent](../Events_And_Data/SimConnect_TransmitClientEvent.htm) function.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_CallDispatch](SimConnect_CallDispatch.htm)
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