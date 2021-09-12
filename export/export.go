package export

import (
	"fmt"
	"glimpse/canvas"
	"glimpse/color"
	"math"
)

const (
	PpmMax    = 255
	PpmFormat = "P3"
)

func Export(c canvas.Canvas) string {
	result := header(c)
	for _, row := range c {
		for _, val := range row {
			result += convertColor(val)
		}
		result += string('\n')
	}
	result += string('\n')
	return result
}

func header(c canvas.Canvas) string {
	return fmt.Sprintf("%s\n%d %d\n%d\n", PpmFormat, len(c), len((c)[0]), PpmMax)
}

func convertColor(c color.Color) string {
	r := rgbScale(c.R())
	g := rgbScale(c.G())
	b := rgbScale(c.B())

	return fmt.Sprintf("%d %d %d ", r, g, b)
}

func rgbScale(v float64) int {
	return int(math.Min(math.Max(0, v*PpmMax), PpmMax))
}
