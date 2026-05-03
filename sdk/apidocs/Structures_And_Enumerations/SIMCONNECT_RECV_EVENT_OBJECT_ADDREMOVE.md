SIMCONNECT\_RECV\_EVENT\_OBJECT\_ADDREMOVE

## SIMCONNECT\_RECV\_EVENT\_OBJECT\_ADDREMOVE

The **SIMCONNECT\_RECV\_EVENT\_OBJECT\_ADDREMOVE** structure is used to return the type and ID of an AI object that has been added or removed from the simulation, by any client.

##### Syntax

```cpp
struct #SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE : SIMCONNECT_RECV_EVENT {
    SIMCONNECT_SIMOBJECT_TYPE  eObjType;
    };
```

##### Members

| Member | Description |
| `eObjType` | Specifies the type of object that was added or removed. One member of the [SIMCONNECT\_SIMOBJECT\_TYPE](SIMCONNECT_SIMOBJECT_TYPE.htm) enumeration. |

##### Remarks

This structure inherits the [`SIMCONNECT_RECV_EVENT`](SIMCONNECT_RECV_EVENT.htm) structure, which inherits the [`SIMCONNECT_RECV`](SIMCONNECT_RECV.htm) structure, and is returned when the dwID parameter of [`SIMCONNECT_RECV`](SIMCONNECT_RECV.htm) is set to [`SIMCONNECT_RECV_ID_EVENT_ADDREMOVE`](SIMCONNECT_RECV_ID.htm). A client can determine whether the object was added or removed from its own event ID that was provided as a parameter to the [SimConnect\_SubscribeToSystemEvent](../General/SimConnect_SubscribeToSystemEvent.htm) function.

The ID of the object added or removed is returned in the `dwData` parameter (a member of the `SIMCONNECT_RECV_EVENT` structure).

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