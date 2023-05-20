package scenes

import (
	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/light"
	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type Scene struct {
	Shapes []shapes.Shape
	Lights []light.Light
}

func Default() *Scene {
	o1 := shapes.NewSphere()
	o1.SetMaterial(materials.NewMaterial(color.New(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0, 0, 0, 1))
	o2 := shapes.NewSphere()
	o2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))

	return &Scene{
		Shapes: []shapes.Shape{
			shapes.Shape(o1), shapes.Shape(o2),
		},
		Lights: []light.Light{
			light.NewLight(tuple.NewPoint(-10, 10, -10), color.New(1, 1, 1)),
		},
	}
}

func New(shapes []shapes.Shape, lights []light.Light) *Scene {
	return &Scene{shapes, lights}
}
