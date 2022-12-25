package cli

import (
	"fmt"
	"image/png"
	"os"
	"rapture/pkg/canvas"
	"rapture/pkg/config"
	"rapture/pkg/geometry"
	"rapture/pkg/grid"
)

func Run(cfg config.RaptureConfig) {
	// Parse file
	features, err := geometry.Parse(cfg.Path)
	if err != nil {
		panic(err)
	}

	// Create grid
	grd := grid.NewGridFromBound(
		features.Extent,
		cfg.Width,
		cfg.Height,
		features.NumFeatures,
	)

	// Add Points
	fmt.Println("Adding points...")
	geometry.MapToGrid(features, cfg.Prop, grd)
	fmt.Printf("Added %d points\n", grd.Cells().Len())

	fmt.Println("Condensing values")
	summary := grd.Cells().Condense(canvas.Density, nil)

	cvs := canvas.NewCanvas(grd, summary)
	fmt.Println("Rendering")
	img := cvs.Render()

	f, err := os.Create(cfg.Output)
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
