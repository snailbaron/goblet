package sdl

import "math"

type InitFlags uint32

const (
	INIT_AUDIO    InitFlags = 0x00000010 // SDL_INIT_AUDIO implies SDL_INIT_EVENTS
	INIT_VIDEO    InitFlags = 0x00000020 // SDL_INIT_VIDEO implies SDL_INIT_EVENTS, should be initialized on the main thread
	INIT_JOYSTICK InitFlags = 0x00000200 // SDL_INIT_JOYSTICK implies SDL_INIT_EVENTS
	INIT_HAPTIC   InitFlags = 0x00001000
	INIT_GAMEPAD  InitFlags = 0x00002000 // SDL_INIT_GAMEPAD implies SDL_INIT_JOYSTICK
	INIT_EVENTS   InitFlags = 0x00004000
	INIT_SENSOR   InitFlags = 0x00008000 // SDL_INIT_SENSOR implies SDL_INIT_EVENTS
	INIT_CAMERA   InitFlags = 0x00010000 // SDL_INIT_CAMERA implies SDL_INIT_EVENTS
)

type WindowFlags uint64

const (
	WINDOW_FULLSCREEN          WindowFlags = 0x0000000000000001 // window is in fullscreen mode
	WINDOW_OPENGL              WindowFlags = 0x0000000000000002 // window usable with OpenGL context
	WINDOW_OCCLUDED            WindowFlags = 0x0000000000000004 // window is occluded
	WINDOW_HIDDEN              WindowFlags = 0x0000000000000008 // window is neither mapped onto the desktop nor shown in the taskbar/dock/window list; SDL_ShowWindow( is required for it to become visible
	WINDOW_BORDERLESS          WindowFlags = 0x0000000000000010 // no window decoration
	WINDOW_RESIZABLE           WindowFlags = 0x0000000000000020 // window can be resized
	WINDOW_MINIMIZED           WindowFlags = 0x0000000000000040 // window is minimized
	WINDOW_MAXIMIZED           WindowFlags = 0x0000000000000080 // window is maximized
	WINDOW_MOUSE_GRABBED       WindowFlags = 0x0000000000000100 // window has grabbed mouse input
	WINDOW_INPUT_FOCUS         WindowFlags = 0x0000000000000200 // window has input focus
	WINDOW_MOUSE_FOCUS         WindowFlags = 0x0000000000000400 // window has mouse focus
	WINDOW_EXTERNAL            WindowFlags = 0x0000000000000800 // window not created by SDL
	WINDOW_MODAL               WindowFlags = 0x0000000000001000 // window is modal
	WINDOW_HIGH_PIXEL_DENSITY  WindowFlags = 0x0000000000002000 // window uses high pixel density back buffer if possible
	WINDOW_MOUSE_CAPTURE       WindowFlags = 0x0000000000004000 // window has mouse captured (unrelated to MOUSE_GRABBED
	WINDOW_MOUSE_RELATIVE_MODE WindowFlags = 0x0000000000008000 // window has relative mode enabled
	WINDOW_ALWAYS_ON_TOP       WindowFlags = 0x0000000000010000 // window should always be above others
	WINDOW_UTILITY             WindowFlags = 0x0000000000020000 // window should be treated as a utility window, not showing in the task bar and window list
	WINDOW_TOOLTIP             WindowFlags = 0x0000000000040000 // window should be treated as a tooltip and does not get mouse or keyboard focus, requires a parent window
	WINDOW_POPUP_MENU          WindowFlags = 0x0000000000080000 // window should be treated as a popup menu, requires a parent window
	WINDOW_KEYBOARD_GRABBED    WindowFlags = 0x0000000000100000 // window has grabbed keyboard input
	WINDOW_FILL_DOCUMENT       WindowFlags = 0x0000000000200000 // window is in fill-document mode (Emscripten only, since SDL 3.4.0
	WINDOW_VULKAN              WindowFlags = 0x0000000010000000 // window usable for Vulkan surface
	WINDOW_METAL               WindowFlags = 0x0000000020000000 // window usable for Metal view
	WINDOW_TRANSPARENT         WindowFlags = 0x0000000040000000 // window with transparent buffer
	WINDOW_NOT_FOCUSABLE       WindowFlags = 0x0000000080000000 // window should not be focusable
)

type EventType uint32

// Application events

const (
	EVENT_QUIT EventType = 0x100 + iota // User-requested quit

	// These application events have special meaning on iOS and Android, see
	// README-ios.md and README-android.md for details.

	// The application is being terminated by the OS. This event must be handled in
	// a callback set with SDL_AddEventWatch().  Called on iOS in
	// applicationWillTerminate() Called on Android in onDestroy().
	EVENT_TERMINATING

	// The application is low on memory, free memory if possible. This event must
	// be handled in a callback set with SDL_AddEventWatch().  Called on iOS in
	// applicationDidReceiveMemoryWarning() Called on Android in onTrimMemory().
	EVENT_LOW_MEMORY

	// The application is about to enter the background. This event must be handled
	// in a callback set with SDL_AddEventWatch().  Called on iOS in
	// applicationWillResignActive() Called on Android in onPause().
	EVENT_WILL_ENTER_BACKGROUND

	// The application did enter the background and may not get CPU for some time.
	// This event must be handled in a callback set with SDL_AddEventWatch().
	// Called on iOS in applicationDidEnterBackground() Called on Android in
	// onPause().
	EVENT_DID_ENTER_BACKGROUND

	// The application is about to enter the foreground. This event must be handled
	// in a callback set with SDL_AddEventWatch().  Called on iOS in
	// applicationWillEnterForeground() Called on Android in onResume().
	EVENT_WILL_ENTER_FOREGROUND

	// The application is now interactive. This event must be handled in a callback
	// set with SDL_AddEventWatch().  Called on iOS in applicationDidBecomeActive()
	// Called on Android in onResume().
	EVENT_DID_ENTER_FOREGROUND

	EVENT_LOCALE_CHANGED // The user's locale preferences have changed.

	EVENT_SYSTEM_THEME_CHANGED // The system theme changed
)

// Display events

const (
	// 0x150 was SDL_DISPLAYEVENT, reserve the number for sdl2-compat.

	EVENT_DISPLAY_ORIENTATION           EventType = 0x151 + iota // Display orientation has changed to data1
	EVENT_DISPLAY_ADDED                                          // Display has been added to the system
	EVENT_DISPLAY_REMOVED                                        // Display has been removed from the system
	EVENT_DISPLAY_MOVED                                          // Display has changed position
	EVENT_DISPLAY_DESKTOP_MODE_CHANGED                           // Display has changed desktop mode
	EVENT_DISPLAY_CURRENT_MODE_CHANGED                           // Display has changed current mode
	EVENT_DISPLAY_CONTENT_SCALE_CHANGED                          // Display has changed content scale
	EVENT_DISPLAY_USABLE_BOUNDS_CHANGED                          // Display has changed usable bounds
	EVENT_DISPLAY_FIRST                 = EVENT_DISPLAY_ORIENTATION
	EVENT_DISPLAY_LAST                  = EVENT_DISPLAY_USABLE_BOUNDS_CHANGED
)

// Window events

const (
	// 0x200 was SDL_WINDOWEVENT, reserve the number for sdl2-compat
	// 0x201 was SDL_SYSWMEVENT, reserve the number for sdl2-compat.

	EVENT_WINDOW_SHOWN  EventType = 0x202 + iota // Window has been shown
	EVENT_WINDOW_HIDDEN                          // Window has been hidden
	// Window has been exposed and should be redrawn, and can be redrawn directly from event watchers for this event. data1 is 1 for live-resize expose events, 0 otherwise.
	EVENT_WINDOW_EXPOSED
	EVENT_WINDOW_MOVED                 // Window has been moved to data1, data2
	EVENT_WINDOW_RESIZED               // Window has been resized to data1xdata2
	EVENT_WINDOW_PIXEL_SIZE_CHANGED    // The pixel size of the window has changed to data1xdata2
	EVENT_WINDOW_METAL_VIEW_RESIZED    // The pixel size of a Metal view associated with the window has changed
	EVENT_WINDOW_MINIMIZED             // Window has been minimized
	EVENT_WINDOW_MAXIMIZED             // Window has been maximized
	EVENT_WINDOW_RESTORED              // Window has been restored to normal size and position
	EVENT_WINDOW_MOUSE_ENTER           // Window has gained mouse focus
	EVENT_WINDOW_MOUSE_LEAVE           // Window has lost mouse focus
	EVENT_WINDOW_FOCUS_GAINED          // Window has gained keyboard focus
	EVENT_WINDOW_FOCUS_LOST            // Window has lost keyboard focus
	EVENT_WINDOW_CLOSE_REQUESTED       // The window manager requests that the window be closed
	EVENT_WINDOW_HIT_TEST              // Window had a hit test that wasn't SDL_HITTEST_NORMAL
	EVENT_WINDOW_ICCPROF_CHANGED       // The window's ICC profile has changed
	EVENT_WINDOW_DISPLAY_CHANGED       // Window has been moved to display data1
	EVENT_WINDOW_DISPLAY_SCALE_CHANGED // Window display scale has been changed
	EVENT_WINDOW_SAFE_AREA_CHANGED     // The window safe area has been changed
	EVENT_WINDOW_OCCLUDED              // The window has been occluded
	EVENT_WINDOW_ENTER_FULLSCREEN      // The window has entered fullscreen mode
	EVENT_WINDOW_LEAVE_FULLSCREEN      // The window has left fullscreen mode
	// The window with the associated ID is being or has been destroyed. If this
	// message is being handled in an event watcher, the window handle is still
	// valid and can still be used to retrieve any properties associated with the
	// window. Otherwise, the handle has already been destroyed and all resources
	// associated with it are invalid.
	EVENT_WINDOW_DESTROYED
	EVENT_WINDOW_HDR_STATE_CHANGED // Window HDR properties have changed
	EVENT_WINDOW_SETTINGS_CHANGED  // Window settings have changed (on visionOS)
	EVENT_WINDOW_FIRST             = EVENT_WINDOW_SHOWN
	EVENT_WINDOW_LAST              = EVENT_WINDOW_SETTINGS_CHANGED
)

