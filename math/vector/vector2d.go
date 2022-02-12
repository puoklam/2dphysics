package vector

import (
	"math"

	"github.com/puoklam/2dphysics/math/float"
)

const (
	ErrNilVector = "nil vector received"
)

const (
	Pi = math.Pi
)

type Vector2D struct {
	X, Y float64
}

func (v *Vector2D) Normalize() {
	mag := Magnitude(v)
	v.X /= mag
	v.Y /= mag
}

func (v *Vector2D) ReverseDir() {
	v.X *= -1
	v.Y *= -1
}

func (v *Vector2D) Add(v2 *Vector2D) {
	validateBinOp(v, v2)
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vector2D) Sub(v2 *Vector2D) {
	validateBinOp(v, v2)
	v.X -= v2.X
	v.Y -= v2.Y
}

func (v *Vector2D) Mul(f float64) {
	validateUnOp(v)
	v.X *= f
	v.Y *= f
}

func (v *Vector2D) Zero() {
	validateUnOp(v)
	v.X, v.Y = 0, 0
}

func (v *Vector2D) Rotate(o *Vector2D, t float64) {
	if t != 0 {
		x, y := v.X-o.X, v.Y-o.Y
		cos, sin := math.Cos(t), math.Sin(t)
		x, y = x*cos-y*sin, x*sin+y*cos
		v.X, v.Y = o.X+x, o.Y+y
	}
}

// validation
func validateUnOp(v *Vector2D) {
	if v == nil {
		panic(ErrNilVector)
	}
}

func validateBinOp(v1, v2 *Vector2D) {
	if v1 == nil || v2 == nil {
		panic(ErrNilVector)
	}
}

// exported functions
func NewVector(x, y float64) *Vector2D {
	return &Vector2D{x, y}
}

func Copy(v *Vector2D) *Vector2D {
	if v == nil {
		return nil
	}
	return NewVector(v.X, v.Y)
}

// TODO: check float64 overflow
func Add(v1, v2 *Vector2D) *Vector2D {
	validateBinOp(v1, v2)
	return &Vector2D{v1.X + v2.X, v1.Y + v2.Y}
}

func Sub(v1, v2 *Vector2D) *Vector2D {
	validateBinOp(v1, v2)
	return &Vector2D{v1.X - v2.X, v1.Y - v2.Y}
}

func Mul(v *Vector2D, f float64) *Vector2D {
	if v == nil {
		panic(ErrNilVector)
	}
	return &Vector2D{v.X * f, v.Y * f}
}

func MulVector2D(v1, v2 *Vector2D) *Vector2D {
	validateBinOp(v1, v2)
	return &Vector2D{v1.X * v2.X, v1.Y * v2.Y}
}

func Dot(v1, v2 *Vector2D) float64 {
	validateBinOp(v1, v2)
	return v1.X*v2.X + v1.Y*v2.Y
}

func Cross(v1, v2 *Vector2D) *Vector2D {
	validateBinOp(v1, v2)
	return &Vector2D{0, 0}
}

func Magnitude(v *Vector2D) float64 {
	validateUnOp(v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MagnitudeSquared(v *Vector2D) float64 {
	validateUnOp(v)
	return v.X*v.X + v.Y*v.Y
}

func Normalize(v *Vector2D) *Vector2D {
	validateUnOp(v)
	cp := Copy(v)
	cp.Normalize()
	return cp
}

func IsUnitVector(v *Vector2D) bool {
	validateUnOp(v)
	mag := MagnitudeSquared(v)
	return float.Equal(mag, 1)
}

func IsZeroVector(v *Vector2D) bool {
	validateUnOp(v)
	mag := MagnitudeSquared(v)
	return float.Equal(mag, 0)
}

func Cos(v1, v2 *Vector2D) float64 {
	dp := Dot(v1, v2)
	mag1, mag2 := Magnitude(v1), Magnitude(v2)
	if float.Equal(mag1, 0) || float.Equal(mag2, 0) {
		return math.NaN()
	}
	return dp / (mag1 * mag2)
}

func Sin(v1, v2 *Vector2D) float64 {
	cp := Cross(v1, v2)
	mag1, mag2 := Magnitude(v1), Magnitude(v2)
	if float.Equal(mag1, 0) || float.Equal(mag2, 0) {
		return math.NaN()
	}
	return Magnitude(cp) / (mag1 * mag2)
}

func Angle(v1, v2 *Vector2D) float64 {
	return math.Acos(Cos(v1, v2))
}

func IsOrthogonal(v1, v2 *Vector2D) bool {
	validateBinOp(v1, v2)
	return float.Equal(Dot(v1, v2), 0)
}

func Rotate(v, o *Vector2D, t float64) *Vector2D {
	cp := Copy(v)
	cp.Rotate(o, t)
	return cp
}

func Projection(v1, v2, o *Vector2D) *Vector2D {
	if v1 == nil || v2 == nil || o == nil {
		panic(ErrNilVector)
	}
	a := Sub(v1, o)
	b := Sub(v2, o)
	t := Dot(a, b) / Dot(b, b)
	return Add(o, Mul(v2, t))
}
