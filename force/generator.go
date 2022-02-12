package force

import (
	"github.com/puoklam/2dphysics/shape"
)

type Generator interface {
	update(body shape.Body, dt float64)
	zero()
}
