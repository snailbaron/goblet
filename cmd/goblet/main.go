package main

import (
	"log"
	"time"

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

	fps := 60
	frameDuration := time.Duration(int64(time.Second) / int64(fps))
	lastFrame := 0

	start := time.Now()
	for {
		done := false
		for !done {
			e, ok := sdl.PollEvent()
			if !ok {
				break
			}

			switch e.(type) {
			case sdl.QuitEvent:
				done = true
			}
		}

		if done {
			break
		}

		now := time.Now()
		frame := int(now.Sub(start) / frameDuration)
		framesPassed := frame - lastFrame
		lastFrame = frame

		if framesPassed > 0 {
			for range framesPassed {
				renderer.SetDrawColor(30, 30, 30, 255)
				renderer.Clear()
			}
		}

		nextFrameTime := start.Add(time.Duration(frame + 1) * frameDuration)
		time.Sleep(nextFrameTime.Sub(now))
	}
}
