package objects

import (
	"fmt"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Sphere struct {
	center    tuple.Tuple
	radius    float64
	transform matrix.Matrix
}

func NewSphere() *Sphere {
	return &Sphere{
		center:    tuple.NewPoint(0, 0, 0),
		radius:    1,
		transform: matrix.NewIdentity(4),
	}
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Shpere(center: %s, radius: %f, transform: %s)", s.center, s.radius, s.transform)
}

func (s *Sphere) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Sphere) GetTransform() matrix.Matrix {
	return s.transform
}

func (s *Sphere) Normal(worldPoint tuple.Tuple) tuple.Tuple {
	inv_mat, err := s.transform.Inverse()
	if err != nil {
		panic(err)
	}
	objectPoint, _ := tuple.Multiply(inv_mat, worldPoint)
	objectNormal := tuple.Subtract(objectPoint, s.center)
	worldNormal, _ := tuple.Multiply(inv_mat.Transpose(), objectNormal)
	return worldNormal.ToVector().Normalize()
}
