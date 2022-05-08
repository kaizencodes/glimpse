package ray

import (
	"fmt"
	"glimpse/calc"
	"glimpse/matrix"
	"glimpse/shapes"
	"glimpse/tuple"
	"math"
	"sort"
	"strconv"
)

const BounceLimit = 5

type Ray struct {
	origin      tuple.Tuple
	direction   tuple.Tuple
	bounceLimit int
}

func (r *Ray) Position(dist float64) tuple.Tuple {
	return tuple.Add(r.origin, r.direction.Scalar(dist))
}

func (r *Ray) String() string {
	return fmt.Sprintf("Ray(origin: %s, direction: %s)", r.origin, r.direction)
}

func (r *Ray) Equal(other *Ray) bool {
	return r.origin.Equal(other.origin) && r.direction.Equal(other.direction)
}

func (r *Ray) Translate(x, y, z float64) *Ray {
	origin, err := tuple.Multiply(matrix.Translation(x, y, z), r.origin)
	if err != nil {
		panic(err)
	}
	return &Ray{origin: origin, direction: r.direction}
}

func (r *Ray) Scale(x, y, z float64) *Ray {
	origin, err := tuple.Multiply(matrix.Scaling(x, y, z), r.origin)
	if err != nil {
		panic(err)
	}
	direction, err := tuple.Multiply(matrix.Scaling(x, y, z), r.direction)
	if err != nil {
		panic(err)
	}
	return &Ray{origin: origin, direction: direction}
}

func (r *Ray) Origin() tuple.Tuple {
	return r.origin
}

func (r *Ray) Direction() tuple.Tuple {
	return r.direction
}

func (r *Ray) BounceLimit() int {
	return r.bounceLimit
}

func (r *Ray) SetBounceLimit(bounceLimit int) {
	r.bounceLimit = bounceLimit
}

func (r *Ray) Intersect(s shapes.Shape) Intersections {
	transform, err := s.Transform().Inverse()
	if err != nil {
		panic(err)
	}
	origin, _ := tuple.Multiply(transform, r.origin)
	direction, _ := tuple.Multiply(transform, r.direction)
	localRay := &Ray{origin, direction, BounceLimit}

	switch s := s.(type) {
	case *shapes.Sphere:
		return localRay.intersectSphere(s)
	case *shapes.Cylinder:
		return localRay.intersectCylinder(s)
	case *shapes.Plane:
		return localRay.intersectPlane(s)
	case *shapes.Cube:
		return localRay.intersectCube(s)
	case *shapes.Group:
		return localRay.intersectGroup(s)
	case *shapes.Triangle:
		return localRay.intersectTriangle(s)
	default:
		panic(fmt.Errorf("not supported shape %T", s))
	}
}