// Keyboard events

const (
	EVENT_KEY_DOWN                EventType = 0x300 + iota // Key pressed
	EVENT_KEY_UP                                           // Key released
	EVENT_TEXT_EDITING                                     // Keyboard text editing (composition)
	EVENT_TEXT_INPUT                                       // Keyboard text input
	EVENT_KEYMAP_CHANGED                                   // Keymap changed due to a system event such as an input language or keyboard layout change.
	EVENT_KEYBOARD_ADDED                                   // A new keyboard has been inserted into the system
	EVENT_KEYBOARD_REMOVED                                 // A keyboard has been removed
	EVENT_TEXT_EDITING_CANDIDATES                          // Keyboard text editing candidates
	EVENT_SCREEN_KEYBOARD_SHOWN                            // The on-screen keyboard has been shown
	EVENT_SCREEN_KEYBOARD_HIDDEN                           // The on-screen keyboard has been hidden
)

// Mouse events

const (
	EVENT_MOUSE_MOTION      EventType = 0x400 + iota // Mouse moved
	EVENT_MOUSE_BUTTON_DOWN                          // Mouse button pressed
	EVENT_MOUSE_BUTTON_UP                            // Mouse button released
	EVENT_MOUSE_WHEEL                                // Mouse wheel motion
	EVENT_MOUSE_ADDED                                // A new mouse has been inserted into the system
	EVENT_MOUSE_REMOVED                              // A mouse has been removed
)

// Joystick events

const (
	EVENT_JOYSTICK_AXIS_MOTION     EventType = 0x600 + iota // Joystick axis motion
	EVENT_JOYSTICK_BALL_MOTION                              // Joystick trackball motion
	EVENT_JOYSTICK_HAT_MOTION                               // Joystick hat position change
	EVENT_JOYSTICK_BUTTON_DOWN                              // Joystick button pressed
	EVENT_JOYSTICK_BUTTON_UP                                // Joystick button released
	EVENT_JOYSTICK_ADDED                                    // A new joystick has been inserted into the system
	EVENT_JOYSTICK_REMOVED                                  // An opened joystick has been removed
	EVENT_JOYSTICK_BATTERY_UPDATED                          // Joystick battery level change
	EVENT_JOYSTICK_UPDATE_COMPLETE                          // Joystick update is complete
)

// Gamepad events

const (
	EVENT_GAMEPAD_AXIS_MOTION          EventType = 0x650 + iota // Gamepad axis motion
	EVENT_GAMEPAD_BUTTON_DOWN                                   // Gamepad button pressed
	EVENT_GAMEPAD_BUTTON_UP                                     // Gamepad button released
	EVENT_GAMEPAD_ADDED                                         // A new gamepad has been inserted into the system
	EVENT_GAMEPAD_REMOVED                                       // A gamepad has been removed
	EVENT_GAMEPAD_REMAPPED                                      // The gamepad mapping was updated
	EVENT_GAMEPAD_TOUCHPAD_DOWN                                 // Gamepad touchpad was touched
	EVENT_GAMEPAD_TOUCHPAD_MOTION                               // Gamepad touchpad finger was moved
	EVENT_GAMEPAD_TOUCHPAD_UP                                   // Gamepad touchpad finger was lifted
	EVENT_GAMEPAD_SENSOR_UPDATE                                 // Gamepad sensor was updated
	EVENT_GAMEPAD_UPDATE_COMPLETE                               // Gamepad update is complete
	EVENT_GAMEPAD_STEAM_HANDLE_UPDATED                          // Gamepad Steam handle has changed
	EVENT_GAMEPAD_CAPSENSE_TOUCH                                // Gamepad capsense was touched
	EVENT_GAMEPAD_CAPSENSE_RELEASE                              // Gamepad capsense was released
)

// Touch events

const (
	EVENT_FINGER_DOWN EventType = 0x700 + iota
	EVENT_FINGER_UP
	EVENT_FINGER_MOTION
	EVENT_FINGER_CANCELED
)

// Pinch events

const (
	EVENT_PINCH_BEGIN  EventType = 0x710 + iota // Pinch gesture started
	EVENT_PINCH_UPDATE                          // Pinch gesture updated
	EVENT_PINCH_END                             // Pinch gesture ended
)

// 0x800, 0x801, and 0x802 were the Gesture events from SDL2. Do not reuse these values! sdl2-compat needs them!

// Clipboard events

const (
	EVENT_CLIPBOARD_UPDATE EventType = 0x900 + iota // The clipboard changed
)

// Drag and drop events

const (
	EVENT_DROP_FILE     EventType = 0x1000 + iota // The system requests a file open
	EVENT_DROP_TEXT                               // text/plain drag-and-drop event
	EVENT_DROP_BEGIN                              // A new set of drops is beginning (NULL filename)
	EVENT_DROP_COMPLETE                           // Current set of drops is now complete (NULL filename)
	EVENT_DROP_POSITION                           // Position while moving over the window
)

// Audio hotplug events

const (
	EVENT_AUDIO_DEVICE_ADDED          EventType = 0x1100 + iota // A new audio device is available
	EVENT_AUDIO_DEVICE_REMOVED                                  // An audio device has been removed.
	EVENT_AUDIO_DEVICE_FORMAT_CHANGED                           // An audio device's format has been changed by the system.
)

// Sensor events

const (
	EVENT_SENSOR_UPDATE EventType = 0x1200 + iota // A sensor was updated
)

// Pressure-sensitive pen events

const (
	EVENT_PEN_PROXIMITY_IN  EventType = 0x1300 + iota // Pressure-sensitive pen has become available
	EVENT_PEN_PROXIMITY_OUT                           // Pressure-sensitive pen has become unavailable
	EVENT_PEN_DOWN                                    // Pressure-sensitive pen touched drawing surface
	EVENT_PEN_UP                                      // Pressure-sensitive pen stopped touching drawing surface
	EVENT_PEN_BUTTON_DOWN                             // Pressure-sensitive pen button pressed
	EVENT_PEN_BUTTON_UP                               // Pressure-sensitive pen button released
	EVENT_PEN_MOTION                                  // Pressure-sensitive pen is moving on the tablet
	EVENT_PEN_AXIS                                    // Pressure-sensitive pen angle/pressure/etc changed
)

// Camera hotplug events

const (
	EVENT_CAMERA_DEVICE_ADDED    EventType = 0x1400 + iota // A new camera device is available
	EVENT_CAMERA_DEVICE_REMOVED                            // A camera device has been removed.
	EVENT_CAMERA_DEVICE_APPROVED                           // A camera device has been approved for use by the user.
	EVENT_CAMERA_DEVICE_DENIED                             // A camera device has been denied for use by the user.
)

// Notification events

const (
	EVENT_NOTIFICATION_ACTION_INVOKED EventType = 0x1500 + iota // A user response to a system notification was received.
)

// Render events

const (
	EVENT_RENDER_TARGETS_RESET EventType = 0x2000 + iota // The render targets have been reset and their contents need to be updated
	EVENT_RENDER_DEVICE_RESET                            // The device has been reset and all textures need to be recreated
	EVENT_RENDER_DEVICE_LOST                             // The device has been lost and can't be recovered.
)

// Reserved events for private platforms

const (
	EVENT_PRIVATE0 EventType = 0x4000 + iota
	EVENT_PRIVATE1
	EVENT_PRIVATE2
	EVENT_PRIVATE3
)

