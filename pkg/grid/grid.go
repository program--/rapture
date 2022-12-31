package grid

import (
	"image"

	"github.com/paulmach/orb"
)

type Grid[T cell_t] struct {
	cells *CellArray[T]
	xAxis *Axis[float64]
	yAxis *Axis[float64]
}

func NewGrid[T cell_t](xAxis *Axis[float64], yAxis *Axis[float64], points uint) *Grid[T] {
	return &Grid[T]{NewCellArray[T](points), xAxis, yAxis}
}

func NewGridFromBound[T cell_t](b *orb.Bound, width uint, height uint, n uint) *Grid[T] {
	xax := NewAxis(b.Min.X(), b.Max.X(), width)
	yax := NewAxis(b.Min.Y(), b.Max.Y(), height)
	return NewGrid[T](xax, yax, n)
}

func (grd *Grid[T]) Bounds() *orb.Bound {
	xmin, xmax := grd.xAxis.Bounds()
	ymin, ymax := grd.yAxis.Bounds()
	return &orb.Bound{
		Max: orb.Point{xmax, ymax},
		Min: orb.Point{xmin, ymin},
	}
}

func (grd *Grid[T]) Width() uint {
	return grd.xAxis.Dim()
}

func (grd *Grid[T]) Height() uint {
	return grd.yAxis.Dim()
}

func (grd *Grid[T]) Rect() image.Rectangle {
	return image.Rect(0, 0, int(grd.Width()), int(grd.Height()))
}

func (grd *Grid[T]) Dim() (width uint, height uint) {
	return grd.Width(), grd.Height()
}

func (grd *Grid[T]) Cells() *CellArray[T] {
	return grd.cells
}

func (grd *Grid[T]) NumCells() uint {
	return grd.Width() * grd.Height()
}

func (grd *Grid[T]) NumFilledCells() uint {
	return grd.cells.Len()
}

func (grd *Grid[T]) NumEmptyCells() uint {
	return grd.NumCells() - grd.NumFilledCells()
}

func (grd *Grid[T]) IndexXY(x float64, y float64) (column int, row int) {
	return grd.xAxis.Index(x, false), grd.yAxis.Index(y, true)
}

func (grd *Grid[T]) Index(p *orb.Point) (column int, row int) {
	return grd.IndexXY(p.X(), p.Y())
}

func (grd *Grid[T]) AddPoint(p *orb.Point, value T) error {
	column, row := grd.Index(p)

	if err := checkGridIndex(column, row, p); err != nil {
		return err
	}

	grd.cells.Add(uint(column), uint(row), value)
	return nil
}

func (grd *Grid[T]) AddSegment(p1 *orb.Point, p2 *orb.Point, value T) error {
	c1, r1 := grd.Index(p1)
	c2, r2 := grd.Index(p2)

	if err := checkGridIndex(c1, r1, p1); err != nil {
		return err
	}

	if err := checkGridIndex(c2, r2, p2); err != nil {
		return err
	}

	cDiff := uint(c2 - c1)
	rDiff := uint(r2 - r1)
	for column := uint(c1); column < uint(c2); column++ {
		row := (rDiff/cDiff)*(column-uint(c1)) + uint(r1)
		grd.cells.Add(column, row, value)
	}

	return nil
}
