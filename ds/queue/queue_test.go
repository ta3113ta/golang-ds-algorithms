package queue

import "testing"

func TestEnqueue(t *testing.T) {
	queue := New[uint8]()
	queue.Enqueue(64)

	if queue.IsEmpty() {
		t.Fatalf("expected IsEmpty = true, but get %v", queue.IsEmpty())
	}
}

func TestDequeue(t *testing.T) {
	queue := New[uint8]()
	queue.Enqueue(32)
	queue.Enqueue(64)

	retrievedDequeue, _ := queue.Dequeue()

	if retrievedDequeue != 32 {
		t.Fatalf("expected 32, but get %v", retrievedDequeue)
	}
}

func TestPeekFront(t *testing.T) {
	queue := New[uint8]()
	queue.Enqueue(8)
	queue.Enqueue(16)

	retrievedPeek, _ := queue.Peek()

	if retrievedPeek != 8 {
		t.Fatalf("expected 8, but get %v", retrievedPeek)
	}
}

func TestSize(t *testing.T) {
	queue := New[uint8]()
	queue.Enqueue(8)
	queue.Enqueue(16)

	if queue.Len() != 2 {
		t.Fatalf("expected 2, but get %v", queue.Len())
	}
}
