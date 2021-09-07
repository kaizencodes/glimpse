package ray

import (
	"fmt"
	"glimpse/tuple"
)

type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

func New(origin, direction tuple.Tuple) Ray {
	return Ray{origin, direction}
}

func (r Ray) Position(dist float64) tuple.Tuple {
	return tuple.Add(r.origin, r.direction.Scalar(dist))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(origin: %s, direction: %s)", r.origin, r.direction)
}
