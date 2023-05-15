package shapes

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

func TestSphereLocalNormalAt(t *testing.T) {
	var tests = []struct {
		sphere   *Sphere
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			sphere:   NewSphere(),
			point:    tuple.NewPoint(1, 0, 0),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			sphere:   NewSphere(),
			point:    tuple.NewPoint(0, 1, 0),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			sphere:   NewSphere(),
			point:    tuple.NewPoint(0, 0, 1),
			expected: tuple.NewVector(0, 0, 1),
		},
		{
			sphere:   NewSphere(),
			point:    tuple.NewPoint(1, 0, 0),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			sphere:   NewSphere(),
			point:    tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
			expected: tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
		},
		{
			sphere:   NewSphere(),
			point:    tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
			expected: tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0).Normalize(),
		},
	}

	for _, test := range tests {
		if got := test.sphere.LocalNormalAt(test.point, Intersection{}); !got.Equal(test.expected) {
			t.Errorf("Sphere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", test.sphere, test.point, got, test.expected)
		}
	}
}

func TestTransformations(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	sphere := NewSphere()
	sphere.SetTransform(matrix.Scaling(2, 2, 2))
	want := Intersections{
		NewIntersection(3.0, sphere),
		NewIntersection(7.0, sphere),
	}

	got := Intersect(sphere, r)
	for i := range got {
		if got[i].t != want[i].t {
			t.Errorf("incorrect t of intersect:\n%s \n \ngot: \n%f. \nexpected: \n%f", r, got[i].t, want[i].t)
		}
		if got[i].shape != want[i].shape {
			t.Errorf("incorrect Shape of intersect:\n%s \n \ngot: \n%s. \nexpected: \n%s", r, got[i].shape, want[i].shape)
		}
	}
}
