package sdl

import "math"

// WindowID is a unique ID for a window.
//
// Defined in [<SDL3/SDL_video.h>].
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_video.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_video.h
type WindowID uint32

// KeyboardID is a unique ID for a keyboard for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// Defined in [<SDL3/SDL_keyboard.h>].
//
// If the keyboard is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_keyboard.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_keyboard.h
type KeyboardID uint32

// MouseID is a unique ID for a mouse for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// Defined in [<SDL3/SDL_mouse.h>].
//
// If the mouse is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_mouse.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_mouse.h
type MouseID uint32

// AudioDeviceID is an SDL Audio Device instance IDs.
//
// Defined in [<SDL3/SDL_audio.h>].
//
// Zero is used to signify an invalid/null device.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_audio.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_audio.h
type AudioDeviceID uint32

// CameraID is a unique ID for a camera device for the time it is connected to
// the system, and is never reused for the lifetime of the application.
//
// Defined in [<SDL3/SDL_camera.h>].
//
// If the device is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_camera.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_camera.h
type CameraID uint32

// DisplayID is a unique ID for a display for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// Defined in [<SDL3/SDL_video.h>].
//
// If the display is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_video.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_video.h
type DisplayID uint32

// JoystickID is a unique ID for a joystick for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// Defined in [<SDL3/SDL_joystick.h>].
//
// If the joystick is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_joystick.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_joystick.h
type JoystickID uint32

// PenID is an SDL pen instance ID.
//
// Defined in [<SDL3/SDL_pen.h>].
//
// Zero is used to signify an invalid/null device.
//
// These show up in pen events when SDL sees input from them. They remain
// consistent as long as SDL can recognize a tool to be the same pen; but if a
// pen's digitizer table is physically detached from the computer, it might get
// a new ID when reconnected, as SDL won't know it's the same device.
//
// These IDs are only stable within a single run of a program; the next time a
// program is run, the pen's ID will likely be different, even if the hardware
// hasn't been disconnected, etc.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_pen.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_pen.h
type PenID uint32

// SensorID is a unique ID for a sensor for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// Defined in [<SDL3/SDL_sensor.h>].
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_sensor.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_sensor.h
type SensorID uint32

// TouchID is a unique ID for a touch device.
//
// Defined in [<SDL3/SDL_touch.h>].
//
// This ID is valid for the time the device is connected to the system, and is
// never reused for the lifetime of the application.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_touch.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_touch.h
type TouchID uint64

// FingerID is a unique ID for a single finger on a touch device.
//
// Defined in [<SDL3/SDL_touch.h>].
//
// This ID is valid for the time the finger (stylus, etc) is touching and will
// be unique for all fingers currently in contact, so this ID tracks the
// lifetime of a single continuous touch. This value may represent an index, a
// pointer, or some other unique ID, depending on the platform.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_touch.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_touch.h
type FingerID uint64

// InitFlags are initialization flags for [Init] and/or [InitSubSystem].
//
// Defined in [<SDL3/SDL_init.h>]
//
// These are the flags which may be passed to [Init](). You should specify the
// subsystems which you will be using in your application.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_init.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_init.h
type InitFlags uint32

// Values for initialization flags.
const (
	// InitAudio implies [InitEvents].
	InitAudio InitFlags = 0x00000010

	// InitVideo implies [InitEvents], should be initialized on the main
	// thread.
	InitVideo InitFlags = 0x00000020

	// InitJoystick implies [InitEvents].
	InitJoystick InitFlags = 0x00000200

	InitHaptic InitFlags = 0x00001000

	// InitGamepad implies [InitJoystick].
	InitGamepad InitFlags = 0x00002000

	InitEvents InitFlags = 0x00004000

	// InitSensor implies [InitEvents].
	InitSensor InitFlags = 0x00008000

	// InitCamera implies [InitEvents].
	InitCamera InitFlags = 0x00010000
)

// WindowFlags are the flags on a window.
//
// Defined in [<SDL3/SDL_video.h>].
//
// These cover a lot of true/false, or on/off, window state. Some of it is
// immutable after being set through [CreateWindow](), some of it can be
// changed on existing windows by the app, and some of it might be altered by
// the user or system outside of the app's control.
//
// When creating windows with [WindowResizable], SDL will constrain resizable
// windows to the dimensions recommended by the compositor to fit it within the
// usable desktop space, although some compositors will do this automatically
// without intervention as well. Use [SetWindowResizable] after creation
// instead if you wish to create a window with a specific size.
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_video.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_video.h
type WindowFlags uint64

const (
	// WindowFullscreen indicates window is in fullscreen mode.
	WindowFullscreen WindowFlags = 0x0000000000000001

	// WindowOpenGL indicates window usable with OpenGL context.
	WindowOpenGL WindowFlags = 0x0000000000000002

	// WindowOccluded indicates window is occluded.
	WindowOccluded WindowFlags = 0x0000000000000004

	// WindowHidden indicates window is neither mapped onto the desktop nor
	// shown in the taskbar/dock/window list; [ShowWindow] is required for it
	// to become visible.
	WindowHidden WindowFlags = 0x0000000000000008

	// WindowBorderless indicates no window decoration.
	WindowBorderless WindowFlags = 0x0000000000000010

	// WindowResizable indicates window can be resized.
	WindowResizable WindowFlags = 0x0000000000000020

	// WindowMinimized indicates window is minimized.
	WindowMinimized WindowFlags = 0x0000000000000040

	// WindowMaximized indicates window is maximized.
	WindowMaximized WindowFlags = 0x0000000000000080

	// WindowMouseGrabbed indicates window has grabbed mouse input.
	WindowMouseGrabbed WindowFlags = 0x0000000000000100

	// WindowInputFocus indicates window has input focus.
	WindowInputFocus WindowFlags = 0x0000000000000200

	// WindowMouseFocus indicates window has mouse focus.
	WindowMouseFocus WindowFlags = 0x0000000000000400

	// WindowExternal indicates window not created by SDL.
	WindowExternal WindowFlags = 0x0000000000000800

	// WindowModal indicates window is modal.
	WindowModal WindowFlags = 0x0000000000001000

	// WindowHighPixelDensity indicates window uses high pixel density back
	// buffer if possible.
	WindowHighPixelDensity WindowFlags = 0x0000000000002000

	// WindowMouseCapture indicates window has mouse captured (unrelated to
	// MOUSE_GRABBED).
	WindowMouseCapture WindowFlags = 0x0000000000004000

	// WindowMouseRelativeMode indicates window has relative mode enabled.
	WindowMouseRelativeMode WindowFlags = 0x0000000000008000

	// WindowAlwaysOnTop indicates window should always be above others.
	WindowAlwaysOnTop WindowFlags = 0x0000000000010000

	// WindowUtility indicates window should be treated as a utility window,
	// not showing in the task bar and window list.
	WindowUtility WindowFlags = 0x0000000000020000

	// WindowTooltip indicates window should be treated as a tooltip and does
	// not get mouse or keyboard focus, requires a parent window.
	WindowTooltip WindowFlags = 0x0000000000040000

	// WindowPopupMenu indicates window should be treated as a popup menu,
	// requires a parent window.
	WindowPopupMenu WindowFlags = 0x0000000000080000

	// WindowKeyboardGrabbed indicates window has grabbed keyboard input.
	WindowKeyboardGrabbed WindowFlags = 0x0000000000100000

	// WindowFillDocument indicates window is in fill-document mode (Emscripten
	// only, since SDL 3.4.0).
	WindowFillDocument WindowFlags = 0x0000000000200000

	// WindowVulkan indicates window usable for Vulkan surface.
	WindowVulkan WindowFlags = 0x0000000010000000

	// WindowMetal indicates window usable for Metal view.
	WindowMetal WindowFlags = 0x0000000020000000

	// WindowTransparent indicates window with transparent buffer.
	WindowTransparent WindowFlags = 0x0000000040000000

	// WindowNotFocusable indicates window should not be focusable.
	WindowNotFocusable WindowFlags = 0x0000000080000000
)

// EventType holds the types of events that can be delivered.
//
// Defined in [<SDL3/SDL_events.h>].
//
// This enum is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type EventType uint32