// Internal events

const (
	EVENT_POLL_SENTINEL EventType = 0x7F00 + iota // Signals the end of an event poll cycle
)

// Events SDL_EVENT_USER through SDL_EVENT_LAST are for your use, and should be
// allocated with SDL_RegisterEvents().
const EVENT_USER EventType = 0x8000

// This last event is only for bounding internal arrays.
const EVENT_LAST EventType = 0xFFFF

type Scancode uint32

const (
	SCANCODE_UNKNOWN Scancode = 0

	// These values are from usage page 0x07 (USB keyboard page).

	SCANCODE_A Scancode = 4
	SCANCODE_B Scancode = 5
	SCANCODE_C Scancode = 6
	SCANCODE_D Scancode = 7
	SCANCODE_E Scancode = 8
	SCANCODE_F Scancode = 9
	SCANCODE_G Scancode = 10
	SCANCODE_H Scancode = 11
	SCANCODE_I Scancode = 12
	SCANCODE_J Scancode = 13
	SCANCODE_K Scancode = 14
	SCANCODE_L Scancode = 15
	SCANCODE_M Scancode = 16
	SCANCODE_N Scancode = 17
	SCANCODE_O Scancode = 18
	SCANCODE_P Scancode = 19
	SCANCODE_Q Scancode = 20
	SCANCODE_R Scancode = 21
	SCANCODE_S Scancode = 22
	SCANCODE_T Scancode = 23
	SCANCODE_U Scancode = 24
	SCANCODE_V Scancode = 25
	SCANCODE_W Scancode = 26
	SCANCODE_X Scancode = 27
	SCANCODE_Y Scancode = 28
	SCANCODE_Z Scancode = 29

	SCANCODE_1 Scancode = 30
	SCANCODE_2 Scancode = 31
	SCANCODE_3 Scancode = 32
	SCANCODE_4 Scancode = 33
	SCANCODE_5 Scancode = 34
	SCANCODE_6 Scancode = 35
	SCANCODE_7 Scancode = 36
	SCANCODE_8 Scancode = 37
	SCANCODE_9 Scancode = 38
	SCANCODE_0 Scancode = 39

	SCANCODE_RETURN    Scancode = 40
	SCANCODE_ESCAPE    Scancode = 41
	SCANCODE_BACKSPACE Scancode = 42
	SCANCODE_TAB       Scancode = 43
	SCANCODE_SPACE     Scancode = 44

	SCANCODE_MINUS        Scancode = 45
	SCANCODE_EQUALS       Scancode = 46
	SCANCODE_LEFTBRACKET  Scancode = 47
	SCANCODE_RIGHTBRACKET Scancode = 48
	// Located at the lower left of the return key on ISO keyboards and at the
	// right end of the QWERTY row on ANSI keyboards.  Produces REVERSE SOLIDUS
	// (backslash) and VERTICAL LINE in a US layout, REVERSE SOLIDUS and VERTICAL
	// LINE in a UK Mac layout, NUMBER SIGN and TILDE in a UK Windows layout,
	// DOLLAR SIGN and POUND SIGN in a Swiss German layout, NUMBER SIGN and
	// APOSTROPHE in a German layout, GRAVE ACCENT and POUND SIGN in a French Mac
	// layout, and ASTERISK and MICRO SIGN in a French Windows layout.
	SCANCODE_BACKSLASH Scancode = 49
	// ISO USB keyboards actually use this code instead of 49 for the same key, but
	// all OSes I've seen treat the two codes identically. So, as an implementer,
	// unless your keyboard generates both of those codes and your OS treats them
	// differently, you should generate SDL_SCANCODE_BACKSLASH instead of this
	// code. As a user, you should not rely on this code because SDL will never
	// generate it with most (all?) keyboards.
	SCANCODE_NONUSHASH  Scancode = 50
	SCANCODE_SEMICOLON  Scancode = 51
	SCANCODE_APOSTROPHE Scancode = 52
	// Located in the top left corner (on both ANSI and ISO keyboards). Produces
	// GRAVE ACCENT and TILDE in a US Windows layout and in US and UK Mac layouts
	// on ANSI keyboards, GRAVE ACCENT and NOT SIGN in a UK Windows layout, SECTION
	// SIGN and PLUS-MINUS SIGN in US and UK Mac layouts on ISO keyboards, SECTION
	// SIGN and DEGREE SIGN in a Swiss German layout (Mac: only on ISO keyboards),
	// CIRCUMFLEX ACCENT and DEGREE SIGN in a German layout (Mac: only on ISO
	// keyboards), SUPERSCRIPT TWO and TILDE in a French Windows layout, COMMERCIAL
	// AT and NUMBER SIGN in a French Mac layout on ISO keyboards, and LESS-THAN
	// SIGN and GREATER-THAN SIGN in a Swiss German, German, or French Mac layout
	// on ANSI keyboards.
	SCANCODE_GRAVE  Scancode = 53
	SCANCODE_COMMA  Scancode = 54
	SCANCODE_PERIOD Scancode = 55
	SCANCODE_SLASH  Scancode = 56

	SCANCODE_CAPSLOCK Scancode = 57

	SCANCODE_F1  Scancode = 58
	SCANCODE_F2  Scancode = 59
	SCANCODE_F3  Scancode = 60
	SCANCODE_F4  Scancode = 61
	SCANCODE_F5  Scancode = 62
	SCANCODE_F6  Scancode = 63
	SCANCODE_F7  Scancode = 64
	SCANCODE_F8  Scancode = 65
	SCANCODE_F9  Scancode = 66
	SCANCODE_F10 Scancode = 67
	SCANCODE_F11 Scancode = 68
	SCANCODE_F12 Scancode = 69

	SCANCODE_PRINTSCREEN Scancode = 70
	SCANCODE_SCROLLLOCK  Scancode = 71
	SCANCODE_PAUSE       Scancode = 72
	// insert on PC, help on some Mac keyboards (but does send code 73, not 117).
	SCANCODE_INSERT   Scancode = 73
	SCANCODE_HOME     Scancode = 74
	SCANCODE_PAGEUP   Scancode = 75
	SCANCODE_DELETE   Scancode = 76
	SCANCODE_END      Scancode = 77
	SCANCODE_PAGEDOWN Scancode = 78
	SCANCODE_RIGHT    Scancode = 79
	SCANCODE_LEFT     Scancode = 80
	SCANCODE_DOWN     Scancode = 81
	SCANCODE_UP       Scancode = 82

	// num lock on PC, clear on Mac keyboards.
	SCANCODE_NUMLOCKCLEAR Scancode = 83
	SCANCODE_KP_DIVIDE    Scancode = 84
	SCANCODE_KP_MULTIPLY  Scancode = 85
	SCANCODE_KP_MINUS     Scancode = 86
	SCANCODE_KP_PLUS      Scancode = 87
	SCANCODE_KP_ENTER     Scancode = 88
	SCANCODE_KP_1         Scancode = 89
	SCANCODE_KP_2         Scancode = 90
	SCANCODE_KP_3         Scancode = 91
	SCANCODE_KP_4         Scancode = 92
	SCANCODE_KP_5         Scancode = 93
	SCANCODE_KP_6         Scancode = 94
	SCANCODE_KP_7         Scancode = 95
	SCANCODE_KP_8         Scancode = 96
	SCANCODE_KP_9         Scancode = 97
	SCANCODE_KP_0         Scancode = 98
	SCANCODE_KP_PERIOD    Scancode = 99

	// This is the additional key that ISO keyboards have over ANSI ones, located
	// between left shift and Z. Produces GRAVE ACCENT and TILDE in a US or UK Mac
	// layout, REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US or UK Windows
	// layout, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German,
	// or French layout.
	SCANCODE_NONUSBACKSLASH Scancode = 100
	SCANCODE_APPLICATION    Scancode = 101 // windows contextual menu, compose
	// The USB document says this is a status flag, not a physical key - but some
	// Mac keyboards do have a power key. */.
	SCANCODE_POWER      Scancode = 102
	SCANCODE_KP_EQUALS  Scancode = 103
	SCANCODE_F13        Scancode = 104
	SCANCODE_F14        Scancode = 105
	SCANCODE_F15        Scancode = 106
	SCANCODE_F16        Scancode = 107
	SCANCODE_F17        Scancode = 108
	SCANCODE_F18        Scancode = 109
	SCANCODE_F19        Scancode = 110
	SCANCODE_F20        Scancode = 111
	SCANCODE_F21        Scancode = 112
	SCANCODE_F22        Scancode = 113
	SCANCODE_F23        Scancode = 114
	SCANCODE_F24        Scancode = 115
	SCANCODE_EXECUTE    Scancode = 116
	SCANCODE_HELP       Scancode = 117 // AL Integrated Help Center
	SCANCODE_MENU       Scancode = 118 // Menu (show menu)
	SCANCODE_SELECT     Scancode = 119
	SCANCODE_STOP       Scancode = 120 // AC Stop
	SCANCODE_AGAIN      Scancode = 121 // AC Redo/Repeat
	SCANCODE_UNDO       Scancode = 122 // AC Undo
	SCANCODE_CUT        Scancode = 123 // AC Cut
	SCANCODE_COPY       Scancode = 124 // AC Copy
	SCANCODE_PASTE      Scancode = 125 // AC Paste
	SCANCODE_FIND       Scancode = 126 // AC Find
	SCANCODE_MUTE       Scancode = 127
	SCANCODE_VOLUMEUP   Scancode = 128
	SCANCODE_VOLUMEDOWN Scancode = 129
	// not sure whether there's a reason to enable these
	//
	//	SDL_SCANCODE_LOCKINGCAPSLOCK = 130,
	//	SDL_SCANCODE_LOCKINGNUMLOCK = 131,
	//	SDL_SCANCODE_LOCKINGSCROLLLOCK = 132,
	SCANCODE_KP_COMMA       Scancode = 133
	SCANCODE_KP_EQUALSAS400 Scancode = 134

	SCANCODE_INTERNATIONAL1 Scancode = 135 // used on Asian keyboards, see footnotes in USB doc
	SCANCODE_INTERNATIONAL2 Scancode = 136
	SCANCODE_INTERNATIONAL3 Scancode = 137 // Yen
	SCANCODE_INTERNATIONAL4 Scancode = 138
	SCANCODE_INTERNATIONAL5 Scancode = 139
	SCANCODE_INTERNATIONAL6 Scancode = 140
	SCANCODE_INTERNATIONAL7 Scancode = 141
	SCANCODE_INTERNATIONAL8 Scancode = 142
	SCANCODE_INTERNATIONAL9 Scancode = 143
	SCANCODE_LANG1          Scancode = 144 // Hangul/English toggle
	SCANCODE_LANG2          Scancode = 145 // Hanja conversion
	SCANCODE_LANG3          Scancode = 146 // Katakana
	SCANCODE_LANG4          Scancode = 147 // Hiragana
	SCANCODE_LANG5          Scancode = 148 // Zenkaku/Hankaku
	SCANCODE_LANG6          Scancode = 149 // reserved
	SCANCODE_LANG7          Scancode = 150 // reserved
	SCANCODE_LANG8          Scancode = 151 // reserved
	SCANCODE_LANG9          Scancode = 152 // reserved

	SCANCODE_ALTERASE   Scancode = 153 // Erase-Eaze
	SCANCODE_SYSREQ     Scancode = 154
	SCANCODE_CANCEL     Scancode = 155 // AC Cancel
	SCANCODE_CLEAR      Scancode = 156
	SCANCODE_PRIOR      Scancode = 157
	SCANCODE_RETURN2    Scancode = 158
	SCANCODE_SEPARATOR  Scancode = 159
	SCANCODE_OUT        Scancode = 160
	SCANCODE_OPER       Scancode = 161
	SCANCODE_CLEARAGAIN Scancode = 162
	SCANCODE_CRSEL      Scancode = 163
	SCANCODE_EXSEL      Scancode = 164

	SCANCODE_KP_00              Scancode = 176
	SCANCODE_KP_000             Scancode = 177
	SCANCODE_THOUSANDSSEPARATOR Scancode = 178
	SCANCODE_DECIMALSEPARATOR   Scancode = 179
	SCANCODE_CURRENCYUNIT       Scancode = 180
	SCANCODE_CURRENCYSUBUNIT    Scancode = 181
	SCANCODE_KP_LEFTPAREN       Scancode = 182
	SCANCODE_KP_RIGHTPAREN      Scancode = 183
	SCANCODE_KP_LEFTBRACE       Scancode = 184
	SCANCODE_KP_RIGHTBRACE      Scancode = 185
	SCANCODE_KP_TAB             Scancode = 186
	SCANCODE_KP_BACKSPACE       Scancode = 187
	SCANCODE_KP_A               Scancode = 188
	SCANCODE_KP_B               Scancode = 189
	SCANCODE_KP_C               Scancode = 190
	SCANCODE_KP_D               Scancode = 191
	SCANCODE_KP_E               Scancode = 192
	SCANCODE_KP_F               Scancode = 193
	SCANCODE_KP_XOR             Scancode = 194
	SCANCODE_KP_POWER           Scancode = 195
	SCANCODE_KP_PERCENT         Scancode = 196
	SCANCODE_KP_LESS            Scancode = 197
	SCANCODE_KP_GREATER         Scancode = 198
	SCANCODE_KP_AMPERSAND       Scancode = 199
	SCANCODE_KP_DBLAMPERSAND    Scancode = 200
	SCANCODE_KP_VERTICALBAR     Scancode = 201
	SCANCODE_KP_DBLVERTICALBAR  Scancode = 202
	SCANCODE_KP_COLON           Scancode = 203
	SCANCODE_KP_HASH            Scancode = 204
	SCANCODE_KP_SPACE           Scancode = 205
	SCANCODE_KP_AT              Scancode = 206
	SCANCODE_KP_EXCLAM          Scancode = 207
	SCANCODE_KP_MEMSTORE        Scancode = 208
	SCANCODE_KP_MEMRECALL       Scancode = 209
	SCANCODE_KP_MEMCLEAR        Scancode = 210
	SCANCODE_KP_MEMADD          Scancode = 211
	SCANCODE_KP_MEMSUBTRACT     Scancode = 212
	SCANCODE_KP_MEMMULTIPLY     Scancode = 213
	SCANCODE_KP_MEMDIVIDE       Scancode = 214
	SCANCODE_KP_PLUSMINUS       Scancode = 215
	SCANCODE_KP_CLEAR           Scancode = 216
	SCANCODE_KP_CLEARENTRY      Scancode = 217
	SCANCODE_KP_BINARY          Scancode = 218
	SCANCODE_KP_OCTAL           Scancode = 219
	SCANCODE_KP_DECIMAL         Scancode = 220
	SCANCODE_KP_HEXADECIMAL     Scancode = 221

	SCANCODE_LCTRL  Scancode = 224
	SCANCODE_LSHIFT Scancode = 225
	SCANCODE_LALT   Scancode = 226 // alt, option
	SCANCODE_LGUI   Scancode = 227 // windows, command (apple), meta
	SCANCODE_RCTRL  Scancode = 228
	SCANCODE_RSHIFT Scancode = 229
	SCANCODE_RALT   Scancode = 230 // alt gr, option
	SCANCODE_RGUI   Scancode = 231 // windows, command (apple), meta

	// I'm not sure if this is really not covered by any of the above, but since
	// there's a special SDL_KMOD_MODE for it I'm adding it here.
	SCANCODE_MODE Scancode = 257

	// These values are mapped from usage page 0x0C (USB consumer page).
	//
	// There are way more keys in the spec than we can represent in the current
	// scancode range, so pick the ones that commonly come up in real world usage.

	SCANCODE_SLEEP Scancode = 258 // Sleep
	SCANCODE_WAKE  Scancode = 259 // Wake

	SCANCODE_CHANNEL_INCREMENT Scancode = 260 // Channel Increment
	SCANCODE_CHANNEL_DECREMENT Scancode = 261 // Channel Decrement

	SCANCODE_MEDIA_PLAY           Scancode = 262 // Play
	SCANCODE_MEDIA_PAUSE          Scancode = 263 // Pause
	SCANCODE_MEDIA_RECORD         Scancode = 264 // Record
	SCANCODE_MEDIA_FAST_FORWARD   Scancode = 265 // Fast Forward
	SCANCODE_MEDIA_REWIND         Scancode = 266 // Rewind
	SCANCODE_MEDIA_NEXT_TRACK     Scancode = 267 // Next Track
	SCANCODE_MEDIA_PREVIOUS_TRACK Scancode = 268 // Previous Track
	SCANCODE_MEDIA_STOP           Scancode = 269 // Stop
	SCANCODE_MEDIA_EJECT          Scancode = 270 // Eject
	SCANCODE_MEDIA_PLAY_PAUSE     Scancode = 271 // Play / Pause
	SCANCODE_MEDIA_SELECT         Scancode = 272 // Media Select

	SCANCODE_AC_NEW        Scancode = 273 // AC New
	SCANCODE_AC_OPEN       Scancode = 274 // AC Open
	SCANCODE_AC_CLOSE      Scancode = 275 // AC Close
	SCANCODE_AC_EXIT       Scancode = 276 // AC Exit
	SCANCODE_AC_SAVE       Scancode = 277 // AC Save
	SCANCODE_AC_PRINT      Scancode = 278 // AC Print
	SCANCODE_AC_PROPERTIES Scancode = 279 // AC Properties

	SCANCODE_AC_SEARCH    Scancode = 280 // AC Search
	SCANCODE_AC_HOME      Scancode = 281 // AC Home
	SCANCODE_AC_BACK      Scancode = 282 // AC Back
	SCANCODE_AC_FORWARD   Scancode = 283 // AC Forward
	SCANCODE_AC_STOP      Scancode = 284 // AC Stop
	SCANCODE_AC_REFRESH   Scancode = 285 // AC Refresh
	SCANCODE_AC_BOOKMARKS Scancode = 286 // AC Bookmarks

	// These are values that are often used on mobile phones.

	// Usually situated below the display on phones and used as a multi-function
	// feature key for selecting a software defined function shown on the bottom
	// left of the display.
	SCANCODE_SOFTLEFT Scancode = 287
	// Usually situated below the display on phones and used as a multi-function
	// feature key for selecting a software defined function shown on the bottom
	// right of the display.
	SCANCODE_SOFTRIGHT Scancode = 288
	SCANCODE_CALL      Scancode = 289 // Used for accepting phone calls.
	SCANCODE_ENDCALL   Scancode = 290 // Used for rejecting phone calls.

	// Add any other keys here.

	SCANCODE_RESERVED Scancode = 400 // 400-500 reserved for dynamic keycodes

	SCANCODE_COUNT Scancode = 512 // not a key, just marks the number of scancodes for array bounds
)

