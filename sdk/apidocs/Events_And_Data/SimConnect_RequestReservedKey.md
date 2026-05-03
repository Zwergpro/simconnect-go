SimConnect\_RequestReservedKey

## SimConnect\_RequestReservedKey

The **SimConnect\_RequestReservedKey** function is used to request a specific keyboard TAB-key combination applies only to this client.

**NOTE**: The current status of this function is NO ERROR, NO RESPONSE.

##### Syntax

```cpp
HRESULT SimConnect_RequestReservedKey(
    HANDLE  hSimConnect,
    SIMCONNECT_CLIENT_EVENT_ID  EventID,
    const char*  szKeyChoice1,
    const char*  szKeyChoice2 = "",
    const char*  szKeyChoice3 = ""
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _EventID_ | Specifies the client defined event ID. | Integer |
| _szKeyChoice1_ | Null-terminated string containing the first key choice. Refer to the list below for all the choices that can be entered for these three parameters. | String |
| _szKeyChoice2_ | Null-terminated string containing the second key choice. | String |
| _szKeyChoice3_ | Null-terminated string containing the third key choice. | String |

The you can find a list of all the available key strings that can be used in the szKeyChoiceN parameter here:

- [Valid Input Strings](../InputEvents/SimConnect_MapInputEventToClientEvent.htm#input_strings)

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Remarks

A successful call to this function will result in a [SIMCONNECT\_RECV\_RESERVED\_KEY](../Structures_And_Enumerations/SIMCONNECT_RECV_RESERVED_KEY.htm) structure being returned, with the key that has been assigned to this client. The first of the three that can be assigned will be the choice, unless all three are already taken, in which case a null string will be returned.

The `szKeyChoice` parameters should be a single character (such as "A"), which is requesting that the key combination TAB-A is reserved for this client. All reserved keys are TAB-key combinations.

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