// Application events.
const (
	// EventQuit indicates a user-requested quit.
	EventQuit EventType = 0x100 + iota

	// These application events have special meaning on iOS and Android, see
	// README-ios.md and README-android.md for details.

	// EventTerminating indicates the application is being terminated by the
	// OS. This event must be handled in a callback set with [AddEventWatch].
	// Called on iOS in applicationWillTerminate(). Called on Android in
	// onDestroy().
	EventTerminating

	// EventLowMemory indicates the application is low on memory, free memory
	// if possible. This event must be handled in a callback set with
	// [AddEventWatch]. Called on iOS in applicationDidReceiveMemoryWarning().
	// Called on Android in onTrimMemory().
	EventLowMemory

	// EventWillEnterBackground indicates the application is about to enter the
	// background. This event must be handled in a callback set with
	// [AddEventWatch]. Called on iOS in applicationWillResignActive(). Called
	// on Android in onPause().
	EventWillEnterBackground

	// EventDidEnterBackground indicates the application did enter the
	// background and may not get CPU for some time. This event must be handled
	// in a callback set with [AddEventWatch]. Called on iOS in
	// applicationDidEnterBackground(). Called on Android in onPause().
	EventDidEnterBackground

	// EventWillEnterForeground indicates the application is about to enter the
	// foreground. This event must be handled in a callback set with
	// [AddEventWatch]. Called on iOS in applicationWillEnterForeground().
	// Called on Android in onResume().
	EventWillEnterForeground

	// EventDidEnterForeground indicates the application is now interactive.
	// This event must be handled in a callback set with [AddEventWatch].
	// Called on iOS in applicationDidBecomeActive(). Called on Android in
	// onResume().
	EventDidEnterForeground

	// EventLocaleChanged indicates the user's locale preferences have changed.
	EventLocaleChanged

	// EventSystemThemeChanged indicates the system theme changed.
	EventSystemThemeChanged
)

// Display events.
const (
	// 0x150 was SDL_DISPLAYEVENT, reserve the number for sdl2-compat.

	// EventDisplayOrientation indicates display orientation has changed to
	// [DisplayEvent.Data1].
	EventDisplayOrientation EventType = 0x151 + iota

	// EventDisplayAdded indicates display has been added to the system.
	EventDisplayAdded

	// EventDisplayRemoved indicates display has been removed from the system.
	EventDisplayRemoved

	// EventDisplayMoved indicates display has changed position.
	EventDisplayMoved

	// EventDisplayDesktopModeChanged indicates display has changed desktop
	// mode.
	EventDisplayDesktopModeChanged

	// EventDisplayCurrentModeChanged indicates display has changed current
	// mode.
	EventDisplayCurrentModeChanged

	// EventDisplayContentScaleChanged indicates display has changed content
	// scale.
	EventDisplayContentScaleChanged

	// EventDisplayUsableBoundsChanged indicates display has changed usable
	// bounds.
	EventDisplayUsableBoundsChanged

	EventDisplayFirst = EventDisplayOrientation
	EventDisplayLast  = EventDisplayUsableBoundsChanged
)

// Window events.
const (
	// 0x200 was SDL_WINDOWEVENT, reserve the number for sdl2-compat
	// 0x201 was SDL_SYSWMEVENT, reserve the number for sdl2-compat.

	// EventWindowShown indicates window has been shown.
	EventWindowShown EventType = 0x202 + iota

	// EventWindowHidden indicates window has been hidden.
	EventWindowHidden

	// EventWindowExposed indicates window has been exposed and should be
	// redrawn, and can be redrawn directly from event watchers for this event.
	// data1 is 1 for live-resize expose events, 0 otherwise..
	EventWindowExposed

	// EventWindowMoved indicates window has been moved to data1, data2.
	EventWindowMoved

	// EventWindowResized indicates window has been resized to data1xdata2.
	EventWindowResized

	// EventWindowPixelSizeChanged indicates the pixel size of the window has
	// changed to data1xdata2.
	EventWindowPixelSizeChanged

	// EventWindowMetalViewResized indicates the pixel size of a Metal view
	// associated with the window has changed.
	EventWindowMetalViewResized

	// EventWindowMinimized indicates window has been minimized.
	EventWindowMinimized

	// EventWindowMaximized indicates window has been maximized.
	EventWindowMaximized

	// EventWindowRestored indicates window has been restored to normal size
	// and position.
	EventWindowRestored

	// EventWindowMouseEnter indicates window has gained mouse focus.
	EventWindowMouseEnter

	// EventWindowMouseLeave indicates window has lost mouse focus.
	EventWindowMouseLeave

	// EventWindowFocusGained indicates window has gained keyboard focus.
	EventWindowFocusGained

	// EventWindowFocusLost indicates window has lost keyboard focus.
	EventWindowFocusLost

	// EventWindowCloseRequested indicates the window manager requests that the
	// window be closed.
	EventWindowCloseRequested

	// EventWindowHitTest indicates window had a hit test that wasn't
	// SDL_HITTEST_NORMAL.
	EventWindowHitTest

	// EventWindowIccprofChanged indicates the window's ICC profile has
	// changed.
	EventWindowIccprofChanged

	// EventWindowDisplayChanged indicates window has been moved to display
	// data1.
	EventWindowDisplayChanged

	// EventWindowDisplayScaleChanged indicates window display scale has been
	// changed.
	EventWindowDisplayScaleChanged

	// EventWindowSafeAreaChanged indicates the window safe area has been
	// changed.
	EventWindowSafeAreaChanged

	// EventWindowOccluded indicates the window has been occluded.
	EventWindowOccluded

	// EventWindowEnterFullscreen indicates the window has entered fullscreen
	// mode.
	EventWindowEnterFullscreen

	// EventWindowLeaveFullscreen indicates the window has left fullscreen
	// mode.
	EventWindowLeaveFullscreen

	// EventWindowDestroyed indicates the window with the associated ID is
	// being or has been destroyed. If this message is being handled in an
	// event watcher, the window handle is still valid and can still be used to
	// retrieve any properties associated with the window. Otherwise, the
	// handle has already been destroyed and all resources associated with it
	// are invalid.
	EventWindowDestroyed

	// EventWindowHdrStateChanged indicates window HDR properties have changed.
	EventWindowHdrStateChanged

	// EventWindowSettingsChanged indicates window settings have changed (on
	// visionOS).
	EventWindowSettingsChanged

	EventWindowFirst = EventWindowShown
	EventWindowLast  = EventWindowSettingsChanged
)

// Keyboard events.
const (
	// EventKeyDown indicates key pressed.
	EventKeyDown EventType = 0x300 + iota

	// EventKeyUp indicates key released.
	EventKeyUp

	// EventTextEditing indicates Keyboard text editing (composition).
	EventTextEditing

	// EventTextInput indicates Keyboard text input.
	EventTextInput

	// EventKeymapChanged indicates Keymap changed due to a system event such
	// as an input language or keyboard layout change..
	EventKeymapChanged

	// EventKeyboardAdded indicates A new keyboard has been inserted into the
	// system.
	EventKeyboardAdded

	// EventKeyboardRemoved indicates A keyboard has been removed.
	EventKeyboardRemoved

	// EventTextEditingCandidates indicates Keyboard text editing candidates.
	EventTextEditingCandidates

	// EventScreenKeyboardShown indicates the on-screen keyboard has been
	// shown.
	EventScreenKeyboardShown

	// EventScreenKeyboardHidden indicates the on-screen keyboard has been
	// hidden.
	EventScreenKeyboardHidden
)

// Mouse events.
const (
	// EventMouseMotion indicates mouse moved.
	EventMouseMotion EventType = 0x400 + iota

	// EventMouseButtonDown indicates mouse button pressed.
	EventMouseButtonDown

	// EventMouseButtonUp indicates mouse button released.
	EventMouseButtonUp

	// EventMouseWheel indicates mouse wheel motion.
	EventMouseWheel

	// EventMouseAdded indicates a new mouse has been inserted into the system.
	EventMouseAdded

	// EventMouseRemoved indicates a mouse has been removed.
	EventMouseRemoved
)

// Joystick events.
const (
	// EventJoystickAxisMotion indicates joystick axis motion.
	EventJoystickAxisMotion EventType = 0x600 + iota

	// EventJoystickBallMotion indicates joystick trackball motion.
	EventJoystickBallMotion

	// EventJoystickHatMotion indicates joystick hat position change.
	EventJoystickHatMotion

	// EventJoystickButtonDown indicates joystick button pressed.
	EventJoystickButtonDown

	// EventJoystickButtonUp indicates joystick button released.
	EventJoystickButtonUp

	// EventJoystickAdded indicates a new joystick has been inserted into the
	// system.
	EventJoystickAdded

	// EventJoystickRemoved indicates an opened joystick has been removed.
	EventJoystickRemoved

	// EventJoystickBatteryUpdated indicates joystick battery level change.
	EventJoystickBatteryUpdated

	// EventJoystickUpdateComplete indicates joystick update is complete.
	EventJoystickUpdateComplete
)

// Gamepad events.
const (
	// EventGamepadAxisMotion indicates gamepad axis motion.
	EventGamepadAxisMotion EventType = 0x650 + iota

	// EventGamepadButtonDown indicates gamepad button pressed.
	EventGamepadButtonDown

	// EventGamepadButtonUp indicates gamepad button released.
	EventGamepadButtonUp

	// EventGamepadAdded indicates a new gamepad has been inserted into the
	// system.
	EventGamepadAdded

	// EventGamepadRemoved indicates a gamepad has been removed.
	EventGamepadRemoved

	// EventGamepadRemapped indicates the gamepad mapping was updated.
	EventGamepadRemapped

	// EventGamepadTouchpadDown indicates gamepad touchpad was touched.
	EventGamepadTouchpadDown

	// EventGamepadTouchpadMotion indicates gamepad touchpad finger was moved.
	EventGamepadTouchpadMotion

	// EventGamepadTouchpadUp indicates gamepad touchpad finger was lifted.
	EventGamepadTouchpadUp

	// EventGamepadSensorUpdate indicates gamepad sensor was updated.
	EventGamepadSensorUpdate

	// EventGamepadUpdateComplete indicates gamepad update is complete.
	EventGamepadUpdateComplete

	// EventGamepadSteamHandleUpdated indicates gamepad Steam handle has
	// changed.
	EventGamepadSteamHandleUpdated

	// EventGamepadCapsenseTouch indicates gamepad capsense was touched.
	EventGamepadCapsenseTouch

	// EventGamepadCapsenseRelease indicates gamepad capsense was released.
	EventGamepadCapsenseRelease
)

