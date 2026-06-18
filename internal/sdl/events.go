package sdl

// #include <SDL3/SDL.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func toStrings(cStrings **C.char, num int32) []string {
	var results []string
	for _, cc := range unsafe.Slice(cStrings, num) {
		results = append(results, C.GoString(cc))
	}
	return results
}

type Event any

// # SDL_AudioDeviceEvent
//
// Audio device event structure (event.adevice.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_AudioDeviceEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_AUDIO_DEVICE_ADDED, or SDL_EVENT_AUDIO_DEVICE_REMOVED, or SDL_EVENT_AUDIO_DEVICE_FORMAT_CHANGED */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_AudioDeviceID which;       /**< SDL_AudioDeviceID for the device being added or removed or changing */
//	    bool recording; /**< false if a playback device, true if a recording device. */
//	    Uint8 padding1;
//	    Uint8 padding2;
//	    Uint8 padding3;
//	} SDL_AudioDeviceEvent;
//
// # Remarks
//
// Note that SDL will send a
// [SDL_EVENT_AUDIO_DEVICE_ADDED](SDL_EVENT_AUDIO_DEVICE_ADDED) event for
// every device it discovers during initialization. After that, this event
// will only arrive when a device is hotplugged during the program's run.
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type AudioDeviceEvent struct {
	// SDL_EVENT_AUDIO_DEVICE_ADDED, or SDL_EVENT_AUDIO_DEVICE_REMOVED, or
	// SDL_EVENT_AUDIO_DEVICE_FORMAT_CHANGED
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// SDL_AudioDeviceID for the device being added or removed or changing
	Which AudioDeviceID

	// false if a playback device, true if a recording device.
	recording bool
}

// # SDL_CameraDeviceEvent
//
// Camera device event structure (event.cdevice.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_CameraDeviceEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_CAMERA_DEVICE_ADDED, SDL_EVENT_CAMERA_DEVICE_REMOVED, SDL_EVENT_CAMERA_DEVICE_APPROVED, SDL_EVENT_CAMERA_DEVICE_DENIED */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_CameraID which;       /**< SDL_CameraID for the device being added or removed or changing */
//	} SDL_CameraDeviceEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type CameraDeviceEvent struct {
	// SDL_EVENT_CAMERA_DEVICE_ADDED, SDL_EVENT_CAMERA_DEVICE_REMOVED,
	// SDL_EVENT_CAMERA_DEVICE_APPROVED, SDL_EVENT_CAMERA_DEVICE_DENIED
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// SDL_CameraID for the device being added or removed or changing
	Which CameraID
}

// # SDL_ClipboardEvent
//
// An event triggered when the clipboard contents have changed (event.clipboard.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_ClipboardEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_CLIPBOARD_UPDATE */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    bool owner;         /**< are we owning the clipboard (internal update) */
//	    Sint32 num_mime_types;   /**< number of mime types */
//	    const char **mime_types; /**< current mime types */
//	} SDL_ClipboardEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type ClipboardEvent struct {
	// SDL_EVENT_CLIPBOARD_UPDATE
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// are we owning the clipboard (internal update)
	Owner bool

	// current mime types
	MIMETypes []string
}

// # SDL_DisplayEvent
//
// Display state change event data (event.display.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_DisplayEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_DISPLAY_* */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_DisplayID displayID;/**< The associated display */
//	    Sint32 data1;       /**< event dependent data */
//	    Sint32 data2;       /**< event dependent data */
//	} SDL_DisplayEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type DisplayEvent struct {
	Type      EventType // SDL_EVENT_DISPLAY_*
	Timestamp uint64    // In nanoseconds, populated using SDL_GetTicksNS()
	DisplayID DisplayID // The associated display
	Data1     int32     // event dependent data
	Data2     int32     // event dependent data
}

// # SDL_DropEvent
//
// An event used to drop text or request a file open by the system (event.drop.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_DropEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_DROP_BEGIN or SDL_EVENT_DROP_FILE or SDL_EVENT_DROP_TEXT or SDL_EVENT_DROP_COMPLETE or SDL_EVENT_DROP_POSITION */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID;    /**< The window that was dropped on, if any */
//	    float x;            /**< X coordinate, relative to window (not on begin) */
//	    float y;            /**< Y coordinate, relative to window (not on begin) */
//	    const char *source; /**< The source app that sent this drop event, or NULL if that isn't available */
//	    const char *data;   /**< The text for SDL_EVENT_DROP_TEXT and the file name for SDL_EVENT_DROP_FILE, NULL for other events */
//	} SDL_DropEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type DropEvent struct {
	// SDL_EVENT_DROP_BEGIN or SDL_EVENT_DROP_FILE or SDL_EVENT_DROP_TEXT or
	// SDL_EVENT_DROP_COMPLETE or SDL_EVENT_DROP_POSITION
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The window that was dropped on, if any
	WindowID WindowID

	// X coordinate, relative to window (not on begin)
	X float32

	// Y coordinate, relative to window (not on begin)
	Y float32

	// The source app that sent this drop event, or NULL if that isn't
	// available
	Source string

	// The text for SDL_EVENT_DROP_TEXT and the file name for
	// SDL_EVENT_DROP_FILE, NULL for other events
	Data string
}

// # SDL_GamepadAxisEvent
//
// Gamepad axis motion event structure (event.gaxis.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_GamepadAxisEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_GAMEPAD_AXIS_MOTION */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which; /**< The joystick instance id */
//	    Uint8 axis;         /**< The gamepad axis (SDL_GamepadAxis) */
//	    Uint8 padding1;
//	    Uint8 padding2;
//	    Uint8 padding3;
//	    Sint16 value;       /**< The axis value (range: -32768 to 32767) */
//	    Uint16 padding4;
//	} SDL_GamepadAxisEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadAxisEvent struct {
	// SDL_EVENT_GAMEPAD_AXIS_MOTION
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID

	// The gamepad axis (SDL_GamepadAxis)
	Axis uint8

	// The axis value (range: -32768 to 32767)
	Value int16
}

// # SDL_GamepadButtonEvent
//
// Gamepad button event structure (event.gbutton.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_GamepadButtonEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_GAMEPAD_BUTTON_DOWN or SDL_EVENT_GAMEPAD_BUTTON_UP */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which; /**< The joystick instance id */
//	    Uint8 button;       /**< The gamepad button (SDL_GamepadButton) */
//	    bool down;      /**< true if the button is pressed */
//	    Uint8 padding1;
//	    Uint8 padding2;
//	} SDL_GamepadButtonEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadButtonEvent struct {
	// SDL_EVENT_GAMEPAD_BUTTON_DOWN or SDL_EVENT_GAMEPAD_BUTTON_UP
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID

	// The gamepad button (SDL_GamepadButton)
	Button uint8

	// true if the button is pressed
	Down bool
}

