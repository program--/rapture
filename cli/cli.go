package cli

import (
	"fmt"
	"image/png"
	"os"
	"rapture/pkg/canvas"
	"rapture/pkg/config"
	"rapture/pkg/geometry"
	"rapture/pkg/grid"

	"github.com/tidwall/geojson"
)

func Run(cfg config.RaptureConfig) {
	// Read File
	s, err := os.ReadFile(cfg.Path)
	if err != nil {
		panic(err)
	}

	// Parse GeoJSON (currently bottleneck)
	fmt.Println("Parsing file...")
	g, err := geojson.Parse(string(s), geojson.DefaultParseOptions)
	if err != nil {
		panic(err)
	}

	// Create Grid
	fmt.Println("Creating grid...")
	grd := grid.NewGridFromGeojson(g, cfg.Width, cfg.Height)

	// Add Points
	fmt.Println("Adding points...")
	geometry.MapToGrid(g, cfg.Prop, grd)
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
