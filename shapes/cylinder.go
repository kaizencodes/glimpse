package shapes

import (
	"fmt"
	"glimpse/calc"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
	"math"
)

type Cylinder struct {
	transform        matrix.Matrix
	material         *materials.Material
	minimum, maximum float64
	closed           bool
	parent           Shape
}

func (s *Cylinder) String() string {
	return fmt.Sprintf("Cylinder(min: %f, max: %f, transform: %s, material: %s)", s.minimum, s.maximum, s.transform, s.material)
}

func (s *Cylinder) Parent() Shape {
	return s.parent
}

func (s *Cylinder) SetParent(other Shape) {
	s.parent = other
}

func (s *Cylinder) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Cylinder) SetMaterial(mat *materials.Material) {
	s.material = mat
}

func (s *Cylinder) SetMinimum(min float64) {
	s.minimum = min
}

func (s *Cylinder) SetMaximum(max float64) {
	s.maximum = max
}

func (s *Cylinder) SetClosed(closed bool) {
	s.closed = closed
}

func (s *Cylinder) Minimum() float64 {
	return s.minimum
}

func (s *Cylinder) Maximum() float64 {
	return s.maximum
}

func (s *Cylinder) Closed() bool {
	return s.closed
}

func (s *Cylinder) Material() *materials.Material {
	return s.material
}

func (s *Cylinder) Transform() matrix.Matrix {
	return s.transform
}

func (s *Cylinder) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
	// compute the square of the distance from the y axis.
	dist := math.Pow(point.X(), 2) + math.Pow(point.Z(), 2)

	if dist < 1 && point.Y() >= s.Maximum()-calc.EPSILON {
		return tuple.NewVector(0, 1, 0)
	} else if dist < 1 && point.Y() <= s.Minimum()+calc.EPSILON {
		return tuple.NewVector(0, -1, 0)
	}

	return tuple.NewVector(point.X(), 0, point.Z())
}

func NewCylinder() *Cylinder {
	return &Cylinder{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
		minimum:   -math.MaxFloat64,
		maximum:   math.MaxFloat64,
		closed:    false,
	}
}
