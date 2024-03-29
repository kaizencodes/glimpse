package materials

import (
	"fmt"
	"math"

	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type PatternType int

const (
	Base PatternType = iota
	Stripe
	Gradient
	Ring
	Checker
	Test
)

type Pattern struct {
	transform matrix.Matrix
	colorAt   func(tuple.Tuple) color.Color // function that determines the color at a point
}

func NewPattern(pattern PatternType, colors ...color.Color) *Pattern {
	switch pattern {
	case Base:
		return newBasePattern(colors[0])
	case Stripe:
		return newStripePattern(colors[0], colors[1])
	case Gradient:
		return newGradientPattern(colors[0], colors[1])
	case Ring:
		return newRingPattern(colors[0], colors[1])
	case Checker:
		return newCheckerPattern(colors[0], colors[1])
	case Test:
		return newTestPattern()
	default:
		panic(fmt.Errorf("not supported pattern: %T", pattern))
	}
}

// return a constant color
func newBasePattern(c color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt:   func(t tuple.Tuple) color.Color { return c },
	}
}

// returns a striped patter with two colors
func newStripePattern(a, b color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			if math.Mod(math.Floor(point.X), 2) == 0 {
				return a
			}
			return b
		},
	}
}

// returns a gradient pattern with two colors
func newGradientPattern(a, b color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			distance := color.Subtract(b, a)
			fraction := point.X - math.Floor(point.X)
			return color.Add(a, distance.Scalar(fraction))
		},
	}
}

// returns a ring pattern with two colors
func newRingPattern(a, b color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			comp := math.Sqrt(math.Pow(point.X, 2) + math.Pow(point.Z, 2))
			if math.Mod(math.Floor(comp), 2) == 0 {
				return a
			}
			return b
		},
	}
}

// returns a checker pattern with two colors
func newCheckerPattern(a, b color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			sum := math.Floor(point.X) + math.Floor(point.Y) + math.Floor(point.Z)
			if math.Mod(sum, 2) == 0 {
				return a
			}
			return b
		},
	}
}

// returns a test pattern
func newTestPattern() *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			return color.New(point.X, point.Y, point.Z)
		},
	}
}
