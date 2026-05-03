SIMCONNECT\_JETWAY\_DATA

## SIMCONNECT\_JETWAY\_DATA

The **SIMCONNECT\_JETWAY\_DATA** structure is used to return information on a single jetway.

##### Syntax

```cpp
struct SIMCONNECT_JETWAY_DATA : public SIMCONNECT_RECV
{
    char AirportIcao[8];
    int ParkingIndex;
    SIMCONNECT_DATA_LATLONALT Lla;
    SIMCONNECT_DATA_PBH Pbh;
    int Status;
    int Door;
    SIMCONNECT_DATA_XYZ ExitDoorRelativePos;    //
    SIMCONNECT_DATA_XYZ MainHandlePos;          //
    SIMCONNECT_DATA_XYZ SecondaryHandle;        //
    SIMCONNECT_DATA_XYZ WheelGroundLock;        //
    DWORD JetwayObjectId;                       //
    DWORD AttachedObjectId;                     //
}
```

##### Members

| Member | Description |
| `AirportIcao` | ICAO code of the airport (will be the same as the one you used to make the request). |
| `ParkingIndex` | Index of the parking space linked to this jetway (will be the same as the one you used to make the request). |
| `Lla` | Lattitude / Longitude / Altitude of the jetway, returned as a [`SIMCONNECT_DATA_LATLONALT`](SIMCONNECT_DATA_LATLONALT.htm) struct. |
| `Pbh` | Pitch / Bank / Heading of the jetway, returned as a [`SIMCONNECT_DATA_PBH`](SIMCONNECT_DATA_PBH.htm) struct. |
| `Status` | The status of the jetway, which will be one of the following:<br>1. 0: JETWAY\_STATUS\_REST<br>2. 1: JETWAY\_STATUS\_APPROACH\_OUTSIDE<br>3. 2: JETWAY\_STATUS\_APPROACH\_DOOR<br>4. 3: JETWAY\_STATUS\_HOOD\_CONNECT<br>5. 4: JETWAY\_STATUS\_HOOD\_DISCONNECT<br>6. 5: JETWAY\_STATUS\_RETRACT\_OUTSIDE <br>7. 6: JETWAY\_STATUS\_RETRACT\_HOME<br>8. 7: JETWAY\_STATUS\_FULLY\_ATTACHED |
| `Door` | The index of the door attached to the jetway. |
| `ExitDoorRelativePos` | Door position relative to aircraft, returned as a [`SIMCONNECT_DATA_XYZ`](SIMCONNECT_DATA_XYZ.htm) struct. |
| `MainHandlePos` | Relative position of IK\_MainHandle (world pos, in meters), returned as a [`SIMCONNECT_DATA_XYZ`](SIMCONNECT_DATA_XYZ.htm) struct. |
| `SecondaryHandle` | Relative position of IK\_SecondaryHandle (world pos, in meters), returned as a [`SIMCONNECT_DATA_XYZ`](SIMCONNECT_DATA_XYZ.htm) struct. |
| `WheelGroundLock` | Relative position of IK\_WheelsGroundLock (world pos, in meters), returned as a [`SIMCONNECT_DATA_XYZ`](SIMCONNECT_DATA_XYZ.htm) struct |
| `JetwayObjectId` | ObjectId of the jetway (used by [`SimConnect_RequestDataOnSimObject`](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm) for example) |
| `AttachedObjectId` | ObjectId of the object (aircraft) attached to the jetway (used by [`SimConnect_RequestDataOnSimObject`](../Events_And_Data/SimConnect_RequestDataOnSimObject.htm) for example) |

##### Remarks

This structure is returned as one element in the [`SIMCONNECT_RECV_JETWAY_DATA`](SIMCONNECT_RECV_JETWAY_DATA.htm) structure.

Related Topics

1. [SimConnect SDK](../../SimConnect_SDK.htm)
2. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
3. [SimConnect\_RequestJetwayData](../Facilities/SimConnect_RequestJetwayData.htm)
4. [SIMCONNECT\_RECV\_JETWAY\_DATA](SIMCONNECT_RECV_JETWAY_DATA.htm)
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