// # SDL_GamepadCapSenseEvent
//
// Gamepad capsense event structure (event.gcapsense.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_GamepadCapSenseEvent
//	{
//	    SDL_EventType type;     /**< SDL_EVENT_GAMEPAD_CAPSENSE_TOUCH or SDL_EVENT_GAMEPAD_CAPSENSE_RELEASE */
//	    Uint32 reserved;
//	    Uint64 timestamp;       /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which;   /**< The joystick instance id */
//	    Uint8 capsense;         /**< The capsense type (SDL_GamepadCapSenseType) */
//	    bool down;              /**< true if the capsense is touched */
//	    Uint8 padding1;
//	    Uint8 padding2;
//	} SDL_GamepadCapSenseEvent;
//
// # Version
//
// This struct is available since SDL 3.6.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadCapSenseEvent struct {
	// SDL_EVENT_GAMEPAD_CAPSENSE_TOUCH or SDL_EVENT_GAMEPAD_CAPSENSE_RELEASE
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID

	// The capsense type (SDL_GamepadCapSenseType)
	Capsense uint8

	// true if the capsense is touched
	Down bool
}

// # SDL_GamepadDeviceEvent
//
// Gamepad device event structure (event.gdevice.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_GamepadDeviceEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_GAMEPAD_ADDED, SDL_EVENT_GAMEPAD_REMOVED, or SDL_EVENT_GAMEPAD_REMAPPED, SDL_EVENT_GAMEPAD_UPDATE_COMPLETE or SDL_EVENT_GAMEPAD_STEAM_HANDLE_UPDATED */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which;       /**< The joystick instance id */
//	} SDL_GamepadDeviceEvent;
//
// # Remarks
//
// Joysticks that are supported gamepads receive both an
// [SDL_JoyDeviceEvent](SDL_JoyDeviceEvent) and an
// [SDL_GamepadDeviceEvent](SDL_GamepadDeviceEvent).
//
// SDL will send GAMEPAD_ADDED events for joysticks that are already plugged
// in during [SDL_Init](SDL_Init)() and are recognized as gamepads. It will
// also send events for joysticks that get gamepad mappings at runtime.
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// # See Also
//
//   - [JoyDeviceEvent]
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadDeviceEvent struct {
	// SDL_EVENT_GAMEPAD_ADDED, SDL_EVENT_GAMEPAD_REMOVED, or
	// SDL_EVENT_GAMEPAD_REMAPPED, SDL_EVENT_GAMEPAD_UPDATE_COMPLETE or
	// SDL_EVENT_GAMEPAD_STEAM_HANDLE_UPDATED
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID
}

// # SDL_GamepadSensorEvent
//
// Gamepad sensor event structure (event.gsensor.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_GamepadSensorEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_GAMEPAD_SENSOR_UPDATE */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which; /**< The joystick instance id */
//	    Sint32 sensor;      /**< The type of the sensor, one of the values of SDL_SensorType */
//	    float data[3];      /**< Up to 3 values from the sensor, as defined in SDL_sensor.h */
//	    Uint64 sensor_timestamp; /**< The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock */
//	} SDL_GamepadSensorEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadSensorEvent struct {
	// SDL_EVENT_GAMEPAD_SENSOR_UPDATE
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID

	// The type of the sensor, one of the values of SDL_SensorType
	Sensor int32

	// Up to 3 values from the sensor, as defined in SDL_sensor.h
	Data [3]float32

	// The timestamp of the sensor reading in nanoseconds, not necessarily
	// synchronized with the system clock
	SensorTimestamp uint64
}

// # SDL_GamepadTouchpadEvent
//
// Gamepad touchpad event structure (event.gtouchpad.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_GamepadTouchpadEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_GAMEPAD_TOUCHPAD_DOWN or SDL_EVENT_GAMEPAD_TOUCHPAD_MOTION or SDL_EVENT_GAMEPAD_TOUCHPAD_UP */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which; /**< The joystick instance id */
//	    Sint32 touchpad;    /**< The index of the touchpad */
//	    Sint32 finger;      /**< The index of the finger on the touchpad */
//	    float x;            /**< Normalized in the range 0...1 with 0 being on the left */
//	    float y;            /**< Normalized in the range 0...1 with 0 being at the top */
//	    float pressure;     /**< Normalized in the range 0...1 */
//	} SDL_GamepadTouchpadEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type GamepadTouchpadEvent struct {
	// SDL_EVENT_GAMEPAD_TOUCHPAD_DOWN or SDL_EVENT_GAMEPAD_TOUCHPAD_MOTION or
	// SDL_EVENT_GAMEPAD_TOUCHPAD_UP
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

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

// # SDL_JoyAxisEvent
//
// Joystick axis motion event structure (event.jaxis.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_JoyAxisEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_JOYSTICK_AXIS_MOTION */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which; /**< The joystick instance id */
//	    Uint8 axis;         /**< The joystick axis index */
//	    Uint8 padding1;
//	    Uint8 padding2;
//	    Uint8 padding3;
//	    Sint16 value;       /**< The axis value (range: -32768 to 32767) */
//	    Uint16 padding4;
//	} SDL_JoyAxisEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyAxisEvent struct {
	// SDL_EVENT_JOYSTICK_AXIS_MOTION
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID

	// The joystick axis index
	Axis uint8

	// The axis value (range: -32768 to 32767)
	Value int16
}

// # SDL_JoyBallEvent
//
// Joystick trackball motion event structure (event.jball.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_JoyBallEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_JOYSTICK_BALL_MOTION */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which; /**< The joystick instance id */
//	    Uint8 ball;         /**< The joystick trackball index */
//	    Uint8 padding1;
//	    Uint8 padding2;
//	    Uint8 padding3;
//	    Sint16 xrel;        /**< The relative motion in the X direction */
//	    Sint16 yrel;        /**< The relative motion in the Y direction */
//	} SDL_JoyBallEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyBallEvent struct {
	// SDL_EVENT_JOYSTICK_BALL_MOTION
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID

	// The joystick trackball index
	Ball uint8

	// The relative motion in the X direction
	Xrel int16

	// The relative motion in the Y direction
	Yrel int16
}

