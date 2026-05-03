SimConnect\_SetDataOnSimObject

## SimConnect\_SetDataOnSimObject

The **SimConnect\_SetDataOnSimObject** function is used to make changes to the data properties of an object.

##### Syntax

```cpp
HRESULT SimConnect_SetDataOnSimObject(
    HANDLE  hSimConnect,
    SIMCONNECT_DATA_DEFINITION_ID  DefineID,
    SIMCONNECT_OBJECT_ID  ObjectID,
    SIMCONNECT_DATA_SET_FLAG  Flags,
    DWORD  ArrayCount,
    DWORD  cbUnitSize,
    void*  pDataSet
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _DefineID_ | Specifies the ID of the client defined data definition. | Integer |
| _ObjectID_ | Specifies the ID of the Microsoft Flight Simulator object that the data should be about. This ID can be `SIMCONNECT_OBJECT_ID_USER` (to specify the user's aircraft) or obtained from a [SIMCONNECT\_RECV\_SIMOBJECT\_DATA\_BYTYPE](../Structures_And_Enumerations/SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE.htm) structure after a call to [SimConnect\_RequestDataOnSimObjectType](SimConnect_RequestDataOnSimObjectType.htm). Also refer to the note on "multiplayer mode" at the end of the remarks for [SimConnect\_RequestDataOnSimObject](SimConnect_RequestDataOnSimObject.htm). | Integer |
| _Flags_ | Null, or one or more of the flags listed below. | Flag |
| _ArrayCount_ | Specifies the number of elements in the data array. A count of zero is interpreted as one element. Ensure that the data array has been initialized completely before transmitting it to Microsoft Flight Simulator. Failure to properly initialize all array elements may result in unexpected behavior. | Integer |
| _cbUnitSize_ | Specifies the size of each element in the data array in bytes. | Integer |
| _pDataSet_ | Pointer to the data that is to be written. If the data is not in tagged format, this should point to the block of data. If the data is in tagged format this should point to the first tag name ( _datumID_), which is always four bytes long, which should be followed by the data itself. Any number of tag name/value pairs can be specified this way, the server will use the _cbUnitSize_ parameter to determine how much data has been sent. | Integer |

The table below shows the possible values for the _flags_ parameter:

| Return value | Description |
| --- | --- |
| SIMCONNECT\_DATA\_SET\_FLAG\_TAGGED | The data to be set is being sent in tagged format. If this flag is not set then the entire client data area should be replaced with new data. Refer to the _pDataSet_ parameter and [SimConnect\_RequestClientData](SimConnect_RequestClientData.htm) for more details on the tagged format. |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

##### Example

```cpp
static enum DATA_DEFINE_ID {
    DEFINITION3
    };

hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION3, "Initial Position", "NULL", SIMCONNECT_DATATYPE_INITPOSITION, 0);
SIMCONNECT_DATA_INITPOSITION Init;
    Init.Altitude = 5000.0;
    Init.Latitude = 47.64210;
    Init.Longitude = -122.13010;
    Init.Pitch = -0.0;
    Init.Bank = -1.0;
    Init.Heading = 180.0;
    Init.OnGround = 0;
    Init.Airspeed = 0;
SimConnect_SetDataOnSimObject(hSimConnect, DEFINITION3, SIMCONNECT_OBJECT_ID_USER, 0, sizeof(Init), &Init);
```

##### Remarks

The data that is set on an object is defined in a data definition (see the [SimConnect\_AddToDataDefinition](SimConnect_AddToDataDefinition.htm) function). This data can include the following structures: [SIMCONNECT\_DATA\_WAYPOINT](../Structures_And_Enumerations/SIMCONNECT_DATA_WAYPOINT.htm), [SIMCONNECT\_DATA\_INITPOSITION](../Structures_And_Enumerations/SIMCONNECT_DATA_INITPOSITION.htm), and [SIMCONNECT\_DATA\_MARKERSTATE](../Structures_And_Enumerations/SIMCONNECT_DATA_MARKERSTATE.htm). Any number of waypoints can be given to an AI object using a single call to this function, and any number of marker state structures can also be combined into an array.

The [Simulation Variables](../../../SimVars/Simulation_Variables.htm) documents include a column indicating whether variables can be written to or not. An exception will be returned if an attempt is made to write to a variable that cannot be set in this way.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AddToDataDefinition](SimConnect_AddToDataDefinition.htm)
4. [SimConnect\_ClearDataDefinition](SimConnect_ClearDataDefinition.htm)
5. [SimConnect\_RequestDataOnSimObject](SimConnect_RequestDataOnSimObject.htm)
6. [SimConnect\_RequestDataOnSimObjectType](SimConnect_RequestDataOnSimObjectType.htm)
7. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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