SimConnect\_FlightSave

## SimConnect\_FlightSave

The **SimConnect\_FlightSave** function is used to save the current state of a flight to a flight file.

##### Syntax

```cpp
HRESULT SimConnect_FlightSave(
    HANDLE  hSimConnect,
    const char*  szFileName,
    const char*  szTitle
    const char*  szDescription,
    DWORD  Flags
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _szFileName_ | Null-terminated string containing the path to the flight file. Flight files have the extension .FLT, but no need to enter an extension here. | String |
| _szTitle_ | Null-terminated string containing the title of the flight file. If this is NULL then the szFileName parameter is used as the title. | String |
| _szDescription_ | Null-terminated string containing the text to enter in the Description field of the flight file. | String |
| _Flags_ | Unused. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

Flight files can be opened using a text editor.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_FlightLoad](SimConnect_FlightLoad.htm)
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