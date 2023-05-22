package canvas

import (
	"github.com/kaizencodes/glimpse/internal/color"
)

type Canvas [][]color.Color

func New(w, h int) Canvas {
	c := make(Canvas, w)
	for i := 0; i < int(w); i++ {
		c[i] = make([]color.Color, h)
	}
	return c
}

func (c Canvas) String() string {
	var result string

	for y := 0; y < len(c[0]); y++ {
		for x := 0; x < len(c); x++ {
			result += c[x][y].String()
		}
		result += string('\n')
	}
	return result
}
