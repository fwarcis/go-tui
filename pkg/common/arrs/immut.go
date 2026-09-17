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

func (s Immutable[T]) At(i int) T {
	return s.elements[i]
}

func (s Immutable[T]) Elements(yield func(T) bool) {
	for i := range s.elements {
		if !yield(s.elements[i]) {
			return
		}
	}
}

func (s Immutable[T]) All(yield func(int, T) bool) {
	for i := range s.elements {
		if !yield(i, s.elements[i]) {
			return
		}
	}
}

func (s Immutable[T]) Len() int {
	return len(s.elements)
}
