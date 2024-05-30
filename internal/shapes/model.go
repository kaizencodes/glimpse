package shapes

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

// Model is a shape that is defined by vertices.
// It is a group of triangles primitives.
type Model struct {
	group  Group
	parent Shape
}

func NewModel(model string) *Model {
	return &Model{
		group: *Parse(model),
	}
}

func (m *Model) String() string {
	return fmt.Sprintf("Model(material: %s, transform: %s)", m.group.material, m.group.transform)
}

func (m *Model) SetTransform(transform matrix.Matrix) {
	m.group.transform = transform
}

func (s *Model) SetMaterial(mat *materials.Material) {
	s.group.material = mat
}

func (m *Model) Material() *materials.Material {
	return m.group.material
}

func (m *Model) Transform() matrix.Matrix {
	return m.group.transform
}

func (m *Model) CalculateBoundingBox() {
	m.group.CalculateBoundingBoxCascade()
}

func (m *Model) BoundingBox() *BoundingBox {
	return m.group.BoundingBox()
}

func (m *Model) Divide(threshold int) {
	m.group.Divide(threshold)
}

func (m *Model) localNormalAt(_point tuple.Tuple, _hit Intersection) tuple.Tuple {
	return tuple.Tuple{}
}

func (m *Model) localIntersect(r *ray.Ray) Intersections {
	return m.group.localIntersect(r)
}

func (m *Model) Parent() Shape {
	return m.parent
}

func (m *Model) SetParent(other Shape) {
	m.parent = other
}

func Parse(input string) *Group {
	group := NewGroup()
	vertices := parseVertices(input)
	normals := parseNormals(input)

	faces := parseFaces(input, vertices, normals)
	for _, face := range faces {
		group.AddChild(face)
	}
	return group
}

func parseVertices(input string) []tuple.Tuple {
	r := regexp.MustCompile("(?m)^v .*\n")
	vertexLines := r.FindAllString(input, -1)
	vertices := []tuple.Tuple{
		tuple.NewPoint(0, 0, 0), // index is 1 based
	}
	for _, line := range vertexLines {
		vertices = append(vertices, tuple.NewPoint(splitVertexLine(line)))
	}
	return vertices
}

func splitVertexLine(line string) (float64, float64, float64) {
	stripNewline := strings.Replace(line, "\n", "", 1)
	split := strings.Split(stripNewline, " ")
	p1, _ := strconv.ParseFloat(split[1], 64)
	p2, _ := strconv.ParseFloat(split[2], 64)
	p3, _ := strconv.ParseFloat(split[3], 64)
	return p1, p2, p3
}

func parseNormals(input string) []tuple.Tuple {
	r := regexp.MustCompile("(?m)^vn.*\n")
	normalLines := r.FindAllString(input, -1)
	normals := []tuple.Tuple{
		tuple.NewVector(0, 0, 0), // index is 1 based
	}
	for _, line := range normalLines {
		normals = append(normals, tuple.NewVector(splitVertexLine(line)))
	}
	return normals
}

func parseFaces(input string, vertices, normals []tuple.Tuple) (faces []*Triangle) {
	r := regexp.MustCompile("(?m)^f.*\n")
	faceLines := r.FindAllString(input, -1)
	for _, line := range faceLines {
		indexes := convertLinesToIndexes(line)

		// fan triangulation
		for i := 0; i < len(indexes)-2; i++ {
			// var face *shapes.Triangle
			if indexes[0][1] != 0 {
				face := NewSmoothTriangle(vertices[indexes[0][0]], vertices[indexes[i+1][0]], vertices[indexes[i+2][0]], normals[indexes[0][1]], normals[indexes[i+1][1]], normals[indexes[i+2][1]])
				faces = append(faces, face)
			} else {
				face := NewTriangle(vertices[indexes[0][0]], vertices[indexes[i+1][0]], vertices[indexes[i+2][0]])
				faces = append(faces, face)
			}
		}
	}
	return faces
}

func convertLinesToIndexes(line string) (indexes [][]int) {
	stripNewline := strings.Replace(line, "\n", "", 1)
	split := strings.Split(stripNewline, " ")
	withNormal, _ := regexp.MatchString(`\d+\/\d*\/\d+`, split[1])
	// skipping the first element as it's the type char e.g. f
	for i := 1; i < len(split); i++ {
		if withNormal {
			n := strings.Split(split[i], "/")
			vertex, _ := strconv.Atoi(n[0])
			normal, _ := strconv.Atoi(n[2])
			indexes = append(indexes, []int{vertex, normal})
		} else {
			vertex, _ := strconv.Atoi(split[i])
			indexes = append(indexes, []int{vertex, 0})
		}
	}

	return indexes
}
