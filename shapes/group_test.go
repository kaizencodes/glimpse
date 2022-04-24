package shapes

import (
	"testing"
)

func TestAddChild(t *testing.T) {
	g := NewGroup()
	s := NewTestShape()

	g.AddChild(s)

	if g.Children()[0] != s {
		t.Errorf("group did not add child.")
	}
}
