package collision

import "github.com/puoklam/physics2d/math/vector"

type Manifold struct {
	Normal   *vector.Vector2D
	Contacts []*vector.Vector2D
	Depth    float64
}

func NewManifold(n *vector.Vector2D, d float64) *Manifold {
	return &Manifold{
		vector.Copy(n),
		make([]*vector.Vector2D, 0),
		d,
	}
}

func (m *Manifold) AddContact(c *vector.Vector2D) {
	m.Contacts = append(m.Contacts, vector.Copy(c))
}
