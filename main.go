/*
Package main is the entry point for the glimpse ray tracer. It reads a scene
file, renders the scene, and writes the result to a PPM file.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kaizencodes/glimpse/internal/export"
	"github.com/kaizencodes/glimpse/internal/renderer"
	"github.com/kaizencodes/glimpse/internal/scenes/builder"
	"github.com/kaizencodes/glimpse/internal/scenes/reader"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	start := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./glimpse <filepath>")
		return 1
	}

	filepath := os.Args[1]

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist\n", filepath)
		return 1
	}

	config, err := reader.Read(filepath)
	if err != nil {
		fmt.Printf("The input file has the following error:\n\n %s\n", err.Error())
		os.Exit(1)
	}

	cam, scene := builder.BuildScene(config)
	fmt.Printf("Rendering initiated\n")

	img := renderer.Render(cam, scene)

	fmt.Printf("File writing initiated\n")

	filename := fmt.Sprintf("renders/render-%d.ppm", time.Now().Unix())
	if err := os.WriteFile(filename, []byte(export.Export(img)), 0666); err != nil {
		fmt.Printf("%e\n", err)
		log.Fatal(err)
	}
	fmt.Printf("File writing completed\n")
	elapsed := time.Since(start)
	fmt.Printf("Rendering took %s\n", elapsed)

	return 0
}
