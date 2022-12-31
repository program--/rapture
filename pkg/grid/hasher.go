package grid

import "math"

type Hasher interface {
	Hash(uint, uint) uint64
	Unhash(uint64) (uint, uint)
}

type SimpleHasher struct{}

func (SimpleHasher) Hash(x uint, y uint) uint64 {
	if x < y {
		return uint64(math.Pow(float64(y), 2.0) + float64(x))
	} else {
		return uint64(math.Pow(float64(x), 2.0) + float64(x) + float64(y))
	}
}

func (SimpleHasher) Unhash(z uint64) (uint, uint) {
	zS := math.Floor(math.Sqrt(float64(z)))
	zP := math.Pow(zS, 2.0)
	zF := float64(z)

	if zF-zP < zS {
		return uint(zF - zP), uint(zS)
	} else {
		return uint(zS), uint(zF - zP - zS)
	}

}
