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

