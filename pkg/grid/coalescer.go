package grid

type Coalescer[T cell_t] interface{ Coaelesce(T, T) T }
