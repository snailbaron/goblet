package sdl

// #include <SDL3/SDL.h>
import "C"

import (
	"unsafe"
)

func toStrings(cStrings **C.char, num int32) []string {
	results := make([]string, num)
	for i, cc := range unsafe.Slice(cStrings, num) {
		results[i] = C.GoString(cc)
	}
	return results
}

type Event struct {
	CommonEvent

	internal any
}

func cast[T any](e *Event) *T {
	t, _ := e.internal.(*T)
	return t
}

func (e *Event) Display() *DisplayEvent {
	return cast[DisplayEvent](e)
}

func (e *Event) Window() *WindowEvent {
	return cast[WindowEvent](e)
}

func (e *Event) KDevice() *KeyboardDeviceEvent {
	return cast[KeyboardDeviceEvent](e)
}

func (e *Event) Key() *KeyboardEvent {
	return cast[KeyboardEvent](e)
}

func (e *Event) Edit() *TextEditingEvent {
	return cast[TextEditingEvent](e)
}

func (e *Event) EditCandidates() *TextEditingCandidatesEvent {
	return cast[TextEditingCandidatesEvent](e)
}

func (e *Event) Text() *TextInputEvent {
	return cast[TextInputEvent](e)
}

func (e *Event) MDevice() *MouseDeviceEvent {
	return cast[MouseDeviceEvent](e)
}

func (e *Event) Motion() *MouseMotionEvent {
	return cast[MouseMotionEvent](e)
}

func (e *Event) Button() *MouseButtonEvent {
	return cast[MouseButtonEvent](e)
}

func (e *Event) Wheel() *MouseWheelEvent {
	return cast[MouseWheelEvent](e)
}

func (e *Event) JDevice() *JoyDeviceEvent {
	return cast[JoyDeviceEvent](e)
}

func (e *Event) JAxis() *JoyAxisEvent {
	return cast[JoyAxisEvent](e)
}

func (e *Event) JBall() *JoyBallEvent {
	return cast[JoyBallEvent](e)
}

func (e *Event) JHat() *JoyHatEvent {
	return cast[JoyHatEvent](e)
}

func (e *Event) JButton() *JoyButtonEvent {
	return cast[JoyButtonEvent](e)
}

func (e *Event) JBattery() *JoyBatteryEvent {
	return cast[JoyBatteryEvent](e)
}

func (e *Event) GDevice() *GamepadDeviceEvent {
	return cast[GamepadDeviceEvent](e)
}

func (e *Event) GAxis() *GamepadAxisEvent {
	return cast[GamepadAxisEvent](e)
}

func (e *Event) GButton() *GamepadButtonEvent {
	return cast[GamepadButtonEvent](e)
}

func (e *Event) GTouchpad() *GamepadTouchpadEvent {
	return cast[GamepadTouchpadEvent](e)
}

func (e *Event) GSensor() *GamepadSensorEvent {
	return cast[GamepadSensorEvent](e)
}

func (e *Event) ADevice() *AudioDeviceEvent {
	return cast[AudioDeviceEvent](e)
}

func (e *Event) CDevice() *CameraDeviceEvent {
	return cast[CameraDeviceEvent](e)
}

func (e *Event) Sensor() *SensorEvent {
	return cast[SensorEvent](e)
}

func (e *Event) Quit() *QuitEvent {
	return cast[QuitEvent](e)
}

func (e *Event) User() *UserEvent {
	return cast[UserEvent](e)
}

func (e *Event) TFinger() *TouchFingerEvent {
	return cast[TouchFingerEvent](e)
}

func (e *Event) Pinch() *PinchFingerEvent {
	return cast[PinchFingerEvent](e)
}

func (e *Event) PProximity() *PenProximityEvent {
	return cast[PenProximityEvent](e)
}

func (e *Event) PTouch() *PenTouchEvent {
	return cast[PenTouchEvent](e)
}

func (e *Event) PMotion() *PenMotionEvent {
	return cast[PenMotionEvent](e)
}

func (e *Event) PButton() *PenButtonEvent {
	return cast[PenButtonEvent](e)
}

func (e *Event) PAxis() *PenAxisEvent {
	return cast[PenAxisEvent](e)
}

func (e *Event) Render() *RenderEvent {
	return cast[RenderEvent](e)
}

func (e *Event) Drop() *DropEvent {
	return cast[DropEvent](e)
}

func (e *Event) Clipboard() *ClipboardEvent {
	return cast[ClipboardEvent](e)
}

// AudioDeviceEvent is an audio device event structure (event.adevice.*).
//
// [EVENT_AUDIO_DEVICE_ADDED], or [EVENT_AUDIO_DEVICE_REMOVED], or
// [EVENT_AUDIO_DEVICE_FORMAT_CHANGED]
//
// Defined in [<SDL3/SDL_events.h>]
//
// Note that SDL will send a [EVENT_AUDIO_DEVICE_ADDED] event for every device
// it discovers during initialization. After that, this event will only arrive
// when a device is hotplugged during the program's run.
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type AudioDeviceEvent struct {
	// AudioDeviceID for the device being added or removed or changing
	Which AudioDeviceID

	// false if a playback device, true if a recording device.
	Recording bool
}

