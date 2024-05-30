package shapes

import (
	"math"
	"sort"
	"strconv"
)

type Intersection struct {
	t, u, v float64
	shape   Shape
}

type Intersections []Intersection

func (i Intersection) Empty() bool {
	return i.t == math.MaxFloat64
}

func (i Intersection) T() float64 {
	return i.t
}

func (i Intersection) Shape() Shape {
	return i.shape
}

func (xs Intersections) String() string {
	var result string

	for i := 0; i < len(xs); i++ {
		result += strconv.FormatFloat(xs[i].t, 'f', -1, 64) + ", "
	}
	return result
}

func (xs Intersections) Sort() {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].t < xs[j].t
	})
}

func (xs Intersections) Hit() Intersection {
	res := Intersection{t: math.MaxFloat64}
	for i := 0; i < len(xs); i++ {
		if xs[i].t < 0 {
			continue
		}
		if xs[i].t < res.t {
			res = xs[i]
		}
	}
	return res
}

func NewIntersectionWithUV(t, u, v float64, obj Shape) Intersection {
	return Intersection{t, u, v, obj}
}

func NewIntersection(t float64, obj Shape) Intersection {
	return Intersection{t, -1, -1, obj}
}
