package force

import (
	"github.com/puoklam/2dphysics/shape"
)

type Generator interface {
	Update(body *shape.Body, dt float64)
	Zero()
}
