package object

import (
	"fmt"
	"ray_tracer/calc"
	"strconv"
)

type Touple struct {
	x, y, z, w float64
}

func (t *Touple) IsPoint() bool {
	return t.w == 1.0
}

func (t *Touple) IsVector() bool {
	return t.w == 0.0
}

func (t *Touple) Equal(other *Touple) bool {
	return calc.FloatEquals(t.x, other.x) && calc.FloatEquals(t.y, other.y) &&
		calc.FloatEquals(t.z, other.z) && calc.FloatEquals(t.w, other.w)
}

func (t *Touple) Add(other *Touple) (*Touple, error) {
	if t.IsPoint() && other.IsPoint() {
		return nil, fmt.Errorf("addition of 2 points is not supported.")
	}
	return &Touple{t.x + other.x, t.y + other.y, t.z + other.z, t.w + other.w}, nil
}

func (t *Touple) Subtract(other *Touple) (*Touple, error) {
	if t.IsVector() && other.IsPoint() {
		return nil, fmt.Errorf("can't subtract a point from a vector.")
	}
	return &Touple{t.x - other.x, t.y - other.y, t.z - other.z, t.w - other.w}, nil
}

func (t *Touple) String() string {
	x := strconv.FormatFloat(t.x, 'f', -1, 64)
	y := strconv.FormatFloat(t.y, 'f', -1, 64)
	z := strconv.FormatFloat(t.z, 'f', -1, 64)
	w := strconv.FormatFloat(t.w, 'f', -1, 64)

	return "Touple(x: " + x + ", y: " + y + ", z: " + z + ", w: " + w + ")"
}

func NewVector(x, y, z float64) *Touple {
	return &Touple{x, y, z, 0.0}
}

func NewPoint(x, y, z float64) *Touple {
	return &Touple{x, y, z, 1.0}
}
