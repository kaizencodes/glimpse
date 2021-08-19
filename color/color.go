package color

import (
	"fmt"
	"ray_tracer/calc"
	"strconv"
)

type Color struct {
	red, green, blue float64
}

func Add(a, b *Color) *Color {
	return &Color{a.red + b.red, a.green + b.green, a.blue + b.blue}
}

func Subtract(a, b *Color) *Color {
	return &Color{a.red - b.red, a.green - b.green, a.blue - b.blue}
}

func Multiply(c *Color, s float64) *Color {
	return &Color{c.red * s, c.green * s, c.blue * s}
}

func HadamardProduct(a, b *Color) *Color {
	return &Color{a.red * b.red, a.green * b.green, a.blue * b.blue}
}

func (c *Color) Equal(other *Color) bool {
	return calc.FloatEquals(c.red, other.red) && calc.FloatEquals(c.green, other.green) &&
		calc.FloatEquals(c.blue, other.blue)
}

func (c *Color) String() string {
	r := strconv.FormatFloat(c.red, 'f', -1, 64)
	g := strconv.FormatFloat(c.green, 'f', -1, 64)
	b := strconv.FormatFloat(c.blue, 'f', -1, 64)

	return fmt.Sprintf("Color(x: %s, y: %s, z: %s)", r, g, b)
}
