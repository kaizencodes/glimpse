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

type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

func (r Ray) Position(dist float64) tuple.Tuple {
	return tuple.Add(r.origin, r.direction.Scalar(dist))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(origin: %s, direction: %s)", r.origin, r.direction)
}

func (r Ray) Equal(other Ray) bool {
	return r.origin.Equal(other.origin) && r.direction.Equal(other.direction)
}

func (r Ray) Translate(x, y, z float64) Ray {
	origin, err := tuple.Multiply(matrix.Translation(x, y, z), r.origin)
	if err != nil {
		panic(err)
	}
	return Ray{origin: origin, direction: r.direction}
}

func (r Ray) Scale(x, y, z float64) Ray {
	origin, err := tuple.Multiply(matrix.Scaling(x, y, z), r.origin)
	if err != nil {
		panic(err)
	}
	direction, err := tuple.Multiply(matrix.Scaling(x, y, z), r.direction)
	if err != nil {
		panic(err)
	}
	return Ray{origin: origin, direction: direction}
}

func (r Ray) Origin() tuple.Tuple {
	return r.origin
}

func (r Ray) Direction() tuple.Tuple {
	return r.direction
}

func (r Ray) Intersect(s shapes.Shape) Intersections {
	transform, err := s.Transform().Inverse()
	if err != nil {
		panic(err)
	}
	origin, _ := tuple.Multiply(transform, r.origin)
	direction, _ := tuple.Multiply(transform, r.direction)
	localRay := Ray{origin, direction}

	switch s := s.(type) {
	case *shapes.Sphere:
		return localRay.intersectSphere(s)
	case *shapes.Plane:
		return localRay.intersectPlane(s)
	default:
		panic(fmt.Errorf("Not supported shape %T", s))
	}
}

func (r Ray) intersectSphere(s *shapes.Sphere) Intersections {
	sphere_to_ray := tuple.Subtract(r.origin, tuple.NewPoint(0, 0, 0))

	a := tuple.Dot(r.direction, r.direction)
	b := 2 * tuple.Dot(r.direction, sphere_to_ray)
	c := tuple.Dot(sphere_to_ray, sphere_to_ray) - 1

	disciminant := math.Pow(b, 2) - 4*a*c

	if disciminant < 0 {
		return Intersections{}
	}

	t1 := (-b - math.Sqrt(disciminant)) / (2 * a)
	t2 := (-b + math.Sqrt(disciminant)) / (2 * a)

	return Intersections{Intersection{t: t1, shape: s}, Intersection{t: t2, shape: s}}
}

func (r Ray) intersectPlane(s *shapes.Plane) Intersections {
	if math.Abs(r.Direction().Y()) < calc.EPSILON {
		return Intersections{}
	}

	t := -r.Origin().Y() / r.Direction().Y()
	return Intersections{
		Intersection{t: t, shape: s},
	}
}

func New(origin, direction tuple.Tuple) Ray {
	return Ray{origin, direction}
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

type Computations struct {
	t         float64
	shape     shapes.Shape
	point     tuple.Tuple
	eyeV      tuple.Tuple
	normalV   tuple.Tuple
	overPoint tuple.Tuple
	inside    bool
}

func (c Computations) T() float64 {
	return c.t
}

func (c Computations) Shape() shapes.Shape {
	return c.shape
}

func (c Computations) Point() tuple.Tuple {
	return c.point
}

func (c Computations) EyeV() tuple.Tuple {
	return c.eyeV
}

func (c Computations) NormalV() tuple.Tuple {
	return c.normalV
}

func (c Computations) OverPoint() tuple.Tuple {
	return c.overPoint
}

func (c Computations) Inside() bool {
	return c.inside
}

func PrepareComputations(i Intersection, r Ray) Computations {
	point := r.Position(i.t)
	normalV := shapes.NormalAt(point, i.shape)
	eyeV := r.Direction().Negate()

	var inside bool
	if tuple.Dot(normalV, eyeV) < 0 {
		inside = true
		normalV = normalV.Negate()
	} else {
		inside = false
	}
	overPoint := tuple.Add(point, normalV.Scalar(calc.EPSILON))

	return Computations{
		t:         i.T(),
		shape:     i.Shape(),
		point:     point,
		eyeV:      eyeV,
		normalV:   normalV,
		inside:    inside,
		overPoint: overPoint,
	}
}
