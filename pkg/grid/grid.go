package grid

import (
	"image"
)

type GridCell struct {
	Val any
	Col int // a.k.a. x-coordinate
	Row int // a.k.a. y-coordinate
}

type Grid struct {
	cells *Cells
	xAxis *Axis
	yAxis *Axis
}

// Returns grid extent as xmin, xmax, ymin, ymax
func (this *Grid) Bounds() (float64, float64, float64, float64) {
	xmin, xmax := this.xAxis.Bounds()
	ymin, ymax := this.yAxis.Bounds()
	return xmin, xmax, ymin, ymax
}

func (this *Grid) Width() int {
	return this.xAxis.Dim()
}

func (this *Grid) Height() int {
	return this.yAxis.Dim()
}

// Returns dimensions of grid as width by height
func (this *Grid) Dim() (int, int) {
	return this.Width(), this.Height()
}

func (this *Grid) Cells() *Cells {
	return this.cells
}

func (this *Grid) NumCells() int {
	return this.Width() * this.Height()
}

// Index a coordinate pair onto a grid. X or Y can be -1 if they are outside the grid.
func (this *Grid) Index(x float64, y float64) (int, int) {
	return this.xAxis.Index(x, false), this.yAxis.Index(y, true)
}

// Adds a point-value to the grid
func (this *Grid) AddCell(x float64, y float64, val any) {
	col, row := this.Index(x, y)
	this.cells.AddCell(col, row, val)
}

// Adds coordinates associated to a line segment to the grid
func (this *Grid) AddLine(x1 float64, y1 float64, x2 float64, y2 float64, val any) {
	col1, row1 := this.Index(x1, y1)
	col2, row2 := this.Index(x2, y2)
	colDiff := col2 - col1
	rowDiff := row2 - row1

	for col_idx := col1; col_idx < col2; col_idx++ {
		row_idx := ((rowDiff)/(colDiff))*(col_idx-col1) + row1
		this.cells.AddCell(col_idx, row_idx, val)
	}
}

func (this *Grid) Rect() image.Rectangle {
	return image.Rect(0, 0, this.Width(), this.Height())
}

func NewGrid(xAxis *Axis, yAxis *Axis, points int) *Grid {
	return &Grid{NewCells(points), xAxis, yAxis}
}
