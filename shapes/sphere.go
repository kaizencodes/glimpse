package shapes

import (
	"fmt"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Sphere struct {
	center    tuple.Tuple
	radius    float64
	transform matrix.Matrix
	material  *Material
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Shpere(center: %s, radius: %f, transform: %s)", s.center, s.radius, s.transform)
}

func (s *Sphere) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Sphere) SetMaterial(mat *Material) {
	s.material = mat
}

func (s *Sphere) Material() *Material {
	return s.material
}

func (s *Sphere) Transform() matrix.Matrix {
	return s.transform
}

func (s *Sphere) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
	return tuple.Subtract(point, s.center)
}

func NewSphere() *Sphere {
	return &Sphere{
		center:    tuple.NewPoint(0, 0, 0),
		radius:    1,
		transform: matrix.DefaultTransform(),
		material:  DefaultMaterial(),
	}
}
