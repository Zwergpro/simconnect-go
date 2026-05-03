SimConnect\_GetLastSentPacketID

## SimConnect\_GetLastSentPacketID

The **SimConnect\_GetLastSentPacketID** function returns the ID of the last packet sent to the SimConnect server.

##### Syntax

```cpp
HRESULT SimConnect_GetLastSentPacketID(
    HANDLE  hSimConnect,
    DWORD*  pdwSendID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _pdwSendID_ | Pointer to a double word containing the ID of the last sent packet. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
DWORD dwLastID;
hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_MY_EVENT, "Custom.Event");
hr = SimConnect_TransmitClientEvent(hSimConnect, 0, EVENT_MY_EVENT, 0, SIMCONNECT_GROUP_PRIORITY_HIGHEST, 0);
hr = SimConnect_GetLastSentPacketID(hSimConnect, &dwLastID);
```

##### Remarks

This function should be used in conjunction with returned structures of type [SIMCONNECT\_RECV\_EXCEPTION](../Structures_And_Enumerations/SIMCONNECT_RECV_EXCEPTION.htm) to help pinpoint errors (exceptions) returned by the server. This is done by matching the send ID returned with the exception, with the number returned by this function and stored appropriately. This function is primarily intended to be used while debugging and testing the client application, rather than in a final retail build.

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