SimConnect\_CallCommBusEvent

## SimConnect\_CallCommBusEvent

The **SimConnect\_CallCommBusEvent** function is used to call a communication (CommBus) event.

**C++****C#**

##### Syntax

```cpp
HRESULT SimConnect_CallCommBusEvent(
    HANDLE hSimConnect,
    const char * EventName,
    SIMCONNECT_COMM_BUS_BROADCAST_TO BroadcastTo,
    DWORD BufferSize,
    const char * Data
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _EventName_ | Specifies the name of the event to call. Name of the event to call. Note that there is no need to register the event in the client prior to calling it. | Integer |
| _BroadcastTo_ | One of the `SIMCONNECT_COMM_BUS_BROADCAST_TO` enum members that defines where the event should be broadcast (WASM, JS, or SimConnect). | Integer |
| _BufferSize_ | This is the size (in bytes) of the `Data` buffer. | Integer |
| _Data_ | This is the data to send alongside the event. The data will be received differently depending on the receiving platform:<br>- SimConnect C# - A string will be received.<br>- SimConnect C++ - A `char*` string and a size will be received.<br>- JavaScript - A string will be received.<br>- WASM - A `char*` string and a size will be received. | String |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Syntax

```cs
void SimConnect::CallCommBusEvent(
    string EventName,
    SIMCONNECT_COMM_BUS_BROADCAST_TO BroadcastTo,
    object Data
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _EventName_ | Specifies the name of the event to call. Note that there is no need to register the event in the client prior to calling it. | String |
| _BroadcastTo_ | One of the `SIMCONNECT_COMM_BUS_BROADCAST_TO` enum members that defines where the event should be broadcast (WASM, JS, or SimConnect). | Enum |
| _Data_ | An abstract object containing the data to send with the event, which will be will be converted with Marshal into a pointer and a size. The data will be received differently depending on the receiving platform:<br>- SimConnect C# - A string will be received.<br>- SimConnect C++ - A `char*` string and a size will be received.<br>- JavaScript - A string will be received.<br>- WASM - A `char*` string and a size will be received. | Object |

##### Return Values

N/A (use a try/catch to test for errors).


##### Remarks

N/A

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)
4. [Communication](Communication.htm)

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