package physics2d

import (
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
			got := body1.Center.Y
			y += tt.want[j]
			if !float.Equal(got, y) {
				t.Errorf("%d. got %v; want %v", i, got, y)
			}
		}
	}
}

// func TestCollide(t *testing.T) {
// 	world := NewWorld(1.0 / 6)
// 	gravity := NewGravity(vector.NewVector(0, -10))

// 	circle1 := shape.NewCircle(vector.NewVector(10, 500), 10, 100)
// 	circle2 := shape.NewCircle(vector.NewVector(10, 300), 20, 200)

// 	world.AddBody(circle1.Body)
// 	world.AddBody(circle2.Body)
// 	// world.AddBody(body2)
// 	world.registry.Add(circle1.Body, gravity)

// 	for j := 0; j < 48; j++ {
// 		fmt.Println(circle1.Center, circle2.Center)
// 		world.Update(1.0 / 6)
// 		t.Errorf("OK")
// 	}
// }
