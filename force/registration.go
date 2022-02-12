package force

import (
	"github.com/puoklam/2dphysics/shape"
)

type Registration struct {
	body      shape.Body
	generator Generator
}

func NewRegistration(body shape.Body, generator Generator) *Registration {
	return &Registration{
		body,
		generator,
	}
}

func EqualRegistration(r1, r2 Registration) bool {
	return r1.generator == r2.generator && r1.body == r2.body
}
