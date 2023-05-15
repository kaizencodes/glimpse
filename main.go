package main

import (
	"fmt"
	"glimpse/export"
	"glimpse/reader"
	"glimpse/world/builder"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./glimpse <filepath>")
		return
	}

	filepath := os.Args[1]

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist\n", filepath)
		return
	}

	config, err := reader.Read(filepath)
	if err != nil {
		fmt.Printf("The input file has the following error:\n\n %e\n", err)
		os.Exit(1)
	}

	cam, world := builder.BuildScene(config)
	fmt.Printf("Rendering initiated\n")

	img := cam.Render(world)

	fmt.Printf("File writing initiated\n")

	filename := fmt.Sprintf("renders/render-%d.ppm", time.Now().Unix())
	if err := os.WriteFile(filename, []byte(export.Export(img)), 0666); err != nil {
		fmt.Printf("%e\n", err)
		log.Fatal(err)
	}
	fmt.Printf("File writing completed\n")
	elapsed := time.Since(start)
	fmt.Printf("Rendering took %s\n", elapsed)
	os.Exit(0)
}
