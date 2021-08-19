package color

import "testing"

func TestAdd(t *testing.T) {
    var tests = []struct {
        left  *Color
        right *Color
        want  *Color
    }{
        {
            left:  &Color{1.0, 1.0, 1.0},
            right: &Color{0.0, 1.5, -1.0},
            want:  &Color{1.0, 2.5, 0.0},
        },
        {
            left:  &Color{-1.0, 1.0, 1.0},
            right: &Color{-2.0, 1.5, 0.0001},
            want:  &Color{-3.0, 2.5, 1.0001},
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
        left  *Color
        right *Color
        want  *Color
    }{
        {
            left:  &Color{1.0, 1.0, 1.0},
            right: &Color{0.0, 1.5, -1.0},
            want:  &Color{1.0, -0.5, 2.0},
        },
        {
            left:  &Color{-1.0, 1.0, 1.0},
            right: &Color{-2.0, 1.5, 0.0001},
            want:  &Color{1.0, -0.5, 0.9999},
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
        tuple  *Color
        scalar float64
        want   *Color
    }{
        {
            tuple:  &Color{1.0, 1.0, 1.0},
            scalar: 1,
            want:   &Color{1.0, 1.0, 1.0},
        },
        {
            tuple:  &Color{-2.0, 1.5, 0.5},
            scalar: 0.5,
            want:   &Color{-1.0, 0.75, 0.25},
        },
        {
            tuple:  &Color{-2.0, 1.5, 0.5},
            scalar: 2,
            want:   &Color{-4.0, 3, 1.0},
        },
    }

    for _, test := range tests {
        if got := Multiply(test.tuple, test.scalar); !got.Equal(test.want) {
            t.Errorf("input: %s + %f \ngot: %s. \nexpected: %s", test.tuple, test.scalar, got, test.want)
        }
    }
}

func TestHadamardProduct(t *testing.T) {
    var tests = []struct {
        left  *Color
        right *Color
        want  *Color
    }{
        {
            left:  &Color{1.0, 0.9, 0.5},
            right: &Color{0.0, 0.5, 0.5},
            want:  &Color{0.0, 0.45, 0.25},
        },
    }

    for _, test := range tests {
        if got := HadamardProduct(test.left, test.right); !got.Equal(test.want) {
            t.Errorf("input: %s - %s \ngot: %s. \nexpected: %s", test.left, test.right, got, test.want)
        }
    }
}
