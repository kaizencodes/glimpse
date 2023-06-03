package renderer

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/camera"
	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/light"
	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/scenes"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

func TestRender(t *testing.T) {
	scene := scenes.Default()
	c := camera.New(11, 11, math.Pi/2)
	transform := camera.ViewTransformation(
		tuple.NewPoint(0, 0, -5),
		tuple.NewPoint(0, 0, 0),
		tuple.NewVector(0, 1, 0),
	)
	c.SetTransform(transform)
	img := Render(c, scene)
	result := img[5][5]
	expected := color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
	if !result.Equal(expected) {
		t.Errorf("Render, expected %s, got %s", expected, result)
	}
}

func TestIntersect(t *testing.T) {
	scene := scenes.Default()
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	sections := intersect(scene, r)
	expected := []float64{4, 4.5, 5.5, 6}
	for i, v := range expected {
		if sections[i].T() != v {
			t.Errorf("incorrect t of intersect:\n%s \n \ngot: \n%f. \nexpected: \n%f", r, sections[i].T(), v)
		}
	}
}

func TestShadeHit(t *testing.T) {
	// Shading an intersection
	scene := scenes.Default()
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	shape := scene.Shapes[0]
	i := shapes.NewIntersection(4, shape)
	comps := prepareComputations(i, r, shapes.Intersections{i})

	result := shadeHit(scene, comps)
	expected := color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// Shading an intersection from the inside
	scene = scenes.Default()
	scene.Lights = []light.Light{
		light.NewLight(tuple.NewPoint(0, 0.25, 0), color.New(1, 1, 1)),
	}
	r = ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	shape = scene.Shapes[1]
	i = shapes.NewIntersection(0.5, shape)
	comps = prepareComputations(i, r, shapes.Intersections{i})

	result = shadeHit(scene, comps)
	expected = color.New(0.9049844720832575, 0.9049844720832575, 0.9049844720832575)
	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// multiple light sources
	scene = scenes.Default()
	scene.Lights = []light.Light{
		light.NewLight(tuple.NewPoint(0, 0.25, 0), color.New(1, 1, 1)),
		light.NewLight(tuple.NewPoint(1, 0, 1), color.New(0.9, 0.7, 0)),
	}
	r = ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	shape = scene.Shapes[1]
	i = shapes.NewIntersection(0.5, shape)
	comps = prepareComputations(i, r, shapes.Intersections{i})

	result = shadeHit(scene, comps)
	expected = color.New(0.19, 0.16999999999999998, 0.1)
	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// shadeHit() is given an intersection in shadow
	scene = scenes.Default()
	scene.Lights = []light.Light{
		light.NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1)),
	}
	s1 := shapes.NewSphere()
	s2 := shapes.NewSphere()
	s2.SetTransform(matrix.Translation(0, 0, 10))
	scene.Shapes = []shapes.Shape{s1, s2}

	r = ray.New(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1))
	i = shapes.NewIntersection(4, s2)
	comps = prepareComputations(i, r, shapes.Intersections{i})

	result = shadeHit(scene, comps)
	expected = color.New(0.1, 0.1, 0.1)
	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// with reflective shapes

	scene = scenes.Default()
	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	shape = shapes.NewPlane()
	shape.SetTransform(matrix.Translation(0, -1, 0))
	mat := materials.DefaultMaterial()
	mat.Reflective = 0.5
	shape.SetMaterial(mat)
	i = shapes.NewIntersection(math.Sqrt(2), shape)
	comps = prepareComputations(i, r, shapes.Intersections{i})
	result = shadeHit(scene, comps)
	expected = color.New(0.876755987245857, 0.924338636811946, 0.8291733376797681)

	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// with refractive shapes

	scene = scenes.Default()
	floor := shapes.NewPlane()
	floor.SetTransform(matrix.Translation(0, -1, 0))
	mat = materials.DefaultMaterial()
	mat.Transparency = 0.5
	mat.RefractiveIndex = 1.5
	floor.SetMaterial(mat)

	ball := shapes.NewSphere()
	mat = materials.DefaultMaterial()
	mat.SetPattern(materials.NewPattern(materials.Base, color.Red()))
	mat.Ambient = 0.5
	ball.SetTransform(matrix.Translation(0, -3.5, -0.5))
	ball.SetMaterial(mat)
	scene.Shapes = []shapes.Shape{scene.Shapes[0], scene.Shapes[1], floor, ball}

	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i = shapes.NewIntersection(math.Sqrt(2), floor)
	comps = prepareComputations(i, r, shapes.Intersections{i})
	result = shadeHit(scene, comps)
	expected = color.New(0.936425388674727, 0.686425388674727, 0.686425388674727)

	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	scene = scenes.Default()
	floor = shapes.NewPlane()
	floor.SetTransform(matrix.Translation(0, -1, 0))
	mat = materials.DefaultMaterial()
	mat.Transparency = 0.5
	mat.Reflective = 0.5
	mat.RefractiveIndex = 1.5
	floor.SetMaterial(mat)

	ball = shapes.NewSphere()
	mat = materials.DefaultMaterial()
	mat.SetPattern(materials.NewPattern(materials.Base, color.Red()))
	mat.Ambient = 0.5
	ball.SetMaterial(mat)
	ball.SetTransform(matrix.Translation(0, -3.5, -0.5))
	scene.Shapes = []shapes.Shape{scene.Shapes[0], scene.Shapes[1], floor, ball}

	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i = shapes.NewIntersection(math.Sqrt(2), floor)
	comps = prepareComputations(i, r, shapes.Intersections{i})
	result = shadeHit(scene, comps)
	expected = color.New(0.9339151403109409, 0.6964342260713607, 0.6924306911127073)

	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}
}

