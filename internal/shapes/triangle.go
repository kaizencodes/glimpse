package shapes

import (
	"fmt"
	"math"

	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

// Triangle is an atomic object that is used to build more complex shapes.
type Triangle struct {
	Model                                  Shape
	P1, P2, P3, E1, E2, N1, N2, N3, Normal tuple.Tuple
	boundingBox                            *BoundingBox
}

func (s *Triangle) String() string {
	return fmt.Sprintf("Triangle(material: %s, transform: %s)", s.Material(), s.Transform())
}

// These are defined to implement the shape interface, no need for them as we use the model's transforms and materials
func (s *Triangle) SetTransform(transform matrix.Matrix) {
}

func (s *Triangle) SetMaterial(mat *materials.Material) {
}

func (s *Triangle) Material() *materials.Material {
	return s.Model.Material()
}

func (s *Triangle) Transform() matrix.Matrix {
	return matrix.DefaultTransform()
}

func (s *Triangle) Parent() Shape {
	return s.Model
}

func (s *Triangle) SetParent(other Shape) {
}

func (s *Triangle) CalculateBoundingBox() {
	s.boundingBox.AddPoint(s.P1)
	s.boundingBox.AddPoint(s.P2)
	s.boundingBox.AddPoint(s.P3)

	TransformBoundingBox(s.boundingBox, s.Transform())
}

func (s *Triangle) BoundingBox() *BoundingBox {
	return s.boundingBox
}

func (s *Triangle) localNormalAt(point tuple.Tuple, hit Intersection) tuple.Tuple {
	if s.smooth() {
		return s.smoothLocalNormalAt(point, hit)
	}

	return s.Normal
}

func (s *Triangle) smoothLocalNormalAt(point tuple.Tuple, hit Intersection) tuple.Tuple {
	return tuple.Add(
		tuple.Add(
			s.N2.Scalar(hit.u),
			s.N3.Scalar(hit.v)),
		s.N1.Scalar(1-hit.u-hit.v))
}

func (s *Triangle) localIntersect(r *ray.Ray) Intersections {
	// Möller–Trumbore algorithm
	directionCrossE2 := tuple.Cross(r.Direction, s.E2)
	determinant := tuple.Dot(s.E1, directionCrossE2)
	xs := Intersections{}
	if math.Abs(determinant) < utils.EPSILON {
		return xs
	}

	f := 1.0 / determinant
	p1ToOrigin := tuple.Subtract(r.Origin, s.P1)
	u := f * tuple.Dot(p1ToOrigin, directionCrossE2)
	if u < 0.0 || u > 1.0 {
		return xs
	}

	originCrossE1 := tuple.Cross(p1ToOrigin, s.E1)
	v := f * tuple.Dot(r.Direction, originCrossE1)
	if v < 0.0 || (u+v) > 1.0 {
		return xs
	}

	t := f * tuple.Dot(s.E2, originCrossE1)
	return Intersections{NewIntersectionWithUV(t, u, v, Shape(s))}
}

func NewTriangle(p1, p2, p3 tuple.Tuple) *Triangle {
	e1 := tuple.Subtract(p2, p1)
	e2 := tuple.Subtract(p3, p1)
	normal := tuple.Cross(e2, e1).Normalize()

	return &Triangle{
		boundingBox: DefaultBoundingBox(),

		P1:     p1,
		P2:     p2,
		P3:     p3,
		E1:     e1,
		E2:     e2,
		Normal: normal,
	}
}

func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 tuple.Tuple) *Triangle {
	e1 := tuple.Subtract(p2, p1)
	e2 := tuple.Subtract(p3, p1)
	normal := tuple.Cross(e2, e1).Normalize()

	return &Triangle{
		boundingBox: DefaultBoundingBox(),

		P1:     p1,
		P2:     p2,
		P3:     p3,
		E1:     e1,
		E2:     e2,
		Normal: normal,
		N1:     n1,
		N2:     n2,
		N3:     n3,
	}
}

func (s *Triangle) smooth() bool {
	emptyVector := tuple.NewVector(0, 0, 0)
	return s.N1 != emptyVector || s.N2 != emptyVector || s.N3 != emptyVector
}
