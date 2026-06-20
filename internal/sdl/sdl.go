package sdl

// #include <SDL3/SDL.h>
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
	C.SDL_DestroyWindow(w.ptr)
	w.ptr = nil
}

func (w *Window) CreateRenderer() (*Renderer, error) {
	cRenderer := C.SDL_CreateRenderer(w.ptr, nil)
	if cRenderer == nil {
		return nil, getError()
	}
	return &Renderer{ptr: cRenderer}, nil
}

func (rr *Renderer) Destroy() {
	C.SDL_DestroyRenderer(rr.ptr)
	rr.ptr = nil
}

func (rr *Renderer) Clear() error {
	if !C.SDL_RenderClear(rr.ptr) {
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

func PollEvent() (*Event, bool) {
	var e C.SDL_Event
	if !C.SDL_PollEvent(&e) {
		return nil, false
	}
	return convertEvent(&e), true
}
