package float

import (
	"math"
)

const (
	EqualityThreshold = 1e-9
	// EqualityThreshold = math.SmallestNonzeroFloat64
)

func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func Equal(a, b float64) bool {
	return math.Abs(a-b) <= EqualityThreshold*Max(1, Max(a, b))
}
