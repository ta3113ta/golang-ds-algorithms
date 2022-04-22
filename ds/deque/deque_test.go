package deque

import (
	"testing"
)

func TestCapacity(t *testing.T) {
	deque := New[uint8](10)

	if deque.Capacity() != 16 {
		t.Fatalf("expected 10, but get %v", deque.Capacity())
	}
}

func TestEmtpyLen(t *testing.T) {
	deque := New[uint8](10)

	if deque.Len() != 0 {
		t.Fatalf("expected 0, but get %v", deque.Len())
	}
}

func TestNotFull(t *testing.T) {
	deque := New[uint8](10)

	if deque.IsFull() {
		t.Fatalf("expected false, but get true")
	}
}

func TestGrow(t *testing.T) {
	deque := New[uint8](5)

	for i := 0; i <= 10; i++ {
		deque.PushFront(uint8(i))
	}

	if deque.Capacity() != 16 {
		t.Fatalf("expected 20, but get %v", deque.Capacity())
	}
}

func TestPushFront(t *testing.T) {
	deque := New[uint8](10)

	deque.PushFront(1)
	deque.PushFront(2)
	deque.PushFront(3)

	if deque.Len() != 3 {
		t.Fatalf("expected 3, but get %v", deque.Len())
	}

	if deque.PeekFront() != uint8(3) {
		t.Fatalf("expected 3, but get %v", deque.PeekFront())
	}
}

func TestPushBack(t *testing.T) {
	deque := New[uint8](10)

	deque.PushBack(1)
	deque.PushBack(2)

	if deque.Len() != 2 {
		t.Fatalf("expected 2, but get %v", deque.Len())
	}

	if deque.PeekBack() != uint8(2) {
		t.Fatalf("expected 2, but get %v", deque.PeekBack())
	}
}

func TestPopFront(t *testing.T) {
	deque := New[uint8](10)

	deque.PushBack(1)
	deque.PushBack(2)

	value := deque.PopFront()
	if value != 1 {
		t.Fatalf("expected 1, but get %v", value)
	}
}

func TestPopBack(t *testing.T) {
	deque := New[uint8](3)

	deque.PushBack(1)
	deque.PushBack(2)

	value := deque.PopBack()
	if value != 2 {
		t.Fatalf("expected 2, but get %v", value)
	}
}

func TestPeekFront(t *testing.T) {
	deque := New[uint8](3)

	deque.PushBack(1)
	deque.PushBack(2)

	if deque.PeekFront() != 1 {
		t.Fatalf("expected 1, but get %v", deque.PeekFront())
	}
}

func TestPeekBack(t *testing.T) {
	deque := New[uint](3)

	deque.PushFront(3)
	deque.PushFront(2)

	if deque.PeekBack() != 3 {
		t.Fatalf("expected 3, but get %v", deque.PeekBack())
	}
}

func TestClear(t *testing.T) {
	deque := New[uint8](3)

	deque.PushBack(1)
	deque.PushBack(2)

	deque.Clear()

	if deque.Len() != 0 {
		t.Fatalf("expected 0, but get %v", deque.Len())
	}
}
