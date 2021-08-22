package canvas

import (
    "glimpse/color"
    "testing"
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
                    color.Color{}, color.Color{},
                },
                []color.Color{
                    color.Color{}, color.Color{},
                },
                []color.Color{
                    color.Color{}, color.Color{},
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

func TestExportToPpm(t *testing.T) {
    var tests = []struct {
        c    Canvas
        want string
    }{
        {
            c: Canvas{
                []color.Color{
                    color.Color{1.5, 0, 0}, color.Color{0, 0.5, 0},
                },
                []color.Color{
                    color.Color{-0.5, 0, 1}, color.Color{},
                },
                []color.Color{
                    color.Color{}, color.Color{},
                },
            },
            want: `
                P3
                3 2
                255
                255 0 0 0 127 0
                0 0 255 0 0 0
                0 0 0 0 0 0
                `,
        },
    }

    for _, test := range tests {
        if got := test.c.ExportToPpm(); got != test.want {
            t.Errorf("canvas \n%s \nexport to ppm \ngot: \n%s. \nexpected: \n%s", test.c, got, test.want)
        }
    }
}
