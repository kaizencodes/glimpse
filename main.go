/*
Package main is the entry point for the glimpse ray tracer. It reads a scene
file, renders the scene, and writes the result to a PPM file.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kaizencodes/glimpse/internal/export"
	"github.com/kaizencodes/glimpse/internal/renderer"
	"github.com/kaizencodes/glimpse/internal/scenes/builder"
	"github.com/kaizencodes/glimpse/internal/scenes/reader"
)

var filePath, outputPath, defaultOutputPath string

func init() {
	defaultOutputPath = "renders/render"

	flag.StringVar(&filePath, "f", "", "Filepath for the yml describing the scene.")
	flag.StringVar(&outputPath, "o", defaultOutputPath, "Output path where the render will be saved. Folder has to exist.")
}

const commandHelp = `Usage:
  glimpse [options]

Description:
  Glimpse is a ray tracer. It reads a scene
  file, renders the scene, and writes the result to a PPM file.

Options:
  -h		Show this help message and exit.
  -f		Filepath for the yml describing the scene.
  -o 		Output path where the render will be saved. Folder has to exist.

Examples:
  command -f /examples/marbles.yml
  command -f /examples/marbles.yml -o /renders/new_marble_render

Additional Information:
  - The -o flag has a default value. It defaults to the renders folder.
  - glimpse will append a timestamp and extension to the output file`

func main() {
	start := time.Now()

	flag.Parse()

	if filePath == "" {
		fmt.Println(commandHelp)
		os.Exit(1)
	}

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist\n", filePath)
		os.Exit(1)
	}

	config, err := reader.Read(filePath)
	if err != nil {
		fmt.Printf("The input file has the following error:\n\n %s\n", err.Error())
		os.Exit(1)
	}

	cam, scene := builder.BuildScene(config)

	img := renderer.Render(cam, scene)

	fmt.Printf("\nWriting to file\n")

	if err := os.WriteFile(fmt.Sprintf(outputPath+"-%s.ppm", time.Now().Format(time.RFC3339Nano)), []byte(export.Export(img)), 0666); err != nil {
		fmt.Printf("%e\n", err)
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Total time: %s\n", elapsed)

	os.Exit(0)
}
