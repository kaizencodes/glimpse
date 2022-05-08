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
	points := parseVertices(input)

	triangles := parseFaces(input, points)
	for _, triangle := range triangles {
		group.AddChild(triangle)
	}
	return group
}

func parseVertices(input string) []tuple.Tuple {
	r := regexp.MustCompile("(?m)^v.*\n")
	vertices := r.FindAllString(input, -1)
	points := []tuple.Tuple{
		tuple.NewPoint(0, 0, 0), // index is 1 based
	}
	for _, line := range vertices {
		points = append(points, tuple.NewPoint(splitVerticeLine(line)))
	}
	return points
}

func splitVerticeLine(line string) (float64, float64, float64) {
	stripNewline := strings.Replace(line, "\n", "", 1)
	split := strings.Split(stripNewline, " ")
	p1, _ := strconv.ParseFloat(split[1], 64)
	p2, _ := strconv.ParseFloat(split[2], 64)
	p3, _ := strconv.ParseFloat(split[3], 64)
	return p1, p2, p3
}

func parseFaces(input string, points []tuple.Tuple) []*shapes.Triangle {
	r := regexp.MustCompile("(?m)^f.*\n")
	faces := r.FindAllString(input, -1)
	triangles := []*shapes.Triangle{}
	for _, line := range faces {
		indexes := convertLinesToIndexes(line)
		// fan triangulation
		for i := 0; i < len(indexes)-2; i++ {
			triangle := shapes.NewTriangle(points[indexes[0]], points[indexes[i+1]], points[indexes[i+2]])
			triangles = append(triangles, triangle)
		}
	}
	return triangles
}

func convertLinesToIndexes(line string) (indexes []int) {
	stripNewline := strings.Replace(line, "\n", "", 1)
	split := strings.Split(stripNewline, " ")

	// skipping the first element as it's the type char e.g. f
	for i := 1; i < len(split); i++ {
		n, _ := strconv.Atoi(split[i])
		indexes = append(indexes, n)
	}
	return indexes
}
