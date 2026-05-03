SimConnect\_ExecuteAction

## SimConnect\_ExecuteAction

The **SimConnect\_ExecuteAction** function can be used to call an action from an XML file with a variable number of parameters. You can find a full list of available actions here: [SimConnect Actions](../SimConnect_Actions.htm)

##### Syntax

```cpp
HRESULT SimConnect_ExecuteAction(
    HANDLE hSimConnect,
    DWORD cbRequestID,
    const char* szActionID,
    DWORD cbUnitSize
    void* pParamValues
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `cbRequestID` | ID of the request to retrieve it in the callback response. | Integer |
| `szActionID` | The name of the action. | Char |
| `cbUnitSize` | The size of the `pParamValues` data. | Integer |
| `pParamValues` | A `void*` that contains all the values packed as byte data. | Void |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

Note that it is possible to get the following `SIMCONNECT_EXCEPTION` values when using this function:

- `SIMCONNECT_EXCEPTION_ACTION_NOT_FOUND` (with offset 2)
- `SIMCONNECT_EXCEPTION_NOT_AN_ACTION` (with offset 2)
- `SIMCONNECT_EXCEPTION_INCORRECT_ACTION_PARAMS` (with offset 4)

##### Remarks

The example below shows one way of how to construct the `pParamValues` parameter for use with this function:

**NOTE**: All strings are of a fixed size (256).

```cpp
HANDLE  hSimConnect;
// SetHighlight (Bool) - FocusDuration (Float) - SetCamera (Bool) - InstrumentPartId (String 256) - InstrumentHtmlId (String 256)
int paramValuesSize = sizeof(bool) + sizeof(float) + sizeof(bool) + 256 + 256;
unsigned char* pData = new unsigned char[paramValuesSize];
unsigned char* pDataBegin = pData;
bool setHighlightVal = TRUE;
*pData = setHighlightVal;
pData += sizeof(bool);
float focusDurationVal = 5.0f;
memcpy(pData, &focusDurationVal, sizeof(float));
pData += sizeof(float);
bool setCameraVal = FALSE;
*pData = setCameraVal;
pData += sizeof(bool);
char instrumentPartIdVal[256] = { "FUEL_SWITCH_TRANSFER" };
memcpy(pData, &instrumentPartIdVal, 256);
pData += 256;
char instrumentHtmlVal[256] = { "" };
memcpy(pData, &instrumentHtmlVal, 256);
pData += 256;
pData = pDataBegin;
SimConnect_ExecuteAction(hSimConnect, 0, "FocusInstrumentAction", paramValuesSize, pData);
```

Alternatively, you can use a struct to form the `pParamValues` parameter:

**NOTE**: The struct must be `#pragma pack (1)` because of struct padding.

```cpp
#pragma pack (1)
struct FocusInstrumentParams
{
    bool highlightPart;
    float focusDuration;
    bool setCamera;
    char instrumentPartId[256];
    char instrumentHtml[256];
};
#pragma pack (0)
///
HANDLE  hSimConnect;
FocusInstrumentParams params;
params.highlightPart = TRUE;
params.focusDuration = 5.0f;
params.setCamera = FALSE;
strcpy_s(params.instrumentPartId, "FUEL_SWITCH_TRANSFER");
strcpy_s(params.instrumentHtml, "\0");
SimConnect_ExecuteAction(hSimConnect, 0, "FocusInstrumentAction", sizeof(params), &params);
```

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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