SimConnect\_MapInputEventToClientEvent

## SimConnect\_MapInputEventToClientEvent

**IMPORTANT!** There is a known bug with `SimConnect_MapInputEventToClientEvent` where the DeviceId could change if a device was unplugged and plugged in again. To preserve backwards compatibility, this function has been deprecated in favour of the `SimConnect_MapInputEventToClientEvent_EX1`, which does not have this same potential issue.The **SimConnect\_MapInputEventToClientEvent** function is used to connect input events (such as keystrokes, joystick or mouse movements) with the sending of appropriate event notifications.

##### Syntax

```cpp
HRESULT SimConnect_MapInputEventToClientEvent(
    HANDLE  hSimConnect,
    SIMCONNECT_INPUT_GROUP_ID  GroupID,
    const char*  pszInputDefinition,
    SIMCONNECT_CLIENT_EVENT_ID  DownEventID,
    DWORD  DownValue = 0,
    SIMCONNECT_CLIENT_EVENT_ID  UpEventID =(SIMCONNECT_CLIENT_EVENT_ID)SIMCONNECT_UNUSED,
    DWORD  UpValue = 0,
    BOOL  bMaskable = FALSE
    );
```

##### Parameters

| Parameter | Description | Type |
| --- | --- | --- |
| _hSimConnect_ | Handle to a SimConnect object. | Integer |
| _GroupID_ | Specifies the ID of the client defined input group that the input event is to be added to. | Integer |
| _pszInputDefinition_ | Pointer to a null-terminated string containing the definition of the input events (keyboard keys, mouse or joystick events, for example). See the Remarks and example below for a range of possibilities. | Integer |
| _DownEventID_ | Specifies the ID of the down, and default, event. This is the client defined event that is triggered when the input event occurs. If only an up event is required, set this to `SIMCONNECT_UNUSED`. | Integer |
| _DownValue_ | Specifies an optional numeric value, which will be returned when the down event occurs. | Integer<br>(OPTIONAL) |
| _UpEventID_ | Specifies the ID of the up event. This is the client defined event that is triggered when the up action occurs. | Integer<br>(OPTIONAL) |
| _UpValue_ | Specifies an optional numeric value, which will be returned when the up event occurs. | Integer<br>(OPTIONAL) |
| _bMaskable_ | If set to true, specifies that the client will mask the event, and no other lower priority clients will receive it. The default is false. | Bool<br>(OPTIONAL) |

##### Return Values

The function returns an **HRESULT**. Possible values include, but are not limited to, those in the following table:

| Return value | Description |
| --- | --- |
| S\_OK | The function succeeded. |
| E\_FAIL | The function failed. |
| VALUE | A numeric value that is returned when one of the optional parameters _DownValue_ or _UpValue_ have been supplied. |

##### Example

```cpp
static enum INPUT_ID {
    INPUT_1,
    };
static enum EVENT_ID {
    EVENT_1,
    EVENT_2,
    EVENT_3
    };

hr = SimConnect_MapClientEventToSimEvent(hSimConnect, EVENT_1, "parking_brakes");
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "a+B", EVENT_1);
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "VK_LCONTROL+A+U", EVENT_1);
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "VK_LSHIFT+VK_LCONTROL+a", EVENT_2);
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "VK_RSHIFT+VK_RCONTROL+a", EVENT_2, 0, EVENT_3);
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "joystick:0:button:0", EVENT_2);
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "joystick:0:button:1", EVENT_3);
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "joystick:0:POV:0", EVENT_3);
hr = SimConnect_MapInputEventToClientEvent(hSimConnect, INPUT_1, "joystick:0:XAxis", EVENT_3);
hr = SimConnect_SetInputGroupPriority(hSimConnect, INPUT_1, SIMCONNECT_GROUP_PRIORITY_HIGHEST);
```

##### Remarks

