SIMCONNECT\_CAMERA\_DATA\_MASK

## SIMCONNECT\_CAMERA\_DATA\_MASK

The **SIMCONNECT\_CAMERA\_DATA\_MASK** enum used to create a bitmask for use with the `SimConnect_CameraSet` function, telling the simulation which camera struct members are being set.

##### Syntax

```cpp
enum SIMCONNECT_CAMERA_DATA_MASK {
    SIMCONNECT_CAMERA_DATA_MASK_NONE         = 0x00;
    SIMCONNECT_CAMERA_DATA_MASK_POSITION     = 0x01;
    SIMCONNECT_CAMERA_DATA_MASK_ROTATION     = 0x02;
    SIMCONNECT_CAMERA_DATA_MASK_TARGETED     = 0x04;
    SIMCONNECT_CAMERA_DATA_MASK_FOV          = 0x08;
    SIMCONNECT_CAMERA_DATA_MASK_REFERENTIAL  = 0x10;
    SIMCONNECT_CAMERA_DATA_MASK_ALL_ROTATION = SIMCONNECT_CAMERA_DATA_MASK_POSITION | SIMCONNECT_CAMERA_DATA_MASK_ROTATION | SIMCONNECT_CAMERA_DATA_MASK_FOV | SIMCONNECT_CAMERA_DATA_MASK_REFERENTIAL;
    SIMCONNECT_CAMERA_DATA_MASK_ALL_TARGETED = SIMCONNECT_CAMERA_DATA_MASK_POSITION | SIMCONNECT_CAMERA_DATA_MASK_TARGETED | SIMCONNECT_CAMERA_DATA_MASK_FOV | SIMCONNECT_CAMERA_DATA_MASK_REFERENTIAL;
    }
```

##### Members

| Member | Description |
| `SIMCONNECT_CAMERA_DATA_MASK_NONE` | No data mask is applied (not used). |
| `SIMCONNECT_CAMERA_DATA_MASK_POSITION` | The camera position flag (position is either in x/y/z/ meters or lat/lon/alt, depending on the referential used). |
| `SIMCONNECT_CAMERA_DATA_MASK_ROTATION` | The camera rotation flag (rotation is given as pitch/bank/heading in degrees). Note that if this is used then `SIMCONNECT_CAMERA_DATA_MASK_TARGETED` should not be used, otherwise `SIMCONNECT_CAMERA_DATA_MASK_TARGETED` takes priority and will overwrite the pitch/bank/heading values. |
| `SIMCONNECT_CAMERA_DATA_MASK_TARGETED` | The camera "target" position flag (target position is either in x/y/z/ meters or lat/lon/alt, depending on the referential used). The camera rotation will be calculated as the direction/orientation between the _position_ and the _target position_. Note that if this is used then `SIMCONNECT_CAMERA_DATA_MASK_ROTATION` is _not_ used, and any values given for this will override the pitch/bank/heading rotation values even if they are included. |
| `SIMCONNECT_CAMERA_DATA_MASK_FOV` | The camera field of view flag (the value used for FOV is _radians_). |
| `SIMCONNECT_CAMERA_DATA_MASK_REFERENTIAL` | The camera referential flag. |
| `SIMCONNECT_CAMERA_DATA_MASK_ALL_ROTATION` | The "all parameters" flag using _rotation_ to set the camera "look at" direction/orientation (ie: get/set all the above parameters, _except_`SIMCONNECT_CAMERA_DATA_MASK_TARGETED`). |
| `SIMCONNECT_CAMERA_DATA_MASK_ALL_TARGETED` | The "all parameters" flag using _target position_ to set the camera "look at" direction/orientation (ie: get/set all the above parameters, _except_`SIMCONNECT_CAMERA_DATA_MASK_ALL_ROTATION`). |

##### Remarks

N/A

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