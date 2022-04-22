// Thanks! https://github.com/gammazero/deque for the idea.
// This deque hasn't been implemented shrinkIfExcess() yet, so keep it simple.

package deque

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const minCapacity = 16

// Deque a double-ended queue implemented with a growable ring buffer.
type Deque[T any] struct {
	buf   []T
	head  int
	tail  int
	count int
}

// WithCapacity new a double-ended queue with the specified capacity.
func New[T any](size int) *Deque[T] {
	capacity := size
	buf := []T{}

	if capacity != 0 {
		bufSize := minCapacity
		for bufSize < capacity {
			bufSize <<= 1
		}
		buf = make([]T, bufSize)
	}

	return &Deque[T]{
		buf:  buf,
		head: 0,
		tail: 0,
	}
}

// Capacity returns the capacity of the deque.
func (d *Deque[T]) Capacity() int {
	return cap(d.buf)
}

// Len returns the length of the deque.
func (d *Deque[T]) Len() int {
	return d.count
}

// IsFull returns true if the deque is full.
func (d *Deque[T]) IsFull() bool {
	return d.count == len(d.buf)
}

// Grow the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full
func (d *Deque[T]) Grow() {
	newBuf := make([]T, d.count<<1)

	if d.tail > d.head {
		copy(newBuf, d.buf[d.head:d.tail])
	} else {
		n := copy(newBuf, d.buf[d.head:])
		copy(newBuf[n:], d.buf[:d.tail])
	}

	d.head = 0
	d.tail = d.count
	d.buf = newBuf
}

// PushFront insert an element at the front of the deque.
func (d *Deque[T]) PushFront(value T) {
	if d.IsFull() {
		d.Grow()
	}

	// Calculate the head position.
	d.head = d.prev(d.head)
	d.buf[d.head] = value
	d.count++
}

// PushBack insert an element at the back of the deque.
func (d *Deque[T]) PushBack(value T) {
	if d.IsFull() {
		d.Grow()
	}

	d.buf[d.tail] = value

	// Calculate the tail position.
	d.tail = d.next(d.tail)
	d.count++
}

// PopFront remove and return the element at the front of the deque.
func (d *Deque[T]) PopFront() T {
	if d.Len() <= 0 {
		panic("deque is empty")
	}

	ret := d.buf[d.head]
	d.buf[d.head] = *new(T)

	// Calculate new head position.
	d.head = d.next(d.head)
	d.count--

	return ret
}

// PopBack remove and return the element at the back of the deque.
func (d *Deque[T]) PopBack() T {
	if d.Len() <= 0 {
		panic("deque is empty")
	}

	// Calculate new tail position.
	d.tail = d.prev(d.tail)

	// Remove value at tail.
	ret := d.buf[d.tail]
	d.buf[d.tail] = *new(T)
	d.count--

	return ret
}

// PeekFront returns the element at the front of the deque.
func (d *Deque[T]) PeekFront() T {
	if d.Len() <= 0 {
		panic("deque is empty")
	}

	return d.buf[d.head]
}

// PeekBack returns the last element of the deque.
func (q *Deque[T]) PeekBack() T {
	if q.Len() <= 0 {
		panic("deque is empty")
	}

	return q.buf[q.prev(q.tail)]
}

// Clear the deque.
func (d *Deque[T]) Clear() {
	// bitwise modulus
	modBits := len(d.buf) - 1
	for h := d.head; h != d.tail; h = (h + 1) & modBits {
		d.buf[h] = *new(T)
	}

	d.head = 0
	d.tail = 0
	d.count = 0
}

// prev return the previous buffer position wrapping around buffer.
func (d *Deque[T]) prev(i int) int {
	return (i - 1) & (len(d.buf) - 1) // bitwise modulus
}

// next return the next buffer position wrapping around buffer.
func (d *Deque[T]) next(i int) int {
	return (i + 1) & (len(d.buf) - 1) // bitwise modulus
}
