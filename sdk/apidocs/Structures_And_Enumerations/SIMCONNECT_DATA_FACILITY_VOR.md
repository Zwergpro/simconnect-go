SIMCONNECT\_DATA\_FACILITY\_VOR

## SIMCONNECT\_DATA\_FACILITY\_VOR

The **SIMCONNECT\_DATA\_FACILITY\_VOR** structure is used to return information on a single VOR station in the facilities cache.

##### Syntax

```cpp
struct SIMCONNECT_DATA_FACILITY_VOR: public SIMCONNECT_DATA_FACILITY_NDB{
    DWORD  Flags;
    float  fLocalizer;
    double  GlideLat;
    double  GlideLon;
    double  GlideAlt;
    float  fGlideSlopeAngle;
    };
```

##### Members

| Member | Description |
| `Flags` | Flags indicating whether the other fields are valid or not.<br>- `SIMCONNECT_RECV_ID_VOR_LIST_HAS_NAV_SIGNAL`(0x1): Set if the station has a NAV transmitter, and if so, GlideLat, GlideLon and GlideAlt contain valid data.<br>- `SIMCONNECT_RECV_ID_VOR_LIST_HAS_LOCALIZER`(0x2): Set if the station transmits an ILS localizer angle, and if so fLocalizer contains valid data.<br>- `SIMCONNECT_RECV_ID_VOR_LIST_HAS_GLIDE_SLOPE`(0x4): Set if the station transmits an ILS approach angle, and if so fGlideSlopeAngle contains valid data.<br>- `SIMCONNECT_RECV_ID_VOR_LIST_HAS_DME`(0x8): Set if the station t transmits a DME signal, and if so the inherited DME fFrequency contains valid data. |
| `fLocalizer` | The ILS localizer angle in degrees. |
| `GlideLat` | The latitude of the glide slope transmitter in degrees. |
| `GlideLon` | The longitude of the glide slope transmitter in degrees. |
| `GlideAlt` | The altitude of the glide slope transmitter in degrees. |
| `fGlideSlopeAngle` | The ILS approach angle in degrees. |

##### Remarks

This structure is returned as one element in the [SIMCONNECT\_RECV\_VOR\_LIST](SIMCONNECT_RECV_VOR_LIST.htm) structure. It inherits all the members from [SIMCONNECT\_DATA\_FACILITY\_NDB](SIMCONNECT_DATA_FACILITY_NDB.htm).

See the remarks for [SimConnect\_RequestFacilitesList](../Facilities/SimConnect_RequestFacilitesList.htm).

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