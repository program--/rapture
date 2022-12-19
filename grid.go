package rapture

import "image"

type GridCell struct {
	val any
	col int // a.k.a. x-coordinate
	row int // a.k.a. y-coordinate
}

type Grid struct {
	cells []GridCell
	xAxis Axis
	yAxis Axis
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

func (this *Grid) Cells() []GridCell {
	return this.cells
}

func (this *Grid) NumCells() int {
	return this.Width() * this.Height()
}

// Index a coordinate pair onto a grid. X or Y can be -1 if they are outside the grid.
func (this *Grid) Index(x float64, y float64) (int, int) {
	return this.xAxis.Index(x), this.yAxis.Index(y)
}

// Adds a point-value to the grid
func (this *Grid) AddCell(x float64, y float64, val any) {
	col, row := this.Index(x, y)
	this.cells = append(this.cells, GridCell{val, col, row})
}

func (this *Grid) Rect() image.Rectangle {
	return image.Rect(0, 0, this.Width(), this.Height())
}

func NewGrid(xAxis Axis, yAxis Axis) *Grid {
	cells := make([]GridCell, 0)
	return &Grid{cells, xAxis, yAxis}
}
