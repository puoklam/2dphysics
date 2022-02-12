package physics2d

import (
	"fmt"
	"testing"

	"github.com/puoklam/physics2d/math/float"
	"github.com/puoklam/physics2d/math/vector"
	"github.com/puoklam/physics2d/shape"
)

func TestGeavity(t *testing.T) {
	tests := []struct {
		in   float64
		want []float64
	}{
		{
			in: -10,
			want: []float64{
				-10,
				-20,
				-30,
			},
		},
	}
	for i, tt := range tests {
		world := NewWorld(1)
		gravity := NewGravity(vector.NewVector(0, tt.in))

		var y float64 = 100
		body1 := shape.NewBody(vector.NewVector(10, y), 0, 1)

		world.AddBody(body1)
		// world.AddBody(body2)
		world.registry.Add(body1, gravity)

		for j := 0; j < len(tt.want); j++ {
			world.Update(1)
			fmt.Println(body1.Center)
			got := body1.Center.Y
			y += tt.want[j]
			if !float.Equal(got, y) {
				t.Errorf("%d. got %v; want %v", i, got, y)
			}
		}
	}
}
