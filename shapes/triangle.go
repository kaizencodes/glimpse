package shapes

import (
	"fmt"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Triangle struct {
	transform                              matrix.Matrix
	material                               *materials.Material
	parent                                 Shape
	p1, p2, p3, e1, e2, n1, n2, n3, normal tuple.Tuple
}

func (t *Triangle) String() string {
	return fmt.Sprintf("Triangle(material: %s, transform: %s)", t.material, t.transform)
}

func (t *Triangle) SetTransform(transform matrix.Matrix) {
	t.transform = transform
}

func (t *Triangle) SetMaterial(mat *materials.Material) {
	t.material = mat
}

func (t *Triangle) Material() *materials.Material {
	return t.material
}

func (t *Triangle) Transform() matrix.Matrix {
	return t.transform
}

func (s *Triangle) Parent() Shape {
	return s.parent
}

func (s *Triangle) SetParent(other Shape) {
	s.parent = other
}

func (t *Triangle) P1() tuple.Tuple {
	return t.p1
}

func (t *Triangle) P2() tuple.Tuple {
	return t.p2
}

func (t *Triangle) P3() tuple.Tuple {
	return t.p3
}

func (t *Triangle) E1() tuple.Tuple {
	return t.e1
}

func (t *Triangle) E2() tuple.Tuple {
	return t.e2
}

func (t *Triangle) Normal() tuple.Tuple {
	return t.normal
}

func (t *Triangle) SetP1(point tuple.Tuple) {
	t.p1 = point
}

func (t *Triangle) SetP2(point tuple.Tuple) {
	t.p2 = point
}

func (t *Triangle) SetP3(point tuple.Tuple) {
	t.p3 = point
}

func (t *Triangle) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
	return t.normal
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
