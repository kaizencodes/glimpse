package ray

import (
	"glimpse/calc"
	"glimpse/matrix"
	"glimpse/shapes"
	"glimpse/tuple"
	"math"
	"testing"
)

func TestPosition(t *testing.T) {
	var tests = []struct {
		ray  *Ray
		dist float64
		want tuple.Tuple
	}{
		{
			ray:  New(tuple.NewPoint(2, 3, 4), tuple.NewVector(1, 0, 0)),
			dist: 0,
			want: tuple.NewPoint(2, 3, 4),
		},
		{
			ray:  New(tuple.NewPoint(2, 3, 4), tuple.NewVector(1, 0, 0)),
			dist: 1,
			want: tuple.NewPoint(3, 3, 4),
		},
		{
			ray:  New(tuple.NewPoint(2, 3, 4), tuple.NewVector(0, 1, 0)),
			dist: -1,
			want: tuple.NewPoint(2, 2, 4),
		},
		{
			ray:  New(tuple.NewPoint(2, 3, 4), tuple.NewVector(0, 0, 1)),
			dist: 2.5,
			want: tuple.NewPoint(2, 3, 6.5),
		},
	}

	for _, test := range tests {
		if got := test.ray.Position(test.dist); !got.Equal(test.want) {
			t.Errorf("ray position:\n%s \n dist: %f. \ngot: \n%s. \nexpected: \n%s", test.ray, test.dist, got, test.want)
		}
	}
}

func TestIntersect(t *testing.T) {
	sphere := shapes.NewSphere()
	plane := shapes.NewPlane()
	cube := shapes.NewCube()
	var tests = []struct {
		ray      *Ray
		s        shapes.Shape
		expected Intersections
	}{
		{
			ray: New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1)),
			s:   sphere,
			expected: Intersections{
				Intersection{t: 4.0, shape: sphere},
				Intersection{t: 6.0, shape: sphere},
			},
		},
		{
			ray: New(tuple.NewPoint(0, 1, -5), tuple.NewVector(0, 0, 1)),
			s:   sphere,
			expected: Intersections{
				Intersection{t: 5.0, shape: sphere},
				Intersection{t: 5.0, shape: sphere},
			},
		},
		{
			ray:      New(tuple.NewPoint(0, 2, -5), tuple.NewVector(0, 0, 1)),
			s:        sphere,
			expected: Intersections{},
		},
		{
			ray: New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1)),
			s:   sphere,
			expected: Intersections{
				Intersection{t: -1.0, shape: sphere},
				Intersection{t: 1.0, shape: sphere},
			},
		},
		{
			ray: New(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1)),
			s:   sphere,
			expected: Intersections{
				Intersection{t: -6.0, shape: sphere},
				Intersection{t: -4.0, shape: sphere},
			},
		},
		{
			// Intersect with a ray parallel to the plane
			ray:      New(tuple.NewPoint(0, 10, 0), tuple.NewVector(0, 0, 1)),
			s:        plane,
			expected: Intersections{},
		},
		{
			// Intersect with a coplanar ray
			ray:      New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1)),
			s:        plane,
			expected: Intersections{},
		},
		{
			// A ray intersecting a plane from above
			ray: New(tuple.NewPoint(0, 1, 0), tuple.NewVector(0, -1, 0)),
			s:   plane,
			expected: Intersections{
				Intersection{t: 1, shape: plane},
			},
		},
		{
			// A ray intersecting a plane from below
			ray: New(tuple.NewPoint(0, -1, 0), tuple.NewVector(0, 1, 0)),
			s:   plane,
			expected: Intersections{
				Intersection{t: 1, shape: plane},
			},
		},
		{
			// +x
			ray: New(tuple.NewPoint(5, 0.5, 0), tuple.NewVector(-1, 0, 0)),
			s:   cube,
			expected: Intersections{
				Intersection{t: 4, shape: cube},
				Intersection{t: 6, shape: cube},
			},
		},
		{
			// -x
			ray: New(tuple.NewPoint(-5, 0.5, 0), tuple.NewVector(1, 0, 0)),
			s:   cube,
			expected: Intersections{
				Intersection{t: 4, shape: cube},
				Intersection{t: 6, shape: cube},
			},
		},
		{
			// +y
			ray: New(tuple.NewPoint(0.5, 5, 0), tuple.NewVector(0, -1, 0)),
			s:   cube,
			expected: Intersections{
				Intersection{t: 4, shape: cube},
				Intersection{t: 6, shape: cube},
			},
		},
		{
			// -y
			ray: New(tuple.NewPoint(0.5, -5, 0), tuple.NewVector(0, 1, 0)),
			s:   cube,
			expected: Intersections{
				Intersection{t: 4, shape: cube},
				Intersection{t: 6, shape: cube},
			},
		},
		{
			// +z
			ray: New(tuple.NewPoint(0.5, 0, 5), tuple.NewVector(0, 0, -1)),
			s:   cube,
			expected: Intersections{
				Intersection{t: 4, shape: cube},
				Intersection{t: 6, shape: cube},
			},
		},
		{
			// -z
			ray: New(tuple.NewPoint(0.5, 0, -5), tuple.NewVector(0, 0, 1)),
			s:   cube,
			expected: Intersections{
				Intersection{t: 4, shape: cube},
				Intersection{t: 6, shape: cube},
			},
		},
		{
			// inside
			ray: New(tuple.NewPoint(0, 0.5, 0), tuple.NewVector(0, 0, 1)),
			s:   cube,
			expected: Intersections{
				Intersection{t: -1, shape: cube},
				Intersection{t: 1, shape: cube},
			},
		},
		// cube misses
		{
			ray:      New(tuple.NewPoint(-2, 0, 0), tuple.NewVector(0.2673, 0.5345, 0.8018)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(0, -2, 0), tuple.NewVector(0.8018, 0.2673, 0.5345)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(0, 0, -2), tuple.NewVector(0.5345, 0.8018, 0.2673)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(2, 0, 2), tuple.NewVector(0, 0, -1)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(0, 2, 2), tuple.NewVector(0, -1, 0)),
			s:        cube,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(2, 2, 0), tuple.NewVector(-1, 0, 0)),
			s:        cube,
			expected: Intersections{},
		},
	}

	for _, test := range tests {
		testIntersection(t, test.s, test.ray, test.expected)
	}
}

