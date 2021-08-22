package canvas

import (
    "glimpse/color"
    "testing"
)

func TestNewCanvas(t *testing.T) {
    var tests = []struct {
        w    int
        h    int
        want *Canvas
    }{
        {
            w: 2,
            h: 3,
            want: &Canvas{
                pane: [][]color.Color{
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
        },
    }

    for _, test := range tests {
        if got := NewCanvas(test.w, test.h); got.String() != test.want.String() {
            t.Errorf("canvas width w:%d, h:%d \ngot: \n%s. \nexpected: \n%s", test.w, test.h, got, test.want)
        }
    }
}

func TestWritePixel(t *testing.T) {
    var tests = []struct {
        pixel color.Color
        w     int
        h     int
        want  *Canvas
    }{
        {
            pixel: color.Color{0.5, 0.3, 0.2},
            w:     1,
            h:     0,
            want: &Canvas{
                pane: [][]color.Color{
                    []color.Color{
                        color.Color{}, color.Color{0.5, 0.3, 0.2},
                    },
                    []color.Color{
                        color.Color{}, color.Color{},
                    },
                    []color.Color{
                        color.Color{}, color.Color{},
                    },
                },
            },
        },
    }

    for _, test := range tests {
        canvas := NewCanvas(2, 3)
        canvas.WritePixel(test.w, test.h, test.pixel)
        if got := canvas.String(); got != test.want.String() {
            t.Errorf("canvas write pixel %s \nat w:%d, h:%d \ngot: \n%s. \nexpected: \n%s", test.pixel.String(), test.w, test.h, got, test.want)
        }
    }
}

func TestExportToPpm(t *testing.T) {
    var tests = []struct {
        c    *Canvas
        want string
    }{
        {
            c: &Canvas{
                pane: [][]color.Color{
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
