package color

import (
	"fmt"
	"glimpse/calc"
	"strconv"
)

type Color struct {
	r, g, b float64
}

func (c Color) Scalar(s float64) Color {
	return Color{c.r * s, c.g * s, c.b * s}
}

func (c Color) Equal(other Color) bool {
	return calc.FloatEquals(c.r, other.r) && calc.FloatEquals(c.g, other.g) &&
		calc.FloatEquals(c.b, other.b)
}

func (c Color) String() string {
	r := strconv.FormatFloat(c.r, 'f', -1, 64)
	g := strconv.FormatFloat(c.g, 'f', -1, 64)
	b := strconv.FormatFloat(c.b, 'f', -1, 64)

	return fmt.Sprintf("(%s, %s, %s)", r, g, b)
}

func (c Color) R() float64 {
	return c.r
}

func (c Color) G() float64 {
	return c.g
}

func (c Color) B() float64 {
	return c.b
}

func New(r, g, b float64) Color {
	return Color{r, g, b}
}

func Add(a, b Color) Color {
	return Color{a.r + b.r, a.g + b.g, a.b + b.b}
}

func Subtract(a, b Color) Color {
	return Color{a.r - b.r, a.g - b.g, a.b - b.b}
}

func HadamardProduct(a, b Color) Color {
	return Color{a.r * b.r, a.g * b.g, a.b * b.b}
}

func Black() Color {
	return Color{}
}
