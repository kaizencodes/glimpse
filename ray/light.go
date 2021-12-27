package ray

import (
	"fmt"
	"glimpse/color"
	"glimpse/shapes"
	"glimpse/tuple"
	"math"
)

type Light struct {
	position  tuple.Tuple
	intensity color.Color
}

func (l Light) String() string {
	return fmt.Sprintf("Light(position: %f, intensity: %f)", l.position, l.intensity)
}

func (l Light) Position() tuple.Tuple {
	return l.position
}

func (l Light) Intensity() color.Color {
	return l.intensity
}

func NewLight(position tuple.Tuple, intensity color.Color) Light {
	return Light{position, intensity}
}

func Lighting(shape shapes.Shape, light Light, point, eyeV, normalV tuple.Tuple, inShadow bool) color.Color {
	mat := shape.Material()
	coloring := shapes.ColorAt(point, shape)
	effectiveColor := color.HadamardProduct(coloring, light.intensity)
	lightV := tuple.Subtract(light.position, point).Normalize()
	ambient := effectiveColor.Scalar(mat.Ambient())
	if inShadow {
		return ambient
	}

	lightDotNormal := tuple.Dot(lightV, normalV)
	var diffuse, specular color.Color
	if lightDotNormal < 0 {
		diffuse = color.Black()
		specular = color.Black()
	} else {
		diffuse = effectiveColor.Scalar(mat.Diffuse() * lightDotNormal)
		reflectV := tuple.Reflect(lightV.Negate(), normalV)
		reflectDotEye := tuple.Dot(reflectV, eyeV)
		if reflectDotEye <= 0 {
			specular = color.Black()
		} else {
			factor := math.Pow(reflectDotEye, mat.Shininess())
			specular = light.intensity.Scalar(mat.Specular() * factor)
		}
	}

	return color.Add(color.Add(ambient, diffuse), specular)
}
