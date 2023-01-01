package grid

import (
	"fmt"

	"github.com/fogleman/gg"
)

func (grd *Grid[T]) Render(output string, padding uint) {
	var nc uint

	ctx := gg.NewContext(int(grd.Width())+2*int(padding), int(grd.Height())+2*int(padding))

	if grd.opts.BackgroundColor != nil {
		ctx.SetColor(grd.opts.BackgroundColor)
		ctx.Clear()
	}

	nc = grd.cAxis.Dim()

	pal := grd.opts.Palette.Colors(nc)
	fmt.Printf("Mapping %d cells\n", grd.NumFilledCells())

	for _, cell := range grd.cells.cells {
		ctx.Push()
		ctx.SetColor(pal[grd.cAxis.Index(cell.value, false)])
		ctx.DrawPoint(float64(cell.column+padding), float64(cell.row+padding), 0.075)
		ctx.Fill()
		ctx.Pop()
	}

	ctx.SavePNG(output)
}
