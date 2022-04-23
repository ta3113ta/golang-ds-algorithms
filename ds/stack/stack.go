package stack

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Pop removes and returns the top element of the stack.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}

	idx := len(*s) - 1
	v := (*s)[idx]
	*s = (*s)[:idx]

	return v, true
}

// Peek returns the top element of the stack without removing it.
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}

	idx := len(*s) - 1
	v := (*s)[idx]

	return v, true
}
