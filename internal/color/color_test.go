package color

import "testing"

func TestAdd(t *testing.T) {
	var tests = []struct {
		left     Color
		right    Color
		expected Color
	}{
		{
			left:     Color{1.0, 1.0, 1.0},
			right:    Color{0.0, 1.5, -1.0},
			expected: Color{1.0, 2.5, 0.0},
		},
		{
			left:     Color{-1.0, 1.0, 1.0},
			right:    Color{-2.0, 1.5, 0.0001},
			expected: Color{-3.0, 2.5, 1.0001},
		},
	}

	for _, test := range tests {
		if got := Add(test.left, test.right); !got.Equal(test.expected) {
			t.Errorf("input: %s + %s \ngot: %s. \nexpected: %s", test.left, test.right, got, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	var tests = []struct {
		left     Color
		right    Color
		expected Color
	}{
		{
			left:     Color{1.0, 1.0, 1.0},
			right:    Color{0.0, 1.5, -1.0},
			expected: Color{1.0, -0.5, 2.0},
		},
		{
			left:     Color{-1.0, 1.0, 1.0},
			right:    Color{-2.0, 1.5, 0.0001},
			expected: Color{1.0, -0.5, 0.9999},
		},
	}

	for _, test := range tests {
		if got := Subtract(test.left, test.right); !got.Equal(test.expected) {
			t.Errorf("input: %s - %s \ngot: %s. \nexpected: %s", test.left, test.right, got, test.expected)
		}
	}
}

func TestScalar(t *testing.T) {
	var tests = []struct {
		tuple    Color
		scalar   float64
		expected Color
	}{
		{
			tuple:    Color{1.0, 1.0, 1.0},
			scalar:   1,
			expected: Color{1.0, 1.0, 1.0},
		},
		{
			tuple:    Color{-2.0, 1.5, 0.5},
			scalar:   0.5,
			expected: Color{-1.0, 0.75, 0.25},
		},
		{
			tuple:    Color{-2.0, 1.5, 0.5},
			scalar:   2,
			expected: Color{-4.0, 3, 1.0},
		},
	}

	for _, test := range tests {
		if got := test.tuple.Scalar(test.scalar); !got.Equal(test.expected) {
			t.Errorf("input: %s + %f \ngot: %s. \nexpected: %s", test.tuple, test.scalar, got, test.expected)
		}
	}
}

func TestHadamardProduct(t *testing.T) {
	var tests = []struct {
		left     Color
		right    Color
		expected Color
	}{
		{
			left:     Color{1.0, 0.9, 0.5},
			right:    Color{0.0, 0.5, 0.5},
			expected: Color{0.0, 0.45, 0.25},
		},
	}

	for _, test := range tests {
		if got := HadamardProduct(test.left, test.right); !got.Equal(test.expected) {
			t.Errorf("input: %s - %s \ngot: %s. \nexpected: %s", test.left, test.right, got, test.expected)
		}
	}
}
