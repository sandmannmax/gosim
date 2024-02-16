package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type object struct {
	x         int32
	y         int32
	w         int32
	h         int32
	velocityX float64
	velocityY float64
}

func (o *object) calculate(dt float64) {
	o.x += int32(o.velocityX * dt)
	o.y += int32(o.velocityY * dt)

	if o.x+o.w >= 800 || o.x < 0 {
		o.velocityX *= -1
	}
	if o.y+o.h >= 600 || o.y < 0 {
		o.velocityY *= -1
	}
}

func (o object) draw(surface *sdl.Surface) {
	color := sdl.Color{R: 255, G: 0, B: 0, A: 255}
	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)
	rect := sdl.Rect{X: o.x, Y: o.y, W: o.w, H: o.h}
	surface.FillRect(&rect, pixel)
}

type engine struct {
	window       *sdl.Window
	windowWidth  int32
	windowHeight int32
	surface      *sdl.Surface
	lastUpdate   uint64
	objects      []*object
	fps          uint32
}

func New(window *sdl.Window, surface *sdl.Surface) engine {
	w, h := window.GetSize()
	e := engine{
		window:       window,
		windowWidth:  w,
		windowHeight: h,
		surface:      surface,
		objects: []*object{
			{x: 0, y: 0, w: 200, h: 200, velocityX: 300, velocityY: 150},
		},
		fps: 60,
	}
	e.initialize()
	return e
}

func (e *engine) initialize() {
	e.lastUpdate = sdl.GetPerformanceCounter()
}

func (e *engine) Render() {
	start := sdl.GetPerformanceCounter()
	performanceFrequency := float64(sdl.GetPerformanceFrequency())
	dt := float64(start-e.lastUpdate) / performanceFrequency

	e.surface.FillRect(nil, 0)

	for _, o := range e.objects {
		o.calculate(dt)
		o.draw(e.surface)
	}

	e.window.UpdateSurface()

	e.lastUpdate = start
	end := sdl.GetPerformanceCounter()
	elapsedMs := float64(end-start) / performanceFrequency * 1000.0
	remaining := 1000.0/float64(e.fps) - elapsedMs
	if remaining < 0 {
		remaining = 0
	}

	sdl.Delay(uint32(remaining))
}
