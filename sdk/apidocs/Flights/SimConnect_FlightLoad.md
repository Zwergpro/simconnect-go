SimConnect\_FlightLoad

## SimConnect\_FlightLoad

The **SimConnect\_FlightLoad** function is used to load an existing flight file.

##### Syntax

```cpp
HRESULT SimConnect_FlightLoad(
    HANDLE  hSimConnect,
    const char*  szFileName
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _szFileName_ | Null-terminated string containing the path to the flight file. | String |

##### Return Values

This function does not return a value.

##### Remarks

Flight files can be opened using a text editor.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_FlightSave](SimConnect_FlightSave.htm)
4. [SimConnect\_FlightPlanLoad](SimConnect_FlightPlanLoad.htm)
5. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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