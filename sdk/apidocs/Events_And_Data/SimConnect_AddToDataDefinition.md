SimConnect\_AddToDataDefinition

## SimConnect\_AddToDataDefinition

The **SimConnect\_AddToDataDefinition** function is used to add a Microsoft Flight Simulator 2024 simulation variable name to a client defined object definition.

##### Syntax

```cpp
HRESULT SimConnect_AddToDataDefinition(
    HANDLE  hSimConnect,
    SIMCONNECT_DATA_DEFINITION_ID  DefineID,
    const char*  DatumName,
    const char*  UnitsName,
    SIMCONNECT_DATATYPE  DatumType = SIMCONNECT_DATATYPE_FLOAT64,
    float  fEpsilon = 0,
    DWORD  DatumID = SIMCONNECT_UNUSED
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _DefineID_ | Specifies the ID of the client defined data definition. | Integer |
| _DatumName_ | Specifies the name of the Microsoft Flight Simulator 2024 simulation variable. See the [Simulation Variables documents](../../../SimVars/Simulation_Variables.htm) for tables of variable names. If an index is required then it should be appended to the variable name following a colon, see the example for DEFINITION\_2 below. Indexes are numbered from 1 (not zero). Simulation variable names are not case-sensitive (so can be entered in upper or lower case). | Integer |
| _UnitsName_ | Specifies the units of the variable. See [Simulation Variable Units](../../../SimVars/Simulation_Variable_Units.htm) for tables of acceptable unit names. It is possible to specify different units to receive the data in, from those specified in the Simulation Variables document. The alternative units must come under the same heading (such as Angular Velocity, or Volume, as specified in the Units of Measurement section of the Simulation Variables document). For strings and structures enter "NULL" for this parameter. | Integer |
| _DatumType_ | One member of the [SIMCONNECT\_DATATYPE](../Structures_And_Enumerations/SIMCONNECT_DATATYPE.htm) enumeration type. This parameter is used to determine what datatype should be used to return the data. The default is `SIMCONNECT_DATATYPE_FLOAT64`. Note that the structure data types are legitimate parameters here. | Integer<br>(OPTIONAL) |
| _fEpsilon_ | If data is requested only when it changes (see the _flags_ parameter of [SimConnect\_RequestDataOnSimObject](SimConnect_RequestDataOnSimObject.htm)) a change will only be reported if it is greater than the value of this parameter (not greater than or equal to). The default is zero, so even the tiniest change will initiate the transmission of data. Set this value appropriately so insignificant changes are not transmitted. This can be used with integer data, the floating point _fEpsilon_ value is first truncated to its integer component before the comparison is made (for example, an _fEpsilon_ value of 2.9 truncates to 2, so a data change of 2 will not trigger a transmission, and a change of 3 will do so). | Float<br>(OPTIONAL) |
| _DatumID_ | Specifies a client defined datum ID. The default is zero. Use this to identify the data received if the data is being returned in tagged format (see the flags parameter of [SimConnect\_RequestDataOnSimObject](SimConnect_RequestDataOnSimObject.htm). There is no need to specify datum IDs if the data is not being returned in tagged format. | Integer<br>(OPTIONAL) |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table.

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |

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
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_2, "GENERAL ENG RPM:1", "rpm");
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_2, "GENERAL ENG RPM:2", "revolutions per minute");
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_2, "GENERAL ENG RPM:3", "degrees per second");
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_2, "GENERAL ENG RPM:4", "minutes per round");
      ....
hr = SimConnect_RequestDataOnSimObject(hSimConnect, REQUEST_1, DEFINITION_1, SIMCONNECT_OBJECT_ID_USER, SIMCONNECT_PERIOD_ONCE);
      ....
    case SIMCONNECT_RECV_ID_SIMOBJECT_DATA:
        {
        SIMCONNECT_RECV_SIMOBJECT_DATA *pObjData = (SIMCONNECT_RECV_SIMOBJECT_DATA*)pData;
        switch(pObjData->dwRequestID)
            {
            case REQUEST_1:
                Struct1 *pS = (Struct1*)&pObjData->dwData;
                break;
            }
        break;
        }
```

