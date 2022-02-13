package physics2d

import (
	"math"

	"github.com/puoklam/physics2d/collision"
	"github.com/puoklam/physics2d/force"
	"github.com/puoklam/physics2d/math/vector"
	"github.com/puoklam/physics2d/shape"
)

type World struct {
	registry *force.Registry
	bodies   []*shape.Body
	dt       float64

	collisions []collision.Collision
}

func NewWorld(dt float64) *World {
	return &World{
		force.NewRegistry(),
		make([]*shape.Body, 0),
		dt,
		make([]collision.Collision, 0),
	}
}

func (w *World) FixedUpdate() {
	// update force
	w.registry.Update(w.dt)

	// find collisions
	w.collisions = nil
	for i, b1 := range w.bodies {
		for _, b2 := range w.bodies[i+1:] {
			if b1.Mass() == shape.InfMass && b2.Mass() == shape.InfMass {
				continue
			}
			if c1, c2 := b1.Collider, b2.Collider; c1 != nil && c2 != nil {
				if m, ok := collision.FindCollision(c1, c2); ok {
					w.collisions = append(w.collisions, collision.Collision{
						Body1:    b1,
						Body2:    b2,
						Manifold: m,
					})
				}
			}
		}
	}

	// resolve collisions (impulse resolution)
	for k := 0; k < 8; k++ {
		for _, c := range w.collisions {
			for j := 0; j < len(c.Manifold.Contacts); j++ {
				applyImpulse(c.Body1, c.Body2, c.Manifold)
			}
		}
	}
	// update velocity
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

func applyImpulse(b1, b2 *shape.Body, m *collision.Manifold) {
	massInv1, massInv2 := b1.MassInv(), b2.MassInv()
	massInvSum := massInv1 + massInv2
	if massInvSum == 0 {
		return
	}

	relVelo := vector.Sub(b2.LinVelo, b1.LinVelo)
	relNormal := vector.Normalize(m.Normal)

	dp := vector.Dot(relVelo, relNormal)
	if dp > 0 {
		return
	}

	e := math.Min(b1.Cor, b2.Cor)
	j := -(1 + e) * dp / massInvSum
	if len(m.Contacts) > 0 && j != 0 {
		j /= float64(len(m.Contacts))
	}

	impulse := vector.Mul(vector.Copy(relNormal), j)
	b1.LinVelo.Add(vector.Mul(vector.Copy(impulse), -massInv1))
	b2.LinVelo.Add(vector.Mul(vector.Copy(impulse), massInv2))
}
