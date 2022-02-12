package shape

import (
	"math"

	"github.com/puoklam/2dphysics/math/vector"
)

const (
	ErrInvalidRadius = "invalid radius"
)

type Shape interface {
	Area() float64
}

/**
 * Circle struct
 */
type Circle struct {
	*Body
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func NewCircle(c *vector.Vector2D, r float64, m float64) *Circle {
	if r < 0 {
		panic(ErrInvalidRadius)
	}
	body := NewBody(vector.Copy(c), 0, m)
	return &Circle{body, r}
}

/**
 * Triangle struct
 */
// type Tri struct {
// 	vertices []*vector.Vector
// }

// func (t *Tri) Area() float64 {
// 	// 1/2 * a * b * sin(c)
// 	// sqrt(s * (s-a) * (s-b) * (s-c)) where s = (a+b+c)/2
// 	a := vector.Sub(t.vertices[1], t.vertices[0])
// 	b := vector.Sub(t.vertices[2], t.vertices[0])
// 	return vector.Magnitude(a) * vector.Magnitude(b) * vector.Sin(a, b) * 0.5
// }

// func NewTri(v1, v2, v3 *vector.Vector) *Tri {
// 	if v1 == nil || v2 == nil || v3 == nil {
// 		panic(vector.ErrNilVector)
// 	}
// 	// cv1, cv2, cv3 := *v1, *v2, *v3
// 	v1, v3, v3 = vector.Copy(v1), vector.Copy(v2), vector.Copy(v3)
// 	return &Tri{[]*vector.Vector{v1, v2, v3}}
// }

/**
 * Rectangle struct
 */
type Rect struct {
	*Body
	halfDiag *vector.Vector2D
}

func (r Rect) Area() float64 {
	return math.Abs((r.Max().Y - r.Min().Y) * (r.Max().X - r.Min().X))
}

func (r Rect) Min() *vector.Vector2D {
	return vector.Sub(r.Center, r.halfDiag)
}

func (r Rect) Max() *vector.Vector2D {
	return vector.Add(r.Center, r.halfDiag)
}

func NewRect(min, max *vector.Vector2D, m float64) *Rect {
	if min == nil || max == nil {
		panic(vector.ErrNilVector)
	}
	diag := vector.Sub(max, min)
	halfDiag := vector.Mul(diag, 0.5)
	center := vector.Add(min, halfDiag)
	body := NewBody(center, 0, m)
	return &Rect{body, halfDiag}
}
