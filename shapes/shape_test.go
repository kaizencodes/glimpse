package shapes

import (
	"glimpse/color"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
	"math"
	"testing"
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
		if result := NormalAt(test.point, test.shape); !result.Equal(test.expected) {
			t.Errorf("test normal:\n%s \n point: %s. \nresult: \n%s. \nexpected: \n%s", test.shape, test.point, result, test.expected)
		}
	}

	// Computing the normal on a translated shape
	shape := NewTestShape()
	shape.SetTransform(matrix.Translation(0, 1, 0))
	point := tuple.NewPoint(0, 1.70711, -0.70711)
	want := tuple.NewVector(0, 0.7071067811865475, -0.7071067811865476)

	if got := NormalAt(point, shape); !got.Equal(want) {
		t.Errorf("test normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", shape, point, got, want)
	}

	// Computing the normal on a transformed shape
	shape = NewTestShape()
	transform, _ := matrix.Multiply(matrix.Scaling(1, 0.5, 1), matrix.RotationZ(math.Pi/5.0))
	shape.SetTransform(transform)
	point = tuple.NewPoint(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0)
	want = tuple.NewVector(0, 0.9701425001453319, -0.24253562503633294)

	if got := NormalAt(point, shape); !got.Equal(want) {
		t.Errorf("test normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", shape, point, got, want)
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
	result := NormalAt(tuple.NewPoint(1.7321, 1.1547, -5.5774), s)
	expected := tuple.NewVector(0.28570368184140726, 0.42854315178114105, -0.8571605294481017)
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

func TestWorldToObject(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(matrix.RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(matrix.Scaling(2, 2, 2))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))
	g2.AddChild(s)
	result := worldToObject(tuple.NewPoint(-2, 0, -10), s)
	expected := tuple.NewPoint(0, 0, -1)
	if !result.Equal(expected) {
		t.Errorf("incorrect point convertion to object space.\nexpected: %s\nresult: %s", expected, result)
	}
}

func TestNormalToWorld(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(matrix.RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(matrix.Scaling(1, 2, 3))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))
	g2.AddChild(s)
	result := normalToWorld(tuple.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3), s)
	expected := tuple.NewVector(0.28571428571428575, 0.42857142857142855, -0.8571428571428571)
	if !result.Equal(expected) {
		t.Errorf("incorrect point convertion to object space.\nexpected: %s\nresult: %s", expected, result)
	}
}
