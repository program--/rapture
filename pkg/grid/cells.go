package grid

import (
	"math"
	"reflect"
)

type Cell struct {
	Val any
	Col int
	Row int
}

type Cells struct {
	Vals []any
	Cols []int // x
	Rows []int // y
}

type GridSummary struct {
	MaxVal float64
	MinVal float64
	AvgVal float64
}

func NewCells(size int) *Cells {
	cells := new(Cells)
	cells.Allocate(size)
	return cells
}

// Allocate size to cells arrays. If cells is already used,
// then initial values are copied to newly allocated arrays.
func (this *Cells) Allocate(size int) {
	tot_cap := this.Cap() + size
	l := this.Len()
	vals := make([]any, l, tot_cap)
	cols := make([]int, l, tot_cap)
	rows := make([]int, l, tot_cap)

	if l > 0 {
		copy(vals, this.Vals)
		copy(cols, this.Cols)
		copy(rows, this.Rows)
	}

	this.Vals = vals
	this.Cols = cols
	this.Rows = rows
}

// Add cell with given col, row, and value
func (this *Cells) AddCell(col int, row int, val any) {
	iter := this.Len()
	if iter >= this.Cap() {
		// Double array allocation
		this.Allocate(this.Cap())
	}

	this.Vals = append(this.Vals, val)
	this.Cols = append(this.Cols, col)
	this.Rows = append(this.Rows, row)
}

// applies a function across all allocated cells
func (this *Cells) Condense(f func(any, any, any) any, opts any) *GridSummary {
	index := make(map[int64]any)
	for i := 0; i < this.Len(); i++ {
		key := elepair(this.Cols[i], this.Rows[i])
		if _, ok := index[key]; ok {
			index[key] = f(index[key], this.Vals[i], opts)
		} else {
			index[key] = f(0.0, this.Vals[i], opts)
		}
	}

	summary := &GridSummary{
		MaxVal: math.Inf(-1),
		MinVal: math.Inf(1),
		AvgVal: 0.0,
	}

	new_len := len(index)
	if new_len != this.Len() {
		this.Cols = make([]int, 0, new_len)
		this.Rows = make([]int, 0, new_len)
		this.Vals = make([]any, 0, new_len)

		for k, v := range index {
			col, row := eleunpair(k)
			this.Cols = append(this.Cols, col)
			this.Rows = append(this.Rows, row)
			this.Vals = append(this.Vals, v)

			rv := reflect.ValueOf(v).Float()
			if rv < summary.MinVal {
				summary.MinVal = rv
			}

			if rv > summary.MaxVal {
				summary.MaxVal = rv
			}

			summary.AvgVal += rv
		}

		summary.AvgVal /= float64(new_len)
	}

	return summary
}

// returns number of cells
func (this *Cells) Len() int {
	return len(this.Cols)
}

// returns allocation size of cells
func (this *Cells) Cap() int {
	return cap(this.Cols)
}

// returns cell at input index
func (this *Cells) At(index int) *Cell {
	return &Cell{
		Val: this.Vals[index],
		Col: this.Cols[index],
		Row: this.Rows[index],
	}
}

func elepair(x int, y int) int64 {
	if x < y {
		return int64(math.Pow(float64(y), 2.0) + float64(x))
	} else {
		return int64(math.Pow(float64(x), 2.0) + float64(x) + float64(y))
	}
}

func eleunpair(z int64) (int, int) {
	sqfl := math.Floor(math.Sqrt(float64(z)))
	sqflsq := math.Pow(sqfl, 2.0)
	zf := float64(z)

	if zf-sqflsq < sqfl {
		return int(zf - sqflsq), int(sqfl)
	} else {
		return int(sqfl), int(zf - sqflsq - sqfl)
	}
}
