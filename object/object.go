package object

import (
	"fmt"
	"math"
	"ray_tracer/calc"
	"strconv"
)

type Tuple struct {
	x, y, z, w float64
}

func (t *Tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t *Tuple) IsVector() bool {
	return t.w == 0.0
}

func (t *Tuple) GetX() float64 {
	return t.x
}

func (t *Tuple) GetY() float64 {
	return t.y
}

func (t *Tuple) GetZ() float64 {
	return t.z
}

func (t *Tuple) Equal(other *Tuple) bool {
	return calc.FloatEquals(t.x, other.x) && calc.FloatEquals(t.y, other.y) &&
		calc.FloatEquals(t.z, other.z) && calc.FloatEquals(t.w, other.w)
}

func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2))
}

func (t *Tuple) Normalize() *Tuple {
	mag := t.Magnitude()
	return &Tuple{t.x / mag, t.y / mag, t.z / mag, t.w / mag}
}

func (t *Tuple) String() string {
	x := strconv.FormatFloat(t.x, 'f', -1, 64)
	y := strconv.FormatFloat(t.y, 'f', -1, 64)
	z := strconv.FormatFloat(t.z, 'f', -1, 64)
	w := strconv.FormatFloat(t.w, 'f', -1, 64)

	return fmt.Sprintf("Tuple(x: %s, y: %s, z: %s, w: %s)", x, y, z, w)
}

func Add(a, b *Tuple) *Tuple {
	return &Tuple{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}
}

func Subtract(a, b *Tuple) *Tuple {
	return &Tuple{a.x - b.x, a.y - b.y, a.z - b.z, a.w - b.w}
}

func Negate(t *Tuple) *Tuple {
	return &Tuple{-t.x, -t.y, -t.z, -t.w}
}

func Multiply(t *Tuple, s float64) *Tuple {
	return &Tuple{t.x * s, t.y * s, t.z * s, t.w * s}
}

func Dot(a *Tuple, b *Tuple) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z + a.w*b.w
}

func Cross(a *Tuple, b *Tuple) *Tuple {
	return &Tuple{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
		0.0,
	}
}

func NewVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}
