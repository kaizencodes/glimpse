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
	}

	for _, test := range tests {
		result := test.ray.Intersect(test.s)
		for i, _ := range result {
			if !calc.FloatEquals(result[i].t, test.expected[i].t) {
				t.Errorf("incorrect t of intersect:\n%s \n \nresult: \n%f. \nexpected: \n%f", test.ray, result[i].t, test.expected[i].t)
			}
			if result[i].shape != test.expected[i].shape {
				t.Errorf("incorrect Shape of intersect:\n%s \n \nresult: \n%s. \nexpected: \n%s", test.ray, result[i].shape, test.expected[i].shape)
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
	for i, _ := range got {
		if got[i].t != want[i].t {
			t.Errorf("incorrect t of intersect:\n%s \n \ngot: \n%f. \nexpected: \n%f", r, got[i].t, want[i].t)
		}
		if got[i].shape != want[i].shape {
			t.Errorf("incorrect Shape of intersect:\n%s \n \ngot: \n%s. \nexpected: \n%s", r, got[i].shape, want[i].shape)
		}
	}
}
