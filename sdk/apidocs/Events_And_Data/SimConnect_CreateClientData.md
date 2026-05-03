SimConnect\_CreateClientData

## SimConnect\_CreateClientData

The **SimConnect\_CreateClientData** function is used to request the creation of a reserved data area for this client.

##### Syntax

```cpp
HRESULT SimConnect_CreateClientData(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_DATA_ID  ClientDataID,
    DWORD  dwSize,
    SIMCONNECT_CREATE_CLIENT_DATA_FLAG  Flags
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _ClientDataID_ | ID of the client data area. Before calling this function, call [SimConnect\_MapClientDataNameToID](SimConnect_MapClientDataNameToID.htm) to map an ID to a unique client area name. | Integer |
| _dwSize_ | Double word containing the size of the data area in bytes. | Integer |
| _Flags_ | Specify the flag `SIMCONNECT_CREATE_CLIENT_DATA_FLAG_READ_ONLY` if the data area can only be written to by this client (the client creating the data area). By default other clients can write to this data area. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

Use this function, along with the other client data functions, to reserve an area of memory for client data on the server, that other clients can have read (or read and write) access to. Specify the contents of the data area with the [SimConnect\_AddToClientDataDefinition](SimConnect_AddToClientDataDefinition.htm) call, and set the actual values with a call to [SimConnect\_SetClientData](SimConnect_SetClientData.htm). Other clients can receive the data with a call to [SimConnect\_RequestClientData](SimConnect_RequestClientData.htm).

The maximum size of a client data area is set by the constant `SIMCONNECT_CLIENTDATA_MAX_SIZE`, which is 8K.There is no maximum number of client data areas, but the total must not exceed 1Mbyte. If a request is made for a client data area greater than `SIMCONNECT_CLIENTDATA_MAX_SIZE` a `SIMCONNECT_EXCEPTION_INVALID_DATA_SIZE` exception is returned. If a request is made for a client data area that will exceed the total maximum memory a `SIMCONNECT_EXCEPTION_OUT_OF_BOUNDS` exception is returned.

One client area can be referenced by any number of client data definitions. Typically the name of the client area, and the data definitions, should be published appropriately so other clients can be written to use them. Care should be taken to give the area a unique name.

Once created, a client data area cannot be deleted or reduced in size. To increase the size of the data area, first close the connection, then re-open it and request the client data area again, using the same name, but with the required size. The additional data area will be initialized to zero, but the previous data will be untouched by this process. Client data persists to the end of the _Microsoft Flight Simulator 2024_ session, and is not removed when the client that created it is closed. It is also possible to change a read-only data area to write-able using this technique.

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SimConnect\_AddToClientDataDefinition](SimConnect_AddToClientDataDefinition.htm)
3. [SimConnect\_ClearClientDataDefinition](SimConnect_ClearClientDataDefinition.htm)
4. [SimConnect\_MapClientDataNameToID](SimConnect_MapClientDataNameToID.htm)
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