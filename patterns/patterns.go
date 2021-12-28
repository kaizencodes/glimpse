package patterns

import (
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/tuple"
	"math"
)

type Pattern struct {
	transform matrix.Matrix
	colorAt   func(tuple.Tuple) color.Color
}

func NewMonoPattern(c color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt:   func(t tuple.Tuple) color.Color { return c },
	}
}

func NewStripePattern(a, b color.Color) *Pattern {
	return &Pattern{
		transform: matrix.DefaultTransform(),
		colorAt: func(point tuple.Tuple) color.Color {
			a, b := a, b
			if math.Mod(math.Floor(point.X()), 2) == 0 {
				return a
			} else {
				return b
			}
		},
	}
}

func NewGradientPattern(a, b color.Color) *Pattern {
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

func (p *Pattern) ColorAt(point tuple.Tuple) color.Color {
	return p.colorAt(point)
}

func (s *Pattern) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Pattern) Transform() matrix.Matrix {
	return s.transform
}