package matrix

import (
    "glimpse/tuple"
    "testing"
)

func TestNew(t *testing.T) {
    var tests = []struct {
        n, m int
        want Matrix
    }{
        {
            n: 3, m: 2,
            want: Matrix{
                []float64{0, 0},
                []float64{0, 0},
                []float64{0, 0},
            },
        },
    }

    for _, test := range tests {
        if got := New(test.n, test.m); got.String() != test.want.String() {
            t.Errorf("matrix size:%d, %d. got: \n%s. \nexpected: \n%s", test.n, test.m, got, test.want)
        }
    }
}

func TestNewIdentity(t *testing.T) {
    var tests = []struct {
        s    int
        want Matrix
    }{
        {
            s: 3,
            want: Matrix{
                []float64{1, 0, 0},
                []float64{0, 1, 0},
                []float64{0, 0, 1},
            },
        },
    }

    for _, test := range tests {
        if got := NewIdentity(test.s); got.String() != test.want.String() {
            t.Errorf("matrix size:%d, got: \n%s. \nexpected: \n%s", test.s, got, test.want)
        }
    }
}

func TestMultiply(t *testing.T) {
    var tests = []struct {
        a    Matrix
        b    Transformable
        want Transformable
    }{
        {
            a: Matrix{
                []float64{1, 2, 3},
                []float64{4, 5, 6},
                []float64{7, 8, 9},
            },
            b: Matrix{
                []float64{1, 2, 3},
                []float64{1, 2, 3},
                []float64{1, 2, 3},
            },
            want: Matrix{
                []float64{6, 12, 18},
                []float64{15, 30, 45},
                []float64{24, 48, 72},
            },
        },
        {
            a: Matrix{
                []float64{1, 2, 3},
                []float64{4, 5, 6},
            },
            b: Matrix{
                []float64{1, 2},
                []float64{1, 2},
                []float64{1, 2},
            },
            want: Matrix{
                []float64{6, 12},
                []float64{15, 30},
            },
        },
        {
            a: Matrix{
                []float64{1, 2, 3, 4},
                []float64{2, 4, 4, 2},
                []float64{8, 6, 4, 1},
                []float64{0, 0, 0, 1},
            },
            b:    tuple.Tuple{1, 2, 3, 1},
            want: tuple.Tuple{18, 24, 33, 1},
        },
    }

    for _, test := range tests {
        got, _ := Multiply(test.a, test.b)
        if got.String() != test.want.String() {
            t.Errorf("matrix multiplication,\na:\n%s\nb:\n%s\ngot:\n%s\nexpected: \n%s", test.a, test.b, got, test.want)
        }
    }
}

func TestInvalidMultiply(t *testing.T) {
    var tests = []struct {
        a Matrix
        b Matrix
    }{
        {
            a: Matrix{
                []float64{1, 2, 3},
                []float64{4, 5, 6},
                []float64{7, 8, 9},
            },
            b: Matrix{
                []float64{1, 2, 3, 4},
                []float64{1, 2, 3, 4},
            },
        },
    }

    for _, test := range tests {
        if _, err := Multiply(test.a, test.b); err == nil {
            t.Errorf("should've raised incompatible matrices error")
        }
    }
}

func TestTranspose(t *testing.T) {
    var tests = []struct {
        a    Matrix
        want Matrix
    }{
        {
            a: Matrix{
                []float64{0, 9, 3, 0},
                []float64{9, 6, 0, 8},
                []float64{1, 8, 2, 3},
                []float64{0, 0, 5, 4},
            },
            want: Matrix{
                []float64{0, 9, 1, 0},
                []float64{9, 6, 8, 0},
                []float64{3, 0, 2, 5},
                []float64{0, 8, 3, 4},
            },
        },
        {
            a: Matrix{
                []float64{1, 0, 0, 0},
                []float64{0, 1, 0, 0},
                []float64{0, 0, 1, 0},
                []float64{0, 0, 0, 1},
            },
            want: Matrix{
                []float64{1, 0, 0, 0},
                []float64{0, 1, 0, 0},
                []float64{0, 0, 1, 0},
                []float64{0, 0, 0, 1},
            },
        },
        {
            a: Matrix{
                []float64{1, 7},
                []float64{9, 6},
                []float64{4, 8},
                []float64{2, 3},
            },
            want: Matrix{
                []float64{1, 9, 4, 2},
                []float64{7, 6, 8, 3},
            },
        },
        {
            a: Matrix{
                []float64{1, 9, 4, 2},
            },
            want: Matrix{
                []float64{1},
                []float64{9},
                []float64{4},
                []float64{2},
            },
        },
    }

    for _, test := range tests {
        if got := test.a.Transpose(); got.String() != test.want.String() {
            t.Errorf("matrix transposition,\na:\n%s\ngot:\n%s\nexpected: \n%s", test.a, got, test.want)
        }
    }
}

func TestDeterminant(t *testing.T) {
    var tests = []struct {
        a    Matrix
        want float64
    }{
        {
            a: Matrix{
                []float64{1, 5},
                []float64{-3, 2},
            },
            want: 17,
        },
        {
            a: Matrix{
                []float64{1, 2, 6},
                []float64{-5, 8, -4},
                []float64{2, 6, 4},
            },
            want: -196,
        },
        {
            a: Matrix{
                []float64{-2, -8, 3, 5},
                []float64{-3, 1, 7, 3},
                []float64{1, 2, -9, 6},
                []float64{-6, 7, 7, -9},
            },
            want: -4071,
        },
    }

    for _, test := range tests {
        if got := test.a.Determinant(); got != test.want {
            t.Errorf("matrix determinant,\na:\n%s\ngot: %f\nexpected: %f", test.a, got, test.want)
        }
    }
}