func TestColorAt(t *testing.T) {
	// The color when a ray misses
	scene := scenes.Default()
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 1, 0))
	result := colorAt(scene, r)
	expected := color.Black()
	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The color when a ray hits
	scene = scenes.Default()
	r = ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	result = colorAt(scene, r)
	expected = color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The color with an intersection behind the ray
	scene = scenes.Default()
	outer := scene.Shapes[0]
	m := outer.Material()
	m.Ambient = 1
	outer.SetMaterial(m)
	outer.Material().SetPattern(m.Pattern())

	inner := scene.Shapes[1]
	m = inner.Material()
	m.Ambient = 1
	inner.SetMaterial(m)
	inner.Material().SetPattern(m.Pattern())

	r = ray.New(tuple.NewPoint(0, 0, 0.75), tuple.NewVector(0, 0, -1))
	result = colorAt(scene, r)
	expected = inner.Material().ColorAt(r.Origin)
	if !result.Equal(expected) {
		t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}
}

func TestRecusingReflection(t *testing.T) {
	scene := scenes.Default()
	scene.Lights = []light.Light{
		light.NewLight(tuple.NewPoint(0, 0, 0), color.New(1, 1, 1)),
	}

	mat := materials.DefaultMaterial()
	mat.Reflective = 1

	lower := shapes.NewPlane()
	lower.SetMaterial(mat)
	lower.SetTransform(matrix.Translation(0, -1, 0))

	upper := shapes.NewPlane()
	upper.SetMaterial(mat)
	upper.SetTransform(matrix.Translation(0, 1, 0))

	scene.Shapes = []shapes.Shape{
		lower, upper,
	}

	r := ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 1, 0))
	// If the limit would not be in place this would run into an infinite recursion.
	colorAt(scene, r)
}

func TestShadowAt(t *testing.T) {
	scene := scenes.Default()
	var tests = []struct {
		scene    *scenes.Scene
		point    tuple.Tuple
		expected bool
	}{
		{
			// There is no shadow when nothing is collinear with point and light
			scene:    scene,
			point:    tuple.NewPoint(0, 10, 0),
			expected: false,
		},
		{
			// The shadow when an object is between the point and the light
			scene:    scene,
			point:    tuple.NewPoint(10, -10, 10),
			expected: true,
		},
		{
			// There is no shadow when an object is behind the light
			scene:    scene,
			point:    tuple.NewPoint(-20, 20, -20),
			expected: false,
		},
		{
			// There is no shadow when an object is behind the point
			scene:    scene,
			point:    tuple.NewPoint(-2, 2, -2),
			expected: false,
		},
	}

	for _, test := range tests {
		result := shadowAt(test.scene, test.point)
		if result != test.expected {
			t.Errorf("ShadowAt,\npoint:\n%s\nresult:\n%t\nexpected: \n%t", test.point, result, test.expected)
		}
	}
}

