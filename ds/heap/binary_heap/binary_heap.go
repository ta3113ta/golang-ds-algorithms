package binary_heap

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Heap[T any] struct {
	count      int
	items      []T
	comparator func(T, T) bool
}

func New[T any](comparator func(T, T) bool) *Heap[T] {
	return &Heap[T]{
		count:      0,
		items:      make([]T, 1),
		comparator: comparator,
	}
}

// NewMin create a new min heap
func NewMin[T constraints.Ordered]() *Heap[T] {
	fn := func(a, b T) bool {
		return a < b
	}

	return New(fn)
}

// NewMax create a new max heap
func NewMax[T constraints.Ordered]() *Heap[T] {
	fn := func(a, b T) bool {
		return a > b
	}

	return New(fn)
}

// Len returns the number of items in the heap.
func (h *Heap[T]) Len() int {
	return h.count
}

// IsEmpty returns true if the heap is empty.
func (h *Heap[T]) IsEmpty() bool {
	return h.count == 0
}

// Add adds an element to the heap.
func (h *Heap[T]) Add(value T) {
	h.count++
	h.items = append(h.items, value)

	// Heapify up
	idx := h.count
	for h.parentIdx(idx) > 0 {
		pdx := h.parentIdx(idx)
		if h.comparator(h.items[idx], h.items[pdx]) {
			// swap
			h.items[idx], h.items[pdx] = h.items[pdx], h.items[idx]
		}
		idx = pdx
	}
}

// Next returns the next element in the heap.
func (h *Heap[T]) Next() T {
	next := h.swapRemove(1)
	h.count--

	if h.count > 0 {
		idx := 1
		for h.childrenPresent(idx) {
			cdx := h.smallestChildIdx(idx)
			if !h.comparator(h.items[idx], h.items[cdx]) {
				// swap
				h.items[idx], h.items[cdx] = h.items[cdx], h.items[idx]
			}
			idx = cdx
		}
	}

	return next
}

// HasNext returns true if the heap has more elements.
func (h *Heap[T]) HasNext() bool {
	return h.count > 0
}

// Peek returns the top element of the heap without removing it.
func (h *Heap[T]) Peek() (T, bool) {
	if h.count == 0 {
		return *new(T), false
	}

	return h.items[1], true
}

func (h *Heap[T]) parentIdx(idx int) int {
	return idx / 2
}

func (h *Heap[T]) childrenPresent(idx int) bool {
	return h.leftChildIdx(idx) <= h.count
}

func (h *Heap[T]) leftChildIdx(idx int) int {
	return idx * 2
}

func (h *Heap[T]) rightChildIdx(idx int) int {
	return h.leftChildIdx(idx) + 1
}

func (h *Heap[T]) smallestChildIdx(idx int) int {
	if h.rightChildIdx(idx) > h.count {
		return h.leftChildIdx(idx)
	}

	ldx := h.leftChildIdx(idx)
	rdx := h.rightChildIdx(idx)

	if h.comparator(h.items[ldx], h.items[rdx]) {
		return ldx
	}

	return rdx
}

func (h *Heap[T]) swapRemove(idx int) T {
	if idx >= len(h.items) {
		errorStr := fmt.Sprintf("index (is %d) should be < len (is %d)", idx, h.count)
		panic(errorStr)
	}

	// get the last item
	value := h.items[idx]

	// swap the last item with the current item
	h.items[idx] = h.items[h.count]

	// remove the last item
	h.items = h.items[:h.count]

	return value
}
