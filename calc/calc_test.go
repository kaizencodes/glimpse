package calc

import "testing"

func TestFloatEquals(t *testing.T) {
    var tests = []struct {
        left  float64
        right float64
        want  bool
    }{
        {1.000000001, 1.000000002, true},
        {1.0000001, 1.0000002, false},
    }

    for _, test := range tests {
        if got := FloatEquals(test.left, test.right); got != test.want {
            t.Errorf("%f == %f got %t, instead of %t", test.left, test.right, got, test.want)
        }
    }
}
