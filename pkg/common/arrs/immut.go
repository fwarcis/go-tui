package arrs

type Immutable[T any] struct {
	elements []T
}

func NewImmutable[T any](
	elements []T,
) Immutable[T] {
	return Immutable[T]{
		elements: elements,
	}
}

func (i Immutable[T]) At(idx int) T {
	return i.elements[idx]
}

func (i Immutable[T]) Elements(yield func(T) bool) {
	for idx := range i.elements {
		if !yield(i.elements[idx]) {
			return
		}
	}
}

func (i Immutable[T]) All(yield func(int, T) bool) {
	for idx := range i.elements {
		if !yield(idx, i.elements[idx]) {
			return
		}
	}
}

func (i Immutable[T]) Len() int {
	return len(i.elements)
}