func TestCylinderIntersect(t *testing.T) {
	cylinder := shapes.NewCylinder()
	var tests = []struct {
		ray      *Ray
		s        shapes.Shape
		expected Intersections
	}{
		// cylinder misses
		{
			ray:      New(tuple.NewPoint(1, 0, 0), tuple.NewVector(0, 1, 0).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 1, 0).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(0, 0, -5), tuple.NewVector(1, 1, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		// A ray strikes a cylinder
		{
			ray: New(tuple.NewPoint(1, 0, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:   cylinder,
			expected: Intersections{
				Intersection{t: 5, shape: cylinder},
				Intersection{t: 5, shape: cylinder},
			},
		},
		{
			ray: New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:   cylinder,
			expected: Intersections{
				Intersection{t: 4, shape: cylinder},
				Intersection{t: 6, shape: cylinder},
			},
		},
		{
			ray: New(tuple.NewPoint(0.5, 0, -5), tuple.NewVector(0.1, 1, 1).Normalize()),
			s:   cylinder,
			expected: Intersections{
				Intersection{t: 6.807981917027314, shape: cylinder},
				Intersection{t: 7.088723439378867, shape: cylinder},
			},
		},
	}

	for _, test := range tests {
		testIntersection(t, test.s, test.ray, test.expected)
	}
}

func TestTruncatedCylinderIntersect(t *testing.T) {
	cylinder := shapes.NewCylinder()
	cylinder.SetMinimum(1)
	cylinder.SetMaximum(2)

	var tests = []struct {
		ray      *Ray
		s        shapes.Shape
		expected Intersections
	}{
		// cylinder misses
		{
			ray:      New(tuple.NewPoint(0, 1.5, 0), tuple.NewVector(0.1, 1, 0).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(0, 3, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(0, 2, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		{
			ray:      New(tuple.NewPoint(0, 1, -5), tuple.NewVector(0, 0, 1).Normalize()),
			s:        cylinder,
			expected: Intersections{},
		},
		// A ray strikes a cylinder
		{
			ray: New(tuple.NewPoint(0, 1.5, -2), tuple.NewVector(0, 0, 1).Normalize()),
			s:   cylinder,
			expected: Intersections{
				Intersection{t: 1, shape: cylinder},
				Intersection{t: 3, shape: cylinder},
			},
		},
	}

	for _, test := range tests {
		testIntersection(t, test.s, test.ray, test.expected)
	}
}

func TestClosedCylinderIntersect(t *testing.T) {
	cylinder := shapes.NewCylinder()
	cylinder.SetMinimum(1)
	cylinder.SetMaximum(2)
	cylinder.SetClosed(true)

	var tests = []struct {
		ray      *Ray
		s        shapes.Shape
		expected int
	}{
		{
			ray:      New(tuple.NewPoint(0, 3, 0), tuple.NewVector(0, -1, 0).Normalize()),
			s:        cylinder,
			expected: 2,
		},
		{
			ray:      New(tuple.NewPoint(0, 3, -2), tuple.NewVector(0, -1, 2).Normalize()),
			s:        cylinder,
			expected: 2,
		},
		{
			ray:      New(tuple.NewPoint(0, 4, -2), tuple.NewVector(0, -1, 1).Normalize()),
			s:        cylinder,
			expected: 2,
		},
		{
			ray:      New(tuple.NewPoint(0, 0, -2), tuple.NewVector(0, 1, 2).Normalize()),
			s:        cylinder,
			expected: 2,
		},
		{
			ray:      New(tuple.NewPoint(0, -1, -2), tuple.NewVector(0, 1, 1).Normalize()),
			s:        cylinder,
			expected: 2,
		},
	}

	for _, test := range tests {
		result := test.ray.Intersect(test.s)
		if len(result) != test.expected {
			t.Errorf("incorrect number of intersections. Result: %d. Expected: %d", len(result), test.expected)
		}
	}
}

func TestEmptyGroupIntersect(t *testing.T) {
	g := shapes.NewGroup()
	r := New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	expected := Intersections{}

	testIntersection(t, g, r, expected)
}

func TestGroupIntersect(t *testing.T) {
	g := shapes.NewGroup()
	s1 := shapes.NewSphere()
	s2 := shapes.NewSphere()
	s2.SetTransform(matrix.Translation(0, 0, -3))
	s3 := shapes.NewSphere()
	s3.SetTransform(matrix.Translation(5, 0, 0))
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)
	r := New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	expected := Intersections{
		Intersection{t: 1, shape: s2},
		Intersection{t: 3, shape: s2},
		Intersection{t: 4, shape: s1},
		Intersection{t: 6, shape: s1},
	}

	testIntersection(t, g, r, expected)
}

func TestTriangleIntersect(t *testing.T) {
	triangle := shapes.NewTriangle(
		tuple.NewPoint(0, 1, 0),
		tuple.NewPoint(-1, 0, 0),
		tuple.NewPoint(1, 0, 0),
	)

	var tests = []struct {
		ray      *Ray
		expected Intersections
	}{
		// Intersecting a ray parallel to the triangle
		{
			ray:      New(tuple.NewPoint(0, -1, -2), tuple.NewVector(0, 1, 0)),
			expected: Intersections{},
		},
		// ray misses the p1-p3 edge
		{
			ray:      New(tuple.NewPoint(1, 1, -2), tuple.NewVector(0, 0, 1)),
			expected: Intersections{},
		},
		// ray misses the p1-p2 edge
		{
			ray:      New(tuple.NewPoint(-1, 1, -2), tuple.NewVector(0, 0, 1)),
			expected: Intersections{},
		},
		// ray misses the p2-p3 edge
		{
			ray:      New(tuple.NewPoint(0, -1, -2), tuple.NewVector(0, 0, 1)),
			expected: Intersections{},
		},
		// ray strikes the triangle
		{
			ray: New(tuple.NewPoint(0, 0.5, -2), tuple.NewVector(0, 0, 1)),
			expected: Intersections{
				Intersection{t: 2, shape: triangle},
			},
		},
	}
	for _, test := range tests {
		testIntersection(t, triangle, test.ray, test.expected)
	}
}

func TestGroupTransformation(t *testing.T) {
	g := shapes.NewGroup()
	g.SetTransform(matrix.Scaling(2, 2, 2))
	s := shapes.NewSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))
	g.AddChild(s)
	r := New(tuple.NewPoint(10, 0, -10), tuple.NewVector(0, 0, 1))
	if len(r.Intersect(g)) != 2 {
		t.Errorf("incorrect transformation")
	}
}

func testIntersection(t *testing.T, s shapes.Shape, r *Ray, expected Intersections) {
	result := r.Intersect(s)
	if len(result) != len(expected) {
		t.Errorf("incorrect number of intersections. Result: %d. Expected: %d", len(result), len(expected))
	} else {
		for i := range result {
			if !calc.FloatEquals(result[i].t, expected[i].t) {
				t.Errorf("incorrect t of intersect:\n%s \n \nresult: \n%f. \nexpected: \n%f", r, result[i].t, expected[i].t)
			}
			if result[i].shape != expected[i].shape {
				t.Errorf("incorrect Shape of intersect:\n%s \n \nresult: \n%s. \nexpected: \n%s", r, result[i].shape, expected[i].shape)
			}
		}
	}
}

func TestHit(t *testing.T) {
	shape := shapes.Shape(shapes.NewSphere())
	var tests = []struct {
		collection Intersections
		want       Intersection
	}{
		{
			collection: Intersections{
				Intersection{t: 1.0, shape: shape},
				Intersection{t: 2.0, shape: shape},
			},
			want: Intersection{t: 1.0, shape: shape},
		},
		{
			collection: Intersections{
				Intersection{t: -1.0, shape: shape},
				Intersection{t: 1.0, shape: shape},
			},
			want: Intersection{t: 1.0, shape: shape},
		},
		{
			collection: Intersections{
				Intersection{t: -2.0, shape: shape},
				Intersection{t: -1.0, shape: shape},
			},
			want: Intersection{t: math.MaxFloat64},
		},
		{
			collection: Intersections{
				Intersection{t: 5.0, shape: shape},
				Intersection{t: 7.0, shape: shape},
				Intersection{t: -3.0, shape: shape},
				Intersection{t: 2.0, shape: shape},
			},
			want: Intersection{t: 2.0, shape: shape},
		},
	}

	for _, test := range tests {
		if got := test.collection.Hit(); got.t != test.want.t {
			t.Errorf("Hit: collection\n%s \ngot: \n%f. \nexpected: \n%f", test.collection, got.t, test.want.t)
		}
	}
}

func TestTranslate(t *testing.T) {
	r := New(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
	want := New(tuple.NewPoint(4, 6, 8), tuple.NewVector(0, 1, 0))
	x, y, z := 3.0, 4.0, 5.0

	if got := r.Translate(x, y, z); !got.Equal(want) {
		t.Errorf("translation(%f, %f, %f),\nray:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, r, got, want)
	}

	x, y, z = 2.0, 3.0, 4.0
	want = New(tuple.NewPoint(2, 6, 12), tuple.NewVector(0, 3, 0))

	if got := r.Scale(x, y, z); !got.Equal(want) {
		t.Errorf("scale(%f, %f, %f),\nray:\n%s\n\ngot:\n%s", x, y, z, r, got)
	}
}

func TestScale(t *testing.T) {
	r := New(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
	want := New(tuple.NewPoint(2, 6, 12), tuple.NewVector(0, 3, 0))
	x, y, z := 2.0, 3.0, 4.0

	if got := r.Scale(x, y, z); !got.Equal(want) {
		t.Errorf("scale(%f, %f, %f),\nray:\n%s\n\ngot:\n%s", x, y, z, r, got)
	}
}

func TestSphereTransformations(t *testing.T) {
	r := New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	sphere := shapes.NewSphere()
	sphere.SetTransform(matrix.Scaling(2, 2, 2))
	want := Intersections{
		Intersection{t: 3.0, shape: sphere},
		Intersection{t: 7.0, shape: sphere},
	}

	got := r.Intersect(sphere)
	for i := range got {
		if got[i].t != want[i].t {
			t.Errorf("incorrect t of intersect:\n%s \n \ngot: \n%f. \nexpected: \n%f", r, got[i].t, want[i].t)
		}
		if got[i].shape != want[i].shape {
			t.Errorf("incorrect Shape of intersect:\n%s \n \ngot: \n%s. \nexpected: \n%s", r, got[i].shape, want[i].shape)
		}
	}
}
