package canvas

import (
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