func (r *Ray) intersectSphere(s *shapes.Sphere) Intersections {
	sphere_to_ray := r.origin.ToVector()

	a := tuple.Dot(r.direction, r.direction)
	b := 2 * tuple.Dot(r.direction, sphere_to_ray)
	c := tuple.Dot(sphere_to_ray, sphere_to_ray) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return Intersections{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return Intersections{Intersection{t: t1, shape: s}, Intersection{t: t2, shape: s}}
}

func (r *Ray) intersectGroup(s *shapes.Group) Intersections {
	xs := Intersections{}
	for _, child := range s.Children() {
		xs = append(xs, r.Intersect(child)...)
	}
	xs.Sort()
	return xs
}

func (r *Ray) intersectTriangle(s *shapes.Triangle) Intersections {
	// Möller–Trumbore algorithm
	directionCrossE2 := tuple.Cross(r.Direction(), s.E2())
	determinant := tuple.Dot(s.E1(), directionCrossE2)
	xs := Intersections{}
	if math.Abs(determinant) < calc.EPSILON {
		return xs
	}

	f := 1.0 / determinant
	p1ToOrigin := tuple.Subtract(r.Origin(), s.P1())
	u := f * tuple.Dot(p1ToOrigin, directionCrossE2)
	if u < 0.0 || u > 1.0 {
		return xs
	}

	originCrossE1 := tuple.Cross(p1ToOrigin, s.E1())
	v := f * tuple.Dot(r.Direction(), originCrossE1)
	if v < 0.0 || (u+v) > 1.0 {
		return xs
	}

	t := f * tuple.Dot(s.E2(), originCrossE1)
	return Intersections{Intersection{t, s}}
}

func (r *Ray) intersectCylinder(s *shapes.Cylinder) Intersections {
	a := math.Pow(r.direction.X(), 2) + math.Pow(r.direction.Z(), 2)
	if calc.FloatEquals(a, 0.0) {
		return intersectionsForCaps(Intersections{}, s, r)
	}

	b := 2*r.origin.X()*r.direction.X() + 2*r.origin.Z()*r.direction.Z()
	c := math.Pow(r.origin.X(), 2) + math.Pow(r.origin.Z(), 2) - 1

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

	y0 := r.Origin().Y() + t0*r.direction.Y()
	if s.Minimum() < y0 && y0 < s.Maximum() {
		xs = append(xs, Intersection{t: t0, shape: s})
	}

	y1 := r.Origin().Y() + t1*r.direction.Y()
	if s.Minimum() < y1 && y1 < s.Maximum() {
		xs = append(xs, Intersection{t: t1, shape: s})
	}

	return intersectionsForCaps(xs, s, r)
}

func intersectionsForCaps(xs Intersections, s *shapes.Cylinder, r *Ray) Intersections {
	// caps only matter if the cylinder is closed, and might possibly be intersected by the ray.
	if !(s.Closed() || calc.FloatEquals(r.Direction().Y(), 0)) {
		return xs
	}
	// check for an intersection with the lower end cap by intersecting the ray with the plane at y=s.minimum
	t := (s.Minimum() - r.origin.Y()) / r.direction.Y()
	if checkCap(r, t) {
		xs = append(xs, Intersection{t: t, shape: s})
	}

	// check for an intersection with the upper end cap by intersecting the ray with the plane at y=cyl.maximum
	t = (s.Maximum() - r.origin.Y()) / r.direction.Y()
	if checkCap(r, t) {
		xs = append(xs, Intersection{t: t, shape: s})
	}
	return xs
}

// checks to see if the intersection at `t` is within a radius
//  of 1 (the radius of your cylinders) from the y axis.
func checkCap(r *Ray, t float64) bool {
	x := r.Origin().X() + t*r.Direction().X()
	z := r.Origin().Z() + t*r.Direction().Z()
	return math.Pow(x, 2)+math.Pow(z, 2) <= 1
}

func (r *Ray) intersectPlane(s *shapes.Plane) Intersections {
	if math.Abs(r.Direction().Y()) < calc.EPSILON {
		return Intersections{}
	}

	t := -r.Origin().Y() / r.Direction().Y()
	return Intersections{
		Intersection{t: t, shape: s},
	}
}

func (r *Ray) intersectCube(s *shapes.Cube) Intersections {
	xMin, xMax := checkAxis(r.origin.X(), r.direction.X())
	yMin, yMax := checkAxis(r.origin.Y(), r.direction.Y())
	zMin, zMax := checkAxis(r.origin.Z(), r.direction.Z())

	min := math.Max(xMin, math.Max(yMin, zMin))
	max := math.Min(xMax, math.Min(yMax, zMax))

	if min > max {
		return Intersections{}
	}

	return Intersections{
		Intersection{t: min, shape: s},
		Intersection{t: max, shape: s},
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

func New(origin, direction tuple.Tuple) *Ray {
	return &Ray{origin, direction, BounceLimit}
}

type Intersection struct {
	t     float64
	shape shapes.Shape
}

type Intersections []Intersection

func (inter Intersection) Empty() bool {
	return inter.t == math.MaxFloat64
}

func (inter Intersection) T() float64 {
	return inter.t
}

func (inter Intersection) Shape() shapes.Shape {
	return inter.shape
}

func (c Intersections) String() string {
	var result string

	for _, section := range c {
		result += strconv.FormatFloat(section.t, 'f', -1, 64) + ", "
	}
	return result
}

func (c Intersections) Sort() {
	sort.Slice(c, func(i, j int) bool {
		return c[i].t < c[j].t
	})
}

func (c Intersections) Hit() Intersection {
	res := Intersection{t: math.MaxFloat64}
	for _, val := range c {
		if val.t < 0 {
			continue
		}
		if val.t < res.t {
			res = val
		}
	}
	return res
}

func NewIntersection(t float64, obj shapes.Shape) Intersection {
	return Intersection{t, obj}
}
