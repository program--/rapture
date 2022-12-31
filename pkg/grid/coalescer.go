package grid

type Coalescer[T cell_t] interface{ Coalesce(T, T) T }

type Accumulator[T numeric_t] struct{}

func (acc Accumulator[T]) Coalesce(a T, b T) T {
	return a + b
}
