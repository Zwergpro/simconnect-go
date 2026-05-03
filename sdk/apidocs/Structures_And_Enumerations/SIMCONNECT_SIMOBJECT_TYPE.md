SIMCONNECT\_SIMOBJECT\_TYPE

## SIMCONNECT\_SIMOBJECT\_TYPE

The **SIMCONNECT\_SIMOBJECT\_TYPE** enumeration type is used with the [SimConnect\_RequestDataOnSimObjectType](../Events_And_Data/SimConnect_RequestDataOnSimObjectType.htm) call to request information on specific or nearby objects.

##### Syntax

```cpp
enum SIMCONNECT_SIMOBJECT_TYPE{
    SIMCONNECT_SIMOBJECT_TYPE_USER = 0,
   SIMCONNECT_SIMOBJECT_TYPE_USER_AIRCRAFT = 0,
  SIMCONNECT_SIMOBJECT_TYPE_ALL,
  SIMCONNECT_SIMOBJECT_TYPE_AIRCRAFT,
  SIMCONNECT_SIMOBJECT_TYPE_HELICOPTER,
  SIMCONNECT_SIMOBJECT_TYPE_BOAT,
  SIMCONNECT_SIMOBJECT_TYPE_GROUND,
  SIMCONNECT_SIMOBJECT_TYPE_HOT_AIR_BALLOON,
  SIMCONNECT_SIMOBJECT_TYPE_ANIMAL,
  SIMCONNECT_SIMOBJECT_TYPE_USER_AVATAR,
  SIMCONNECT_SIMOBJECT_TYPE_USER_CURRENT,
    };
```

##### Members

| Member | Description |
| `SIMCONNECT_SIMOBJECT_TYPE_USER` | Specifies the user aircraft (this is a legacy value, as its naming is ambiguous as to whether the user is the avatar or the aircraft). |
| `SIMCONNECT_SIMOBJECT_TYPE_USER_AIRCRAFT` | Specifies the user aircraft. |
| `SIMCONNECT_SIMOBJECT_TYPE_ALL` | Specifies all AI controlled objects. |
| `SIMCONNECT_SIMOBJECT_TYPE_AIRCRAFT` | Specifies all aircraft. |
| `SIMCONNECT_SIMOBJECT_TYPE_HELICOPTER` | Specifies all helicopters. |
| `SIMCONNECT_SIMOBJECT_TYPE_BOAT` | Specifies all AI controlled boats. |
| `SIMCONNECT_SIMOBJECT_TYPE_GROUND` | Specifies all AI controlled ground vehicles. |
| `SIMCONNECT_SIMOBJECT_TYPE_HOT_AIR_BALLOON` | Specifies all hot air balloons. |
| `SIMCONNECT_SIMOBJECT_TYPE_ANIMAL` | Specifies all animals. |
| `SIMCONNECT_SIMOBJECT_TYPE_USER_AVATAR` | Specifies the user avatar. |
| `SIMCONNECT_SIMOBJECT_TYPE_USER_CURRENT` | Specifies the user, regardless of whether it is the aircraft or avatar. |

##### Remarks

This enum is used to specify the return of the object IDs of all objects created using the AI creation functions, whether they are created by this client, other clients, or Microsoft Flight Simulator 2024 itself. It can also be used to specify the return the object ID of the user aircraft. However it cannot be used to specify the IDs of objects like cars moving on freeways, which are not controlled by the AI component.

See the remarks and examples for [SimConnect\_AddToDataDefinition](../Events_And_Data/SimConnect_AddToDataDefinition.htm).

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