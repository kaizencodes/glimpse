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
