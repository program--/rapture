package canvas

import (
	"image"
	"image/color"
	"rapture/pkg/grid"
	"reflect"
)

type Canvas struct {
	grid *grid.Grid
}

func (this *Canvas) Render() image.Image {
	img := image.NewRGBA(this.grid.Rect())
	this.Fill(img, color.RGBA{R: 0, G: 0, B: 0, A: 255})
	filled_cells := this.grid.Cells()
	for i := 0; i < filled_cells.Len(); i++ {
		cell := filled_cells.At(i)
		img.Set(cell.Col, cell.Row, this.Color(cell.Val))
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

func (this *Canvas) Color(val any) color.Color {
	v := reflect.ValueOf(val).Float()
	return color.RGBA{R: 250, G: 0, B: 0, A: uint8(v*100 + 100)}
}

func NewCanvas(grid *grid.Grid) *Canvas {
	return &Canvas{grid: grid}
}
