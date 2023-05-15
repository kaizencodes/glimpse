package camera

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/calc"
	"github.com/kaizencodes/glimpse/color"
	"github.com/kaizencodes/glimpse/matrix"
	"github.com/kaizencodes/glimpse/ray"
	"github.com/kaizencodes/glimpse/tuple"
	"github.com/kaizencodes/glimpse/world"
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
		if result := New(test.width, test.height, test.fov).PixelSize(); !calc.FloatEquals(result, test.expected) {
			t.Errorf("camera PixelSize expected %f, got %f", test.expected, result)
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
			c:        New(201, 101, math.Pi/2),
			x:        100,
			y:        50,
			expected: ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, -1)),
		},
		{
			c:        New(201, 101, math.Pi/2),
			x:        0,
			y:        0,
			expected: ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0.6651864261194509, 0.33259321305972545, -0.6685123582500481)),
		},
	}

	for _, test := range tests {
		if result := test.c.RayForPixel(test.x, test.y); !result.Equal(test.expected) {
			t.Errorf("RayForPixel expected %s, got %s", test.expected, result)
		}
	}

	c := New(201, 101, math.Pi/2)
	transform, _ := matrix.Multiply(matrix.RotationY(math.Pi/4), matrix.Translation(0, -2, 5))
	c.SetTransform(transform)
	expected := ray.NewRay(tuple.NewPoint(0, 2, -5), tuple.NewVector(0.7071067811865474, 0, -0.7071067811865478))

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
			from:     tuple.NewPoint(0, 0, 0),
			to:       tuple.NewPoint(0, 0, -1),
			up:       tuple.NewVector(0, 1, 0),
			expected: matrix.NewIdentity(4),
		},
		{
			from:     tuple.NewPoint(0, 0, 0),
			to:       tuple.NewPoint(0, 0, 1),
			up:       tuple.NewVector(0, 1, 0),
			expected: matrix.Scaling(-1, 1, -1),
		},
		{
			from:     tuple.NewPoint(0, 0, 8),
			to:       tuple.NewPoint(0, 0, 0),
			up:       tuple.NewVector(0, 1, 0),
			expected: matrix.Translation(0, 0, -8),
		},
		{
			from: tuple.NewPoint(1, 3, 2),
			to:   tuple.NewPoint(4, -2, 8),
			up:   tuple.NewVector(1, 1, 0),
			expected: matrix.Matrix{
				[]float64{-0.5070925528371099, 0.5070925528371099, 0.6761234037828132, -2.366431913239846},
				[]float64{0.7677159338596801, 0.6060915267313263, 0.12121830534626524, -2.8284271247461894},
				[]float64{-0.35856858280031806, 0.5976143046671968, -0.7171371656006361, 0},
				[]float64{0, 0, 0, 1},
			},
		},
	}

	for _, test := range tests {
		result := ViewTransformation(test.from, test.to, test.up)
		if !result.Equal(test.expected) {
			t.Errorf("ViewTransformation,\nto:\n%s\nfrom:\n%s\nup:\n%s\nresult:\n%s\nexpected: \n%s", test.to, test.from, test.up, result, test.expected)
		}
	}
}

func TestRender(t *testing.T) {
	w := world.Default()
	c := New(11, 11, math.Pi/2)
	transform := ViewTransformation(
		tuple.NewPoint(0, 0, -5),
		tuple.NewPoint(0, 0, 0),
		tuple.NewVector(0, 1, 0),
	)
	c.SetTransform(transform)
	img := c.Render(w)
	result := img[5][5]
	expected := color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
	if !result.Equal(expected) {
		t.Errorf("Render, expected %s, got %s", expected, result)
	}
}
