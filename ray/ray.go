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
	case *shapes.Plane:
		return localRay.intersectPlane(s)
	case *shapes.Cube:
		return localRay.intersectCube(s)
	default:
		panic(fmt.Errorf("Not supported shape %T", s))
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
