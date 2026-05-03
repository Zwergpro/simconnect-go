SimConnect\_GetInputEvent

## SimConnect\_GetInputEvent

The **SimConnect\_GetInputEvent** function is used to retrieve the value of a specific input event (identified by its hash).

##### Syntax

```cpp
HRESULT SimConnect_GetInputEvent(
    HANDLE hSimConnect,
    SIMCONNECT_DATA_REQUEST_ID RequestID,
    DWORD Hash,
)
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `RequestID` | The ID that will identify the current request in the response event. | Integer |
| `Hash` | Hash ID that will identify the desired inputEvent. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

This function might throw one of the following `SIMCONNECT_EXCEPTION`:

- `SIMCONNECT_EXCEPTION_ERROR` in case of internal errors
- `SIMCONNECT_EXCEPTION_GET_INPUT_EVENT_FAILED` if the given hash in wrong

##### Remarks

This function will generate a [`SIMCONNECT_RECV_GET_INPUT_EVENT`](../Structures_And_Enumerations/SIMCONNECT_RECV_GET_INPUT_EVENT.htm) response with the value of the input event referenced by the given hash (you can get the hashes for the available input events using [`SimConnect_EnumerateInputEvents`](SimConnect_EnumerateInputEvents.htm)).

##### Note for C\#

In the example code below, `data.eType` is castable to `data.Value[0]` in the case of `SIMCONNECT_INPUT_EVENT_TYPE.DOUBLE` and castable to `SimConnect.GetInputEventString` in case of `SIMCONNECT_INPUT_EVENT_TYPE.STRING`. This is the _default_ behaviour and you do not need to use `RegisterStruct` to "enable" it. You can, however, override it by using `RegisterStruct` with your own custom struct.

```cpp
// First, we call the function
private void M_oSimConnect_OnRecvEnumerateInputEvents(SimConnect sender, SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS data)
    {
        for (int i = 0; i < data.dwArraySize; ++i)
        {
            SIMCONNECT_INPUT_EVENT_DESCRIPTOR msg = (SIMCONNECT_INPUT_EVENT_DESCRIPTOR) data.rgData[i];
            if (msg.Name == "XXX" && msg.Type == SIMCONNECT_DATATYPE.FLOAT64)
            {
                m_oSimConnect.GetInputEvent(ReqID.Get, a.Hash);
            }
            else if (a.Name == "YYY" && msg.Type == SIMCONNECT_DATATYPE_STRING256)
            {
                m_oSimConnect.GetInputEvent(ReqID.Get, a.Hash);
            }
            m_oSimConnect.EnumerateInputEventParams(a.Hash);
        }
    }
// Callback
private void M_oSimConnect_OnRecvGetInputEvent(SimConnect sender, SIMCONNECT_RECV_GET_INPUT_EVENT data)
    {
        switch (data.eType)
        {
            case SIMCONNECT_INPUT_EVENT_TYPE.DOUBLE:
                double d = (double)data.Value[0];
                Console.WriteLine("Receive Double: " + d.ToString());
                break;
            case SIMCONNECT_INPUT_EVENT_TYPE.STRING:
                SimConnect.InputEventString str = (SimConnect.InputEventString)data.Value[0];
                Console.WriteLine("Receive String: " + str.value.ToString());
                break;
            case SIMCONNECT_INPUT_EVENT_TYPE.NONE:
                Debug.Assert(false);
                break;
        }
    }
```

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_RECV\_GET\_INPUT\_EVENT](../Structures_And_Enumerations/SIMCONNECT_RECV_GET_INPUT_EVENT.htm)
4. [SIMCONNECT\_RECV\_ENUMERATE\_INPUT\_EVENTS](../Structures_And_Enumerations/SIMCONNECT_RECV_ENUMERATE_INPUT_EVENTS.htm)
5. [SIMCONNECT\_INPUT\_EVENT\_DESCRIPTOR](../Structures_And_Enumerations/SIMCONNECT_INPUT_EVENT_DESCRIPTOR.htm)
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