type Keycode uint32

func ScancodeToKeycode(scancode Scancode) Keycode {
	return Keycode(scancode) | SDLK_SCANCODE_MASK
}

const (
	SDLK_EXTENDED_MASK        Keycode = (1 << 29)
	SDLK_SCANCODE_MASK        Keycode = (1 << 30)
	SDLK_UNKNOWN              Keycode = 0x00000000 // 0
	SDLK_RETURN               Keycode = 0x0000000d // '\r'
	SDLK_ESCAPE               Keycode = 0x0000001b // '\x1B'
	SDLK_BACKSPACE            Keycode = 0x00000008 // '\b'
	SDLK_TAB                  Keycode = 0x00000009 // '\t'
	SDLK_SPACE                Keycode = 0x00000020 // ' '
	SDLK_EXCLAIM              Keycode = 0x00000021 // '!'
	SDLK_DBLAPOSTROPHE        Keycode = 0x00000022 // '"'
	SDLK_HASH                 Keycode = 0x00000023 // '#'
	SDLK_DOLLAR               Keycode = 0x00000024 // '$'
	SDLK_PERCENT              Keycode = 0x00000025 // '%'
	SDLK_AMPERSAND            Keycode = 0x00000026 // '&'
	SDLK_APOSTROPHE           Keycode = 0x00000027 // '\''
	SDLK_LEFTPAREN            Keycode = 0x00000028 // '('
	SDLK_RIGHTPAREN           Keycode = 0x00000029 // ')'
	SDLK_ASTERISK             Keycode = 0x0000002a // '*'
	SDLK_PLUS                 Keycode = 0x0000002b // '+'
	SDLK_COMMA                Keycode = 0x0000002c // ','
	SDLK_MINUS                Keycode = 0x0000002d // '-'
	SDLK_PERIOD               Keycode = 0x0000002e // '.'
	SDLK_SLASH                Keycode = 0x0000002f // '/'
	SDLK_0                    Keycode = 0x00000030 // '0'
	SDLK_1                    Keycode = 0x00000031 // '1'
	SDLK_2                    Keycode = 0x00000032 // '2'
	SDLK_3                    Keycode = 0x00000033 // '3'
	SDLK_4                    Keycode = 0x00000034 // '4'
	SDLK_5                    Keycode = 0x00000035 // '5'
	SDLK_6                    Keycode = 0x00000036 // '6'
	SDLK_7                    Keycode = 0x00000037 // '7'
	SDLK_8                    Keycode = 0x00000038 // '8'
	SDLK_9                    Keycode = 0x00000039 // '9'
	SDLK_COLON                Keycode = 0x0000003a // ':'
	SDLK_SEMICOLON            Keycode = 0x0000003b // ';'
	SDLK_LESS                 Keycode = 0x0000003c // '<'
	SDLK_EQUALS               Keycode = 0x0000003d // '='
	SDLK_GREATER              Keycode = 0x0000003e // '>'
	SDLK_QUESTION             Keycode = 0x0000003f // '?'
	SDLK_AT                   Keycode = 0x00000040 // '@'
	SDLK_LEFTBRACKET          Keycode = 0x0000005b // '['
	SDLK_BACKSLASH            Keycode = 0x0000005c // '\\'
	SDLK_RIGHTBRACKET         Keycode = 0x0000005d // ']'
	SDLK_CARET                Keycode = 0x0000005e // '^'
	SDLK_UNDERSCORE           Keycode = 0x0000005f // '_'
	SDLK_GRAVE                Keycode = 0x00000060 // '`'
	SDLK_A                    Keycode = 0x00000061 // 'a'
	SDLK_B                    Keycode = 0x00000062 // 'b'
	SDLK_C                    Keycode = 0x00000063 // 'c'
	SDLK_D                    Keycode = 0x00000064 // 'd'
	SDLK_E                    Keycode = 0x00000065 // 'e'
	SDLK_F                    Keycode = 0x00000066 // 'f'
	SDLK_G                    Keycode = 0x00000067 // 'g'
	SDLK_H                    Keycode = 0x00000068 // 'h'
	SDLK_I                    Keycode = 0x00000069 // 'i'
	SDLK_J                    Keycode = 0x0000006a // 'j'
	SDLK_K                    Keycode = 0x0000006b // 'k'
	SDLK_L                    Keycode = 0x0000006c // 'l'
	SDLK_M                    Keycode = 0x0000006d // 'm'
	SDLK_N                    Keycode = 0x0000006e // 'n'
	SDLK_O                    Keycode = 0x0000006f // 'o'
	SDLK_P                    Keycode = 0x00000070 // 'p'
	SDLK_Q                    Keycode = 0x00000071 // 'q'
	SDLK_R                    Keycode = 0x00000072 // 'r'
	SDLK_S                    Keycode = 0x00000073 // 's'
	SDLK_T                    Keycode = 0x00000074 // 't'
	SDLK_U                    Keycode = 0x00000075 // 'u'
	SDLK_V                    Keycode = 0x00000076 // 'v'
	SDLK_W                    Keycode = 0x00000077 // 'w'
	SDLK_X                    Keycode = 0x00000078 // 'x'
	SDLK_Y                    Keycode = 0x00000079 // 'y'
	SDLK_Z                    Keycode = 0x0000007a // 'z'
	SDLK_LEFTBRACE            Keycode = 0x0000007b // '{'
	SDLK_PIPE                 Keycode = 0x0000007c // '|'
	SDLK_RIGHTBRACE           Keycode = 0x0000007d // '}'
	SDLK_TILDE                Keycode = 0x0000007e // '~'
	SDLK_DELETE               Keycode = 0x0000007f // '\x7F'
	SDLK_PLUSMINUS            Keycode = 0x000000b1 // '\xB1'
	SDLK_CAPSLOCK             Keycode = 0x40000039 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CAPSLOCK)
	SDLK_F1                   Keycode = 0x4000003a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F1)
	SDLK_F2                   Keycode = 0x4000003b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F2)
	SDLK_F3                   Keycode = 0x4000003c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F3)
	SDLK_F4                   Keycode = 0x4000003d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F4)
	SDLK_F5                   Keycode = 0x4000003e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F5)
	SDLK_F6                   Keycode = 0x4000003f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F6)
	SDLK_F7                   Keycode = 0x40000040 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F7)
	SDLK_F8                   Keycode = 0x40000041 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F8)
	SDLK_F9                   Keycode = 0x40000042 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F9)
	SDLK_F10                  Keycode = 0x40000043 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F10)
	SDLK_F11                  Keycode = 0x40000044 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F11)
	SDLK_F12                  Keycode = 0x40000045 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F12)
	SDLK_PRINTSCREEN          Keycode = 0x40000046 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRINTSCREEN)
	SDLK_SCROLLLOCK           Keycode = 0x40000047 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SCROLLLOCK)
	SDLK_PAUSE                Keycode = 0x40000048 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAUSE)
	SDLK_INSERT               Keycode = 0x40000049 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_INSERT)
	SDLK_HOME                 Keycode = 0x4000004a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HOME)
	SDLK_PAGEUP               Keycode = 0x4000004b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEUP)
	SDLK_END                  Keycode = 0x4000004d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_END)
	SDLK_PAGEDOWN             Keycode = 0x4000004e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEDOWN)
	SDLK_RIGHT                Keycode = 0x4000004f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RIGHT)
	SDLK_LEFT                 Keycode = 0x40000050 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LEFT)
	SDLK_DOWN                 Keycode = 0x40000051 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DOWN)
	SDLK_UP                   Keycode = 0x40000052 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UP)
	SDLK_NUMLOCKCLEAR         Keycode = 0x40000053 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_NUMLOCKCLEAR)
	SDLK_KP_DIVIDE            Keycode = 0x40000054 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DIVIDE)
	SDLK_KP_MULTIPLY          Keycode = 0x40000055 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MULTIPLY)
	SDLK_KP_MINUS             Keycode = 0x40000056 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MINUS)
	SDLK_KP_PLUS              Keycode = 0x40000057 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUS)
	SDLK_KP_ENTER             Keycode = 0x40000058 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_ENTER)
	SDLK_KP_1                 Keycode = 0x40000059 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_1)
	SDLK_KP_2                 Keycode = 0x4000005a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_2)
	SDLK_KP_3                 Keycode = 0x4000005b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_3)
	SDLK_KP_4                 Keycode = 0x4000005c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_4)
	SDLK_KP_5                 Keycode = 0x4000005d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_5)
	SDLK_KP_6                 Keycode = 0x4000005e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_6)
	SDLK_KP_7                 Keycode = 0x4000005f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_7)
	SDLK_KP_8                 Keycode = 0x40000060 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_8)
	SDLK_KP_9                 Keycode = 0x40000061 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_9)
	SDLK_KP_0                 Keycode = 0x40000062 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_0)
	SDLK_KP_PERIOD            Keycode = 0x40000063 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERIOD)
	SDLK_APPLICATION          Keycode = 0x40000065 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_APPLICATION)
	SDLK_POWER                Keycode = 0x40000066 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_POWER)
	SDLK_KP_EQUALS            Keycode = 0x40000067 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALS)
	SDLK_F13                  Keycode = 0x40000068 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F13)
	SDLK_F14                  Keycode = 0x40000069 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F14)
	SDLK_F15                  Keycode = 0x4000006a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F15)
	SDLK_F16                  Keycode = 0x4000006b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F16)
	SDLK_F17                  Keycode = 0x4000006c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F17)
	SDLK_F18                  Keycode = 0x4000006d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F18)
	SDLK_F19                  Keycode = 0x4000006e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F19)
	SDLK_F20                  Keycode = 0x4000006f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F20)
	SDLK_F21                  Keycode = 0x40000070 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F21)
	SDLK_F22                  Keycode = 0x40000071 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F22)
	SDLK_F23                  Keycode = 0x40000072 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F23)
	SDLK_F24                  Keycode = 0x40000073 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F24)
	SDLK_EXECUTE              Keycode = 0x40000074 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXECUTE)
	SDLK_HELP                 Keycode = 0x40000075 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HELP)
	SDLK_MENU                 Keycode = 0x40000076 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MENU)
	SDLK_SELECT               Keycode = 0x40000077 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SELECT)
	SDLK_STOP                 Keycode = 0x40000078 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_STOP)
	SDLK_AGAIN                Keycode = 0x40000079 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AGAIN)
	SDLK_UNDO                 Keycode = 0x4000007a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UNDO)
	SDLK_CUT                  Keycode = 0x4000007b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CUT)
	SDLK_COPY                 Keycode = 0x4000007c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_COPY)
	SDLK_PASTE                Keycode = 0x4000007d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PASTE)
	SDLK_FIND                 Keycode = 0x4000007e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_FIND)
	SDLK_MUTE                 Keycode = 0x4000007f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MUTE)
	SDLK_VOLUMEUP             Keycode = 0x40000080 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEUP)
	SDLK_VOLUMEDOWN           Keycode = 0x40000081 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEDOWN)
	SDLK_KP_COMMA             Keycode = 0x40000085 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COMMA)
	SDLK_KP_EQUALSAS400       Keycode = 0x40000086 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALSAS400)
	SDLK_ALTERASE             Keycode = 0x40000099 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_ALTERASE)
	SDLK_SYSREQ               Keycode = 0x4000009a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SYSREQ)
	SDLK_CANCEL               Keycode = 0x4000009b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CANCEL)
	SDLK_CLEAR                Keycode = 0x4000009c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEAR)
	SDLK_PRIOR                Keycode = 0x4000009d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRIOR)
	SDLK_RETURN2              Keycode = 0x4000009e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RETURN2)
	SDLK_SEPARATOR            Keycode = 0x4000009f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SEPARATOR)
	SDLK_OUT                  Keycode = 0x400000a0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OUT)
	SDLK_OPER                 Keycode = 0x400000a1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OPER)
	SDLK_CLEARAGAIN           Keycode = 0x400000a2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEARAGAIN)
	SDLK_CRSEL                Keycode = 0x400000a3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CRSEL)
	SDLK_EXSEL                Keycode = 0x400000a4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXSEL)
	SDLK_KP_00                Keycode = 0x400000b0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_00)
	SDLK_KP_000               Keycode = 0x400000b1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_000)
	SDLK_THOUSANDSSEPARATOR   Keycode = 0x400000b2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_THOUSANDSSEPARATOR)
	SDLK_DECIMALSEPARATOR     Keycode = 0x400000b3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DECIMALSEPARATOR)
	SDLK_CURRENCYUNIT         Keycode = 0x400000b4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYUNIT)
	SDLK_CURRENCYSUBUNIT      Keycode = 0x400000b5 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYSUBUNIT)
	SDLK_KP_LEFTPAREN         Keycode = 0x400000b6 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTPAREN)
	SDLK_KP_RIGHTPAREN        Keycode = 0x400000b7 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTPAREN)
	SDLK_KP_LEFTBRACE         Keycode = 0x400000b8 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTBRACE)
	SDLK_KP_RIGHTBRACE        Keycode = 0x400000b9 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTBRACE)
	SDLK_KP_TAB               Keycode = 0x400000ba // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_TAB)
	SDLK_KP_BACKSPACE         Keycode = 0x400000bb // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BACKSPACE)
	SDLK_KP_A                 Keycode = 0x400000bc // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_A)
	SDLK_KP_B                 Keycode = 0x400000bd // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_B)
	SDLK_KP_C                 Keycode = 0x400000be // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_C)
	SDLK_KP_D                 Keycode = 0x400000bf // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_D)
	SDLK_KP_E                 Keycode = 0x400000c0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_E)
	SDLK_KP_F                 Keycode = 0x400000c1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_F)
	SDLK_KP_XOR               Keycode = 0x400000c2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_XOR)
	SDLK_KP_POWER             Keycode = 0x400000c3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_POWER)
	SDLK_KP_PERCENT           Keycode = 0x400000c4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERCENT)
	SDLK_KP_LESS              Keycode = 0x400000c5 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LESS)
	SDLK_KP_GREATER           Keycode = 0x400000c6 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_GREATER)
	SDLK_KP_AMPERSAND         Keycode = 0x400000c7 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AMPERSAND)
	SDLK_KP_DBLAMPERSAND      Keycode = 0x400000c8 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLAMPERSAND)
	SDLK_KP_VERTICALBAR       Keycode = 0x400000c9 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_VERTICALBAR)
	SDLK_KP_DBLVERTICALBAR    Keycode = 0x400000ca // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLVERTICALBAR)
	SDLK_KP_COLON             Keycode = 0x400000cb // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COLON)
	SDLK_KP_HASH              Keycode = 0x400000cc // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HASH)
	SDLK_KP_SPACE             Keycode = 0x400000cd // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_SPACE)
	SDLK_KP_AT                Keycode = 0x400000ce // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AT)
	SDLK_KP_EXCLAM            Keycode = 0x400000cf // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EXCLAM)
	SDLK_KP_MEMSTORE          Keycode = 0x400000d0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSTORE)
	SDLK_KP_MEMRECALL         Keycode = 0x400000d1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMRECALL)
	SDLK_KP_MEMCLEAR          Keycode = 0x400000d2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMCLEAR)
	SDLK_KP_MEMADD            Keycode = 0x400000d3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMADD)
	SDLK_KP_MEMSUBTRACT       Keycode = 0x400000d4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSUBTRACT)
	SDLK_KP_MEMMULTIPLY       Keycode = 0x400000d5 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMMULTIPLY)
	SDLK_KP_MEMDIVIDE         Keycode = 0x400000d6 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMDIVIDE)
	SDLK_KP_PLUSMINUS         Keycode = 0x400000d7 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUSMINUS)
	SDLK_KP_CLEAR             Keycode = 0x400000d8 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEAR)
	SDLK_KP_CLEARENTRY        Keycode = 0x400000d9 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEARENTRY)
	SDLK_KP_BINARY            Keycode = 0x400000da // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BINARY)
	SDLK_KP_OCTAL             Keycode = 0x400000db // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_OCTAL)
	SDLK_KP_DECIMAL           Keycode = 0x400000dc // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DECIMAL)
	SDLK_KP_HEXADECIMAL       Keycode = 0x400000dd // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HEXADECIMAL)
	SDLK_LCTRL                Keycode = 0x400000e0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LCTRL)
	SDLK_LSHIFT               Keycode = 0x400000e1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LSHIFT)
	SDLK_LALT                 Keycode = 0x400000e2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LALT)
	SDLK_LGUI                 Keycode = 0x400000e3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LGUI)
	SDLK_RCTRL                Keycode = 0x400000e4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RCTRL)
	SDLK_RSHIFT               Keycode = 0x400000e5 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RSHIFT)
	SDLK_RALT                 Keycode = 0x400000e6 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RALT)
	SDLK_RGUI                 Keycode = 0x400000e7 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RGUI)
	SDLK_MODE                 Keycode = 0x40000101 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MODE)
	SDLK_SLEEP                Keycode = 0x40000102 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SLEEP)
	SDLK_WAKE                 Keycode = 0x40000103 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_WAKE)
	SDLK_CHANNEL_INCREMENT    Keycode = 0x40000104 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CHANNEL_INCREMENT)
	SDLK_CHANNEL_DECREMENT    Keycode = 0x40000105 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CHANNEL_DECREMENT)
	SDLK_MEDIA_PLAY           Keycode = 0x40000106 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PLAY)
	SDLK_MEDIA_PAUSE          Keycode = 0x40000107 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PAUSE)
	SDLK_MEDIA_RECORD         Keycode = 0x40000108 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_RECORD)
	SDLK_MEDIA_FAST_FORWARD   Keycode = 0x40000109 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_FAST_FORWARD)
	SDLK_MEDIA_REWIND         Keycode = 0x4000010a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_REWIND)
	SDLK_MEDIA_NEXT_TRACK     Keycode = 0x4000010b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_NEXT_TRACK)
	SDLK_MEDIA_PREVIOUS_TRACK Keycode = 0x4000010c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PREVIOUS_TRACK)
	SDLK_MEDIA_STOP           Keycode = 0x4000010d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_STOP)
	SDLK_MEDIA_EJECT          Keycode = 0x4000010e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_EJECT)
	SDLK_MEDIA_PLAY_PAUSE     Keycode = 0x4000010f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PLAY_PAUSE)
	SDLK_MEDIA_SELECT         Keycode = 0x40000110 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_SELECT)
	SDLK_AC_NEW               Keycode = 0x40000111 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_NEW)
	SDLK_AC_OPEN              Keycode = 0x40000112 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_OPEN)
	SDLK_AC_CLOSE             Keycode = 0x40000113 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_CLOSE)
	SDLK_AC_EXIT              Keycode = 0x40000114 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_EXIT)
	SDLK_AC_SAVE              Keycode = 0x40000115 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_SAVE)
	SDLK_AC_PRINT             Keycode = 0x40000116 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_PRINT)
	SDLK_AC_PROPERTIES        Keycode = 0x40000117 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_PROPERTIES)
	SDLK_AC_SEARCH            Keycode = 0x40000118 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_SEARCH)
	SDLK_AC_HOME              Keycode = 0x40000119 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_HOME)
	SDLK_AC_BACK              Keycode = 0x4000011a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BACK)
	SDLK_AC_FORWARD           Keycode = 0x4000011b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_FORWARD)
	SDLK_AC_STOP              Keycode = 0x4000011c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_STOP)
	SDLK_AC_REFRESH           Keycode = 0x4000011d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_REFRESH)
	SDLK_AC_BOOKMARKS         Keycode = 0x4000011e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BOOKMARKS)
	SDLK_SOFTLEFT             Keycode = 0x4000011f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SOFTLEFT)
	SDLK_SOFTRIGHT            Keycode = 0x40000120 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SOFTRIGHT)
	SDLK_CALL                 Keycode = 0x40000121 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CALL)
	SDLK_ENDCALL              Keycode = 0x40000122 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_ENDCALL)
	SDLK_LEFT_TAB             Keycode = 0x20000001 // Extended key Left Tab
	SDLK_LEVEL5_SHIFT         Keycode = 0x20000002 // Extended key Level 5 Shift
	SDLK_MULTI_KEY_COMPOSE    Keycode = 0x20000003 // Extended key Multi-key Compose
	SDLK_LMETA                Keycode = 0x20000004 // Extended key Left Meta
	SDLK_RMETA                Keycode = 0x20000005 // Extended key Right Meta
	SDLK_LHYPER               Keycode = 0x20000006 // Extended key Left Hyper
	SDLK_RHYPER               Keycode = 0x20000007 // Extended key Right Hyper
)

