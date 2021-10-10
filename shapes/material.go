package shapes

import (
	"fmt"
	"glimpse/color"
)

type Material struct {
	color                                 color.Color
	ambient, diffuse, specular, shininess float64
}

func (mat Material) Color() color.Color {
	return mat.color
}

func (mat Material) Ambient() float64 {
	return mat.ambient
}

func (mat Material) Diffuse() float64 {
	return mat.diffuse
}

func (mat Material) Specular() float64 {
	return mat.specular
}

func (mat Material) Shininess() float64 {
	return mat.shininess
}

func (mat Material) String() string {
	return fmt.Sprintf("Material(color: %s\n, ambient: %f, diffuse: %f, specular: %f, shininess: %f,)",
		mat.color,
		mat.ambient,
		mat.diffuse,
		mat.specular,
		mat.shininess,
	)
}

func DefaultMaterial() Material {
	return Material{
		color:     color.New(1, 1, 1),
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}

func NewMaterial(col color.Color, ambient, diffuse, specular, shininess float64) Material {
	return Material{col, ambient, diffuse, specular, shininess}
}
