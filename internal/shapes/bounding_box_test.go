package shapes

import (
	"testing"

	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestAddPoint(t *testing.T) {
	// Adding points to an empty bounding box
	box := DefaultBoundingBox()
	box.AddPoint(tuple.NewPoint(-5, 2, 0))
	box.AddPoint(tuple.NewPoint(7, 2, -3))
	expected := NewBoundingBox(tuple.NewPoint(7, 2, 0), tuple.NewPoint(-5, 2, -3))

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}
