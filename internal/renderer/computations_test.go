package renderer

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestPrepareComputations(t *testing.T) {
	// The hit, when an intersection occurs on the outside
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := shapes.NewSphere()
	i := shapes.NewIntersection(4, s)
	comps := prepareComputations(i, r, shapes.Intersections{i})
	point := tuple.NewPoint(0, 0, -1)
	eyeV := tuple.NewVector(0, 0, -1)
	normalV := tuple.NewVector(0, 0, -1)
	inside := false

	testComputation(t, comps, s, i, point, eyeV, normalV, inside)
}

func TestPrepareComputations2(t *testing.T) {
	// The hit, when an intersection occurs on the inside
	r := ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	s := shapes.NewSphere()
	i := shapes.NewIntersection(1, s)
	comps := prepareComputations(i, r, shapes.Intersections{i})
	point := tuple.NewPoint(0, 0, 1)
	eyeV := tuple.NewVector(0, 0, -1)
	normalV := tuple.NewVector(0, 0, -1)
	inside := true

	testComputation(t, comps, s, i, point, eyeV, normalV, inside)
}

func TestPrepareComputations3(t *testing.T) {
	// The hit should offset the point
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := shapes.NewSphere()
	s.SetTransform(matrix.Translation(0, 0, 1))
	i := shapes.NewIntersection(5, s)
	comps := prepareComputations(i, r, shapes.Intersections{i})
	point := tuple.NewPoint(0, 0, 0)
	eyeV := tuple.NewVector(0, 0, -1)
	normalV := tuple.NewVector(0, 0, -1)
	inside := false

	testComputation(t, comps, s, i, point, eyeV, normalV, inside)
	if comps.OverPoint.Z > -utils.EPSILON/2.0 {
		t.Errorf("incorrect OverPoint.Z %f > %f", comps.OverPoint.Z, -utils.EPSILON/2)
	}

	if comps.Point.Z < comps.OverPoint.Z {
		t.Errorf("incorrect Z %f < OverPoint.Z %f", comps.Point.Z, comps.OverPoint.Z)
	}
}

func TestPrepareComputations4(t *testing.T) {
	// The under point is offset below the surface
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := shapes.NewGlassSphere()
	s.SetTransform(matrix.Translation(0, 0, 1))
	i := shapes.NewIntersection(5, s)
	comps := prepareComputations(i, r, shapes.Intersections{i})
	eps := utils.EPSILON / 2.0
	if comps.UnderPoint.Z < eps {
		t.Errorf("incorrect UnderPoint.Z %f < %f", comps.UnderPoint.Z, utils.EPSILON/2)
	}

	if comps.Point.Z > comps.UnderPoint.Z {
		t.Errorf("incorrect Z %f > UnderPoint.Z %f", comps.Point.Z, comps.UnderPoint.Z)
	}
}

func TestPrepareComputations5(t *testing.T) {
	// Precomputing the reflection vector

	r := ray.New(tuple.NewPoint(0, 1, -1), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	p := shapes.NewPlane()
	i := shapes.NewIntersection(math.Sqrt(2), p)
	comps := prepareComputations(i, r, shapes.Intersections{i})
	reflectV := tuple.NewVector(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	if comps.ReflectV != reflectV {
		t.Errorf("incorrect reflection vector, expected %f, got: %f", reflectV, comps.ReflectV)
	}

	// Preparing the normal on a smooth triangle
	triangle := shapes.NewSmoothTriangle(
		tuple.NewPoint(0, 1, 0),
		tuple.NewPoint(-1, 0, 0),
		tuple.NewPoint(1, 0, 0),
		tuple.NewVector(0, 1, 0),
		tuple.NewVector(-1, 0, 0),
		tuple.NewVector(1, 0, 0),
	)
	triangle.Model = shapes.NewGroup()
	r = ray.New(tuple.NewPoint(-0.2, 0.3, -2), tuple.NewVector(0, 0, 1))
	hit := shapes.NewIntersectionWithUV(1, 0.45, 0.25, triangle)
	xs := shapes.Intersections{hit}
	result := prepareComputations(hit, r, xs).NormalV
	expected := tuple.NewVector(-0.5547001962252291, 0.8320502943378437, 0)
	if result != expected {
		t.Errorf("hit not passed to shape NormalAt")
	}
}

func testComputation(t *testing.T, comps Computations, shape shapes.Shape, i shapes.Intersection, point, eyeV, normalV tuple.Tuple, inside bool) {
	if comps.T != i.T() {
		t.Errorf("incorrect T, expected %f, got: %f", i.T(), comps.T)
	}

	if comps.Shape != shape {
		t.Errorf("incorrect T, expected %s, got: %s", shape, comps.Shape)
	}

	if comps.Point != point {
		t.Errorf("incorrect point, expected %s, got: %s", point, comps.Point)
	}

	if comps.EyeV != eyeV {
		t.Errorf("incorrect eyeV, expected %s, got: %s", eyeV, comps.EyeV)
	}

	if comps.NormalV != normalV {
		t.Errorf("incorrect normalV, expected %s, got: %s", normalV, comps.NormalV)
	}

	if comps.Inside != inside {
		t.Errorf("incorrect inside, expected %t, got: %t", inside, comps.Inside)
	}
}

func TestSchlick(t *testing.T) {
	// The schlick approximation under total internal reflection
	sphere := shapes.NewGlassSphere()
	r := ray.New(tuple.NewPoint(0, 0, math.Sqrt(2)/2), tuple.NewVector(0, 1, 0))
	xs := shapes.Intersections{
		shapes.NewIntersection(-math.Sqrt(2)/2, sphere),
		shapes.NewIntersection(math.Sqrt(2)/2, sphere),
	}
	comps := prepareComputations(xs[1], r, xs)
	result := comps.schlick()
	expected := 1.0

	if !utils.FloatEquals(result, expected) {
		t.Errorf("incorrect reflectance:\nresult: \n%f. \nexpected: \n%f", result, expected)
	}
}

func TestSchlick2(t *testing.T) {
	// The schlick approximation with a perpendicular viewing angle
	sphere := shapes.NewGlassSphere()
	r := ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 1, 0))
	xs := shapes.Intersections{
		shapes.NewIntersection(-1, sphere),
		shapes.NewIntersection(1, sphere),
	}
	comps := prepareComputations(xs[1], r, xs)
	result := comps.schlick()
	expected := 0.04

	if !utils.FloatEquals(result, expected) {
		t.Errorf("incorrect reflectance:\nresult: \n%f. \nexpected: \n%f", result, expected)
	}
}

