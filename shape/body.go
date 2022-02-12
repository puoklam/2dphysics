package shape

import (
	"github.com/puoklam/2dphysics/math/vector"
)

type Body struct {
	Center   *vector.Vector2D
	Rotation float64
	force    *vector.Vector2D
	mass     float64
	massInv  float64

	linVelo *vector.Vector2D
	angVelo *vector.Vector2D
}

func (b *Body) Mass() float64 {
	return b.mass
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
	b.force.Add(f)
}

func (b *Body) Update(dt float64) {
	if b.mass == 0 {
		return
	}
	// update velocity
	a := vector.Mul(b.force, b.massInv)
	b.linVelo.Add(vector.Mul(a, dt))

	// update position
	b.Center.Add(vector.Mul(b.linVelo, dt))

	b.ClearForce()
}

func (b *Body) ClearForce() {
	b.force.Zero()
}

func (b *Body) Get() *vector.Vector2D {
	return b.linVelo
}

func NewBody(c *vector.Vector2D, r, m float64) *Body {
	body := &Body{
		c,
		r,
		vector.NewVector(0, 0),
		0,
		0,
		vector.NewVector(0, 0),
		vector.NewVector(0, 0),
	}
	body.SetMass(m)
	return body
}