// Touch events.
const (
	EventFingerDown EventType = 0x700 + iota
	EventFingerUp
	EventFingerMotion
	EventFingerCanceled
)

// Pinch events.
const (
	// EventPinchBegin indicates pinch gesture started.
	EventPinchBegin EventType = 0x710 + iota

	// EventPinchUpdate indicates pinch gesture updated.
	EventPinchUpdate

	// EventPinchEnd indicates pinch gesture ended.
	EventPinchEnd
)

// 0x800, 0x801, and 0x802 were the Gesture events from SDL2. Do not reuse
// these values! sdl2-compat needs them!

// Clipboard events.
const (
	// EventClipboardUpdate indicates the clipboard changed.
	EventClipboardUpdate EventType = 0x900 + iota
)

// Drag and drop events.
const (
	// EventDropFile indicates the system requests a file open.
	EventDropFile EventType = 0x1000 + iota

	// EventDropText indicates text/plain drag-and-drop event.
	EventDropText

	// EventDropBegin indicates a new set of drops is beginning (NULL filename).
	EventDropBegin

	// EventDropComplete indicates current set of drops is now complete (NULL
	// filename).
	EventDropComplete

	// EventDropPosition indicates position while moving over the window.
	EventDropPosition
)

// Audio hotplug events.
const (
	// EventAudioDeviceAdded indicates a new audio device is available.
	EventAudioDeviceAdded EventType = 0x1100 + iota

	// EventAudioDeviceRemoved indicates an audio device has been removed.
	EventAudioDeviceRemoved

	// EventAudioDeviceFormatChanged indicates an audio device's format has
	// been changed by the system.
	EventAudioDeviceFormatChanged
)

// Sensor events.
const (
	// EventSensorUpdate indicates a sensor was updated.
	EventSensorUpdate EventType = 0x1200 + iota
)

// Pressure-sensitive pen events.
const (
	// EventPenProximityIn indicates pressure-sensitive pen has become
	// available.
	EventPenProximityIn EventType = 0x1300 + iota

	// EventPenProximityOut indicates pressure-sensitive pen has become
	// unavailable.
	EventPenProximityOut

	// EventPenDown indicates pressure-sensitive pen touched drawing surface.
	EventPenDown

	// EventPenUp indicates pressure-sensitive pen stopped touching drawing
	// surface.
	EventPenUp

	// EventPenButtonDown indicates pressure-sensitive pen button pressed.
	EventPenButtonDown

	// EventPenButtonUp indicates pressure-sensitive pen button released.
	EventPenButtonUp

	// EventPenMotion indicates pressure-sensitive pen is moving on the tablet.
	EventPenMotion

	// EventPenAxis indicates pressure-sensitive pen angle/pressure/etc
	// changed.
	EventPenAxis
)

// Camera hotplug events.
const (
	// EventCameraDeviceAdded indicates a new camera device is available.
	EventCameraDeviceAdded EventType = 0x1400 + iota

	// EventCameraDeviceRemoved indicates a camera device has been removed.
	EventCameraDeviceRemoved

	// EventCameraDeviceApproved indicates a camera device has been approved
	// for use by the user.
	EventCameraDeviceApproved

	// EventCameraDeviceDenied indicates a camera device has been denied for
	// use by the user.
	EventCameraDeviceDenied
)

// Notification events.
const (
	// EventNotificationActionInvoked indicates a user response to a system
	// notification was received.
	EventNotificationActionInvoked EventType = 0x1500 + iota
)

// Render events.
const (
	// EventRenderTargetsReset indicates the render targets have been reset and
	// their contents need to be updated.
	EventRenderTargetsReset EventType = 0x2000 + iota

	// EventRenderDeviceReset indicates the device has been reset and all
	// textures need to be recreated.
	EventRenderDeviceReset

	// EventRenderDeviceLost indicates the device has been lost and can't be
	// recovered.
	EventRenderDeviceLost
)

// Reserved events for private platforms.
const (
	EventPrivate0 EventType = 0x4000 + iota
	EventPrivate1
	EventPrivate2
	EventPrivate3
)

// Internal events.
const (
	// EventPollSentinel indicates signals the end of an event poll cycle.
	EventPollSentinel EventType = 0x7F00 + iota
)

// Events SDL_EVENT_USER through SDL_EVENT_LAST are for your use, and should be
// allocated with SDL_RegisterEvents().
const (
	EventUser EventType = 0x8000

	// EventLast is only for bounding internal arrays.
	EventLast EventType = 0xFFFF
)

// Scancode is the SDL keyboard scancode representation.
//
// Defined in [<SDL3/SDL_scancode.h>].
//
// An SDL scancode is the physical representation of a key on the keyboard,
// independent of language and keyboard mapping.
//
// Values of this type are used to represent keyboard keys, among other places
// in the scancode field of the [KeyboardEvent] structure.
//
// The values in this enumeration are based on the USB usage page standard:
// https://usb.org/sites/default/files/hut1_5.pdf.
//
// This enum is available since SDL 3.2.0.
//
// [<SDL3/SDL_scancode.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_scancode.h
type Scancode uint32

