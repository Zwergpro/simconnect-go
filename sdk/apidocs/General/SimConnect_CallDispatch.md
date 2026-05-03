SimConnect\_CallDispatch

## SimConnect\_CallDispatch

The **SimConnect\_CallDispatch** function is used to process the next SimConnect message received, through the specified callback function.

##### Syntax

```cpp
HRESULT SimConnect_CallDispatch(
    HANDLE  hSimConnect,
    DispatchProc  pfcnDispatch,
    void *  pContext
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _pfcnDispatch_ | Specifies the callback function. For a definition of the function see `DispatchProc`. | Integer |
| _pContext_ | Specifies a pointer that the client can define that will be returned in the callback. This is used in particular by managed code clients to pass a _this_ pointer to the callback. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
int quit = 0;
while( quit == 0 ) {
    hr = SimConnect_CallDispatch(hSimConnect, MyDispatchProc, NULL);
    };
```

##### Remarks

It is important to call this function sufficiently frequently that the queue of information received from the server is processed (typically it is coded within a **while** loop that terminates when the application is exited).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_GetNextDispatch](SimConnect_GetNextDispatch.htm)
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