The maximum number of events that can be added to an input group is 1000. For the keyboard the input definition can include a maximum of two modifiers ( **VK\_LCONTROL**, **VK\_RCONTROL**, **VK\_LSHIFT**, **VK\_RSHIFT**, **VK\_LMENU** (Alt Left) and **VK\_RMENU** (Alt Right)) and two keys (case sensitive). For joysticks the input definition is in the form "`joystick:n:input[:i]`". Where `n` is the joystick number (starting from 0), `input` is the input name, and `i` is an optional index number that might be required by the input name (`joystick:0:button:0` for example). The input name can be one in the following table:

| Input Name | Description | Range of values |
| Button | One of the joystick buttons, configured from 0. | Not applicable |
| POV | Point of view switch (often called the hat switch). | 0 facing ahead 4500 forward right 9000 right 13500 rear right 18000 rear 22500 rear left 27000 left 31500 forward left |
| Slider | The variable position slider on the joystick. | The actual values returned can vary widely on the joystick, though the limits are 32K (pulled back to the limit) to -32K (maximum forward limit). |
| XAxis, YAxis or ZAxis | Movement of the joystick in the X, Y, or Z directions. For most joysticks the movement is left or right for the XAxis and forward or backward for the YAxis, with no values for the ZAxis. | The limits in the Y axis are 32K (pulled back) to -32K (pushed forward). The limits in the X axis are -32K (full left) to 32K (full right). Depending on the joystick though, the limits may be significantly less than these values. |
| RxAxis, RyAxis, or RzAxis | Rotation of the joystick about the X, Y, or Z axis. For most joysticks there is only rotational movement around the Z axis, with no values for the X or Y axis. | For the Z axis, the limits are -32K (rotated left to the maximum) to 32K (rotated right to the maximum). Again, actual limits depend on the joystick. |

For keyboard hits, usually no further information other than the key has been pressed is necessary for the client to process the event appropriately. For joystick events, other than button events, it is also important to know the extent of the movement (or position of the hat switch, or of the slider). This information is returned with the event in the _dwData_ parameter of a [SIMCONNECT\_RECV\_EVENT](../Structures_And_Enumerations/SIMCONNECT_RECV_EVENT.htm) structure.

For button, hat switch, or keyboard events, one event is transmitted to the client, or two if an up event is specified, when the input event occurs. If joystick axis, rotation or slider events are requested, then an event is transmitted for these six times per second whether the joystick is actually moved or not, unless the value for these is zero, in which case events are not transmitted until the joystick is moved from this position. Joystick and keyboard events are only transmitted when a flight is loaded, not while the user is navigating the shell of the product.

#### Valid Input Strings

The list below contains every valid input string that can be used:

"VK\_BACK"

"VK\_TAB"

"VK\_CLEAR"

"VK\_RETURN"

"VK\_PAUSE"

"VK\_CAPITAL"

"VK\_KANA"

"VK\_IME\_ON"

"VK\_JUNJA"

"VK\_FINAL"

"VK\_KANJI"

"VK\_IME\_OFF"

"VK\_ESCAPE"

"VK\_CONVERT"

"VK\_NONCONVERT"

"VK\_ACCEPT"

"VK\_MODECHANGE"

"VK\_SPACE"

"VK\_PRIOR"

"VK\_NEXT"

"VK\_END"

"VK\_HOME"

"VK\_LEFT"

"VK\_UP"

"VK\_RIGHT"

"VK\_DOWN"

"VK\_SELECT"

"VK\_PRINT"

"VK\_EXECUTE"

"VK\_SNAPSHOT"

"VK\_INSERT"

"VK\_DELETE"

"VK\_HELP"

"0"

"1"

"2"

"3"

"4"

"5"

"6"

"7"

"8"

"9"

"A"

"B"

"C"

"D"

"E"

"F"

"G"

"H"

"I"

"J"

"K"

"L"

"M"

"N"

"O"

"P"

"Q"

"R"

"S"

"T"

"U"

"V"

"W"

"X"

"Y"

"Z"

"VK\_LWIN"

"VK\_RWIN"

"VK\_APPS"

"VK\_SLEEP"

"VK\_NUMPAD0"

"VK\_NUMPAD1"

"VK\_NUMPAD2"

"VK\_NUMPAD3"

"VK\_NUMPAD4"

"VK\_NUMPAD5"

"VK\_NUMPAD6"

"VK\_NUMPAD7"

