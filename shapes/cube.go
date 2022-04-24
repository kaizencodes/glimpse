package shapes

import (
	"fmt"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
	"math"
)

type Cube struct {
	transform matrix.Matrix
	material  *materials.Material
	parent    Shape
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
	x, y, z := math.Abs(point.X()), math.Abs(point.Y()), math.Abs(point.Z())
	max := math.Max(x, math.Max(y, z))

	if max == x {
		return tuple.NewVector(point.X(), 0, 0)
	} else if max == y {
		return tuple.NewVector(0, point.Y(), 0)
	}
	return tuple.NewVector(0, 0, point.Z())
}

func NewCube() *Cube {
	return &Cube{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
	}
}
