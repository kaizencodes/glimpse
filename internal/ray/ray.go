// contains ray struct and related methods.
package ray

import (
	"fmt"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

// BounceLimit is the maximum number of bounces a ray can make. This is used to prevent infinite recursion.
// Also it is used to save time when rendering.
const BounceLimit = 5

type Ray struct {
	Origin      tuple.Tuple // point
	Direction   tuple.Tuple // vector
	BounceLimit int
}

func New(origin, direction tuple.Tuple) *Ray {
	return &Ray{Origin: origin, Direction: direction, BounceLimit: BounceLimit}
}

// Position returns the point at a given distance along the ray.
func (r *Ray) Position(dist float64) tuple.Tuple {
	return tuple.Add(r.Origin, r.Direction.Scalar(dist))
}

// String returns a string representation of the ray.
func (r *Ray) String() string {
	return fmt.Sprintf("Ray(Origin: %s, Direction: %s)", r.Origin, r.Direction)
}

// equality check between two rays.
func (r *Ray) Equal(other *Ray) bool {
	return r.Origin.Equal(other.Origin) && r.Direction.Equal(other.Direction)
}

// Translate applies a translation matrix to the ray. Moving it in the 3d space.
func (r *Ray) Translate(x, y, z float64) *Ray {
	origin := tuple.Multiply(matrix.Translation(x, y, z), r.Origin)
	return &Ray{Origin: origin, Direction: r.Direction}
}

// Scale applies a scaling matrix to the ray.
func (r *Ray) Scale(x, y, z float64) *Ray {
	origin := tuple.Multiply(matrix.Scaling(x, y, z), r.Origin)
	direction := tuple.Multiply(matrix.Scaling(x, y, z), r.Direction)
	return &Ray{Origin: origin, Direction: direction}
}
