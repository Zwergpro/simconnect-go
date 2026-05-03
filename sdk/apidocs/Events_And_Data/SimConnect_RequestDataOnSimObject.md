SimConnect\_RequestDataOnSimObject

## SimConnect\_RequestDataOnSimObject

The **SimConnect\_RequestDataOnSimObject** function is used to request when the SimConnect client is to receive data values for a specific object.

##### Syntax

```cpp
HRESULT SimConnect_RequestDataOnSimObject(
    HANDLE  hSimConnect,
    SIMCONNECT_DATA_REQUEST_ID  RequestID,
    SIMCONNECT_DATA_DEFINITION_ID  DefineID,
    SIMCONNECT_OBJECT_ID  ObjectID,
    SIMCONNECT_PERIOD  Period,
    SIMCONNECT_DATA_REQUEST_FLAG  Flags = 0,
    DWORD  origin = 0,
    DWORD  interval = 0,
    DWORD  limit = 0
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _RequestID_ | Specifies the ID of the client defined request. This is used later by the client to identify which data has been received. This value should be unique for each request, as the server will map your _RequestID_ to your last received _DefineID_, which means that in the case of two calls with identical _RequestID_ but different _DefineID_, the second call will overwrite the first one (assuming a _Period_ other than `SIMCONNECT_PERIOD_ONCE`). | Integer |
| _DefineID_ | Specifies the ID of the client defined [data definition](SimConnect_AddToDataDefinition.htm). | Integer |
| _ObjectID_ | Specifies the ID of the Microsoft Flight Simulator 2024 object that the data should be about. This ID can be `SIMCONNECT_OBJECT_ID_USER` (to specify the user's aircraft) or obtained from a `SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE` structure after a call to `SimConnect_RequestDataOnSimObjectType`. Also refer to the note on developing clients for multiplayer mode in the Remarks section below. | Integer |
| _Period_ | One member of the `SIMCONNECT_PERIOD` enumeration type, specifying how often the data is to be sent by the server and received by the client. | Integer |
| _Flags_ | A DWORD containing one or more of the values shown in the table below this one. | Integer<br>(OPTIONAL) |
| _origin_ | The number of _Period_ events that should elapse before transmission of the data begins. The default is zero, which means transmissions will start immediately. | Integer<br>(OPTIONAL) |
| _interval_ | The number of _Period_ events that should elapse between transmissions of the data. The default is zero, which means the data is transmitted every _Period_. | Integer<br>(OPTIONAL) |
| _limit_ | The number of times the data should be transmitted before this communication is ended. The default is zero, which means the data should be transmitted endlessly. | Integer<br>(OPTIONAL) |

The following table shows the available values for _F_ _lags_:

| Flag value | Description |
| --- | --- |
| 0 | The default, data will be sent strictly according to the defined period. |
| `SIMCONNECT_DATA_REQUEST_FLAG_CHANGED` | Data will only be sent to the client when one or more values have changed. If this is the only flag set, then all the variables in a data definition will be returned if just one of the values changes. |
| `SIMCONNECT_DATA_REQUEST_FLAG_TAGGED` | Requested data will be sent in tagged format (datum ID/value pairs). Tagged format requires that a datum reference ID is returned along with the data value, in order that the client code is able to identify the variable. This flag is usually set in conjunction with the previous flag, but it can be used on its own to return all the values in a data definition in datum ID/value pairs. See the `SIMCONNECT_RECV_SIMOBJECT_DATA` structure for more details. |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| `S_OK` | The function succeeded. |
| `E_FAIL` | The function failed. |

##### Example

```cpp
static enum DATA_DEFINE_ID {
    DEFINITION_1,
    DEFINITION_2
    };
static enum DATA_REQUEST_ID {
    REQUEST_1,
    REQUEST_2,
    };
struct Struct1
{
    double  kohlsmann;
    double  altitude;
    double  latitude;
    double  longitude;
};

hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_1, "Kohlsman setting hg", "inHg");
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_1, "Indicated Altitude", "feet");
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_1, "Plane Latitude", "degrees");
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_1, "Plane Longitude", "degrees");
      ....
hr = SimConnect_RequestDataOnSimObject(hSimConnect, REQUEST_2, DEFINITION_1, SIMCONNECT_OBJECT_ID_USER, SIMCONNECT_PERIOD_SECOND);
      ....
    case SIMCONNECT_RECV_ID_SIMOBJECT_DATA:
    {
    SIMCONNECT_RECV_SIMOBJECT_DATA *pObjData = (SIMCONNECT_RECV_SIMOBJECT_DATA*) pData;
    switch(pObjData->dwRequestID)
        {
        case REQUEST_2:
            Struct1 *pS = (Struct1*)&pObjData->dwData;
            break;
        }
    break;
    }
```

##### Remarks

Changing the _Period_ parameter or changing the content of a data definition has a higher performance cost than changing the _origin_, _interval_ or _limit_ parameters. So to temporarily turn off data requests, especially for short periods of time, consider setting the _interval_ parameter to a very high value (such as `DWORD _MAX`). If changes are required to a data definition, consider setting the _Period_ parameter to `SIMCONNECT_PERIOD_NEVER` (see the `SIMCONNECT_PERIOD` enumeration) before making the changes, and then turning on the appropriate period after the changes have been made.

It is possible to change the period of a request, by re-sending the `SimConnect_RequestDataOnSimObject` call with the same _RequestID_, _DefineID_ and _ObjectID_ parameters, but with a new period. The one exception to this is the new period cannot be `SIMCONNECT_PERIOD_ONCE`, in this case a request with a new _RequestID_ should be sent.

Data is always transmitted with the `SimConnect_RequestDataOnSimObject` function, so if timing only notifications are required, use the `SimConnect_SubscribeToSystemEvent` function.

Note that variable length strings should not be used in data definitions, except where the _Period_ parameter has been set to `SIMCONNECT_PERIOD_ONCE`.

One method of testing whether the user has changed aircraft type is to use this function to return the title of the user aircraft, and note that if it changes, the user has changed the type of aircraft (all aircraft types have unique title strings, including those simply with different color schemes).

If boolean data has been requested as part of a data definition, note that the only reliable numeric equivalent is that 0 is returned for False. Non-zero values, especially both 1 and -1, are used to indicate True.

**NOTE**: **Multiplayer Mode** \- When developing a client for use in multiplayer mode it is not safe to use the ID number for the user aircraft returned by the function `SimConnect_RequestDataOnSimObjectType`, as the actual number can change depending on several factors, including the number of users involved in the multiplayer flight. Always use the constant `SIMCONNECT_OBJECT_ID_USER` for the _ObjectID_ parameter if the SimConnect client is to work in multiplayer mode. Also note that in this mode it is good practice to remove any calls associated with periodic data on AI objects and to remove subscriptions to AI objects.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AddToDataDefinition](SimConnect_AddToDataDefinition.htm)
4. [SimConnect\_ClearDataDefinition](SimConnect_ClearDataDefinition.htm)
5. [SimConnect\_RequestDataOnSimObjectType](SimConnect_RequestDataOnSimObjectType.htm)
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