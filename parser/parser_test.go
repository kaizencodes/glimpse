package parser

import (
	"glimpse/tuple"
	"testing"
)

func TestParseVertices(t *testing.T) {
	input := `v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0`

	result := parseVertices(input)
	expected := []tuple.Tuple{
		tuple.NewPoint(0, 0, 0), // index is 1 based
		tuple.NewPoint(-1, 1, 0),
		tuple.NewPoint(-1, 0.5, 0),
		tuple.NewPoint(1, 0, 0),
		tuple.NewPoint(1, 1, 0),
	}
	for k, v := range result {
		if v != expected[k] {
			t.Errorf("Incorrect parsing. expected \n%s \n got %s", expected[k], v)
		}
	}

}

func TestParseFaces(t *testing.T) {
	input := `v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0
f 1 2 3
f 1 3 4
`

	points := []tuple.Tuple{
		tuple.NewPoint(0, 0, 0), // index is 1 based
		tuple.NewPoint(-1, 1, 0),
		tuple.NewPoint(-1, 0, 0),
		tuple.NewPoint(1, 0, 0),
		tuple.NewPoint(1, 1, 0),
	}
	vertices := parseFaces(input, points)
	if !vertices[0].A().Equal(points[1]) {
		t.Errorf("Incorrect parsing. expected vertice point A to be \n%s \n got %s", points[1], vertices[0].A())
	}
	if !vertices[0].B().Equal(points[2]) {
		t.Errorf("Incorrect parsing. expected vertice point B to be \n%s \n got %s", points[2], vertices[0].B())
	}
	if !vertices[0].C().Equal(points[3]) {
		t.Errorf("Incorrect parsing. expected vertice point C to be \n%s \n got %s", points[3], vertices[0].C())
	}
	if !vertices[1].A().Equal(points[1]) {
		t.Errorf("Incorrect parsing. expected vertice point A to be \n%s \n got %s", points[1], vertices[1].A())
	}
	if !vertices[1].B().Equal(points[3]) {
		t.Errorf("Incorrect parsing. expected vertice point B to be \n%s \n got %s", points[3], vertices[1].B())
	}
	if !vertices[1].C().Equal(points[4]) {
		t.Errorf("Incorrect parsing. expected vertice point C to be \n%s \n got %s", points[4], vertices[1].C())
	}
}
