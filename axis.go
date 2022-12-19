package rapture

import "sort"

type Axis struct {
	length int
	from   float64
	to     float64
}

// Returns axis extent as min, max
func (this *Axis) Bounds() (float64, float64) {
	return this.from, this.to
}

func (this *Axis) Dim() int {
	return this.length
}

// Index a degree value onto the axis. Returns -1 if the degree value is off the axis
func (this *Axis) Index(degree float64) int {
	idx := sort.SearchFloat64s(this.Seq(), degree)
	if idx == this.length {
		return -1
	}

	return idx
}

// Returns the Axis as a sequence of equally spaced degree values
func (this *Axis) Seq() []float64 {
	diff := (this.to - this.from) / float64(this.length-1)
	vec := make([]float64, this.length)
	vec[0] = this.from
	for i := 1; i < this.length-1; i++ {
		vec[i] = vec[i-1] + diff
	}
	vec[this.length-1] = this.to
	return vec
}

func NewAxis(from float64, to float64, length int) *Axis {
	return &Axis{length, from, to}
}
