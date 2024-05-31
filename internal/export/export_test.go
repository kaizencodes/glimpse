package export

import (
	"testing"

	"github.com/kaizencodes/glimpse/internal/canvas"
	"github.com/kaizencodes/glimpse/internal/color"
)

func TestExport(t *testing.T) {
	var tests = []struct {
		c        canvas.Canvas
		expected []byte
	}{
		{
			c: canvas.Canvas{
				[]color.Color{
					color.New(1.5, 0, 0), color.New(-0.5, 0, 1), color.New(0, 0, 0),
				},
				[]color.Color{
					color.New(0, 0.5, 0), color.New(0, 0, 0), color.New(0, 0, 0),
				},
			},
			expected: []byte(`P3
2 3
255
255 0 0 0 127 0 
0 0 255 0 0 0 
0 0 0 0 0 0 

`),
		},
	}

	for _, test := range tests {
		if got := Export(test.c); string(got) != string(test.expected) {
			t.Errorf("canvas \n%s \nexport to ppm \ngot: \n%s. \nexpected: \n%s", test.c, got, test.expected)
		}
	}
}
