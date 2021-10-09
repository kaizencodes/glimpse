package objects

import (
	"glimpse/matrix"
	"glimpse/tuple"
)

type Object interface {
	Material() Material
	SetMaterial(m Material)
	Transform() matrix.Matrix
	SetTransform(transform matrix.Matrix)
	Normal(worldPoint tuple.Tuple) tuple.Tuple
}
