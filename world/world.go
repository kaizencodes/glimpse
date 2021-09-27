package world

import (
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/objects"
	"glimpse/ray"
	"glimpse/tuple"
)

type World struct {
	objects []objects.Object
	light   ray.Light
}

func Default() World {
	o1 := objects.NewSphere()
	o1.SetMaterial(objects.NewMaterial(color.New(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0))
	o2 := objects.NewSphere()
	o2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))

	return World{
		objects: []objects.Object{
			objects.Object(o1), objects.Object(o2),
		},
		light: ray.NewLight(tuple.NewPoint(-10, 10, -10), color.New(1, 1, 1)),
	}
}

func New(objects []objects.Object, light ray.Light) World {
	return World{objects, light}
}

func (w World) Intersect(r ray.Ray) ray.Intersections {
	coll := ray.Intersections{}
	for _, o := range w.objects {
		coll = append(coll, ray.Intersect(r, o)...)
	}
	coll.Sort()

	return coll
}
