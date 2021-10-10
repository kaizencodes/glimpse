package shapes

import (
    "glimpse/tuple"
    "math"
    "testing"
)

func TestLocalNormalAt(t *testing.T) {
    var tests = []struct {
        sphere *Sphere
        point  tuple.Tuple
        want   tuple.Tuple
    }{
        {
            sphere: NewSphere(),
            point:  tuple.NewPoint(1, 0, 0),
            want:   tuple.NewVector(1, 0, 0),
        },
        {
            sphere: NewSphere(),
            point:  tuple.NewPoint(0, 1, 0),
            want:   tuple.NewVector(0, 1, 0),
        },
        {
            sphere: NewSphere(),
            point:  tuple.NewPoint(0, 0, 1),
            want:   tuple.NewVector(0, 0, 1),
        },
        {
            sphere: NewSphere(),
            point:  tuple.NewPoint(1, 0, 0),
            want:   tuple.NewVector(1, 0, 0),
        },
        {
            sphere: NewSphere(),
            point:  tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
            want:   tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
        },
        {
            sphere: NewSphere(),
            point:  tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
            want:   tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0).Normalize(),
        },
    }

    for _, test := range tests {
        if got := test.sphere.LocalNormalAt(test.point); !got.Equal(test.want) {
            t.Errorf("shpere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", test.sphere, test.point, got, test.want)
        }
    }
}
