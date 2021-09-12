package ray

import (
	"fmt"
	"glimpse/color"
	"glimpse/objects"
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

func NewLight(position tuple.Tuple, intensity color.Color) Light {
	return Light{position, intensity}
}

func Lighting(mat objects.Material, light Light, point, eyeV, normalV tuple.Tuple) color.Color {
	effectiveColor := color.HadamardProduct(mat.Color(), light.intensity)
	lightV := tuple.Subtract(light.position, point).Normalize()
	ambient := effectiveColor.Scalar(mat.Ambient())
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
