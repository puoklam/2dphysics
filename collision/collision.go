package collision

import (
	"math"

	"github.com/puoklam/physics2d/math/vector"
	"github.com/puoklam/physics2d/shape"
)

type Collision struct {
	Body1, Body2 *shape.Body
	Manifold     *Manifold
}

const (
	TBC = iota
	Circle
	Rect
)

func FindCollision(c1, c2 shape.Collider) (*Manifold, bool) {
	type1, type2 := TBC, TBC
	switch c1.(type) {
	case *shape.Circle:
		type1 = Circle
	case *shape.Rect:
		type1 = Rect
	}
	switch c2.(type) {
	case *shape.Circle:
		type2 = Circle
	case *shape.Rect:
		type2 = Rect
	}
	if type1 == TBC || type2 == TBC {
		return nil, false
	}
	if type1 == Circle && type2 == Circle {
		circle1, _ := c1.(*shape.Circle)
		circle2, _ := c2.(*shape.Circle)
		return findColOfCircles(circle1, circle2)
	} else {
		panic("Shape not supported")
	}
}

func findColOfCircles(c1, c2 *shape.Circle) (*Manifold, bool) {
	// sum radii, distance
	sr, dist := c1.Radius+c2.Radius, vector.Sub(c2.Center, c1.Center)

	if vector.MagnitudeSquared(dist)-sr*sr > 0 {
		return nil, false
	}
	depth := math.Abs(vector.Magnitude(dist)-sr) * 0.5
	normal := vector.Normalize(dist)

	// contact point
	contact := vector.Add(vector.Copy(c1.Center), vector.Mul(normal, c1.Radius-depth))
	res := NewManifold(normal, depth)
	res.AddContact(contact)
	return res, true
}
