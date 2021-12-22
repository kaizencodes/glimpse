package patterns

import (
	"glimpse/color"
	"glimpse/tuple"
	"testing"
)

func TestStripeAt(t *testing.T) {
	pattern := StripePattern(color.White(), color.Black())
	var tests = []struct {
		point    tuple.Tuple
		expected color.Color
	}{
		// A stripe pattern is constant in Y
		{
			point:    tuple.NewPoint(0, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0, 1, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0, 2, 0),
			expected: color.White(),
		},
		// A stripe pattern is constant in Z
		{
			point:    tuple.NewPoint(0, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0, 0, 1),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0, 0, 2),
			expected: color.White(),
		},
		// A stripe pattern alternates in X
		{
			point:    tuple.NewPoint(0, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0.9, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(1, 0, 0),
			expected: color.Black(),
		},
		{
			point:    tuple.NewPoint(-0.1, 0, 0),
			expected: color.Black(),
		},
		{
			point:    tuple.NewPoint(-1, 0, 0),
			expected: color.Black(),
		},
		{
			point:    tuple.NewPoint(-1.1, 0, 0),
			expected: color.White(),
		},
	}

	for _, test := range tests {
		if result := StripeAt(pattern, test.point); test.expected.Equal(result) {
			t.Errorf("StripeAt:%s, result: \n%s. \nexpected: \n%s", test.point, result, test.expected)
		}
	}
}