type Keymod uint16

const (
	KMOD_NONE   Keymod = 0x0000                      // no modifier is applicable.
	KMOD_LSHIFT Keymod = 0x0001                      // the left Shift key is down.
	KMOD_RSHIFT Keymod = 0x0002                      // the right Shift key is down.
	KMOD_LEVEL5 Keymod = 0x0004                      // the Level 5 Shift key is down.
	KMOD_LCTRL  Keymod = 0x0040                      // the left Ctrl (Control) key is down.
	KMOD_RCTRL  Keymod = 0x0080                      // the right Ctrl (Control) key is down.
	KMOD_LALT   Keymod = 0x0100                      // the left Alt key is down.
	KMOD_RALT   Keymod = 0x0200                      // the right Alt key is down.
	KMOD_LGUI   Keymod = 0x0400                      // the left GUI key (often the Windows key) is down.
	KMOD_RGUI   Keymod = 0x0800                      // the right GUI key (often the Windows key) is down.
	KMOD_NUM    Keymod = 0x1000                      // the Num Lock key (may be located on an extended keypad) is down.
	KMOD_CAPS   Keymod = 0x2000                      // the Caps Lock key is down.
	KMOD_MODE   Keymod = 0x4000                      // the !AltGr key is down.
	KMOD_SCROLL Keymod = 0x8000                      // the Scroll Lock key is down.
	KMOD_CTRL   Keymod = (KMOD_LCTRL | KMOD_RCTRL)   // Any Ctrl key is down.
	KMOD_SHIFT  Keymod = (KMOD_LSHIFT | KMOD_RSHIFT) // Any Shift key is down.
	KMOD_ALT    Keymod = (KMOD_LALT | KMOD_RALT)     // Any Alt key is down.
	KMOD_GUI    Keymod = (KMOD_LGUI | KMOD_RGUI)     // Any GUI key is down.
)

