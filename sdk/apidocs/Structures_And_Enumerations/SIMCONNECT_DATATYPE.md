SIMCONNECT\_DATATYPE

## SIMCONNECT\_DATATYPE

The **SIMCONNECT\_DATATYPE** enumeration type is used with the [SimConnect\_AddToDataDefinition](../Events_And_Data/SimConnect_AddToDataDefinition.htm) call to specify the data type that the server should use to return the specified data to the client.

##### Syntax

```cpp
enum SIMCONNECT_DATATYPE{
    SIMCONNECT_DATATYPE_INVALID,
    SIMCONNECT_DATATYPE_INT32,
    SIMCONNECT_DATATYPE_INT64,
    SIMCONNECT_DATATYPE_FLOAT32,
    SIMCONNECT_DATATYPE_FLOAT64,
    SIMCONNECT_DATATYPE_STRING8,
    SIMCONNECT_DATATYPE_STRING32,
    SIMCONNECT_DATATYPE_STRING64,
    SIMCONNECT_DATATYPE_STRING128,
    SIMCONNECT_DATATYPE_STRING256,
    SIMCONNECT_DATATYPE_STRING260,
    SIMCONNECT_DATATYPE_STRINGV,
    SIMCONNECT_DATATYPE_INITPOSITION,
    SIMCONNECT_DATATYPE_MARKERSTATE,
    SIMCONNECT_DATATYPE_WAYPOINT,
    SIMCONNECT_DATATYPE_LATLONALT,
    SIMCONNECT_DATATYPE_XYZ
    };
```

##### Members

| Member | Description |
| `SIMCONNECT_DATATYPE_FLOAT32` | Specifies a 32 bit or 64 bit floating point number. |
| `SIMCONNECT_DATATYPE_FLOAT64` |
| `SIMCONNECT_DATATYPE_INT32` | Specifies a 32 bit _unsigned_ integer value. |
| `SIMCONNECT_DATATYPE_INT64` | Specifies a 64 bit _signed_ integer value. |
| `SIMCONNECT_DATATYPE_STRING8` | Specifies strings of the given length (8 characters to 260 characters) |
| `SIMCONNECT_DATATYPE_STRING32` |
| `SIMCONNECT_DATATYPE_STRING64` |
| `SIMCONNECT_DATATYPE_STRING128` |
| `SIMCONNECT_DATATYPE_STRING256` |
| `SIMCONNECT_DATATYPE_STRING260` |
| `SIMCONNECT_DATATYPE_STRINGV` | Specifies a variable length string. |
| `SIMCONNECT_DATATYPE_INITPOSITION` | Specifies the `SIMCONNECT_DATA_INITPOSITION` structure. |
| `SIMCONNECT_DATATYPE_MARKERSTATE` | Specifies the `SIMCONNECT_DATA_MARKERSTATE` structure. |
| `SIMCONNECT_DATATYPE_WAYPOINT` | Specifies the `SIMCONNECT_DATA_WAYPOINT` structure. |
| `SIMCONNECT_DATATYPE_LATLONALT` | Specifies the `SIMCONNECT_DATA_LATLONALT` structure. |
| `SIMCONNECT_DATATYPE_XYZ` | Specifies the `SIMCONNECT_DATA_XYZ` structure. |

##### Remarks

The three structures in the list of data types can only be used as input (using [SimConnect\_SetDataOnSimObject](../Events_And_Data/SimConnect_SetDataOnSimObject.htm)) and not to receive requested data.

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