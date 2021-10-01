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

func (w *World) Objects() []objects.Object {
	return w.objects
}

func (w *World) Light() ray.Light {
	return w.light
}

func (w *World) SetLight(l ray.Light) {
	w.light = l
}

func (w *World) ColorAt(r ray.Ray) color.Color {
	hit := w.intersect(r).Hit()
	if hit.Empty() {
		return color.Black()
	}

	return w.shadeHit(ray.PrepareComputations(hit, r))
}

func (w *World) intersect(r ray.Ray) ray.Intersections {
	coll := ray.Intersections{}
	for _, o := range w.objects {
		coll = append(coll, r.Intersect(o)...)
	}
	coll.Sort()

	return coll
}

func (w *World) shadeHit(comps ray.Computations) color.Color {
	return ray.Lighting(
		comps.Object().Material(),
		w.Light(),
		comps.Point(),
		comps.EyeV(),
		comps.NormalV(),
	)
}

func Default() *World {
	o1 := objects.NewSphere()
	o1.SetMaterial(objects.NewMaterial(color.New(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0))
	o2 := objects.NewSphere()
	o2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))

	return &World{
		objects: []objects.Object{
			objects.Object(o1), objects.Object(o2),
		},
		light: ray.NewLight(tuple.NewPoint(-10, 10, -10), color.New(1, 1, 1)),
	}
}

func New(objects []objects.Object, light ray.Light) *World {
	return &World{objects, light}
}