type PowerState int8

const (
	POWERSTATE_ERROR      PowerState = -1 + iota // error determining power status
	POWERSTATE_UNKNOWN                           // cannot determine power status
	POWERSTATE_ON_BATTERY                        // Not plugged in, running on the battery
	POWERSTATE_NO_BATTERY                        // Plugged in, no battery available
	POWERSTATE_CHARGING                          // Plugged in, charging battery
	POWERSTATE_CHARGED                           // Plugged in, battery charged
)

type MouseButtonFlags uint32

const (
	BUTTON_LEFT   MouseButtonFlags = 1
	BUTTON_MIDDLE MouseButtonFlags = 2
	BUTTON_RIGHT  MouseButtonFlags = 3
	BUTTON_X1     MouseButtonFlags = 4
	BUTTON_X2     MouseButtonFlags = 5

	BUTTON_LMASK  MouseButtonFlags = 1 << (BUTTON_LEFT - 1)
	BUTTON_MMASK  MouseButtonFlags = 1 << (BUTTON_MIDDLE - 1)
	BUTTON_RMASK  MouseButtonFlags = 1 << (BUTTON_RIGHT - 1)
	BUTTON_X1MASK MouseButtonFlags = 1 << (BUTTON_X1 - 1)
	BUTTON_X2MASK MouseButtonFlags = 1 << (BUTTON_X2 - 1)
)

