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
func (c *Cells) Allocate(size int) {
	tot_cap := c.Cap() + size
	l := c.Len()
	vals := make([]any, l, tot_cap)
	cols := make([]int, l, tot_cap)
	rows := make([]int, l, tot_cap)

	if l > 0 {
		copy(vals, c.Vals)
		copy(cols, c.Cols)
		copy(rows, c.Rows)
	}

	c.Vals = vals
	c.Cols = cols
	c.Rows = rows
}

// Add cell with given col, row, and value
func (c *Cells) AddCell(col int, row int, val any) {
	iter := c.Len()
	if iter >= c.Cap() {
		// Double array allocation
		c.Allocate(c.Cap())
	}

	c.Vals = append(c.Vals, val)
	c.Cols = append(c.Cols, col)
	c.Rows = append(c.Rows, row)
}

// applies a function across all allocated cells
func (c *Cells) Condense(f func(any, any, any) any, opts any) *GridSummary {
	index := make(map[int64]any)
	for i := 0; i < c.Len(); i++ {
		key := elepair(c.Cols[i], c.Rows[i])
		if _, ok := index[key]; ok {
			index[key] = f(index[key], c.Vals[i], opts)
		} else {
			index[key] = f(0.0, c.Vals[i], opts)
		}
	}

	summary := &GridSummary{
		MaxVal: math.Inf(-1),
		MinVal: math.Inf(1),
		AvgVal: 0.0,
	}

	new_len := len(index)
	if new_len != c.Len() {
		c.Cols = make([]int, 0, new_len)
		c.Rows = make([]int, 0, new_len)
		c.Vals = make([]any, 0, new_len)

		for k, v := range index {
			col, row := eleunpair(k)
			c.Cols = append(c.Cols, col)
			c.Rows = append(c.Rows, row)
			c.Vals = append(c.Vals, v)

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
func (c *Cells) Len() int {
	return len(c.Cols)
}

// returns allocation size of cells
func (c *Cells) Cap() int {
	return cap(c.Cols)
}

// returns cell at input index
func (c *Cells) At(index int) *Cell {
	return &Cell{
		Val: c.Vals[index],
		Col: c.Cols[index],
		Row: c.Rows[index],
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
