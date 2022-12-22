package cli

import (
	"fmt"
	"image/png"
	"os"
	"rapture/pkg/canvas"
	"rapture/pkg/geometry"
	"rapture/pkg/grid"

	"github.com/tidwall/geojson"
)

func Run(path string, width int, height int, output string) {
	s, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Parsing file...")
	g, err := geojson.Parse(string(s), geojson.DefaultParseOptions)
	if err != nil {
		panic(err)
	}

	// Get Extent
	fmt.Println("Getting extent...")
	rect := g.Rect()
	xmin := rect.Min.X
	ymin := rect.Min.Y
	xmax := rect.Max.X
	ymax := rect.Max.Y

	// Setup Grid
	fmt.Println("Setting up grid...")
	xax := grid.NewAxis(xmin, xmax, width)
	yax := grid.NewAxis(ymin, ymax, height)
	grd := grid.NewGrid(xax, yax, g.NumPoints())

	// Add Points
	fmt.Println("Adding points...")
	summary := geometry.MapToGrid(g, grd)

	fmt.Printf("Added %d points\n", grd.Cells().Len())

	fmt.Println("Condensing values")
	grd.Cells().Condense(canvas.Density, summary.MaxVal)

	cvs := canvas.NewCanvas(grd)
	fmt.Println("Rendering")
	img := cvs.Render()

	f, err := os.Create(output)
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
