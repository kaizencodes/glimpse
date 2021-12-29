package materials

import (
	"fmt"
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/tuple"
	"math"
)

type PatternType int

const (
	Base PatternType = iota
	Stripe
	Gradient
	Ring
	Checker
)

type Pattern struct {
	transform matrix.Matrix
	colorAt   func(tuple.Tuple) color.Color
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
	default:
		panic(fmt.Errorf("Not supported pattern: %T", pattern))
	}
}

func newBasePattern(c color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt:   func(t tuple.Tuple) color.Color { return c },
	}
}

func newStripePattern(a, b color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			a, b := a, b
			if math.Mod(math.Floor(point.X()), 2) == 0 {
				return a
			}
			return b
		},
	}
}

func newGradientPattern(a, b color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			a, b := a, b
			distance := color.Subtract(b, a)
			fraction := point.X() - math.Floor(point.X())
			return color.Add(a, distance.Scalar(fraction))
		},
	}
}

func newRingPattern(a, b color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			a, b := a, b
			comp := math.Sqrt(math.Pow(point.X(), 2) + math.Pow(point.Z(), 2))
			if math.Mod(math.Floor(comp), 2) == 0 {
				return a
			}
			return b
		},
	}
}

func newCheckerPattern(a, b color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			a, b := a, b
			sum := math.Floor(point.X()) + math.Floor(point.Y()) + math.Floor(point.Z())
			if math.Mod(math.Floor(sum), 2) == 0 {
				return a
			}
			return b
		},
	}
}
