package canvas

import (
	"image"
	"image/color"
	"rapture/pkg/grid"
)

type Canvas struct {
	grid       grid.Grid
	colorspace color.Model
}

func (this *Canvas) Render() image.Image {
	img := image.NewRGBA(this.grid.Rect())
	filled_cells := this.grid.Cells()
	for i := 0; i < len(filled_cells); i++ {
		cell := filled_cells[i]
		img.Set(cell.Col, cell.Row, this.Color(cell.Val))
	}

	return img
}

func (this *Canvas) Color(val any) color.Color {
	panic("not implemented")
}