type MouseWheelDirection uint8

const (
	MOUSEWHEEL_NORMAL  MouseWheelDirection = iota /**< The scroll direction is normal */
	MOUSEWHEEL_FLIPPED                            /**< The scroll direction is flipped / natural */
)

type PenInputFlags uint32

const (
	PEN_INPUT_DOWN         PenInputFlags = 1 << 0  /**< pen is pressed down */
	PEN_INPUT_BUTTON_1     PenInputFlags = 1 << 1  /**< button 1 is pressed */
	PEN_INPUT_BUTTON_2     PenInputFlags = 1 << 2  /**< button 2 is pressed */
	PEN_INPUT_BUTTON_3     PenInputFlags = 1 << 3  /**< button 3 is pressed */
	PEN_INPUT_BUTTON_4     PenInputFlags = 1 << 4  /**< button 4 is pressed */
	PEN_INPUT_BUTTON_5     PenInputFlags = 1 << 5  /**< button 5 is pressed */
	PEN_INPUT_ERASER_TIP   PenInputFlags = 1 << 30 /**< eraser tip is used */
	PEN_INPUT_IN_PROXIMITY PenInputFlags = 1 << 31 /**< pen is in proximity (since SDL 3.4.0) */
)

type PenAxis uint16

const (
	PEN_AXIS_PRESSURE            PenAxis = iota /**< Pen pressure.  Unidirectional: 0 to 1.0 */
	PEN_AXIS_XTILT                              /**< Pen horizontal tilt angle.  Bidirectional: -90.0 to 90.0 (left-to-right). */
	PEN_AXIS_YTILT                              /**< Pen vertical tilt angle.  Bidirectional: -90.0 to 90.0 (top-to-down). */
	PEN_AXIS_DISTANCE                           /**< Pen distance to drawing surface.  Unidirectional: 0.0 to 1.0 */
	PEN_AXIS_ROTATION                           /**< Pen barrel rotation.  Bidirectional: -180 to 179.9 (clockwise, 0 is facing up, -180.0 is facing down). */
	PEN_AXIS_SLIDER                             /**< Pen finger wheel or slider (e.g., Airbrush Pen).  Unidirectional: 0 to 1.0 */
	PEN_AXIS_TANGENTIAL_PRESSURE                /**< Pressure from squeezing the pen ("barrel pressure"). */
	PEN_AXIS_COUNT                              /**< Total known pen axis types in this version of SDL. This number may grow in future releases! */
)