// CameraDeviceEvent is a camera device event structure (event.cdevice.*).
//
// [EVENT_CAMERA_DEVICE_ADDED], [EVENT_CAMERA_DEVICE_REMOVED],
// [EVENT_CAMERA_DEVICE_APPROVED], [EVENT_CAMERA_DEVICE_DENIED]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type CameraDeviceEvent struct {
	// CameraID for the device being added or removed or changing
	Which CameraID
}

// ClipboardEvent is an event triggered when the clipboard contents have
// changed (event.clipboard.*).
//
// [EVENT_CLIPBOARD_UPDATE]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type ClipboardEvent struct {
	// are we owning the clipboard (internal update)
	Owner bool

	// current mime types
	MIMETypes []string
}

// CommonEvent defines fields shared by every event.
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type CommonEvent struct {
	// CommonEvent type, shared with all events, Uint32 to cover user events
	// which are not in the EventType enumeration
	Type EventType

	// In nanoseconds, populated using GetTicksNS()
	Timestamp uint64
}

// DisplayEvent contains display state change event data (event.display.*).
//
// EVENT_DISPLAY_*
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type DisplayEvent struct {
	// The associated display
	DisplayID DisplayID

	// event dependent data
	Data1 int32

	// event dependent data
	Data2 int32
}

// DropEvent is an event used to drop text or request a file open by the system
// (event.drop.*).
//
// [EVENT_DROP_BEGIN] or [EVENT_DROP_FILE] or [EVENT_DROP_TEXT] or
// [EVENT_DROP_COMPLETE] or [EVENT_DROP_POSITION]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type DropEvent struct {
	// The window that was dropped on, if any
	WindowID WindowID

	// X coordinate, relative to window (not on begin)
	X float32

	// Y coordinate, relative to window (not on begin)
	Y float32

	// The source app that sent this drop event, or NULL if that isn't
	// available
	Source string

	// The text for EVENT_DROP_TEXT and the file name for EVENT_DROP_FILE, NULL
	// for other events
	Data string
}

// GamepadAxisEvent defines gamepad axis motion event structure
// (event.gaxis.*).
//
// [EVENT_GAMEPAD_AXIS_MOTION]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadAxisEvent struct {
	// The joystick instance id
	Which JoystickID

	// The gamepad axis
	Axis GamepadAxis

	// The axis value (range: -32768 to 32767)
	Value int16
}

// GamepadButtonEvent is a gamepad button event structure (event.gbutton.*).
//
// [EVENT_GAMEPAD_BUTTON_DOWN] or [EVENT_GAMEPAD_BUTTON_UP]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadButtonEvent struct {
	// The joystick instance id
	Which JoystickID

	// The gamepad button (GamepadButton)
	Button uint8

	// true if the button is pressed
	Down bool
}

// GamepadDeviceEvent is a gamepad device event structure (event.gdevice.*).
//
// [EVENT_GAMEPAD_ADDED], [EVENT_GAMEPAD_REMOVED], or
// [EVENT_GAMEPAD_REMAPPED], [EVENT_GAMEPAD_UPDATE_COMPLETE] or
// [EVENT_GAMEPAD_STEAM_HANDLE_UPDATED]
//
// Defined in [<SDL3/SDL_events.h>]
//
// Joysticks that are supported gamepads receive both an [JoyDeviceEvent] and
// an [GamepadDeviceEvent].
//
// SDL will send [EVENT_GAMEPAD_ADDED] events for joysticks that are already
// plugged in during [Init]() and are recognized as gamepads. It will also send
// events for joysticks that get gamepad mappings at runtime.
//
// This struct is available since SDL 3.2.0.
//
// See Also:
//
//   - [JoyDeviceEvent]
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadDeviceEvent struct {
	// The joystick instance id
	Which JoystickID
}

// GamepadSensorEvent is a gamepad sensor event structure (event.gsensor.*).
//
// [EVENT_GAMEPAD_SENSOR_UPDATE]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadSensorEvent struct {
	// The joystick instance id
	Which JoystickID

	// The type of the sensor, one of the values of SensorType
	Sensor SensorType

	// Up to 3 values from the sensor, as defined in SDL_sensor.h
	Data [3]float32

	// The timestamp of the sensor reading in nanoseconds, not necessarily
	// synchronized with the system clock
	SensorTimestamp uint64
}

// GamepadTouchpadEvent is a gamepad touchpad event structure
// (event.gtouchpad.*).
//
// [EVENT_GAMEPAD_TOUCHPAD_DOWN] or [EVENT_GAMEPAD_TOUCHPAD_MOTION] or
// [EVENT_GAMEPAD_TOUCHPAD_UP].
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadTouchpadEvent struct {
	// The joystick instance id
	Which JoystickID

	// The index of the touchpad
	Touchpad int32

	// The index of the finger on the touchpad
	Finger int32

	// Normalized in the range 0...1 with 0 being on the left
	X float32

	// Normalized in the range 0...1 with 0 being at the top
	Y float32

	// Normalized in the range 0...1
	Pressure float32
}

// JoyAxisEvent is a joystick axis motion event structure (event.jaxis.*).
//
// [EVENT_JOYSTICK_AXIS_MOTION]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyAxisEvent struct {
	// The joystick instance id
	Which JoystickID

	// The joystick axis index
	Axis uint8

	// The axis value (range: -32768 to 32767)
	Value int16
}

