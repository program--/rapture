package cli

import (
	"fmt"
	"rapture/pkg/config"
	"rapture/pkg/geometry"
	"rapture/pkg/grid"
)

func Run(cfg config.RaptureConfig) {
	// Parse file
	fmt.Println("Parsing")
	features, err := geometry.NewFeatureCollection(cfg.Path)
	if err != nil {
		panic(err)
	}

	// Create grid
	fmt.Println("Creating grid...")
	grd := grid.NewGridFromBound[float64](features.Extent, cfg.Width, cfg.Height, uint(len(features.Features)))

	// Add Points
	fmt.Println("Adding points...")
	n, err := grd.MapFeatures(features, "POPULATION_2020")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added %d points\n", n)

	// fmt.Println("Condensing values")
	// summary := grd.Cells().Condense(canvas.Density, nil)

	grd.Render(cfg.Output)
}
