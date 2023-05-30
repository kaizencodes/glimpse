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

type Triangle struct {
	transform                              matrix.Matrix
	material                               *materials.Material
	parent                                 Shape
	P1, P2, P3, E1, E2, N1, N2, N3, Normal tuple.Tuple
}

func (s *Triangle) String() string {
	return fmt.Sprintf("Triangle(material: %s, transform: %s)", s.material, s.transform)
}

func (s *Triangle) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Triangle) SetMaterial(mat *materials.Material) {
	s.material = mat
}

func (s *Triangle) Material() *materials.Material {
	return s.material
}

func (s *Triangle) Transform() matrix.Matrix {
	return s.transform
}

func (s *Triangle) Parent() Shape {
	return s.parent
}

func (s *Triangle) SetParent(other Shape) {
	s.parent = other
}

func (s *Triangle) LocalNormalAt(point tuple.Tuple, hit Intersection) tuple.Tuple {
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
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
		P1:        p1,
		P2:        p2,
		P3:        p3,
		E1:        e1,
		E2:        e2,
		Normal:    normal,
	}
}

func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 tuple.Tuple) *Triangle {
	e1 := tuple.Subtract(p2, p1)
	e2 := tuple.Subtract(p3, p1)
	normal := tuple.Cross(e2, e1).Normalize()

	return &Triangle{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
		P1:        p1,
		P2:        p2,
		P3:        p3,
		E1:        e1,
		E2:        e2,
		Normal:    normal,
		N1:        n1,
		N2:        n2,
		N3:        n3,
	}
}

func (s *Triangle) smooth() bool {
	emptyVector := tuple.NewVector(0, 0, 0)
	return s.N1 != emptyVector || s.N2 != emptyVector || s.N3 != emptyVector
}
