package materials

import (
	"glimpse/color"
	"glimpse/tuple"
	"testing"
)

func TestStripePattern(t *testing.T) {
	pattern := newStripePattern(color.White(), color.Black())
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
		if result := pattern.colorAt(test.point); !test.expected.Equal(result) {
			t.Errorf("ColorAt:%s, result: \n%s. \nexpected: \n%s", test.point, result, test.expected)
		}
	}
}

func TestGradientPattern(t *testing.T) {
	// A gradient linearly interpolates between colors
	pattern := newGradientPattern(color.White(), color.Black())
	var tests = []struct {
		point    tuple.Tuple
		expected color.Color
	}{
		{
			point:    tuple.NewPoint(0, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0.25, 0, 0),
			expected: color.New(0.75, 0.75, 0.75),
		},
		{
			point:    tuple.NewPoint(0.5, 0, 0),
			expected: color.New(0.5, 0.5, 0.5),
		},
		{
			point:    tuple.NewPoint(0.75, 0, 0),
			expected: color.New(0.25, 0.25, 0.25),
		},
	}

	for _, test := range tests {
		if result := pattern.colorAt(test.point); !test.expected.Equal(result) {
			t.Errorf("ColorAt:%s, result: \n%s. \nexpected: \n%s", test.point, result, test.expected)
		}
	}
}

func TestRingPattern(t *testing.T) {
	// A ring should extend in both x and z
	pattern := newRingPattern(color.White(), color.Black())
	var tests = []struct {
		point    tuple.Tuple
		expected color.Color
	}{
		{
			point:    tuple.NewPoint(0, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(1, 0, 0),
			expected: color.Black(),
		},
		{
			point:    tuple.NewPoint(0, 0, 1),
			expected: color.Black(),
		},
		{
			point:    tuple.NewPoint(0.708, 0, 0.708),
			expected: color.Black(),
		},
	}

	for _, test := range tests {
		if result := pattern.colorAt(test.point); !test.expected.Equal(result) {
			t.Errorf("ColorAt:%s, result: \n%s. \nexpected: \n%s", test.point, result, test.expected)
		}
	}
}

func TestCheckerPattern(t *testing.T) {
	pattern := newCheckerPattern(color.White(), color.Black())
	var tests = []struct {
		point    tuple.Tuple
		expected color.Color
	}{
		// Checkers should repeat in x
		{
			point:    tuple.NewPoint(0, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0.99, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(1.01, 0, 0),
			expected: color.Black(),
		},
		// Checkers should repeat in y
		{
			point:    tuple.NewPoint(0, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0, 0.99, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0, 1.01, 0),
			expected: color.Black(),
		},
		// Checkers should repeat in y
		{
			point:    tuple.NewPoint(0, 0, 0),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0, 0, 0.99),
			expected: color.White(),
		},
		{
			point:    tuple.NewPoint(0, 0, 1.01),
			expected: color.Black(),
		},
	}

	for _, test := range tests {
		if result := pattern.colorAt(test.point); !test.expected.Equal(result) {
			t.Errorf("ColorAt:%s, result: \n%s. \nexpected: \n%s", test.point, result, test.expected)
		}
	}
}

func TestTestPattern(t *testing.T) {
	pattern := newTestPattern()
	var tests = []struct {
		point    tuple.Tuple
		expected color.Color
	}{
		{
			point:    tuple.NewPoint(0, 0, 0),
			expected: color.Black(),
		},
		{
			point:    tuple.NewPoint(0.99, 0.01, 0.5),
			expected: color.New(0.99, 0.01, 0.5),
		},
	}

	for _, test := range tests {
		if result := pattern.colorAt(test.point); !test.expected.Equal(result) {
			t.Errorf("ColorAt:%s, result: \n%s. \nexpected: \n%s", test.point, result, test.expected)
		}
	}
}
