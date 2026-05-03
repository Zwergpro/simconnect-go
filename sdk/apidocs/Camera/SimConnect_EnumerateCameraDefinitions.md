SimConnect\_EnumerateCameraDefinitions

## SimConnect\_EnumerateCameraDefinitions

The **SimConnect\_EnumerateCameraDefinitions** function is used to retrieve an array of all the defined camera names.

##### Syntax

```cpp
HRESULT SimConnect_EnumerateCameraDefinitions(
    HANDLE hSimConnect,
    SIMCONNECT_DATA_REQUEST_ID RequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _RequestID_ | Specifies the client defined request ID. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

Request information on cameras:

```cpp
HANDLE  hSimConnect = NULL;
int expectedDefReceived = -1;
int defReceived = 0;
void main()
{
    // Open SimConnect connection
    if (SUCCEEDED(SimConnect_Open(&hSimConnect, "Enumerate Camera Def Names", NULL, 0, 0, 0)))
    {
        SimConnect_EnumerateCameraDefinitions(hSimConnect, 0);
        // Call dispatch proc while we didn't receive all
        while (expectedDefReceived != defReceived)
        {
            SimConnect_CallDispatch(hSimConnect, MyDispatchProc, NULL);
            Sleep(1);
        }
    }
}
```

User defined dispatch function:

```cpp
void CALLBACK MyDispatchProc(SIMCONNECT_RECV* pData, DWORD cbData, void* pContext)
{
    switch (pData->dwID)
    {
        case SIMCONNECT_RECV_ID_CAMERA_DEFINITION_LIST:
        {
            SIMCONNECT_RECV_CAMERA_DEFINITION_LIST* list = (SIMCONNECT_RECV_CAMERA_DEFINITION_LIST*)pData;
            for (int i = 0; i < list->dwArraySize; ++i)
            {
                std::cout << list->rgData[i].Str << std::endl;
            }
            expectedDefReceived = list->dwOutOf;
            defReceived++;
            break;
        }
    }
}
```

##### Remarks

This function will generate a `SIMCONNECT_RECV_CAMERA_DEFINITION_LIST` containing an array with all the different unique camera ID strings.

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