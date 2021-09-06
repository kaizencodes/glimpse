package tuple

import (
	"fmt"
	"glimpse/calc"
	"glimpse/matrix"
	"math"
	"strconv"
)

type Tuple struct {
	X, Y, Z, W float64
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t Tuple) GetX() float64 {
	return t.X
}

func (t Tuple) GetY() float64 {
	return t.Y
}

func (t Tuple) GetZ() float64 {
	return t.Z
}

func (t Tuple) Equal(other Tuple) bool {
	return calc.FloatEquals(t.X, other.X) && calc.FloatEquals(t.Y, other.Y) &&
		calc.FloatEquals(t.Z, other.Z) && calc.FloatEquals(t.W, other.W)
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
}

func (t Tuple) Normalize() Tuple {
	mag := t.Magnitude()
	return Tuple{t.X / mag, t.Y / mag, t.Z / mag, t.W / mag}
}

func (t Tuple) String() string {
	X := strconv.FormatFloat(t.X, 'f', -1, 64)
	Y := strconv.FormatFloat(t.Y, 'f', -1, 64)
	Z := strconv.FormatFloat(t.Z, 'f', -1, 64)
	W := strconv.FormatFloat(t.W, 'f', -1, 64)

	return fmt.Sprintf("Tuple(X: %s, Y: %s, Z: %s, W: %s)", X, Y, Z, W)
}

func Add(a, b Tuple) Tuple {
	return Tuple{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

func Subtract(a, b Tuple) Tuple {
	return Tuple{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

func Multiply(a matrix.Matrix, b Tuple) (Tuple, error) {
	mat := matrix.New(4, 1)
	for i, v := range b.ToSlice() {
		mat[i][0] = float64(v)
	}
	mat, err := matrix.Multiply(a, mat)
	if err != nil {
		return Tuple{}, err
	}
	mat = mat.Transpose()
	return Tuple{mat[0][0], mat[0][1], mat[0][2], mat[0][3]}, nil
}

func (t Tuple) Scalar(s float64) Tuple {
	return Tuple{t.X * s, t.Y * s, t.Z * s, t.W * s}
}

func Negate(t Tuple) Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func Dot(a Tuple, b Tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

func Cross(a Tuple, b Tuple) Tuple {
	return Tuple{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
		0.0,
	}
}

func (t Tuple) ToSlice() []float64 {
	return []float64{t.X, t.Y, t.Z, t.W}
}

func NewVector(X, Y, Z float64) Tuple {
	return Tuple{X, Y, Z, 0.0}
}

func NewPoint(X, Y, Z float64) Tuple {
	return Tuple{X, Y, Z, 1.0}
}
