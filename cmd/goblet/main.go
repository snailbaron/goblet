package main

import (
	"log"

	"github.com/snailbaron/goblet/internal/sdl"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := config.Save(); err != nil {
		log.Fatal(err)
	}

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

	timer := NewFrameTimer(config.FPS)
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
			if err := renderer.SetDrawColor(30, 30, 30, 255); err != nil {
				log.Fatal(err)
			}
			if err := renderer.Clear(); err != nil {
				log.Fatal(err)
			}
			if err := renderer.Present(); err != nil {
				log.Fatal(err)
			}
		}

		timer.Relax()
	}
}
