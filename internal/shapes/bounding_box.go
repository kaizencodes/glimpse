package shapes

import (
	"math"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
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
			W: 1,
		},
		Max: tuple.Tuple{
			X: math.Inf(-1),
			Y: math.Inf(-1),
			Z: math.Inf(-1),
			W: 1,
		},
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

func (b *BoundingBox) AddBox(box *BoundingBox) {
	b.AddPoint(box.Min)
	b.AddPoint(box.Max)
}

func (b *BoundingBox) ContainsPoint(p tuple.Tuple) bool {
	return b.Min.X <= p.X && p.X <= b.Max.X &&
		b.Min.Y <= p.Y && p.Y <= b.Max.Y &&
		b.Min.Z <= p.Z && p.Z <= b.Max.Z
}

func (b *BoundingBox) ContainsBox(box *BoundingBox) bool {
	return b.ContainsPoint(box.Min) && b.ContainsPoint(box.Max)
}

func TransformBoundingBox(b *BoundingBox, m matrix.Matrix) {
	box := DefaultBoundingBox()
	points := [8]tuple.Tuple{
		tuple.NewPoint(b.Min.X, b.Min.Y, b.Min.Z),
		tuple.NewPoint(b.Min.X, b.Min.Y, b.Max.Z),
		tuple.NewPoint(b.Min.X, b.Max.Y, b.Min.Z),
		tuple.NewPoint(b.Min.X, b.Max.Y, b.Max.Z),
		tuple.NewPoint(b.Max.X, b.Min.Y, b.Min.Z),
		tuple.NewPoint(b.Max.X, b.Min.Y, b.Max.Z),
		tuple.NewPoint(b.Max.X, b.Max.Y, b.Min.Z),
		tuple.NewPoint(b.Max.X, b.Max.Y, b.Max.Z),
	}

	for _, p := range points {
		box.AddPoint(tuple.Multiply(m, p))
	}
	b.Min = box.Min
	b.Max = box.Max
}

func BoxIntersection(box *BoundingBox, r *ray.Ray) bool {
	intersections := aABBIntersect(NewTestShape(), r, box.Min, box.Max)
	return len(intersections) > 0
}

// Splits the bounding box into two even smaller boxes.
// The split is always along the longest axis.
// If the axis have the same length, the x axis is chosen.
func (b *BoundingBox) Split() (left, right *BoundingBox) {
	dx := b.Max.X - b.Min.X
	dy := b.Max.Y - b.Min.Y
	dz := b.Max.Z - b.Min.Z

	var midMin, midMax tuple.Tuple
	greatest := math.Max(dx, math.Max(dy, dz))
	switch greatest {
	case dx:
		midMin = tuple.NewPoint(b.Min.X+dx/2, b.Min.Y, b.Min.Z)
		midMax = tuple.NewPoint(b.Max.X-dx/2, b.Max.Y, b.Max.Z)
	case dy:
		midMin = tuple.NewPoint(b.Min.X, b.Min.Y+dy/2, b.Min.Z)
		midMax = tuple.NewPoint(b.Max.X, b.Max.Y-dy/2, b.Max.Z)
	case dz:
		midMin = tuple.NewPoint(b.Min.X, b.Min.Y, b.Min.Z+dz/2)
		midMax = tuple.NewPoint(b.Max.X, b.Max.Y, b.Max.Z-dz/2)
	}

	left = NewBoundingBox(b.Min, midMax)
	right = NewBoundingBox(midMin, b.Max)
	return left, right
}
