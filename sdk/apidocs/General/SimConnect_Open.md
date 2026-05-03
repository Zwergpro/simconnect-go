SimConnect\_Open

## SimConnect\_Open

The **SimConnect\_Open** function is used to send a request to the _Microsoft Flight Simulator_ server to open up communications with a new client.

##### Syntax

```cpp
HRESULT SimConnect_Open(
    HANDLE*  phSimConnect,
    LPCSTR  szName,
    HWND  hWnd,
    DWORD  UserEventWin32,
    HANDLE  hEventHandle,
    DWORD  ConfigIndex
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _phSimConnect_ | Pointer to a handle to a SimConnect object. | Integer |
| _szName_ | Pointer to a null-terminated string containing an appropriate name for the client program. | Integer |
| _hWnd_ | Handle to a Windows object. Set this to NULL if the handle is not being used.<br>_UserEventWin32_ | Integer |
| _UserEventWin32_ | Code number that the client can specify. Set this to 0 if it is not being used. | Integer |
| _hEventHandle_ | A Windows Event handle. A client can be written to respond to Windows Events, rather than use a polling and callback system, which can be a more efficient process if the client does not have to respond very frequently to changes in data in Microsoft Flight Simulator 2024. | Integer |
| _ConfigIndex_ | The configuration index. The `SimConnect.cfg` file can contain a number of configurations, identified in sections with the `[SimConnect]` or `[SimConnect.N]` titles. Setting this configuration index indicates which configuration settings to use for this SimConnect client. This is useful for applications that communicate with a number of different machines that are running _Microsoft Flight Simulator 2024_. The default configuration index is zero (matching a `[SimConnect]` entry in a `SimConnect.cfg` file). Note the E\_INVALIDARG return value. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |
| E\_INVALIDARG | A SimConnect section in the `SimConnect.cfg` file did not contain the config index requested in the parameters. |

##### Example

```cpp
HRESULT hr; HANDLE hSimConnect = NULL;
hr = SimConnect_Open(&hSimConnect, "Your Application Name", NULL, 0, 0, SIMCONNECT_OPEN_CONFIGINDEX_LOCAL);
```

##### Remarks

Most client applications will have one [SimConnect\_Open](SimConnect_Open.htm) call, and one corresponding [SimConnect\_Close](SimConnect_Close.htm) call. However in some applications, multiplayer in particular, multiple [SimConnect\_Open](SimConnect_Open.htm) calls may be necessary, in which case an array or list of handles will need to be maintained, and closed appropriately.

A client can optionally examine the [SIMCOMMENT\_RECV\_OPEN](../Structures_And_Enumerations/SIMCONNECT_RECV_OPEN.htm) structure that is returned after a call to [SimConnect\_Open](SimConnect_Open.htm). This structure gives versioning and build information that should be useful when multiple versions of SimConnect and multiple versions of _Microsoft Flight Simulator 2024_ that support it, are available.

If a remote client successfully establishes a link with Flight Simulator, but at some later time the network connection is lost, SimConnect functions will return the `NTSTATUS` error `STATUS_REMOTE_DISCONNECT` (0xC000013CL).

The [SIMCONNECT\_EXCEPTION\_VERSION\_MISMATCH](../Structures_And_Enumerations/SIMCONNECT_EXCEPTION.htm) exception will be returned if a versioning error has occurred, typically when a client built on a newer version of the SimConnect client dll attempts to work with an older version of the SimConnect server. If this exception is received the number 4 is returned in the _dwIndex_ parameter of the [SIMCONNECT\_RECV\_EXCEPTION](../Structures_And_Enumerations/SIMCONNECT_RECV_EXCEPTION.htm) structure.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_Close](SimConnect_Close.htm)
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