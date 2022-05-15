package shapes

import (
	"fmt"
	"glimpse/calc"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/ray"
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

func (s *Cube) LocalIntersect(r *ray.Ray) Intersections {
	xMin, xMax := checkAxis(r.Origin().X(), r.Direction().X())
	yMin, yMax := checkAxis(r.Origin().Y(), r.Direction().Y())
	zMin, zMax := checkAxis(r.Origin().Z(), r.Direction().Z())

	min := math.Max(xMin, math.Max(yMin, zMin))
	max := math.Min(xMax, math.Min(yMax, zMax))

	if min > max {
		return Intersections{}
	}

	return Intersections{
		NewIntersection(min, s),
		NewIntersection(max, s),
	}
}

func checkAxis(origin, direction float64) (min, max float64) {
	minNumerator := -1 - origin
	maxNumerator := 1 - origin
	if math.Abs(direction) >= calc.EPSILON {
		min = minNumerator / direction
		max = maxNumerator / direction
	} else {
		min = minNumerator * math.MaxFloat64
		max = maxNumerator * math.MaxFloat64
	}

	if min > max {
		min, max = max, min
	}

	return min, max
}

func (s *Cube) Parent() Shape {
	return s.parent
}

func (s *Cube) SetParent(other Shape) {
	s.parent = other
}

func NewCube() *Cube {
	return &Cube{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
	}
}
