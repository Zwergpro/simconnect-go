SimConnect\_RequestClientData

## SimConnect\_RequestClientData

The **SimConnect\_RequestClientData** function is used to request that the specified data in an area created by another client be sent to this client.

##### Syntax

```cpp
HRESULT SimConnect_RequestClientData(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_DATA_ID  ClientDataID,
    SIMCONNECT_DATA_REQUEST_ID  RequestID,
    SIMCONNECT_CLIENT_DATA_DEFINITION_ID  DefineID,
    SIMCONNECT_CLIENT_DATA_PERIOD  Period = SIMCONNECT_CLIENT_DATA_PERIOD_ONCE,
    SIMCONNECT_CLIENT_DATA_REQUEST_FLAG  Flags = 0,
    DWORD  origin = 0,
    DWORD  interval = 0,
    DWORD  limit = 0
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _ClientDataID_ | Specifies the ID of the client data area. Before calling this function for the first time on one client area, call [SimConnect\_MapClientDataNameToID](SimConnect_MapClientDataNameToID.htm) to map an ID to the unique client data area name. This name must match the name specified by the client creating the data area with the [SimConnect\_MapClientDataNameToID](SimConnect_MapClientDataNameToID.htm) and [SimConnect\_CreateClientData](SimConnect_CreateClientData.htm) functions. | Integer |
| _RequestID_ | Specifies the ID of the client-defined request. This is used later by the client to identify which data has been received. This value should be unique for each request, re-using a _RequestID_ will overwrite any previous request using the same ID. | Integer |
| _DefineID_ | Specifies the ID of the client-defined data definition. This definition specifies the data that should be sent to the client. | Integer |
| _Period_ | One member of the [SIMCONNECT\_CLIENT\_DATA\_PERIOD](../Structures_And_Enumerations/SIMCONNECT_CLIENT_DATA_PERIOD.htm) enumeration type, specifying how often the data is to be sent by the server and received by the client. | Integer<br>(OPTIONAL) |
| _Flags_ | A DWORD containing one or more of the values from the table below. | Integer<br>(OPTIONAL) |
| _origin_ | The number of _Period_ events that should elapse before transmission of the data begins. The default is zero, which means transmissions will start immediately. | Integer<br>(OPTIONAL) |
| _interval_ | The number of _Period_ events that should elapse between transmissions of the data. The default is zero, which means the data is transmitted every _Period_. | Integer<br>(OPTIONAL) |
| _limit_ | The number of times the data should be transmitted before this communication is ended. The default is zero, which means the data should be transmitted endlessly. | Integer<br>(OPTIONAL) |

The table below lists the available values that can be used with the _Flags_ parameter:

| Flag Value | Description |
| --- | --- |
| 0 | The default, data will be sent strictly according to the defined period. |
| SIMCONNECT\_CLIENT\_DATA\_REQUEST\_FLAG\_CHANGED | Data will only be sent to the client when one or more values have changed. If this is the only flag set, then all the variables in a data definition will be returned if just one of the values changes. |
| SIMCONNECT\_CLIENT\_DATA\_REQUEST\_FLAG\_TAGGED | Requested data will be sent in tagged format (datum ID/value pairs). Tagged format requires that a datum reference ID is returned along with the data value, in order that the client code is able to identify the variable. This flag is usually set in conjunction with the previous flag, but it can be used on its own to return all the values in a data definition in datum ID/value pairs. See the [SIMCONNECT\_RECV\_CLIENT\_DATA](../Structures_And_Enumerations/SIMCONNECT_RECV_CLIENT_DATA.htm) structure for more details. |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

A data definition must be specified, using the [SimConnect\_AddToClientDataDefinition](SimConnect_AddToClientDataDefinition.htm) function, before this function can be called If the data definition exceeds the size of the client data area on the server, then the extra bytes will be filled with zeros, an error will not be returned.

The data will be returned in a [SIMCONNECT\_RECV\_CLIENT\_DATA](../Structures_And_Enumerations/SIMCONNECT_RECV_CLIENT_DATA.htm) structure.

See the remarks for [SimConnect\_RequestDataOnSimObject](SimConnect_RequestDataOnSimObject.htm), as the two functions work in a very similar manner.

This function has been updated for the SP1a release of the SDK, expanding on its functionality.

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SimConnect\_AddToClientDataDefinition](SimConnect_AddToClientDataDefinition.htm)
3. [SimConnect\_ClearClientDataDefinition](SimConnect_ClearClientDataDefinition.htm)
4. [SimConnect\_CreateClientData](SimConnect_CreateClientData.htm)
5. [SimConnect\_MapClientDataNameToID](SimConnect_MapClientDataNameToID.htm)
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