package parser

import (
	"glimpse/shapes"
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

	assertFace(vertices[0], points[1], points[2], points[3], t)
	assertFace(vertices[1], points[1], points[3], points[4], t)
}

func TestTriangulatingPolygons(t *testing.T) {
	input := `v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0
v 0 2 0
f 1 2 3 4 5
`

	points := []tuple.Tuple{
		tuple.NewPoint(0, 0, 0), // index is 1 based
		tuple.NewPoint(-1, 1, 0),
		tuple.NewPoint(-1, 0, 0),
		tuple.NewPoint(1, 0, 0),
		tuple.NewPoint(1, 1, 0),
		tuple.NewPoint(0, 2, 0),
	}
	vertices := parseFaces(input, points)
	assertFace(vertices[0], points[1], points[2], points[3], t)
	assertFace(vertices[1], points[1], points[3], points[4], t)
	assertFace(vertices[2], points[1], points[4], points[5], t)
}

func assertFace(face *shapes.Triangle, p1, p2, p3 tuple.Tuple, t *testing.T) {
	if !face.P1().Equal(p1) {
		t.Errorf("Incorrect parsing. expected vertice point P1 to be \n%s \n got %s", p1, face.P1())
	}
	if !face.P2().Equal(p2) {
		t.Errorf("Incorrect parsing. expected vertice point P2 to be \n%s \n got %s", p2, face.P2())
	}
	if !face.P3().Equal(p3) {
		t.Errorf("Incorrect parsing. expected vertice point P3 to be \n%s \n got %s", p3, face.P3())
	}
}
