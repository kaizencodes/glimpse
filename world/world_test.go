package world

import (
	"glimpse/color"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/ray"
	"glimpse/shapes"
	"glimpse/tuple"
	"math"
	"testing"
)

func TestIntersect(t *testing.T) {
	w := Default()
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	sections := w.intersect(r)
	expected := []float64{4, 4.5, 5.5, 6}
	for i, v := range expected {
		if sections[i].T() != v {
			t.Errorf("incorrect t of intersect:\n%s \n \ngot: \n%f. \nexpected: \n%f", r, sections[i].T(), v)
		}
	}
}

func TestShadeHit(t *testing.T) {
	w := Default()
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	shape := w.Shapes()[0]
	i := ray.NewIntersection(4, shape)
	comps := ray.PrepareComputations(i, r, ray.Intersections{i})

	result := w.shadeHit(comps)
	expected := color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
	if result != expected {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	w = Default()
	w.SetLights([]ray.Light{
		ray.NewLight(tuple.NewPoint(0, 0.25, 0), color.New(1, 1, 1)),
	})
	r = ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	shape = w.Shapes()[1]
	i = ray.NewIntersection(0.5, shape)
	comps = ray.PrepareComputations(i, r, ray.Intersections{i})

	result = w.shadeHit(comps)
	expected = color.New(0.9049844720832575, 0.9049844720832575, 0.9049844720832575)
	if result != expected {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// multiple light sources
	w = Default()
	w.SetLights([]ray.Light{
		ray.NewLight(tuple.NewPoint(0, 0.25, 0), color.New(1, 1, 1)),
		ray.NewLight(tuple.NewPoint(1, 0, 1), color.New(0.9, 0.7, 0)),
	})
	r = ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	shape = w.Shapes()[1]
	i = ray.NewIntersection(0.5, shape)
	comps = ray.PrepareComputations(i, r, ray.Intersections{i})

	result = w.shadeHit(comps)
	expected = color.New(0.19, 0.16999999999999998, 0.1)
	if result != expected {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	w = Default()
	w.SetLights([]ray.Light{
		ray.NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1)),
	})
	s1 := shapes.NewSphere()
	s2 := shapes.NewSphere()
	s2.SetTransform(matrix.Translation(0, 0, 10))
	w.SetShapes([]shapes.Shape{s1, s2})

	r = ray.New(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1))
	i = ray.NewIntersection(4, s2)
	comps = ray.PrepareComputations(i, r, ray.Intersections{i})

	result = w.shadeHit(comps)
	expected = color.New(0.1, 0.1, 0.1)
	if result != expected {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// with reflective shapes

	w = Default()
	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	shape = shapes.NewPlane()
	shape.SetTransform(matrix.Translation(0, -1, 0))
	mat := materials.DefaultMaterial()
	mat.SetReflective(0.5)
	shape.SetMaterial(mat)
	i = ray.NewIntersection(math.Sqrt(2), shape)
	comps = ray.PrepareComputations(i, r, ray.Intersections{i})
	result = w.shadeHit(comps)
	expected = color.New(0.876755987245857, 0.924338636811946, 0.8291733376797681)

	if result != expected {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}
}

func TestColorAt(t *testing.T) {
	w := Default()
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 1, 0))
	result := w.ColorAt(r)
	expected := color.Black()
	if result != expected {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	w = Default()
	r = ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	result = w.ColorAt(r)
	expected = color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
	if result != expected {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	w = Default()
	outer := w.Shapes()[0]
	m := outer.Material()
	m.SetAmbient(1)
	outer.SetMaterial(m)
	outer.Material().SetPattern(m.Pattern())

	inner := w.Shapes()[1]
	m = inner.Material()
	m.SetAmbient(1)
	inner.SetMaterial(m)
	inner.Material().SetPattern(m.Pattern())

	r = ray.New(tuple.NewPoint(0, 0, 0.75), tuple.NewVector(0, 0, -1))
	result = w.ColorAt(r)
	expected = inner.Material().ColorAt(r.Origin())
	if result != expected {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}
}

func TestRecusingReflection(t *testing.T) {
	w := Default()
	w.SetLights([]ray.Light{
		ray.NewLight(tuple.NewPoint(0, 0, 0), color.New(1, 1, 1)),
	})

	mat := materials.DefaultMaterial()
	mat.SetReflective(1)

	lower := shapes.NewPlane()
	lower.SetMaterial(mat)
	lower.SetTransform(matrix.Translation(0, -1, 0))

	upper := shapes.NewPlane()
	upper.SetMaterial(mat)
	upper.SetTransform(matrix.Translation(0, 1, 0))

	w.SetShapes([]shapes.Shape{
		lower, upper,
	})

	r := ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 1, 0))
	// If the limit would not be in place this would run into an infinite recursion.
	w.ColorAt(r)
}

func TestShadowAt(t *testing.T) {
	w := Default()
	var tests = []struct {
		w        *World
		point    tuple.Tuple
		expected bool
	}{
		{
			w:        w,
			point:    tuple.NewPoint(0, 10, 0),
			expected: false,
		},
		{
			w:        w,
			point:    tuple.NewPoint(10, -10, 10),
			expected: true,
		},
		{
			w:        w,
			point:    tuple.NewPoint(-20, 20, -20),
			expected: false,
		},
		{
			w:        w,
			point:    tuple.NewPoint(-2, 2, -2),
			expected: false,
		},
	}

	for _, test := range tests {
		result := test.w.shadowAt(test.point)
		if result != test.expected {
			t.Errorf("ShadowAt,\npoint:\n%s\nresult:\n%t\nexpected: \n%t", test.point, result, test.expected)
		}
	}
}

func TestReflectedColor(t *testing.T) {
	// The reflected color for a nonreflective material
	w := Default()
	r := ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	shape := w.Shapes()[1]
	mat := shape.Material()
	mat.SetAmbient(1)
	i := ray.NewIntersection(1, shape)
	comps := ray.PrepareComputations(i, r, ray.Intersections{i})
	result := w.reflectedColor(comps)
	expected := color.Black()

	if result != expected {
		t.Errorf("incorrect reflected color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The reflected color for a reflective material

	w = Default()
	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	shape = shapes.NewPlane()
	shape.SetTransform(matrix.Translation(0, -1, 0))
	mat = materials.DefaultMaterial()
	mat.SetReflective(0.5)
	shape.SetMaterial(mat)
	i = ray.NewIntersection(math.Sqrt(2), shape)
	comps = ray.PrepareComputations(i, r, ray.Intersections{i})
	result = w.reflectedColor(comps)
	expected = color.New(0.1903305982643556, 0.23791324783044449, 0.14274794869826668)

	if result != expected {
		t.Errorf("incorrect reflected color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// Returns when ray has reached the maximum recursive depth.

	w = Default()
	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	r.SetBounceLimit(0)
	shape = shapes.NewPlane()
	shape.SetTransform(matrix.Translation(0, -1, 0))
	mat = materials.DefaultMaterial()
	mat.SetReflective(0.5)
	shape.SetMaterial(mat)
	i = ray.NewIntersection(math.Sqrt(2), shape)
	comps = ray.PrepareComputations(i, r, ray.Intersections{i})
	result = w.reflectedColor(comps)
	expected = color.Black()

	if result != expected {
		t.Errorf("incorrect reflected color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}
}
