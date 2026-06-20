package main

import (
	"log"

	"github.com/snailbaron/goblet/internal/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Goblet", 800, 600, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	renderer, err := window.CreateRenderer()
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Destroy()

	timer := NewFrameTimer(60)
	for {
		done := false
		for !done {
			e, ok := sdl.PollEvent()
			if !ok {
				break
			}

			switch e.Type {
			case sdl.EVENT_QUIT:
				done = true
			case sdl.EVENT_KEY_DOWN:
				k := e.Key()
				if k.Key == sdl.SDLK_ESCAPE {
					done = true
				}
			}
		}

		if done {
			break
		}

		if framesPassed := timer.Cutoff(); framesPassed > 0 {
			for range framesPassed {
				renderer.SetDrawColor(30, 30, 30, 255)
				renderer.Clear()
			}
		}

		timer.Relax()
	}
}
