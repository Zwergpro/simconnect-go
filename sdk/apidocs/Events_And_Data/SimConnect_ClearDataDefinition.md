SimConnect\_ClearDataDefinition

## SimConnect\_ClearDataDefinition

The **SimConnect\_ClearDataDefinition** function is used to remove all simulation variables from a client defined data definition.

##### Syntax

```cpp
HRESULT SimConnect_ClearDataDefinition(
    HANDLE  hSimConnect,
    SIMCONNECT_DATA_DEFINITION_ID  DefineID
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _DefineID_ | Specifies the ID of the client defined data definition. | Integer |

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
      ....
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_1, "Kohlsman setting hg", "inHg");
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_1, "Indicated Altitude", "feet");
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_1, "Plane Latitude", "degrees");
hr = SimConnect_AddToDataDefinition(hSimConnect, DEFINITION_1, "Plane Longitude", "degrees");
      ....
hr = SimConnect_ClearDataDefinition(hSimConnect, DEFINITION_1);
      ....
```

##### Remarks

Use this function to permanently delete a data definition. To temporarily suspend data requests see the remarks for the [SimConnect\_RequestDataOnSimObject](SimConnect_RequestDataOnSimObject.htm) function.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_AddToDataDefinition](SimConnect_AddToDataDefinition.htm)
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