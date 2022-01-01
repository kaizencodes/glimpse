package shapes

import (
	"fmt"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Cube struct {
	transform matrix.Matrix
	material  *materials.Material
}

func (s *Cube) String() string {
	return fmt.Sprintf("Cube(material: %s, transform: %s)", s.material, s.transform)
}

func (s *Cube) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Cube) SetMaterial(mat *materials.Material) {
	s.material = mat
}

func (s *Cube) Material() *materials.Material {
	return s.material
}

func (s *Cube) Transform() matrix.Matrix {
	return s.transform
}

func (s *Cube) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
	return tuple.NewVector(0, 1, 0)
}

func NewCube() *Cube {
	return &Cube{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
	}
}
