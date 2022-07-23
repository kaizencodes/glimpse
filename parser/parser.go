package parser

import (
	"glimpse/shapes"
	"glimpse/tuple"
	"regexp"
	"strconv"
	"strings"
)

func Parse(input string) *shapes.Group {
	group := shapes.NewGroup()
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
	vertexLines := r.FindAllString(input, -1)
	normals := []tuple.Tuple{
		tuple.NewVector(0, 0, 0), // index is 1 based
	}
	for _, line := range vertexLines {
		normals = append(normals, tuple.NewVector(splitVertexLine(line)))
	}
	return normals
}

func parseFaces(input string, vertices, normals []tuple.Tuple) (faces []*shapes.Triangle) {
	r := regexp.MustCompile("(?m)^f.*\n")
	faceLines := r.FindAllString(input, -1)
	// faces := []*shapes.Triangle{}
	for _, line := range faceLines {
		indexes := convertLinesToIndexes(line)
		// fan triangulation
		for i := 0; i < len(indexes)-2; i++ {
			// var face *shapes.Triangle
			if indexes[0][1] != 0 {
				face := shapes.NewSmoothTriangle(vertices[indexes[0][0]], vertices[indexes[i+1][0]], vertices[indexes[i+2][0]], vertices[indexes[0][1]], vertices[indexes[i+1][1]], vertices[indexes[i+2][1]])
				faces = append(faces, face)
			} else {
				face := shapes.NewTriangle(vertices[indexes[0][0]], vertices[indexes[i+1][0]], vertices[indexes[i+2][0]])
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
