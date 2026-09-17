package arrs

type Immutable[T any] struct {
	values []T
}

func NewImmutable[T any](
	values []T,
) Immutable[T] {
	return Immutable[T]{
		values: values,
	}
}

func (i Immutable[T]) At(idx int) T {
	return i.values[idx]
}

func (i Immutable[T]) Values(yield func(T) bool) {
	for idx := range i.values {
		if !yield(i.values[idx]) {
			return
		}
	}
}

func (i Immutable[T]) All(yield func(int, T) bool) {
	for idx := range i.values {
		if !yield(idx, i.values[idx]) {
			return
		}
	}
}

func (i Immutable[T]) Len() int {
	return len(i.values)
}