// # SDL_JoyBatteryEvent
//
// Joystick battery level change event structure (event.jbattery.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_JoyBatteryEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_JOYSTICK_BATTERY_UPDATED */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which; /**< The joystick instance id */
//	    SDL_PowerState state; /**< The joystick battery state */
//	    int percent;          /**< The joystick battery percent charge remaining */
//	} SDL_JoyBatteryEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyBatteryEvent struct {
	// SDL_EVENT_JOYSTICK_BATTERY_UPDATED
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID

	// The joystick battery state
	State PowerState

	// The joystick battery percent charge remaining
	Percent int
}

// # SDL_JoyButtonEvent
//
// Joystick button event structure (event.jbutton.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_JoyButtonEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_JOYSTICK_BUTTON_DOWN or SDL_EVENT_JOYSTICK_BUTTON_UP */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which; /**< The joystick instance id */
//	    Uint8 button;       /**< The joystick button index */
//	    bool down;      /**< true if the button is pressed */
//	    Uint8 padding1;
//	    Uint8 padding2;
//	} SDL_JoyButtonEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyButtonEvent struct {
	// SDL_EVENT_JOYSTICK_BUTTON_DOWN or SDL_EVENT_JOYSTICK_BUTTON_UP
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID

	// The joystick button index
	Button uint8

	// true if the button is pressed
	Down bool
}

// # SDL_JoyDeviceEvent
//
// Joystick device event structure (event.jdevice.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_JoyDeviceEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_JOYSTICK_ADDED or SDL_EVENT_JOYSTICK_REMOVED or SDL_EVENT_JOYSTICK_UPDATE_COMPLETE */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which;       /**< The joystick instance id */
//	} SDL_JoyDeviceEvent;
//
// # Remarks
//
// SDL will send JOYSTICK_ADDED events for devices that are already plugged in
// during [SDL_Init](SDL_Init).
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// # See Also
//
//   - [GamepadDeviceEvent]
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyDeviceEvent struct {
	// SDL_EVENT_JOYSTICK_ADDED or SDL_EVENT_JOYSTICK_REMOVED or
	// SDL_EVENT_JOYSTICK_UPDATE_COMPLETE
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID
}

// # SDL_JoyHatEvent
//
// Joystick hat position change event structure (event.jhat.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_JoyHatEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_JOYSTICK_HAT_MOTION */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_JoystickID which; /**< The joystick instance id */
//	    Uint8 hat;          /**< The joystick hat index */
//	    Uint8 value;        /**< The hat position value.
//	                         *   \sa SDL_HAT_LEFTUP SDL_HAT_UP SDL_HAT_RIGHTUP
//	                         *   \sa SDL_HAT_LEFT SDL_HAT_CENTERED SDL_HAT_RIGHT
//	                         *   \sa SDL_HAT_LEFTDOWN SDL_HAT_DOWN SDL_HAT_RIGHTDOWN
//	                         *
//	                         *   Note that zero means the POV is centered.
//	                         */
//	    Uint8 padding1;
//	    Uint8 padding2;
//	} SDL_JoyHatEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type JoyHatEvent struct {
	// SDL_EVENT_JOYSTICK_HAT_MOTION
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The joystick instance id
	Which JoystickID

	// The joystick hat index
	Hat uint8

	// The hat position value.
	//
	// Note that zero means the POV is centered.
	Value JoystickHat
}

// # SDL_KeyboardDeviceEvent
//
// Keyboard device event structure (event.kdevice.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_KeyboardDeviceEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_KEYBOARD_ADDED or SDL_EVENT_KEYBOARD_REMOVED */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_KeyboardID which;   /**< The keyboard instance id */
//	} SDL_KeyboardDeviceEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type KeyboardDeviceEvent struct {
	// SDL_EVENT_KEYBOARD_ADDED or SDL_EVENT_KEYBOARD_REMOVED
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The keyboard instance id
	Which KeyboardID
}

// # SDL_KeyboardEvent
//
// Keyboard button event structure (event.key.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_KeyboardEvent
//	{
//	    SDL_EventType type;     /**< SDL_EVENT_KEY_DOWN or SDL_EVENT_KEY_UP */
//	    Uint32 reserved;
//	    Uint64 timestamp;       /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID;  /**< The window with keyboard focus, if any */
//	    SDL_KeyboardID which;   /**< The keyboard instance id, or 0 if unknown or virtual */
//	    SDL_Scancode scancode;  /**< SDL physical key code */
//	    SDL_Keycode key;        /**< SDL virtual key code */
//	    SDL_Keymod mod;         /**< current key modifiers */
//	    Uint16 raw;             /**< The platform dependent scancode for this event */
//	    bool down;              /**< true if the key is pressed */
//	    bool repeat;            /**< true if this is a key repeat */
//	} SDL_KeyboardEvent;
//
// # Remarks
//
// The `key` is the base [SDL_Keycode](SDL_Keycode) generated by pressing the
// `scancode` using the current keyboard layout, applying any options
// specified in [SDL_HINT_KEYCODE_OPTIONS](SDL_HINT_KEYCODE_OPTIONS). You can
// get the [SDL_Keycode](SDL_Keycode) corresponding to the event scancode and
// modifiers directly from the keyboard layout, bypassing
// [SDL_HINT_KEYCODE_OPTIONS](SDL_HINT_KEYCODE_OPTIONS), by calling
// [SDL_GetKeyFromScancode](SDL_GetKeyFromScancode)().
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// # See Also
//
//   - [GetKeyFromScancode]
//   - [HINT_KEYCODE_OPTIONS]
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type KeyboardEvent struct {
	// SDL_EVENT_KEY_DOWN or SDL_EVENT_KEY_UP
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

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

// # SDL_MouseButtonEvent
//
// Mouse button event structure (event.button.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_MouseButtonEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_MOUSE_BUTTON_DOWN or SDL_EVENT_MOUSE_BUTTON_UP */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The window with mouse focus, if any */
//	    SDL_MouseID which;  /**< The mouse instance id in relative mode, SDL_TOUCH_MOUSEID for touch events, or 0 */
//	    Uint8 button;       /**< The mouse button index */
//	    bool down;          /**< true if the button is pressed */
//	    Uint8 clicks;       /**< 1 for single-click, 2 for double-click, etc. */
//	    Uint8 padding;
//	    float x;            /**< X coordinate, relative to window */
//	    float y;            /**< Y coordinate, relative to window */
//	} SDL_MouseButtonEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type MouseButtonEvent struct {
	// SDL_EVENT_MOUSE_BUTTON_DOWN or SDL_EVENT_MOUSE_BUTTON_UP
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The window with mouse focus, if any
	WindowID WindowID

	// The mouse instance id in relative mode, SDL_TOUCH_MOUSEID for touch
	// events, or 0
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

// # SDL_MouseDeviceEvent
//
// Mouse device event structure (event.mdevice.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_MouseDeviceEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_MOUSE_ADDED or SDL_EVENT_MOUSE_REMOVED */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_MouseID which;  /**< The mouse instance id */
//	} SDL_MouseDeviceEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type MouseDeviceEvent struct {
	// SDL_EVENT_MOUSE_ADDED or SDL_EVENT_MOUSE_REMOVED
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The mouse instance id
	Which MouseID
}

