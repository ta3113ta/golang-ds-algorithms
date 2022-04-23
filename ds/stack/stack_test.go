package stack

import "testing"

func TestPush(t *testing.T) {
	stack := New[uint8]()

	stack.Push(1)
	stack.Push(2)

	if len(*stack) != 2 {
		t.Errorf("Expected stack length to be 2, got %d", len(*stack))
	}
}

func TestPop(t *testing.T) {
	stack := New[uint8]()

	stack.Push(1)
	stack.Push(2)

	v, ok := stack.Pop()
	if !ok {
		t.Error("Expected Pop to return true")
	}

	if v != 2 {
		t.Errorf("Expected value to be 2, got %d", v)
	}
}

func TestPeek(t *testing.T) {
	stack := New[uint8]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	v, ok := stack.Peek()
	if !ok {
		t.Error("Expected Peek to return true")
	}

	if v != 3 {
		t.Errorf("Expected value to be 3, got %d", v)
	}
}
