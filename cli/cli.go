package cli

import (
	"fmt"
	"image/png"
	"os"
	"rapture/pkg/canvas"
	"rapture/pkg/grid"

	"github.com/tidwall/geojson"
)

func Run(path string) {
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
	xax := grid.NewAxis(xmin, xmax, 3200)
	yax := grid.NewAxis(ymin, ymax, 3200)
	grd := grid.NewGrid(xax, yax)

	// Add Points
	fmt.Println("Adding points...")
	g.ForEach(func(geom geojson.Object) bool {
		pt := geom.Center()
		grd.AddCell(pt.X, pt.Y, nil)
		return true
	})
	fmt.Printf("Added %d points\n", len(grd.Cells()))

	cvs := canvas.NewCanvas(grd)
	fmt.Println("Rendering")
	img := cvs.Render()

	f, _ := os.Create("example/pts.png")
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