// # SDL_MouseMotionEvent
//
// Mouse motion event structure (event.motion.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_MouseMotionEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_MOUSE_MOTION */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The window with mouse focus, if any */
//	    SDL_MouseID which;  /**< The mouse instance id in relative mode, SDL_TOUCH_MOUSEID for touch events, SDL_PEN_MOUSEID for pen events, or 0 */
//	    SDL_MouseButtonFlags state;       /**< The current button state */
//	    float x;            /**< X coordinate, relative to window */
//	    float y;            /**< Y coordinate, relative to window */
//	    float xrel;         /**< The relative motion in the X direction */
//	    float yrel;         /**< The relative motion in the Y direction */
//	} SDL_MouseMotionEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type MouseMotionEvent struct {
	// SDL_EVENT_MOUSE_MOTION
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The window with mouse focus, if any
	WindowID WindowID

	// The mouse instance id in relative mode, SDL_TOUCH_MOUSEID for touch
	// events, SDL_PEN_MOUSEID for pen events, or 0
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

// # SDL_MouseWheelEvent
//
// Mouse wheel event structure (event.wheel.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_MouseWheelEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_MOUSE_WHEEL */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The window with mouse focus, if any */
//	    SDL_MouseID which;  /**< The mouse instance id in relative mode or 0 */
//	    float x;            /**< The amount scrolled horizontally, positive to the right and negative to the left */
//	    float y;            /**< The amount scrolled vertically, positive away from the user and negative toward the user */
//	    SDL_MouseWheelDirection direction; /**< Set to one of the SDL_MOUSEWHEEL_* defines. When FLIPPED the values in X and Y will be opposite. Multiply by -1 to change them back */
//	    float mouse_x;      /**< X coordinate, relative to window */
//	    float mouse_y;      /**< Y coordinate, relative to window */
//	    Sint32 integer_x;   /**< The amount scrolled horizontally, accumulated to whole scroll "ticks" (added in 3.2.12) */
//	    Sint32 integer_y;   /**< The amount scrolled vertically, accumulated to whole scroll "ticks" (added in 3.2.12) */
//	} SDL_MouseWheelEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type MouseWheelEvent struct {
	// SDL_EVENT_MOUSE_WHEEL
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

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

	// Set to one of the SDL_MOUSEWHEEL_* defines. When FLIPPED the values in X
	// and Y will be opposite. Multiply by -1 to change them back
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

// # SDL_NotificationEvent
//
// Notification dialog event structure (event.notification.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_NotificationEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_NOTIFICATION_ACTION_INVOKED */
//	    Uint32 reserved;
//	    Uint64 timestamp;         /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_NotificationID which; /**< The ID of the notification that generated this event. */
//	    const char *action_id;    /**< The identifier string of the action invoked in the notification dialog. */
//	} SDL_NotificationEvent;
//
// # Remarks
//
// An `action_id` value of 'default' for an
// [SDL_EVENT_NOTIFICATION_ACTION_INVOKED](SDL_EVENT_NOTIFICATION_ACTION_INVOKED)
// event indicates that the notification was interacted with without selecting
// a specific action (e.g. the body of the notification was clicked on).
//
// # Version
//
// This struct is available since SDL 3.6.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type NotificationEvent struct {
	// SDL_EVENT_NOTIFICATION_ACTION_INVOKED
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The ID of the notification that generated this event.
	Which NotificationID

	// The identifier string of the action invoked in the notification dialog.
	ActionId string
}

