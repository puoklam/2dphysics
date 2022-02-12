package physics2d

import (
	"github.com/puoklam/physics2d/force"
	"github.com/puoklam/physics2d/math/vector"
	"github.com/puoklam/physics2d/shape"
)

type World struct {
	registry *force.Registry
	bodies   []*shape.Body
	dt       float64
}

func NewWorld(dt float64) *World {
	return &World{
		force.NewRegistry(),
		make([]*shape.Body, 0),
		dt,
	}
}

func (w *World) FixedUpdate() {
	w.registry.Update(w.dt)

	for _, body := range w.bodies {
		body.Update(w.dt)
	}
}

func (w *World) Update(dt float64) {
	w.FixedUpdate()
}

func (w *World) AddBody(body *shape.Body) {
	w.bodies = append(w.bodies, body)
}

type Gravity struct {
	a *vector.Vector2D
}

func (g Gravity) Update(body *shape.Body, dt float64) {
	body.AddForce(vector.Mul(g.a, body.Mass()))
}

func (g Gravity) Zero() {
}

func NewGravity(a *vector.Vector2D) Gravity {
	return Gravity{vector.Copy(a)}
}
