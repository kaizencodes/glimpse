package canvas

import (
	"fmt"
	"glimpse/color"
)

type Canvas [][]color.Color

const (
	PpmFormat = "P3"
	PpmMax    = 255
)

func New(w, h int) Canvas {
	c := make(Canvas, h)
	for i := 0; i < int(h); i++ {
		c[i] = make([]color.Color, w)
	}
	return c
}

func (c Canvas) ExportToPpm() string {
	result := fmt.Sprintf("%s\n%d %d\n%d\n", PpmFormat, len(c), len((c)[0]), PpmMax)
	for _, row := range c {
		for _, val := range row {
			result += val.ConvertToPpm()
		}
		result += string('\n')
	}
	result += string('\n')
	return result
}

func (c Canvas) String() string {
	var result string

	for _, row := range c {
		for _, val := range row {
			result += val.String()
		}
		result += string('\n')
	}
	return result
}
