Communication API

## COMMUNICATION API

The following SimConnect functions and enums/structs are used for for communicating data between various SimConnect clients, WebAssembly modules, and JavaScript gauges. You can find information on the other platform communication functions from the following pages:

- [JavaScript Communication API](../../../JavaScript/Communication_API/Communication_API.htm)
- [WebAssembly Communication API](../../../WASM/Communication_API/Communication_API.htm)

| Function | Description |
| `SimConnect_CallCommBusEvent` | Used to call a communication (CommBus) event. |
| `SimConnect_SubscribeToCommBusEvent` | Used to subscribe the client to a communication (CommBus) event |
| `SimConnect_UnsubscribeToCommBusEvent` | Used to unsubscribe the client from a communication (CommBus) event. |

For a list of other SimConnect functions please see the [SimConnect API Reference](../../SimConnect_API_Reference.htm).

When using this API it's worth noting the following:

- JSON is a good way to format data for communication between the different platforms. In C++, a version of [RapidJSON](https://rapidjson.org/) is provided as part of the SDK - with some fixes to work in Web Assembly - and can help you when sharing information.

- You can add an ID in function arguments to know the caller, and ensure you are dealing with the correct events.

- All calls are asynchronous.

You can find a sample project to use as a reference when using the Communication API here:

- [CommBus](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm#commbus)

There are also debugging tools available in DevMode which can be used to help debug Flow Events and Communication API events, which are described on the following page:

- [Debug Platform Dispatcher](../../../../2_DevMode/Menus/Debug/Debug_Platform_Dispatcher.htm)

### Calling And Registering Events

The communication (CommBus) API is cross platform, and as such, can be used to call _and_ register events on the different supported platforms. Below you can find brief overviews for how this can be done.

**C++****C#**

#### Register an event in SimConnect

Open a SimConnect connection:

```cpp
HRESULT hr = E_FAIL;
HANDLE hSimConnect = NULL;
hr = SimConnect_Open(&hSimConnect, "Connection name", NULL, 0, 0, 0);
```

Create a dispatch function:

```cpp
std::string str;
void CALLBACK Dispatch(SIMCONNECT_RECV* pData, DWORD cbData, void* pContext)
{
    switch (pData->dwID)
    {
        case SIMCONNECT_RECV_ID_COMM_BUS: // The received data are a comm bus event
        {
            SIMCONNECT_RECV_COMM_BUS* pCommBusEvt = (SIMCONNECT_RECV_COMM_BUS*)pData;
            switch (pCommBusEvt->uEventID) // Check which event this is
            {
                case EVENT_COMM_BUS:
                {
                    if (pCommBusEvt->dwOutOf == 1)
                    {
                        PrintReceivedData(pCommBusEvt->rgData);
                    }
                    else
                    {
                        receptionBuffer += data.rgData;
                        if (pCommBusEvt->dwEntryNumber + 1 == pCommBusEvt->dwOutOf)
                        {
                            PrintReceivedData(receptionBuffer);
                            receptionBuffer = "";
                        }
                    }
                    break;
                }
            }
        }
    }
}
```

Call the dispatch:

```cpp
while( !quit )
{
    SimConnect_CallDispatch(hSimConnect, Dispatch, NULL);
    Sleep(1000);
}
```

**NOTE**: In WebAssembly, there is no needs to loop on `SimConnect_CallDispatch`, the given callback will be called each frame.

Register the event:

```cpp
HRESULT hr = E_FAIL;
hr = SimConnect_SubscribeToCommBusEvent(hSimConnect, EVENT_COMM_BUS, "SimConnectEventName");
```

#### Call JS from SimConnectCall JS from SimConnect

In your JS file you need to create the callback:

```js
myJsCallback(args) { ... }
```

You then need to register the view listener:

```js
this.commBusListener = RegisterCommBusListener();
```

Next you would connect the callback to an event name:

```js
this.commBuListener.on("MyJsCallback", this.myJsCallback);
```

In the SimConnect client you need to open a SimConnect connection:

```cpp
HRESULT hr = E_FAIL;
HANDLE hSimConnect = NULL;
hr = SimConnect_Open(&hSimConnect, "Connection name", NULL, 0, 0, 0);
```

Call the event:

```cpp
HRESULT hr = E_FAIL;
hr = SimConnect_CallCommBusEvent(hSimConnect, eventName, broadcastTo, strlen(buffer) + 1, buffer);
```

#### Call WASM From SimConnectCall WASM From SimConnect

Create a WASM callback:

```wasm
static void MyWasmCallback(const char* args, unsigned int size, void* ctx) { ... }
```

Register the event:

```wasm
fsCommBusRegister("MyWasmCallback", MyWasmCallback);
```

Open a SimConnect connection:

```cpp
HRESULT hr = E_FAIL;
HANDLE hSimConnect = NULL;
hr = SimConnect_Open(&hSimConnect, "Connection name", NULL, 0, 0, 0);
```

Call the WASM event:

```cpp
HRESULT hr = E_FAIL;
hr = SimConnect_CallCommBusEvent(hSimConnect, eventName, broadcastTo, strlen(buffer) + 1, buffer);
```

#### Call SimConnect (C++) From SimConnect (C++)Call SimConnect (C++) From SimConnect (C++)

Open a SimConnect connection:

```cpp
HRESULT hr = E_FAIL;
HANDLE hSimConnect = NULL;
hr = SimConnect_Open(&hSimConnect, "Connection name", NULL, 0, 0, 0);
```

Create a dispatch function:

```cpp
std::string str;
void CALLBACK Dispatch(SIMCONNECT_RECV* pData, DWORD cbData, void* pContext)
{
    switch (pData->dwID)
    {
        case SIMCONNECT_RECV_ID_COMM_BUS: // The received data are a comm bus event
        {
            SIMCONNECT_RECV_COMM_BUS* pCommBusEvt = (SIMCONNECT_RECV_COMM_BUS*)pData;
            switch (pCommBusEvt->uEventID) // Check which event this is
            {
                case EVENT_COMM_BUS:
                {
                    if (pCommBusEvt->dwOutOf == 1)
                    {
                        PrintReceivedData(pCommBusEvt->rgData);
                    }
                    else
                    {
                        receptionBuffer += data.rgData;
                        if (pCommBusEvt->dwEntryNumber + 1 == pCommBusEvt->dwOutOf)
                        {
                            PrintReceivedData(receptionBuffer);
                            receptionBuffer = "";
                        }
                    }
                break;
            }
        }
    }
}
```

Call the dispatch:

```cpp
while( !quit )
{
SimConnect_CallDispatch(hSimConnect, Dispatch, NULL);
Sleep(1000);
}
```

Register the SimConnect event:

```cpp
HRESULT hr = E_FAIL;
hr = SimConnect_SubscribeToCommBusEvent(hSimConnect, EVENT_COMM_BUS, "SimConnectEventName");
```

Open a Simconnect connection:

```cpp
HRESULT hr = E_FAIL;
HANDLE hSimConnect = NULL;
hr = SimConnect_Open(&hSimConnect, "Connection name", NULL, 0, 0, 0);
```

Call the SimConnect event:

```cpp
HRESULT hr = E_FAIL;
hr = SimConnect_CallCommBusEvent(hSimConnect, eventName, broadcastTo, strlen(buffer) + 1, buffer);
```

#### Call SimConnect (C\#) From SimConnect (C++)Call SimConnect (C\#) From SimConnect (C++)

Open a SimConnect (C#) connection:

```cs
SimConnect m_oSimConnect = new SimConnect("Connection name", m_hWnd, WM_USER_SIMCONNECT, null, 0);
```

Create a CommBus callback:

```cs
string receptionBuffer;
void OnCommBusEvent(SimConnect sender, SIMCONNECT_RECV_COMM_BUS data)
{
    switch (data.uEventID)
    {
        case ((int)Events.CommBusEvent):
            {
                if (data.dwOutOf == 1)
                {
                    PrintReceivedData(data.rgData);
                }
                else
                {
                    receptionBuffer += data.rgData;
                    if (data.dwEntryNumber + 1 == data.dwOutOf)
                    {
                        PrintReceivedData(receptionBuffer);
                        receptionBuffer = "";
                    }
                }
            }
            break;
        default:
            break;
    }
}
```

Register a callback on Recv CommBus:

```cs
m_oSimConnect.OnRecvCommBus += new SimConnect.RecvCommBusEventHandler(OnCommBusEvent);
```

Register the SimConnect (C#) event:

```cs
m_oSimConnect.SubscribeToCommBusEvent(Events.CommBusEvent, "SimConnectEventName");
```

Open a SimConnect (C++) connection:

```cpp
HRESULT hr = E_FAIL;
HANDLE hSimConnect = NULL;
hr = SimConnect_Open(&hSimConnect, "Connection name", NULL, 0, 0, 0);
```

Call the SimConnect (C++) event:

```cpp
HRESULT hr = E_FAIL;
hr = SimConnect_CallCommBusEvent(hSimConnect, "SimConnectEventName", broadcastTo, strlen(buffer) + 1, buffer);
```

#### Register an event in SimConnect

Open a SimConnect connection:

```cs
SimConnect m_oSimConnect = new SimConnect("Connection name", m_hWnd, WM_USER_SIMCONNECT, null, 0);
```

Create a CommBus callback:

```cs
string receptionBuffer;
void OnCommBusEvent(SimConnect sender, SIMCONNECT_RECV_COMM_BUS data)
{
    switch (data.uEventID)
    {
        case ((int)Events.CommBusEvent):
            {
                if (data.dwOutOf == 1)
                {
                    PrintReceivedData(data.rgData);
                }
                else
                {
                    receptionBuffer += data.rgData;
                    if (data.dwEntryNumber + 1 == data.dwOutOf)
                    {
                        PrintReceivedData(receptionBuffer);
                        receptionBuffer = "";
                    }
                }
            }
            break;
        default:
            break;
    }
}
```

Register a callback on Recv CommBus:

```cs
m_oSimConnect.OnRecvCommBus += new SimConnect.RecvCommBusEventHandler(OnCommBusEvent);
```

Register the event:

```cs
m_oSimConnect.SubscribeToCommBusEvent(Events.CommBusEvent, "SimConnectEventName");
```

#### Call JS from SimConnectCall JS from SimConnect

In your JS file you need to create the callback:

```js
myJsCallback(args) { ... }
```

You then need to register the view listener:

```js
this.commBusListener = RegisterCommBusListener();
```

Next you would connect the callback to an event name:

```js
this.commBuListener.on("MyJsCallback", this.myJsCallback);
```

In the SimConnect client you need to open a SimConnect connection:

```cs
SimConnect m_oSimConnect = new SimConnect("Connection name", m_hWnd, WM_USER_SIMCONNECT, null, 0);
```

Call the event:

```cs
String message = "something";
m_oSimConnect.CallCommBusEvent("MyJsCallback", broadcastTo, message);
```

#### Call WASM from SimConnectCall WASM from SimConnect

Create a WASM callback:

```wasm
static void MyWasmCallback(const char* args, unsigned int size, void* ctx) { ... }
```

Register the event:

```wasm
fsCommBusRegister("MyWasmCallback", MyWasmCallback);
```

Open a SimConnect connection:

```cs
SimConnect m_oSimConnect = new SimConnect("Connection name", m_hWnd, WM_USER_SIMCONNECT, null, 0);
```

Call the WASM event:

```cs
String message = "something";
m_oSimConnect.CallCommBusEvent("MyWasmCallback", broadcastTo, message);
```

#### Call SimConnect (C++) From SimConnect (C\#)Call SimConnect (C++) From SimConnect (C\#)

Open a SimConnect (C++) connection:

```cs
HRESULT hr = E_FAIL;
HANDLE hSimConnect = NULL;
hr = SimConnect_Open(&hSimConnect, "Connection name", NULL, 0, 0, 0);
```

Create a dispatch function (C++):

```cs
std::string str;
void CALLBACK Dispatch(SIMCONNECT_RECV* pData, DWORD cbData, void* pContext)
{
    switch (pData->dwID)
    {
        case SIMCONNECT_RECV_ID_COMM_BUS: // The received data are a comm bus event
        {
            SIMCONNECT_RECV_COMM_BUS* pCommBusEvt = (SIMCONNECT_RECV_COMM_BUS*)pData;
            switch (pCommBusEvt->uEventID) // Check which event this is
            {
                case EVENT_COMM_BUS:
                {
                    if (pCommBusEvt->dwOutOf == 1)
                    {
                        PrintReceivedData(pCommBusEvt->rgData);
                    }
                    else
                    {
                        receptionBuffer += data.rgData;
                        if (pCommBusEvt->dwEntryNumber + 1 == pCommBusEvt->dwOutOf)
                        {
                            PrintReceivedData(receptionBuffer);
                            receptionBuffer = "";
                        }
                    }
                break;
            }
        }
    }
}
```

Call the dispatch (C++):

```cs
while( !quit )
{
SimConnect_CallDispatch(hSimConnect, Dispatch, NULL);
Sleep(1000);
}
```

Register the SimConnect (C++) event:

```cs
HRESULT hr = E_FAIL;
hr = SimConnect_SubscribeToCommBusEvent(hSimConnect, EVENT_COMM_BUS, "SimConnectEventName");
```

Open a SimConnect (C#) connection:

```cs
SimConnect m_oSimConnect = new SimConnect("Connection name", m_hWnd, WM_USER_SIMCONNECT, null, 0);
```

Call the SimConnect (C#) event:

```cs
String message = "truc";
m_oSimConnect.CallCommBusEvent("SimConnectEventName", broadcastTo, message);
```

#### Call SimConnect (C\#) From SimConnect (C\#)Call SimConnect (C\#) From SimConnect (C\#)

Open a SimConnect connection:

```cs
SimConnect m_oSimConnect = new SimConnect("Connection name", m_hWnd, WM_USER_SIMCONNECT, null, 0);
```

Create a CommBus callback:

```cs
string receptionBuffer;
void OnCommBusEvent(SimConnect sender, SIMCONNECT_RECV_COMM_BUS data)
{
    switch (data.uEventID)
    {
        case ((int)Events.CommBusEvent):
            {
                if (data.dwOutOf == 1)
                {
                    PrintReceivedData(data.rgData);
                }
                else
                {
                    receptionBuffer += data.rgData;
                    if (data.dwEntryNumber + 1 == data.dwOutOf)
                    {
                        PrintReceivedData(receptionBuffer);
                        receptionBuffer = "";
                    }
                }
            }
            break;
        default:
            break;
    }
}
```

Register a callback on Recv CommBus:

```cs
m_oSimConnect.OnRecvCommBus += new SimConnect.RecvCommBusEventHandler(OnCommBusEvent);
```

Register the SimConnect event:

```cs
m_oSimConnect.SubscribeToCommBusEvent(Events.CommBusEvent, "SimConnectEventName");
```

Open a SimConnect connection:

```cs
SimConnect m_oSimConnect = new SimConnect("Connection name", m_hWnd, WM_USER_SIMCONNECT, null, 0);
```

Call the SimConnect event:

```cs
String message = "something";
m_oSimConnect.CallCommBusEvent("SimConnectEventName", broadcastTo, message);
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