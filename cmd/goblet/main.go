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

	screen := NewScreen(config)
	if err := screen.Init(); err != nil {
		log.Panic(err)
	}
	defer screen.Destroy()

	screen.AddWidget(&Button{
		Position:        sdl.FRect{X: 200, Y: 100, W: 200, H: 100},
		BorderColor:     NewRGBColor(10, 10, 10),
		BorderWidth:     5,
		BackgroundColor: NewRGBColor(200, 200, 150),
	})

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
			for range framesPassed {
				if err := screen.Update(timer.Delta()); err != nil {
					log.Panic(err)
				}
			}

			if err := screen.Render(); err != nil {
				log.Panic(err)
			}
		}

		timer.Relax()
	}
}
