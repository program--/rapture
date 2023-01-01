package util

type Coalescer[T Cell_t] interface{ Coalesce(T, T) T }

type Accumulator[T Numeric_t] struct{}

func (acc Accumulator[T]) Coalesce(a T, b T) T {
	return a + b
}
