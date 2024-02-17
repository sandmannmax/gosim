package engine

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type particle struct {
	position     vector3
	velocity     vector3
	acceleration vector3
	damping      float64
	inverseMass  float64
	color        uint32
}

func (p *particle) calculate(dt float64) {
	p.position.addScaled(p.velocity, dt)

	acceleration := p.acceleration
	acceleration.add(vector3{y: -10}) // add gravity

	p.velocity.addScaled(acceleration, dt)
	p.velocity.mul(math.Pow(p.damping, dt)) // impose drag

	// if (p.position.x+float64(p.rect.W) >= 800 && p.velocity.x > 0) || (p.position.x <= 0 && p.velocity.x < 0) {
	// 	p.velocity.x *= -1
	// }
	// if (p.position.y+float64(p.rect.H) >= 600 && p.velocity.y > 0) || (p.position.y <= 0 && p.velocity.y < 0) {
	// 	p.velocity.y *= -1
	// }
}

func (p particle) draw(renderer *sdl.Renderer) {
	factor := 1.0
	x := float32(p.position.x * factor)
	y := float32(-p.position.y * factor)

	color := sdl.Color{255, 255, 255, 255}

	var halfWidth float32 = 8.0
	renderer.RenderGeometry(nil, []sdl.Vertex{
		{sdl.FPoint{x - halfWidth, y - halfWidth}, color, sdl.FPoint{0, 0}},
		{sdl.FPoint{x - halfWidth, y + halfWidth}, color, sdl.FPoint{0, 0}},
		{sdl.FPoint{x + halfWidth, y + halfWidth}, color, sdl.FPoint{0, 0}},
	}, nil)
	renderer.RenderGeometry(nil, []sdl.Vertex{
		{sdl.FPoint{x - halfWidth, y - halfWidth}, color, sdl.FPoint{0, 0}},
		{sdl.FPoint{x + halfWidth, y + halfWidth}, color, sdl.FPoint{0, 0}},
		{sdl.FPoint{x + halfWidth, y - halfWidth}, color, sdl.FPoint{0, 0}},
	}, nil)
}
