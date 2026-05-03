SimConnect\_InsertString

## SimConnect\_InsertString

The **SimConnect\_InsertString** function is used to assist in adding variable length strings to a structure.

##### Syntax

```cpp
HRESULT SimConnect_InsertString(
    char*  pDest,
    DWORD  cbDest,
    void**  ppEnd,
    DWORD*  pcbStringV,
    const char*  pSource
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _pDest_ | Pointer to where the source string is to be written in the destination object. | Pointer |
| _cbDest_ | The size of the remaining space in the destination object. | Integer |
| _ppEnd_ | Pointer to a pointer, (usually a pointer to a char pointer). On return the pointer locates the end of the string in the structure, and hence the starting position for any other string to be included in the structure. | Pointer |
| _pcbStringV_ | Pointer to a DWORD. On returning this DWORD will contain the size of the source string in bytes. | Pointer |
| _pSource_ | Pointer to the source string. | Pointer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

This function does not communicate with the SimConnect server, but is a helper function to assist in the handling of variable length strings. Its counterpart is the [SimConnect\_RetrieveString](SimConnect_RetrieveString.htm) function.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_RetrieveString](SimConnect_RetrieveString.htm)
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