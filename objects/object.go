package objects

import "glimpse/matrix"

type Object interface {
	Material() Material
	SetMaterial(m Material)
	Transform() matrix.Matrix
	SetTransform(transform matrix.Matrix)
}
