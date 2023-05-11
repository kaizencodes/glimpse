package builder

import (
	"glimpse/camera"
	"glimpse/color"
	"glimpse/light"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/shapes"
	"glimpse/tuple"
	"glimpse/world"
	cfg "glimpse/world/config"
)

func BuildScene(scene cfg.Scene) (*camera.Camera, *world.World) {
	cam := buildCamera(scene.Camera)
	world := world.Default()
	world.SetLights(buildLight(scene.Light))
	world.SetShapes(buildObjects(scene.Objects))

	return cam, world
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

func buildLight(config cfg.Light) []light.Light {
	return []light.Light{
		light.NewLight(
			tuple.NewPointFromSlice(config.Position),
			color.FromSlice(config.Intensity),
		),
	}
}

func buildObjects(config []cfg.Object) []shapes.Shape {
	var shapes []shapes.Shape
	for _, obj := range config {
		shapes = append(shapes, buildObject(obj))
	}
	return shapes
}

func buildObject(config cfg.Object) shapes.Shape {
	var shape shapes.Shape

	switch config.Type {
	case "sphere":
		shape = shapes.NewSphere()
	case "plane":
		shape = shapes.NewPlane()
	case "cube":
		shape = shapes.NewCube()
	case "cylinder":
		cylinder := shapes.NewCylinder()

		cylinder.SetMinimum(config.Minimum)
		cylinder.SetMaximum(config.Maximum)
		cylinder.SetClosed(config.Closed)

		shape = cylinder
	default:
		panic("Unknown shape type")
	}

	shape.SetMaterial(buildMaterial(config.Material))
	shape.SetTransform(buildTransforms(config.Transform))

	return shape
}

func buildTransforms(config []cfg.Transform) matrix.Matrix {
	var transforms matrix.Matrix

	transforms = buildTransform(config[0])
	if len(config) == 1 {
		return transforms
	}

	for i := 1; i < len(config); i++ {
		transforms, _ = matrix.Multiply(buildTransform(config[i]), transforms)
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
	material := materials.NewMaterial(
		color.FromSlice(config.Color),
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
	}

	return material
}

func buildPattern(config cfg.Pattern) *materials.Pattern {
	var pattern *materials.Pattern

	switch config.Type {
	case "stripe":
		materials.NewPattern(
			materials.Stripe,
			color.FromSlice(config.Colors[0]),
			color.FromSlice(config.Colors[1]),
		)
	case "gradient":
		materials.NewPattern(
			materials.Gradient,
			color.FromSlice(config.Colors[0]),
			color.FromSlice(config.Colors[1]),
		)
	case "ring":
		materials.NewPattern(
			materials.Ring,
			color.FromSlice(config.Colors[0]),
			color.FromSlice(config.Colors[1]),
		)
	case "checkers":
		materials.NewPattern(
			materials.Checker,
			color.FromSlice(config.Colors[0]),
			color.FromSlice(config.Colors[1]),
		)
	default:
		materials.NewPattern(
			materials.Base,
			color.FromSlice(config.Colors[0]),
		)
	}
	return pattern
}
