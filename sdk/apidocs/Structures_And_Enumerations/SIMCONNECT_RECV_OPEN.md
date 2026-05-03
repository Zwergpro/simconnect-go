SIMCONNECT\_RECV\_OPEN

## SIMCONNECT\_RECV\_OPEN

The **SIMCONNECT\_RECV\_OPEN** structure is used to return information to the client, after a successful call to [SimConnect\_Open](../General/SimConnect_Open.htm).

##### Syntax

```cpp
struct SIMCONNECT_RECV_OPEN : public SIMCONNECT_RECV {
    char  szApplicationName[256];
    DWORD  dwApplicationVersionMajor;
    DWORD  dwApplicationVersionMinor;
    DWORD  dwApplicationBuildMajor;
    DWORD  dwApplicationBuildMinor;
    DWORD  dwSimConnectVersionMajor;
    DWORD  dwSimConnectVersionMinor;
    DWORD  dwSimConnectBuildMajor;
    DWORD  dwSimConnectBuildMinor;
    DWORD  dwReserved1;
    DWORD  dwReserved2;
    };
```

##### Members

| Member | Description |
| `szApplicationName[256]` | Null-terminated string containing the application name. |
| `dwApplicationVersionMajor` | Double word containing the application version major number. |
| `dwApplicationVersionMinor` | Double word containing the application version minor number. |
| `dwApplicationBuildMajor` | Double word containing the application build major number. |
| `dwApplicationBuildMinor` | Double word containing the application build minor number. |
| `dwSimConnectVersionMajor` | Double word containing the SimConnect version major number. |
| `dwSimConnectVersionMinor` | Double word containing the SimConnect version minor number. |
| `dwSimConnectBuildMajor` | Double word containing the SimConnect build major number. |
| `dwSimConnectBuildMinor` | Double word containing the SimConnect build minor number. |
| `dwReserved1` | Reserved. |
| `dwReserved2` | Reserved. |

##### Remarks

This structure inherits the `SIMCONNECT_RECV` structure and is returned when the `dwID` parameter of `SIMCONNECT_RECV` is set to [SIMCONNECT\_RECV\_ID\_OPEN](SIMCONNECT_RECV_ID.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_Open](../General/SimConnect_Open.htm)
4. [SIMCONNECT\_RECV](SIMCONNECT_RECV.htm)
5. [SIMCONNECT\_RECV\_OPEN](SIMCONNECT_RECV_OPEN.htm)
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