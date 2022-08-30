package set

type Set[T comparable] struct{ base map[T]bool }

func NewSet[T comparable]() Set[T] {
	return Set[T]{base: make(map[T]bool)}
}

func (s *Set[T]) Insert(value T) bool {
	if s.Contains(value) {
		return true
	}

	s.base[value] = true
	return true
}

func (s *Set[T]) Contains(value T) bool {
	return s.base[value]
}

func (s *Set[T]) Remove(value T) bool {
	if !s.Contains(value) {
		return false
	}

	delete(s.base, value)
	return true
}

func (s *Set[T]) Clear() {
	for k := range s.base {
		delete(s.base, k)
	}
}

func (s *Set[T]) Len() int {
	return len(s.base)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}
