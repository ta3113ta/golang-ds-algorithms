package queue

type Queue[T any] struct {
	data []T
}

// New create a new queue
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue add an element to the end of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.data = append(q.data, item)
}

// Dequeue remove the first element in the queue
func (q *Queue[T]) Dequeue() (value T, ok bool) {
	if q.IsEmpty() {
		return
	}

	value = q.data[0]
	q.data = q.data[1:]
	ok = true

	return
}

// Peek return the first element in the queue
func (q *Queue[T]) Peek() (value T, ok bool) {
	if q.IsEmpty() {
		return
	}

	value = q.data[0]
	ok = true

	return
}

// Len return the length of the queue
func (q *Queue[T]) Len() int32 {
	return int32(len(q.data))
}

// Empty check if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}

// Clear clear all elements in the queue
func (q *Queue[T]) Clear() {
	q.data = []T{}
}
