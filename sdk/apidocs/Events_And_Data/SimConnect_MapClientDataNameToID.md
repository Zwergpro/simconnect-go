SimConnect\_MapClientDataNameToID

## SimConnect\_MapClientDataNameToID

The **SimConnect\_MapClientDataNameToID** function is used to associate an ID with a named client data area.

##### Syntax

```cpp
HRESULT SimConnect_MapClientDataNameToID(
      HANDLE  hSimConnect,
      const char*  szClientDataName,
      SIMCONNECT_CLIENT_DATA_ID  ClientDataID
      );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _szClientDataName_ | Null-terminated string containing the client data area name. This is the name that another client will use to specify the data area. The name is not case-sensitive. If the name requested is already in use by another addon, a error will be returned. | Integer |
| _ClientDataID_ | A unique ID for the client data area, specified by the client. If the ID number is already in use, an error will be returned. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |
| SIMCONNECT\_EXCEPTION\_ALREADY\_CREATED | The name requested is already in use by another addon. |
| SIMCONNECT\_EXCEPTION\_DUPLICATE\_ID | The ID number is already in use. |

##### Remarks

This function should be called once for each client data area: the client setting up the data should call it just before a call to [SimConnect\_CreateClientData](SimConnect_CreateClientData.htm), and the clients requesting the data should call it before any calls to [SimConnect\_RequestClientData](SimConnect_RequestClientData.htm) are made. The name given to a client data area must be unique, however by mapping an ID number to the name, calls to the functions to set and request the data are made more efficient.

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SimConnect\_AddToClientDataDefinition](SimConnect_AddToClientDataDefinition.htm)
3. [SimConnect\_ClearClientDataDefinition](SimConnect_ClearClientDataDefinition.htm)
4. [SimConnect\_CreateClientData](SimConnect_CreateClientData.htm)
5. [SimConnect\_RequestClientData](SimConnect_RequestClientData.htm)
6. [SimConnect\_SetClientData](SimConnect_SetClientData.htm)

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