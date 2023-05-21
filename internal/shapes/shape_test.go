package shapes

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestNormalAt(t *testing.T) {
	var tests = []struct {
		shape    Shape
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			shape:    NewTestShape(),
			point:    tuple.NewPoint(1, 0, 0),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			shape:    NewTestShape(),
			point:    tuple.NewPoint(0, 1, 0),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			shape:    NewTestShape(),
			point:    tuple.NewPoint(0, 0, 1),
			expected: tuple.NewVector(0, 0, 1),
		},
		{
			shape:    NewTestShape(),
			point:    tuple.NewPoint(1, 0, 0),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			shape:    NewTestShape(),
			point:    tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
			expected: tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
		},
		{
			shape:    NewTestShape(),
			point:    tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
			expected: tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0).Normalize(),
		},
	}

	for _, test := range tests {
		if result := NormalAt(test.point, test.shape, Intersection{}); !result.Equal(test.expected) {
			t.Errorf("test normal:\n%s \n point: %s. \nresult: \n%s. \nexpected: \n%s", test.shape, test.point, result, test.expected)
		}
	}

	// Computing the normal on a translated shape
	shape := NewTestShape()
	shape.SetTransform(matrix.Translation(0, 1, 0))
	point := tuple.NewPoint(0, 1.70711, -0.70711)
	expected := tuple.NewVector(0, 0.7071067811865475, -0.7071067811865476)

	if got := NormalAt(point, shape, Intersection{}); !got.Equal(expected) {
		t.Errorf("test normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", shape, point, got, expected)
	}

	// Computing the normal on a transformed shape
	shape = NewTestShape()
	transform := matrix.Multiply(matrix.Scaling(1, 0.5, 1), matrix.RotationZ(math.Pi/5.0))
	shape.SetTransform(transform)
	point = tuple.NewPoint(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0)
	expected = tuple.NewVector(0, 0.9701425001453319, -0.24253562503633294)

	if got := NormalAt(point, shape, Intersection{}); !got.Equal(expected) {
		t.Errorf("test normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", shape, point, got, expected)
	}

	// Finding the normal on a child object
	g1 := NewGroup()
	g1.SetTransform(matrix.RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(matrix.Scaling(1, 2, 3))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))
	g2.AddChild(s)
	result := NormalAt(tuple.NewPoint(1.7321, 1.1547, -5.5774), s, Intersection{})
	expected = tuple.NewVector(0.28570368184140726, 0.42854315178114105, -0.8571605294481017)
	if !result.Equal(expected) {
		t.Errorf("incorrect point convertion to object space.\nexpected: %s\nresult: %s", expected, result)
	}
}

func TestColorAt(t *testing.T) {
	// Stripes with an object transformation
	shape := NewTestShape()
	shape.SetTransform(matrix.Scaling(2, 2, 2))
	mat := shape.Material()

	mat.SetPattern(materials.NewPattern(materials.Stripe, color.White(), color.Black()))

	point := tuple.NewPoint(1.5, 0, 0)
	expected := color.White()

	if result := ColorAt(point, shape); !result.Equal(expected) {
		t.Errorf("test color for:\n%s \n at point: %s. \nresult: \n%s. \nexpected: \n%s", shape, point, result, expected)
	}

	// Stripes with a pattern transformation
	shape = NewTestShape()
	mat = shape.Material()
	mat.SetPattern(materials.NewPattern(materials.Stripe, color.White(), color.Black()))
	mat.SetTransform(matrix.Scaling(2, 2, 2))

	point = tuple.NewPoint(1.5, 0, 0)
	expected = color.White()

	if result := ColorAt(point, shape); !result.Equal(expected) {
		t.Errorf("test color for:\n%s \n at point: %s. \nresult: \n%s. \nexpected: \n%s", shape, point, result, expected)
	}

	// Stripes with both an object and a pattern transformation
	shape = NewTestShape()
	shape.SetTransform(matrix.Scaling(2, 2, 2))
	mat = shape.Material()
	mat.SetPattern(materials.NewPattern(materials.Stripe, color.White(), color.Black()))
	mat.SetTransform(matrix.Translation(0.5, 0, 0))

	point = tuple.NewPoint(2.5, 0, 0)
	expected = color.White()

	if result := ColorAt(point, shape); !result.Equal(expected) {
		t.Errorf("test color for:\n%s \n at point: %s. \nresult: \n%s. \nexpected: \n%s", shape, point, result, expected)
	}
}

func TestSceneToObject(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(matrix.RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(matrix.Scaling(2, 2, 2))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))
	g2.AddChild(s)
	result := sceneToObject(tuple.NewPoint(-2, 0, -10), s)
	expected := tuple.NewPoint(0, 0, -1)
	if !result.Equal(expected) {
		t.Errorf("incorrect point convertion to object space.\nexpected: %s\nresult: %s", expected, result)
	}
}

func TestNormalToScene(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(matrix.RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(matrix.Scaling(1, 2, 3))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))
	g2.AddChild(s)
	result := normalToScene(tuple.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3), s)
	expected := tuple.NewVector(0.28571428571428575, 0.42857142857142855, -0.8571428571428571)
	if !result.Equal(expected) {
		t.Errorf("incorrect point convertion to object space.\nexpected: %s\nresult: %s", expected, result)
	}
}

