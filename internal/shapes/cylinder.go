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

func (s *Cylinder) LocalNormalAt(point tuple.Tuple, _hit Intersection) tuple.Tuple {
	// compute the square of the distance from the y axis.
	dist := math.Pow(point.X(), 2) + math.Pow(point.Z(), 2)

	if dist < 1 && point.Y() >= s.Maximum()-utils.EPSILON {
		return tuple.NewVector(0, 1, 0)
	} else if dist < 1 && point.Y() <= s.Minimum()+utils.EPSILON {
		return tuple.NewVector(0, -1, 0)
	}

	return tuple.NewVector(point.X(), 0, point.Z())
}

func (s *Cylinder) LocalIntersect(r *ray.Ray) Intersections {
	a := math.Pow(r.Direction().X(), 2) + math.Pow(r.Direction().Z(), 2)
	if utils.FloatEquals(a, 0.0) {
		return s.intersectionsForCaps(Intersections{}, r)
	}

	b := 2*r.Origin().X()*r.Direction().X() + 2*r.Origin().Z()*r.Direction().Z()
	c := math.Pow(r.Origin().X(), 2) + math.Pow(r.Origin().Z(), 2) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return Intersections{}
	}

	t0 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t1 := (-b + math.Sqrt(discriminant)) / (2 * a)

	xs := Intersections{}

	if t0 > t1 {
		t0, t1 = t1, t0
	}

	y0 := r.Origin().Y() + t0*r.Direction().Y()
	if s.Minimum() < y0 && y0 < s.Maximum() {
		xs = append(xs, NewIntersection(t0, s))
	}

	y1 := r.Origin().Y() + t1*r.Direction().Y()
	if s.Minimum() < y1 && y1 < s.Maximum() {
		xs = append(xs, NewIntersection(t1, s))
	}

	return s.intersectionsForCaps(xs, r)
}

func (s *Cylinder) intersectionsForCaps(xs Intersections, r *ray.Ray) Intersections {
	// caps only matter if the cylinder is closed, and might possibly be intersected by the ray.
	if !(s.Closed() || utils.FloatEquals(r.Direction().Y(), 0)) {
		return xs
	}
	// check for an intersection with the lower end cap by intersecting the ray with the plane at y=s.minimum
	t := (s.Minimum() - r.Origin().Y()) / r.Direction().Y()
	if checkCap(r, t) {
		xs = append(xs, NewIntersection(t, s))
	}

	// check for an intersection with the upper end cap by intersecting the ray with the plane at y=cyl.maximum
	t = (s.Maximum() - r.Origin().Y()) / r.Direction().Y()
	if checkCap(r, t) {
		xs = append(xs, NewIntersection(t, s))
	}
	return xs
}

// checks to see if the intersection at `t` is within a radius
//  of 1 (the radius of your cylinders) from the y axis.
func checkCap(r *ray.Ray, t float64) bool {
	x := r.Origin().X() + t*r.Direction().X()
	z := r.Origin().Z() + t*r.Direction().Z()
	return math.Pow(x, 2)+math.Pow(z, 2) <= 1
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