// Scancode values.
const (
	ScancodeUnknown Scancode = 0

	// These values are from usage page 0x07 (USB keyboard page).

	ScancodeA Scancode = 4
	ScancodeB Scancode = 5
	ScancodeC Scancode = 6
	ScancodeD Scancode = 7
	ScancodeE Scancode = 8
	ScancodeF Scancode = 9
	ScancodeG Scancode = 10
	ScancodeH Scancode = 11
	ScancodeI Scancode = 12
	ScancodeJ Scancode = 13
	ScancodeK Scancode = 14
	ScancodeL Scancode = 15
	ScancodeM Scancode = 16
	ScancodeN Scancode = 17
	ScancodeO Scancode = 18
	ScancodeP Scancode = 19
	ScancodeQ Scancode = 20
	ScancodeR Scancode = 21
	ScancodeS Scancode = 22
	ScancodeT Scancode = 23
	ScancodeU Scancode = 24
	ScancodeV Scancode = 25
	ScancodeW Scancode = 26
	ScancodeX Scancode = 27
	ScancodeY Scancode = 28
	ScancodeZ Scancode = 29

	Scancode1 Scancode = 30
	Scancode2 Scancode = 31
	Scancode3 Scancode = 32
	Scancode4 Scancode = 33
	Scancode5 Scancode = 34
	Scancode6 Scancode = 35
	Scancode7 Scancode = 36
	Scancode8 Scancode = 37
	Scancode9 Scancode = 38
	Scancode0 Scancode = 39

	ScancodeReturn    Scancode = 40
	ScancodeEscape    Scancode = 41
	ScancodeBackspace Scancode = 42
	ScancodeTab       Scancode = 43
	ScancodeSpace     Scancode = 44

	ScancodeMinus        Scancode = 45
	ScancodeEquals       Scancode = 46
	ScancodeLeftbracket  Scancode = 47
	ScancodeRightbracket Scancode = 48
	// ScancodeBackslash is located at the lower left of the return key on ISO
	// keyboards and at the right end of the QWERTY row on ANSI keyboards.
	// Produces REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US layout,
	// REVERSE SOLIDUS and VERTICAL LINE in a UK Mac layout, NUMBER SIGN and
	// TILDE in a UK Windows layout, DOLLAR SIGN and POUND SIGN in a Swiss
	// German layout, NUMBER SIGN and APOSTROPHE in a German layout, GRAVE
	// ACCENT and POUND SIGN in a French Mac layout, and ASTERISK and MICRO
	// SIGN in a French Windows layout.
	ScancodeBackslash Scancode = 49
	// ScancodeNonushash is actually used by ISO USB keyboards instead of 49
	// for the same key, but all OSes I've seen treat the two codes
	// identically. So, as an implementer, unless your keyboard generates both
	// of those codes and your OS treats them differently, you should generate
	// SDL_SCANCODE_BACKSLASH instead of this code. As a user, you should not
	// rely on this code because SDL will never generate it with most (all?)
	// keyboards.
	ScancodeNonushash  Scancode = 50
	ScancodeSemicolon  Scancode = 51
	ScancodeApostrophe Scancode = 52
	// ScancodeGrave is located in the top left corner (on both ANSI and ISO
	// keyboards).  Produces GRAVE ACCENT and TILDE in a US Windows layout and
	// in US and UK Mac layouts on ANSI keyboards, GRAVE ACCENT and NOT SIGN in
	// a UK Windows layout, SECTION SIGN and PLUS-MINUS SIGN in US and UK Mac
	// layouts on ISO keyboards, SECTION SIGN and DEGREE SIGN in a Swiss German
	// layout (Mac: only on ISO keyboards), CIRCUMFLEX ACCENT and DEGREE SIGN
	// in a German layout (Mac: only on ISO keyboards), SUPERSCRIPT TWO and
	// TILDE in a French Windows layout, COMMERCIAL AT and NUMBER SIGN in a
	// French Mac layout on ISO keyboards, and LESS-THAN SIGN and GREATER-THAN
	// SIGN in a Swiss German, German, or French Mac layout on ANSI keyboards.
	ScancodeGrave  Scancode = 53
	ScancodeComma  Scancode = 54
	ScancodePeriod Scancode = 55
	ScancodeSlash  Scancode = 56

	ScancodeCapslock Scancode = 57

	ScancodeF1  Scancode = 58
	ScancodeF2  Scancode = 59
	ScancodeF3  Scancode = 60
	ScancodeF4  Scancode = 61
	ScancodeF5  Scancode = 62
	ScancodeF6  Scancode = 63
	ScancodeF7  Scancode = 64
	ScancodeF8  Scancode = 65
	ScancodeF9  Scancode = 66
	ScancodeF10 Scancode = 67
	ScancodeF11 Scancode = 68
	ScancodeF12 Scancode = 69

	ScancodePrintscreen Scancode = 70
	ScancodeScrolllock  Scancode = 71
	ScancodePause       Scancode = 72
	// ScancodeInsert is insert on PC, help on some Mac keyboards (but does
	// send code 73, not 117).
	ScancodeInsert   Scancode = 73
	ScancodeHome     Scancode = 74
	ScancodePageup   Scancode = 75
	ScancodeDelete   Scancode = 76
	ScancodeEnd      Scancode = 77
	ScancodePagedown Scancode = 78
	ScancodeRight    Scancode = 79
	ScancodeLeft     Scancode = 80
	ScancodeDown     Scancode = 81
	ScancodeUp       Scancode = 82

	// ScancodeNumlockclear is num lock on PC, clear on Mac keyboards.
	ScancodeNumlockclear Scancode = 83
	ScancodeKPDivide     Scancode = 84
	ScancodeKPMultiply   Scancode = 85
	ScancodeKPMinus      Scancode = 86
	ScancodeKPPlus       Scancode = 87
	ScancodeKPEnter      Scancode = 88
	ScancodeKP1          Scancode = 89
	ScancodeKP2          Scancode = 90
	ScancodeKP3          Scancode = 91
	ScancodeKP4          Scancode = 92
	ScancodeKP5          Scancode = 93
	ScancodeKP6          Scancode = 94
	ScancodeKP7          Scancode = 95
	ScancodeKP8          Scancode = 96
	ScancodeKP9          Scancode = 97
	ScancodeKP0          Scancode = 98
	ScancodeKPPeriod     Scancode = 99

	// ScancodeNonusbackslash is the additional key that ISO keyboards have
	// over ANSI ones, located between left shift and Z. Produces GRAVE ACCENT
	// and TILDE in a US or UK Mac layout, REVERSE SOLIDUS (backslash) and
	// VERTICAL LINE in a US or UK Windows layout, and LESS-THAN SIGN and
	// GREATER-THAN SIGN in a Swiss German, German, or French layout.
	ScancodeNonusbackslash Scancode = 100
	ScancodeApplication    Scancode = 101 // windows contextual menu, compose
	// ScancodePower is a status flag, not a physical key, the USB document
	// says, but some Mac keyboards do have a power key.
	ScancodePower      Scancode = 102
	ScancodeKPEquals   Scancode = 103
	ScancodeF13        Scancode = 104
	ScancodeF14        Scancode = 105
	ScancodeF15        Scancode = 106
	ScancodeF16        Scancode = 107
	ScancodeF17        Scancode = 108
	ScancodeF18        Scancode = 109
	ScancodeF19        Scancode = 110
	ScancodeF20        Scancode = 111
	ScancodeF21        Scancode = 112
	ScancodeF22        Scancode = 113
	ScancodeF23        Scancode = 114
	ScancodeF24        Scancode = 115
	ScancodeExecute    Scancode = 116
	ScancodeHelp       Scancode = 117 // AL Integrated Help Center
	ScancodeMenu       Scancode = 118 // Menu (show menu)
	ScancodeSelect     Scancode = 119
	ScancodeStop       Scancode = 120 // AC Stop
	ScancodeAgain      Scancode = 121 // AC Redo/Repeat
	ScancodeUndo       Scancode = 122 // AC Undo
	ScancodeCut        Scancode = 123 // AC Cut
	ScancodeCopy       Scancode = 124 // AC Copy
	ScancodePaste      Scancode = 125 // AC Paste
	ScancodeFind       Scancode = 126 // AC Find
	ScancodeMute       Scancode = 127
	ScancodeVolumeup   Scancode = 128
	ScancodeVolumedown Scancode = 129

	// not sure whether there's a reason to enable these
	//
	//	SDL_SCANCODE_LOCKINGCAPSLOCK = 130,
	//	SDL_SCANCODE_LOCKINGNUMLOCK = 131,
	//	SDL_SCANCODE_LOCKINGSCROLLLOCK = 132,

	ScancodeKPComma       Scancode = 133
	ScancodeKPEqualsas400 Scancode = 134

	// ScancodeInternational1 is used on Asian keyboards, see footnotes in USB
	// doc.
	ScancodeInternational1 Scancode = 135

	ScancodeInternational2 Scancode = 136
	ScancodeInternational3 Scancode = 137 // Yen
	ScancodeInternational4 Scancode = 138
	ScancodeInternational5 Scancode = 139
	ScancodeInternational6 Scancode = 140
	ScancodeInternational7 Scancode = 141
	ScancodeInternational8 Scancode = 142
	ScancodeInternational9 Scancode = 143
	ScancodeLang1          Scancode = 144 // Hangul/English toggle
	ScancodeLang2          Scancode = 145 // Hanja conversion
	ScancodeLang3          Scancode = 146 // Katakana
	ScancodeLang4          Scancode = 147 // Hiragana
	ScancodeLang5          Scancode = 148 // Zenkaku/Hankaku
	ScancodeLang6          Scancode = 149 // reserved
	ScancodeLang7          Scancode = 150 // reserved
	ScancodeLang8          Scancode = 151 // reserved
	ScancodeLang9          Scancode = 152 // reserved

	ScancodeAlterase   Scancode = 153 // Erase-Eaze
	ScancodeSysreq     Scancode = 154
	ScancodeCancel     Scancode = 155 // AC Cancel
	ScancodeClear      Scancode = 156
	ScancodePrior      Scancode = 157
	ScancodeReturn2    Scancode = 158
	ScancodeSeparator  Scancode = 159
	ScancodeOut        Scancode = 160
	ScancodeOper       Scancode = 161
	ScancodeClearagain Scancode = 162
	ScancodeCrsel      Scancode = 163
	ScancodeExsel      Scancode = 164

	ScancodeKP00               Scancode = 176
	ScancodeKP000              Scancode = 177
	ScancodeThousandsseparator Scancode = 178
	ScancodeDecimalseparator   Scancode = 179
	ScancodeCurrencyunit       Scancode = 180
	ScancodeCurrencysubunit    Scancode = 181
	ScancodeKPLeftparen        Scancode = 182
	ScancodeKPRightparen       Scancode = 183
	ScancodeKPLeftbrace        Scancode = 184
	ScancodeKPRightbrace       Scancode = 185
	ScancodeKPTab              Scancode = 186
	ScancodeKPBackspace        Scancode = 187
	ScancodeKPA                Scancode = 188
	ScancodeKPB                Scancode = 189
	ScancodeKPC                Scancode = 190
	ScancodeKPD                Scancode = 191
	ScancodeKPE                Scancode = 192
	ScancodeKPF                Scancode = 193
	ScancodeKPXor              Scancode = 194
	ScancodeKPPower            Scancode = 195
	ScancodeKPPercent          Scancode = 196
	ScancodeKPLess             Scancode = 197
	ScancodeKPGreater          Scancode = 198
	ScancodeKPAmpersand        Scancode = 199
	ScancodeKPDblampersand     Scancode = 200
	ScancodeKPVerticalbar      Scancode = 201
	ScancodeKPDblverticalbar   Scancode = 202
	ScancodeKPColon            Scancode = 203
	ScancodeKPHash             Scancode = 204
	ScancodeKPSpace            Scancode = 205
	ScancodeKPAt               Scancode = 206
	ScancodeKPExclam           Scancode = 207
	ScancodeKPMemstore         Scancode = 208
	ScancodeKPMemrecall        Scancode = 209
	ScancodeKPMemclear         Scancode = 210
	ScancodeKPMemadd           Scancode = 211
	ScancodeKPMemsubtract      Scancode = 212
	ScancodeKPMemmultiply      Scancode = 213
	ScancodeKPMemdivide        Scancode = 214
	ScancodeKPPlusminus        Scancode = 215
	ScancodeKPClear            Scancode = 216
	ScancodeKPClearentry       Scancode = 217
	ScancodeKPBinary           Scancode = 218
	ScancodeKPOctal            Scancode = 219
	ScancodeKPDecimal          Scancode = 220
	ScancodeKPHexadecimal      Scancode = 221

	ScancodeLctrl  Scancode = 224
	ScancodeLshift Scancode = 225
	ScancodeLalt   Scancode = 226 // alt, option
	ScancodeLgui   Scancode = 227 // windows, command (apple), meta
	ScancodeRctrl  Scancode = 228
	ScancodeRshift Scancode = 229
	ScancodeRalt   Scancode = 230 // alt gr, option
	ScancodeRgui   Scancode = 231 // windows, command (apple), meta

	// I'm not sure if this is really not covered by any of the above, but since
	// there's a special SDL_KMOD_MODE for it I'm adding it here.

	ScancodeMode Scancode = 257

	// These values are mapped from usage page 0x0C (USB consumer page).
	//
	// There are way more keys in the spec than we can represent in the current
	// scancode range, so pick the ones that commonly come up in real world usage.

	ScancodeSleep Scancode = 258 // Sleep
	ScancodeWake  Scancode = 259 // Wake

	ScancodeChannelIncrement Scancode = 260 // Channel Increment
	ScancodeChannelDecrement Scancode = 261 // Channel Decrement

	ScancodeMediaPlay          Scancode = 262 // Play
	ScancodeMediaPause         Scancode = 263 // Pause
	ScancodeMediaRecord        Scancode = 264 // Record
	ScancodeMediaFastForward   Scancode = 265 // Fast Forward
	ScancodeMediaRewind        Scancode = 266 // Rewind
	ScancodeMediaNextTrack     Scancode = 267 // Next Track
	ScancodeMediaPreviousTrack Scancode = 268 // Previous Track
	ScancodeMediaStop          Scancode = 269 // Stop
	ScancodeMediaEject         Scancode = 270 // Eject
	ScancodeMediaPlayPause     Scancode = 271 // Play / Pause
	ScancodeMediaSelect        Scancode = 272 // Media Select

	ScancodeAcNew        Scancode = 273 // AC New
	ScancodeAcOpen       Scancode = 274 // AC Open
	ScancodeAcClose      Scancode = 275 // AC Close
	ScancodeAcExit       Scancode = 276 // AC Exit
	ScancodeAcSave       Scancode = 277 // AC Save
	ScancodeAcPrint      Scancode = 278 // AC Print
	ScancodeAcProperties Scancode = 279 // AC Properties

	ScancodeAcSearch    Scancode = 280 // AC Search
	ScancodeAcHome      Scancode = 281 // AC Home
	ScancodeAcBack      Scancode = 282 // AC Back
	ScancodeAcForward   Scancode = 283 // AC Forward
	ScancodeAcStop      Scancode = 284 // AC Stop
	ScancodeAcRefresh   Scancode = 285 // AC Refresh
	ScancodeAcBookmarks Scancode = 286 // AC Bookmarks

	// These are values that are often used on mobile phones.

	// ScancodeSoftleft is usually situated below the display on phones and
	// used as a multi-function feature key for selecting a software defined
	// function shown on the bottom left of the display.
	ScancodeSoftleft Scancode = 287
	// ScancodeSoftright is usually situated below the display on phones and
	// used as a multi-function feature key for selecting a software defined
	// function shown on the bottom right of the display.
	ScancodeSoftright Scancode = 288
	ScancodeCall      Scancode = 289 // Used for accepting phone calls.
	ScancodeEndcall   Scancode = 290 // Used for rejecting phone calls.

	// Add any other keys here.

	ScancodeReserved Scancode = 400 // 400-500 reserved for dynamic keycodes

	ScancodeCount Scancode = 512 // not a key, just marks the number of scancodes for array bounds
)

