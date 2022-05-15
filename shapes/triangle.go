package shapes

import (
	"fmt"
	"glimpse/calc"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/ray"
	"glimpse/tuple"
	"math"
)

type Triangle struct {
	transform                              matrix.Matrix
	material                               *materials.Material
	parent                                 Shape
	p1, p2, p3, e1, e2, n1, n2, n3, normal tuple.Tuple
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

func (s *Triangle) P1() tuple.Tuple {
	return s.p1
}

func (s *Triangle) P2() tuple.Tuple {
	return s.p2
}

func (s *Triangle) P3() tuple.Tuple {
	return s.p3
}

func (s *Triangle) E1() tuple.Tuple {
	return s.e1
}

func (s *Triangle) E2() tuple.Tuple {
	return s.e2
}

func (s *Triangle) Normal() tuple.Tuple {
	return s.normal
}

func (s *Triangle) SetP1(point tuple.Tuple) {
	s.p1 = point
}

func (s *Triangle) SetP2(point tuple.Tuple) {
	s.p2 = point
}

func (s *Triangle) SetP3(point tuple.Tuple) {
	s.p3 = point
}

func (s *Triangle) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
	return s.normal
}

func (s *Triangle) LocalIntersect(r *ray.Ray) Intersections {
	// Möller–Trumbore algorithm
	directionCrossE2 := tuple.Cross(r.Direction(), s.E2())
	determinant := tuple.Dot(s.E1(), directionCrossE2)
	xs := Intersections{}
	if math.Abs(determinant) < calc.EPSILON {
		return xs
	}

	f := 1.0 / determinant
	p1ToOrigin := tuple.Subtract(r.Origin(), s.P1())
	u := f * tuple.Dot(p1ToOrigin, directionCrossE2)
	if u < 0.0 || u > 1.0 {
		return xs
	}

	originCrossE1 := tuple.Cross(p1ToOrigin, s.E1())
	v := f * tuple.Dot(r.Direction(), originCrossE1)
	if v < 0.0 || (u+v) > 1.0 {
		return xs
	}

	t := f * tuple.Dot(s.E2(), originCrossE1)
	return Intersections{NewIntersection(t, s)}
}

func NewTriangle(p1, p2, p3 tuple.Tuple) *Triangle {
	e1 := tuple.Subtract(p2, p1)
	e2 := tuple.Subtract(p3, p1)
	normal := tuple.Cross(e2, e1).Normalize()

	return &Triangle{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
		p1:        p1,
		p2:        p2,
		p3:        p3,
		e1:        e1,
		e2:        e2,
		normal:    normal,
	}
}
