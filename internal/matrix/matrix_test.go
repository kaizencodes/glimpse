package matrix

import (
	"testing"
)

func TestNewEmpty(t *testing.T) {
	var tests = []struct {
		n, m     int
		expected Matrix
	}{
		{
			n: 3, m: 2,
			expected: Matrix{
				data: []float64{
					0, 0,
					0, 0,
					0, 0,
				},
				row_size: 4,
				col_size: 2,
			},
		},
	}

	for _, test := range tests {
		if got := NewEmpty(test.n, test.m); got.String() != test.expected.String() {
			t.Errorf("matrix size:%d, %d. got: \n%s. \nexpected: \n%s", test.n, test.m, got, test.expected)
		}
	}
}

func TestNewIdentity(t *testing.T) {
	var tests = []struct {
		s        int
		expected Matrix
	}{
		{
			s: 3,
			expected: Matrix{
				data: []float64{
					1, 0, 0,
					0, 1, 0,
					0, 0, 1,
				},
				row_size: 3,
				col_size: 3,
			},
		},
	}

	for _, test := range tests {
		if got := NewIdentity(test.s); got.String() != test.expected.String() {
			t.Errorf("matrix size:%d, got: \n%s. \nexpected: \n%s", test.s, got, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	var tests = []struct {
		a        Matrix
		b        Matrix
		expected Matrix
	}{
		{
			a: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
				row_size: 3,
				col_size: 3,
			},
			b: Matrix{
				data: []float64{
					1, 2, 3,
					1, 2, 3,
					1, 2, 3,
				},
				row_size: 3,
				col_size: 3,
			},
			expected: Matrix{
				data: []float64{
					6, 12, 18,
					15, 30, 45,
					24, 48, 72,
				},
				row_size: 3,
				col_size: 3,
			},
		},
		{
			a: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
				},
				row_size: 2,
				col_size: 3,
			},
			b: Matrix{
				data: []float64{
					1, 2,
					1, 2,
					1, 2,
				},
				row_size: 3,
				col_size: 2,
			},
			expected: Matrix{
				data: []float64{
					6, 12,
					15, 30,
				},
				row_size: 2,
				col_size: 2,
			},
		},
	}

	for _, test := range tests {
		got := Multiply(test.a, test.b)
		if got.String() != test.expected.String() {
			t.Errorf("matrix multiplication,\na:\n%s\nb:\n%s\ngot:\n%s\nexpected: \n%s", test.a, test.b, got, test.expected)
		}
	}
}

func TestInvalidMultiply(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic for invalid multiplication")
		}
	}()

	var tests = []struct {
		a Matrix
		b Matrix
	}{
		{
			a: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
				row_size: 3,
				col_size: 3,
			},
			b: Matrix{
				data: []float64{
					1, 2, 3, 4,
					1, 2, 3, 4,
				},
				row_size: 2,
				col_size: 4,
			},
		},
	}

	for _, test := range tests {
		Multiply(test.a, test.b)
	}
}

func TestTranspose(t *testing.T) {
	var tests = []struct {
		a        Matrix
		expected Matrix
	}{
		{
			a: Matrix{
				data: []float64{
					0, 9, 3, 0,
					9, 6, 0, 8,
					1, 8, 2, 3,
					0, 0, 5, 4,
				},
				row_size: 4,
				col_size: 4,
			},
			expected: Matrix{
				data: []float64{
					0, 9, 1, 0,
					9, 6, 8, 0,
					3, 0, 2, 5,
					0, 8, 3, 4,
				},
				row_size: 4,
				col_size: 4,
			},
		},
		{
			a: Matrix{
				data: []float64{
					1, 0, 0, 0,
					0, 1, 0, 0,
					0, 0, 1, 0,
					0, 0, 0, 1,
				},
				row_size: 4,
				col_size: 4,
			},
			expected: Matrix{
				data: []float64{
					1, 0, 0, 0,
					0, 1, 0, 0,
					0, 0, 1, 0,
					0, 0, 0, 1,
				},
				row_size: 4,
				col_size: 4,
			},
		},
		{
			a: Matrix{
				data: []float64{
					1, 7,
					9, 6,
					4, 8,
					2, 3,
				},
				row_size: 4,
				col_size: 2,
			},
			expected: Matrix{
				data: []float64{
					1, 9, 4, 2,
					7, 6, 8, 3,
				},
				row_size: 2,
				col_size: 4,
			},
		},
		{
			a: Matrix{
				data:     []float64{1, 9, 4, 2},
				row_size: 1,
				col_size: 4,
			},
			expected: Matrix{
				data: []float64{
					1,
					9,
					4,
					2,
				},
				row_size: 4,
				col_size: 1,
			},
		},
	}

	for _, test := range tests {
		if got := test.a.Transpose(); got.String() != test.expected.String() {
			t.Errorf("matrix transposition,\na:\n%s\ngot:\n%s\nexpected: \n%s", test.a, got, test.expected)
		}
	}
}
