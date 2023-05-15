package shapes

import (
	"testing"

	"github.com/kaizencodes/glimpse/ray"
	"github.com/kaizencodes/glimpse/tuple"
)

func TestNewTriangle(t *testing.T) {
	p1 := tuple.NewPoint(0, 1, 0)
	p2 := tuple.NewPoint(-1, 0, 0)
	p3 := tuple.NewPoint(1, 0, 0)
	e1 := tuple.NewVector(-1, -1, 0)
	e2 := tuple.NewVector(1, -1, 0)
	normal := tuple.NewVector(0, 0, -1)

	triangle := NewTriangle(p1, p2, p3)

	if triangle.P1() != p1 {
		t.Errorf("Triangle P1\ngot: \n%s. \nexpected: \n%s", triangle.P1(), p1)
	}
	if triangle.P2() != p2 {
		t.Errorf("Triangle P2\ngot: \n%s. \nexpected: \n%s", triangle.P2(), p2)
	}
	if triangle.P3() != p3 {
		t.Errorf("Triangle P3\ngot: \n%s. \nexpected: \n%s", triangle.P3(), p3)
	}
	if triangle.E1() != e1 {
		t.Errorf("Triangle E1\ngot: \n%s. \nexpected: \n%s", triangle.E1(), e1)
	}
	if triangle.E2() != e2 {
		t.Errorf("Triangle E2\ngot: \n%s. \nexpected: \n%s", triangle.E2(), e2)
	}
	if triangle.Normal() != normal {
		t.Errorf("Triangle normal\ngot: \n%s. \nexpected: \n%s", triangle.Normal(), normal)
	}
}

func TestLocalNormalAt(t *testing.T) {
	triangle := NewTriangle(
		tuple.NewPoint(0, 1, 0),
		tuple.NewPoint(-1, 0, 0),
		tuple.NewPoint(1, 0, 0),
	)

	var tests = []struct {
		point tuple.Tuple
	}{
		{
			point: tuple.NewPoint(0, 0.5, 0),
		},
		{
			point: tuple.NewPoint(-0.5, 0.75, 0),
		},
		{
			point: tuple.NewPoint(0.5, 0.25, 0),
		},
	}

	for _, test := range tests {
		if result := triangle.LocalNormalAt(test.point, Intersection{}); !result.Equal(triangle.normal) {
			t.Errorf("Triangle normal:\n point: %s. \nresult: \n%s. \nexpected: \n%s", test.point, result, triangle.normal)
		}
	}
}

func TestLocalNormalAtWithSmoothTriangle(t *testing.T) {
	triangle := NewSmoothTriangle(
		tuple.NewPoint(0, 1, 0),
		tuple.NewPoint(-1, 0, 0),
		tuple.NewPoint(1, 0, 0),
		tuple.NewVector(0, 1, 0),
		tuple.NewVector(-1, 0, 0),
		tuple.NewVector(1, 0, 0),
	)

	// A smooth triangle uses u/v to interpolate the normal

	intersection := NewIntersectionWithUV(1, 0.45, 0.25, triangle)
	expected := tuple.NewVector(-0.5547001962252291, 0.8320502943378437, 0)
	if result := NormalAt(tuple.NewPoint(0, 0, 0), triangle, intersection); !result.Equal(expected) {
		t.Errorf("Triangle normal:\n point: %s. \nresult: \n%s. \nexpected: \n%s", tuple.NewPoint(0, 0, 0), result, expected)
	}
}

func TestTriangleIntersect(t *testing.T) {
	triangle := NewTriangle(
		tuple.NewPoint(0, 1, 0),
		tuple.NewPoint(-1, 0, 0),
		tuple.NewPoint(1, 0, 0),
	)

	var tests = []struct {
		ray      *ray.Ray
		expected Intersections
	}{
		// Intersecting a ray parallel to the triangle
		{
			ray:      ray.NewRay(tuple.NewPoint(0, -1, -2), tuple.NewVector(0, 1, 0)),
			expected: Intersections{},
		},
		// ray misses the p1-p3 edge
		{
			ray:      ray.NewRay(tuple.NewPoint(1, 1, -2), tuple.NewVector(0, 0, 1)),
			expected: Intersections{},
		},
		// ray misses the p1-p2 edge
		{
			ray:      ray.NewRay(tuple.NewPoint(-1, 1, -2), tuple.NewVector(0, 0, 1)),
			expected: Intersections{},
		},
		// ray misses the p2-p3 edge
		{
			ray:      ray.NewRay(tuple.NewPoint(0, -1, -2), tuple.NewVector(0, 0, 1)),
			expected: Intersections{},
		},
		// ray strikes the triangle
		{
			ray: ray.NewRay(tuple.NewPoint(0, 0.5, -2), tuple.NewVector(0, 0, 1)),
			expected: Intersections{
				NewIntersection(2, triangle),
			},
		},
	}
	for _, test := range tests {
		testIntersection(t, triangle, test.ray, test.expected)
	}
}
