package ray

import (
    "glimpse/color"
    "glimpse/matrix"
    "glimpse/objects"
    "glimpse/tuple"
    "math"
    "testing"
)

func TestPosition(t *testing.T) {
    var tests = []struct {
        ray  Ray
        dist float64
        want tuple.Tuple
    }{
        {
            ray:  New(tuple.NewPoint(2, 3, 4), tuple.NewVector(1, 0, 0)),
            dist: 0,
            want: tuple.NewPoint(2, 3, 4),
        },
        {
            ray:  New(tuple.NewPoint(2, 3, 4), tuple.NewVector(1, 0, 0)),
            dist: 1,
            want: tuple.NewPoint(3, 3, 4),
        },
        {
            ray:  New(tuple.NewPoint(2, 3, 4), tuple.NewVector(0, 1, 0)),
            dist: -1,
            want: tuple.NewPoint(2, 2, 4),
        },
        {
            ray:  New(tuple.NewPoint(2, 3, 4), tuple.NewVector(0, 0, 1)),
            dist: 2.5,
            want: tuple.NewPoint(2, 3, 6.5),
        },
    }

    for _, test := range tests {
        if got := test.ray.Position(test.dist); !got.Equal(test.want) {
            t.Errorf("ray position:\n%s \n dist: %f. \ngot: \n%s. \nexpected: \n%s", test.ray, test.dist, got, test.want)
        }
    }
}

func TestIntersect(t *testing.T) {
    shpere := objects.NewSphere()
    var tests = []struct {
        ray  Ray
        s    *objects.Sphere
        want Intersections
    }{
        {
            ray: New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1)),
            s:   shpere,
            want: Intersections{
                Intersection{t: 4.0, object: shpere},
                Intersection{t: 6.0, object: shpere},
            },
        },
        {
            ray: New(tuple.NewPoint(0, 1, -5), tuple.NewVector(0, 0, 1)),
            s:   shpere,
            want: Intersections{
                Intersection{t: 5.0, object: shpere},
                Intersection{t: 5.0, object: shpere},
            },
        },
        {
            ray:  New(tuple.NewPoint(0, 2, -5), tuple.NewVector(0, 0, 1)),
            s:    shpere,
            want: Intersections{},
        },
        {
            ray: New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1)),
            s:   shpere,
            want: Intersections{
                Intersection{t: -1.0, object: shpere},
                Intersection{t: 1.0, object: shpere},
            },
        },
        {
            ray: New(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1)),
            s:   shpere,
            want: Intersections{
                Intersection{t: -6.0, object: shpere},
                Intersection{t: -4.0, object: shpere},
            },
        },
    }

    for _, test := range tests {
        got := Intersect(test.ray, test.s)
        for i, _ := range got {
            if got[i].t != test.want[i].t {
                t.Errorf("incorrect t of intersect:\n%s \n \ngot: \n%f. \nexpected: \n%f", test.ray, got[i].t, test.want[i].t)
            }
            if got[i].object != test.want[i].object {
                t.Errorf("incorrect object of intersect:\n%s \n \ngot: \n%s. \nexpected: \n%s", test.ray, got[i].object, test.want[i].object)
            }
        }
    }
}

func TestHit(t *testing.T) {
    shpere := objects.NewSphere()
    var tests = []struct {
        collection Intersections
        want       Intersection
    }{
        {
            collection: Intersections{
                Intersection{t: 1.0, object: shpere},
                Intersection{t: 2.0, object: shpere},
            },
            want: Intersection{t: 1.0, object: shpere},
        },
        {
            collection: Intersections{
                Intersection{t: -1.0, object: shpere},
                Intersection{t: 1.0, object: shpere},
            },
            want: Intersection{t: 1.0, object: shpere},
        },
        {
            collection: Intersections{
                Intersection{t: -2.0, object: shpere},
                Intersection{t: -1.0, object: shpere},
            },
            want: Intersection{t: math.MaxFloat64},
        },
        {
            collection: Intersections{
                Intersection{t: 5.0, object: shpere},
                Intersection{t: 7.0, object: shpere},
                Intersection{t: -3.0, object: shpere},
                Intersection{t: 2.0, object: shpere},
            },
            want: Intersection{t: 2.0, object: shpere},
        },
    }

    for _, test := range tests {
        if got := Hit(test.collection); got.t != test.want.t {
            t.Errorf("Hit: collection\n%s \ngot: \n%f. \nexpected: \n%f", test.collection, got.t, test.want.t)
        }
    }
}

