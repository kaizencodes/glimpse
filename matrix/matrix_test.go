package matrix

import (
    "testing"
)

func TestNew(t *testing.T) {
    var tests = []struct {
        s    int
        want Matrix
    }{
        {
            s: 3,
            want: Matrix{
                []float64{0, 0, 0},
                []float64{0, 0, 0},
                []float64{0, 0, 0},
            },
        },
    }

    for _, test := range tests {
        if got := New(test.s); got.String() != test.want.String() {
            t.Errorf("matrix size:%d, got: \n%s. \nexpected: \n%s", test.s, got, test.want)
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
        b    Matrix
        want Matrix
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
    }

    for _, test := range tests {
        if got, _ := Multiply(test.a, test.b); got.String() != test.want.String() {
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
    }

    for _, test := range tests {
        if got := Transpose(test.a); got.String() != test.want.String() {
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
        if got := Determinant(test.a); got != test.want {
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
        if got := Submatrix(test.a, test.col, test.row); got.String() != test.want.String() {
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
        if got := Minor(test.a, test.col, test.row); got != test.want {
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
        if got := Cofactor(test.a, test.col, test.row); got != test.want {
            t.Errorf("Cofactor,\na:\n%s\n col: %d\n row: %d\ngot: %f\nexpected: %f", test.a, test.col, test.row, got, test.want)
        }
    }
}
