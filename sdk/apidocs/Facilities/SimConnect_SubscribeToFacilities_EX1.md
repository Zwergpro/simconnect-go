SimConnect\_SubscribeToFacilities\_EX1

## SimConnect\_SubscribeToFacilities\_EX1

The **SimConnect\_SubscribeToFacilities\_EX1** function is used to request notifications when a facility of a certain type is added to the facilities cache, with the ability to specify callbacks.

##### Syntax

```cpp
HRESULT SimConnect_SubscribeToFacilities_EX1(
    HANDLE hSimConnect,
    SIMCONNECT_FACILITY_LIST_TYPE type,
    SIMCONNECT_DATA_REQUEST_ID newElemInRangeRequestID,
    SIMCONNECT_DATA_REQUEST_ID oldElemOutRangeRequestID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| `hSimConnect` | Handle to a SimConnect object. | Integer |
| `type` | Specifies one member of the [`SIMCONNECT_FACILITY_LIST_TYPE`](../Structures_And_Enumerations/SIMCONNECT_FACILITY_LIST_TYPE.htm) enumeration type. | Enum |
| `newElemInRangeRequestID` | Request Id for messages about new elements considered to be in range of the reality bubble. If -1 is used, then the client won't receive messages for elements coming into range.<br>**NOTE**: This cannot be set to -1 if `oldElemOutRangeRequestID` is also -1. | Integer |
| `oldElemOutRangeRequestID` | Request Id for messages about an element that is newly considered out of range of the reality bubble. If -1 is used, then the client won't receive messages for elements coming into range.<br>**NOTE**: This cannot be set to -1 if `newElemInRangeRequestID` is also -1. | Integer |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| `S_OK` | The function succeeded. |
| `E_FAIL` | The function failed. |

##### Remarks

This function is a version of the [`SimConnect_SubscribeToFacilities`](SimConnect_SubscribeToFacilities.htm) and works in almost the exact same way with one exception. Using `SIMCONNECT_FACILITY_LIST_TYPE_AIRPORT` will now return a list of the airports within the current reality bubble, _not_ the entire world. You can also specify the callbacks to use when a facility either enters or leaves the reality bubble. Note that to unsubscribe you should use the companion function [`SimConnect_UnsubscribeToFacilities_EX1`](SimConnect_UnsubscribeToFacilities_EX1.htm).

##### See Also

- [SimConnect API Reference](../../SimConnect_API_Reference.htm)
- [SimConnect\_UnsubscribeToFacilities\_EX1](SimConnect_UnsubscribeToFacilities_EX1.htm)
- [SimConnect\_RequestFacilitiesList](SimConnect_SubscribeToFacilities.htm)
- [SimConnect\_SubscribeToFacilities](SimConnect_SubscribeToFacilities.htm)

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