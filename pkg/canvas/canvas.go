package canvas

import (
	"fmt"
	"image"
	"image/color"
	"rapture/pkg/colors"
	"rapture/pkg/grid"
	"reflect"
)

type Canvas struct {
	grid *grid.Grid
	ramp *colors.ColorRamp
	summ *grid.GridSummary
}

func (cvs *Canvas) Render() image.Image {
	img := image.NewRGBA(cvs.grid.Rect())
	cvs.Fill(img, color.RGBA{R: 0, G: 0, B: 0, A: 255})

	color_axis := grid.NewAxis(cvs.summ.MinVal, cvs.summ.MaxVal, cvs.ramp.Steps())

	filled_cells := cvs.grid.Cells()
	for i := 0; i < filled_cells.Len(); i++ {
		cell := filled_cells.At(i)
		cellVal := reflect.ValueOf(cell.Val).Float()
		colorIndex := color_axis.Index(cellVal, true)
		cellColor := cvs.Color(colorIndex)

		if cellColor == nil {
			panic("cell color is nil")
		}

		img.Set(cell.Col, cell.Row, cellColor)
	}

	return img
}

func (cvs *Canvas) Fill(img *image.RGBA, c color.Color) {
	for row := 0; row < cvs.grid.Height(); row++ {
		for col := 0; col < cvs.grid.Width(); col++ {
			img.Set(col, row, c)
		}
	}
}

func (cvs *Canvas) Color(idx int) color.Color {
	return cvs.ramp.Generate(idx)
}

func NewCanvas(grid *grid.Grid, s *grid.GridSummary) *Canvas {
	fc, _ := colors.ColorFromHex("#FFFF00")
	ec, _ := colors.ColorFromHex("#000000")
	fmt.Println(*s)
	return &Canvas{
		grid: grid,
		ramp: colors.NewColorRamp(
			fc,
			ec,
			100,
		),
		summ: s,
	}
}
