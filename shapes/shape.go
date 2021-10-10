package shapes

import (
	"glimpse/matrix"
	"glimpse/tuple"
)

type Shape interface {
	Material() Material
	SetMaterial(m Material)
	Transform() matrix.Matrix
	SetTransform(transform matrix.Matrix)
	Normal(worldPoint tuple.Tuple) tuple.Tuple
}
