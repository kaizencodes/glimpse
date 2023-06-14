package tuple

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestIsPoint(t *testing.T) {
	var tests = []struct {
		input    Tuple
		expected bool
	}{
		{
			// if w is 1, it is a point.
			input:    Tuple{4.3, -4.2, 3.1, 1.0},
			expected: true,
		},
		{
			// if w is 0, it is a vector.
			input:    Tuple{-3.3, 3.2, 3.1, 0.0},
			expected: false,
		},
	}

	for _, test := range tests {
		if got := test.input.IsPoint(); got != test.expected {
			t.Errorf("expected IsPoint() to be %t, got %t", got, test.expected)
		}
	}
}

func TestIsVector(t *testing.T) {
	var tests = []struct {
		input    Tuple
		expected bool
	}{
		{
			// if w is 1, it is a point.
			input:    Tuple{4.3, -4.2, 3.1, 1.0},
			expected: false,
		},
		{
			// if w is 0, it is a vector.
			input:    Tuple{-3.3, 3.2, 3.1, 0.0},
			expected: true,
		},
	}

	for _, test := range tests {
		if got := test.input.IsVector(); got != test.expected {
			t.Errorf("expected IsVector() to be %t, got %t", got, test.expected)
		}
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		left     Tuple
		right    Tuple
		expected bool
	}{
		{
			left:     Tuple{1.0, 1.0, 1.0, 1.0},
			right:    Tuple{1.0, 1.0, 1.0, 1.0},
			expected: true,
		},
		{
			left:     Tuple{1.0, 1.0, 1.0, 1.0},
			right:    Tuple{0.0, 1.0, 1.0, 1.0},
			expected: false,
		},
		{
			left:     Tuple{1.0, 1.0, 1.0, 1.0},
			right:    Tuple{1.0, 0.0, 1.0, 1.0},
			expected: false,
		},
		{
			left:     Tuple{1.0, 1.0, 1.0, 1.0},
			right:    Tuple{1.0, 1.0, 0.0, 1.0},
			expected: false,
		},
		{
			left:     Tuple{1.0, 1.0, 1.0, 1.0},
			right:    Tuple{1.0, 1.0, 1.0, 0.0},
			expected: false,
		},
	}

	for _, test := range tests {
		if got := test.left.Equal(test.right); got != test.expected {
			t.Errorf("expected tuples to be equal, but they were not.")
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		left     Tuple
		right    Tuple
		expected Tuple
	}{
		{
			left:     Tuple{1.0, 1.0, 1.0, 1.0},
			right:    Tuple{0.0, 1.5, -1.0, 0.0},
			expected: Tuple{1.0, 2.5, 0.0, 1.0},
		},
		{
			left:     Tuple{-1.0, 1.0, 1.0, 0.0},
			right:    Tuple{-2.0, 1.5, 0.0001, 0.0},
			expected: Tuple{-3.0, 2.5, 1.0001, 0.0},
		},
		{
			left:     Tuple{1.0, 1.0, 1.0, 0.0},
			right:    Tuple{0.0, 1.5, -1.0, 1.0},
			expected: Tuple{1.0, 2.5, 0.0, 1.0},
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
		left     Tuple
		right    Tuple
		expected Tuple
	}{
		{
			left:     Tuple{1.0, 1.0, 1.0, 1.0},
			right:    Tuple{0.0, 1.5, -1.0, 0.0},
			expected: Tuple{1.0, -0.5, 2.0, 1.0},
		},
		{
			left:     Tuple{-1.0, 1.0, 1.0, 0.0},
			right:    Tuple{-2.0, 1.5, 0.0001, 0.0},
			expected: Tuple{1.0, -0.5, 0.9999, 0.0},
		},
		{
			left:     Tuple{1.0, 1.0, 1.0, 1.0},
			right:    Tuple{0.0, 1.5, -1.0, 1.0},
			expected: Tuple{1.0, -0.5, 2.0, 0.0},
		},
	}

	for _, test := range tests {
		if got := Subtract(test.left, test.right); !got.Equal(test.expected) {
			t.Errorf("input: %s - %s \ngot: %s. \nexpected: %s", test.left, test.right, got, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	var tests = []struct {
		a        matrix.Matrix
		b        Tuple
		expected Tuple
	}{
		{
			a: matrix.New(4, 4,
				[]float64{
					1, 2, 3, 4,
					2, 4, 4, 2,
					8, 6, 4, 1,
					0, 0, 0, 1,
				},
			),
			b:        Tuple{1, 2, 3, 1},
			expected: Tuple{18, 24, 33, 1},
		},
	}

	for _, test := range tests {
		got := Multiply(test.a, test.b)
		if got.String() != test.expected.String() {
			t.Errorf("multiplication,\na:\n%s\nb:\n%s\ngot:\n%s\nexpected: \n%s", test.a, test.b, got, test.expected)
		}
	}
}

func TestScalar(t *testing.T) {
	var tests = []struct {
		tuple    Tuple
		scalar   float64
		expected Tuple
	}{
		{
			tuple:    Tuple{1.0, 1.0, 1.0, 1.0},
			scalar:   1,
			expected: Tuple{1.0, 1.0, 1.0, 1.0},
		},
		{
			tuple:    Tuple{-2.0, 1.5, 0.5, 0.0},
			scalar:   0.5,
			expected: Tuple{-1.0, 0.75, 0.25, 0.0},
		},
		{
			tuple:    Tuple{-2.0, 1.5, 0.5, 0.0},
			scalar:   2,
			expected: Tuple{-4.0, 3, 1.0, 0.0},
		},
	}

	for _, test := range tests {
		if got := test.tuple.Scalar(test.scalar); !got.Equal(test.expected) {
			t.Errorf("input: %s + %f \ngot: %s. \nexpected: %s", test.tuple, test.scalar, got, test.expected)
		}
	}
}

func TestNegate(t *testing.T) {
	var tests = []struct {
		input    Tuple
		expected Tuple
	}{
		{
			input:    Tuple{-1.0, 1.0, 1.0, 0.0},
			expected: Tuple{1.0, -1.0, -1.0, 0.0},
		},
		{
			input:    Tuple{-1.345, 3.45, -31.45, 1.0},
			expected: Tuple{1.345, -3.45, 31.45, -1.0},
		},
	}

	for _, test := range tests {
		if got := test.input.Negate(); !got.Equal(test.expected) {
			t.Errorf("Negating: %s\n got: %s. \nexpected: %s", test.input, got, test.expected)
		}
	}
}

func TestMagnitude(t *testing.T) {
	var tests = []struct {
		input    Tuple
		expected float64
	}{
		{
			input:    Tuple{0.0, 1.0, 0.0, 0.0},
			expected: 1.0,
		},
		{
			input:    Tuple{0.0, 0.0, 1.0, 0.0},
			expected: 1.0,
		},
		{
			input:    Tuple{1.0, 2.0, 3.0, 0.0},
			expected: math.Sqrt(14.0),
		},
		{
			input:    Tuple{-1.0, -2.0, -3.0, 0.0},
			expected: math.Sqrt(14.0),
		},
	}

	for _, test := range tests {
		if got := test.input.Magnitude(); !utils.FloatEquals(got, test.expected) {
			t.Errorf("Magnitude of %s \ngot: %f. \nexpected: %f", test.input, got, test.expected)
		}
	}
}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		input    Tuple
		expected Tuple
	}{
		{
			input:    Tuple{4.0, 0.0, 0.0, 0.0},
			expected: Tuple{1.0, 0.0, 0.0, 0.0},
		},
		{
			input: Tuple{1.0, 2.0, 3.0, 0.0},
			// (1/√14, 2/√14, 3/√14, 0)
			expected: Tuple{0.2672612419124244, 0.5345224838248488, 0.8017837257372732, 0.0},
		},
	}

	for _, test := range tests {
		if got := test.input.Normalize(); !got.Equal(test.expected) {
			t.Errorf("Normalizing %s \ngot: %s. \nexpected: %s", test.input, got, test.expected)
		}
	}
}

func TestDot(t *testing.T) {
	var tests = []struct {
		left     Tuple
		right    Tuple
		expected float64
	}{
		{
			left:     Tuple{1.0, 2.0, 3.0, 0.0},
			right:    Tuple{2.0, 3.0, 4.0, 0.0},
			expected: 20.0,
		},
	}

	for _, test := range tests {
		if got := Dot(test.left, test.right); !utils.FloatEquals(got, test.expected) {
			t.Errorf("Dot product of %s and %s\ngot: %f. \nexpected: %f", test.left, test.right, got, test.expected)
		}
	}
}

func TestCross(t *testing.T) {
	var tests = []struct {
		left     Tuple
		right    Tuple
		expected Tuple
	}{
		{
			left:     Tuple{1.0, 2.0, 3.0, 0.0},
			right:    Tuple{2.0, 3.0, 4.0, 0.0},
			expected: Tuple{-1.0, 2.0, -1.0, 0.0},
		},
		{
			left:     Tuple{2.0, 3.0, 4.0, 0.0},
			right:    Tuple{1.0, 2.0, 3.0, 0.0},
			expected: Tuple{1.0, -2.0, 1.0, 0.0},
		},
	}

	for _, test := range tests {
		if got := Cross(test.left, test.right); !got.Equal(test.expected) {
			t.Errorf("Cross product of %s and %s\ngot: %f. \nexpected: %f", test.left, test.right, got, test.expected)
		}
	}
}

func TestReflect(t *testing.T) {
	var tests = []struct {
		in       Tuple
		normal   Tuple
		expected Tuple
	}{
		{
			in:       Tuple{1.0, -1.0, 0.0, 0.0},
			normal:   Tuple{0.0, 1.0, 0.0, 0.0},
			expected: Tuple{1.0, 1.0, 0.0, 0.0},
		},
		{
			in:       Tuple{0.0, -1.0, 0.0, 0.0},
			normal:   Tuple{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0.0, 0.0},
			expected: Tuple{1.0, 0.0, 0.0, 0.0},
		},
	}

	for _, test := range tests {
		if got := Reflect(test.in, test.normal); !got.Equal(test.expected) {
			t.Errorf("Reflect of %s and %s\ngot: %s. \nexpected: %s", test.in, test.normal, got, test.expected)
		}
	}
}
