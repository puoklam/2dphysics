package force

import (
	"github.com/puoklam/physics2d/shape"
)

type Generator interface {
	Update(body *shape.Body, dt float64)
	Zero()
}
