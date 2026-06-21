// Package sdl is a Go wrapper for the [SDL library].
//
// [SDL library]: https://www.libsdl.org
package sdl

// #include <SDL3/SDL.h>
// #include <SDL3_image/SDL_image.h>
//
// #include <stdlib.h>
import "C"

import "unsafe"

// Error is an error type indicating an SDL error. It stores the error message
// from SDL_GetError.
type Error struct {
	message string
}

func (e Error) Error() string {
	return e.message
}

func getError() Error {
	return Error{message: C.GoString(C.SDL_GetError())}
}

// Init initializes the SDL library.
//
// Defined in [<SDL3/SDL_init.h>].
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// SDL_Init() simply forwards to calling SDL_InitSubSystem(). Therefore, the
// two may be used interchangeably. Though for readability of your code
// SDL_InitSubSystem() might be preferred.
//
// The file I/O (for example: SDL_IOFromFile) and threading (SDL_CreateThread)
// subsystems are initialized by default. Message boxes
// (SDL_ShowSimpleMessageBox) also attempt to work without initializing the
// video subsystem, in hopes of being useful in showing an error dialog when
// SDL_Init fails. You must specifically initialize other subsystems if you use
// them in your application.
//
// Logging (such as SDL_Log) works without initialization, too.
//
// flags may be any of the following OR'd together:
//
//   - SDL_INIT_AUDIO: audio subsystem; automatically initializes the events
//     subsystem
//   - SDL_INIT_VIDEO: video subsystem; automatically initializes the events
//     subsystem, should be initialized on the main thread.
//   - SDL_INIT_JOYSTICK: joystick subsystem; automatically initializes the
//     events subsystem
//   - SDL_INIT_HAPTIC: haptic (force feedback) subsystem
//   - SDL_INIT_GAMEPAD: gamepad subsystem; automatically initializes the
//     joystick subsystem
//   - SDL_INIT_EVENTS: events subsystem
//   - SDL_INIT_SENSOR: sensor subsystem; automatically initializes the events
//     subsystem
//   - SDL_INIT_CAMERA: camera subsystem; automatically initializes the events
//     subsystem
//
// Subsystem initialization is ref-counted, you must call SDL_QuitSubSystem()
// for each SDL_InitSubSystem() to correctly shutdown a subsystem manually (or
// call SDL_Quit() to force shutdown). If a subsystem is already loaded then
// this call will increase the ref-count and return.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// [<SDL3/SDL_init.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_init.h
func Init(flags InitFlags) error {
	if !C.SDL_Init(C.SDL_InitFlags(flags)) {
		return getError()
	}
	return nil
}

// Quit cleans up all initialized subsystems.
//
// Defined in [<SDL3/SDL_init.h>].
//
// You should call this function even if you have already shutdown each
// initialized subsystem with [QuitSubSystem]. It is safe to call this function
// even in the case of errors in initialization.
//
// You can use this function with atexit() to ensure that it is run when your
// application is shutdown, but it is not wise to do this from a library or
// other dynamically loaded code.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// [<SDL3/SDL_init.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_init.h
func Quit() {
	C.SDL_Quit()
}

// Window is an opaque handle to a window.
type Window struct{ ptr *C.SDL_Window }

// Renderer represents rendering state.
type Renderer struct{ ptr *C.SDL_Renderer }

