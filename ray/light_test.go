package ray

import (
    "glimpse/color"
    "glimpse/shapes"
    "glimpse/tuple"
    "math"
    "testing"
)

func TestLighting(t *testing.T) {
    var tests = []struct {
        eyeV     tuple.Tuple
        normalV  tuple.Tuple
        light    Light
        inShadow bool
        expected color.Color
    }{
        {
            eyeV:     tuple.NewVector(0, 0, -1),
            normalV:  tuple.NewVector(0, 0, -1),
            light:    NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1)),
            inShadow: false,
            expected: color.New(1.9, 1.9, 1.9),
        },
        {
            eyeV:     tuple.NewVector(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0),
            normalV:  tuple.NewVector(0, 0, -1),
            light:    NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1)),
            inShadow: false,
            expected: color.New(1.0, 1.0, 1.0),
        },
        {
            eyeV:     tuple.NewVector(0, 0, -1),
            normalV:  tuple.NewVector(0, 0, -1),
            light:    NewLight(tuple.NewPoint(0, 10, -10), color.New(1, 1, 1)),
            inShadow: false,
            expected: color.New(0.7363961030678927, 0.7363961030678927, 0.7363961030678927),
        },
        {
            eyeV:     tuple.NewVector(0, -math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0),
            normalV:  tuple.NewVector(0, 0, -1),
            light:    NewLight(tuple.NewPoint(0, 10, -10), color.New(1, 1, 1)),
            inShadow: false,
            expected: color.New(1.6363961030678928, 1.6363961030678928, 1.6363961030678928),
        },
        {
            eyeV:     tuple.NewVector(0, 0, -1),
            normalV:  tuple.NewVector(0, 0, -1),
            light:    NewLight(tuple.NewPoint(0, 0, 10), color.New(1, 1, 1)),
            inShadow: false,
            expected: color.New(0.1, 0.1, 0.1),
        },
        {
            eyeV:     tuple.NewVector(0, 0, -1),
            normalV:  tuple.NewVector(0, 0, -1),
            light:    NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1)),
            inShadow: true,
            expected: color.New(0.1, 0.1, 0.1),
        },
    }

    mat := shapes.DefaultMaterial()
    pos := tuple.NewPoint(0, 0, 0)
    for _, test := range tests {
        if got := Lighting(mat, test.light, pos, test.eyeV, test.normalV, test.inShadow); !got.Equal(test.expected) {
            t.Errorf("Lighting:\n light: %s \neyeV: %s \nnormalV: %s\ninShadow: %t\ngot: \n%s. \nexpected: \n%s", test.light, test.eyeV, test.normalV, test.inShadow, got, test.expected)
        }
    }
}