// JoyBallEvent is a joystick trackball motion event structure (event.jball.*).
//
// [EVENT_JOYSTICK_BALL_MOTION]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyBallEvent struct {
	// The joystick instance id
	Which JoystickID

	// The joystick trackball index
	Ball uint8

	// The relative motion in the X direction
	Xrel int16

	// The relative motion in the Y direction
	Yrel int16
}

// JoyBatteryEvent is a joystick battery level change event structure
// (event.jbattery.*).
//
// [EVENT_JOYSTICK_BATTERY_UPDATED]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyBatteryEvent struct {
	// The joystick instance id
	Which JoystickID

	// The joystick battery state
	State PowerState

	// The joystick battery percent charge remaining
	Percent int
}

// JoyButtonEvent is a joystick button event structure (event.jbutton.*).
//
// [EVENT_JOYSTICK_BUTTON_DOWN] or [EVENT_JOYSTICK_BUTTON_UP]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyButtonEvent struct {
	// The joystick instance id
	Which JoystickID

	// The joystick button index
	Button uint8

	// true if the button is pressed
	Down bool
}

// JoyDeviceEvent is a joystick device event structure (event.jdevice.*).
//
// [EVENT_JOYSTICK_ADDED] or [EVENT_JOYSTICK_REMOVED] or
// [EVENT_JOYSTICK_UPDATE_COMPLETE]
//
// Defined in [<SDL3/SDL_events.h>]
//
// SDL will send JOYSTICK_ADDED events for devices that are already plugged in
// during [Init].
//
// This struct is available since SDL 3.2.0.
//
// See Also:
//
//   - [GamepadDeviceEvent]
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyDeviceEvent struct {
	// The joystick instance id
	Which JoystickID
}

// JoyHatEvent is a joystick hat position change event structure
// (event.jhat.*).
//
// [EVENT_JOYSTICK_HAT_MOTION]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyHatEvent struct {
	// The joystick instance id
	Which JoystickID

	// The joystick hat index
	Hat uint8

	// The hat position value.
	//
	// Note that zero means the POV is centered.
	Value JoystickHat
}

// KeyboardDeviceEvent is a keyboard device event structure (event.kdevice.*).
//
// [EVENT_KEYBOARD_ADDED] or [EVENT_KEYBOARD_REMOVED]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type KeyboardDeviceEvent struct {
	// The keyboard instance id
	Which KeyboardID
}

// KeyboardEvent is a keyboard button event structure (event.key.*).
//
// [EVENT_KEY_DOWN] or [EVENT_KEY_UP]
//
// Defined in [<SDL3/SDL_events.h>]
//
// The `key` is the base [Keycode] generated by pressing the `scancode` using
// the current keyboard layout, applying any options specified in
// [HINT_KEYCODE_OPTIONS]. You can get the [Keycode] corresponding to the event
// scancode and modifiers directly from the keyboard layout, bypassing
// [HINT_KEYCODE_OPTIONS], by calling [GetKeyFromScancode]().
//
// This struct is available since SDL 3.2.0.
//
// See Also:
//
//   - [GetKeyFromScancode]
//   - [HINT_KEYCODE_OPTIONS]
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type KeyboardEvent struct {
	// The window with keyboard focus, if any
	WindowID WindowID

	// The keyboard instance id, or 0 if unknown or virtual
	Which KeyboardID

	// SDL physical key code
	Scancode Scancode

	// SDL virtual key code
	Key Keycode

	// current key modifiers
	Mod Keymod

	// The platform dependent scancode for this event
	Raw uint16

	// true if the key is pressed
	Down bool

	// true if this is a key repeat
	Repeat bool
}

// MouseButtonEvent is a mouse button event structure (event.button.*).
//
// [EVENT_MOUSE_BUTTON_DOWN] or [EVENT_MOUSE_BUTTON_UP]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type MouseButtonEvent struct {
	// The window with mouse focus, if any
	WindowID WindowID

	// The mouse instance id in relative mode, TOUCH_MOUSEID for touch events,
	// or 0
	Which MouseID

	// The mouse button index
	Button uint8

	// true if the button is pressed
	Down bool

	// 1 for single-click, 2 for double-click, etc.
	Clicks uint8

	// X coordinate, relative to window
	X float32

	// Y coordinate, relative to window
	Y float32
}

// MouseDeviceEvent is a mouse device event structure (event.mdevice.*).
//
// [EVENT_MOUSE_ADDED] or [EVENT_MOUSE_REMOVED]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type MouseDeviceEvent struct {
	// The mouse instance id
	Which MouseID
}

// MouseMotionEvent is a mouse motion event structure (event.motion.*).
//
// [EVENT_MOUSE_MOTION]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type MouseMotionEvent struct {
	// The window with mouse focus, if any
	WindowID WindowID

	// The mouse instance id in relative mode, TOUCH_MOUSEID for touch events,
	// PEN_MOUSEID for pen events, or 0
	Which MouseID

	// The current button state
	State MouseButtonFlags

	// X coordinate, relative to window
	X float32

	// Y coordinate, relative to window
	Y float32

	// The relative motion in the X direction
	Xrel float32

	// The relative motion in the Y direction
	Yrel float32
}

