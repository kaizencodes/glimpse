package shapes

import (
    "glimpse/tuple"
    "math"
    "testing"
)

func TestSphereLocalNormalAt(t *testing.T) {
    var tests = []struct {
        sphere   *Sphere
        point    tuple.Tuple
        expected tuple.Tuple
    }{
        {
            sphere:   NewSphere(),
            point:    tuple.NewPoint(1, 0, 0),
            expected: tuple.NewVector(1, 0, 0),
        },
        {
            sphere:   NewSphere(),
            point:    tuple.NewPoint(0, 1, 0),
            expected: tuple.NewVector(0, 1, 0),
        },
        {
            sphere:   NewSphere(),
            point:    tuple.NewPoint(0, 0, 1),
            expected: tuple.NewVector(0, 0, 1),
        },
        {
            sphere:   NewSphere(),
            point:    tuple.NewPoint(1, 0, 0),
            expected: tuple.NewVector(1, 0, 0),
        },
        {
            sphere:   NewSphere(),
            point:    tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
            expected: tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
        },
        {
            sphere:   NewSphere(),
            point:    tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
            expected: tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0).Normalize(),
        },
    }

    for _, test := range tests {
        if got := test.sphere.LocalNormalAt(test.point); !got.Equal(test.expected) {
            t.Errorf("shpere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", test.sphere, test.point, got, test.expected)
        }
    }
}
