SimConnect\_RetrieveString

## SimConnect\_RetrieveString

The **SimConnect\_RetrieveString** function is used to assist in retrieving variable length strings from a structure.

##### Syntax

```cpp
HRESULT SimConnect_RetrieveString(
    SIMCONNECT_RECV*  pData,
    DWORD  cbData,
    void*  pStringV,
    char**  ppszString,
    DWORD*  pcbString
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _pData_ | Pointer to a SIMCONNECT\_RECV structure, containing the data. | Pointer |
| _cbData_ | The size of the structure that inherits the SIMCONNECT\_RECV structure, in bytes. | Integer |
| _pStringV_ | Pointer to a the start of the variable length string within the structure. | Pointer |
| _ppszString_ | Specifies a pointer to a pointer to a character buffer that should be large enough to contain the maximum length of string that might be returned. On return this buffer should contain the retrieved string. | Pointer |
| _pcbString_ | Pointer to a DWORD. On return this contains the length of the string in bytes. | Pointer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
struct StructVS {
    char title[1];
    }
StructVS *pS = (StructVS*)&pObjData->dwData;
char *pszTitle;
DWORD cbTitle;
hr = SimConnect_RetrieveString(pData, cbData, &pS->strings, &pszTitle, &cbTitle)))
```

##### Remarks

This function does not communicate with the SimConnect server, but is a helper function to assist in the handling of variable length strings. Its counterpart is the [SimConnect\_InsertString](SimConnect_InsertString.htm) function. Note that this function works in the case where an empty string is in the structure returned by the server.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [Programming SimConnect Clients Using Managed Code](../../Programming_SimConnect_Clients_Using_Managed_Code.htm)
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