// # SDL_PenAxisEvent
//
// Pressure-sensitive pen pressure / angle event structure (event.paxis.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_PenAxisEvent
//	{
//	    SDL_EventType type;     /**< SDL_EVENT_PEN_AXIS */
//	    Uint32 reserved;
//	    Uint64 timestamp;       /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID;  /**< The window with pen focus, if any */
//	    SDL_PenID which;        /**< The pen instance id */
//	    SDL_PenInputFlags pen_state;   /**< Complete pen input state at time of event */
//	    float x;                /**< X coordinate, relative to window */
//	    float y;                /**< Y coordinate, relative to window */
//	    SDL_PenAxis axis;       /**< Axis that has changed */
//	    float value;            /**< New value of axis */
//	} SDL_PenAxisEvent;
//
// # Remarks
//
// You might get some of these events even if the pen isn't touching the
// tablet.
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenAxisEvent struct {
	// SDL_EVENT_PEN_AXIS
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

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

// # SDL_PenButtonEvent
//
// Pressure-sensitive pen button event structure (event.pbutton.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_PenButtonEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_PEN_BUTTON_DOWN or SDL_EVENT_PEN_BUTTON_UP */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The window with mouse focus, if any */
//	    SDL_PenID which;        /**< The pen instance id */
//	    SDL_PenInputFlags pen_state;   /**< Complete pen input state at time of event */
//	    float x;                /**< X coordinate, relative to window */
//	    float y;                /**< Y coordinate, relative to window */
//	    Uint8 button;       /**< The pen button index (first button is 1). */
//	    bool down;      /**< true if the button is pressed */
//	} SDL_PenButtonEvent;
//
// # Remarks
//
// This is for buttons on the pen itself that the user might click. The pen
// itself pressing down to draw triggers a
// [SDL_EVENT_PEN_DOWN](SDL_EVENT_PEN_DOWN) event instead.
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenButtonEvent struct {
	// SDL_EVENT_PEN_BUTTON_DOWN or SDL_EVENT_PEN_BUTTON_UP
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

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

// # SDL_PenMotionEvent
//
// Pressure-sensitive pen motion event structure (event.pmotion.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_PenMotionEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_PEN_MOTION */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The window with pen focus, if any */
//	    SDL_PenID which;        /**< The pen instance id */
//	    SDL_PenInputFlags pen_state;   /**< Complete pen input state at time of event */
//	    float x;                /**< X coordinate, relative to window */
//	    float y;                /**< Y coordinate, relative to window */
//	} SDL_PenMotionEvent;
//
// # Remarks
//
// Depending on the hardware, you may get motion events when the pen is not
// touching a tablet, for tracking a pen even when it isn't drawing. You
// should listen for [SDL_EVENT_PEN_DOWN](SDL_EVENT_PEN_DOWN) and
// [SDL_EVENT_PEN_UP](SDL_EVENT_PEN_UP) events, or check `pen_state &
// SDL_PEN_INPUT_DOWN` to decide if a pen is "drawing" when dealing with pen
// motion.
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenMotionEvent struct {
	// SDL_EVENT_PEN_MOTION
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

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

// # SDL_PenProximityEvent
//
// Pressure-sensitive pen proximity event structure (event.pproximity.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_PenProximityEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_PEN_PROXIMITY_IN or SDL_EVENT_PEN_PROXIMITY_OUT */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The window with pen focus, if any */
//	    SDL_PenID which;        /**< The pen instance id */
//	} SDL_PenProximityEvent;
//
// # Remarks
//
// When a pen becomes visible to the system (it is close enough to a tablet,
// etc), SDL will send an
// [SDL_EVENT_PEN_PROXIMITY_IN](SDL_EVENT_PEN_PROXIMITY_IN) event with the new
// pen's ID. This ID is valid until the pen leaves proximity again (has been
// removed from the tablet's area, the tablet has been unplugged, etc). If the
// same pen reenters proximity again, it will be given a new ID.
//
// Note that "proximity" means "close enough for the tablet to know the tool
// is there." The pen touching and lifting off from the tablet while not
// leaving the area are handled by [SDL_EVENT_PEN_DOWN](SDL_EVENT_PEN_DOWN)
// and [SDL_EVENT_PEN_UP](SDL_EVENT_PEN_UP).
//
// Not all platforms have a window associated with the pen during proximity
// events. Some wait until motion/button/etc events to offer this info.
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenProximityEvent struct {
	// SDL_EVENT_PEN_PROXIMITY_IN or SDL_EVENT_PEN_PROXIMITY_OUT
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The window with pen focus, if any
	WindowID WindowID

	// The pen instance id
	Which PenID
}

// # SDL_PenTouchEvent
//
// Pressure-sensitive pen touched event structure (event.ptouch.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_PenTouchEvent
//	{
//	    SDL_EventType type;     /**< SDL_EVENT_PEN_DOWN or SDL_EVENT_PEN_UP */
//	    Uint32 reserved;
//	    Uint64 timestamp;       /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID;  /**< The window with pen focus, if any */
//	    SDL_PenID which;        /**< The pen instance id */
//	    SDL_PenInputFlags pen_state;   /**< Complete pen input state at time of event */
//	    float x;                /**< X coordinate, relative to window */
//	    float y;                /**< Y coordinate, relative to window */
//	    bool eraser;        /**< true if eraser end is used (not all pens support this). */
//	    bool down;          /**< true if the pen is touching or false if the pen is lifted off */
//	} SDL_PenTouchEvent;
//
// # Remarks
//
// These events come when a pen touches a surface (a tablet, etc), or lifts
// off from one.
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PenTouchEvent struct {
	// SDL_EVENT_PEN_DOWN or SDL_EVENT_PEN_UP
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

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

// # SDL_PinchFingerEvent
//
// Pinch event structure (event.pinch.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_PinchFingerEvent
//	{
//	    SDL_EventType type; /**< ::SDL_EVENT_PINCH_BEGIN or ::SDL_EVENT_PINCH_UPDATE or ::SDL_EVENT_PINCH_END */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    float scale;        /**< The scale change since the last SDL_EVENT_PINCH_UPDATE. Scale < 1 is "zoom out". Scale > 1 is "zoom in". */
//	    SDL_WindowID windowID; /**< The window underneath the finger, if any */
//	} SDL_PinchFingerEvent;
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type PinchFingerEvent struct {
	// SDL_EVENT_PINCH_BEGIN or SDL_EVENT_PINCH_UPDATE or SDL_EVENT_PINCH_END
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The scale change since the last SDL_EVENT_PINCH_UPDATE. Scale < 1 is
	// ``zoom out''. Scale > 1 is ``zoom in''.
	Scale float32

	// The window underneath the finger, if any
	WindowID WindowID
}

// # SDL_QuitEvent
//
// The "quit requested" event
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_QuitEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_QUIT */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	} SDL_QuitEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type QuitEvent struct {
	// SDL_EVENT_QUIT
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64
}

// # SDL_RenderEvent
//
// Renderer event structure (event.render.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_RenderEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_RENDER_TARGETS_RESET, SDL_EVENT_RENDER_DEVICE_RESET, SDL_EVENT_RENDER_DEVICE_LOST */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The window containing the renderer in question. */
//	} SDL_RenderEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type RenderEvent struct {
	// SDL_EVENT_RENDER_TARGETS_RESET, SDL_EVENT_RENDER_DEVICE_RESET,
	// SDL_EVENT_RENDER_DEVICE_LOST
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The window containing the renderer in question.
	WindowID WindowID
}

// # SDL_SensorEvent
//
// Sensor event structure (event.sensor.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_SensorEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_SENSOR_UPDATE */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_SensorID which; /**< The instance ID of the sensor */
//	    float data[6];      /**< Up to 6 values from the sensor - additional values can be queried using SDL_GetSensorData() */
//	    Uint64 sensor_timestamp; /**< The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock */
//	} SDL_SensorEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type SensorEvent struct {
	// SDL_EVENT_SENSOR_UPDATE
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The instance ID of the sensor
	Which SensorID

	// Up to 6 values from the sensor - additional values can be queried using
	// SDL_GetSensorData()
	Data [6]float32

	// The timestamp of the sensor reading in nanoseconds, not necessarily
	// synchronized with the system clock
	SensorTimestamp uint64
}

// # SDL_TextEditingCandidatesEvent
//
// Keyboard IME candidates event structure (event.edit_candidates.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_TextEditingCandidatesEvent
//	{
//	    SDL_EventType type;         /**< SDL_EVENT_TEXT_EDITING_CANDIDATES */
//	    Uint32 reserved;
//	    Uint64 timestamp;           /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID;      /**< The window with keyboard focus, if any */
//	    const char * const *candidates;    /**< The list of candidates, or NULL if there are no candidates available */
//	    Sint32 num_candidates;      /**< The number of strings in `candidates` */
//	    Sint32 selected_candidate;  /**< The index of the selected candidate, or -1 if no candidate is selected */
//	    bool horizontal;          /**< true if the list is horizontal, false if it's vertical */
//	    Uint8 padding1;
//	    Uint8 padding2;
//	    Uint8 padding3;
//	} SDL_TextEditingCandidatesEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type TextEditingCandidatesEvent struct {
	// SDL_EVENT_TEXT_EDITING_CANDIDATES
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The window with keyboard focus, if any
	WindowID WindowID

	// The list of candidates, or NULL if there are no candidates available
	Candidates []string

	// The index of the selected candidate, or -1 if no candidate is selected
	SelectedCandidate int32

	// true if the list is horizontal, false if it's vertical
	Horizontal bool
}