// MouseWheelEvent is a mouse wheel event structure (event.wheel.*).
//
// [EVENT_MOUSE_WHEEL]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type MouseWheelEvent struct {
	// The window with mouse focus, if any
	WindowID WindowID

	// The mouse instance id in relative mode or 0
	Which MouseID

	// The amount scrolled horizontally, positive to the right and negative to
	// the left
	X float32

	// The amount scrolled vertically, positive away from the user and negative
	// toward the user
	Y float32

	// Set to one of the MOUSEWHEEL_* defines. When FLIPPED the values in X and
	// Y will be opposite. Multiply by -1 to change them back
	Direction MouseWheelDirection

	// X coordinate, relative to window
	MouseX float32

	// Y coordinate, relative to window
	MouseY float32

	// The amount scrolled horizontally, accumulated to whole scroll ``ticks''
	// (added in 3.2.12)
	IntegerX int32

	// The amount scrolled vertically, accumulated to whole scroll ``ticks''
	// (added in 3.2.12)
	IntegerY int32
}

// PenAxisEvent is a pressure-sensitive pen pressure / angle event structure
// (event.paxis.*).
//
// [EVENT_PEN_AXIS]
//
// Defined in [<SDL3/SDL_events.h>]
//
// You might get some of these events even if the pen isn't touching the
// tablet.
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenAxisEvent struct {
	// The window with pen focus, if any
	WindowID WindowID

	// The pen instance id
	Which PenID

	// Complete pen input state at time of event
	PenState PenInputFlags

	// X coordinate, relative to window
	X float32

	// Y coordinate, relative to window
	Y float32

	// Axis that has changed
	Axis PenAxis

	// New value of axis
	Value float32
}

// PenButtonEvent is a pressure-sensitive pen button event structure
// (event.pbutton.*).
//
// [EVENT_PEN_BUTTON_DOWN] or [EVENT_PEN_BUTTON_UP]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This is for buttons on the pen itself that the user might click. The pen
// itself pressing down to draw triggers a [EVENT_PEN_DOWN] event instead.
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenButtonEvent struct {
	// The window with mouse focus, if any
	WindowID WindowID

	// The pen instance id
	Which PenID

	// Complete pen input state at time of event
	PenState PenInputFlags

	// X coordinate, relative to window
	X float32

	// Y coordinate, relative to window
	Y float32

	// The pen button index (first button is 1).
	Button uint8

	// true if the button is pressed
	Down bool
}

// PenMotionEvent is a pressure-sensitive pen motion event structure
// (event.pmotion.*).
//
// [EVENT_PEN_MOTION]
//
// Defined in [<SDL3/SDL_events.h>]
//
// Depending on the hardware, you may get motion events when the pen is not
// touching a tablet, for tracking a pen even when it isn't drawing. You should
// listen for [EVENT_PEN_DOWN] and [EVENT_PEN_UP] events, or check `pen_state &
// PEN_INPUT_DOWN` to decide if a pen is "drawing" when dealing with pen
// motion.
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenMotionEvent struct {
	// The window with pen focus, if any
	WindowID WindowID

	// The pen instance id
	Which PenID

	// Complete pen input state at time of event
	PenState PenInputFlags

	// X coordinate, relative to window
	X float32

	// Y coordinate, relative to window
	Y float32
}

// PenProximityEvent is a pressure-sensitive pen proximity event structure
// (event.pproximity.*).
//
// [EVENT_PEN_PROXIMITY_IN] or [EVENT_PEN_PROXIMITY_OUT]
//
// Defined in [<SDL3/SDL_events.h>]
//
// When a pen becomes visible to the system (it is close enough to a tablet,
// etc), SDL will send an [EVENT_PEN_PROXIMITY_IN] event with the new pen's ID.
// This ID is valid until the pen leaves proximity again (has been removed from
// the tablet's area, the tablet has been unplugged, etc). If the same pen
// reenters proximity again, it will be given a new ID.
//
// Note that "proximity" means "close enough for the tablet to know the tool is
// there." The pen touching and lifting off from the tablet while not leaving
// the area are handled by [EVENT_PEN_DOWN] and [EVENT_PEN_UP].
//
// Not all platforms have a window associated with the pen during proximity
// events. Some wait until motion/button/etc events to offer this info.
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenProximityEvent struct {
	// The window with pen focus, if any
	WindowID WindowID

	// The pen instance id
	Which PenID
}

// PenTouchEvent is a pressure-sensitive pen touched event structure
// (event.ptouch.*).
//
// [EVENT_PEN_DOWN] or [EVENT_PEN_UP]
//
// Defined in [<SDL3/SDL_events.h>]
//
// These events come when a pen touches a surface (a tablet, etc), or lifts
// off from one.
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenTouchEvent struct {
	// The window with pen focus, if any
	WindowID WindowID

	// The pen instance id
	Which PenID

	// Complete pen input state at time of event
	PenState PenInputFlags

	// X coordinate, relative to window
	X float32

	// Y coordinate, relative to window
	Y float32

	// true if eraser end is used (not all pens support this).
	Eraser bool

	// true if the pen is touching or false if the pen is lifted off
	Down bool
}

// PinchFingerEvent is a pinch event structure (event.pinch.*).
//
// [EVENT_PINCH_BEGIN] or [EVENT_PINCH_UPDATE] or [EVENT_PINCH_END]
//
// Defined in [<SDL3/SDL_events.h>]
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PinchFingerEvent struct {
	// The scale change since the last [EVENT_PINCH_UPDATE]. Scale < 1 is
	// ``zoom out''. Scale > 1 is ``zoom in''.
	Scale float32

	// The window underneath the finger, if any
	WindowID WindowID
}

