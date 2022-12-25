package grid

import (
	"image"

	"github.com/paulmach/orb"
)

type Grid struct {
	cells *Cells
	xAxis *Axis
	yAxis *Axis
}

// Returns grid extent as xmin, xmax, ymin, ymax
func (grd *Grid) Bounds() (float64, float64, float64, float64) {
	xmin, xmax := grd.xAxis.Bounds()
	ymin, ymax := grd.yAxis.Bounds()
	return xmin, xmax, ymin, ymax
}

func (grd *Grid) Width() int {
	return grd.xAxis.Dim()
}

func (grd *Grid) Height() int {
	return grd.yAxis.Dim()
}

// Returns dimensions of grid as width by height
func (grd *Grid) Dim() (int, int) {
	return grd.Width(), grd.Height()
}

func (grd *Grid) Cells() *Cells {
	return grd.cells
}

func (grd *Grid) NumCells() int {
	return grd.Width() * grd.Height()
}

// Index a coordinate pair onto a grid. X or Y can be -1 if they are outside the grid.
func (grd *Grid) Index(x float64, y float64) (int, int) {
	return grd.xAxis.Index(x, false), grd.yAxis.Index(y, true)
}

// Adds a point-value to the grid
func (grd *Grid) AddCell(x float64, y float64, val any) {
	col, row := grd.Index(x, y)
	grd.cells.AddCell(col, row, val)
}

// Adds coordinates associated to a line segment to the grid
func (grd *Grid) AddLine(x1 float64, y1 float64, x2 float64, y2 float64, val any) {
	col1, row1 := grd.Index(x1, y1)
	col2, row2 := grd.Index(x2, y2)
	colDiff := col2 - col1
	rowDiff := row2 - row1

	for col_idx := col1; col_idx < col2; col_idx++ {
		row_idx := ((rowDiff)/(colDiff))*(col_idx-col1) + row1
		grd.cells.AddCell(col_idx, row_idx, val)
	}
}

func (grd *Grid) Rect() image.Rectangle {
	return image.Rect(0, 0, grd.Width(), grd.Height())
}

func NewGrid(xAxis *Axis, yAxis *Axis, points int) *Grid {
	return &Grid{NewCells(points), xAxis, yAxis}
}

// Creates a new grid from a given geojson object.
func NewGridFromBound(b orb.Bound, width int, height int, n int) *Grid {
	// Get extent
	xmin := b.Min.X()
	ymin := b.Min.Y()
	xmax := b.Max.X()
	ymax := b.Max.Y()

	// Setup Axes
	xax := NewAxis(xmin, xmax, width)
	yax := NewAxis(ymin, ymax, height)

	// Setup Grid
	return NewGrid(xax, yax, n)
}
