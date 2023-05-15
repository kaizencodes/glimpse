package utils

import "testing"

func TestFloatEquals(t *testing.T) {
	var tests = []struct {
		left     float64
		right    float64
		expected bool
	}{
		{1.000000001, 1.000000002, true},
		{1.0000001, 1.0000002, false},
	}

	for _, test := range tests {
		if result := FloatEquals(test.left, test.right); result != test.expected {
			t.Errorf("%f == %f result %t, expected %t", test.left, test.right, result, test.expected)
		}
	}
}
