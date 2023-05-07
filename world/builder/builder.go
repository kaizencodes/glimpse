package builder

import (
	"glimpse/camera"
	"glimpse/color"
	"glimpse/light"
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
	}
	// shape.SetMaterial(buildMaterial(config.Material))
	transform := buildTransforms(config.Transform)
	shape.SetTransform(transform)

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
	}
	return transform
}

// func buildMaterial(config cfg.Material) *shapes.Material {
// 	material := shapes.NewMaterial(
// 		color.FromSlice(config.Color),
// 		config.Ambient,
// 		config.Diffuse,
// 		config.Specular,
// 		config.Shininess,
// 		config.Reflective,
// 		config.Transparency,
// 		config.RefractiveIndex,
// 	)
// 	return material
// }

// func buildWorld(config cfg.Config) world.World {
// 	world := *world.Default()
// 	// world.SetLights(buildLight(config.Light))
// 	return world
// }