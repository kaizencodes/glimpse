package object

import (
	"math"
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

func (t *Touple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2))
}

func (t *Touple) Normalize() *Touple {
	mag := t.Magnitude()
	return &Touple{t.x / mag, t.y / mag, t.z / mag, t.w / mag}
}

func (t *Touple) String() string {
	x := strconv.FormatFloat(t.x, 'f', -1, 64)
	y := strconv.FormatFloat(t.y, 'f', -1, 64)
	z := strconv.FormatFloat(t.z, 'f', -1, 64)
	w := strconv.FormatFloat(t.w, 'f', -1, 64)

	return "Touple(x: " + x + ", y: " + y + ", z: " + z + ", w: " + w + ")"
}

func Add(a, b *Touple) *Touple {
	return &Touple{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}
}

func Subtract(a, b *Touple) *Touple {
	return &Touple{a.x - b.x, a.y - b.y, a.z - b.z, a.w - b.w}
}

func Negate(t *Touple) *Touple {
	return &Touple{-t.x, -t.y, -t.z, -t.w}
}

func Multiply(t *Touple, s float64) *Touple {
	return &Touple{t.x * s, t.y * s, t.z * s, t.w * s}
}

func Dot(t *Touple, other *Touple) float64 {
	return t.x*other.x + t.y*other.y + t.z*other.z + t.w*other.w
}

func NewVector(x, y, z float64) *Touple {
	return &Touple{x, y, z, 0.0}
}

func NewPoint(x, y, z float64) *Touple {
	return &Touple{x, y, z, 1.0}
}
