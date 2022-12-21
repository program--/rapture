package cli

import (
	"fmt"
	"image/png"
	"os"
	"rapture/pkg/canvas"
	"rapture/pkg/grid"

	"github.com/tidwall/geojson"
)

func Run(path string, width int, height int, output string) {
	s, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	g, err := geojson.Parse(string(s), geojson.DefaultParseOptions)
	if err != nil {
		panic(err)
	}

	// Get Extent
	rect := g.Rect()
	xmin := rect.Min.X
	ymin := rect.Min.Y
	xmax := rect.Max.X
	ymax := rect.Max.Y

	// Setup Grid
	xax := grid.NewAxis(xmin, xmax, width)
	yax := grid.NewAxis(ymin, ymax, height)
	grd := grid.NewGrid(xax, yax, g.NumPoints())

	// Add Points
	fmt.Println("Adding points...")
	g.ForEach(func(geom geojson.Object) bool {
		pt := geom.Center()
		grd.AddCell(pt.X, pt.Y, 1)
		return true
	})

	fmt.Printf("Added %d points\n", grd.Cells().Len())

	fmt.Println("Condensing values")
	grd.Cells().Condense(canvas.Density)

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
