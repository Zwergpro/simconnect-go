SimConnect\_Close

## SimConnect\_Close

The **SimConnect\_Close** function is used to request that the communication with the server is ended.

##### Syntax

```cpp
HRESULT SimConnect_Close(
    HANDLE  hSimConnect
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. This should only happen if a the `hSimConnect` parameter is erroneous. |

##### Example

```cpp
hr = SimConnect_Close(hSimConnect);
```

##### Remarks

When a SimConnect client is closed, the server will remove all objects, menu items, group definitions and so on, defined or requested by that client, so there is no need to remove them explicitly in the client code.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_Open](SimConnect_Open.htm)
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