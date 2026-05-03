SIMCONNECT\_POSITION\_REFERENTIAL

## SIMCONNECT\_POSITION\_REFERENTIAL

The **SIMCONNECT\_POSITION\_REFERENTIAL** enum is used to define what reference is being used for the camera position.

##### Syntax

```cpp
enum SIMCONNECT_POSITION_REFERENTIAL
{
    SIMCONNECT_POSITION_REFERENTIAL_NONE,
    SIMCONNECT_POSITION_REFERENTIAL_AIRCRAFT,
    SIMCONNECT_POSITION_REFERENTIAL_WORLD,
    SIMCONNECT_POSITION_REFERENTIAL_EYEPOINT
};
```

##### Members

| Member | Description |
| `SIMCONNECT_POSITION_REFERENTIAL_NONE` | No reference (this should not be used). |
| `SIMCONNECT_POSITION_REFERENTIAL_AIRCRAFT` | The position is an offset from the aircraft [Datum Reference Point](../../../../5_Content_Configuration/Modular_SimObjects/Aircraft/Aircraft.htm#DRP) (expressed in meters). |
| `SIMCONNECT_POSITION_REFERENTIAL_WORLD` | The position is expressed using latitude, longitude, and altitude. |
| `SIMCONNECT_POSITION_REFERENTIAL_EYEPOINT` | The position is an offset from the aircraft pilot eyepoint (expressed in meters). |

##### Remarks

This enum will be used by the `SimConnect_CameraGet` and other camera functions.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect Samples](../../../../7_Samples_Tutorials/Samples/VisualStudio/SimConnect_Samples.htm)
4. [Camera](../Camera/Camera.htm)

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