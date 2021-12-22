package patterns

import (
	"glimpse/color"
	"glimpse/tuple"
	"math"
)

type Pattern struct {
	colors []color.Color
}

func StripePattern(a, b color.Color) Pattern {
	return Pattern{colors: []color.Color{a, b}}
}

func StripeAt(p Pattern, point tuple.Tuple) color.Color {
	if math.Mod(math.Floor(point.X()), 2) == 0 {
		return p.colors[1]
	} else {
		return p.colors[0]
	}
}