func TestTranslate(t *testing.T) {
    ray := New(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
    want := New(tuple.NewPoint(4, 6, 8), tuple.NewVector(0, 1, 0))
    x, y, z := 3.0, 4.0, 5.0

    if got := ray.Translate(x, y, z); !got.Equal(want) {
        t.Errorf("translation(%f, %f, %f),\nray:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, ray, got, want)
    }

    x, y, z = 2.0, 3.0, 4.0
    want = New(tuple.NewPoint(2, 6, 12), tuple.NewVector(0, 3, 0))

    if got := ray.Scale(x, y, z); !got.Equal(want) {
        t.Errorf("scale(%f, %f, %f),\nray:\n%s\n\ngot:\n%s", x, y, z, ray, got)
    }
}

func TestScale(t *testing.T) {
    ray := New(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
    want := New(tuple.NewPoint(2, 6, 12), tuple.NewVector(0, 3, 0))
    x, y, z := 2.0, 3.0, 4.0

    if got := ray.Scale(x, y, z); !got.Equal(want) {
        t.Errorf("scale(%f, %f, %f),\nray:\n%s\n\ngot:\n%s", x, y, z, ray, got)
    }
}

func TestShpereTransformations(t *testing.T) {
    ray := New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
    sphere := objects.NewSphere()
    sphere.SetTransform(matrix.GetScaling(2, 2, 2))
    want := Intersections{
        Intersection{t: 3.0, object: sphere},
        Intersection{t: 7.0, object: sphere},
    }

    got := Intersect(ray, sphere)
    for i, _ := range got {
        if got[i].t != want[i].t {
            t.Errorf("incorrect t of intersect:\n%s \n \ngot: \n%f. \nexpected: \n%f", ray, got[i].t, want[i].t)
        }
        if got[i].object != want[i].object {
            t.Errorf("incorrect object of intersect:\n%s \n \ngot: \n%s. \nexpected: \n%s", ray, got[i].object, want[i].object)
        }
    }
}

func TestLighting(t *testing.T) {
    var tests = []struct {
        eyeV    tuple.Tuple
        normalV tuple.Tuple
        light   PointLight
        want    color.Color
    }{
        {
            eyeV:    tuple.NewVector(0, 0, -1),
            normalV: tuple.NewVector(0, 0, -1),
            light:   NewPointLight(tuple.NewPoint(0, 0, -10), color.Color{1, 1, 1}),
            want:    color.Color{1.9, 1.9, 1.9},
        },
        {
            eyeV:    tuple.NewVector(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0),
            normalV: tuple.NewVector(0, 0, -1),
            light:   NewPointLight(tuple.NewPoint(0, 0, -10), color.Color{1, 1, 1}),
            want:    color.Color{1.0, 1.0, 1.0},
        },
        {
            eyeV:    tuple.NewVector(0, 0, -1),
            normalV: tuple.NewVector(0, 0, -1),
            light:   NewPointLight(tuple.NewPoint(0, 10, -10), color.Color{1, 1, 1}),
            want:    color.Color{0.7363961030678927, 0.7363961030678927, 0.7363961030678927},
        },
        {
            eyeV:    tuple.NewVector(0, -math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0),
            normalV: tuple.NewVector(0, 0, -1),
            light:   NewPointLight(tuple.NewPoint(0, 10, -10), color.Color{1, 1, 1}),
            want:    color.Color{1.6363961030678928, 1.6363961030678928, 1.6363961030678928},
        },
        {
            eyeV:    tuple.NewVector(0, 0, -1),
            normalV: tuple.NewVector(0, 0, -1),
            light:   NewPointLight(tuple.NewPoint(0, 0, 10), color.Color{1, 1, 1}),
            want:    color.Color{0.1, 0.1, 0.1},
        },
    }

    mat := objects.DefaultMaterial()
    pos := tuple.NewPoint(0, 0, 0)
    for _, test := range tests {
        if got := Lighting(mat, test.light, pos, test.eyeV, test.normalV); !got.Equal(test.want) {
            t.Errorf("Lighting:\n light: %s \neyeV: %s \nnormalV: %s\ngot: \n%s. \nexpected: \n%s", test.light, test.eyeV, test.normalV, got, test.want)
        }
    }
}
