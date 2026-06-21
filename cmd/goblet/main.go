// Goblet is the main executable.
package main

import (
	"log"

	"github.com/snailbaron/goblet/internal/sdl"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Panic(err)
	}
	if err := config.Save(); err != nil {
		log.Panic(err)
	}

	if err := sdl.Init(sdl.InitVideo | sdl.InitEvents); err != nil {
		log.Panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Goblet", config.WindowWidth, config.WindowHeight, 0)
	if err != nil {
		log.Panic(err)
	}
	defer window.Destroy()

	renderer, err := window.CreateRenderer()
	if err != nil {
		log.Panic(err)
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
			case sdl.EventQuit:
				done = true
			case sdl.EventKeyDown:
				k := e.Key()
				if k.Key == sdl.KEscape {
					done = true
				}
			default:
			}
		}

		if done {
			break
		}

		if framesPassed := timer.Cutoff(); framesPassed > 0 {
			r, g, b, a := config.BackgroundColor[0],
				config.BackgroundColor[1], config.BackgroundColor[2],
				config.BackgroundColor[3]
			if err := renderer.SetDrawColor(r, g, b, a); err != nil {
				log.Panic(err)
			}

			if err := renderer.Clear(); err != nil {
				log.Panic(err)
			}

			if err := renderer.Present(); err != nil {
				log.Panic(err)
			}
		}

		timer.Relax()
	}
}
