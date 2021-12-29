package shapes

import (
	"glimpse/tuple"
	"testing"
)

func TestPlaneLocalNormalAt(t *testing.T) {
	var tests = []struct {
		plane    *Plane
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			plane:    NewPlane(),
			point:    tuple.NewPoint(0, 0, 0),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			plane:    NewPlane(),
			point:    tuple.NewPoint(10, 0, -10),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			plane:    NewPlane(),
			point:    tuple.NewPoint(-5, 0, 150),
			expected: tuple.NewVector(0, 1, 0),
		},
	}

	for _, test := range tests {
		if got := test.plane.LocalNormalAt(test.point); !got.Equal(test.expected) {
			t.Errorf("Sphere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", test.plane, test.point, got, test.expected)
		}
	}
}
