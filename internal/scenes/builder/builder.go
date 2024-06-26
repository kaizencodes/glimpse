package builder

import (
	"fmt"
	"os"

	"github.com/kaizencodes/glimpse/internal/camera"
	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/light"
	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/projectpath"
	"github.com/kaizencodes/glimpse/internal/scenes"
	cfg "github.com/kaizencodes/glimpse/internal/scenes/config"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

func BuildScene(config cfg.Scene) (*camera.Camera, *scenes.Scene) {
	cam := buildCamera(config.Camera)
	scene := scenes.Default()
	scene.Lights = buildLights(config.Lights)
	scene.Shapes = buildObjects(config.Objects)

	return cam, scene
}

func buildCamera(config cfg.Camera) *camera.Camera {
	cam := camera.New(
		int(config.Width),
		int(config.Height),
		config.Fov,
	)
	cam.SetTransform(camera.ViewTransformation(
		tuple.NewVectorFromSlice(config.From),
		tuple.NewVectorFromSlice(config.To),
		tuple.NewVectorFromSlice(config.Up),
	))
	return cam
}

func buildLights(config []cfg.Light) []light.Light {
	var lights []light.Light
	for i := 0; i < len(config); i++ {
		l := light.NewLight(
			tuple.NewPointFromSlice(config[i].Position),
			color.FromSlice(config[i].Intensity),
		)
		lights = append(lights, l)
	}
	return lights
}

func buildObjects(config []cfg.Object) []shapes.Shape {
	var shapes []shapes.Shape
	for i := 0; i < len(config); i++ {
		shapes = append(shapes, buildObject(config[i]))
	}
	return shapes
}

func buildObject(config cfg.Object) shapes.Shape {
	var shape shapes.Shape

	switch config.Type {
	case "sphere":
		shape = shapes.NewSphere()

		shape.SetMaterial(buildMaterial(config.Material))
		shape.SetTransform(buildTransforms(config.Transform))
	case "plane":
		shape = shapes.NewPlane()

		shape.SetMaterial(buildMaterial(config.Material))
		shape.SetTransform(buildTransforms(config.Transform))
	case "cube":
		shape = shapes.NewCube()

		shape.SetMaterial(buildMaterial(config.Material))
		shape.SetTransform(buildTransforms(config.Transform))
	case "cylinder":
		cylinder := shapes.NewCylinder()

		cylinder.Minimum = config.Minimum
		cylinder.Maximum = config.Maximum
		cylinder.Closed = config.Closed

		shape = cylinder

		shape.SetMaterial(buildMaterial(config.Material))
		shape.SetTransform(buildTransforms(config.Transform))
	case "model":
		data, err := os.ReadFile(projectpath.Root + config.File)
		if err != nil {
			panic(fmt.Sprintf("Object file could not be read: %s\n%s", config.File, err.Error()))
		}
		model := shapes.NewModel(string(data))

		model.SetMaterial(buildMaterial(config.Material))
		model.SetTransform(buildTransforms(config.Transform))

		model.CalculateBoundingBox()

		shape = model
	case "group":
		group := shapes.NewGroup()
		group.AddChild(buildObjects(config.Children)...)
		group.SetTransform(buildTransforms(config.Transform))
		group.CalculateBoundingBoxCascade()

		group.Divide(10)
		shape = group
	default:
		panic("Unknown shape type")
	}

	return shape
}

func buildTransforms(config []cfg.Transform) matrix.Matrix {
	var transforms matrix.Matrix

	// If there are no transforms, return the identity matrix.
	if len(config) == 0 {
		return matrix.DefaultTransform()
	}

	// If there is only one transform, just build it and return.
	// Saves a bit of computation, since we are not multiplying by the identity matrix.
	transforms = buildTransform(config[len(config)-1])
	if len(config) == 1 {
		return transforms
	}

	// Multiply the transforms in reverse order, since the first one is loaded, we start from the second to last.
	for i := len(config) - 2; i >= 0; i-- {
		transforms = matrix.Multiply(transforms, buildTransform(config[i]))
	}

	return transforms
}

func buildTransform(config cfg.Transform) matrix.Matrix {
	var transform matrix.Matrix
	switch config.Type {
	case "scale":
		transform = matrix.Scaling(config.Values[0], config.Values[1], config.Values[2])
	case "translate":
		transform = matrix.Translation(config.Values[0], config.Values[1], config.Values[2])
	case "rotate-x":
		transform = matrix.RotationX(config.Values[0])
	case "rotate-y":
		transform = matrix.RotationY(config.Values[0])
	case "rotate-z":
		transform = matrix.RotationZ(config.Values[0])
	}
	return transform
}

func buildMaterial(config cfg.Material) *materials.Material {
	var col color.Color

	if len(config.Color) == 0 {
		col = color.White()
	} else {
		col = color.FromSlice(config.Color)
	}

	material := materials.NewMaterial(
		col,
		config.Ambient,
		config.Diffuse,
		config.Specular,
		config.Shininess,
		config.Reflective,
		config.Transparency,
		config.RefractiveIndex,
	)

	if config.Pattern.Type != "" {
		material.SetPattern(buildPattern(config.Pattern))
		material.SetTransform(buildTransforms(config.Pattern.Transform))
	}

	return material
}

func buildPattern(config cfg.Pattern) *materials.Pattern {
	var pattern *materials.Pattern

	switch config.Type {
	case "stripe":
		pattern = materials.NewPattern(
			materials.Stripe,
			color.FromSlice(config.Colors[0]),
			color.FromSlice(config.Colors[1]),
		)
	case "gradient":
		pattern = materials.NewPattern(
			materials.Gradient,
			color.FromSlice(config.Colors[0]),
			color.FromSlice(config.Colors[1]),
		)
	case "ring":
		pattern = materials.NewPattern(
			materials.Ring,
			color.FromSlice(config.Colors[0]),
			color.FromSlice(config.Colors[1]),
		)
	case "checker":
		pattern = materials.NewPattern(
			materials.Checker,
			color.FromSlice(config.Colors[0]),
			color.FromSlice(config.Colors[1]),
		)
	default:
		pattern = materials.NewPattern(
			materials.Base,
			color.FromSlice(config.Colors[0]),
		)
	}
	return pattern
}