##### Remarks

The maximum number of entries in a data definition is 1000.

##### L Vars

It is possible to get and/or set an [RPN](../../../../6_Programming_APIs/Reverse_Polish_Notation.htm) "L" variable through SimConnect using the `SimConnect_AddToDataDefinition` function, for example:

```cpp
SimConnect_AddToDataDefinition(hSimConnect, DataDefinitionID, "L:VARIABLE_NAME", "number", SIMCONNECT_DATATYPE_FLOAT64);
```

However, when working with "L" vars, it is important to note that:

- "L" variables are shared by everytheing without any priority, essentially making them "global" in scope. They can be used by audio, animation, cockpit, etc... which means that the "L" vars you get using SimConnect may be the same that are used by - for example - audio, so if you update this var, the audio will also be updated. Futhermore, if two SimConnect clients try to access the same "L" Var, the last one which updates the Var is correct and sets the new value for everything.
- It is possible to give a [unit](../../../SimVars/Simulation_Variable_Units.htm) to an "L" var, in which case the variable's type will change and a conversion will be made if necessary.
- "L" variables are _only_ FLOAT64.

##### I and O Vars

It is possible to get and/or  set [RPN](../../../Reverse_Polish_Notation.htm) "I" and "O" variables through SimConnect using the `SimConnect_AddToDataDefinition` function, for example:

```cpp
SimConnect_AddToDataDefinition(hSimConnect, DataDefinitionID, "I:Path:To:Component@alias:I_VARIABLE_NAME", "number", SIMCONNECT_DATATYPE_FLOAT64);
SimConnect_AddToDataDefinition(hSimConnect, DataDefinitionID, "O:Path:To:Component@alias:O_VARIABLE_NAME", "number", SIMCONNECT_DATATYPE_FLOAT64);

```

However, when working with "I" vars, it is important to note that:

- The `DatumName` for "I" and "O" format is :

  - The letter describing the type of the variable followed by a colon ("I:" or "O:")
  - The path to the component, with colon as separator. The path can contain an _optional_ alias prefixed by an "@" to to found the component faster. The path is then followed by a colon. ("Path:To:Component:" or "Path:To:Component@alias:")
  - The name of the variable
- "I" variables are local to an instrument and accessible to all the component of the instrument's hierarchy.
- "O" variables are local to a component and accessible only by this component
- If at an instant t, the component **isn't valid**, the value of the variable will be 0 and setting the value will have no effect. (note that during the loading of the simobject the component isn't valid)
- If the simobject change or is reloaded, the variable is reset.
- "I" and "O" variables are _only_ FLOAT64.

##### Z (or L:1) variables

It is possible to get and/or  set an [RPN](../../../Reverse_Polish_Notation.htm) "Z" (also called "L:1") variable through SimConnect using the `SimConnect_AddToDataDefinition` function, for example:

```cpp
SimConnect_AddToDataDefinition(hSimConnect, DataDefinitionID, "Z:VARIABLE_NAME", "number", SIMCONNECT_DATATYPE_FLOAT64);
SimConnect_AddToDataDefinition(hSimConnect, DataDefinitionID, "L:1:VARIABLE_NAME", "number", SIMCONNECT_DATATYPE_FLOAT64);
```

However, when working with "Z" vars, it is important to note that:

- "Z" variables are local to a simobject
- If the simobject is loading, setting the variable will have no effect, and getting its value always return 0.
- If the simobject change or is reloaded, the variable is reset.
- "Z" variables are _only_ FLOAT64

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_ClearDataDefinition](SimConnect_ClearDataDefinition.htm)
4. [SimConnect\_RequestDataOnSimObject](SimConnect_RequestDataOnSimObject.htm)
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