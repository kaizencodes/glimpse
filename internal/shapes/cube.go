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

func (s *Cube) localNormalAt(point tuple.Tuple, _hit Intersection) tuple.Tuple {
	x, y, z := math.Abs(point.X), math.Abs(point.Y), math.Abs(point.Z)
	max := math.Max(x, math.Max(y, z))

	if max == x {
		return tuple.NewVector(point.X, 0, 0)
	} else if max == y {
		return tuple.NewVector(0, point.Y, 0)
	}
	return tuple.NewVector(0, 0, point.Z)
}

func (s *Cube) localIntersect(r *ray.Ray) Intersections {
	return aABBIntersect(s, r, tuple.NewPoint(-1, -1, -1), tuple.NewPoint(1, 1, 1))
}

func aABBIntersect(s Shape, r *ray.Ray, minPoint, maxPoint tuple.Tuple) Intersections {
	xMin, xMax := checkAxis(r.Origin.X, r.Direction.X, minPoint.X, maxPoint.X)
	yMin, yMax := checkAxis(r.Origin.Y, r.Direction.Y, minPoint.Y, maxPoint.Y)
	zMin, zMax := checkAxis(r.Origin.Z, r.Direction.Z, minPoint.Z, maxPoint.Z)

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

func checkAxis(origin, direction, min, max float64) (float64, float64) {
	minNumerator := min - origin
	maxNumerator := max - origin
	if math.Abs(direction) >= utils.EPSILON {
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
