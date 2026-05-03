SIMCONNECT\_RECV\_EXCEPTION

## SIMCONNECT\_RECV\_EXCEPTION

The **SIMCONNECT\_RECV\_EXCEPTION** structure is used with the `SIMCONNECT_EXCEPTION` enumeration type to return information on an error that has occurred.

##### Syntax

```cpp
struct SIMCONNECT_RECV_EXCEPTION : public SIMCONNECT_RECV {
    DWORD  dwException;
    DWORD  dwSendID;
    DWORD  dwIndex;
    };
```

##### Members

| Member | Description |
| `dwException` | One member of the `SIMCONNECT_EXCEPTION` enumeration type, indicating which error has occurred. |
| `dwSendID` | The ID of the packet that contained the error, see Remarks below.<br>Special case: `UNKNOWN_SENDID = 0`.<br>Note that if this special case is returned, there has been an _internal_ problem. |
| `dwIndex` | The index number (starting at 1) of the first parameter that caused an error.<br>Special case: `UNKNOWN_INDEX = 0`. |

##### Remarks

This structure inherits the `SIMCONNECT_RECV` structure and is returned when the `dwID` parameter of `SIMCONNECT_RECV` is set to [`SIMCONNECT_RECV_ID_EXCEPTION`](SIMCONNECT_RECV_ID.htm). In order to match the `dwSendID` parameter returned here, with the ID of a request, use the `SimConnect_GetLastSentPacketID` call after each request is made.

Note that the `HRESULT` errors returned after each API call do not involve any communication with the SimConnect server, but are simply client-side errors that are returned immediately. Test for exceptions to check for server-side errors.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_RECV\_ID](SIMCONNECT_RECV_ID.htm)
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