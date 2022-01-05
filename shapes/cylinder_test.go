package shapes

import (
	"glimpse/tuple"
	"testing"
)

func TestCylinderLocalNormalAt(t *testing.T) {
	cylinder := NewCylinder()
	var tests = []struct {
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			point:    tuple.NewPoint(1, 0, 0),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			point:    tuple.NewPoint(0, 5, -1),
			expected: tuple.NewVector(0, 0, -1),
		},
		{
			point:    tuple.NewPoint(0, -2, 1),
			expected: tuple.NewVector(0, 0, 1),
		},
		{
			point:    tuple.NewPoint(-1, 1, 0),
			expected: tuple.NewVector(-1, 0, 0),
		},
	}

	for _, test := range tests {
		if result := cylinder.LocalNormalAt(test.point); !result.Equal(test.expected) {
			t.Errorf("Cylinder normal: \nresult: \n%s. \nexpected: \n%s", result, test.expected)
		}
	}
}

func TestClosedCylinderLocalNormalAt(t *testing.T) {
	cylinder := NewCylinder()
	cylinder.SetMinimum(1)
	cylinder.SetMaximum(2)
	cylinder.SetClosed(true)

	var tests = []struct {
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			point:    tuple.NewPoint(0, 1, 0),
			expected: tuple.NewVector(0, -1, 0),
		},
		{
			point:    tuple.NewPoint(0.5, 1, 0),
			expected: tuple.NewVector(0, -1, 0),
		},
		{
			point:    tuple.NewPoint(0, 1, 0.5),
			expected: tuple.NewVector(0, -1, 0),
		},
		{
			point:    tuple.NewPoint(0, 2, 0),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			point:    tuple.NewPoint(0.5, 2, 0),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			point:    tuple.NewPoint(0, 2, 0.5),
			expected: tuple.NewVector(0, 1, 0),
		},
	}

	for _, test := range tests {
		if result := cylinder.LocalNormalAt(test.point); !result.Equal(test.expected) {
			t.Errorf("Cylinder normal: \nresult: \n%s. \nexpected: \n%s", result, test.expected)
		}
	}
}
