package canvas

import (
	"fmt"
	"glimpse/color"
)

type Canvas struct {
	pane [][]color.Color
}

const (
	PpmFormat = "P3"
	PpmMax    = 255
)

func NewCanvas(w, h int) *Canvas {
	c := &Canvas{}
	c.pane = make([][]color.Color, h)
	for i := 0; i < int(h); i++ {
		c.pane[i] = make([]color.Color, w)
	}
	return c
}

func (c *Canvas) WritePixel(w, h int, pixel color.Color) {
	c.pane[h][w] = pixel
}

func (c *Canvas) ExportToPpm() string {
	result := fmt.Sprintf("%s\n%d %d\n%d\n", PpmFormat, len(c.pane), len(c.pane[0]), PpmMax)
	for _, row := range c.pane {
		for _, val := range row {
			result += val.ConvertToPpm()
		}
		result += string('\n')
	}
	result += string('\n')
	return result
}

func (c *Canvas) String() string {
	result := ""
	for _, row := range c.pane {
		for _, val := range row {
			result += val.String()
		}
		result += string('\n')
	}
	return result
}
