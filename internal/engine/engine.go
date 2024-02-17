package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type engine struct {
	window       *sdl.Window
	windowWidth  int32
	windowHeight int32
	renderer     *sdl.Renderer
	lastUpdate   uint64
	particles    []*particle
	fps          uint32
}

func New(window *sdl.Window, renderer *sdl.Renderer) engine {
	w, h := window.GetSize()
	e := engine{
		window:       window,
		windowWidth:  w,
		windowHeight: h,
		renderer:     renderer,
		particles:    []*particle{},
		fps:          60,
	}
	e.initialize()
	return e
}

func (e *engine) initialize() {
	e.lastUpdate = sdl.GetPerformanceCounter()
}

func (e *engine) AddParticle(x float64, y float64, color uint32) {
	e.particles = append(e.particles, &particle{
		position:     vector3{x: x, y: y, z: 0},
		velocity:     vector3{x: 0, y: 0, z: 0},
		acceleration: vector3{},
		damping:      0.995,
		inverseMass:  1,
		color:        color,
	})
}

func (e *engine) Render() {
	start := sdl.GetPerformanceCounter()
	performanceFrequency := float64(sdl.GetPerformanceFrequency())
	dt := float64(start-e.lastUpdate) / performanceFrequency

	e.renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
	e.renderer.Clear()

	for _, p := range e.particles {
		p.calculate(dt)
		p.draw(e.renderer)
	}

	e.renderer.Present()

	e.lastUpdate = start
	end := sdl.GetPerformanceCounter()
	elapsedMs := float64(end-start) / performanceFrequency * 1000.0
	remaining := 1000.0/float64(e.fps) - elapsedMs
	if remaining < 0 {
		remaining = 0
	}

	sdl.Delay(uint32(remaining))
}
