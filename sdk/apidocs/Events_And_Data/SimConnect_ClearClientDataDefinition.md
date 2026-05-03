SimConnect\_ClearClientDataDefinition

## SimConnect\_ClearClientDataDefinition

The **SimConnect\_ClearClientDataDefinition** function is used to clear the definition of the specified client data.

##### Syntax

```cpp
HRESULT SimConnect_ClearClientDataDefinition(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_DATA_DEFINITION_ID  DefineID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _DefineID_ | Specifies the ID of the client defined client data definition. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

N/A

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SimConnect\_AddToClientDataDefinition](SimConnect_AddToClientDataDefinition.htm)
3. [SimConnect\_CreateClientData](SimConnect_CreateClientData.htm)
4. [SimConnect\_MapClientDataNameToID](SimConnect_MapClientDataNameToID.htm)
5. [SimConnect\_SetClientData](SimConnect_SetClientData.htm)
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