package shapes

import (
	"math"
	"testing"
)

func TestHit(t *testing.T) {
	shape := Shape(NewSphere())
	var tests = []struct {
		collection Intersections
		want       Intersection
	}{
		{
			collection: Intersections{
				NewIntersection(1.0, shape),
				NewIntersection(2.0, shape),
			},
			want: Intersection{t: 1.0, shape: shape},
		},
		{
			collection: Intersections{
				NewIntersection(-1.0, shape),
				NewIntersection(1.0, shape),
			},
			want: NewIntersection(1.0, shape),
		},
		{
			collection: Intersections{
				NewIntersection(-2.0, shape),
				NewIntersection(-1.0, shape),
			},
			want: Intersection{t: math.MaxFloat64},
		},
		{
			collection: Intersections{
				NewIntersection(5.0, shape),
				NewIntersection(7.0, shape),
				NewIntersection(-3.0, shape),
				NewIntersection(2.0, shape),
			},
			want: NewIntersection(2.0, shape),
		},
	}

	for _, test := range tests {
		if got := test.collection.Hit(); got.t != test.want.t {
			t.Errorf("Hit: collection\n%s \ngot: \n%f. \nexpected: \n%f", test.collection, got.t, test.want.t)
		}
	}
}
