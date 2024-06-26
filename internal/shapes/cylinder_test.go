package shapes

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestCylinderLocalNormalAt(t *testing.T) {
	cylinder := NewCylinder()
	var tests = []struct {
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			point:    tuple.NewPoint(1, 0, 0),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			point:    tuple.NewPoint(0, 5, -1),
			expected: tuple.NewVector(0, 0, -1),
		},
		{
			point:    tuple.NewPoint(0, -2, 1),
			expected: tuple.NewVector(0, 0, 1),
		},
		{
			point:    tuple.NewPoint(-1, 1, 0),
			expected: tuple.NewVector(-1, 0, 0),
		},
	}

	for _, test := range tests {
		if result := cylinder.localNormalAt(test.point, Intersection{}); !result.Equal(test.expected) {
			t.Errorf("Cylinder normal: \nresult: \n%s. \nexpected: \n%s", result, test.expected)
		}
	}
}

func TestClosedCylinderLocalNormalAt(t *testing.T) {
	cylinder := NewCylinder()
	cylinder.Minimum = 1
	cylinder.Maximum = 2
	cylinder.Closed = true

	var tests = []struct {
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			point:    tuple.NewPoint(0, 1, 0),
			expected: tuple.NewVector(0, -1, 0),
		},
		{
			point:    tuple.NewPoint(0.5, 1, 0),
			expected: tuple.NewVector(0, -1, 0),
		},
		{
			point:    tuple.NewPoint(0, 1, 0.5),
			expected: tuple.NewVector(0, -1, 0),
		},
		{
			point:    tuple.NewPoint(0, 2, 0),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			point:    tuple.NewPoint(0.5, 2, 0),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			point:    tuple.NewPoint(0, 2, 0.5),
			expected: tuple.NewVector(0, 1, 0),
		},
	}

	for _, test := range tests {
		if result := cylinder.localNormalAt(test.point, Intersection{}); !result.Equal(test.expected) {
			t.Errorf("Cylinder normal: \nresult: \n%s. \nexpected: \n%s", result, test.expected)
		}
	}
}

func TestLocalIntersect(t *testing.T) {
	cylinder := NewCylinder()
	var tests = []struct {
		ray      *ray.Ray
		s        Shape
		expected Intersections
	}{
		// cylinder misses
		{
			ray:      ray.New(tuple.NewPoint(1, 0, 0), tuple.NewVector(0, 1, 0).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 1, 0).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(1, 1, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		// A ray strikes a cylinder
		{
			ray: ray.New(tuple.NewPoint(1, 0, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:   cylinder,
			expected: Intersections{
				NewIntersection(5, cylinder),
				NewIntersection(5, cylinder),
			},
		},
		{
			ray: ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:   cylinder,
			expected: Intersections{
				NewIntersection(4, cylinder),
				NewIntersection(6, cylinder),
			},
		},
		{
			ray: ray.New(tuple.NewPoint(0.5, 0, -5), tuple.NewVector(0.1, 1, 1).Normalize()),
			s:   cylinder,
			expected: Intersections{
				NewIntersection(6.807981917027314, cylinder),
				NewIntersection(7.088723439378867, cylinder),
			},
		},
	}

	for _, test := range tests {
		testIntersection(t, test.s, test.ray, test.expected)
	}
}

func TestTruncatedCylinderIntersect(t *testing.T) {
	cylinder := NewCylinder()
	cylinder.Minimum = 1
	cylinder.Maximum = 2

	var tests = []struct {
		ray      *ray.Ray
		s        Shape
		expected Intersections
	}{
		// cylinder misses
		{
			ray:      ray.New(tuple.NewPoint(0, 1.5, 0), tuple.NewVector(0.1, 1, 0).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      ray.New(tuple.NewPoint(0, 3, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      ray.New(tuple.NewPoint(0, 2, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      ray.New(tuple.NewPoint(0, 1, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		// A ray strikes a cylinder
		{
			ray: ray.New(tuple.NewPoint(0, 1.5, -2), tuple.NewVector(0, 0, 1).Normalize()),
			s:   cylinder,
			expected: Intersections{
				NewIntersection(1, cylinder),
				NewIntersection(3, cylinder),
			},
		},
	}

	for _, test := range tests {
		testIntersection(t, test.s, test.ray, test.expected)
	}
}

func TestClosedCylinderIntersect(t *testing.T) {
	cylinder := NewCylinder()
	cylinder.Minimum = 1
	cylinder.Maximum = 2
	cylinder.Closed = true

	var tests = []struct {
		ray      *ray.Ray
		s        Shape
		expected int
	}{
		{
			ray:      ray.New(tuple.NewPoint(0, 3, 0), tuple.NewVector(0, -1, 0).Normalize()),
			s:        cylinder,
			expected: 2,
		},
		{
			ray:      ray.New(tuple.NewPoint(0, 3, -2), tuple.NewVector(0, -1, 2).Normalize()),
			s:        cylinder,
			expected: 2,
		},
		{
			ray:      ray.New(tuple.NewPoint(0, 4, -2), tuple.NewVector(0, -1, 1).Normalize()),
			s:        cylinder,
			expected: 2,
		},
		{
			ray:      ray.New(tuple.NewPoint(0, 0, -2), tuple.NewVector(0, 1, 2).Normalize()),
			s:        cylinder,
			expected: 2,
		},
		{
			ray:      ray.New(tuple.NewPoint(0, -1, -2), tuple.NewVector(0, 1, 1).Normalize()),
			s:        cylinder,
			expected: 2,
		},
	}

	for _, test := range tests {
		result := test.s.localIntersect(test.ray)
		if len(result) != test.expected {
			t.Errorf("incorrect number of intersections. Result: %d. Expected: %d", len(result), test.expected)
		}
	}
}

func TestBoundingBoxForCylinder(t *testing.T) {
	//  A cylinder has a bounding box
	c := NewCylinder()
	c.CalculateBoundingBox()
	expected := NewBoundingBox(tuple.NewPoint(-1, math.Inf(-1), -1), tuple.NewPoint(1, math.Inf(1), 1))

	for _, diff := range utils.Compare(c.BoundingBox(), expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundingBoxForLimitedCylinder(t *testing.T) {
	//  A cylinder has a bounding box
	c := NewCylinder()
	c.Minimum = -5
	c.Maximum = 3
	c.CalculateBoundingBox()
	box := c.BoundingBox()
	expected := NewBoundingBox(tuple.NewPoint(-1, -5, -1), tuple.NewPoint(1, 3, 1))

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}
