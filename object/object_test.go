package object

import "testing"

func TestTouple(t *testing.T) {
    var tests = []struct {
        input *Touple
        want  bool
    }{
        {
            input: &Touple{4.3, -4.2, 3.1, 1.0},
            want:  true,
        },
        {
            input: &Touple{-3.3, 3.2, 3.1, 0.0},
            want:  false,
        },
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

func TestEqual(t *testing.T) {
    var tests = []struct {
        left  *Touple
        right *Touple
        want  bool
    }{
        {
            left:  &Touple{1.0, 1.0, 1.0, 1.0},
            right: &Touple{1.0, 1.0, 1.0, 1.0},
            want:  true,
        },
        {
            left:  &Touple{1.0, 1.0, 1.0, 1.0},
            right: &Touple{0.0, 1.0, 1.0, 1.0},
            want:  false,
        },
        {
            left:  &Touple{1.0, 1.0, 1.0, 1.0},
            right: &Touple{1.0, 0.0, 1.0, 1.0},
            want:  false,
        },
        {
            left:  &Touple{1.0, 1.0, 1.0, 1.0},
            right: &Touple{1.0, 1.0, 0.0, 1.0},
            want:  false,
        },
        {
            left:  &Touple{1.0, 1.0, 1.0, 1.0},
            right: &Touple{1.0, 1.0, 1.0, 0.0},
            want:  false,
        },
    }

    for _, test := range tests {
        if got := test.left.Equal(test.right); got != test.want {
            t.Errorf("expected touples to be equal, but they were not.")
        }
    }
}

func TestAdd(t *testing.T) {
    var tests = []struct {
        left  *Touple
        right *Touple
        want  *Touple
    }{
        {
            left:  &Touple{1.0, 1.0, 1.0, 1.0},
            right: &Touple{0.0, 1.5, -1.0, 0.0},
            want:  &Touple{1.0, 2.5, 0.0, 1.0},
        },
        {
            left:  &Touple{-1.0, 1.0, 1.0, 0.0},
            right: &Touple{-2.0, 1.5, 0.0001, 0.0},
            want:  &Touple{-3.0, 2.5, 1.0001, 0.0},
        },
        {
            left:  &Touple{1.0, 1.0, 1.0, 0.0},
            right: &Touple{0.0, 1.5, -1.0, 1.0},
            want:  &Touple{1.0, 2.5, 0.0, 1.0},
        },
    }

    for _, test := range tests {
        if got := Add(test.left, test.right); !got.Equal(test.want) {
            t.Errorf("input: %s + %s \ngot: %s. \nexpected: %s", test.left, test.right, got, test.want)
        }
    }
}

func TestSubtract(t *testing.T) {
    var tests = []struct {
        left  *Touple
        right *Touple
        want  *Touple
    }{
        {
            left:  &Touple{1.0, 1.0, 1.0, 1.0},
            right: &Touple{0.0, 1.5, -1.0, 0.0},
            want:  &Touple{1.0, -0.5, 2.0, 1.0},
        },
        {
            left:  &Touple{-1.0, 1.0, 1.0, 0.0},
            right: &Touple{-2.0, 1.5, 0.0001, 0.0},
            want:  &Touple{1.0, -0.5, 0.9999, 0.0},
        },
        {
            left:  &Touple{1.0, 1.0, 1.0, 1.0},
            right: &Touple{0.0, 1.5, -1.0, 1.0},
            want:  &Touple{1.0, -0.5, 2.0, 0.0},
        },
    }

    for _, test := range tests {
        if got := Subtract(test.left, test.right); !got.Equal(test.want) {
            t.Errorf("input: %s - %s \ngot: %s. \nexpected: %s", test.left, test.right, got, test.want)
        }
    }
}

func TestNegate(t *testing.T) {
    var tests = []struct {
        input *Touple
        want  *Touple
    }{
        {
            input: &Touple{-1.0, 1.0, 1.0, 0.0},
            want:  &Touple{1.0, -1.0, -1.0, 0.0},
        },
        {
            input: &Touple{-1.345, 3.45, -31.45, 1.0},
            want:  &Touple{1.345, -3.45, 31.45, -1.0},
        },
    }

    for _, test := range tests {
        if got := Negate(test.input); !got.Equal(test.want) {
            t.Errorf("Negating: %s\n got: %s. \nexpected: %s", test.input, got, test.want)
        }
    }
}
