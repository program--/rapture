package grid

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

func NewCells(size int) *Cells {
	cells := new(Cells)
	cells.Allocate(size)
	return cells
}

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

func (this *Cells) Len() int {
	return len(this.Cols)
}

func (this *Cells) Cap() int {
	return cap(this.Cols)
}

func (this *Cells) At(index int) *Cell {
	return &Cell{
		Val: this.Vals[index],
		Col: this.Cols[index],
		Row: this.Rows[index],
	}
}
