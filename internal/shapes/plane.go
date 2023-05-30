package shapes

import (
	"fmt"
	"math"

	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

type Plane struct {
	transform matrix.Matrix
	material  *materials.Material
	parent    Shape
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

func (s *Plane) LocalNormalAt(point tuple.Tuple, _hit Intersection) tuple.Tuple {
	return tuple.NewVector(0, 1, 0)
}

func (s *Plane) localIntersect(r *ray.Ray) Intersections {
	if math.Abs(r.Direction.Y) < utils.EPSILON {
		return Intersections{}
	}

	t := -r.Origin.Y / r.Direction.Y
	return Intersections{
		NewIntersection(t, s),
	}
}

func (s *Plane) Parent() Shape {
	return s.parent
}

func (s *Plane) SetParent(other Shape) {
	s.parent = other
}

func NewPlane() *Plane {
	return &Plane{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
	}
}
