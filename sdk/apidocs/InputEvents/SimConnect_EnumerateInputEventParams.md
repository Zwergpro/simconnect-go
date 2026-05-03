SimConnect\_EnumerateInputEventParams

## SimConnect\_EnumerateInputEventParams

The **SimConnect\_EnumerateInputEventParams** function is used to retrieve a list of all parameters from an input event.

##### Syntax

```cpp
HRESULT SimConnect_EnumerateInputEvents(
    HANDLE hSimConnect,
    SIMCONNECT_DATA_REQUEST_ID RequestID
)
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | String |
| `RequestID` | The ID that will identify the current request in the response event |  |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

This function generate a [`SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS`](../Structures_And_Enumerations/SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS.htm) response where the Value member will contain a string. The string will contain a list of associated input event parameters separated by a semi-colon ";".

##### Example

The following basic example of use is from the DA62, for the `AS1000_MID_COM_1_Mic_Position` instrument, with the HashID "11675888408130357189":

```cpp
SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS:
{
Hash = 11675888408130357189;
Value = ";FLOAT64"
}
```

If an instrument has two arguments (in this example a `FLOAT64` and a `STRING`) then you would have to do the following:

```cpp
SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS:
{
Hash = <Hash>;
Value = ";FLOAT64;char[256]"
}
```

Then you would use the `SetInputEvent` function as follow:

```cpp
unsigned char byteValues[2048] = "";
// Create your value
value = "<myFloatValueInBytes><myStringValueInBytes>;"
m_simConnectObject.SetInputEvent(<Hash>, value);
```

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)
4. [SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENT\_PARAMS](../Structures_And_Enumerations/SIMCONNECT_RECV_ENUMERATE_INPUT_EVENT_PARAMS.htm)

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