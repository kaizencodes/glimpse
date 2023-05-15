package ray

import (
	"fmt"

	"github.com/kaizencodes/glimpse/matrix"
	"github.com/kaizencodes/glimpse/tuple"
)

const BounceLimit = 5

type Ray struct {
	origin      tuple.Tuple
	direction   tuple.Tuple
	bounceLimit int
}

func (r *Ray) Position(dist float64) tuple.Tuple {
	return tuple.Add(r.origin, r.direction.Scalar(dist))
}

func (r *Ray) String() string {
	return fmt.Sprintf("Ray(origin: %s, direction: %s)", r.origin, r.direction)
}

func (r *Ray) Equal(other *Ray) bool {
	return r.origin.Equal(other.origin) && r.direction.Equal(other.direction)
}

func (r *Ray) Translate(x, y, z float64) *Ray {
	origin, err := tuple.Multiply(matrix.Translation(x, y, z), r.origin)
	if err != nil {
		panic(err)
	}
	return &Ray{origin: origin, direction: r.direction}
}

func (r *Ray) Scale(x, y, z float64) *Ray {
	origin, err := tuple.Multiply(matrix.Scaling(x, y, z), r.origin)
	if err != nil {
		panic(err)
	}
	direction, err := tuple.Multiply(matrix.Scaling(x, y, z), r.direction)
	if err != nil {
		panic(err)
	}
	return &Ray{origin: origin, direction: direction}
}

func (r *Ray) Origin() tuple.Tuple {
	return r.origin
}

func (r *Ray) Direction() tuple.Tuple {
	return r.direction
}

func (r *Ray) BounceLimit() int {
	return r.bounceLimit
}

func (r *Ray) SetBounceLimit(bounceLimit int) {
	r.bounceLimit = bounceLimit
}

func NewRay(origin, direction tuple.Tuple) *Ray {
	return &Ray{origin, direction, BounceLimit}
}
