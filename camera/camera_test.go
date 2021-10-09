package camera

import (
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/ray"
	"glimpse/tuple"
	"glimpse/world"
	"math"
	"testing"
)

func TestPixelSize(t *testing.T) {
	var tests = []struct {
		height   int
		width    int
		fov      float64
		expected float64
	}{
		{
			height:   200,
			width:    125,
			fov:      math.Pi / 2,
			expected: 0.01,
		},
		{
			height:   125,
			width:    200,
			fov:      math.Pi / 2,
			expected: 0.01,
		},
	}

	for _, test := range tests {
		if result := New(test.width, test.height, test.fov).PixelSize(); result != test.expected {
			t.Errorf("camera PixelSize expected %f, got %f", test.expected, result)
		}
	}
}

func TestRayForPixel(t *testing.T) {
	var tests = []struct {
		c        *Camera
		x, y     int
		expected ray.Ray
	}{
		{
			c:        New(201, 101, math.Pi/2),
			x:        100,
			y:        50,
			expected: ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, -1)),
		},
		{
			c:        New(201, 101, math.Pi/2),
			x:        0,
			y:        0,
			expected: ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0.6651864261194509, 0.33259321305972545, -0.6685123582500481)),
		},
	}

	for _, test := range tests {
		if result := test.c.RayForPixel(test.x, test.y); result != test.expected {
			t.Errorf("RayForPixel expected %s, got %s", test.expected, result)
		}
	}

	c := New(201, 101, math.Pi/2)
	transform, _ := matrix.Multiply(matrix.RotationY(math.Pi/4), matrix.Translation(0, -2, 5))
	c.SetTransform(transform)
	expected := ray.New(tuple.NewPoint(0, 2, -5), tuple.NewVector(0.7071067811865474, 0, -0.7071067811865478))

	if result := c.RayForPixel(100, 50); result != expected {
		t.Errorf("RayForPixel expected %s, got %s", expected, result)
	}
}

func TestRender(t *testing.T) {
	w := world.Default()
	c := New(11, 11, math.Pi/2)
	transform := world.ViewTransformation(
		tuple.NewPoint(0, 0, -5),
		tuple.NewPoint(0, 0, 0),
		tuple.NewVector(0, 1, 0),
	)
	c.SetTransform(transform)
	img := c.Render(w)
	result := img[5][5]
	expected := color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
	if result != expected {
		t.Errorf("Render, expected %s, got %s", expected, result)
	}
}