// CreateWindow creates a window with the specified dimensions and flags.
//
// Defined in [<SDL3/SDL_video.h>].
//
// Returns the window that was created or NULL on failure; call SDL_GetError()
// for more information.
//
// The window size is a request and may be different than expected based on the
// desktop layout and window manager policies. Your application should be
// prepared to handle a window of any size.
//
// flags may be any of the following OR'd together:
//
//   - [WindowFullscreen]: fullscreen window at desktop resolution
//   - [WindowOpengl]: window usable with an OpenGL context
//   - [WindowHidden]: window is not visible
//   - [WindowBorderless]: no window decoration
//   - [WindowResizable]: window can be resized
//   - [WindowMinimized]: window is minimized
//   - [WindowMaximized]: window is maximized
//   - [WindowMouseGrabbed]: window has grabbed mouse focus
//   - [WindowInputFocus]: window has input focus
//   - [WindowMouseFocus]: window has mouse focus
//   - [WindowExternal]: window not created by SDL
//   - [WindowModal]: window is modal
//   - [WindowHighPixelDensity]: window uses high pixel density back buffer if
//     possible
//   - [WindowMouseCapture]: window has mouse captured (unrelated to
//     MOUSE_GRABBED)
//   - [WindowAlwaysOnTop]: window should always be above others
//   - [WindowUtility]: window should be treated as a utility window, not
//     showing in the task bar and window list
//   - [WindowTooltip]: window should be treated as a tooltip and does not get
//     mouse or keyboard focus, requires a parent window
//   - [WindowPopupMenu]: window should be treated as a popup menu, requires a
//     parent window
//   - [WindowKeyboardGrabbed]: window has grabbed keyboard input
//   - [WindowVulkan]: window usable with a Vulkan instance
//   - [WindowMetal]: window usable with a Metal instance
//   - [WindowTransparent]: window with transparent buffer
//   - [WindowNotFocusable]: window should not be focusable
//
// The SDL_Window will be shown if SDL_WINDOW_HIDDEN is not set. If hidden at
// creation time, SDL_ShowWindow() can be used to show it later.
//
// On Apple's macOS, you must set the NSHighResolutionCapable Info.plist
// property to YES, otherwise you will not receive a High-DPI OpenGL canvas.
//
// The window pixel size may differ from its window coordinate size if the
// window is on a high pixel density display. Use SDL_GetWindowSize() to query
// the client area's size in window coordinates, and
// SDL_GetWindowSizeInPixels() or SDL_GetRenderOutputSize() to query the
// drawable size in pixels. Note that the drawable size can vary after the
// window is created and should be queried again if you get an
// SDL_EVENT_WINDOW_PIXEL_SIZE_CHANGED event.
//
// If the window is created with any of the SDL_WINDOW_OPENGL or
// SDL_WINDOW_VULKAN flags, then the corresponding LoadLibrary function
// (SDL_GL_LoadLibrary or SDL_Vulkan_LoadLibrary) is called and the
// corresponding UnloadLibrary function is called by SDL_DestroyWindow().
//
// If SDL_WINDOW_VULKAN is specified and there isn't a working Vulkan driver,
// SDL_CreateWindow() will fail, because SDL_Vulkan_LoadLibrary() will fail.
//
// If SDL_WINDOW_METAL is specified on an OS that does not support Metal,
// SDL_CreateWindow() will fail.
//
// If you intend to use this window with an SDL_Renderer, you should use
// SDL_CreateWindowAndRenderer() instead of this function, to avoid window
// flicker.
//
// On non-Apple devices, SDL requires you to either not link to the Vulkan
// loader or link to a dynamic library version. This limitation may be removed
// in a future version of SDL.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// [<SDL3/SDL_video.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_video.h
func CreateWindow(title string, w, h int, flags WindowFlags) (*Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cWindow := C.SDL_CreateWindow(
		cTitle, C.int(w), C.int(h), C.SDL_WindowFlags(flags))
	if cWindow == nil {
		return nil, getError()
	}
	return &Window{ptr: cWindow}, nil
}

// Destroy destroys a window.
//
// Defined in [<SDL3/SDL_video.h>].
//
// Any child windows owned by the window will be recursively destroyed as well.
//
// Note that on some platforms, the visible window may not actually be removed
// from the screen until the SDL event loop is pumped again, even though the
// SDL_Window is no longer valid after this call.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// [<SDL3/SDL_video.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_video.h
func (w *Window) Destroy() {
	if w.ptr != nil {
		C.SDL_DestroyWindow(w.ptr)
		w.ptr = nil
	}
}

