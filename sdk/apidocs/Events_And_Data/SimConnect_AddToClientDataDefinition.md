SimConnect\_AddToClientDataDefinition

## SimConnect\_AddToClientDataDefinition

The **SimConnect\_AddToClientDataDefinition** function is used to add an offset and a size in bytes, or a type, to a client data definition.

##### Syntax

```cpp
HRESULT SimConnect_AddToClientDataDefinition(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_DATA_DEFINITION_ID  DefineID,
    DWORD  dwOffset,
    DWORD  dwSizeOrType,
    float  fEpsilon = 0,
    DWORD  DatumID = SIMCONNECT_UNUSED
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _DefineID_ | Specifies the ID of the client-defined client data definition. | Integer |
| _dwOffset_ | Double word containing the offset into the client area, where the new addition is to start. Set this to `SIMCONNECT_CLIENTDATAOFFSET_AUTO` for the offsets to be calculated by the SimConnect server. |  |
| _dwSizeOrType_ | Double word containing either the size of the client data in bytes, or one of the constant values defined in the table below (note that these definitions have a negative value, all positive values will be treated as a size parameter). | Integer |
| _fEpsilon_ | If data is requested only when it changes (see the _flags_ parameter of [SimConnect\_RequestClientData](SimConnect_RequestClientData.htm), a change will only be reported if it is greater than the value of this parameter (not greater than or equal to). The default is zero, so even the tiniest change will initiate the transmission of data. Set this value appropriately so insignificant changes are not transmitted. This can be used with integer data, the floating point _fEpsilon_ value is first truncated to its integer component before the comparison is made (for example, an _fEpsilon_ value of 2.9 truncates to 2, so a data change of 2 will not trigger a transmission, and a change of 3 will do so). This parameter only applies if one of the six constant values listed above has been set in the _dwSizeOrType_ parameter, if a size has been specified SimConnect has no record of the type of data being sent, so cannot do a meaningful comparison of values. | Float<br>(OPTIONAL) |
| _DatumID_ | Specifies a client defined datum ID. The default is zero. Use this to identify the data received if the data is being returned in tagged format (see the flags parameter of [SimConnect\_RequestClientData](SimConnect_RequestClientData.htm). There is no need to specify datum IDs if the data is not being returned in tagged format. | Integer<br>(OPTIONAL) |

The following table shows the different _dwSizeOrType_ constants that can be used:

| Constant | Value |
| --- | --- |
| `SIMCONNECT_CLIENTDATATYPE_INT8` | -1 |
| `SIMCONNECT_CLIENTDATATYPE_INT16` | -2 |
| `SIMCONNECT_CLIENTDATATYPE_INT32` | -3 |
| `SIMCONNECT_CLIENTDATATYPE_INT64` | -4 |
| `SIMCONNECT_CLIENTDATATYPE_FLOAT32` | -5 |
| `SIMCONNECT_CLIENTDATATYPE_FLOAT64` | -6 |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

This function must be called before a client data area can be written to or read from. Typically this function would be called once for each variable that is going to be read or written. Note that an error will not be given if the size of a data definition exceeds the size of the client area - this is to allow for the case where definitions are specified by one client before the relevant client area is created by another.

The most efficient use of client data areas is to group data that changes at the same time into the same data area. Minor performance improvements are gained by not using tagged data, or the _fEpsilon_ parameter, if they are not needed.

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SimConnect\_ClearClientDataDefinition](SimConnect_ClearClientDataDefinition.htm)
3. [SimConnect\_CreateClientData](SimConnect_CreateClientData.htm)
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