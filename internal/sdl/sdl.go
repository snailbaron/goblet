package sdl

// #include <SDL3/SDL.h>
// #include <SDL3_image/SDL_image.h>
//
// #include <stdlib.h>
import "C"

import "unsafe"

type Error struct {
	message string
}

func (e Error) Error() string {
	return e.message
}

func getError() Error {
	return Error{message: C.GoString(C.SDL_GetError())}
}

func Init(flags InitFlags) error {
	if !C.SDL_Init(C.SDL_InitFlags(flags)) {
		return getError()
	}
	return nil
}

func Quit() {
	C.SDL_Quit()
}

type Window struct{ ptr *C.SDL_Window }
type Renderer struct{ ptr *C.SDL_Renderer }

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

func (w *Window) Destroy() {
	if w.ptr != nil {
		C.SDL_DestroyWindow(w.ptr)
		w.ptr = nil
	}
}

func (w *Window) CreateRenderer() (*Renderer, error) {
	cRenderer := C.SDL_CreateRenderer(w.ptr, nil)
	if cRenderer == nil {
		return nil, getError()
	}
	return &Renderer{ptr: cRenderer}, nil
}

func (rr *Renderer) Destroy() {
	if rr.ptr != nil {
		C.SDL_DestroyRenderer(rr.ptr)
		rr.ptr = nil
	}
}

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

func (rr *Renderer) SetDrawColor(r, g, b, a uint8) error {
	if !C.SDL_SetRenderDrawColor(
		rr.ptr, C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)) {

		return getError()
	}
	return nil
}

func (rr *Renderer) LoadTexture(file string) (*Texture, error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	t := C.IMG_LoadTexture(rr.ptr, cFile)
	if t == nil {
		return nil, getError()
	}
	return &Texture{ptr: t}, nil
}

func PollEvent() (*Event, bool) {
	var e C.SDL_Event
	if !C.SDL_PollEvent(&e) {
		return nil, false
	}
	return convertEvent(&e), true
}

// Get the number of nanoseconds since SDL library initialization.
//
// Returns an unsigned 64-bit value representing the number of nanoseconds
// since the SDL library initialized.
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

type Texture struct {
	ptr *C.struct_SDL_Texture
}

func (t *Texture) Destroy() {
	if t.ptr != nil {
		C.SDL_DestroyTexture(t.ptr)
		t.ptr = nil
	}
}
