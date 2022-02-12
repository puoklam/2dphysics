package force

import (
	"github.com/puoklam/2dphysics/shape"
)

type Registry struct {
	registry []*Registration
}

func NewRegistry() *Registry {
	return &Registry{make([]*Registration, 0)}
}

func (r *Registry) Add(body *shape.Body, generator Generator) {
	registration := NewRegistration(body, generator)
	r.registry = append(r.registry, registration)
}

func (r *Registry) Remove(body *shape.Body, generator Generator) {
	registration := NewRegistration(body, generator)
	r.registry = append(r.registry, registration)

	i, n := -1, len(r.registry)
	for j, rg := range r.registry {
		if EqualRegistration(*rg, *registration) {
			i = j
			break
		}
	}
	if i != -1 {
		copy(r.registry[i:], r.registry[i+1:])
		// order doesn't matter
		// r.registry[i], r.registry[n-1] = r.registry[n-1], r.registry[i]
		r.registry = r.registry[:n-1]
	}
}

func (r *Registry) Clear() {
	r.registry = nil
}

func (r *Registry) Update(dt float64) {
	for _, rg := range r.registry {
		rg.generator.Update(rg.body, dt)
	}
}

func (r *Registry) Zero(dt float64) {
	for _, rg := range r.registry {
		rg.generator.Zero()
	}
}
