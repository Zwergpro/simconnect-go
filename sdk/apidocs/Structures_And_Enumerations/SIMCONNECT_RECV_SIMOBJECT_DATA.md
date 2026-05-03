SIMCONNECT\_RECV\_SIMOBJECT\_DATA

## SIMCONNECT\_RECV\_SIMOBJECT\_DATA

The **SIMCONNECT\_RECV\_SIMOBJECT\_DATA** structure will be received by the client after a successful call to [SimConnect\_RequestDataOnSimObject](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm) or [SimConnect\_RequestDataOnSimObjectType](../Events_And_Data/SimConnect_RequestDataOnSimObjectType.htm).

##### Syntax

```cpp
struct SIMCONNECT_RECV_SIMOBJECT_DATA : public SIMCONNECT_RECV {
    DWORD  dwRequestID;
    DWORD  dwObjectID;
    DWORD  dwDefineID;
    DWORD  dwFlags;
    DWORD  dwentrynumber;
    DWORD  dwoutof;
    DWORD  dwDefineCount;
    DWORD  dwData;
    };
```

##### Members

| Member | Description |
| `dwRequestID` | The ID of the client defined request. |
| `dwObjectID` | The ID of the client defined object. |
| `dwDefineID` | The ID of the client defined data definition. |
| `dwFlags` | The flags that were set for this data request, see [SimConnect\_RequestDataOnSimObject](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm) for a description of the flags. This parameter will always be set to zero if the call was [SimConnect\_RequestDataOnSimObjectType](../Events_And_Data/SimConnect_RequestDataOnSimObjectType.htm). |
| `dwentrynumber` | If multiple objects are being returned, this is the index number of this object out of a total of `dwoutof`. This will always be 1 if the call was [SimConnect\_RequestDataOnSimObject](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm), and can be 0 or more if the call was [SimConnect\_RequestDataOnSimObjectType](../Events_And_Data/SimConnect_RequestDataOnSimObjectType.htm). |
| `dwoutof` | The total number of objects being returned. Note that `dwentrynumber` and `dwoutof` start with 1 not 0, so if two objects are being returned dwentrynumber and dwoutof pairs will be 1,2 and 2,2 for the two objects. This will always be 1 if the call was [SimConnect\_RequestDataOnSimObject](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm), and can be 0 or more if the call was [SimConnect\_RequestDataOnSimObjectType](../Events_And_Data/SimConnect_RequestDataOnSimObjectType.htm). |
| `dwDefineCount` | The number of 8-byte elements in the `dwData` array. |
| `dwData` | A data array containing information on a specified object in 8-byte (double word) elements. The length of the array is `dwDefineCount`.<br>The format of this buffer will depend of presence (or absence) of the `SIMCONNECT_DATA_REQUEST_FLAG_TAGGED` flag sent with the [SimConnect\_RequestDataOnSimObject](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm).<br>`SIMCONNECT_DATA_REQUEST_FLAG_TAGGED` disabled:<br>1. The various elements of the data definition are stored right next to each other, with no padding.<br>2. Data will look like: `DATA_0 DATA_1 DATA_2...`<br>`SIMCONNECT_DATA_REQUEST_FLAG_TAGGED` enabled:<br>1. Before each element of the data, the index of the elements will be stored as a DWORD. Then the actual data will be stored with no padding.<br>2. Data will look like: `DWORD_0 DATA_0 DWORD_1 DATA_1 DWORD_2 DATA_2...` |

##### Remarks

This structure inherits the `SIMCONNECT_RECV` structure and is returned when the `dwID` parameter of `SIMCONNECT_RECV` is set to `SIMCONNECT_RECV_ID_SIMOBJECT_DATA`.

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