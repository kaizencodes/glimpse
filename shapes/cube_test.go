package shapes

import (
	"glimpse/tuple"
	"testing"
)

func TestCubeLocalNormalAt(t *testing.T) {
	var tests = []struct {
		shape    *Cube
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(1, 0.5, -0.8),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(-1, -0.2, 0.9),
			expected: tuple.NewVector(-1, 0, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(-0.4, 1, -0.1),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(0.3, -1, -0.7),
			expected: tuple.NewVector(0, -1, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(-0.6, 0.3, 1),
			expected: tuple.NewVector(0, 0, 1),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(0.4, 0.4, -1),
			expected: tuple.NewVector(0, 0, -1),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(1, 1, 1),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(-1, -1, -1),
			expected: tuple.NewVector(-1, 0, 0),
		},
	}

	for _, test := range tests {
		if got := test.shape.LocalNormalAt(test.point); !got.Equal(test.expected) {
			t.Errorf("Sphere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", test.shape, test.point, got, test.expected)
		}
	}
}
