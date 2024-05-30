package shapes

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestPlaneLocalNormalAt(t *testing.T) {
	// The normal of a plane is constant everywhere.
	var tests = []struct {
		plane    *Plane
		point    tuple.Tuple
		expected tuple.Tuple
	}{
		{
			plane:    NewPlane(),
			point:    tuple.NewPoint(0, 0, 0),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			plane:    NewPlane(),
			point:    tuple.NewPoint(10, 0, -10),
			expected: tuple.NewVector(0, 1, 0),
		},
		{
			plane:    NewPlane(),
			point:    tuple.NewPoint(-5, 0, 150),
			expected: tuple.NewVector(0, 1, 0),
		},
	}

	for _, test := range tests {
		if got := test.plane.localNormalAt(test.point, Intersection{}); !got.Equal(test.expected) {
			t.Errorf("Plane normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", test.plane, test.point, got, test.expected)
		}
	}
}

func TestBoundingBoxForPlane(t *testing.T) {
	//  A plane has a bounding box
	p := NewPlane()
	p.CalculateBoundingBox()
	box := p.BoundingBox()
	expected := NewBoundingBox(
		tuple.NewPoint(math.Inf(-1), 0, math.Inf(-1)),
		tuple.NewPoint(math.Inf(1), 0, math.Inf(1)),
	)

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}
