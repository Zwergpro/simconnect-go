SIMCONNECT\_CONTROLLER\_ITEM

## SIMCONNECT\_CONTROLLER\_ITEM

The **SIMCONNECT\_CONTROLLER\_ITEM** struct contains data related to a single controller currently connected to the simulation.

##### Syntax

```cpp
struct SIMCONNECT_CONTROLLER_ITEM{
    SIMCONNECT_STRING(DeviceName, 256);
    unsigned int DeviceId;
    unsigned int ProductId;
    unsigned int CompositeID;
    SIMCONNECT_VERSION_BASE_TYPE HardwareVersion;
};
```

##### Members

| Member | Description |
| `DeviceName` | A string that gives the descriptive name for the device. |
| `DeviceId` | The device ID. |
| `ProductId` | The product ID. |
| `CompositeID` | ID of the USB composite device (for when devices have the same `ProductId`, but there are multiple recognised parts on the same device) |
| `HardwareVersion` | The version data for the hardware, returned as a `SIMCONNECT_VERSION_BASE_TYPE` struct. |

##### Remarks

See `SimConnect_EnumerateControllers` for more information.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_EnumerateControllers](../InputEvents/SimConnect_EnumerateControllers.htm)
4. [SIMCONNECT\_RECV\_CONTROLLERS\_LIST](SIMCONNECT_RECV_CONTROLLERS_LIST.htm)
5. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)

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