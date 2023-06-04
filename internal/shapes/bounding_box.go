package shapes

import (
	"math"

	"github.com/kaizencodes/glimpse/internal/tuple"
)

type BoundingBox struct {
	Min, Max tuple.Tuple
}

func NewBoundingBox(min, max tuple.Tuple) *BoundingBox {
	return &BoundingBox{Min: min, Max: max}
}

func DefaultBoundingBox() *BoundingBox {
	return &BoundingBox{
		Min: tuple.Tuple{
			X: math.Inf(1),
			Y: math.Inf(1),
			Z: math.Inf(1),
		},
		Max: tuple.Tuple{X: math.Inf(-1), Y: math.Inf(-1), Z: math.Inf(-1)},
	}
}

func (b *BoundingBox) AddPoint(p tuple.Tuple) {
	b.Min.X = math.Min(b.Min.X, p.X)
	b.Min.Y = math.Min(b.Min.Y, p.Y)
	b.Min.Z = math.Min(b.Min.Z, p.Z)

	b.Max.X = math.Max(b.Max.X, p.X)
	b.Max.Y = math.Max(b.Max.Y, p.Y)
	b.Max.Z = math.Max(b.Max.Z, p.Z)
}
