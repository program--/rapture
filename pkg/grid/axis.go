package grid

import "sort"

type Axis struct {
	length int
	from   float64
	to     float64
	cache  *[]float64
}

// Returns axis extent as min, max
func (ax *Axis) Bounds() (float64, float64) {
	return ax.from, ax.to
}

// Returns the dimension of the axis (a.k.a. its length)
func (ax *Axis) Dim() int {
	return ax.length
}

// Index a degree value onto the axis. Returns -1 if the degree value is off the axis
func (ax *Axis) Index(degree float64, reverse bool) int {
	idx := sort.SearchFloat64s(ax.Seq(), degree)
	if idx == ax.length {
		return -1
	}

	if reverse {
		return ax.length - idx
	} else {
		return idx
	}
}

// Returns the Axis as a sequence of equally spaced degree values
func (ax *Axis) Seq() []float64 {
	var vec []float64

	if ax.cache == nil {
		diff := (ax.to - ax.from) / float64(ax.length-1)
		vec = make([]float64, 0, ax.length)

		for i := 0; i < ax.length; i++ {
			vec = append(vec, ax.from+(diff*float64(i)))
		}

		ax.cache = &vec
	}

	return *ax.cache
}

func NewAxis(from float64, to float64, length int) *Axis {
	return &Axis{length, from, to, nil}
}
