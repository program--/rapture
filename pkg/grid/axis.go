package grid

import "sort"

type Axis struct {
	length int
	from   float64
	to     float64
	cache  *[]float64
}

// Returns axis extent as min, max
func (this *Axis) Bounds() (float64, float64) {
	return this.from, this.to
}

func (this *Axis) Dim() int {
	return this.length
}

// Index a degree value onto the axis. Returns -1 if the degree value is off the axis
func (this *Axis) Index(degree float64, reverse bool) int {
	idx := sort.SearchFloat64s(this.Seq(), degree)
	if idx == this.length {
		return -1
	}

	if reverse {
		return this.length - idx
	} else {
		return idx
	}
}

// Returns the Axis as a sequence of equally spaced degree values
func (this *Axis) Seq() []float64 {
	var vec []float64

	if this.cache == nil {
		diff := (this.to - this.from) / float64(this.length-1)
		vec = make([]float64, 0, this.length)

		for i := 0; i < this.length; i++ {
			vec = append(vec, this.from+(diff*float64(i)))
		}

		this.cache = &vec
	}

	return *this.cache
}

func NewAxis(from float64, to float64, length int) *Axis {
	return &Axis{length, from, to, nil}
}