// CreateRenderer creates a 2D rendering context for a window.
//
// Defined in [<SDL3/SDL_render.h>].
//
// Returns a valid rendering context or NULL if there was an error; call
// SDL_GetError() for more information.
//
// If you want a specific renderer, you can specify its name here. A list of
// available renderers can be obtained by calling SDL_GetRenderDriver()
// multiple times, with indices from 0 to SDL_GetNumRenderDrivers()-1. If you
// don't need a specific renderer, specify NULL and SDL will attempt to choose
// the best option for you, based on what is available on the user's system.
//
// If name is a comma-separated list, SDL will try each name, in the order
// listed, until one succeeds or all of them fail.
//
// By default the rendering size matches the window size in pixels, but you can
// call SDL_SetRenderLogicalPresentation() to change the content size and
// scaling options.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// [<SDL3/SDL_render.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_render.h
func (w *Window) CreateRenderer() (*Renderer, error) {
	cRenderer := C.SDL_CreateRenderer(w.ptr, nil)
	if cRenderer == nil {
		return nil, getError()
	}
	return &Renderer{ptr: cRenderer}, nil
}

// Destroy destroys the rendering context for a window and free all associated
// textures.
//
// Defined in [<SDL3/SDL_render.h>].
//
// This should be called before destroying the associated window.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// See Also:
//   - SDL_CreateRenderer
//
// [<SDL3/SDL_render.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_render.h
func (rr *Renderer) Destroy() {
	if rr.ptr != nil {
		C.SDL_DestroyRenderer(rr.ptr)
		rr.ptr = nil
	}
}

// Clear clears the current rendering target with the drawing color.
//
// Defined in [<SDL3/SDL_render.h>].
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function clears the entire rendering target, ignoring the viewport and
// the clip rectangle. Note, that clearing will also set/fill all pixels of the
// rendering target to current renderer draw color, so make sure to invoke
// [Renderer.SetDrawColor] when needed.
//
// [<SDL3/SDL_render.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_render.h
func (rr *Renderer) Clear() error {
	if !C.SDL_RenderClear(rr.ptr) {
		return getError()
	}
	return nil
}

// Present updates the screen with any rendering performed since the previous
// call.
//
// Defined in [<SDL3/SDL_render.h>].
//
// SDL's rendering functions operate on a backbuffer; that is, calling a
// rendering function such as SDL_RenderLine() does not directly put a line on
// the screen, but rather updates the backbuffer. As such, you compose your
// entire scene and present the composed backbuffer to the screen as a complete
// picture.
//
// Therefore, when using SDL's rendering API, one does all drawing intended for
// the frame, and then calls this function once per frame to present the final
// drawing to the user.
//
// The backbuffer should be considered invalidated after each present; do not
// assume that previous contents will exist between frames. You are strongly
// encouraged to call SDL_RenderClear() to initialize the backbuffer before
// starting each new frame's drawing, even if you plan to overwrite every
// pixel.
//
// Please note, that in case of rendering to a texture - there is no need to
// call SDL_RenderPresent after drawing needed objects to a texture, and should
// not be done; you are only required to change back the rendering target to
// default via SDL_SetRenderTarget(renderer, NULL) afterwards, as textures by
// themselves do not have a concept of backbuffers. Calling SDL_RenderPresent
// while rendering to a texture will fail.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// [<SDL3/SDL_render.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_render.h
func (rr *Renderer) Present() error {
	if !C.SDL_RenderPresent(rr.ptr) {
		return getError()
	}
	return nil
}

// SetDrawColor sets the color used for drawing operations.
//
// Defined in [<SDL3/SDL_render.h>].
//
// Set the color for drawing or filling rectangles, lines, and points, and for
// [Renderer.Clear].
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// [<SDL3/SDL_render.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_render.h
func (rr *Renderer) SetDrawColor(r, g, b, a uint8) error {
	if !C.SDL_SetRenderDrawColor(
		rr.ptr, C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)) {

		return getError()
	}
	return nil
}

