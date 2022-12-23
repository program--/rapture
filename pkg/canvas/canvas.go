package canvas

import (
	"fmt"
	"image"
	"image/color"
	"rapture/pkg/grid"
	"reflect"
)

type Canvas struct {
	grid *grid.Grid
	ramp *ColorRamp
	summ *grid.GridSummary
}

func (this *Canvas) Render() image.Image {
	img := image.NewRGBA(this.grid.Rect())
	this.Fill(img, color.RGBA{R: 0, G: 0, B: 0, A: 255})

	color_axis := grid.NewAxis(this.summ.MinVal, this.summ.MaxVal, this.ramp.steps)

	filled_cells := this.grid.Cells()
	for i := 0; i < filled_cells.Len(); i++ {
		cell := filled_cells.At(i)
		cellVal := reflect.ValueOf(cell.Val).Float()
		colorIndex := color_axis.Index(cellVal, true)
		cellColor := this.Color(colorIndex)

		if cellColor == nil {
			panic("cell color is nil")
		}

		img.Set(cell.Col, cell.Row, cellColor)
	}

	return img
}

func (this *Canvas) Fill(img *image.RGBA, c color.Color) {
	for row := 0; row < this.grid.Height(); row++ {
		for col := 0; col < this.grid.Width(); col++ {
			img.Set(col, row, c)
		}
	}
}

func (this *Canvas) Color(idx int) color.Color {
	return this.ramp.Generate(idx)
}

func NewCanvas(grid *grid.Grid, s *grid.GridSummary) *Canvas {
	fc, _ := ColorFromHex("#FFFF00")
	ec, _ := ColorFromHex("#000000")
	fmt.Println(*s)
	return &Canvas{
		grid: grid,
		ramp: NewColorRamp(
			fc,
			ec,
			100,
		),
		summ: s,
	}
}