// Keycode is the SDL virtual key representation.
//
// Defined in [<SDL3/SDL_keycode.h>].
//
// Values of this type are used to represent keyboard keys using the current
// layout of the keyboard. These values include Unicode values representing the
// unmodified character that would be generated by pressing the key, or a K*
// constant for those keys that do not generate characters.
//
// A special exception is the number keys at the top of the keyboard which map
// by default to K0...K9 on AZERTY layouts.
//
// Keys with the SDLK_EXTENDED_MASK bit set do not map to a scancode or Unicode
// code point.
//
// Many common keycodes are listed below, but this list is not exhaustive.
//
// This datatype is available since SDL 3.2.0.
//
// See Also:
//
//   - [HINT_KEYCODE_OPTIONS]
//
// [<SDL3/SDL_keycode.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_keycode.h
type Keycode uint32

// ScancodeToKeycode converts a [Scancode] to [Keycode].
func ScancodeToKeycode(scancode Scancode) Keycode {
	return Keycode(scancode) | KScancodeMask
}

// Keycode values.
const (
	KExtendedMask       Keycode = 1 << 29
	KScancodeMask       Keycode = 1 << 30
	KUnknown            Keycode = 0x00000000 // 0
	KReturn             Keycode = 0x0000000d // '\r'
	KEscape             Keycode = 0x0000001b // '\x1B'
	KBackspace          Keycode = 0x00000008 // '\b'
	KTab                Keycode = 0x00000009 // '\t'
	KSpace              Keycode = 0x00000020 // ' '
	KExclaim            Keycode = 0x00000021 // '!'
	KDblapostrophe      Keycode = 0x00000022 // '"'
	KHash               Keycode = 0x00000023 // '#'
	KDollar             Keycode = 0x00000024 // '$'
	KPercent            Keycode = 0x00000025 // '%'
	KAmpersand          Keycode = 0x00000026 // '&'
	KApostrophe         Keycode = 0x00000027 // '\''
	KLeftparen          Keycode = 0x00000028 // '('
	KRightparen         Keycode = 0x00000029 // ')'
	KAsterisk           Keycode = 0x0000002a // '*'
	KPlus               Keycode = 0x0000002b // '+'
	KComma              Keycode = 0x0000002c // ','
	KMinus              Keycode = 0x0000002d // '-'
	KPeriod             Keycode = 0x0000002e // '.'
	KSlash              Keycode = 0x0000002f // '/'
	K0                  Keycode = 0x00000030 // '0'
	K1                  Keycode = 0x00000031 // '1'
	K2                  Keycode = 0x00000032 // '2'
	K3                  Keycode = 0x00000033 // '3'
	K4                  Keycode = 0x00000034 // '4'
	K5                  Keycode = 0x00000035 // '5'
	K6                  Keycode = 0x00000036 // '6'
	K7                  Keycode = 0x00000037 // '7'
	K8                  Keycode = 0x00000038 // '8'
	K9                  Keycode = 0x00000039 // '9'
	KColon              Keycode = 0x0000003a // ':'
	KSemicolon          Keycode = 0x0000003b // ';'
	KLess               Keycode = 0x0000003c // '<'
	KEquals             Keycode = 0x0000003d // '='
	KGreater            Keycode = 0x0000003e // '>'
	KQuestion           Keycode = 0x0000003f // '?'
	KAt                 Keycode = 0x00000040 // '@'
	KLeftbracket        Keycode = 0x0000005b // '['
	KBackslash          Keycode = 0x0000005c // '\\'
	KRightbracket       Keycode = 0x0000005d // ']'
	KCaret              Keycode = 0x0000005e // '^'
	KUnderscore         Keycode = 0x0000005f // '_'
	KGrave              Keycode = 0x00000060 // '`'
	KA                  Keycode = 0x00000061 // 'a'
	KB                  Keycode = 0x00000062 // 'b'
	KC                  Keycode = 0x00000063 // 'c'
	KD                  Keycode = 0x00000064 // 'd'
	KE                  Keycode = 0x00000065 // 'e'
	KF                  Keycode = 0x00000066 // 'f'
	KG                  Keycode = 0x00000067 // 'g'
	KH                  Keycode = 0x00000068 // 'h'
	KI                  Keycode = 0x00000069 // 'i'
	KJ                  Keycode = 0x0000006a // 'j'
	KK                  Keycode = 0x0000006b // 'k'
	KL                  Keycode = 0x0000006c // 'l'
	KM                  Keycode = 0x0000006d // 'm'
	KN                  Keycode = 0x0000006e // 'n'
	KO                  Keycode = 0x0000006f // 'o'
	KP                  Keycode = 0x00000070 // 'p'
	KQ                  Keycode = 0x00000071 // 'q'
	KR                  Keycode = 0x00000072 // 'r'
	KS                  Keycode = 0x00000073 // 's'
	KT                  Keycode = 0x00000074 // 't'
	KU                  Keycode = 0x00000075 // 'u'
	KV                  Keycode = 0x00000076 // 'v'
	KW                  Keycode = 0x00000077 // 'w'
	KX                  Keycode = 0x00000078 // 'x'
	KY                  Keycode = 0x00000079 // 'y'
	KZ                  Keycode = 0x0000007a // 'z'
	KLeftbrace          Keycode = 0x0000007b // '{'
	KPipe               Keycode = 0x0000007c // '|'
	KRightbrace         Keycode = 0x0000007d // '}'
	KTilde              Keycode = 0x0000007e // '~'
	KDelete             Keycode = 0x0000007f // '\x7F'
	KPlusminus          Keycode = 0x000000b1 // '\xB1'
	KCapslock           Keycode = 0x40000039 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CAPSLOCK)
	KF1                 Keycode = 0x4000003a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F1)
	KF2                 Keycode = 0x4000003b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F2)
	KF3                 Keycode = 0x4000003c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F3)
	KF4                 Keycode = 0x4000003d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F4)
	KF5                 Keycode = 0x4000003e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F5)
	KF6                 Keycode = 0x4000003f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F6)
	KF7                 Keycode = 0x40000040 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F7)
	KF8                 Keycode = 0x40000041 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F8)
	KF9                 Keycode = 0x40000042 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F9)
	KF10                Keycode = 0x40000043 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F10)
	KF11                Keycode = 0x40000044 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F11)
	KF12                Keycode = 0x40000045 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F12)
	KPrintscreen        Keycode = 0x40000046 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRINTSCREEN)
	KScrolllock         Keycode = 0x40000047 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SCROLLLOCK)
	KPause              Keycode = 0x40000048 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAUSE)
	KInsert             Keycode = 0x40000049 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_INSERT)
	KHome               Keycode = 0x4000004a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HOME)
	KPageup             Keycode = 0x4000004b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEUP)
	KEnd                Keycode = 0x4000004d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_END)
	KPagedown           Keycode = 0x4000004e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEDOWN)
	KRight              Keycode = 0x4000004f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RIGHT)
	KLeft               Keycode = 0x40000050 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LEFT)
	KDown               Keycode = 0x40000051 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DOWN)
	KUp                 Keycode = 0x40000052 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UP)
	KNumlockclear       Keycode = 0x40000053 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_NUMLOCKCLEAR)
	KKpDivide           Keycode = 0x40000054 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DIVIDE)
	KKpMultiply         Keycode = 0x40000055 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MULTIPLY)
	KKpMinus            Keycode = 0x40000056 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MINUS)
	KKpPlus             Keycode = 0x40000057 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUS)
	KKpEnter            Keycode = 0x40000058 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_ENTER)
	KKp1                Keycode = 0x40000059 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_1)
	KKp2                Keycode = 0x4000005a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_2)
	KKp3                Keycode = 0x4000005b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_3)
	KKp4                Keycode = 0x4000005c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_4)
	KKp5                Keycode = 0x4000005d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_5)
	KKp6                Keycode = 0x4000005e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_6)
	KKp7                Keycode = 0x4000005f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_7)
	KKp8                Keycode = 0x40000060 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_8)
	KKp9                Keycode = 0x40000061 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_9)
	KKp0                Keycode = 0x40000062 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_0)
	KKpPeriod           Keycode = 0x40000063 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERIOD)
	KApplication        Keycode = 0x40000065 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_APPLICATION)
	KPower              Keycode = 0x40000066 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_POWER)
	KKpEquals           Keycode = 0x40000067 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALS)
	KF13                Keycode = 0x40000068 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F13)
	KF14                Keycode = 0x40000069 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F14)
	KF15                Keycode = 0x4000006a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F15)
	KF16                Keycode = 0x4000006b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F16)
	KF17                Keycode = 0x4000006c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F17)
	KF18                Keycode = 0x4000006d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F18)
	KF19                Keycode = 0x4000006e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F19)
	KF20                Keycode = 0x4000006f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F20)
	KF21                Keycode = 0x40000070 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F21)
	KF22                Keycode = 0x40000071 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F22)
	KF23                Keycode = 0x40000072 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F23)
	KF24                Keycode = 0x40000073 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F24)
	KExecute            Keycode = 0x40000074 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXECUTE)
	KHelp               Keycode = 0x40000075 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HELP)
	KMenu               Keycode = 0x40000076 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MENU)
	KSelect             Keycode = 0x40000077 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SELECT)
	KStop               Keycode = 0x40000078 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_STOP)
	KAgain              Keycode = 0x40000079 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AGAIN)
	KUndo               Keycode = 0x4000007a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UNDO)
	KCut                Keycode = 0x4000007b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CUT)
	KCopy               Keycode = 0x4000007c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_COPY)
	KPaste              Keycode = 0x4000007d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PASTE)
	KFind               Keycode = 0x4000007e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_FIND)
	KMute               Keycode = 0x4000007f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MUTE)
	KVolumeup           Keycode = 0x40000080 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEUP)
	KVolumedown         Keycode = 0x40000081 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEDOWN)
	KKpComma            Keycode = 0x40000085 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COMMA)
	KKpEqualsas400      Keycode = 0x40000086 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALSAS400)
	KAlterase           Keycode = 0x40000099 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_ALTERASE)
	KSysreq             Keycode = 0x4000009a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SYSREQ)
	KCancel             Keycode = 0x4000009b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CANCEL)
	KClear              Keycode = 0x4000009c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEAR)
	KPrior              Keycode = 0x4000009d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRIOR)
	KReturn2            Keycode = 0x4000009e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RETURN2)
	KSeparator          Keycode = 0x4000009f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SEPARATOR)
	KOut                Keycode = 0x400000a0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OUT)
	KOper               Keycode = 0x400000a1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OPER)
	KClearagain         Keycode = 0x400000a2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEARAGAIN)
	KCrsel              Keycode = 0x400000a3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CRSEL)
	KExsel              Keycode = 0x400000a4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXSEL)
	KKp00               Keycode = 0x400000b0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_00)
	KKp000              Keycode = 0x400000b1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_000)
	KThousandsseparator Keycode = 0x400000b2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_THOUSANDSSEPARATOR)
	KDecimalseparator   Keycode = 0x400000b3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DECIMALSEPARATOR)
	KCurrencyunit       Keycode = 0x400000b4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYUNIT)
	KCurrencysubunit    Keycode = 0x400000b5 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYSUBUNIT)
	KKpLeftparen        Keycode = 0x400000b6 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTPAREN)
	KKpRightparen       Keycode = 0x400000b7 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTPAREN)
	KKpLeftbrace        Keycode = 0x400000b8 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTBRACE)
	KKpRightbrace       Keycode = 0x400000b9 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTBRACE)
	KKpTab              Keycode = 0x400000ba // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_TAB)
	KKpBackspace        Keycode = 0x400000bb // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BACKSPACE)
	KKpA                Keycode = 0x400000bc // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_A)
	KKpB                Keycode = 0x400000bd // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_B)
	KKpC                Keycode = 0x400000be // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_C)
	KKpD                Keycode = 0x400000bf // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_D)
	KKpE                Keycode = 0x400000c0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_E)
	KKpF                Keycode = 0x400000c1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_F)
	KKpXor              Keycode = 0x400000c2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_XOR)
	KKpPower            Keycode = 0x400000c3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_POWER)
	KKpPercent          Keycode = 0x400000c4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERCENT)
	KKpLess             Keycode = 0x400000c5 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LESS)
	KKpGreater          Keycode = 0x400000c6 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_GREATER)
	KKpAmpersand        Keycode = 0x400000c7 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AMPERSAND)
	KKpDblampersand     Keycode = 0x400000c8 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLAMPERSAND)
	KKpVerticalbar      Keycode = 0x400000c9 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_VERTICALBAR)
	KKpDblverticalbar   Keycode = 0x400000ca // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLVERTICALBAR)
	KKpColon            Keycode = 0x400000cb // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COLON)
	KKpHash             Keycode = 0x400000cc // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HASH)
	KKpSpace            Keycode = 0x400000cd // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_SPACE)
	KKpAt               Keycode = 0x400000ce // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AT)
	KKpExclam           Keycode = 0x400000cf // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EXCLAM)
	KKpMemstore         Keycode = 0x400000d0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSTORE)
	KKpMemrecall        Keycode = 0x400000d1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMRECALL)
	KKpMemclear         Keycode = 0x400000d2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMCLEAR)
	KKpMemadd           Keycode = 0x400000d3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMADD)
	KKpMemsubtract      Keycode = 0x400000d4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSUBTRACT)
	KKpMemmultiply      Keycode = 0x400000d5 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMMULTIPLY)
	KKpMemdivide        Keycode = 0x400000d6 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMDIVIDE)
	KKpPlusminus        Keycode = 0x400000d7 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUSMINUS)
	KKpClear            Keycode = 0x400000d8 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEAR)
	KKpClearentry       Keycode = 0x400000d9 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEARENTRY)
	KKpBinary           Keycode = 0x400000da // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BINARY)
	KKpOctal            Keycode = 0x400000db // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_OCTAL)
	KKpDecimal          Keycode = 0x400000dc // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DECIMAL)
	KKpHexadecimal      Keycode = 0x400000dd // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HEXADECIMAL)
	KLctrl              Keycode = 0x400000e0 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LCTRL)
	KLshift             Keycode = 0x400000e1 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LSHIFT)
	KLalt               Keycode = 0x400000e2 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LALT)
	KLgui               Keycode = 0x400000e3 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LGUI)
	KRctrl              Keycode = 0x400000e4 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RCTRL)
	KRshift             Keycode = 0x400000e5 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RSHIFT)
	KRalt               Keycode = 0x400000e6 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RALT)
	KRgui               Keycode = 0x400000e7 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RGUI)
	KMode               Keycode = 0x40000101 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MODE)
	KSleep              Keycode = 0x40000102 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SLEEP)
	KWake               Keycode = 0x40000103 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_WAKE)
	KChannelIncrement   Keycode = 0x40000104 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CHANNEL_INCREMENT)
	KChannelDecrement   Keycode = 0x40000105 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CHANNEL_DECREMENT)
	KMediaPlay          Keycode = 0x40000106 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PLAY)
	KMediaPause         Keycode = 0x40000107 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PAUSE)
	KMediaRecord        Keycode = 0x40000108 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_RECORD)
	KMediaFastForward   Keycode = 0x40000109 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_FAST_FORWARD)
	KMediaRewind        Keycode = 0x4000010a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_REWIND)
	KMediaNextTrack     Keycode = 0x4000010b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_NEXT_TRACK)
	KMediaPreviousTrack Keycode = 0x4000010c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PREVIOUS_TRACK)
	KMediaStop          Keycode = 0x4000010d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_STOP)
	KMediaEject         Keycode = 0x4000010e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_EJECT)
	KMediaPlayPause     Keycode = 0x4000010f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PLAY_PAUSE)
	KMediaSelect        Keycode = 0x40000110 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_SELECT)
	KAcNew              Keycode = 0x40000111 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_NEW)
	KAcOpen             Keycode = 0x40000112 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_OPEN)
	KAcClose            Keycode = 0x40000113 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_CLOSE)
	KAcExit             Keycode = 0x40000114 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_EXIT)
	KAcSave             Keycode = 0x40000115 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_SAVE)
	KAcPrint            Keycode = 0x40000116 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_PRINT)
	KAcProperties       Keycode = 0x40000117 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_PROPERTIES)
	KAcSearch           Keycode = 0x40000118 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_SEARCH)
	KAcHome             Keycode = 0x40000119 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_HOME)
	KAcBack             Keycode = 0x4000011a // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BACK)
	KAcForward          Keycode = 0x4000011b // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_FORWARD)
	KAcStop             Keycode = 0x4000011c // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_STOP)
	KAcRefresh          Keycode = 0x4000011d // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_REFRESH)
	KAcBookmarks        Keycode = 0x4000011e // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BOOKMARKS)
	KSoftleft           Keycode = 0x4000011f // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SOFTLEFT)
	KSoftright          Keycode = 0x40000120 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SOFTRIGHT)
	KCall               Keycode = 0x40000121 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CALL)
	KEndcall            Keycode = 0x40000122 // SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_ENDCALL)
	KLeftTab            Keycode = 0x20000001 // Extended key Left Tab
	KLevel5Shift        Keycode = 0x20000002 // Extended key Level 5 Shift
	KMultiKeyCompose    Keycode = 0x20000003 // Extended key Multi-key Compose
	KLmeta              Keycode = 0x20000004 // Extended key Left Meta
	KRmeta              Keycode = 0x20000005 // Extended key Right Meta
	KLhyper             Keycode = 0x20000006 // Extended key Left Hyper
	KRhyper             Keycode = 0x20000007 // Extended key Right Hyper
)

