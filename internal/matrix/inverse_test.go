package matrix

import (
	"testing"
)

func TestInverse(t *testing.T) {
	var tests = []struct {
		a        Matrix
		expected Matrix
	}{
		{
			a: Matrix{
				data: []float64{
					-5, 2, 6, -8,
					1, -5, 1, 8,
					7, 7, -6, -7,
					1, -3, 7, 4,
				},
				row_size: 4,
				col_size: 4,
			},
			expected: Matrix{
				data: []float64{
					0.21804511278195488, 0.45112781954887216, 0.24060150375939848, -0.045112781954887216,
					-0.8082706766917294, -1.4567669172932332, -0.44360902255639095, 0.5206766917293233,
					-0.07894736842105263, -0.2236842105263158, -0.05263157894736842, 0.19736842105263158,
					-0.5225563909774437, -0.8139097744360902, -0.3007518796992481, 0.30639097744360905,
				},
				row_size: 4,
				col_size: 4,
			},
		},
		{
			a: Matrix{
				data: []float64{
					8, -5, 9, 2,
					7, 5, 6, 1,
					-6, 0, 9, 6,
					-3, 0, -9, -4,
				},
				row_size: 4,
				col_size: 4,
			},
			expected: Matrix{
				data: []float64{
					-0.15384615384615385, -0.15384615384615385, -0.28205128205128205, -0.5384615384615384,
					-0.07692307692307693, 0.12307692307692308, 0.02564102564102564, 0.03076923076923077,
					0.358974358974359, 0.358974358974359, 0.4358974358974359, 0.9230769230769231,
					-0.6923076923076923, -0.6923076923076923, -0.7692307692307693, -1.9230769230769231,
				},
				row_size: 4,
				col_size: 4,
			},
		},
		{
			a: Matrix{
				data: []float64{
					9, 3, 0, 9,
					-5, -2, -6, -3,
					-4, 9, 6, 4,
					-7, 6, 6, 2,
				},
				row_size: 4,
				col_size: 4,
			},
			expected: Matrix{
				data: []float64{
					-0.040740740740740744, -0.07777777777777778, 0.14444444444444443, -0.2222222222222222,
					-0.07777777777777778, 0.03333333333333333, 0.36666666666666664, -0.3333333333333333,
					-0.029012345679012345, -0.14629629629629629, -0.10925925925925926, 0.12962962962962962,
					0.17777777777777778, 0.06666666666666667, -0.26666666666666666, 0.3333333333333333,
				},
				row_size: 4,
				col_size: 4,
			},
		},
	}

	for _, test := range tests {
		if result := test.a.Inverse(); result.String() != test.expected.String() {
			t.Errorf("matrix inverse,\na:\n%s\ngot:\n%s\nexpected: \n%s", test.a, result, test.expected)
		}
	}
}

func TestNonInvertibleMatrix(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic for non-invertible matrix")
		}
	}()

	Matrix{
		data: []float64{
			-4, 2, -2, -3,
			9, 6, 2, 6,
			0, -5, 1, -5,
			0, 0, 0, 0,
		},
		row_size: 4,
		col_size: 4,
	}.Inverse()

}

func TestDeterminant(t *testing.T) {
	var tests = []struct {
		m        Matrix
		expected float64
	}{
		{
			m: Matrix{
				data: []float64{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 10, 11, 12,
					13, 14, 15, 16,
				},
				row_size: 4,
				col_size: 4,
			},
			expected: 0,
		},
		{
			m: Matrix{
				data: []float64{
					1, 0, 0, 0,
					0, 1, 0, 0,
					0, 0, 1, 0,
					0, 0, 0, 1,
				},
				row_size: 4,
				col_size: 4,
			},
			expected: 1,
		},
		{
			m: Matrix{
				data: []float64{
					-2, -8, 3, 5,
					-3, 1, 7, 3,
					1, 2, -9, 6,
					-6, 7, 7, -9,
				},
				row_size: 4,
				col_size: 4,
			},
			expected: -4071,
		},
		{
			m: Matrix{
				data: []float64{
					2, 4, 1, 3,
					1, 2, 2, 1,
					3, 4, 4, 2,
					1, 2, 1, 3,
				},
				row_size: 4,
				col_size: 4,
			},
			expected: 10,
		},
	}

	for _, test := range tests {
		if got := determinant(test.m); got != test.expected {
			t.Errorf("matrix determinant,\na:\n%s\ngot: %f\nexpected: %f", test.m, got, test.expected)
		}
	}
}

func TestSubDet(t *testing.T) {
	var tests = []struct {
		m        Matrix
		row, col int
		expected float64
	}{
		{
			m: Matrix{
				data: []float64{
					1, 2, 3, 4,
					5, 6, 7, 8,
					9, 10, 11, 12,
					13, 14, 15, 16,
				},
				row_size: 4,
				col_size: 4,
			},
			row:      0,
			col:      0,
			expected: 0,
		},
		{
			m: Matrix{
				data: []float64{
					6, 1, 1, 3,
					4, -2, 5, 1,
					2, 8, 7, 6,
					3, 1, 9, 7,
				},
				row_size: 4,
				col_size: 4,
			},
			row:      2,
			col:      2,
			expected: -85,
		},
	}

	for _, test := range tests {
		if got := subDet(test.m, test.row, test.col); got != test.expected {
			t.Errorf("matrix determinant,\na:\n%s\ngot: %f\nexpected: %f", test.m, got, test.expected)
		}
	}
}
