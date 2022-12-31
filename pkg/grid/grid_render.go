package grid

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/mazznoer/colorgrad"
)

func (grd *Grid[T]) Render(output string) {
	ctx := gg.NewContext(int(grd.Width()), int(grd.Height()))
	ctx.SetColor(color.Black)
	ctx.Clear()
	pal := colorgrad.OrRd().ColorfulColors(100)
	fmt.Printf("Mapping %d cells\n", grd.NumFilledCells())
	for i, cell := range grd.cells.cells {
		ctx.Push()
		ctx.SetColor(pal[i%100])
		ctx.DrawPoint(float64(cell.column), float64(cell.row), 0.075)
		ctx.Fill()
		ctx.Pop()
	}

	ctx.SavePNG(output)
}