func TestIntersect(t *testing.T) {
	sphere := NewSphere()
	plane := NewPlane()
	cube := NewCube()
	var tests = []struct {
		ray      *ray.Ray
		s        Shape
		expected Intersections
	}{
		{
			ray: ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1)),
			s:   sphere,
			expected: Intersections{
				NewIntersection(4.0, sphere),
				NewIntersection(6.0, sphere),
			},
		},
		{
			ray: ray.NewRay(tuple.NewPoint(0, 1, -5), tuple.NewVector(0, 0, 1)),
			s:   sphere,
			expected: Intersections{
				NewIntersection(5.0, sphere),
				NewIntersection(5.0, sphere),
			},
		},
		{
			ray:      ray.NewRay(tuple.NewPoint(0, 2, -5), tuple.NewVector(0, 0, 1)),
			s:        sphere,
			expected: Intersections{},
		},
		{
			ray: ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1)),
			s:   sphere,
			expected: Intersections{
				NewIntersection(-1.0, sphere),
				NewIntersection(1.0, sphere),
			},
		},
		{
			ray: ray.NewRay(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1)),
			s:   sphere,
			expected: Intersections{
				NewIntersection(-6.0, sphere),
				NewIntersection(-4.0, sphere),
			},
		},
		{
			// Intersect with a ray parallel to the plane
			ray:      ray.NewRay(tuple.NewPoint(0, 10, 0), tuple.NewVector(0, 0, 1)),
			s:        plane,
			expected: Intersections{},
		},
		{
			// Intersect with a coplanar ray
			ray:      ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1)),
			s:        plane,
			expected: Intersections{},
		},
		{
			// A ray intersecting a plane from above
			ray: ray.NewRay(tuple.NewPoint(0, 1, 0), tuple.NewVector(0, -1, 0)),
			s:   plane,
			expected: Intersections{
				NewIntersection(1, plane),
			},
		},
		{
			// A ray intersecting a plane from below
			ray: ray.NewRay(tuple.NewPoint(0, -1, 0), tuple.NewVector(0, 1, 0)),
			s:   plane,
			expected: Intersections{
				NewIntersection(1, plane),
			},
		},
		{
			// +x
			ray: ray.NewRay(tuple.NewPoint(5, 0.5, 0), tuple.NewVector(-1, 0, 0)),
			s:   cube,
			expected: Intersections{
				NewIntersection(4, cube),
				NewIntersection(6, cube),
			},
		},
		{
			// -x
			ray: ray.NewRay(tuple.NewPoint(-5, 0.5, 0), tuple.NewVector(1, 0, 0)),
			s:   cube,
			expected: Intersections{
				NewIntersection(4, cube),
				NewIntersection(6, cube),
			},
		},
		{
			// +y
			ray: ray.NewRay(tuple.NewPoint(0.5, 5, 0), tuple.NewVector(0, -1, 0)),
			s:   cube,
			expected: Intersections{
				NewIntersection(4, cube),
				NewIntersection(6, cube),
			},
		},
		{
			// -y
			ray: ray.NewRay(tuple.NewPoint(0.5, -5, 0), tuple.NewVector(0, 1, 0)),
			s:   cube,
			expected: Intersections{
				NewIntersection(4, cube),
				NewIntersection(6, cube),
			},
		},
		{
			// +z
			ray: ray.NewRay(tuple.NewPoint(0.5, 0, 5), tuple.NewVector(0, 0, -1)),
			s:   cube,
			expected: Intersections{
				NewIntersection(4, cube),
				NewIntersection(6, cube),
			},
		},
		{
			// -z
			ray: ray.NewRay(tuple.NewPoint(0.5, 0, -5), tuple.NewVector(0, 0, 1)),
			s:   cube,
			expected: Intersections{
				NewIntersection(4, cube),
				NewIntersection(6, cube),
			},
		},
		{
			// inside
			ray: ray.NewRay(tuple.NewPoint(0, 0.5, 0), tuple.NewVector(0, 0, 1)),
			s:   cube,
			expected: Intersections{
				NewIntersection(-1, cube),
				NewIntersection(1, cube),
			},
		},
		// cube misses
		{
			ray:      ray.NewRay(tuple.NewPoint(-2, 0, 0), tuple.NewVector(0.2673, 0.5345, 0.8018)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      ray.NewRay(tuple.NewPoint(0, -2, 0), tuple.NewVector(0.8018, 0.2673, 0.5345)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      ray.NewRay(tuple.NewPoint(0, 0, -2), tuple.NewVector(0.5345, 0.8018, 0.2673)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      ray.NewRay(tuple.NewPoint(2, 0, 2), tuple.NewVector(0, 0, -1)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      ray.NewRay(tuple.NewPoint(0, 2, 2), tuple.NewVector(0, -1, 0)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      ray.NewRay(tuple.NewPoint(2, 2, 0), tuple.NewVector(-1, 0, 0)),
			s:        cube,
			expected: Intersections{},
		},
	}

	for _, test := range tests {
		testIntersection(t, test.s, test.ray, test.expected)
	}
}

func testIntersection(t *testing.T, s Shape, r *ray.Ray, expected Intersections) {
	if result := Intersect(s, r); len(result) != len(expected) {
		t.Errorf("incorrect number of intersections. Result: %d. Expected: %d", len(result), len(expected))
	} else {
		for i := range result {
			if !utils.FloatEquals(result[i].t, expected[i].t) {
				t.Errorf("incorrect t of intersect:\n%s \n \nresult: \n%f. \nexpected: \n%f", r, result[i].t, expected[i].t)
			}
			if result[i].shape != expected[i].shape {
				t.Errorf("incorrect Shape of intersect:\n%s \n \nresult: \n%s. \nexpected: \n%s", r, result[i].shape, expected[i].shape)
			}
		}
	}
}
