package shapes

import (
	"fmt"

	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type TestShape struct {
	transform matrix.Matrix
	material  *materials.Material
	parent    Shape
}

func (s *TestShape) String() string {
	return fmt.Sprintf("Shape(material: %s\n, transform: %s)", s.Material(), s.Transform())
}

func (s *TestShape) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *TestShape) SetMaterial(mat *materials.Material) {
	s.material = mat
}

func (s *TestShape) Material() *materials.Material {
	return s.material
}

func (s *TestShape) Transform() matrix.Matrix {
	return s.transform
}

func (s *TestShape) Parent() Shape {
	return s.parent
}

func (s *TestShape) SetParent(other Shape) {
	s.parent = other
}

func (s *TestShape) LocalNormalAt(point tuple.Tuple, _hit Intersection) tuple.Tuple {
	return point.ToVector()
}

func (s *TestShape) localIntersect(r *ray.Ray) Intersections {
	return Intersections{}
}

func NewTestShape() *TestShape {
	return &TestShape{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
	}
}
