package shapes

import (
	"fmt"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Plane struct {
	transform matrix.Matrix
	material  *materials.Material
}

func (s *Plane) String() string {
	return fmt.Sprintf("Plane(material: %s, transform: %s)", s.material, s.transform)
}

func (s *Plane) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Plane) SetMaterial(mat *materials.Material) {
	s.material = mat
}

func (s *Plane) Material() *materials.Material {
	return s.material
}

func (s *Plane) Transform() matrix.Matrix {
	return s.transform
}

func (s *Plane) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
	return tuple.NewVector(0, 1, 0)
}

func NewPlane() *Plane {
	return &Plane{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
	}
}