// QuitEvent is the “quit requested” event.
//
// [EVENT_QUIT]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type QuitEvent struct{}

// RenderEvent is a renderer event structure (event.render.*).
//
// [EVENT_RENDER_TARGETS_RESET], [EVENT_RENDER_DEVICE_RESET],
// [EVENT_RENDER_DEVICE_LOST]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type RenderEvent struct {
	// The window containing the renderer in question.
	WindowID WindowID
}

// SensorEvent is a sensor event structure (event.sensor.*).
//
// [EVENT_SENSOR_UPDATE]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type SensorEvent struct {
	// The instance ID of the sensor
	Which SensorID

	// Up to 6 values from the sensor - additional values can be queried using
	// GetSensorData()
	Data [6]float32

	// The timestamp of the sensor reading in nanoseconds, not necessarily
	// synchronized with the system clock
	SensorTimestamp uint64
}

// TextEditingCandidatesEvent is a keyboard IME candidates event structure
// (event.edit_candidates.*).
//
// [EVENT_TEXT_EDITING_CANDIDATES]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type TextEditingCandidatesEvent struct {
	// The window with keyboard focus, if any
	WindowID WindowID

	// The list of candidates, or NULL if there are no candidates available
	Candidates []string

	// The index of the selected candidate, or -1 if no candidate is selected
	SelectedCandidate int32

	// true if the list is horizontal, false if it's vertical
	Horizontal bool
}

// TextEditingEvent is a keyboard text editing event structure (event.edit.*).
//
// [EVENT_TEXT_EDITING]
//
// Defined in [<SDL3/SDL_events.h>]
//
// The start cursor is the position, in UTF-8 characters, where new typing
// will be inserted into the editing text. The length is the number of UTF-8
// characters that will be replaced by new typing.
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type TextEditingEvent struct {
	// The window with keyboard focus, if any
	WindowID WindowID

	// The editing text
	Text string

	// The start cursor of selected editing text, or -1 if not set
	Start int32

	// The length of selected editing text, or -1 if not set
	Length int32
}

// TextInputEvent is a keyboard text input event structure (event.text.*).
//
// [EVENT_TEXT_INPUT]
//
// Defined in [<SDL3/SDL_events.h>]
//
// This event will never be delivered unless text input is enabled by calling
// [StartTextInput](). Text input is disabled by default!
//
// This struct is available since SDL 3.2.0.
//
// See Also:
//
//   - [StartTextInput]
//   - [StopTextInput]
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type TextInputEvent struct {
	// The window with keyboard focus, if any
	WindowID WindowID

	// The input text, UTF-8 encoded
	Text string
}

// TouchFingerEvent is a touch finger event structure (event.tfinger.*).
//
// [EVENT_FINGER_DOWN], [EVENT_FINGER_UP], [EVENT_FINGER_MOTION], or
// [EVENT_FINGER_CANCELED]
//
// Defined in [<SDL3/SDL_events.h>]
//
// Coordinates in this event are normalized. `x` and `y` are normalized to a
// range between 0.0f and 1.0f, relative to the window, so (0,0) is the top
// left and (1,1) is the bottom right. Delta coordinates `dx` and `dy` are
// normalized in the ranges of -1.0f (traversed all the way from the bottom or
// right to all the way up or left) to 1.0f (traversed all the way from the
// top or left to all the way down or right).
//
// Note that while the coordinates are _normalized_, they are not _clamped_,
// which means in some circumstances you can get a value outside of this
// range. For example, a renderer using logical presentation might give a
// negative value when the touch is in the letterboxing. Some platforms might
// report a touch outside of the window, which will also be outside of the
// range.
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type TouchFingerEvent struct {
	// The touch device id
	TouchID TouchID

	FingerID FingerID

	// Normalized in the range 0...1
	X float32

	// Normalized in the range 0...1
	Y float32

	// Normalized in the range -1...1
	Dx float32

	// Normalized in the range -1...1
	Dy float32

	// Normalized in the range 0...1
	Pressure float32

	// The window underneath the finger, if any
	WindowID WindowID
}

// UserEvent is a user-defined event type (event.user.*).
//
// [EVENT_USER] through [EVENT_LAST], Uint32 because these are not in the
// [EventType] enumeration
//
// Defined in [<SDL3/SDL_events.h>]
//
// This event is unique; it is never created by SDL, but only by the
// application. The event can be pushed onto the event queue using
// [PushEvent](). The contents of the structure members are completely up to
// the programmer; the only requirement is that “type” is a value obtained from
// [RegisterEvents]().
//
// This struct is available since SDL 3.2.0.
//
// Code Examples
//
//	extern Sint32 my_event_code;
//	extern void *significant_data;
//	extern void *some_other_data;
//
//	const Uint32 myEventType = SDL_RegisterEvents(1);
//	if (myEventType != 0) {
//	    SDL_Event event;
//	    SDL_zero(event);
//	    event.type = myEventType;
//	    event.user.code = my_event_code;
//	    event.user.data1 = significant_data;
//	    event.user.data2 = some_other_data;
//	    SDL_PushEvent(&event);
//	}
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type UserEvent struct {
	// The associated window if any
	WindowID WindowID

	// User defined event code
	Code int32

	// User defined data pointer
	Data1 uintptr

	// User defined data pointer
	Data2 uintptr
}

