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
			// The normal on a sphere at a point on the x axis.
			sphere:   NewSphere(),
			point:    tuple.NewPoint(1, 0, 0),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			// The normal on a sphere at a point on the y axis.
			sphere:   NewSphere(),
			point:    tuple.NewPoint(0, 1, 0),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			// The normal on a sphere at a point on the z axis.
			sphere:   NewSphere(),
			point:    tuple.NewPoint(0, 0, 1),
			expected: tuple.NewVector(0, 0, 1),
		},
		{
			// The normal on a sphere at a non-axial point.
			sphere:   NewSphere(),
			point:    tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
			expected: tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
		},
		{
			// The normal is a normalized vector.
			sphere:   NewSphere(),
			point:    tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
			expected: tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0).Normalize(),
		},
	}

	for _, test := range tests {
		if result := test.sphere.LocalNormalAt(test.point, Intersection{}); !result.Equal(test.expected) {
			t.Errorf("Sphere normal:\n%s \n point: %s. \nresult: \n%s. \nexpected: \n%s", test.sphere, test.point, result, test.expected)
		}
	}
}

func TestIntersectWithScaledSphere(t *testing.T) {
	// Intersecting a scaled sphere with a ray
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	sphere := NewSphere()
	sphere.SetTransform(matrix.Scaling(2, 2, 2))
	expected := Intersections{
		NewIntersection(3.0, sphere),
		NewIntersection(7.0, sphere),
	}

	result := Intersect(sphere, r)
	for i := range result {
		if result[i].t != expected[i].t {
			t.Errorf("incorrect t of intersect:\n%s \n \nresult: \n%f. \nexpected: \n%f", r, result[i].t, expected[i].t)
		}
		if result[i].shape != expected[i].shape {
			t.Errorf("incorrect Shape of intersect:\n%s \n \nresult: \n%s. \nexpected: \n%s", r, result[i].shape, expected[i].shape)
		}
	}
}

func TestIntersectWithTranslatedSphere(t *testing.T) {
	// Intersecting a translated sphere with a ray
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	sphere := NewSphere()
	sphere.SetTransform(matrix.Translation(5, 0, 0))

	result := Intersect(sphere, r)
	if len(result) != 0 {
		t.Errorf("incorrect number of intersections:\n%s \n \nresult: \n%d. \nexpected: \n%d", r, len(result), 0)
	}
}
