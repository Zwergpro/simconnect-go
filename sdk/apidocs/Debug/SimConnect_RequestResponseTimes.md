SimConnect\_RequestResponseTimes

## SimConnect\_RequestResponseTimes

The **SimConnect\_RequestResponseTimes** function is used to provide some data on the performance of the client-server connection.

##### Syntax

```cpp
HRESULT SimConnect_RequestResponseTimes(
    HANDLE  hSimConnect,
    DWORD  nCount,
    float*  fElapsedSeconds
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _nCount_ | Integer containing the number of elements in the array of floats. This should be set to five for the full range of timings, but can be less if only the first few are of interest. There is no point creating an array of greater than five floats. | Integer |
| _fElapsedSeconds_ | An array of _nCoun_ t floats, containing the times. The five elements will contain the following:<br>**0** \- total round trip time<br>**1** \- time from the request till the packet is sent<br>**2** \- time from the request till the packet is received by the server<br>**3** \- time from the request till the response is made by the server<br>**4** \- time from the server response to the client receives the packet. | Array |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
int quit = 0;
bool fTesting = true;
      ....
while( quit == 0 )
    {
    hr = SimConnect_CallDispatch(hSimConnect, MyDispatchProc, NULL);
    Sleep(0);
    if (fTesting)
        {
        fTesting = false;
        float fElapsedSeconds[5];
        hr = SimConnect_RequestResponseTimes(hSimConnect, 5, &fElapsedSeconds[0]);
          ....
        }
    }
```

##### Remarks

This function should not be used as part of a final application, as it is costly in performance, but is available to help provide some performance data that can be used while building an testing a client application.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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