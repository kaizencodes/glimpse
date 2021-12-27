package patterns

import (
	"fmt"
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/tuple"
	"math"
)

type Pattern interface {
	ColorAt(point tuple.Tuple) color.Color
	Transform() matrix.Matrix
	SetTransform(transform matrix.Matrix)
	String() string
}

type StripePattern struct {
	a, b      color.Color
	transform matrix.Matrix
}

func NewStripePattern(a, b color.Color) *StripePattern {
	return &StripePattern{a: a, b: b, transform: matrix.DefaultTransform()}
}

func (p *StripePattern) ColorAt(point tuple.Tuple) color.Color {
	if math.Mod(math.Floor(point.X()), 2) == 0 {
		return p.a
	} else {
		return p.b
	}
}

func (p *StripePattern) String() string {
	return fmt.Sprintf("StripePattern(a: %s, b: %s)", p.a, p.b)
}

type MonoPattern struct {
	color     color.Color
	transform matrix.Matrix
}

func NewMonoPattern(c color.Color) *MonoPattern {
	return &MonoPattern{color: c, transform: matrix.DefaultTransform()}
}

func (p *MonoPattern) ColorAt(point tuple.Tuple) color.Color {
	return p.color
}

func (p *MonoPattern) String() string {
	return fmt.Sprintf("MonoPattern(color: %s)", p.color)
}

func (s *MonoPattern) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *MonoPattern) Transform() matrix.Matrix {
	return s.transform
}

func (s *StripePattern) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *StripePattern) Transform() matrix.Matrix {
	return s.transform
}
