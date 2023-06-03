package shapes

import (
	"fmt"
	"math"

	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
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

func (s *Sphere) localNormalAt(point tuple.Tuple, _hit Intersection) tuple.Tuple {
	return tuple.Subtract(point, tuple.NewPoint(0, 0, 0))
}

func (s *Sphere) localIntersect(r *ray.Ray) Intersections {
	sphere_to_ray := r.Origin.ToVector()

	a := tuple.Dot(r.Direction, r.Direction)
	b := 2 * tuple.Dot(r.Direction, sphere_to_ray)
	c := tuple.Dot(sphere_to_ray, sphere_to_ray) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return Intersections{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return Intersections{NewIntersection(t1, s), NewIntersection(t2, s)}
}

func (s *Sphere) Parent() Shape {
	return s.parent
}

func (s *Sphere) SetParent(other Shape) {
	s.parent = other
}

func NewSphere() *Sphere {
	return &Sphere{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
	}
}

// A helper for producing a sphere with a glassy material
func NewGlassSphere() *Sphere {
	mat := materials.DefaultMaterial()
	mat.Transparency = 1.0
	mat.RefractiveIndex = 1.5
	return &Sphere{
		transform: matrix.DefaultTransform(),
		material:  mat,
	}
}