// Keymod represents valid key modifiers (possibly OR'd together).
//
// Defined in [<SDL3/SDL_keycode.h>].
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_keycode.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_keycode.h
type Keymod uint16

const (
	// KmodNone indicates no modifier is applicable.
	KmodNone Keymod = 0x0000

	// KmodLshift indicates the left Shift key is down.
	KmodLshift Keymod = 0x0001

	// KmodRshift indicates the right Shift key is down.
	KmodRshift Keymod = 0x0002

	// KmodLevel5 indicates the Level 5 Shift key is down.
	KmodLevel5 Keymod = 0x0004

	// KmodLctrl indicates the left Ctrl (Control) key is down.
	KmodLctrl Keymod = 0x0040

	// KmodRctrl indicates the right Ctrl (Control) key is down.
	KmodRctrl Keymod = 0x0080

	// KmodLalt indicates the left Alt key is down.
	KmodLalt Keymod = 0x0100

	// KmodRalt indicates the right Alt key is down.
	KmodRalt Keymod = 0x0200

	// KmodLgui indicates the left GUI key (often the Windows key) is down.
	KmodLgui Keymod = 0x0400

	// KmodRgui indicates the right GUI key (often the Windows key) is down.
	KmodRgui Keymod = 0x0800

	// KmodNum indicates the Num Lock key (may be located on an extended keypad) is down.
	KmodNum Keymod = 0x1000

	// KmodCaps indicates the Caps Lock key is down.
	KmodCaps Keymod = 0x2000

	// KmodMode indicates the !AltGr key is down.
	KmodMode Keymod = 0x4000

	// KmodScroll indicates the Scroll Lock key is down.
	KmodScroll Keymod = 0x8000

	// KmodCtrl indicates Any Ctrl key is down.
	KmodCtrl Keymod = (KmodLctrl | KmodRctrl)

	// KmodShift indicates Any Shift key is down.
	KmodShift Keymod = (KmodLshift | KmodRshift)

	// KmodAlt indicates Any Alt key is down.
	KmodAlt Keymod = (KmodLalt | KmodRalt)

	// KmodGui indicates Any GUI key is down.
	KmodGui Keymod = (KmodLgui | KmodRgui)
)

