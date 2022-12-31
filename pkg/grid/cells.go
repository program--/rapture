package grid

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

type Cell[T cell_t] struct {
	value  T
	column uint
	row    uint
}

type CellSummary[T cell_t] struct {
	Max T
	Min T
	Avg T
}

type CellArray[T cell_t] struct {
	mu      sync.Mutex
	cells   []Cell[T]
	summary *CellSummary[T]
}

func NewCellArray[T cell_t](size uint) *CellArray[T] {
	cells := new(CellArray[T])
	cells.Allocate(size)
	return cells
}

// returns allocation size of cells
func (c *CellArray[T]) Cap() uint {
	return uint(cap(c.cells))
}

// returns number of cells
func (c *CellArray[T]) Len() uint {
	return uint(len(c.cells))
}

// returns a cell at a given index
func (c *CellArray[T]) At(index int) *Cell[T] {
	return &c.cells[index]
}

// Allocate size to cells arrays. If cells is already used,
// then initial values are copied to newly allocated arrays.
func (c *CellArray[T]) Allocate(size uint) {
	totalCapacity := c.Cap() + size
	currentLength := c.Len()
	newCellsArray := make([]Cell[T], currentLength, totalCapacity)
	if currentLength > 0 {
		copy(newCellsArray, c.cells)
	}

	c.cells = newCellsArray
}

// Add cell with given col, row, and value
func (c *CellArray[T]) Add(column uint, row uint, value T) {
	if c.Len() >= c.Cap() {
		c.Allocate(c.Cap()) // Double allocation
	}

	c.mu.Lock() // mapping is done in a goroutine
	c.cells = append(c.cells, Cell[T]{value, column, row})
	c.mu.Unlock()
}

// Applies some operation defined in reducer across all hashed points
func (c *CellArray[T]) Condense(hasher Hasher, reducer Coalescer[T]) {
	wg := sync.WaitGroup{}
	newLen := uint64(0)
	cellMap := sync.Map{}
	cellMapReduce := func(k uint64, v T, r Coalescer[T]) {
		defer wg.Done()
		if value, ok := cellMap.Load(k); ok {
			cellMap.Store(k, r.Coalesce(value.(T), v))
		} else {
			cellMap.Store(k, r.Coalesce(T(0), v))
			atomic.AddUint64(&newLen, 1)
		}
	}

	// Perform condensing
	wg.Add(int(c.Len()))
	for _, cell := range c.cells {
		key := hasher.Hash(cell.column, cell.row)
		go cellMapReduce(key, cell.value, reducer)
	}
	wg.Wait()

	// Rematerialize cells into new array with length <= original array
	newCells := make([]Cell[T], 0, newLen)
	cellMap.Range(func(key any, value any) bool {
		column, row := hasher.Unhash(key.(uint64))
		cell := Cell[T]{value.(T), column, row}
		newCells = append(newCells, cell)
		return true
	})
}

func (c *CellArray[T]) Summarise() *CellSummary[T] {
	if c.summary == nil {
		switch any(c.cells[0].value).(type) {
		case string:
			panic("summaries for strings is not implemented")
		}

		max := T(math.Inf(-1))
		min := T(math.Inf(1))
		avg := T(0)

		for _, cell := range c.cells {
			if cell.value > max {
				max = cell.value
			}

			if cell.value < min {
				min = cell.value
			}

			avg += cell.value
		}

		avg = avg / T(len(c.cells))

		c.summary = &CellSummary[T]{max, min, avg}
		fmt.Printf("Created grid summary: {Max: %v, Min: %v, Avg: %v}", c.summary.Max, c.summary.Min, c.summary.Avg)
	}

	return c.summary
}
