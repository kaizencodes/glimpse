package shapes

import (
	"fmt"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Sphere struct {
	transform matrix.Matrix
	material  *materials.Material
	parent    Shape
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Sphere(transform: %s, material: %s)", s.transform, s.material)
}

func (s *Sphere) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Sphere) SetMaterial(mat *materials.Material) {
	s.material = mat
}

func (s *Sphere) Material() *materials.Material {
	return s.material
}

func (s *Sphere) Transform() matrix.Matrix {
	return s.transform
}

func (s *Sphere) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
	return tuple.Subtract(point, tuple.NewPoint(0, 0, 0))
}

func NewSphere() *Sphere {
	return &Sphere{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
	}
}

func NewGlassSphere() *Sphere {
	mat := materials.DefaultMaterial()
	mat.SetTransparency(1)
	mat.SetRefractiveIndex(1.5)
	return &Sphere{
		transform: matrix.DefaultTransform(),
		material:  mat,
	}
}