// PowerState represents the basic state for the system's power supply.
//
// Defined in [<SDL3/SDL_power.h>].
//
// These are results returned by [GetPowerInfo].
//
// This enum is available since SDL 3.2.0.
//
// [<SDL3/SDL_power.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_power.h
type PowerState int8

// PowerState values.
const (
	// PowerstateError indicates error determining power status.
	PowerstateError PowerState = -1 + iota

	// PowerstateUnknown indicates cannot determine power status.
	PowerstateUnknown

	// PowerstateOnBattery indicates not plugged in, running on the battery.
	PowerstateOnBattery

	// PowerstateNoBattery indicates plugged in, no battery available.
	PowerstateNoBattery

	// PowerstateCharging indicates plugged in, charging battery.
	PowerstateCharging

	// PowerstateCharged indicates plugged in, battery charged.
	PowerstateCharged
)

// MouseButtonFlags is a bitmask of pressed mouse buttons, as reported by
// [GetMouseState], etc.
//
// Defined in [<SDL3/SDL_mouse.h>].
//
//   - Button 1: Left mouse button
//   - Button 2: Middle mouse button
//   - Button 3: Right mouse button
//   - Button 4: Side mouse button 1
//   - Button 5: Side mouse button 2
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_mouse.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_mouse.h
type MouseButtonFlags uint32

// MouseButtonFlags values.
const (
	ButtonLeft   MouseButtonFlags = 1
	ButtonMiddle MouseButtonFlags = 2
	ButtonRight  MouseButtonFlags = 3
	ButtonX1     MouseButtonFlags = 4
	ButtonX2     MouseButtonFlags = 5

	ButtonLmask  MouseButtonFlags = 1 << (ButtonLeft - 1)
	ButtonMmask  MouseButtonFlags = 1 << (ButtonMiddle - 1)
	ButtonRmask  MouseButtonFlags = 1 << (ButtonRight - 1)
	ButtonX1mask MouseButtonFlags = 1 << (ButtonX1 - 1)
	ButtonX2mask MouseButtonFlags = 1 << (ButtonX2 - 1)
)

// MouseWheelDirection indicates scroll direction types for the Scroll event.
//
// Defined in [<SDL3/SDL_mouse.h>].
//
// This enum is available since SDL 3.2.0.
//
// [<SDL3/SDL_mouse.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_mouse.h
type MouseWheelDirection uint8

// MouseWheelDirection values.
const (
	// MousewheelNormal indicates the scroll direction is normal.
	MousewheelNormal MouseWheelDirection = iota

	// MousewheelFlipped indicates the scroll direction is flipped / natural.
	MousewheelFlipped
)

// PenInputFlags are pen input flags, as reported by various pen events'
// pen_state field.
//
// Defined in [<SDL3/SDL_pen.h>].
//
// This datatype is available since SDL 3.2.0.
//
// [<SDL3/SDL_pen.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_pen.h
type PenInputFlags uint32

// PenInputFlags values.
const (
	// PenInputDown indicates pen is pressed down.
	PenInputDown PenInputFlags = 1 << 0

	// PenInputButton1 indicates button 1 is pressed.
	PenInputButton1 PenInputFlags = 1 << 1

	// PenInputButton2 indicates button 2 is pressed.
	PenInputButton2 PenInputFlags = 1 << 2

	// PenInputButton3 indicates button 3 is pressed.
	PenInputButton3 PenInputFlags = 1 << 3

	// PenInputButton4 indicates button 4 is pressed.
	PenInputButton4 PenInputFlags = 1 << 4

	// PenInputButton5 indicates button 5 is pressed.
	PenInputButton5 PenInputFlags = 1 << 5

	// PenInputEraserTip indicates eraser tip is used.
	PenInputEraserTip PenInputFlags = 1 << 30

	// PenInputInProximity indicates pen is in proximity (since SDL 3.4.0).
	PenInputInProximity PenInputFlags = 1 << 31
)

