package patterns

import (
	"fmt"
	"glimpse/color"
	"glimpse/tuple"
	"math"
)

type Pattern interface {
	ColorAt(point tuple.Tuple) color.Color
	String() string
}

type StripePattern struct {
	a, b color.Color
}

func NewStripePattern(a, b color.Color) StripePattern {
	return StripePattern{a, b}
}

func (p StripePattern) ColorAt(point tuple.Tuple) color.Color {
	if math.Mod(math.Floor(point.X()), 2) == 0 {
		return p.a
	} else {
		return p.b
	}
}

func (p StripePattern) String() string {
	return fmt.Sprintf("StripePattern(a: %s, b: %s)", p.a, p.b)
}

type MonoPattern struct {
	color color.Color
}

func NewMonoPattern(c color.Color) MonoPattern {
	return MonoPattern{color: c}
}

func (p MonoPattern) ColorAt(point tuple.Tuple) color.Color {
	return p.color
}

func (p MonoPattern) String() string {
	return fmt.Sprintf("MonoPattern(color: %s)", p.color)
}
