package color

import (
	"fmt"
	"strconv"

	"github.com/kaizencodes/glimpse/internal/utils"
)

type Color struct {
	R, G, B float64
}

func New(r, g, b float64) Color {
	return Color{r, g, b}
}

func FromSlice(s []float64) Color {
	return Color{s[0], s[1], s[2]}
}

func (c Color) Scalar(s float64) Color {
	return Color{c.R * s, c.G * s, c.B * s}
}

func (c Color) Equal(other Color) bool {
	return utils.FloatEquals(c.R, other.R) && utils.FloatEquals(c.G, other.G) &&
		utils.FloatEquals(c.B, other.B)
}

func (c Color) String() string {
	r := strconv.FormatFloat(c.R, 'f', -1, 64)
	g := strconv.FormatFloat(c.G, 'f', -1, 64)
	b := strconv.FormatFloat(c.B, 'f', -1, 64)

	return fmt.Sprintf("(%s, %s, %s)", r, g, b)
}

func Add(a, b Color) Color {
	return Color{a.R + b.R, a.G + b.G, a.B + b.B}
}

func Subtract(a, b Color) Color {
	return Color{a.R - b.R, a.G - b.G, a.B - b.B}
}

func HadamardProduct(a, b Color) Color {
	return Color{a.R * b.R, a.G * b.G, a.B * b.B}
}

func Black() Color {
	return Color{}
}

func White() Color {
	return Color{1, 1, 1}
}

func Red() Color {
	return Color{1, 0, 0}
}

func Green() Color {
	return Color{0, 1, 0}
}

func Blue() Color {
	return Color{0, 0, 1}
}
