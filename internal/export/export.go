// Exports the finished canvas into a PPM file.
package export

import (
	"bytes"
	"fmt"
	"math"

	"github.com/kaizencodes/glimpse/internal/canvas"
	"github.com/kaizencodes/glimpse/internal/color"
)

const (
	PpmMax    = 255
	PpmFormat = "P3"
)

func Export(c canvas.Canvas) string {
	var result bytes.Buffer
	result.WriteString(header(c))

	for y := 0; y < len(c[0]); y++ {
		for x := 0; x < len(c); x++ {
			result.WriteString(convertColor(c[x][y]))
		}
		result.WriteByte('\n')
	}
	result.WriteByte('\n')
	return result.String()
}

func header(c canvas.Canvas) string {
	return fmt.Sprintf("%s\n%d %d\n%d\n", PpmFormat, len(c), len((c)[0]), PpmMax)
}

func convertColor(c color.Color) string {
	r := rgbScale(c.R)
	g := rgbScale(c.G)
	b := rgbScale(c.B)

	return fmt.Sprintf("%d %d %d ", r, g, b)
}

// colors have a 0-1 range values, this method converts it to the ppm compatible 0-255 rgb range
func rgbScale(v float64) int {
	return int(math.Min(math.Max(0, v*PpmMax), PpmMax))
}
