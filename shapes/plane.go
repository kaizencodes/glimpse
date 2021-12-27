package shapes

import (
	"fmt"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Plane struct {
	center    tuple.Tuple
	transform matrix.Matrix
	material  *Material
}

func (s *Plane) String() string {
	return fmt.Sprintf("Plane(center: %s, material: %s, transform: %s)", s.center, s.material, s.transform)
}

func (s *Plane) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Plane) SetMaterial(mat *Material) {
	s.material = mat
}

func (s *Plane) Material() *Material {
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
		center:    tuple.NewPoint(0, 0, 0),
		transform: matrix.DefaultTransform(),
		material:  DefaultMaterial(),
	}
}
