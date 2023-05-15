package parser

import (
	"testing"

	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

func TestParseVertices(t *testing.T) {
	input := `v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0
vt 1 0
vt 1 0`

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

func TestParseNormals(t *testing.T) {
	input := `v 1 0 0
v 1 1 0
vn 0 0 1
vn 0.707 0 -0.707
vn 1 2 3`

	result := parseNormals(input)
	expected := []tuple.Tuple{
		tuple.NewVector(0, 0, 0), // index is 1 based
		tuple.NewVector(0, 0, 1),
		tuple.NewVector(0.707, 0, -0.707),
		tuple.NewVector(1, 2, 3),
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
	vertices := parseFaces(input, points, []tuple.Tuple{})

	assertFace(vertices[0], points[1], points[2], points[3], t)
	assertFace(vertices[1], points[1], points[3], points[4], t)
}

func TestParseFacesWithNormals(t *testing.T) {
	input := `v 0 1 0
v -1 0 0
v 1 0 0
vn -1 0 0
vn 1 0 0
vn 0 1 0
f 1//3 2//1 3//2
f 1/0/3 2/102/1 3/14/2
`

	vertices := []tuple.Tuple{
		tuple.NewPoint(0, 0, 0), // index is 1 based
		tuple.NewPoint(0, 1, 0),
		tuple.NewPoint(-1, 0, 0),
		tuple.NewPoint(1, 0, 0),
	}
	normals := []tuple.Tuple{
		tuple.NewVector(0, 0, 0), // index is 1 based
		tuple.NewVector(-1, 0, 0),
		tuple.NewVector(1, 0, 0),
		tuple.NewVector(0, 1, 0),
	}
	faces := parseFaces(input, vertices, normals)

	assertFace(faces[0], vertices[1], vertices[2], vertices[3], t)
	assertFace(faces[1], vertices[1], vertices[2], vertices[3], t)
	assertFaceNormal(faces[0], vertices[3], vertices[1], vertices[2], t)
	assertFaceNormal(faces[1], vertices[3], vertices[1], vertices[2], t)
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
	vertices := parseFaces(input, points, []tuple.Tuple{})
	assertFace(vertices[0], points[1], points[2], points[3], t)
	assertFace(vertices[1], points[1], points[3], points[4], t)
	assertFace(vertices[2], points[1], points[4], points[5], t)
}

func assertFace(face *shapes.Triangle, p1, p2, p3 tuple.Tuple, t *testing.T) {
	if !face.P1().Equal(p1) {
		t.Errorf("Incorrect parsing. expected vertex point P1 to be \n%s \n got %s", p1, face.P1())
	}
	if !face.P2().Equal(p2) {
		t.Errorf("Incorrect parsing. expected vertex point P2 to be \n%s \n got %s", p2, face.P2())
	}
	if !face.P3().Equal(p3) {
		t.Errorf("Incorrect parsing. expected vertex point P3 to be \n%s \n got %s", p3, face.P3())
	}
}

func assertFaceNormal(face *shapes.Triangle, n1, n2, n3 tuple.Tuple, t *testing.T) {
	if !face.N1().Equal(n1) {
		t.Errorf("Incorrect parsing. expected vertex normal N1 to be \n%s \n got %s", n1, face.N1())
	}
	if !face.N2().Equal(n2) {
		t.Errorf("Incorrect parsing. expected vertex normal N2 to be \n%s \n got %s", n2, face.N2())
	}
	if !face.N3().Equal(n3) {
		t.Errorf("Incorrect parsing. expected vertex normal N3 to be \n%s \n got %s", n3, face.N3())
	}
}
