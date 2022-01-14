package float

import (
	"math"
)

const (
	EqualityThreshold = 1e-9
)

func Equal(a, b float64) bool {
	return math.Abs(a-b) <= EqualityThreshold
}
