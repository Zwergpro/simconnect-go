SIMCONNECT\_RECV\_EVENT\_FILENAME

## SIMCONNECT\_RECV\_EVENT\_FILENAME

The **SIMCONNECT\_RECV\_EVENT\_FILENAME** structure is used to return a filename and an event ID to the client.

##### Syntax

```cpp
struct SIMCONNECT_RECV_EVENT_FILENAME : SIMCONNECT_RECV_EVENT {
    char  szFileName[MAX_PATH];
    DWORD  dwFlags;
    };
```

##### Members

| Member | Description |
| `szFileName[MAX_PATH]` | The returned filename. |
| `dwFlags` | Reserved, should be 0. |

##### Remarks

This structure inherits the `SIMCONNECT_RECV_EVENT` structure, and is used to attach a filename to the returned event.

When the `SIMCONNECT_RECV` structure `dwID` parameter is set to [`SIMCONNECT_RECV_EVENT_FILENAME`](SIMCONNECT_RECV_EVENT_FILENAME.htm), this structure is returned.

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