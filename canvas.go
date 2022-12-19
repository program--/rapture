package rapture

import (
	"image"
	"image/color"
)

type Canvas struct {
	grid       Grid
	colorspace color.Model
}

func (this *Canvas) Render() image.Image {
	img := image.NewRGBA(this.grid.Rect())
	filled_cells := this.grid.Cells()
	for i := 0; i < len(filled_cells); i++ {
		cell := filled_cells[i]
		img.Set(cell.col, cell.row, this.Color(cell.val))
	}

	return img
}

func (this *Canvas) Color(val any) color.Color {
	panic("not implemented")
}
