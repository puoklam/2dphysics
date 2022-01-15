package shape

import (
	"math"

	"github.com/puoklam/2dphysics/math/vector"
)

const (
	InvalidRadius = "invalid radius"
)

type Shape interface {
	Area() float64
}

/**
 * Circle struct
 */
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func NewCircle(r float64) *Circle {
	if r < 0 {
		panic(InvalidRadius)
	}
	return &Circle{r}
}

/**
 * Triangle struct
 */
type Tri struct {
	vertices []*vector.Vector
}

func (t *Tri) Area() float64 {
	// 1/2 * a * b * sin(c)
	// sqrt(s * (s-a) * (s-b) * (s-c)) where s = (a+b+c)/2
	a := vector.Sub(t.vertices[1], t.vertices[0])
	b := vector.Sub(t.vertices[2], t.vertices[0])
	return vector.Magnitude(a) * vector.Magnitude(b) * vector.Sin(a, b) * 0.5
}

func NewTri(v1, v2, v3 *vector.Vector) *Tri {
	if v1 == nil || v2 == nil || v3 == nil {
		panic(vector.NilVector)
	}
	cv1, cv2, cv3 := *v1, *v2, *v3
	return &Tri{[]*vector.Vector{&cv1, &cv2, &cv3}}
}

/**
 * Rectangle struct
 */
type Rect struct {
	Min, Max *vector.Vector
}

func (r *Rect) Area() float64 {
	return math.Abs((r.Max.Dims[1] - r.Min.Dims[1]) * (r.Max.Dims[0] - r.Min.Dims[0]))
}

func NewRect(min, max *vector.Vector) *Rect {
	if min == nil || max == nil {
		panic(vector.NilVector)
	}
	cmin, cmax := *min, *max
	return &Rect{&cmin, &cmax}
}
