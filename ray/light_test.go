package ray

import (
	"glimpse/color"
	"glimpse/patterns"
	"glimpse/shapes"
	"glimpse/tuple"
	"math"
	"testing"
)

func TestLighting(t *testing.T) {
	var tests = []struct {
		eyeV     tuple.Tuple
		normalV  tuple.Tuple
		light    Light
		inShadow bool
		expected color.Color
	}{
		{
			eyeV:     tuple.NewVector(0, 0, -1),
			normalV:  tuple.NewVector(0, 0, -1),
			light:    NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1)),
			inShadow: false,
			expected: color.New(1.9, 1.9, 1.9),
		},
		{
			eyeV:     tuple.NewVector(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0),
			normalV:  tuple.NewVector(0, 0, -1),
			light:    NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1)),
			inShadow: false,
			expected: color.New(1.0, 1.0, 1.0),
		},
		{
			eyeV:     tuple.NewVector(0, 0, -1),
			normalV:  tuple.NewVector(0, 0, -1),
			light:    NewLight(tuple.NewPoint(0, 10, -10), color.New(1, 1, 1)),
			inShadow: false,
			expected: color.New(0.7363961030678927, 0.7363961030678927, 0.7363961030678927),
		},
		{
			eyeV:     tuple.NewVector(0, -math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0),
			normalV:  tuple.NewVector(0, 0, -1),
			light:    NewLight(tuple.NewPoint(0, 10, -10), color.New(1, 1, 1)),
			inShadow: false,
			expected: color.New(1.6363961030678928, 1.6363961030678928, 1.6363961030678928),
		},
		{
			eyeV:     tuple.NewVector(0, 0, -1),
			normalV:  tuple.NewVector(0, 0, -1),
			light:    NewLight(tuple.NewPoint(0, 0, 10), color.New(1, 1, 1)),
			inShadow: false,
			expected: color.New(0.1, 0.1, 0.1),
		},
		{
			eyeV:     tuple.NewVector(0, 0, -1),
			normalV:  tuple.NewVector(0, 0, -1),
			light:    NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1)),
			inShadow: true,
			expected: color.New(0.1, 0.1, 0.1),
		},
	}

	shape := shapes.NewSphere()
	pos := tuple.NewPoint(0, 0, 0)
	for _, test := range tests {
		if got := Lighting(shape, test.light, pos, test.eyeV, test.normalV, test.inShadow); !got.Equal(test.expected) {
			t.Errorf("Lighting:\n light: %s \neyeV: %s \nnormalV: %s\ninShadow: %t\ngot: \n%s. \nexpected: \n%s", test.light, test.eyeV, test.normalV, test.inShadow, got, test.expected)
		}
	}

	// Lightning with a pattern applied

	eyeV := tuple.NewVector(0, 0, -1)
	normalV := tuple.NewVector(0, 0, -1)
	light := NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1))
	inShadow := false
	pattern := patterns.NewStripePattern(color.White(), color.Black())
	ambientMat := shapes.NewMaterial(pattern, 1, 0, 0, 0)
	shape.SetMaterial(ambientMat)
	pos1 := tuple.NewPoint(0.9, 0, 0)
	pos2 := tuple.NewPoint(1.1, 0, 0)

	color1 := Lighting(shape, light, pos1, eyeV, normalV, inShadow)
	color2 := Lighting(shape, light, pos2, eyeV, normalV, inShadow)

	if !color1.Equal(color.White()) {
		t.Errorf("Lighting:\n light: %s \neyeV: %s \nnormalV: %s\ninShadow: %t\ngot: \n%s. \nexpected: \n%s", light, eyeV, normalV, inShadow, color1, color.White())
	}
	if !color2.Equal(color.Black()) {
		t.Errorf("Lighting:\n light: %s \neyeV: %s \nnormalV: %s\ninShadow: %t\ngot: \n%s. \nexpected: \n%s", light, eyeV, normalV, inShadow, color1, color.Black())
	}

}