// PenAxis holds pen axis indices.
//
// Defined in [<SDL3/SDL_pen.h>].
//
// These are the valid values for the axis field in [PenAxisEvent]. All axes
// are either normalised to 0..1 or report a (positive or negative) angle in
// degrees, with 0.0 representing the centre. Not all pens/backends support all
// axes: unsupported axes are always zero.
//
// To convert angles for tilt and rotation into vector representation, use
// SDL_sinf on the XTILT, YTILT, or ROTATION component, for example:
//
//	SDL_sinf(xtilt * SDL_PI_F / 180.0)
//
// This enum is available since SDL 3.2.0.
//
// [<SDL3/SDL_pen.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_pen.h
type PenAxis uint16

// PenAxis values.
const (
	// PenAxisPressure indicates pen pressure. Unidirectional: 0 to 1.0.
	PenAxisPressure PenAxis = iota

	// PenAxisXtilt indicates pen horizontal tilt angle. Bidirectional: -90.0
	// to 90.0 (left-to-right).
	PenAxisXtilt

	// PenAxisYtilt indicates pen vertical tilt angle. Bidirectional: -90.0 to
	// 90.0 (top-to-down).
	PenAxisYtilt

	// PenAxisDistance indicates pen distance to drawing surface.
	// Unidirectional: 0.0 to 1.0.
	PenAxisDistance

	// PenAxisRotation indicates pen barrel rotation. Bidirectional: -180 to
	// 179.9 (clockwise, 0 is facing up, -180.0 is facing down).
	PenAxisRotation

	// PenAxisSlider indicates pen finger wheel or slider (e.g., Airbrush Pen).
	// Unidirectional: 0 to 1.0.
	PenAxisSlider

	// PenAxisTangentialPressure indicates Pressure from squeezing the pen
	// ("barrel pressure").
	PenAxisTangentialPressure

	// PenAxisCount indicates Total known pen axis types in this version of
	// SDL. This number may grow in future releases!
	PenAxisCount
)

// JoystickHat represents the current hat position. It is simply an Uint8 in
// SDL.
type JoystickHat uint8

// JoystickHat values.
const (
	HatCentered  JoystickHat = 0x00
	HatUp        JoystickHat = 0x01
	HatRight     JoystickHat = 0x02
	HatDown      JoystickHat = 0x04
	HatLeft      JoystickHat = 0x08
	HatRightup   JoystickHat = (HatRight | HatUp)
	HatRightdown JoystickHat = (HatRight | HatDown)
	HatLeftup    JoystickHat = (HatLeft | HatUp)
	HatLeftdown  JoystickHat = (HatLeft | HatDown)
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

// GamepadAxisInvalid is the invalid GamepadAxis value.
const GamepadAxisInvalid GamepadAxis = math.MaxUint8

// GamepadAxis values.
const (
	GamepadAxisLeftx GamepadAxis = iota
	GamepadAxisLefty
	GamepadAxisRightx
	GamepadAxisRighty
	GamepadAxisLeftTrigger
	GamepadAxisRightTrigger
	GamepadAxisCount
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

// GamepadButtonInvalid is the invalid GamepadButton value.
const GamepadButtonInvalid GamepadButton = math.MaxUint8

// GamepadButton values.
const (
	// GamepadButtonSouth is the bottom face button (e.g. Xbox A button).
	GamepadButtonSouth GamepadButton = iota
	// GamepadButtonEast is the right face button (e.g. Xbox B button).
	GamepadButtonEast
	// GamepadButtonWest is the left face button (e.g. Xbox X button).
	GamepadButtonWest
	// GamepadButtonNorth is the top face button (e.g. Xbox Y button).
	GamepadButtonNorth
	GamepadButtonBack
	GamepadButtonGuide
	GamepadButtonStart
	GamepadButtonLeftStick
	GamepadButtonRightStick
	GamepadButtonLeftShoulder
	GamepadButtonRightShoulder
	GamepadButtonDpadUp
	GamepadButtonDpadDown
	GamepadButtonDpadLeft
	GamepadButtonDpadRight
	// GamepadButtonMisc1 is the additional button (e.g. Xbox Series X share
	// button, PS5 microphone button, Nintendo Switch Pro capture button, Steam
	// Controller QAM button, Amazon Luna microphone button, Google Stadia
	// capture button).
	GamepadButtonMisc1
	// GamepadButtonRightPaddle1 is the upper or primary paddle, under your
	// right hand (e.g. Xbox Elite paddle P1, DualSense Edge RB button, Right
	// Joy-Con SR button, Steam Controller R4 button).
	GamepadButtonRightPaddle1
	// GamepadButtonLeftPaddle1 is the upper or primary paddle, under your left
	// hand (e.g. Xbox Elite paddle P3, DualSense Edge LB button, Left Joy-Con
	// SL button, Steam Controller L4 button).
	GamepadButtonLeftPaddle1
	// GamepadButtonRightPaddle2 is the lower or secondary paddle, under your
	// right hand (e.g. Xbox Elite paddle P2, DualSense Edge right Fn button,
	// Right Joy-Con SL button, Steam Controller R5 button).
	GamepadButtonRightPaddle2
	// GamepadButtonLeftPaddle2 is the lower or secondary paddle, under your
	// left hand (e.g. Xbox Elite paddle P4, DualSense Edge left Fn button,
	// Left Joy-Con SR button, Steam Controller L5 button).
	GamepadButtonLeftPaddle2
	// GamepadButtonTouchpad is the PS4/PS5 touchpad button.
	GamepadButtonTouchpad
	// GamepadButtonMisc2 is an additional button.
	GamepadButtonMisc2
	// GamepadButtonMisc3 is an additional button (e.g. Nintendo GameCube left
	// trigger click).
	GamepadButtonMisc3
	// GamepadButtonMisc4 is an additional button (e.g. Nintendo GameCube right
	// trigger click).
	GamepadButtonMisc4
	// GamepadButtonMisc5 is an additional button.
	GamepadButtonMisc5
	// GamepadButtonMisc6 is an additional button.
	GamepadButtonMisc6
	GamepadButtonCount
)

// SensorType represents the different sensors defined by SDL.
//
// Defined in [<SDL3/SDL_sensor.h>].
//
// Additional sensors may be available, using platform dependent semantics.
//
// Here are the additional Android sensors:
//
// https://developer.android.com/reference/android/hardware/SensorEvent.html#values
//
// Accelerometer sensor notes:
//
// The accelerometer returns the current acceleration in SI meters per second
// squared. This measurement includes the force of gravity, so a device at rest
// will have an value of SDL_STANDARD_GRAVITY away from the center of the
// earth, which is a positive Y value.
//
//   - values[0]: Acceleration on the x axis
//   - values[1]: Acceleration on the y axis
//   - values[2]: Acceleration on the z axis
//
// For phones and tablets held in natural orientation and game controllers held
// in front of you, the axes are defined as follows:
//
//   - -X ... +X : left ... right
//   - -Y ... +Y : bottom ... top
//   - -Z ... +Z : farther ... closer
//
// The accelerometer axis data is not changed when the device is rotated.
//
// Gyroscope sensor notes:
//
// The gyroscope returns the current rate of rotation in radians per second.
// The rotation is positive in the counter-clockwise direction. That is, an
// observer looking from a positive location on one of the axes would see
// positive rotation on that axis when it appeared to be rotating
// counter-clockwise.
//
//   - values[0]: Angular speed around the x axis (pitch)
//   - values[1]: Angular speed around the y axis (yaw)
//   - values[2]: Angular speed around the z axis (roll)
//
// For phones and tablets held in natural orientation and game controllers held
// in front of you, the axes are defined as follows:
//
//   - -X ... +X : left ... right
//   - -Y ... +Y : bottom ... top
//   - -Z ... +Z : farther ... closer
//
// The gyroscope axis data is not changed when the device is rotated.
//
// This enum is available since SDL 3.2.0.
//
// [<SDL3/SDL_sensor.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_sensor.h
type SensorType int32

// SensorType values.
const (
	// SensorInvalid is returned for an invalid sensor.
	SensorInvalid = -1 + iota
	// SensorUnknown indicates unknown sensor type.
	SensorUnknown
	// SensorAccel indicates accelerometer.
	SensorAccel
	// SensorGyro indicates gyroscope.
	SensorGyro
	// SensorAccelL indicates accelerometer for left Joy-Con controller and Wii
	// nunchuk.
	SensorAccelL
	// SensorGyroL indicates gyroscope for left Joy-Con controller.
	SensorGyroL
	// SensorAccelR indicates accelerometer for right Joy-Con controller.
	SensorAccelR
	// SensorGyroR indicates gyroscope for right Joy-Con controller.
	SensorGyroR
	SensorCount
)
