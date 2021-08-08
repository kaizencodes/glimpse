package object

import "testing"

func TestTouple(t *testing.T) {
    var tests = []struct {
        input *Touple
        want  bool
    }{
        {&Touple{4.3, -4.2, 3.1, 1.0}, true},
        {&Touple{-3.3, 3.2, 3.1, 0.0}, false},
    }

    for _, test := range tests {
        if got := test.input.IsPoint(); got != test.want {
            t.Errorf("expected IsPoint() to be %t, got %t", got, test.want)
        }
    }
}

func TestNewVector(t *testing.T) {
    test := NewVector(1, 2, 3)

    if test.IsPoint() {
        t.Errorf("expected to get vector, got point")
    }
}

func TestNewPoint(t *testing.T) {
    test := NewPoint(1, 2, 3)

    if test.IsVector() {
        t.Errorf("expected to get point, got vector")
    }
}
