SimConnect\_FlightPlanLoad

## SimConnect\_FlightPlanLoad

The **SimConnect\_FlightPlanLoad** function is used to load an existing flight plan file.

**NOTE**: The current status of this function is NO ERROR, NO RESPONSE.

##### Syntax

```cpp
HRESULT SimConnect_FlightPlanLoad(
    HANDLE  hSimConnect,
    const char*  szFileName
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _szFileName_ | Null-terminated string containing the path to the flight plan file. Flight plans have the extension .PLN, but no need to enter an extension here. The easiest way to create flight plans is to create them from within Microsoft Flight Simulator itself, and then save them off for use by the user or AI controlled aircraft. | String |

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
4. [SimConnect\_FlightSave](SimConnect_FlightSave.htm)
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