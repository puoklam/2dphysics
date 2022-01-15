package vector

import (
	"math"

	"github.com/puoklam/2dphysics/math/float"
)

// error msg
const (
	NilVector    = "nil vector received"
	DimsNotMatch = "dimensions not match"
	InvalidDims  = "invalid dimensions"
)

// math const
const (
	Pi = math.Pi
)

type Vector struct {
	Dims []float64
}

// method set
// func (v *Vector) AddDims(vals ...float64) {
// 	v.Dims = append(v.Dims, vals...)
// }

func (v *Vector) Normalize() {
	mag := Magnitude(v)
	for i := range v.Dims {
		v.Dims[i] /= mag
	}
}

func (v *Vector) ReverseDir() {
	for i := range v.Dims {
		v.Dims[i] *= -1
	}
}

// validation
func validateUnOp(v *Vector) {
	if v == nil {
		panic(NilVector)
	}
}

func validateBinOp(v1, v2 *Vector) {
	if v1 == nil || v2 == nil {
		panic(NilVector)
	}
	if len(v1.Dims) != len(v2.Dims) {
		panic(DimsNotMatch)
	}
}

// func validateTrigonOp(v1, v2 *Vector) error {
// 	if err := validateBinOp(v1, v2); err != nil {
// 		return err
// 	}
// 	if IsZeroVector(v1) || IsZeroVector(v2) {
// 		return errors.New("zero vector received")
// 	}
// 	return nil
// }

// func binReduce(v1, v2 *Vector, fn func(i int) interface{}, r interface{}) {

// }

// exported functions
func NewVector(vals ...float64) *Vector {
	if vals == nil {
		return &Vector{[]float64{}}
	}
	dims := make([]float64, 0, len(vals))
	dims = append(dims, vals...)
	return &Vector{dims}
}

// check float64 overflow
func Add(v1, v2 *Vector) *Vector {
	validateBinOp(v1, v2)
	dims := make([]float64, 0, len(v1.Dims))
	for i := 0; i < cap(dims); i++ {
		dims = append(dims, v1.Dims[i]+v2.Dims[i])
	}
	return &Vector{dims}
}

func Sub(v1, v2 *Vector) *Vector {
	validateBinOp(v1, v2)
	dims := make([]float64, 0, len(v1.Dims))
	for i := 0; i < cap(dims); i++ {
		dims = append(dims, v1.Dims[i]-v2.Dims[i])
	}
	return &Vector{dims}
}

func Mul(v *Vector, f float64) *Vector {
	if v == nil {
		panic(NilVector)
	}
	dims := make([]float64, 0, len(v.Dims))
	for i := 0; i < cap(dims); i++ {
		dims = append(dims, v.Dims[i]*f)
	}
	return &Vector{dims}
}

func Dot(v1, v2 *Vector) float64 {
	validateBinOp(v1, v2)
	var sum float64
	for i := 0; i < len(v1.Dims); i++ {
		sum += v1.Dims[i] * v2.Dims[i]
	}
	return sum
}

func Cross(v1, v2 *Vector) *Vector {
	validateBinOp(v1, v2)
	if len(v1.Dims) > 3 || len(v2.Dims) > 3 {
		panic(InvalidDims)
	}
	dims := make([]float64, 3)
	switch len(v1.Dims) {
	// no op
	// case 0:
	// case 1:
	case 2:
		dims[2] = v1.Dims[0]*v2.Dims[1] - v1.Dims[1]*v2.Dims[0]
	case 3:
		dims[0] = v1.Dims[1]*v2.Dims[2] - v1.Dims[2]*v2.Dims[1]
		dims[1] = v1.Dims[2]*v2.Dims[0] - v1.Dims[0]*v2.Dims[2]
		dims[2] = v1.Dims[0]*v2.Dims[1] - v1.Dims[1]*v2.Dims[0]
	}
	return &Vector{dims}
}

func Magnitude(v *Vector) float64 {
	validateUnOp(v)
	var sum float64
	for _, dim := range v.Dims {
		sum += dim * dim
	}
	return math.Sqrt(sum)
}

func Normalize(v *Vector) *Vector {
	validateUnOp(v)
	dims := make([]float64, 0, len(v.Dims))
	mag := Magnitude(v)
	for _, dim := range v.Dims {
		dims = append(dims, dim/mag)
	}
	return &Vector{dims}
}

func IsUnitVector(v *Vector) bool {
	validateUnOp(v)
	mag := Magnitude(v)
	return float.Equal(mag, 1)
}

func IsZeroVector(v *Vector) bool {
	validateUnOp(v)
	mag := Magnitude(v)
	return float.Equal(mag, 0)
}

func Cos(v1, v2 *Vector) float64 {
	dp := Dot(v1, v2)
	mag1, mag2 := Magnitude(v1), Magnitude(v2)
	if float.Equal(mag1, 0) || float.Equal(mag2, 0) {
		return math.NaN()
	}
	return dp / (mag1 * mag2)
}

func Sin(v1, v2 *Vector) float64 {
	cp := Cross(v1, v2)
	mag1, mag2 := Magnitude(v1), Magnitude(v2)
	if float.Equal(mag1, 0) || float.Equal(mag2, 0) {
		return math.NaN()
	}
	return Magnitude(cp) / (mag1 * mag2)
}

func Angle(v1, v2 *Vector) float64 {
	return math.Acos(Cos(v1, v2))
}

func IsOrthogonal(v *Vector, other ...*Vector) bool {
	if other == nil {
		panic(NilVector)
	}
	for _, vct := range other {
		if !float.Equal(Dot(v, vct), 0) {
			return false
		}
	}
	return true
}
