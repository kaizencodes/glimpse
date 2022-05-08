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
	a, _ := strconv.ParseFloat(split[1], 64)
	b, _ := strconv.ParseFloat(split[2], 64)
	c, _ := strconv.ParseFloat(split[3], 64)
	return a, b, c
}

func parseFaces(input string, points []tuple.Tuple) []*shapes.Triangle {
	r := regexp.MustCompile("(?m)^f.*\n")
	faces := r.FindAllString(input, -1)
	triangles := []*shapes.Triangle{}
	for _, line := range faces {
		a, b, c := splitFaceLine(line)
		triangle := shapes.NewTriangle(points[a], points[b], points[c])
		triangles = append(triangles, triangle)
	}
	return triangles
}

func splitFaceLine(line string) (int, int, int) {
	stripNewline := strings.Replace(line, "\n", "", 1)
	split := strings.Split(stripNewline, " ")
	a, _ := strconv.Atoi(split[1])
	b, _ := strconv.Atoi(split[2])
	c, _ := strconv.Atoi(split[3])
	return a, b, c
}
