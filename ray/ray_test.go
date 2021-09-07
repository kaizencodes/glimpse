package ray

import (
    "glimpse/tuple"
    "testing"
)

func TestPosition(t *testing.T) {
    var tests = []struct {
        ray  Ray
        dist float64
        want tuple.Tuple
    }{
        {
            ray:  New(tuple.Tuple{2, 3, 4, 1}, tuple.Tuple{1, 0, 0, 0}),
            dist: 0,
            want: tuple.Tuple{2, 3, 4, 1},
        },
        {
            ray:  New(tuple.Tuple{2, 3, 4, 1}, tuple.Tuple{1, 0, 0, 0}),
            dist: 1,
            want: tuple.Tuple{3, 3, 4, 1},
        },
        {
            ray:  New(tuple.Tuple{2, 3, 4, 1}, tuple.Tuple{0, 1, 0, 0}),
            dist: -1,
            want: tuple.Tuple{2, 2, 4, 1},
        },
        {
            ray:  New(tuple.Tuple{2, 3, 4, 1}, tuple.Tuple{0, 0, 1, 0}),
            dist: 2.5,
            want: tuple.Tuple{2, 3, 6.5, 1},
        },
    }

    for _, test := range tests {
        if got := test.ray.Position(test.dist); !got.Equal(test.want) {
            t.Errorf("ray position:\n%s \n dist: %f. \ngot: \n%s. \nexpected: \n%s", test.ray, test.dist, got, test.want)
        }
    }
}
