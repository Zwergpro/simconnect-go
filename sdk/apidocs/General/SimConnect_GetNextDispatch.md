SimConnect\_GetNextDispatch

## SimConnect\_GetNextDispatch

The **SimConnect\_GetNextDispatch** function is used to process the next SimConnect message received, without the use of a callback function.

##### Syntax

```cpp
HRESULT SimConnect_GetNextDispatch(
    HANDLE  hSimConnect,
    SIMCONNECT_RECV**  ppData,
    DWORD*  pcbData
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _ppData_ | Pointer to a pointer to a data buffer, initially to be treated as a [SIMCONNECT\_RECV](../Structures_And_Enumerations/SIMCONNECT_RECV.htm) structure. If you are going to make a copy of the data buffer (which is maintained by the SimConnect client library) make sure that the defined buffer is large enough (the size of the returned data structure is one member of the [SIMCONNECT\_RECV](../Structures_And_Enumerations/SIMCONNECT_RECV.htm) structure. | Integer |
| _pcbData_ | Pointer to the size of the data buffer, in bytes. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
SIMCONNECT_RECV* pData;
DWORD cbData;

hr = SimConnect_GetNextDispatch(hSimConnect, &pData, &cbData);
if (SUCCEEDED(hr))
{
    switch(pData->dwID)
        {
        case SIMCONNECT_RECV_ID_OPEN:
            SIMCONNECT_RECV_OPEN *openData = (SIMCONNECT_RECV_OPEN*) pData;
            break;
        case SIMCONNECT_RECV_ID_EVENT:
            SIMCONNECT_RECV_EVENT *evt = (SIMCONNECT_RECV_EVENT*) pData;
            break;
        case SIMCONNECT_RECV_ID_EVENT_FILENAME:
            SIMCONNECT_RECV_EVENT_FILENAME *evt = (SIMCONNECT_RECV_EVENT_FILENAME*) pData;
            break;
        case SIMCONNECT_RECV_ID_EVENT_OBJECT_ADDREMOVE:
            SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE *evt = (SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE*) pData;
            break;
        case SIMCONNECT_RECV_ID_EVENT_FRAME:
            SIMCONNECT_RECV_EVENT_FRAME *evt = (SIMCONNECT_RECV_EVENT_FRAME*) pData;
            break;
        case SIMCONNECT_RECV_ID_SIMOBJECT_DATA:
            SIMCONNECT_RECV_SIMOBJECT_DATA *pObjData = (SIMCONNECT_RECV_SIMOBJECT_DATA*) pData;
            break;
        case SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE:
            SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE *pObjData = (SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE*) pData;
            break;
        case SIMCONNECT_RECV_ID_QUIT:
            break;
        case SIMCONNECT_RECV_ID_EXCEPTION:
            SIMCONNECT_RECV_EXCEPTION *except = (SIMCONNECT_RECV_EXCEPTION*) pData;
            break;
        case SIMCONNECT_RECV_ID_WEATHER_OBSERVATION:
            SIMCONNECT_RECV_WEATHER_OBSERVATION* pWxData = (SIMCONNECT_RECV_WEATHER_OBSERVATION*) pData;
            const char* pszMETAR = (const char*) (pWxData+1);
            break;
        default:
            break;
        }
    }
```

##### Remarks

It is important to call this function sufficiently frequently that the queue of information received from the server is processed. If there are no messages in the queue, the **\[dwID\]** parameter will be set to `SIMCONNECT_RECV_ID_NULL`.

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