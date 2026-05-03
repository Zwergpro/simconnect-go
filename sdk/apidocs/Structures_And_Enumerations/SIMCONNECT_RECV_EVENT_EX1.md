SIMCONNECT\_RECV\_EVENT\_EX1

## SIMCONNECT\_RECV\_EVENT\_EX1

The **SIMCONNECT\_RECV\_EVENT\_EX1** structure is used to return an event ID to the client, along with up to 5 parameters.

##### Syntax

```cpp
struct SIMCONNECT_RECV_EVENT_EX1 : public SIMCONNECT_RECV {
    DWORD  uGroupID;
    DWORD  uEventID;
    DWORD  dwData0;
    DWORD  dwData1;
    DWORD  dwData2;
    DWORD  dwData3;
    DWORD  dwData4;
    };
```

##### Members

| Member | Description |
| `uGroupID` | The ID of the client defined group, or the special case value: `UNKNOWN_GROUP` (which equals `DWORD_MAX`). |
| `uEventID` | The ID of the client defined event that has been requested (such as EVENT\_1 or EVENT\_BRAKES). |
| `dwData`<br>to<br>`dwData` | Each of these 5 paramters corresponds to a parameter that was passed along with the event. |

##### Remarks

This structure inherits the [SIMCONNECT\_RECV](SIMCONNECT_RECV.htm) structure and is returned when the dwID parameter of [SIMCONNECT\_RECV](SIMCONNECT_RECV.htm) is set to [SIMCONNECT\_RECV\_ID\_EVENT](SIMCONNECT_RECV_ID.htm).

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SIMCONNECT\_RECV\_ID](SIMCONNECT_RECV_ID.htm)
4. [SIMCONNECT\_RECV\_EVENT](SIMCONNECT_RECV_EVENT.htm)
5. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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