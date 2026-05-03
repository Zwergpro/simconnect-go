SimConnect\_SetClientData

## SimConnect\_SetClientData

The **SimConnect\_SetClientData** function is used to write one or more units of data to a client data area.

##### Syntax

```cpp
HRESULT SimConnect_SetClientData(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_DATA_ID  ClientDataID,
    SIMCONNECT_CLIENT_DATA_DEFINITION_ID  DefineID,
    DWORD  Flags,
    DWORD  dwReserved,
    DWORD  cbUnitSize,
    void*  pDataSet
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _ClientDataID_ | Specifies the ID of the client data area. | Integer |
| _DefineID_ | Specifies the ID of the client defined client data definition. | Integer |
| _Flags_ | Null, or one or more of the flags listed in the table below. | Flag |
| _dwReserved_ | Reserved for future use. Set to zero. | Integer |
| _cbUnitSize_ | Specifies the size of the data set in bytes. The server will check that this size matches exactly the size of the data definition provided in the DefineID parameter. An exception will be returned if this is not the case. | Integer |
| _pDataSet_ | Pointer to the data that is to be written. If the data is not in tagged format, this should point to the block of client data. If the data is in tagged format this should point to the first tag name ( _datumID_), which is always four bytes long, which should be followed by the data itself. Any number of tag name/value pairs can be specified this way, the server will use the _cbUnitSize_ parameter to determine how much data has been sent. | Integer |

The following table shows the different _flags_ that can be used:

| Flag | Description |
| --- | --- |
| NULL | Do nothing. |
| SIMCONNECT\_DATA\_SET\_FLAG\_TAGGED | The data to be set is being sent tagged format. If this flag is not set then the entire client data area should be replaced with new data. Refer to the _pDataSet_ parameter and [SimConnect\_RequestClientData](SimConnect_RequestClientData.htm) for more details on the tagged format. |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

A data definition must be specified, using the [SimConnect\_AddToClientDataDefinition](SimConnect_AddToClientDataDefinition.htm) function, before data can be set.

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SimConnect\_AddToClientDataDefinition](SimConnect_AddToClientDataDefinition.htm)
3. [SimConnect\_ClearClientDataDefinition](SimConnect_ClearClientDataDefinition.htm)
4. [SimConnect\_CreateClientData](SimConnect_CreateClientData.htm)
5. [SimConnect\_MapClientDataNameToID](SimConnect_MapClientDataNameToID.htm)
6. [SimConnect\_RequestClientData](SimConnect_RequestClientData.htm)

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