type WindowID uint32
type KeyboardID uint32
type MouseID uint32
type AudioDeviceID uint32
type CameraID uint32
type DisplayID uint32
type JoystickID uint32
type NotificationID uint32
type PenID uint32
type SensorID uint32
type TouchID uint64
type FingerID uint64

type JoystickHat uint8

const (
	HAT_CENTERED  JoystickHat = 0x00
	HAT_UP        JoystickHat = 0x01
	HAT_RIGHT     JoystickHat = 0x02
	HAT_DOWN      JoystickHat = 0x04
	HAT_LEFT      JoystickHat = 0x08
	HAT_RIGHTUP   JoystickHat = (HAT_RIGHT | HAT_UP)
	HAT_RIGHTDOWN JoystickHat = (HAT_RIGHT | HAT_DOWN)
	HAT_LEFTUP    JoystickHat = (HAT_LEFT | HAT_UP)
	HAT_LEFTDOWN  JoystickHat = (HAT_LEFT | HAT_DOWN)
)

// GamepadAxis is the list of axes available on a gamepad.
//
// Defined in <SDL3/SDL_gamepad.h>
//
// Thumbstick axis values range from SDL_JOYSTICK_AXIS_MIN to
// SDL_JOYSTICK_AXIS_MAX, and are centered within ~8000 of zero, though
// advanced UI will allow users to set or autodetect the dead zone, which
// varies between gamepads.
//
// Trigger axis values range from 0 (released) to SDL_JOYSTICK_AXIS_MAX (fully
// pressed) when reported by SDL_GetGamepadAxis(). Note that this is not the
// same range that will be reported by the lower-level SDL_GetJoystickAxis().
//
// [<SDL3/SDL_gamepad.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_gamepad.h
type GamepadAxis uint8

const (
	GAMEPAD_AXIS_INVALID GamepadAxis = math.MaxUint8
	GAMEPAD_AXIS_LEFTX   GamepadAxis = iota
	GAMEPAD_AXIS_LEFTY
	GAMEPAD_AXIS_RIGHTX
	GAMEPAD_AXIS_RIGHTY
	GAMEPAD_AXIS_LEFT_TRIGGER
	GAMEPAD_AXIS_RIGHT_TRIGGER
	GAMEPAD_AXIS_COUNT
)

// GamepadButton is the list of buttons available on a gamepad.
//
// Defined in <SDL3/SDL_gamepad.h>
//
// For controllers that use a diamond pattern for the face buttons, the
// south/east/west/north buttons below correspond to the locations in the
// diamond pattern. For Xbox controllers, this would be A/B/X/Y, for Nintendo
// Switch controllers, this would be B/A/Y/X, for GameCube controllers this
// would be A/X/B/Y, for PlayStation controllers this would be
// Cross/Circle/Square/Triangle.
//
// For controllers that don't use a diamond pattern for the face buttons, the
// south/east/west/north buttons indicate the buttons labeled A, B, C, D, or 1,
// 2, 3, 4, or for controllers that aren't labeled, they are the primary,
// secondary, etc. buttons.
//
// The activate action is often the south button and the cancel action is often
// the east button, but in some regions this is reversed, so your game should
// allow remapping actions based on user preferences.
//
// You can query the labels for the face buttons using
// SDL_GetGamepadButtonLabel()
//
// This enum is available since SDL 3.2.0.
//
// [<SDL3/SDL_gamepad.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_gamepad.h
type GamepadButton uint8

const (
	GAMEPAD_BUTTON_INVALID GamepadButton = math.MaxUint8
	// Bottom face button (e.g. Xbox A button).
	GAMEPAD_BUTTON_SOUTH GamepadButton = iota
	// Right face button (e.g. Xbox B button).
	GAMEPAD_BUTTON_EAST
	// Left face button (e.g. Xbox X button).
	GAMEPAD_BUTTON_WEST
	// Top face button (e.g. Xbox Y button).
	GAMEPAD_BUTTON_NORTH
	GAMEPAD_BUTTON_BACK
	GAMEPAD_BUTTON_GUIDE
	GAMEPAD_BUTTON_START
	GAMEPAD_BUTTON_LEFT_STICK
	GAMEPAD_BUTTON_RIGHT_STICK
	GAMEPAD_BUTTON_LEFT_SHOULDER
	GAMEPAD_BUTTON_RIGHT_SHOULDER
	GAMEPAD_BUTTON_DPAD_UP
	GAMEPAD_BUTTON_DPAD_DOWN
	GAMEPAD_BUTTON_DPAD_LEFT
	GAMEPAD_BUTTON_DPAD_RIGHT
	// Additional button (e.g. Xbox Series X share button, PS5 microphone
	// button, Nintendo Switch Pro capture button, Steam Controller QAM button,
	// Amazon Luna microphone button, Google Stadia capture button).
	GAMEPAD_BUTTON_MISC1
	// Upper or primary paddle, under your right hand (e.g. Xbox Elite paddle
	// P1, DualSense Edge RB button, Right Joy-Con SR button, Steam Controller
	// R4 button).
	GAMEPAD_BUTTON_RIGHT_PADDLE1
	// Upper or primary paddle, under your left hand (e.g. Xbox Elite paddle
	// P3, DualSense Edge LB button, Left Joy-Con SL button, Steam Controller
	// L4 button).
	GAMEPAD_BUTTON_LEFT_PADDLE1
	// Lower or secondary paddle, under your right hand (e.g. Xbox Elite paddle
	// P2, DualSense Edge right Fn button, Right Joy-Con SL button, Steam
	// Controller R5 button).
	GAMEPAD_BUTTON_RIGHT_PADDLE2
	// Lower or secondary paddle, under your left hand (e.g. Xbox Elite paddle
	// P4, DualSense Edge left Fn button, Left Joy-Con SR button, Steam
	// Controller L5 button).
	GAMEPAD_BUTTON_LEFT_PADDLE2
	// PS4/PS5 touchpad button.
	GAMEPAD_BUTTON_TOUCHPAD
	// Additional button.
	GAMEPAD_BUTTON_MISC2
	// Additional button (e.g. Nintendo GameCube left trigger click).
	GAMEPAD_BUTTON_MISC3
	// Additional button (e.g. Nintendo GameCube right trigger click).
	GAMEPAD_BUTTON_MISC4
	// Additional button.
	GAMEPAD_BUTTON_MISC5
	// Additional button.
	GAMEPAD_BUTTON_MISC6
	GAMEPAD_BUTTON_COUNT
)

type SensorType int32

const (
	// Returned for an invalid sensor.
	SENSOR_INVALID = -1 + iota
	// Unknown sensor type.
	SENSOR_UNKNOWN
	// Accelerometer.
	SENSOR_ACCEL
	// Gyroscope.
	SENSOR_GYRO
	// Accelerometer for left Joy-Con controller and Wii nunchuk.
	SENSOR_ACCEL_L
	// Gyroscope for left Joy-Con controller.
	SENSOR_GYRO_L
	// Accelerometer for right Joy-Con controller.
	SENSOR_ACCEL_R
	// Gyroscope for right Joy-Con controller.
	SENSOR_GYRO_R
	SENSOR_COUNT
)
