package ray

import (
	"fmt"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

const BounceLimit = 5

type Ray struct {
	Origin      tuple.Tuple
	Direction   tuple.Tuple
	BounceLimit int
}

func (r *Ray) Position(dist float64) tuple.Tuple {
	return tuple.Add(r.Origin, r.Direction.Scalar(dist))
}

func (r *Ray) String() string {
	return fmt.Sprintf("Ray(Origin: %s, Direction: %s)", r.Origin, r.Direction)
}

func (r *Ray) Equal(other *Ray) bool {
	return r.Origin.Equal(other.Origin) && r.Direction.Equal(other.Direction)
}

func (r *Ray) Translate(x, y, z float64) *Ray {
	origin := tuple.Multiply(matrix.Translation(x, y, z), r.Origin)
	return &Ray{Origin: origin, Direction: r.Direction}
}

func (r *Ray) Scale(x, y, z float64) *Ray {
	origin := tuple.Multiply(matrix.Scaling(x, y, z), r.Origin)
	direction := tuple.Multiply(matrix.Scaling(x, y, z), r.Direction)
	return &Ray{Origin: origin, Direction: direction}
}

func NewRay(origin, direction tuple.Tuple) *Ray {
	return &Ray{Origin: origin, Direction: direction, BounceLimit: BounceLimit}
}
