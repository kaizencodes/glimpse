package color

import (
	"fmt"
	"glimpse/calc"
	"math"
	"strconv"
)

type Color struct {
	R, G, B float64
}

func Add(a, b *Color) *Color {
	return &Color{a.R + b.R, a.G + b.G, a.B + b.B}
}

func Subtract(a, b *Color) *Color {
	return &Color{a.R - b.R, a.G - b.G, a.B - b.B}
}

func Multiply(c *Color, s float64) *Color {
	return &Color{c.R * s, c.G * s, c.B * s}
}

func HadamardProduct(a, b *Color) *Color {
	return &Color{a.R * b.R, a.G * b.G, a.B * b.B}
}

func (c *Color) Equal(other *Color) bool {
	return calc.FloatEquals(c.R, other.R) && calc.FloatEquals(c.G, other.G) &&
		calc.FloatEquals(c.B, other.B)
}

func (c *Color) String() string {
	r := strconv.FormatFloat(c.R, 'f', -1, 64)
	g := strconv.FormatFloat(c.G, 'f', -1, 64)
	b := strconv.FormatFloat(c.B, 'f', -1, 64)

	return fmt.Sprintf("(%s, %s, %s)", r, g, b)
}

func (c *Color) ConvertToPpm() string {
	PpmMax := 255.0
	r := int(math.Min(math.Max(0, c.R*PpmMax), PpmMax))
	g := int(math.Min(math.Max(0, c.G*PpmMax), PpmMax))
	b := int(math.Min(math.Max(0, c.B*PpmMax), PpmMax))

	return fmt.Sprintf("%d %d %d ", r, g, b)
}
