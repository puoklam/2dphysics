package shape

import (
	"testing"

	"github.com/puoklam/2dphysics/math/float"
	"github.com/puoklam/2dphysics/math/vector"
)

func TestPolygonArea(t *testing.T) {
	tests := []struct {
		in   []*vector.Vector2D
		out  float64
		want bool
	}{
		// {
		// 	in: []*vector.Vector{
		// 		vector.NewVector(0, 0),
		// 		vector.NewVector(6, 0),
		// 		vector.NewVector(2, 4),
		// 	},
		// 	out:  12,
		// 	want: true,
		// },
		{
			in: []*vector.Vector2D{
				vector.NewVector(1, 1),
				vector.NewVector(3, 5),
			},
			out:  8,
			want: true,
		},
	}
	for i, tt := range tests {
		var s Shape
		if len(tt.in) == 2 {
			s = NewRect(tt.in[0], tt.in[1])
		} else if len(tt.in) == 3 {
			// s = NewTri(tt.in[0], tt.in[1], tt.in[2])
		}
		got := float.Equal(s.Area(), tt.out)
		if got != tt.want {
			t.Errorf("%d. got %v; want %v", i, got, tt.want)
		}
	}
}

func TestCircleArea(t *testing.T) {
	tests := []struct {
		in   float64
		out  float64
		want bool
	}{
		{
			in:   5,
			out:  78.53981633974483,
			want: true,
		},
	}
	for i, tt := range tests {
		c := NewCircle(vector.NewVector(0, 0), tt.in)
		got := float.Equal(c.Area(), tt.out)
		if got != tt.want {
			t.Errorf("%d. got %v; want %v", i, got, tt.want)
		}
	}
}
