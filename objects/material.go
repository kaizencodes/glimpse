package objects

import "glimpse/color"

type Material struct {
	color                                 color.Color
	ambient, diffuse, specular, shininess float64
}

func DefaultMaterial() Material {
	return Material{
		color:     color.Color{1, 1, 1},
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}

func NewMaterial(col color.Color, ambient, diffuse, specular, shininess float64) Material {
	return Material{col, ambient, diffuse, specular, shininess}
}

func (mat Material) GetColor() color.Color {
	return mat.color
}

func (mat Material) GetAmbient() float64 {
	return mat.ambient
}

func (mat Material) GetDiffuse() float64 {
	return mat.diffuse
}

func (mat Material) GetSpecular() float64 {
	return mat.specular
}

func (mat Material) GetShininess() float64 {
	return mat.shininess
}
