package ray

import (
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
    shpere := NewShpere()
    var tests = []struct {
        ray  Ray
        s    *Sphere
        want Intersections
    }{
        {
            ray: New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1)),
            s:   &shpere,
            want: Intersections{
                Intersection{t: 4.0, object: &shpere},
                Intersection{t: 6.0, object: &shpere},
            },
        },
        {
            ray: New(tuple.NewPoint(0, 1, -5), tuple.NewVector(0, 0, 1)),
            s:   &shpere,
            want: Intersections{
                Intersection{t: 5.0, object: &shpere},
                Intersection{t: 5.0, object: &shpere},
            },
        },
        {
            ray:  New(tuple.NewPoint(0, 2, -5), tuple.NewVector(0, 0, 1)),
            s:    &shpere,
            want: Intersections{},
        },
        {
            ray: New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1)),
            s:   &shpere,
            want: Intersections{
                Intersection{t: -1.0, object: &shpere},
                Intersection{t: 1.0, object: &shpere},
            },
        },
        {
            ray: New(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1)),
            s:   &shpere,
            want: Intersections{
                Intersection{t: -6.0, object: &shpere},
                Intersection{t: -4.0, object: &shpere},
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
    shpere := NewShpere()
    var tests = []struct {
        collection Intersections
        want       Intersection
    }{
        {
            collection: Intersections{
                Intersection{t: 1.0, object: &shpere},
                Intersection{t: 2.0, object: &shpere},
            },
            want: Intersection{t: 1.0, object: &shpere},
        },
        {
            collection: Intersections{
                Intersection{t: -1.0, object: &shpere},
                Intersection{t: 1.0, object: &shpere},
            },
            want: Intersection{t: 1.0, object: &shpere},
        },
        {
            collection: Intersections{
                Intersection{t: -2.0, object: &shpere},
                Intersection{t: -1.0, object: &shpere},
            },
            want: Intersection{t: math.MaxFloat64},
        },
        {
            collection: Intersections{
                Intersection{t: 5.0, object: &shpere},
                Intersection{t: 7.0, object: &shpere},
                Intersection{t: -3.0, object: &shpere},
                Intersection{t: 2.0, object: &shpere},
            },
            want: Intersection{t: 2.0, object: &shpere},
        },
    }

    for _, test := range tests {
        if got := Hit(test.collection); got.t != test.want.t {
            t.Errorf("Hit: collection\n%s \ngot: \n%f. \nexpected: \n%f", test.collection, got.t, test.want.t)
        }
    }
}