func TestReflectedColor(t *testing.T) {
	// The reflected color for a nonreflective material
	scene := scenes.Default()
	r := ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	shape := scene.Shapes[1]
	mat := shape.Material()
	mat.Ambient = 1
	i := shapes.NewIntersection(1, shape)
	comps := prepareComputations(i, r, shapes.Intersections{i})
	result := reflectedColor(scene, comps)
	expected := color.Black()

	if !result.Equal(expected) {
		t.Errorf("incorrect reflected color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The reflected color for a reflective material

	scene = scenes.Default()
	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	shape = shapes.NewPlane()
	shape.SetTransform(matrix.Translation(0, -1, 0))
	mat = materials.DefaultMaterial()
	mat.Reflective = 0.5
	shape.SetMaterial(mat)
	i = shapes.NewIntersection(math.Sqrt(2), shape)
	comps = prepareComputations(i, r, shapes.Intersections{i})
	result = reflectedColor(scene, comps)
	expected = color.New(0.1903305982643556, 0.23791324783044449, 0.14274794869826668)

	if !result.Equal(expected) {
		t.Errorf("incorrect reflected color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// Returns when ray has reached the maximum recursive depth.

	scene = scenes.Default()
	r = ray.New(tuple.NewPoint(0, 0, -3), tuple.NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	r.BounceLimit = 0
	shape = shapes.NewPlane()
	shape.SetTransform(matrix.Translation(0, -1, 0))
	mat = materials.DefaultMaterial()
	mat.Reflective = 0.5
	shape.SetMaterial(mat)
	i = shapes.NewIntersection(math.Sqrt(2), shape)
	comps = prepareComputations(i, r, shapes.Intersections{i})
	result = reflectedColor(scene, comps)
	expected = color.Black()

	if !result.Equal(expected) {
		t.Errorf("incorrect reflected color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}
}

func TestRefractedColor(t *testing.T) {
	// The refracted color with an opaque surface
	scene := scenes.Default()
	shape := scene.Shapes[0]
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	xs := shapes.Intersections{
		shapes.NewIntersection(4, shape),
		shapes.NewIntersection(6, shape),
	}
	comps := prepareComputations(xs[0], r, xs)
	result := refractedColor(scene, comps)
	expected := color.Black()

	if !result.Equal(expected) {
		t.Errorf("incorrect refracted color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The refracted color at the maximum recursive depth.

	scene = scenes.Default()
	shape = scene.Shapes[0]
	mat := shape.Material()
	mat.Transparency = 1
	mat.RefractiveIndex = 1.5
	shape.SetMaterial(mat)
	r = ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	r.BounceLimit = 0
	xs = shapes.Intersections{
		shapes.NewIntersection(4, shape),
		shapes.NewIntersection(6, shape),
	}
	comps = prepareComputations(xs[0], r, xs)
	result = refractedColor(scene, comps)
	expected = color.Black()

	if !result.Equal(expected) {
		t.Errorf("incorrect refracted color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The refracted color under total internal reflection.

	scene = scenes.Default()
	shape = scene.Shapes[0]
	mat = shape.Material()
	mat.Transparency = 1
	mat.RefractiveIndex = 1.5
	shape.SetMaterial(mat)
	r = ray.New(tuple.NewPoint(0, 0, math.Sqrt(2)/2), tuple.NewVector(0, 1, 0))
	xs = shapes.Intersections{
		shapes.NewIntersection(-math.Sqrt(2)/2, shape),
		shapes.NewIntersection(math.Sqrt(2)/2, shape),
	}
	comps = prepareComputations(xs[1], r, xs)
	result = refractedColor(scene, comps)
	expected = color.Black()

	if !result.Equal(expected) {
		t.Errorf("incorrect refracted color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

	// The refracted color with a refracted ray.

	scene = scenes.Default()
	a := scene.Shapes[0]
	mat = a.Material()
	mat.Ambient = 1.0
	mat.SetPattern(materials.NewPattern(materials.Test))
	a.SetMaterial(mat)
	b := scene.Shapes[1]
	mat = b.Material()
	mat.Transparency = 1.0
	mat.RefractiveIndex = 1.5
	mat.SetPattern(materials.NewPattern(materials.Test))
	b.SetMaterial(mat)

	r = ray.New(tuple.NewPoint(0, 0, 0.1), tuple.NewVector(0, 1, 0))
	xs = shapes.Intersections{
		shapes.NewIntersection(-0.9899, a),
		shapes.NewIntersection(-0.4899, b),
		shapes.NewIntersection(0.4899, b),
		shapes.NewIntersection(0.9899, a),
	}
	comps = prepareComputations(xs[2], r, xs)
	result = refractedColor(scene, comps)
	expected = color.New(0, 0.9988846826559641, 0.04721642463480325)

	if !result.Equal(expected) {
		t.Errorf("incorrect refracted color:\nresult: \n%s. \nexpected: \n%s", result, expected)
	}

}