func TestSubmatrix(t *testing.T) {
    var tests = []struct {
        a    Matrix
        col  int
        row  int
        want Matrix
    }{
        {
            a: Matrix{
                []float64{1, 5, 0},
                []float64{-3, 2, 7},
                []float64{0, 6, -3},
            },
            col: 0,
            row: 2,
            want: Matrix{
                []float64{-3, 2},
                []float64{0, 6},
            },
        },
        {
            a: Matrix{
                []float64{-6, 1, 1, 6},
                []float64{-8, 5, 8, 6},
                []float64{-1, 0, 8, 2},
                []float64{-7, 1, -1, 1},
            },
            col: 2,
            row: 1,
            want: Matrix{
                []float64{-6, 1, 6},
                []float64{-8, 8, 6},
                []float64{-7, -1, 1},
            },
        },
    }

    for _, test := range tests {
        if got := test.a.Submatrix(test.col, test.row); got.String() != test.want.String() {
            t.Errorf("submatrix,\na:\n%s\n col: %d\n row: %d\ngot:\n%s\nexpected: \n%s", test.a, test.col, test.row, got, test.want)
        }
    }
}

func TestMinor(t *testing.T) {
    var tests = []struct {
        a    Matrix
        col  int
        row  int
        want float64
    }{
        {
            a: Matrix{
                []float64{3, 5, 0},
                []float64{2, -1, -7},
                []float64{6, -1, 5},
            },
            col:  1,
            row:  0,
            want: 25,
        },
        {
            a: Matrix{
                []float64{3, 5, 0},
                []float64{2, -1, -7},
                []float64{6, -1, 5},
            },
            col:  0,
            row:  2,
            want: 4,
        },
    }

    for _, test := range tests {
        if got := test.a.Minor(test.col, test.row); got != test.want {
            t.Errorf("Minor,\na:\n%s\n col: %d\n row: %d\ngot: %f\nexpected: %f", test.a, test.col, test.row, got, test.want)
        }
    }
}

func TestCofactor(t *testing.T) {
    var tests = []struct {
        a    Matrix
        col  int
        row  int
        want float64
    }{
        {
            a: Matrix{
                []float64{3, 5, 0},
                []float64{2, -1, -7},
                []float64{6, -1, 5},
            },
            col:  0,
            row:  2,
            want: 4,
        },
        {
            a: Matrix{
                []float64{3, 5, 0},
                []float64{2, -1, -7},
                []float64{6, -1, 5},
            },
            col:  1,
            row:  0,
            want: -25,
        },
    }

    for _, test := range tests {
        if got := test.a.Cofactor(test.col, test.row); got != test.want {
            t.Errorf("Cofactor,\na:\n%s\n col: %d\n row: %d\ngot: %f\nexpected: %f", test.a, test.col, test.row, got, test.want)
        }
    }
}

func TestInverse(t *testing.T) {
    var tests = []struct {
        a    Matrix
        want Matrix
    }{
        {
            a: Matrix{
                []float64{-5, 2, 6, -8},
                []float64{1, -5, 1, 8},
                []float64{7, 7, -6, -7},
                []float64{1, -3, 7, 4},
            },
            want: Matrix{
                []float64{0.21804511278195488, 0.45112781954887216, 0.24060150375939848, -0.045112781954887216},
                []float64{-0.8082706766917294, -1.4567669172932332, -0.44360902255639095, 0.5206766917293233},
                []float64{-0.07894736842105263, -0.2236842105263158, -0.05263157894736842, 0.19736842105263158},
                []float64{-0.5225563909774437, -0.8139097744360902, -0.3007518796992481, 0.30639097744360905},
            },
        },
        {
            a: Matrix{
                []float64{8, -5, 9, 2},
                []float64{7, 5, 6, 1},
                []float64{-6, 0, 9, 6},
                []float64{-3, 0, -9, -4},
            },
            want: Matrix{
                []float64{-0.15384615384615385, -0.15384615384615385, -0.28205128205128205, -0.5384615384615384},
                []float64{-0.07692307692307693, 0.12307692307692308, 0.02564102564102564, 0.03076923076923077},
                []float64{0.358974358974359, 0.358974358974359, 0.4358974358974359, 0.9230769230769231},
                []float64{-0.6923076923076923, -0.6923076923076923, -0.7692307692307693, -1.9230769230769231},
            },
        },
        {
            a: Matrix{
                []float64{9, 3, 0, 9},
                []float64{-5, -2, -6, -3},
                []float64{-4, 9, 6, 4},
                []float64{-7, 6, 6, 2},
            },
            want: Matrix{
                []float64{-0.040740740740740744, -0.07777777777777778, 0.14444444444444443, -0.2222222222222222},
                []float64{-0.07777777777777778, 0.03333333333333333, 0.36666666666666664, -0.3333333333333333},
                []float64{-0.029012345679012345, -0.14629629629629629, -0.10925925925925926, 0.12962962962962962},
                []float64{0.17777777777777778, 0.06666666666666667, -0.26666666666666666, 0.3333333333333333},
            },
        },
    }

    for _, test := range tests {
        if got, _ := test.a.Inverse(); got.String() != test.want.String() {
            t.Errorf("matrix inverse,\na:\n%s\ngot:\n%s\nexpected: \n%s", test.a, got, test.want)
        }
    }

    non_invertible := Matrix{
        []float64{-4, 2, -2, -3},
        []float64{9, 6, 2, 6},
        []float64{0, -5, 1, -5},
        []float64{0, 0, 0, 0},
    }

    if _, error := non_invertible.Inverse(); error == nil {
        t.Errorf("Should return non invertible matrix error")
    }
}
