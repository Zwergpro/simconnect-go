SIMCONNECT\_RECV\_EVENT

## SIMCONNECT\_RECV\_EVENT

The **SIMCONNECT\_RECV\_EVENT** structure is used to return an event ID to the client.

##### Syntax

```cpp
struct SIMCONNECT_RECV_EVENT : public SIMCONNECT_RECV {
    DWORD  uGroupID;
    DWORD  uEventID;
    DWORD  dwData;
    };
```

##### Members

| Member | Description |
| `uGroupID` | The ID of the client defined group, or the special case value: `UNKNOWN_GROUP` (which equals `DWORD_MAX`). |
| `uEventID` | The ID of the client defined event that has been requested (such as EVENT\_1 or EVENT\_BRAKES). |
| `dwData` | This value is usually zero, but some events require further qualification. For example, joystick movement events require a movement value in addition to the notification that the joystick has been moved (see [SimConnect\_MapInputEventToClientEvent](../InputEvents/SimConnect_MapInputEventToClientEvent.htm) for more information). |

##### Remarks

This structure inherits the [SIMCONNECT\_RECV](SIMCONNECT_RECV.htm) structure and is returned when the dwID parameter of [SIMCONNECT\_RECV](SIMCONNECT_RECV.htm) is set to [SIMCONNECT\_RECV\_ID\_EVENT](SIMCONNECT_RECV_ID.htm). This structure is inherited by several other structures:

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SIMCONNECT\_RECV\_EVENT\_FILENAME](SIMCONNECT_RECV_EVENT_FILENAME.htm)
3. [SIMCONNECT\_RECV\_EVENT\_FRAME](SIMCONNECT_RECV_EVENT_FRAME.htm)
4. [SIMCONNECT\_RECV\_EVENT](SIMCONNECT_RECV_EVENT.htm)
5. [SIMCONNECT\_RECV\_EVENT\_OBJECT\_ADDREMOVE](SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE.htm)
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