// # SDL_TextEditingEvent
//
// Keyboard text editing event structure (event.edit.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_TextEditingEvent
//	{
//	    SDL_EventType type;         /**< SDL_EVENT_TEXT_EDITING */
//	    Uint32 reserved;
//	    Uint64 timestamp;           /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID;      /**< The window with keyboard focus, if any */
//	    const char *text;           /**< The editing text */
//	    Sint32 start;               /**< The start cursor of selected editing text, or -1 if not set */
//	    Sint32 length;              /**< The length of selected editing text, or -1 if not set */
//	} SDL_TextEditingEvent;
//
// # Remarks
//
// The start cursor is the position, in UTF-8 characters, where new typing
// will be inserted into the editing text. The length is the number of UTF-8
// characters that will be replaced by new typing.
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type TextEditingEvent struct {
	// SDL_EVENT_TEXT_EDITING
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The window with keyboard focus, if any
	WindowID WindowID

	// The editing text
	Text string

	// The start cursor of selected editing text, or -1 if not set
	Start int32

	// The length of selected editing text, or -1 if not set
	Length int32
}

// # SDL_TextInputEvent
//
// Keyboard text input event structure (event.text.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_TextInputEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_TEXT_INPUT */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The window with keyboard focus, if any */
//	    const char *text;   /**< The input text, UTF-8 encoded */
//	} SDL_TextInputEvent;
//
// # Remarks
//
// This event will never be delivered unless text input is enabled by calling
// [SDL_StartTextInput](SDL_StartTextInput)(). Text input is disabled by
// default!
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// # See Also
//
//   - [StartTextInput]
//   - [StopTextInput]
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type TextInputEvent struct {
	// SDL_EVENT_TEXT_INPUT
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The window with keyboard focus, if any
	WindowID WindowID

	// The input text, UTF-8 encoded
	Text string
}

// # SDL_TouchFingerEvent
//
// Touch finger event structure (event.tfinger.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>](https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h)
//
// # Syntax
//
//	typedef struct SDL_TouchFingerEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_FINGER_DOWN, SDL_EVENT_FINGER_UP, SDL_EVENT_FINGER_MOTION, or SDL_EVENT_FINGER_CANCELED */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_TouchID touchID; /**< The touch device id */
//	    SDL_FingerID fingerID;
//	    float x;            /**< Normalized in the range 0...1 */
//	    float y;            /**< Normalized in the range 0...1 */
//	    float dx;           /**< Normalized in the range -1...1 */
//	    float dy;           /**< Normalized in the range -1...1 */
//	    float pressure;     /**< Normalized in the range 0...1 */
//	    SDL_WindowID windowID; /**< The window underneath the finger, if any */
//	} SDL_TouchFingerEvent;
//
// # Remarks
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
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type TouchFingerEvent struct {
	// SDL_EVENT_FINGER_DOWN, SDL_EVENT_FINGER_UP, SDL_EVENT_FINGER_MOTION, or
	// SDL_EVENT_FINGER_CANCELED
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

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

// # SDL_UserEvent
//
// A user-defined event type (event.user.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_UserEvent
//	{
//	    Uint32 type;        /**< SDL_EVENT_USER through SDL_EVENT_LAST, Uint32 because these are not in the SDL_EventType enumeration */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The associated window if any */
//	    Sint32 code;        /**< User defined event code */
//	    void *data1;        /**< User defined data pointer */
//	    void *data2;        /**< User defined data pointer */
//	} SDL_UserEvent;
//
// # Remarks
//
// This event is unique; it is never created by SDL, but only by the
// application. The event can be pushed onto the event queue using
// [SDL_PushEvent](SDL_PushEvent)(). The contents of the structure members are
// completely up to the programmer; the only requirement is that “type” is
// a value obtained from [SDL_RegisterEvents](SDL_RegisterEvents)().
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// # Code Examples
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
	// SDL_EVENT_USER through SDL_EVENT_LAST, Uint32 because these are not in
	// the SDL_EventType enumeration
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The associated window if any
	WindowID WindowID

	// User defined event code
	Code int32

	// User defined data pointer
	Data1 *any

	// User defined data pointer
	Data2 *any
}

// # SDL_WindowEvent
//
// Window state change event data (event.window.*)
//
// # Header File
//
// Defined in [<SDL3/SDL_events.h>]
//
// # Syntax
//
//	typedef struct SDL_WindowEvent
//	{
//	    SDL_EventType type; /**< SDL_EVENT_WINDOW_* */
//	    Uint32 reserved;
//	    Uint64 timestamp;   /**< In nanoseconds, populated using SDL_GetTicksNS() */
//	    SDL_WindowID windowID; /**< The associated window */
//	    Sint32 data1;       /**< event dependent data */
//	    Sint32 data2;       /**< event dependent data */
//	} SDL_WindowEvent;
//
// # Version
//
// This struct is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
type WindowEvent struct {
	// SDL_EVENT_WINDOW_*
	Type EventType

	// In nanoseconds, populated using SDL_GetTicksNS()
	Timestamp uint64

	// The associated window
	WindowID WindowID

	// event dependent data
	Data1 int32

	// event dependent data
	Data2 int32
}

