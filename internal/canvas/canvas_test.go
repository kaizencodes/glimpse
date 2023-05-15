package canvas

import (
	"testing"

	"github.com/kaizencodes/glimpse/internal/color"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		w    int
		h    int
		want Canvas
	}{
		{
			w: 2,
			h: 3,
			want: Canvas{
				[]color.Color{
					color.Black(), color.Black(), color.Black(),
				},
				[]color.Color{
					color.Black(), color.Black(), color.Black(),
				},
			},
		},
	}

	for _, test := range tests {
		if got := New(test.w, test.h); got.String() != test.want.String() {
			t.Errorf("canvas width w:%d, h:%d \ngot: \n%s. \nexpected: \n%s", test.w, test.h, got, test.want)
		}
	}
}