// LoadTexture loads an image from a filesystem path into a texture.
//
// Defined in [<SDL3_image/SDL_image.h>].
//
// An SDL_Texture represents an image in GPU memory, usable by SDL's 2D Render
// API. This can be significantly more efficient than using a CPU-bound
// SDL_Surface if you don't need to manipulate the image directly after loading
// it.
//
// If the loaded image has transparency or a colorkey, a texture with an alpha
// channel will be created. Otherwise, SDL_image will attempt to create an
// SDL_Texture in the most format that most reasonably represents the image
// data (but in many cases, this will just end up being 32-bit RGB or 32-bit
// RGBA).
//
// There is a separate function to read files from an SDL_IOStream, if you need
// an i/o abstraction to provide data from anywhere instead of a simple
// filesystem read; that function is IMG_LoadTexture_IO().
//
// If you would rather decode an image to an SDL_Surface (a buffer of pixels in
// CPU memory), call IMG_Load() instead.
//
// When done with the returned texture, the app should dispose of it with a
// call to SDL_DestroyTexture().
//
// This function is available since SDL_image 3.0.0.
//
// See Also
//   - IMG_LoadTextureTyped_IO
//   - IMG_LoadTexture_IO
//
// [<SDL3_image/SDL_image.h>]: https://github.com/libsdl-org/SDL_image/blob/main/include/SDL3_image/SDL_image.h
func (rr *Renderer) LoadTexture(file string) (*Texture, error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	t := C.IMG_LoadTexture(rr.ptr, cFile)
	if t == nil {
		return nil, getError()
	}

	return &Texture{ptr: t}, nil
}

// PollEvent polls for currently pending events.
//
// Defined in [<SDL3/SDL_events.h>].
//
// Returns true if this got an event or false if there are none available.
//
// If event is not NULL, the next event is removed from the queue and stored in
// the SDL_Event structure pointed to by event.
//
// If event is NULL, it simply returns true if there is an event in the queue,
// but will not remove it from the queue.
//
// As this function may implicitly call SDL_PumpEvents(), you can only call
// this function in the thread that initialized the video subsystem.
//
// SDL_PollEvent() is the favored way of receiving system events since it can
// be done from the main loop and does not suspend the main loop while waiting
// on an event to be posted.
//
// The common practice is to fully process the event queue once every frame,
// usually as a first step before updating the game's state:
//
//	while (game_is_still_running) {
//	    SDL_Event event;
//	    while (SDL_PollEvent(&event)) {  // poll until all events are handled!
//	        // decide what to do with this event.
//	    }
//
//	    // update game state, draw the current frame
//	}
//
// Note that Windows (and possibly other platforms) has a quirk about how it
// handles events while dragging/resizing a window, which can cause this
// function to block for significant amounts of time. Technical explanations
// and solutions are discussed on the wiki:
//
// https://wiki.libsdl.org/SDL3/AppFreezeDuringDrag
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// [<SDL3/SDL_events.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_events.h
func PollEvent() (*Event, bool) {
	var e C.SDL_Event
	if !C.SDL_PollEvent(&e) { //nolint:gocritic
		return nil, false
	}
	return convertEvent(&e), true
}

// GetTicksNS returns the number of nanoseconds since the SDL library
// initialization.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// Defined in <SDL3/SDL_timer.h>
//
// [<SDL3/SDL_timer.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_timer.h
func GetTicksNS() uint64 {
	return uint64(C.SDL_GetTicksNS())
}

// Texture is an efficient driver-specific representation of pixel data.
//
// Defined in [<SDL3/SDL_render.h>].
//
// This struct is available since SDL 3.2.0.
//
// See Also
//   - SDL_CreateTexture
//   - SDL_CreateTextureFromSurface
//   - SDL_CreateTextureWithProperties
//   - [Texture.Destroy]
//
// [<SDL3/SDL_render.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_render.h
type Texture struct {
	ptr *C.struct_SDL_Texture
}

// Destroy destroys the specified texture.
//
// Defined in [<SDL3/SDL_render.h>].
//
// Passing NULL or an otherwise invalid texture will set the SDL error message
// to "Invalid texture".
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// See Also
//   - SDL_CreateTexture
//   - SDL_CreateTextureFromSurface
//
// [<SDL3/SDL_render.h>]: https://github.com/libsdl-org/SDL/blob/main/include/SDL3/SDL_render.h
func (t *Texture) Destroy() {
	if t.ptr != nil {
		C.SDL_DestroyTexture(t.ptr)
		t.ptr = nil
	}
}
