SIMCONNECT\_RECV\_RESERVED\_KEY

## SIMCONNECT\_RECV\_RESERVED\_KEY

The **SIMCONNECT\_RECV\_RESERVED\_KEY** structure is used with the [SimConnect\_RequestReservedKey](../Events_And_Data/SimConnect_RequestReservedKey.htm) function to return the reserved key combination.

##### Syntax

```cpp
struct SIMCONNECT_RECV_RESERVED_KEY : public SIMCONNECT_RECV {
    char  szChoiceReserved[30];
    char  szReservedKey[50];
     };
```

##### Members

| Member | Description |
| `szChoiceReserved[30]` | Null-terminated string containing the key that has been reserved. This will be identical to the string entered as one of the choices for the [SimConnect\_RequestReservedKey](../Events_And_Data/SimConnect_RequestReservedKey.htm) function. |
| `szReservedKey[50]` | Null-terminated string containing the reserved key combination. This will be an uppercase string containing all the modifiers that apply. For example, if the client program requests "q", and the choice is accepted, then this parameter will contain "TAB+Q". If the client program requests "Q", then this parameter will contain "SHIFT+TAB+Q". This string could then appear, for example, in a dialog from the client application, informing a user of the appropriate help key. |

##### Remarks

This structure inherits the `SIMCONNECT_RECV` structure and is returned when the `dwID` parameter of `SIMCONNECT_RECV` is set to `SIMCONNECT_RECV_ID_RESERVED_KEY`.

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