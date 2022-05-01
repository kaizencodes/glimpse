package shapes

import (
	"fmt"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Triangle struct {
	transform               matrix.Matrix
	material                *materials.Material
	parent                  Shape
	a, b, c, e1, e2, normal tuple.Tuple
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

func (t *Triangle) A() tuple.Tuple {
	return t.a
}

func (t *Triangle) B() tuple.Tuple {
	return t.b
}

func (t *Triangle) C() tuple.Tuple {
	return t.c
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

func (t *Triangle) SetA(point tuple.Tuple) {
	t.a = point
}

func (t *Triangle) SetB(point tuple.Tuple) {
	t.b = point
}

func (t *Triangle) SetC(point tuple.Tuple) {
	t.c = point
}

func (t *Triangle) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
	return t.normal
}

func NewTriangle(a, b, c tuple.Tuple) *Triangle {
	e1 := tuple.Subtract(b, a)
	e2 := tuple.Subtract(c, a)
	normal := tuple.Cross(e2, e1).Normalize()

	return &Triangle{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
		a:         a,
		b:         b,
		c:         c,
		e1:        e1,
		e2:        e2,
		normal:    normal,
	}
}
