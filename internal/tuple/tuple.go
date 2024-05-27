// Ordered list of numbers. It has 2 use cases in the codebase:
// 1. Representing a point in 3D space
// 2. Representing a vector in 3D space
package tuple

import (
	"fmt"
	"math"
	"strconv"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/utils"
)

// Glimpse uses a left-handed coordinates.
// with the y axis pointing up, and the x axis pointing to the right,
// the z axis points away from the camera.
type Tuple struct {
	X, Y, Z, W float64
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t Tuple) ToVector() Tuple {
	return Tuple{t.X, t.Y, t.Z, 0}
}

func (t Tuple) Equal(other Tuple) bool {
	return utils.FloatEquals(t.X, other.X) && utils.FloatEquals(t.Y, other.Y) &&
		utils.FloatEquals(t.Z, other.Z) && utils.FloatEquals(t.W, other.W)
}

// Used to calculate the magnitude of a vector with the Pythagoras’ theorem.
func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
}

// Normalization is the process of taking an arbitrary vector and converting it into a unit vector.
// A normalized vector has magnitude of 1.
func (t Tuple) Normalize() Tuple {
	mag := t.Magnitude()
	return Tuple{t.X / mag, t.Y / mag, t.Z / mag, t.W / mag}
}

func (t Tuple) String() string {
	x := strconv.FormatFloat(t.X, 'f', -1, 64)
	y := strconv.FormatFloat(t.Y, 'f', -1, 64)
	z := strconv.FormatFloat(t.Z, 'f', -1, 64)
	w := strconv.FormatFloat(t.W, 'f', -1, 64)

	return fmt.Sprintf("Tuple(x: %s, y: %s, z: %s, w: %s)", x, y, z, w)
}

// Used to multiply vectors by a scalar.
func (t Tuple) Scalar(s float64) Tuple {
	return Tuple{t.X * s, t.Y * s, t.Z * s, t.W * s}
}

// Used to get the opposite of vectors.
func (t Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t Tuple) ToSlice() []float64 {
	return []float64{t.X, t.Y, t.Z, t.W}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func NewVectorFromSlice(s []float64) Tuple {
	return Tuple{s[0], s[1], s[2], 0}
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func NewPointFromSlice(s []float64) Tuple {
	return Tuple{s[0], s[1], s[2], 1}
}

// Adding together a point and a vector results in a point.
// Adding together two vectors results in a vector.
// TODO: Adding together two points results in an error.
func Add(a, b Tuple) Tuple {
	return Tuple{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

// Subtracting 2 vectors results in a vector, the change in direction between the two.
// Subtracting 2 points result in a vector pointing from b to a.
// Subtracting a vector from a point results in a point, the point is "moved back" by the vector.
// TODO: subtracting a a point from a vector results in an error.
func Subtract(a, b Tuple) Tuple {
	return Tuple{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

func Multiply(a matrix.Matrix, b Tuple) Tuple {
	mat := matrix.New(4, 1, b.ToSlice())
	mat = matrix.Multiply(a, mat)
	// implicitly transpose the result.
	return Tuple{mat.At(0, 0), mat.At(1, 0), mat.At(2, 0), mat.At(3, 0)}
}

// The dot product of two vectors.
func Dot(a Tuple, b Tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

// The cross product of two vectors.
func Cross(a Tuple, b Tuple) Tuple {
	return Tuple{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
		0.0,
	}
}

func Reflect(incoming, normal Tuple) Tuple {
	return Subtract(incoming, normal.Scalar(2.0*Dot(incoming, normal)))
}
