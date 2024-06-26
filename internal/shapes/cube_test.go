package shapes

import (
	"testing"

	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestCubelocalNormalAt(t *testing.T) {
	var tests = []struct {
		shape    *Cube
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(1, 0.5, -0.8),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(-1, -0.2, 0.9),
			expected: tuple.NewVector(-1, 0, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(-0.4, 1, -0.1),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(0.3, -1, -0.7),
			expected: tuple.NewVector(0, -1, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(-0.6, 0.3, 1),
			expected: tuple.NewVector(0, 0, 1),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(0.4, 0.4, -1),
			expected: tuple.NewVector(0, 0, -1),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(1, 1, 1),
			expected: tuple.NewVector(1, 0, 0),
		},
		{
			shape:    NewCube(),
			point:    tuple.NewPoint(-1, -1, -1),
			expected: tuple.NewVector(-1, 0, 0),
		},
	}

	for _, test := range tests {
		if got := test.shape.localNormalAt(test.point, Intersection{}); !got.Equal(test.expected) {
			t.Errorf("Cube normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", test.shape, test.point, got, test.expected)
		}
	}
}

func TestBoundingBoxForCube(t *testing.T) {
	//  A cube has a bounding box
	c := NewCube()
	c.CalculateBoundingBox()
	box := c.BoundingBox()
	expected := NewBoundingBox(tuple.NewPoint(-1, -1, -1), tuple.NewPoint(1, 1, 1))

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}
