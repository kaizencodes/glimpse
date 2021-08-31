package matrix

import (
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
                []Element{0, 0},
                []Element{0, 0},
                []Element{0, 0},
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
                []Element{1, 0, 0},
                []Element{0, 1, 0},
                []Element{0, 0, 1},
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
                []Element{1, 2, 3},
                []Element{4, 5, 6},
                []Element{7, 8, 9},
            },
            b: Matrix{
                []Element{1, 2, 3},
                []Element{1, 2, 3},
                []Element{1, 2, 3},
            },
            want: Matrix{
                []Element{6, 12, 18},
                []Element{15, 30, 45},
                []Element{24, 48, 72},
            },
        },
        {
            a: Matrix{
                []Element{1, 2, 3},
                []Element{4, 5, 6},
            },
            b: Matrix{
                []Element{1, 2},
                []Element{1, 2},
                []Element{1, 2},
            },
            want: Matrix{
                []Element{6, 12},
                []Element{15, 30},
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
                []Element{1, 2, 3},
                []Element{4, 5, 6},
                []Element{7, 8, 9},
            },
            b: Matrix{
                []Element{1, 2, 3, 4},
                []Element{1, 2, 3, 4},
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
                []Element{0, 9, 3, 0},
                []Element{9, 6, 0, 8},
                []Element{1, 8, 2, 3},
                []Element{0, 0, 5, 4},
            },
            want: Matrix{
                []Element{0, 9, 1, 0},
                []Element{9, 6, 8, 0},
                []Element{3, 0, 2, 5},
                []Element{0, 8, 3, 4},
            },
        },
        {
            a: Matrix{
                []Element{1, 0, 0, 0},
                []Element{0, 1, 0, 0},
                []Element{0, 0, 1, 0},
                []Element{0, 0, 0, 1},
            },
            want: Matrix{
                []Element{1, 0, 0, 0},
                []Element{0, 1, 0, 0},
                []Element{0, 0, 1, 0},
                []Element{0, 0, 0, 1},
            },
        },
        {
            a: Matrix{
                []Element{1, 7},
                []Element{9, 6},
                []Element{4, 8},
                []Element{2, 3},
            },
            want: Matrix{
                []Element{1, 9, 4, 2},
                []Element{7, 6, 8, 3},
            },
        },
        {
            a: Matrix{
                []Element{1, 9, 4, 2},
            },
            want: Matrix{
                []Element{1},
                []Element{9},
                []Element{4},
                []Element{2},
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
        want Element
    }{
        {
            a: Matrix{
                []Element{1, 5},
                []Element{-3, 2},
            },
            want: 17,
        },
        {
            a: Matrix{
                []Element{1, 2, 6},
                []Element{-5, 8, -4},
                []Element{2, 6, 4},
            },
            want: -196,
        },
        {
            a: Matrix{
                []Element{-2, -8, 3, 5},
                []Element{-3, 1, 7, 3},
                []Element{1, 2, -9, 6},
                []Element{-6, 7, 7, -9},
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
                []Element{1, 5, 0},
                []Element{-3, 2, 7},
                []Element{0, 6, -3},
            },
            col: 0,
            row: 2,
            want: Matrix{
                []Element{-3, 2},
                []Element{0, 6},
            },
        },
        {
            a: Matrix{
                []Element{-6, 1, 1, 6},
                []Element{-8, 5, 8, 6},
                []Element{-1, 0, 8, 2},
                []Element{-7, 1, -1, 1},
            },
            col: 2,
            row: 1,
            want: Matrix{
                []Element{-6, 1, 6},
                []Element{-8, 8, 6},
                []Element{-7, -1, 1},
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
        want Element
    }{
        {
            a: Matrix{
                []Element{3, 5, 0},
                []Element{2, -1, -7},
                []Element{6, -1, 5},
            },
            col:  1,
            row:  0,
            want: 25,
        },
        {
            a: Matrix{
                []Element{3, 5, 0},
                []Element{2, -1, -7},
                []Element{6, -1, 5},
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
        want Element
    }{
        {
            a: Matrix{
                []Element{3, 5, 0},
                []Element{2, -1, -7},
                []Element{6, -1, 5},
            },
            col:  0,
            row:  2,
            want: 4,
        },
        {
            a: Matrix{
                []Element{3, 5, 0},
                []Element{2, -1, -7},
                []Element{6, -1, 5},
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
                []Element{-5, 2, 6, -8},
                []Element{1, -5, 1, 8},
                []Element{7, 7, -6, -7},
                []Element{1, -3, 7, 4},
            },
            want: Matrix{
                []Element{0.21804511278195488, 0.45112781954887216, 0.24060150375939848, -0.045112781954887216},
                []Element{-0.8082706766917294, -1.4567669172932332, -0.44360902255639095, 0.5206766917293233},
                []Element{-0.07894736842105263, -0.2236842105263158, -0.05263157894736842, 0.19736842105263158},
                []Element{-0.5225563909774437, -0.8139097744360902, -0.3007518796992481, 0.30639097744360905},
            },
        },
        {
            a: Matrix{
                []Element{8, -5, 9, 2},
                []Element{7, 5, 6, 1},
                []Element{-6, 0, 9, 6},
                []Element{-3, 0, -9, -4},
            },
            want: Matrix{
                []Element{-0.15384615384615385, -0.15384615384615385, -0.28205128205128205, -0.5384615384615384},
                []Element{-0.07692307692307693, 0.12307692307692308, 0.02564102564102564, 0.03076923076923077},
                []Element{0.358974358974359, 0.358974358974359, 0.4358974358974359, 0.9230769230769231},
                []Element{-0.6923076923076923, -0.6923076923076923, -0.7692307692307693, -1.9230769230769231},
            },
        },
        {
            a: Matrix{
                []Element{9, 3, 0, 9},
                []Element{-5, -2, -6, -3},
                []Element{-4, 9, 6, 4},
                []Element{-7, 6, 6, 2},
            },
            want: Matrix{
                []Element{-0.040740740740740744, -0.07777777777777778, 0.14444444444444443, -0.2222222222222222},
                []Element{-0.07777777777777778, 0.03333333333333333, 0.36666666666666664, -0.3333333333333333},
                []Element{-0.029012345679012345, -0.14629629629629629, -0.10925925925925926, 0.12962962962962962},
                []Element{0.17777777777777778, 0.06666666666666667, -0.26666666666666666, 0.3333333333333333},
            },
        },
    }

    for _, test := range tests {
        if got, _ := test.a.Inverse(); got.String() != test.want.String() {
            t.Errorf("matrix inverse,\na:\n%s\ngot:\n%s\nexpected: \n%s", test.a, got, test.want)
        }
    }

    non_invertible := Matrix{
        []Element{-4, 2, -2, -3},
        []Element{9, 6, 2, 6},
        []Element{0, -5, 1, -5},
        []Element{0, 0, 0, 0},
    }

    if _, error := non_invertible.Inverse(); error == nil {
        t.Errorf("Should return non invertible matrix error")
    }
}