"VK\_NUMPAD8"

"VK\_NUMPAD9"

"VK\_MULTIPLY"

"VK\_ADD"

"VK\_SEPARATOR"

"VK\_SUBTRACT"

"VK\_DECIMAL"

"VK\_DIVIDE"

"VK\_F1"

"VK\_F2"

"VK\_F3"

"VK\_F4"

"VK\_F5"

"VK\_F6"

"VK\_F7"

"VK\_F8"

"VK\_F9"

"VK\_F10"

"VK\_F11"

"VK\_F12"

"VK\_F13"

"VK\_F14"

"VK\_F15"

"VK\_F16"

"VK\_F17"

"VK\_F18"

"VK\_F19"

"VK\_F20"

"VK\_F21"

"VK\_F22"

"VK\_F23"

"VK\_F24"

"VK\_NUMLOCK"

"VK\_SCROLL"

"VK\_OEM\_FJ\_JISHO"

"VK\_OEM\_FJ\_MASSHOU"

"VK\_OEM\_FJ\_TOUROKU"

"VK\_OEM\_FJ\_LOYA"

"VK\_OEM\_FJ\_ROYA"

"VK\_LSHIFT"

"VK\_RSHIFT"

"VK\_LCONTROL"

"VK\_RCONTROL"

"VK\_LMENU"

"VK\_RMENU"

"VK\_BROWSER\_BACK"

"VK\_BROWSER\_FORWARD"

"VK\_BROWSER\_REFRESH"

"VK\_BROWSER\_STOP"

"VK\_BROWSER\_SEARCH"

"VK\_BROWSER\_FAVORITES"

"VK\_BROWSER\_HOME"

"VK\_VOLUME\_MUTE"

"VK\_VOLUME\_DOWN"

"VK\_VOLUME\_UP"

"VK\_MEDIA\_NEXT\_TRACK"

"VK\_MEDIA\_PREV\_TRACK"

"VK\_MEDIA\_STOP"

"VK\_MEDIA\_PLAY\_PAUSE"

"VK\_LAUNCH\_MAIL"

"VK\_LAUNCH\_MEDIA\_SELECT"

"VK\_LAUNCH\_APP1"

"VK\_LAUNCH\_APP2"

"VK\_SEMICOLON"

"VK\_PLUS"

"VK\_COMMA"

"VK\_MINUS"

"VK\_PERIOD"

"VK\_SLASH"

"VK\_TILDE"

"VK\_LBRACKET"

"VK\_BACKSLASH"

"VK\_RBRACKET"

"VK\_QUOTE"

"VK\_0xDF"

"VK\_0xE0"

"VK\_OEM\_AX"

"VK\_OEM\_102"

"VK\_ICO\_HELP"

"VK\_ICO\_00"

"VK\_PROCESSKEY"

"VK\_ICO\_CLEAR"

"VK\_PACKET"

"VK\_0xE8"

"VK\_OEM\_RESET"

"VK\_OEM\_JUMP"

"VK\_OEM\_PA1"

"VK\_OEM\_PA2"

"VK\_OEM\_PA3"

"VK\_OEM\_WSCTRL"

"VK\_OEM\_CUSEL"

"VK\_OEM\_ATTN"

"VK\_OEM\_FINISH"

"VK\_OEM\_COPY"

"VK\_OEM\_AUTO"

"VK\_OEM\_ENLW"

"VK\_OEM\_BACKTAB"

"VK\_ATTN"

"VK\_CRSEL"

"VK\_EXSEL"

"VK\_EREOF"

"VK\_PLAY"

"VK\_ZOOM"

"VK\_PA1"

"VK\_OEM\_CLEAR"

Related Topics

1. [SimConnect API Reference](../../SimConnect_API_Reference.htm)
2. [SimConnect\_SetInputGroupPriority](../Events_And_Data/SimConnect_SetInputGroupPriority.htm)
3. [SimConnect\_RemoveInputEvent](../Events_And_Data/SimConnect_RemoveInputEvent.htm)
4. [SimConnect\_ClearInputGroup](../Events_And_Data/SimConnect_ClearInputGroup.htm)
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