func convertEvent(cEvent *C.SDL_Event) Event {
	cCommonEvent := (*C.SDL_CommonEvent)(unsafe.Pointer(cEvent))
	t := EventType(cCommonEvent._type)

	switch {
	case t >= EVENT_DISPLAY_FIRST && t <= EVENT_DISPLAY_LAST:
		e := (*C.SDL_DisplayEvent)(unsafe.Pointer(cEvent))
		return DisplayEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			DisplayID: DisplayID(e.displayID),
			Data1:     int32(e.data1),
			Data2:     int32(e.data2),
		}
	case t >= EVENT_WINDOW_FIRST && t <= EVENT_WINDOW_LAST:
		e := (*C.SDL_WindowEvent)(unsafe.Pointer(cEvent))
		return WindowEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Data1:     int32(e.data1),
			Data2:     int32(e.data2),
		}
	case t == EVENT_KEYBOARD_ADDED || t == EVENT_KEYBOARD_REMOVED:
		e := (*C.SDL_KeyboardDeviceEvent)(unsafe.Pointer(cEvent))
		return KeyboardDeviceEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     KeyboardID(e.which),
		}
	case t == EVENT_KEY_DOWN || t == EVENT_KEY_UP:
		e := (*C.SDL_KeyboardEvent)(unsafe.Pointer(cEvent))
		return KeyboardEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Which:     KeyboardID(e.which),
			Scancode:  Scancode(e.scancode),
			Key:       Keycode(e.key),
			Mod:       Keymod(e.mod),
			Raw:       uint16(e.raw),
			Down:      bool(e.down),
			Repeat:    bool(e.repeat),
		}
	case t == EVENT_TEXT_EDITING:
		e := (*C.SDL_TextEditingEvent)(unsafe.Pointer(cEvent))
		return TextEditingEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Text:      C.GoString(e.text),
			Start:     int32(e.start),
			Length:    int32(e.length),
		}
	case t == EVENT_TEXT_EDITING_CANDIDATES:
		e := (*C.SDL_TextEditingCandidatesEvent)(unsafe.Pointer(cEvent))
		return TextEditingCandidatesEvent{
			Type:              EventType(e._type),
			Timestamp:         uint64(e.timestamp),
			WindowID:          WindowID(e.windowID),
			Candidates:        toStrings(e.candidates, int32(e.num_candidates)),
			SelectedCandidate: int32(e.selected_candidate),
			Horizontal:        bool(e.horizontal),
		}
	case t == EVENT_TEXT_INPUT:
		e := (*C.SDL_TextInputEvent)(unsafe.Pointer(cEvent))
		return TextInputEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Text:      C.GoString(e.text),
		}
	case t == EVENT_MOUSE_ADDED || t == EVENT_MOUSE_REMOVED:
		e := (*C.SDL_MouseDeviceEvent)(unsafe.Pointer(cEvent))
		return MouseDeviceEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     MouseID(e.which),
		}
	case t == EVENT_MOUSE_MOTION:
		e := (*C.SDL_MouseMotionEvent)(unsafe.Pointer(cEvent))
		return MouseMotionEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Which:     MouseID(e.which),
			State:     MouseButtonFlags(e.state),
			X:         float32(e.x),
			Y:         float32(e.y),
			Xrel:      float32(e.xrel),
			Yrel:      float32(e.yrel),
		}
	case t == EVENT_MOUSE_BUTTON_DOWN || t == EVENT_MOUSE_BUTTON_UP:
		e := (*C.SDL_MouseButtonEvent)(unsafe.Pointer(cEvent))
		return MouseButtonEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Which:     MouseID(e.which),
			Button:    uint8(e.button),
			Down:      bool(e.down),
			Clicks:    uint8(e.clicks),
			X:         float32(e.x),
			Y:         float32(e.y),
		}
	case t == EVENT_MOUSE_WHEEL:
		e := (*C.SDL_MouseWheelEvent)(unsafe.Pointer(cEvent))
		return MouseWheelEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
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
		return JoyDeviceEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
		}
	case t == EVENT_JOYSTICK_AXIS_MOTION:
		e := (*C.SDL_JoyAxisEvent)(unsafe.Pointer(cEvent))
		return JoyAxisEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
			Axis:      uint8(e.axis),
			Value:     int16(e.value),
		}
	case t == EVENT_JOYSTICK_BALL_MOTION:
		e := (*C.SDL_JoyBallEvent)(unsafe.Pointer(cEvent))
		return JoyBallEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
			Ball:      uint8(e.ball),
			Xrel:      int16(e.xrel),
			Yrel:      int16(e.yrel),
		}
	case t == EVENT_JOYSTICK_HAT_MOTION:
		e := (*C.SDL_JoyHatEvent)(unsafe.Pointer(cEvent))
		return JoyHatEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
			Hat:       uint8(e.hat),
			Value:     JoystickHat(e.value),
		}
	case t == EVENT_JOYSTICK_BUTTON_DOWN || t == EVENT_JOYSTICK_BUTTON_UP:
		e := (*C.SDL_JoyButtonEvent)(unsafe.Pointer(cEvent))
		return JoyButtonEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
			Button:    uint8(e.button),
			Down:      bool(e.down),
		}
	case t == EVENT_JOYSTICK_BATTERY_UPDATED:
		e := (*C.SDL_JoyBatteryEvent)(unsafe.Pointer(cEvent))
		return JoyBatteryEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
			State:     PowerState(e.state),
			Percent:   int(e.percent),
		}
	case t == EVENT_GAMEPAD_ADDED || t == EVENT_GAMEPAD_REMOVED ||
		t == EVENT_GAMEPAD_REMAPPED || t == EVENT_GAMEPAD_UPDATE_COMPLETE ||
		t == EVENT_GAMEPAD_STEAM_HANDLE_UPDATED:

		e := (*C.SDL_GamepadDeviceEvent)(unsafe.Pointer(cEvent))
		return GamepadDeviceEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
		}
	case t == EVENT_GAMEPAD_AXIS_MOTION:
		e := (*C.SDL_GamepadAxisEvent)(unsafe.Pointer(cEvent))
		return GamepadAxisEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
			Axis:      uint8(e.axis),
			Value:     int16(e.value),
		}
	case t == EVENT_GAMEPAD_BUTTON_DOWN || t == EVENT_GAMEPAD_BUTTON_UP:
		e := (*C.SDL_GamepadButtonEvent)(unsafe.Pointer(cEvent))
		return GamepadButtonEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
			Button:    uint8(e.button),
			Down:      bool(e.down),
		}
	case t == EVENT_GAMEPAD_TOUCHPAD_DOWN ||
		t == EVENT_GAMEPAD_TOUCHPAD_MOTION || t == EVENT_GAMEPAD_TOUCHPAD_UP:

		e := (*C.SDL_GamepadTouchpadEvent)(unsafe.Pointer(cEvent))
		return GamepadTouchpadEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     JoystickID(e.which),
			Touchpad:  int32(e.touchpad),
			Finger:    int32(e.finger),
			X:         float32(e.x),
			Y:         float32(e.y),
			Pressure:  float32(e.pressure),
		}
	case t == EVENT_GAMEPAD_SENSOR_UPDATE:
		e := (*C.SDL_GamepadSensorEvent)(unsafe.Pointer(cEvent))
		return GamepadSensorEvent{
			Type:            EventType(e._type),
			Timestamp:       uint64(e.timestamp),
			Which:           JoystickID(e.which),
			Sensor:          int32(e.sensor),
			Data:            *(*[3]float32)(unsafe.Pointer(&e.data)),
			SensorTimestamp: uint64(e.sensor_timestamp),
		}
	//case t == EVENT_GAMEPAD_CAPSENSE_TOUCH ||
	//	t == EVENT_GAMEPAD_CAPSENSE_RELEASE:

	//	e := (*C.SDL_GamepadCapSenseEvent)(unsafe.Pointer(cEvent))
	//	return GamepadCapSenseEvent{
	//		Type:      EventType(e._type),
	//		Timestamp: uint64(e.timestamp),
	//		Which:     JoystickID(e.which),
	//		Capsense:  uint8(e.capsense),
	//		Down:      bool(e.down),
	//	}
	case t == EVENT_AUDIO_DEVICE_ADDED || t == EVENT_AUDIO_DEVICE_REMOVED ||
		t == EVENT_AUDIO_DEVICE_FORMAT_CHANGED:

		e := (*C.SDL_AudioDeviceEvent)(unsafe.Pointer(cEvent))
		return AudioDeviceEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     AudioDeviceID(e.which),
			recording: bool(e.recording),
		}
	case t == EVENT_CAMERA_DEVICE_ADDED || t == EVENT_CAMERA_DEVICE_REMOVED ||
		t == EVENT_CAMERA_DEVICE_APPROVED || t == EVENT_CAMERA_DEVICE_DENIED:

		e := (*C.SDL_CameraDeviceEvent)(unsafe.Pointer(cEvent))
		return CameraDeviceEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Which:     CameraID(e.which),
		}
	case t == EVENT_SENSOR_UPDATE:
		e := (*C.SDL_SensorEvent)(unsafe.Pointer(cEvent))
		return SensorEvent{
			Type:            EventType(e._type),
			Timestamp:       uint64(e.timestamp),
			Which:           SensorID(e.which),
			Data:            *(*[6]float32)(unsafe.Pointer(&e.data)),
			SensorTimestamp: uint64(e.sensor_timestamp),
		}
	case t == EVENT_QUIT:
		e := (*C.SDL_QuitEvent)(unsafe.Pointer(cEvent))
		return QuitEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
		}
	case t >= EVENT_USER && t <= EVENT_LAST:
		e := (*C.SDL_UserEvent)(unsafe.Pointer(cEvent))
		return UserEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Code:      int32(e.code),
			Data1:     (*any)(e.data1),
			Data2:     (*any)(e.data2),
		}
	case t == EVENT_FINGER_DOWN || t == EVENT_FINGER_UP ||
		t == EVENT_FINGER_MOTION || t == EVENT_FINGER_CANCELED:

		e := (*C.SDL_TouchFingerEvent)(unsafe.Pointer(cEvent))
		return TouchFingerEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			TouchID:   TouchID(e.touchID),
			FingerID:  FingerID(e.fingerID),
			X:         float32(e.x),
			Y:         float32(e.y),
			Dx:        float32(e.dx),
			Dy:        float32(e.dy),
			Pressure:  float32(e.pressure),
			WindowID:  WindowID(e.windowID),
		}
	case t == EVENT_PINCH_BEGIN || t == EVENT_PINCH_UPDATE ||
		t == EVENT_PINCH_END:

		e := (*C.SDL_PinchFingerEvent)(unsafe.Pointer(cEvent))
		return PinchFingerEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Scale:     float32(e.scale),
			WindowID:  WindowID(e.windowID),
		}
	case t == EVENT_PEN_PROXIMITY_IN || t == EVENT_PEN_PROXIMITY_OUT:
		e := (*C.SDL_PenProximityEvent)(unsafe.Pointer(cEvent))
		return PenProximityEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Which:     PenID(e.which),
		}
	case t == EVENT_PEN_DOWN || t == EVENT_PEN_UP:
		e := (*C.SDL_PenTouchEvent)(unsafe.Pointer(cEvent))
		return PenTouchEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Which:     PenID(e.which),
			PenState:  PenInputFlags(e.pen_state),
			X:         float32(e.x),
			Y:         float32(e.y),
			Eraser:    bool(e.eraser),
			Down:      bool(e.down),
		}
	case t == EVENT_PEN_MOTION:
		e := (*C.SDL_PenMotionEvent)(unsafe.Pointer(cEvent))
		return PenMotionEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Which:     PenID(e.which),
			PenState:  PenInputFlags(e.pen_state),
			X:         float32(e.x),
			Y:         float32(e.y),
		}
	case t == EVENT_PEN_BUTTON_DOWN || t == EVENT_PEN_BUTTON_UP:
		e := (*C.SDL_PenButtonEvent)(unsafe.Pointer(cEvent))
		return PenButtonEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Which:     PenID(e.which),
			PenState:  PenInputFlags(e.pen_state),
			X:         float32(e.x),
			Y:         float32(e.y),
			Button:    uint8(e.button),
			Down:      bool(e.down),
		}
	case t == EVENT_PEN_AXIS:
		e := (*C.SDL_PenAxisEvent)(unsafe.Pointer(cEvent))
		return PenAxisEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			Which:     PenID(e.which),
			PenState:  PenInputFlags(e.pen_state),
			X:         float32(e.x),
			Y:         float32(e.y),
			Axis:      PenAxis(e.axis),
			Value:     float32(e.value),
		}
	case t == EVENT_RENDER_TARGETS_RESET || t == EVENT_RENDER_DEVICE_RESET ||
		t == EVENT_RENDER_DEVICE_LOST:

		e := (*C.SDL_RenderEvent)(unsafe.Pointer(cEvent))
		return RenderEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
		}
	case t == EVENT_DROP_BEGIN || t == EVENT_DROP_FILE ||
		t == EVENT_DROP_TEXT || t == EVENT_DROP_COMPLETE || t == EVENT_DROP_POSITION:

		e := (*C.SDL_DropEvent)(unsafe.Pointer(cEvent))
		return DropEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			WindowID:  WindowID(e.windowID),
			X:         float32(e.x),
			Y:         float32(e.y),
			Source:    C.GoString(e.source),
			Data:      C.GoString(e.data),
		}
	case t == EVENT_CLIPBOARD_UPDATE:
		e := (*C.SDL_ClipboardEvent)(unsafe.Pointer(cEvent))
		return ClipboardEvent{
			Type:      EventType(e._type),
			Timestamp: uint64(e.timestamp),
			Owner:     bool(e.owner),
			MIMETypes: toStrings(e.mime_types, int32(e.num_mime_types)),
		}
	//case t == EVENT_NOTIFICATION_ACTION_INVOKED:
	//	e := (*C.SDL_NotificationEvent)(unsafe.Pointer(cEvent))
	//	return NotificationEvent{
	//		Type:      EventType(e._type),
	//		Timestamp: uint64(e.timestamp),
	//		Which:     NotificationID(e.which),
	//		ActionId:  C.GoString(e.action_id),
	//	}
	}

	panic(fmt.Sprintf("unknown event type %d", t))
}