// WindowEvent is a window state change event data (event.window.*).
//
// EVENT_WINDOW_*
//
// Defined in [<SDL3/SDL_events.h>]
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type WindowEvent struct {
	// The associated window
	WindowID WindowID

	// event dependent data
	Data1 int32

	// event dependent data
	Data2 int32
}

func convertEvent(cEvent *C.SDL_Event) *Event {
	cCommonEvent := (*C.SDL_CommonEvent)(unsafe.Pointer(cEvent))

	result := &Event{
		CommonEvent: CommonEvent{
			Type:      EventType(cCommonEvent._type),
			Timestamp: uint64(cCommonEvent.timestamp),
		},
		internal: nil,
	}
	t := result.Type

	switch {
	case t >= EVENT_DISPLAY_FIRST && t <= EVENT_DISPLAY_LAST:
		e := (*C.SDL_DisplayEvent)(unsafe.Pointer(cEvent))
		result.internal = &DisplayEvent{
			DisplayID: DisplayID(e.displayID),
			Data1:     int32(e.data1),
			Data2:     int32(e.data2),
		}
	case t >= EVENT_WINDOW_FIRST && t <= EVENT_WINDOW_LAST:
		e := (*C.SDL_WindowEvent)(unsafe.Pointer(cEvent))
		result.internal = &WindowEvent{
			WindowID: WindowID(e.windowID),
			Data1:    int32(e.data1),
			Data2:    int32(e.data2),
		}
	case t == EVENT_KEYBOARD_ADDED || t == EVENT_KEYBOARD_REMOVED:
		e := (*C.SDL_KeyboardDeviceEvent)(unsafe.Pointer(cEvent))
		result.internal = &KeyboardDeviceEvent{
			Which: KeyboardID(e.which),
		}
	case t == EVENT_KEY_DOWN || t == EVENT_KEY_UP:
		e := (*C.SDL_KeyboardEvent)(unsafe.Pointer(cEvent))
		result.internal = &KeyboardEvent{
			WindowID: WindowID(e.windowID),
			Which:    KeyboardID(e.which),
			Scancode: Scancode(e.scancode),
			Key:      Keycode(e.key),
			Mod:      Keymod(e.mod),
			Raw:      uint16(e.raw),
			Down:     bool(e.down),
			Repeat:   bool(e.repeat),
		}
	case t == EVENT_TEXT_EDITING:
		e := (*C.SDL_TextEditingEvent)(unsafe.Pointer(cEvent))
		result.internal = &TextEditingEvent{
			WindowID: WindowID(e.windowID),
			Text:     C.GoString(e.text),
			Start:    int32(e.start),
			Length:   int32(e.length),
		}
	case t == EVENT_TEXT_EDITING_CANDIDATES:
		e := (*C.SDL_TextEditingCandidatesEvent)(unsafe.Pointer(cEvent))
		result.internal = &TextEditingCandidatesEvent{
			WindowID:          WindowID(e.windowID),
			Candidates:        toStrings(e.candidates, int32(e.num_candidates)),
			SelectedCandidate: int32(e.selected_candidate),
			Horizontal:        bool(e.horizontal),
		}
	case t == EVENT_TEXT_INPUT:
		e := (*C.SDL_TextInputEvent)(unsafe.Pointer(cEvent))
		result.internal = &TextInputEvent{
			WindowID: WindowID(e.windowID),
			Text:     C.GoString(e.text),
		}
	case t == EVENT_MOUSE_ADDED || t == EVENT_MOUSE_REMOVED:
		e := (*C.SDL_MouseDeviceEvent)(unsafe.Pointer(cEvent))
		result.internal = &MouseDeviceEvent{
			Which: MouseID(e.which),
		}
	case t == EVENT_MOUSE_MOTION:
		e := (*C.SDL_MouseMotionEvent)(unsafe.Pointer(cEvent))
		result.internal = &MouseMotionEvent{
			WindowID: WindowID(e.windowID),
			Which:    MouseID(e.which),
			State:    MouseButtonFlags(e.state),
			X:        float32(e.x),
			Y:        float32(e.y),
			Xrel:     float32(e.xrel),
			Yrel:     float32(e.yrel),
		}
	case t == EVENT_MOUSE_BUTTON_DOWN || t == EVENT_MOUSE_BUTTON_UP:
		e := (*C.SDL_MouseButtonEvent)(unsafe.Pointer(cEvent))
		result.internal = &MouseButtonEvent{
			WindowID: WindowID(e.windowID),
			Which:    MouseID(e.which),
			Button:   uint8(e.button),
			Down:     bool(e.down),
			Clicks:   uint8(e.clicks),
			X:        float32(e.x),
			Y:        float32(e.y),
		}
	case t == EVENT_MOUSE_WHEEL:
		e := (*C.SDL_MouseWheelEvent)(unsafe.Pointer(cEvent))
		result.internal = &MouseWheelEvent{
			WindowID:  WindowID(e.windowID),
			Which:     MouseID(e.which),
			X:         float32(e.x),
			Y:         float32(e.y),
			Direction: MouseWheelDirection(e.direction),
			MouseX:    float32(e.mouse_x),
			MouseY:    float32(e.mouse_y),
			IntegerX:  int32(e.integer_x),
			IntegerY:  int32(e.integer_y),
		}
	case t == EVENT_JOYSTICK_ADDED || t == EVENT_JOYSTICK_REMOVED ||
		t == EVENT_JOYSTICK_UPDATE_COMPLETE:

		e := (*C.SDL_JoyDeviceEvent)(unsafe.Pointer(cEvent))
		result.internal = &JoyDeviceEvent{
			Which: JoystickID(e.which),
		}
	case t == EVENT_JOYSTICK_AXIS_MOTION:
		e := (*C.SDL_JoyAxisEvent)(unsafe.Pointer(cEvent))
		result.internal = &JoyAxisEvent{
			Which: JoystickID(e.which),
			Axis:  uint8(e.axis),
			Value: int16(e.value),
		}
	case t == EVENT_JOYSTICK_BALL_MOTION:
		e := (*C.SDL_JoyBallEvent)(unsafe.Pointer(cEvent))
		result.internal = &JoyBallEvent{
			Which: JoystickID(e.which),
			Ball:  uint8(e.ball),
			Xrel:  int16(e.xrel),
			Yrel:  int16(e.yrel),
		}
	case t == EVENT_JOYSTICK_HAT_MOTION:
		e := (*C.SDL_JoyHatEvent)(unsafe.Pointer(cEvent))
		result.internal = &JoyHatEvent{
			Which: JoystickID(e.which),
			Hat:   uint8(e.hat),
			Value: JoystickHat(e.value),
		}
	case t == EVENT_JOYSTICK_BUTTON_DOWN || t == EVENT_JOYSTICK_BUTTON_UP:
		e := (*C.SDL_JoyButtonEvent)(unsafe.Pointer(cEvent))
		result.internal = &JoyButtonEvent{
			Which:  JoystickID(e.which),
			Button: uint8(e.button),
			Down:   bool(e.down),
		}
	case t == EVENT_JOYSTICK_BATTERY_UPDATED:
		e := (*C.SDL_JoyBatteryEvent)(unsafe.Pointer(cEvent))
		result.internal = &JoyBatteryEvent{
			Which:   JoystickID(e.which),
			State:   PowerState(e.state),
			Percent: int(e.percent),
		}
	case t == EVENT_GAMEPAD_ADDED || t == EVENT_GAMEPAD_REMOVED ||
		t == EVENT_GAMEPAD_REMAPPED || t == EVENT_GAMEPAD_UPDATE_COMPLETE ||
		t == EVENT_GAMEPAD_STEAM_HANDLE_UPDATED:

		e := (*C.SDL_GamepadDeviceEvent)(unsafe.Pointer(cEvent))
		result.internal = &GamepadDeviceEvent{
			Which: JoystickID(e.which),
		}
	case t == EVENT_GAMEPAD_AXIS_MOTION:
		e := (*C.SDL_GamepadAxisEvent)(unsafe.Pointer(cEvent))
		result.internal = &GamepadAxisEvent{
			Which: JoystickID(e.which),
			Axis:  GamepadAxis(e.axis),
			Value: int16(e.value),
		}
	case t == EVENT_GAMEPAD_BUTTON_DOWN || t == EVENT_GAMEPAD_BUTTON_UP:
		e := (*C.SDL_GamepadButtonEvent)(unsafe.Pointer(cEvent))
		result.internal = &GamepadButtonEvent{
			Which:  JoystickID(e.which),
			Button: uint8(e.button),
			Down:   bool(e.down),
		}
	case t == EVENT_GAMEPAD_TOUCHPAD_DOWN ||
		t == EVENT_GAMEPAD_TOUCHPAD_MOTION || t == EVENT_GAMEPAD_TOUCHPAD_UP:

		e := (*C.SDL_GamepadTouchpadEvent)(unsafe.Pointer(cEvent))
		result.internal = &GamepadTouchpadEvent{
			Which:    JoystickID(e.which),
			Touchpad: int32(e.touchpad),
			Finger:   int32(e.finger),
			X:        float32(e.x),
			Y:        float32(e.y),
			Pressure: float32(e.pressure),
		}
	case t == EVENT_GAMEPAD_SENSOR_UPDATE:
		e := (*C.SDL_GamepadSensorEvent)(unsafe.Pointer(cEvent))
		result.internal = &GamepadSensorEvent{
			Which:           JoystickID(e.which),
			Sensor:          SensorType(e.sensor),
			Data:            *(*[3]float32)(unsafe.Pointer(&e.data)),
			SensorTimestamp: uint64(e.sensor_timestamp),
		}
	case t == EVENT_AUDIO_DEVICE_ADDED || t == EVENT_AUDIO_DEVICE_REMOVED ||
		t == EVENT_AUDIO_DEVICE_FORMAT_CHANGED:

		e := (*C.SDL_AudioDeviceEvent)(unsafe.Pointer(cEvent))
		result.internal = &AudioDeviceEvent{
			Which:     AudioDeviceID(e.which),
			Recording: bool(e.recording),
		}
	case t == EVENT_CAMERA_DEVICE_ADDED || t == EVENT_CAMERA_DEVICE_REMOVED ||
		t == EVENT_CAMERA_DEVICE_APPROVED || t == EVENT_CAMERA_DEVICE_DENIED:

		e := (*C.SDL_CameraDeviceEvent)(unsafe.Pointer(cEvent))
		result.internal = &CameraDeviceEvent{
			Which: CameraID(e.which),
		}
	case t == EVENT_SENSOR_UPDATE:
		e := (*C.SDL_SensorEvent)(unsafe.Pointer(cEvent))
		result.internal = &SensorEvent{
			Which:           SensorID(e.which),
			Data:            *(*[6]float32)(unsafe.Pointer(&e.data)),
			SensorTimestamp: uint64(e.sensor_timestamp),
		}
	case t == EVENT_QUIT:
		result.internal = &QuitEvent{}
	case t >= EVENT_USER && t <= EVENT_LAST:
		e := (*C.SDL_UserEvent)(unsafe.Pointer(cEvent))
		result.internal = &UserEvent{
			WindowID: WindowID(e.windowID),
			Code:     int32(e.code),
			Data1:    uintptr(e.data1),
			Data2:    uintptr(e.data2),
		}
	case t == EVENT_FINGER_DOWN || t == EVENT_FINGER_UP ||
		t == EVENT_FINGER_MOTION || t == EVENT_FINGER_CANCELED:

		e := (*C.SDL_TouchFingerEvent)(unsafe.Pointer(cEvent))
		result.internal = &TouchFingerEvent{
			TouchID:  TouchID(e.touchID),
			FingerID: FingerID(e.fingerID),
			X:        float32(e.x),
			Y:        float32(e.y),
			Dx:       float32(e.dx),
			Dy:       float32(e.dy),
			Pressure: float32(e.pressure),
			WindowID: WindowID(e.windowID),
		}
	case t == EVENT_PINCH_BEGIN || t == EVENT_PINCH_UPDATE ||
		t == EVENT_PINCH_END:

		e := (*C.SDL_PinchFingerEvent)(unsafe.Pointer(cEvent))
		result.internal = &PinchFingerEvent{
			Scale:    float32(e.scale),
			WindowID: WindowID(e.windowID),
		}
	case t == EVENT_PEN_PROXIMITY_IN || t == EVENT_PEN_PROXIMITY_OUT:
		e := (*C.SDL_PenProximityEvent)(unsafe.Pointer(cEvent))
		result.internal = &PenProximityEvent{
			WindowID: WindowID(e.windowID),
			Which:    PenID(e.which),
		}
	case t == EVENT_PEN_DOWN || t == EVENT_PEN_UP:
		e := (*C.SDL_PenTouchEvent)(unsafe.Pointer(cEvent))
		result.internal = &PenTouchEvent{
			WindowID: WindowID(e.windowID),
			Which:    PenID(e.which),
			PenState: PenInputFlags(e.pen_state),
			X:        float32(e.x),
			Y:        float32(e.y),
			Eraser:   bool(e.eraser),
			Down:     bool(e.down),
		}
	case t == EVENT_PEN_MOTION:
		e := (*C.SDL_PenMotionEvent)(unsafe.Pointer(cEvent))
		result.internal = &PenMotionEvent{
			WindowID: WindowID(e.windowID),
			Which:    PenID(e.which),
			PenState: PenInputFlags(e.pen_state),
			X:        float32(e.x),
			Y:        float32(e.y),
		}
	case t == EVENT_PEN_BUTTON_DOWN || t == EVENT_PEN_BUTTON_UP:
		e := (*C.SDL_PenButtonEvent)(unsafe.Pointer(cEvent))
		result.internal = &PenButtonEvent{
			WindowID: WindowID(e.windowID),
			Which:    PenID(e.which),
			PenState: PenInputFlags(e.pen_state),
			X:        float32(e.x),
			Y:        float32(e.y),
			Button:   uint8(e.button),
			Down:     bool(e.down),
		}
	case t == EVENT_PEN_AXIS:
		e := (*C.SDL_PenAxisEvent)(unsafe.Pointer(cEvent))
		result.internal = &PenAxisEvent{
			WindowID: WindowID(e.windowID),
			Which:    PenID(e.which),
			PenState: PenInputFlags(e.pen_state),
			X:        float32(e.x),
			Y:        float32(e.y),
			Axis:     PenAxis(e.axis),
			Value:    float32(e.value),
		}
	case t == EVENT_RENDER_TARGETS_RESET || t == EVENT_RENDER_DEVICE_RESET ||
		t == EVENT_RENDER_DEVICE_LOST:

		e := (*C.SDL_RenderEvent)(unsafe.Pointer(cEvent))
		result.internal = &RenderEvent{
			WindowID: WindowID(e.windowID),
		}
	case t == EVENT_DROP_BEGIN || t == EVENT_DROP_FILE ||
		t == EVENT_DROP_TEXT || t == EVENT_DROP_COMPLETE || t == EVENT_DROP_POSITION:

		e := (*C.SDL_DropEvent)(unsafe.Pointer(cEvent))
		result.internal = &DropEvent{
			WindowID: WindowID(e.windowID),
			X:        float32(e.x),
			Y:        float32(e.y),
			Source:   C.GoString(e.source),
			Data:     C.GoString(e.data),
		}
	case t == EVENT_CLIPBOARD_UPDATE:
		e := (*C.SDL_ClipboardEvent)(unsafe.Pointer(cEvent))
		result.internal = &ClipboardEvent{
			Owner:     bool(e.owner),
			MIMETypes: toStrings(e.mime_types, int32(e.num_mime_types)),
		}
	}

	return result
}
