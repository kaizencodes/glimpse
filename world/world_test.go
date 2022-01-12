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
	if !result.Equal(expected) {
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
	if !result.Equal(expected) {
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
	if !result.Equal(expected) {
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
	if !result.Equal(expected) {
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

	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// with refractive shapes

	w = Default()
	floor := shapes.NewPlane()
	floor.SetTransform(matrix.Translation(0, -1, 0))
	mat = materials.DefaultMaterial()
	mat.SetTransparency(0.5)
	mat.SetRefractiveIndex(1.5)
	floor.SetMaterial(mat)

	ball := shapes.NewSphere()
	mat = materials.DefaultMaterial()
	mat.SetPattern(materials.NewPattern(materials.Base, color.Red()))
	mat.SetAmbient(0.5)
	ball.SetTransform(matrix.Translation(0, -3.5, -0.5))
	ball.SetMaterial(mat)
	w.SetShapes([]shapes.Shape{w.shapes[0], w.shapes[1], floor, ball})

	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i = ray.NewIntersection(math.Sqrt(2), floor)
	comps = ray.PrepareComputations(i, r, ray.Intersections{i})
	result = w.shadeHit(comps)
	expected = color.New(0.936425388674727, 0.686425388674727, 0.686425388674727)

	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	w = Default()
	floor = shapes.NewPlane()
	floor.SetTransform(matrix.Translation(0, -1, 0))
	mat = materials.DefaultMaterial()
	mat.SetTransparency(0.5)
	mat.SetReflective(0.5)
	mat.SetRefractiveIndex(1.5)
	floor.SetMaterial(mat)

	ball = shapes.NewSphere()
	mat = materials.DefaultMaterial()
	mat.SetPattern(materials.NewPattern(materials.Base, color.Red()))
	mat.SetAmbient(0.5)
	ball.SetMaterial(mat)
	ball.SetTransform(matrix.Translation(0, -3.5, -0.5))
	w.SetShapes([]shapes.Shape{w.shapes[0], w.shapes[1], floor, ball})

	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i = ray.NewIntersection(math.Sqrt(2), floor)
	comps = ray.PrepareComputations(i, r, ray.Intersections{i})
	result = w.shadeHit(comps)
	expected = color.New(0.9339151403109409, 0.6964342260713607, 0.6924306911127073)

	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}
}

func TestColorAt(t *testing.T) {
	w := Default()
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 1, 0))
	result := w.ColorAt(r)
	expected := color.Black()
	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	w = Default()
	r = ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	result = w.ColorAt(r)
	expected = color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
	if !result.Equal(expected) {
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
	if !result.Equal(expected) {
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

	if !result.Equal(expected) {
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

	if !result.Equal(expected) {
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

	if !result.Equal(expected) {
		t.Errorf("incorrect reflected color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}
}

func TestRefractedColor(t *testing.T) {
	// The refracted color with an opaque surface
	w := Default()
	shape := w.Shapes()[0]
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	xs := ray.Intersections{
		ray.NewIntersection(4, shape),
		ray.NewIntersection(6, shape),
	}
	comps := ray.PrepareComputations(xs[0], r, xs)
	result := w.refractedColor(comps)
	expected := color.Black()

	if !result.Equal(expected) {
		t.Errorf("incorrect refracted color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The refracted color at the maximum recursive depth.

	w = Default()
	shape = w.Shapes()[0]
	mat := shape.Material()
	mat.SetTransparency(1)
	mat.SetRefractiveIndex(1.5)
	shape.SetMaterial(mat)
	r = ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	r.SetBounceLimit(0)
	xs = ray.Intersections{
		ray.NewIntersection(4, shape),
		ray.NewIntersection(6, shape),
	}
	comps = ray.PrepareComputations(xs[0], r, xs)
	result = w.refractedColor(comps)
	expected = color.Black()

	if !result.Equal(expected) {
		t.Errorf("incorrect refracted color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The refracted color under total internal reflection.

	w = Default()
	shape = w.Shapes()[0]
	mat = shape.Material()
	mat.SetTransparency(1)
	mat.SetRefractiveIndex(1.5)
	shape.SetMaterial(mat)
	r = ray.New(tuple.NewPoint(0, 0, math.Sqrt(2)/2), tuple.NewVector(0, 1, 0))
	xs = ray.Intersections{
		ray.NewIntersection(-math.Sqrt(2)/2, shape),
		ray.NewIntersection(math.Sqrt(2)/2, shape),
	}
	comps = ray.PrepareComputations(xs[1], r, xs)
	result = w.refractedColor(comps)
	expected = color.Black()

	if !result.Equal(expected) {
		t.Errorf("incorrect refracted color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The refracted color with a refracted ray.

	w = Default()
	a := w.Shapes()[0]
	mat = a.Material()
	mat.SetAmbient(1.0)
	mat.SetPattern(materials.NewPattern(materials.Test))
	a.SetMaterial(mat)
	b := w.Shapes()[1]
	mat = b.Material()
	mat.SetTransparency(1.0)
	mat.SetRefractiveIndex(1.5)
	mat.SetPattern(materials.NewPattern(materials.Test))
	b.SetMaterial(mat)

	r = ray.New(tuple.NewPoint(0, 0, 0.1), tuple.NewVector(0, 1, 0))
	xs = ray.Intersections{
		ray.NewIntersection(-0.9899, a),
		ray.NewIntersection(-0.4899, b),
		ray.NewIntersection(0.4899, b),
		ray.NewIntersection(0.9899, a),
	}
	comps = ray.PrepareComputations(xs[2], r, xs)
	result = w.refractedColor(comps)
	expected = color.New(0, 0.9988846826559641, 0.04721642463480325)

	if !result.Equal(expected) {
		t.Errorf("incorrect refracted color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

}
