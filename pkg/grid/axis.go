package grid

import (
	"sort"
)

func NewAxis[T numeric_t](from T, to T, length uint) *Axis[T] {
	return &Axis[T]{length, from, to, nil}
}

type Axis[T numeric_t] struct {
	// Length of axis (i.e. number of elements in
	// sequence between (inclusively) fields from and to.)
	length uint

	// Starting value in axis
	from T

	// Ending value in axis
	to T

	// Cached sequence array pointer
	cache *[]T
}

// Returns axis extent as min, max
func (ax *Axis[T]) Bounds() (T, T) {
	return ax.from, ax.to
}

// Returns the dimension of the axis (a.k.a. its length)
func (ax *Axis[T]) Dim() uint {
	return ax.length
}

// Index a degree value onto the axis. Returns -1 if the degree value is off the axis
func (ax *Axis[T]) Index(v T, reverse bool) int {
	seq := ax.Seq()
	var idx = sort.Search(int(ax.length), func(i int) bool { return seq[i] >= v })

	if idx != -1 && reverse {
		return int(ax.length - uint(idx))
	}

	return idx
}

// Returns the Axis as a sequence of equally spaced degree values
func (ax *Axis[T]) Seq() []T {
	if ax.cache == nil {
		vect := make([]T, 0, ax.length)
		diff := T(float64(ax.to-ax.from) / float64(ax.length-1))
		for i := uint(0); i < ax.length; i++ {
			vect = append(vect, ax.from+(diff*T(i)))
		}
		ax.cache = &vect
	}

	return *ax.cache
}
