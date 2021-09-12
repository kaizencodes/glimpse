package ray

import (
	"fmt"
	"glimpse/matrix"
	"glimpse/objects"
	"glimpse/tuple"
	"math"
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

func New(origin, direction tuple.Tuple) Ray {
	return Ray{origin, direction}
}

type Intersection struct {
	t      float64
	object *objects.Sphere
}

type Intersections []Intersection

func (inter Intersection) Empty() bool {
	return inter.t == math.MaxFloat64
}

func (inter Intersection) GetT() float64 {
	return inter.t
}

func (inter Intersection) GetObject() *objects.Sphere {
	return inter.object
}

func (c Intersections) String() string {
	var result string

	for _, section := range c {
		result += strconv.FormatFloat(section.t, 'f', -1, 64) + ", "
	}
	return result
}

func Intersect(r Ray, s *objects.Sphere) Intersections {
	transform, err := s.Transform().Inverse()
	if err != nil {
		panic(err)
	}
	origin, _ := tuple.Multiply(transform, r.origin)
	direction, _ := tuple.Multiply(transform, r.direction)
	ray2 := Ray{origin, direction}

	sphere_to_ray := tuple.Subtract(ray2.origin, tuple.NewPoint(0, 0, 0))

	a := tuple.Dot(ray2.direction, ray2.direction)
	b := 2 * tuple.Dot(ray2.direction, sphere_to_ray)
	c := tuple.Dot(sphere_to_ray, sphere_to_ray) - 1

	disciminant := math.Pow(b, 2) - 4*a*c

	if disciminant < 0 {
		return Intersections{}
	}

	t1 := (-b - math.Sqrt(disciminant)) / (2 * a)
	t2 := (-b + math.Sqrt(disciminant)) / (2 * a)

	return Intersections{Intersection{t: t1, object: s}, Intersection{t: t2, object: s}}
}

func Hit(coll Intersections) Intersection {
	res := Intersection{t: math.MaxFloat64}
	for _, val := range coll {
		if val.t < 0 {
			continue
		}
		if val.t < res.t {
			res = val
		}
	}
	return res
}
