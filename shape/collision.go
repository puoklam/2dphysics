package shape

import (
	"github.com/puoklam/2dphysics/math/float"
	"github.com/puoklam/2dphysics/math/vector"
)

type Collider interface {
}

func IsPointOnLine(p *vector.Vector2D, l *Line) bool {
	if l.Start.X == l.End.X {
		// vertical line
		return float.Equal(p.X, l.Start.X)
	}
	m := l.Slope()
	c := l.Start.Y - m*l.Start.X
	return float.Equal(p.Y, m*p.X+c)
}

func IsPointInCircle(p *vector.Vector2D, c *Circle) bool {
	// check radius (squared radius)
	return float.Equal(vector.MagnitudeSquared(p), c.Radius*c.Radius)
}

func IsPointInRect(p *vector.Vector2D, r *Rect) bool {
	// return p.Dims[0] >= r.Min().Dims[0] && p.Dims[0] <= r.Max().Dims[0] && p.Dims[1] >= r.Min().Dims[1] && p.Dims[1] <= r.Max().Dims[1]
	rotated := vector.Rotate(p, r.Center, -r.Rotation)
	return rotated.X >= r.Min().X && rotated.X <= r.Max().X && rotated.Y >= r.Min().Y && rotated.Y <= r.Max().Y
}

func IsLineInCircle(l *Line, c *Circle) bool {
	if IsPointInCircle(l.Start, c) || IsPointInCircle(l.End, c) {
		return true
	}
	proj := vector.Projection(l.End, c.Center, l.Start)
	return IsPointInCircle(proj, c)
}

func IsLineInRect(l *Line, r *Rect) bool {
	if IsPointInRect(l.Start, r) || IsPointInRect(l.End, r) {
		return true
	}
	// rorate line by -angle rad
	rs := vector.Rotate(l.Start, r.Center, r.Rotation)
	re := vector.Rotate(l.End, r.Center, r.Rotation)
	rotated := NewLine(rs, re)

	u := vector.Normalize(vector.Sub(rotated.End, rotated.Start))
	if u.X != 0 {
		u.X = 1 / u.X
	}
	if u.Y != 0 {
		u.Y = 1 / u.Y
	}
	a := vector.MulVector2D(r.Max(), u)
	b := vector.MulVector2D(r.Min(), u)

	tMin := float.Max(float.Min(a.X, b.X), float.Min(a.Y, b.Y))
	tMax := float.Min(float.Max(a.X, b.X), float.Max(a.Y, b.Y))

	if tMax < 0 || tMin > tMax {
		return false
	}

	t := tMax
	if tMin >= 0 {
		t = tMin
	}
	return t > 0 && t*t < l.DistanceSquared()
}
