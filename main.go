package main

import (
	"fmt"

	"github.com/sandmannmax/gosim/internal/engine"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("gosim", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	e := engine.New(window, renderer)

	running := true
	for running {
		e.Render()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
				break
			case *sdl.MouseButtonEvent:
				mbe := event.(*sdl.MouseButtonEvent)
				if mbe.Type == sdl.MOUSEBUTTONDOWN {
					e.AddParticle(float64(mbe.X), -float64(mbe.Y), 0)
				}
				break
			}
		}
	}
}