func TestSchlick3(t *testing.T) {
	// The schlick approximation with small angle and n2 > n1
	sphere := shapes.NewGlassSphere()
	r := ray.New(tuple.NewPoint(0, 0.99, -2), tuple.NewVector(0, 0, 1))
	xs := shapes.Intersections{
		shapes.NewIntersection(1.8589, sphere),
	}
	comps := prepareComputations(xs[0], r, xs)
	result := comps.schlick()
	expected := 0.4887308101221217

	if result != expected {
		t.Errorf("incorrect reflectance:\nresult: \n%f. \nexpected: \n%f", result, expected)
	}
}

func TestRefraction(t *testing.T) {
	// Finding N1 and N2 at various intersections
	a := shapes.NewGlassSphere()
	a.SetTransform(matrix.Scaling(2, 2, 2))
	a.Material().RefractiveIndex = 1.5

	b := shapes.NewGlassSphere()
	b.SetTransform(matrix.Translation(0, 0, -0.25))
	b.Material().RefractiveIndex = 2.0

	c := shapes.NewGlassSphere()
	c.SetTransform(matrix.Translation(0, 0, 0.25))
	c.Material().RefractiveIndex = 2.5

	r := ray.New(tuple.NewPoint(0, 0, -4), tuple.NewVector(0, 0, 1))
	xs := shapes.Intersections{
		shapes.NewIntersection(2.0, a),
		shapes.NewIntersection(2.75, b),
		shapes.NewIntersection(3.25, c),
		shapes.NewIntersection(4.75, b),
		shapes.NewIntersection(5.25, c),
		shapes.NewIntersection(6.0, a),
	}

	var tests = []struct {
		ray                    *ray.Ray
		computations           Computations
		expectedN1, expectedN2 float64
	}{
		{
			computations: prepareComputations(xs[0], r, xs),
			expectedN1:   1.0,
			expectedN2:   1.5,
		},
		{
			computations: prepareComputations(xs[1], r, xs),
			expectedN1:   1.5,
			expectedN2:   2.0,
		},
		{
			computations: prepareComputations(xs[2], r, xs),
			expectedN1:   2.0,
			expectedN2:   2.5,
		},
		{
			computations: prepareComputations(xs[3], r, xs),
			expectedN1:   2.5,
			expectedN2:   2.5,
		},
		{
			computations: prepareComputations(xs[4], r, xs),
			expectedN1:   2.5,
			expectedN2:   1.5,
		},
		{
			computations: prepareComputations(xs[5], r, xs),
			expectedN1:   1.5,
			expectedN2:   1.0,
		},
	}

	for _, test := range tests {
		if test.computations.N1 != test.expectedN1 {
			t.Errorf("incorrect n1 :\nresult: \n%f. \nexpected: \n%f", test.computations.N1, test.expectedN1)
		}

		if test.computations.N2 != test.expectedN2 {
			t.Errorf("incorrect n2 :\nresult: \n%f. \nexpected: \n%f", test.computations.N2, test.expectedN2)
		}
	}
}
