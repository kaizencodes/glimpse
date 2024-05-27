// camera is the viewport through which the scene is viewed.
package camera

import (
	"math"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type Camera struct {
	Width, Height int     // the with and height of the image in pixels.
	Fov           float64 // An angle that describes how much the camera can see.
	// when the value is small it the view will be zoomed in.
	pixelSize             float64       // the size of a pixel in world units.
	halfWidth, halfHeight float64       // helper variables to avoid repeated calculations.
	transform             matrix.Matrix // transformation matrix to position the camera.
}

func New(width, height int, fov float64) *Camera {
	c := Camera{
		Width:     width,
		Height:    height,
		Fov:       fov,
		transform: matrix.DefaultTransform(),
	}

	halfView := math.Tan(fov / 2.0)
	aspect := float64(width) / float64(height)
	if aspect >= 1 {
		c.halfWidth = halfView
		c.halfHeight = halfView / aspect
	} else {
		c.halfWidth = halfView * aspect
		c.halfHeight = halfView
	}

	c.pixelSize = (c.halfWidth * 2.0) / float64(c.Width)

	return &c
}

func (c *Camera) Transform() matrix.Matrix {
	return c.transform
}

func (c *Camera) SetTransform(m matrix.Matrix) {
	c.transform = m
}

// RayForPixel computes the ray that passes through the camera pixel (x, y).
// TODO: move to the renderer package
func (c *Camera) RayForPixel(x, y int) *ray.Ray {
	// the offset from the edge of the canvas to the pixel's center
	xOffset := (float64(x) + 0.5) * c.pixelSize
	yOffset := (float64(y) + 0.5) * c.pixelSize

	// the untransformed coordinates of the pixel in global space.
	// the camera looks toward -z, so +x is to the left.
	sceneX := c.halfWidth - xOffset
	sceneY := c.halfHeight - yOffset

	invTransform := c.transform.Inverse()

	// using the camera matrix, transform the canvas point and the origin,
	// and then compute the ray's direction vector. The canvas is at z=-1
	pixel := tuple.Multiply(invTransform, tuple.NewPoint(sceneX, sceneY, -1))
	origin := tuple.Multiply(invTransform, tuple.NewPoint(0, 0, 0))
	direction := tuple.Subtract(pixel, origin).Normalize()

	return ray.New(origin, direction)
}

// Transforms the camera position,
// from is a point in global space, where the camera is at,
// to is a vector in global space, where the camera is looking,
// up is a vector in global space, indicating which direction is up.
// TODO: make this private and call inside the new method to set the transform. Not used anywhere else.
func ViewTransformation(from, to, up tuple.Tuple) matrix.Matrix {
	forward := tuple.Subtract(to, from).Normalize()
	left := tuple.Cross(forward, up.Normalize())
	trueUp := tuple.Cross(left, forward)

	data := []float64{
		left.X, left.Y, left.Z, 0,
		trueUp.X, trueUp.Y, trueUp.Z, 0,
		-forward.X, -forward.Y, -forward.Z, 0,
		0, 0, 0, 1,
	}

	orientation := matrix.New(4, 4, data)

	result := matrix.Multiply(orientation, matrix.Translation(-from.X, -from.Y, -from.Z))

	return result
}
