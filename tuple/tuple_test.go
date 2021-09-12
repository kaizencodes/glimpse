package tuple

import (
    "glimpse/calc"
    "glimpse/matrix"
    "math"
    "testing"
)

func TestTuple(t *testing.T) {
    var tests = []struct {
        input Tuple
        want  bool
    }{
        {
            input: Tuple{4.3, -4.2, 3.1, 1.0},
            want:  true,
        },
        {
            input: Tuple{-3.3, 3.2, 3.1, 0.0},
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
        left  Tuple
        right Tuple
        want  bool
    }{
        {
            left:  Tuple{1.0, 1.0, 1.0, 1.0},
            right: Tuple{1.0, 1.0, 1.0, 1.0},
            want:  true,
        },
        {
            left:  Tuple{1.0, 1.0, 1.0, 1.0},
            right: Tuple{0.0, 1.0, 1.0, 1.0},
            want:  false,
        },
        {
            left:  Tuple{1.0, 1.0, 1.0, 1.0},
            right: Tuple{1.0, 0.0, 1.0, 1.0},
            want:  false,
        },
        {
            left:  Tuple{1.0, 1.0, 1.0, 1.0},
            right: Tuple{1.0, 1.0, 0.0, 1.0},
            want:  false,
        },
        {
            left:  Tuple{1.0, 1.0, 1.0, 1.0},
            right: Tuple{1.0, 1.0, 1.0, 0.0},
            want:  false,
        },
    }

    for _, test := range tests {
        if got := test.left.Equal(test.right); got != test.want {
            t.Errorf("expected tuples to be equal, but they were not.")
        }
    }
}

func TestAdd(t *testing.T) {
    var tests = []struct {
        left  Tuple
        right Tuple
        want  Tuple
    }{
        {
            left:  Tuple{1.0, 1.0, 1.0, 1.0},
            right: Tuple{0.0, 1.5, -1.0, 0.0},
            want:  Tuple{1.0, 2.5, 0.0, 1.0},
        },
        {
            left:  Tuple{-1.0, 1.0, 1.0, 0.0},
            right: Tuple{-2.0, 1.5, 0.0001, 0.0},
            want:  Tuple{-3.0, 2.5, 1.0001, 0.0},
        },
        {
            left:  Tuple{1.0, 1.0, 1.0, 0.0},
            right: Tuple{0.0, 1.5, -1.0, 1.0},
            want:  Tuple{1.0, 2.5, 0.0, 1.0},
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
        left  Tuple
        right Tuple
        want  Tuple
    }{
        {
            left:  Tuple{1.0, 1.0, 1.0, 1.0},
            right: Tuple{0.0, 1.5, -1.0, 0.0},
            want:  Tuple{1.0, -0.5, 2.0, 1.0},
        },
        {
            left:  Tuple{-1.0, 1.0, 1.0, 0.0},
            right: Tuple{-2.0, 1.5, 0.0001, 0.0},
            want:  Tuple{1.0, -0.5, 0.9999, 0.0},
        },
        {
            left:  Tuple{1.0, 1.0, 1.0, 1.0},
            right: Tuple{0.0, 1.5, -1.0, 1.0},
            want:  Tuple{1.0, -0.5, 2.0, 0.0},
        },
    }

    for _, test := range tests {
        if got := Subtract(test.left, test.right); !got.Equal(test.want) {
            t.Errorf("input: %s - %s \ngot: %s. \nexpected: %s", test.left, test.right, got, test.want)
        }
    }
}

func TestMultiply(t *testing.T) {
    var tests = []struct {
        a    matrix.Matrix
        b    Tuple
        want Tuple
    }{
        {
            a: matrix.Matrix{
                []float64{1, 2, 3, 4},
                []float64{2, 4, 4, 2},
                []float64{8, 6, 4, 1},
                []float64{0, 0, 0, 1},
            },
            b:    Tuple{1, 2, 3, 1},
            want: Tuple{18, 24, 33, 1},
        },
    }

    for _, test := range tests {
        got, _ := Multiply(test.a, test.b)
        if got.String() != test.want.String() {
            t.Errorf("multiplication,\na:\n%s\nb:\n%s\ngot:\n%s\nexpected: \n%s", test.a, test.b, got, test.want)
        }
    }
}

func TestScalar(t *testing.T) {
    var tests = []struct {
        tuple  Tuple
        scalar float64
        want   Tuple
    }{
        {
            tuple:  Tuple{1.0, 1.0, 1.0, 1.0},
            scalar: 1,
            want:   Tuple{1.0, 1.0, 1.0, 1.0},
        },
        {
            tuple:  Tuple{-2.0, 1.5, 0.5, 0.0},
            scalar: 0.5,
            want:   Tuple{-1.0, 0.75, 0.25, 0.0},
        },
        {
            tuple:  Tuple{-2.0, 1.5, 0.5, 0.0},
            scalar: 2,
            want:   Tuple{-4.0, 3, 1.0, 0.0},
        },
    }

    for _, test := range tests {
        if got := test.tuple.Scalar(test.scalar); !got.Equal(test.want) {
            t.Errorf("input: %s + %f \ngot: %s. \nexpected: %s", test.tuple, test.scalar, got, test.want)
        }
    }
}

func TestNegate(t *testing.T) {
    var tests = []struct {
        input Tuple
        want  Tuple
    }{
        {
            input: Tuple{-1.0, 1.0, 1.0, 0.0},
            want:  Tuple{1.0, -1.0, -1.0, 0.0},
        },
        {
            input: Tuple{-1.345, 3.45, -31.45, 1.0},
            want:  Tuple{1.345, -3.45, 31.45, -1.0},
        },
    }

    for _, test := range tests {
        if got := test.input.Negate(); !got.Equal(test.want) {
            t.Errorf("Negating: %s\n got: %s. \nexpected: %s", test.input, got, test.want)
        }
    }
}

func TestMagnitude(t *testing.T) {
    var tests = []struct {
        input Tuple
        want  float64
    }{
        {
            input: Tuple{0.0, 1.0, 0.0, 0.0},
            want:  1.0,
        },
        {
            input: Tuple{0.0, 0.0, 1.0, 0.0},
            want:  1.0,
        },
        {
            input: Tuple{1.0, 2.0, 3.0, 0.0},
            want:  math.Sqrt(14.0),
        },
    }

    for _, test := range tests {
        if got := test.input.Magnitude(); !calc.FloatEquals(got, test.want) {
            t.Errorf("Magnitude of %s \ngot: %f. \nexpected: %f", test.input, got, test.want)
        }
    }
}

func TestNormalize(t *testing.T) {
    var tests = []struct {
        input Tuple
        want  Tuple
    }{
        {
            input: Tuple{4.0, 0.0, 0.0, 0.0},
            want:  Tuple{1.0, 0.0, 0.0, 0.0},
        },
        {
            input: Tuple{1.0, 2.0, 3.0, 0.0},
            want:  Tuple{0.2672612419124244, 0.5345224838248488, 0.8017837257372732, 0.0},
        },
    }

    for _, test := range tests {
        if got := test.input.Normalize(); !got.Equal(test.want) {
            t.Errorf("Normalizing %s \ngot: %s. \nexpected: %s", test.input, got, test.want)
        }
    }
}

func TestDot(t *testing.T) {
    var tests = []struct {
        left  Tuple
        right Tuple
        want  float64
    }{
        {
            left:  Tuple{1.0, 2.0, 3.0, 0.0},
            right: Tuple{2.0, 3.0, 4.0, 0.0},
            want:  20.0,
        },
    }

    for _, test := range tests {
        if got := Dot(test.left, test.right); !calc.FloatEquals(got, test.want) {
            t.Errorf("Dot product of %s and %s\ngot: %f. \nexpected: %f", test.left, test.right, got, test.want)
        }
    }
}

func TestCross(t *testing.T) {
    var tests = []struct {
        left  Tuple
        right Tuple
        want  Tuple
    }{
        {
            left:  Tuple{1.0, 2.0, 3.0, 0.0},
            right: Tuple{2.0, 3.0, 4.0, 0.0},
            want:  Tuple{-1.0, 2.0, -1.0, 0.0},
        },
        {
            left:  Tuple{2.0, 3.0, 4.0, 0.0},
            right: Tuple{1.0, 2.0, 3.0, 0.0},
            want:  Tuple{1.0, -2.0, 1.0, 0.0},
        },
    }

    for _, test := range tests {
        if got := Cross(test.left, test.right); !got.Equal(test.want) {
            t.Errorf("Cross product of %s and %s\ngot: %f. \nexpected: %f", test.left, test.right, got, test.want)
        }
    }
}

func TestReflect(t *testing.T) {
    var tests = []struct {
        in     Tuple
        normal Tuple
        want   Tuple
    }{
        {
            in:     Tuple{1.0, -1.0, 0.0, 0.0},
            normal: Tuple{0.0, 1.0, 0.0, 0.0},
            want:   Tuple{1.0, 1.0, 0.0, 0.0},
        },
        {
            in:     Tuple{0.0, -1.0, 0.0, 0.0},
            normal: Tuple{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0.0, 0.0},
            want:   Tuple{1.0, 0.0, 0.0, 0.0},
        },
    }

    for _, test := range tests {
        if got := Reflect(test.in, test.normal); !got.Equal(test.want) {
            t.Errorf("Reflect of %s and %s\ngot: %s. \nexpected: %s", test.in, test.normal, got, test.want)
        }
    }
}
