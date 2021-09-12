package objects

import (
    "glimpse/matrix"
    "glimpse/tuple"
    "math"
    "testing"
)

func TestNormal(t *testing.T) {
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
        if got := test.sphere.Normal(test.point); !got.Equal(test.want) {
            t.Errorf("shpere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", test.sphere, test.point, got, test.want)
        }
    }

    sphere := NewSphere()
    sphere.SetTransform(matrix.Translation(0, 1, 0))
    point := tuple.NewPoint(0, 1.70711, -0.70711)
    want := tuple.NewVector(0, 0.7071067811865475, -0.7071067811865476)

    if got := sphere.Normal(point); !got.Equal(want) {
        t.Errorf("shpere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", sphere, point, got, want)
    }

    transform, _ := matrix.Multiply(matrix.Scaling(1, 0.5, 1), matrix.RotationZ(math.Pi/5.0))
    sphere.SetTransform(transform)
    point = tuple.NewPoint(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0)
    want = tuple.NewVector(0, 0.9701425001453319, -0.24253562503633294)

    if got := sphere.Normal(point); !got.Equal(want) {
        t.Errorf("shpere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", sphere, point, got, want)
    }
}
