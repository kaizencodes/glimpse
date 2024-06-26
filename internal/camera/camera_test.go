package camera

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestPixelSize(t *testing.T) {
	var tests = []struct {
		height   int
		width    int
		fov      float64
		expected float64
	}{
		{
			// The pixel size for a horizontal canvas
			height:   200,
			width:    125,
			fov:      math.Pi / 2,
			expected: 0.01,
		},
		{
			// The pixel size for a vertical canvas
			height:   125,
			width:    200,
			fov:      math.Pi / 2,
			expected: 0.01,
		},
	}

	for _, test := range tests {
		if result := New(test.width, test.height, test.fov).pixelSize; !utils.FloatEquals(result, test.expected) {
			t.Errorf("camera pixelSize expected %f, got %f", test.expected, result)
		}
	}
}

func TestRayForPixel(t *testing.T) {
	var tests = []struct {
		c        *Camera
		x, y     int
		expected *ray.Ray
	}{
		{
			// Constructing a ray through the center of the canvas
			c:        New(201, 101, math.Pi/2),
			x:        100,
			y:        50,
			expected: ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, -1)),
		},
		{
			// Constructing a ray through a corner of the canvas
			c:        New(201, 101, math.Pi/2),
			x:        0,
			y:        0,
			expected: ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0.6651864261194509, 0.33259321305972545, -0.6685123582500481)),
		},
	}

	for _, test := range tests {
		if result := test.c.RayForPixel(test.x, test.y); !result.Equal(test.expected) {
			t.Errorf("RayForPixel expected %s, got %s", test.expected, result)
		}
	}

	// Constructing a ray when the camera is transformed
	c := New(201, 101, math.Pi/2)
	transform := matrix.Multiply(matrix.RotationY(math.Pi/4), matrix.Translation(0, -2, 5))
	c.SetTransform(transform)
	// (√2/2, 0, -√2/2)
	expected := ray.New(tuple.NewPoint(0, 2, -5), tuple.NewVector(0.7071067811865474, 0, -0.7071067811865478))

	if result := c.RayForPixel(100, 50); !result.Equal(expected) {
		t.Errorf("RayForPixel expected %s, got %s", expected, result)
	}
}

func TestViewTransformation(t *testing.T) {
	var tests = []struct {
		from     tuple.Tuple
		to       tuple.Tuple
		up       tuple.Tuple
		expected matrix.Matrix
	}{
		{
			// The transformation matrix for the default orientation
			from:     tuple.NewPoint(0, 0, 0),
			to:       tuple.NewPoint(0, 0, -1),
			up:       tuple.NewVector(0, 1, 0),
			expected: matrix.DefaultTransform(),
		},
		{
			// A view transformation matrix looking in positive z direction
			from:     tuple.NewPoint(0, 0, 0),
			to:       tuple.NewPoint(0, 0, 1),
			up:       tuple.NewVector(0, 1, 0),
			expected: matrix.Scaling(-1, 1, -1),
		},
		{
			// The view transformation moves the world
			from:     tuple.NewPoint(0, 0, 8),
			to:       tuple.NewPoint(0, 0, 0),
			up:       tuple.NewVector(0, 1, 0),
			expected: matrix.Translation(0, 0, -8),
		},
		{
			// An arbitrary view transformation
			from: tuple.NewPoint(1, 3, 2),
			to:   tuple.NewPoint(4, -2, 8),
			up:   tuple.NewVector(1, 1, 0),
			expected: matrix.New(4, 4,
				[16]float64{
					-0.5070925528371099, 0.5070925528371099, 0.6761234037828131, -2.366431913239846,
					0.76771593385968, 0.6060915267313263, 0.12121830534626525, -2.8284271247461894,
					-0.35856858280031806, 0.5976143046671968, -0.7171371656006361, 0,
					0, 0, 0, 1,
				},
			),
		},
	}

	for _, test := range tests {
		result := ViewTransformation(test.from, test.to, test.up)
		if !result.Equal(test.expected) {
			t.Errorf("ViewTransformation,\nto:\n%s\nfrom:\n%s\nup:\n%s\nresult:\n%s\nexpected: \n%s", test.to, test.from, test.up, result, test.expected)
		}
	}
}
