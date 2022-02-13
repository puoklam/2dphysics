package shape

import (
	"github.com/puoklam/physics2d/math/vector"
)

const (
	InfMass = 0
)

type Collider interface {
}

type Body struct {
	Center   *vector.Vector2D
	Rotation float64
	force    *vector.Vector2D
	mass     float64
	massInv  float64

	LinVelo *vector.Vector2D
	AngVelo *vector.Vector2D

	// TODO: extract collider out, now self referencing (circle.Body.Collider = circle)
	Collider Collider
	// coef. of restitution
	Cor float64
}

func (b *Body) Mass() float64 {
	return b.mass
}

func (b *Body) MassInv() float64 {
	return b.massInv
}

func (b *Body) SetMass(m float64) {
	if m == 0 {
		b.mass = 0
		b.massInv = 0
		return
	}
	b.mass = m
	b.massInv = 1 / m
}

func (b *Body) AddForce(f *vector.Vector2D) {
	b.force.Add(vector.Copy(f))
}

func (b *Body) Update(dt float64) {
	if b.mass == 0 {
		return
	}
	// update velocity
	a := vector.Mul(b.force, b.massInv)
	b.LinVelo.Add(vector.Mul(a, dt))

	// update position
	b.Center.Add(vector.Mul(b.LinVelo, dt))

	b.ClearForce()
}

func (b *Body) ClearForce() {
	b.force.Zero()
}

func NewBody(c *vector.Vector2D, r, m float64) *Body {
	body := &Body{
		vector.Copy(c),
		r,
		vector.NewVector(0, 0),
		0,
		0,
		vector.NewVector(0, 0),
		vector.NewVector(0, 0),
		nil,
		1,
	}
	body.SetMass(m)
	return body
}
