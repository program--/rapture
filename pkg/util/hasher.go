package util

import "math"

type Hasher interface {
	Hash(uint, uint) uint64
	Unhash(uint64) (uint, uint)
}

// Implements a simple hasher to get a unique value for each pairing of XY values.
// Space filling curves could be used as well, though I'm not sure there's any benefit.
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

type MortonHasher struct{}

func (MortonHasher) split(a uint) uint64 {
	x := uint64(a) & 0x1fffff
	x = (x | x<<32) & 0x1f00000000ffff
	x = (x | x<<16) & 0x1f0000ff0000ff
	x = (x | x<<8) & 0x100f00f00f00f00f
	x = (x | x<<4) & 0x10c30c30c30c30c3
	x = (x | x<<2) & 0x1249249249249249
	return x
}

func (MortonHasher) compact(z uint64) uint {
	a := z & 0x1249249249249249
	a = (a ^ (a >> 1)) & 0x10c30c30c30c30c3
	a = (a ^ (a >> 2)) & 0x100f00f00f00f00f
	a = (a ^ (a >> 4)) & 0x1f0000ff0000ff
	a = (a ^ (a >> 8)) & 0x1f00000000ffff
	a = (a ^ (a >> 16)) & 0x1fffff
	return uint(a)
}

func (m MortonHasher) Hash(x uint, y uint) uint64 {
	return uint64(0) | m.split(x) | (m.split(y) << 1)
}

func (m MortonHasher) Unhash(z uint64) (uint, uint) {
	return m.compact(z), m.compact(z >> 1)
}
