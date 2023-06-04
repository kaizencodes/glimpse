package shapes

import (
	"math"

	"github.com/kaizencodes/glimpse/internal/matrix"
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

// Creates a bounding box for a shape, in object space
func BoundFor(shape Shape) *BoundingBox {
	box := DefaultBoundingBox()
	switch s := shape.(type) {
	case *Sphere:
		box.Min = tuple.NewPoint(-1, -1, -1)
		box.Max = tuple.NewPoint(1, 1, 1)
	case *Plane:
		box.Min = tuple.NewPoint(math.Inf(-1), 0, math.Inf(-1))
		box.Max = tuple.NewPoint(math.Inf(1), 0, math.Inf(1))
	case *Cube:
		box.Min = tuple.NewPoint(-1, -1, -1)
		box.Max = tuple.NewPoint(1, 1, 1)
	case *Cylinder:
		box.Min = tuple.NewPoint(-1, math.Max(s.Minimum, math.Inf(-1)), -1)
		box.Max = tuple.NewPoint(1, math.Min(s.Maximum, math.Inf(1)), 1)
	case *Triangle:
		box.AddPoint(s.P1)
		box.AddPoint(s.P2)
		box.AddPoint(s.P3)
	case *TestShape:
		box.Min = tuple.NewPoint(-1, -1, -1)
		box.Max = tuple.NewPoint(1, 1, 1)
	}

	return box
}

func Transform(b *BoundingBox, m matrix.Matrix) *BoundingBox {
	box := DefaultBoundingBox()
	points := []tuple.Tuple{
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
	return box
}
