package binary_heap

import "testing"

func TestEmptyHeap(t *testing.T) {
	heap := NewMax[int32]()
	if heap.HasNext() {
		t.Error("Heap should be empty")
	}
}

func TestMinHeap(t *testing.T) {
	heap := NewMin[int32]()
	heap.Add(4)
	heap.Add(2)
	heap.Add(9)
	heap.Add(11)

	if heap.Len() != 4 {
		t.Error("Heap should have 4 elements")
	}

	expected := []int32{2, 4, 9, 11}
	for _, v := range expected {
		if heap.Next() != v {
			t.Errorf("Heap should have %d as next element", v)
		}
	}

	heap.Add(1)
	if !heap.HasNext() {
		t.Error("Heap should have next element")
	}
}

func TestMaxHeap(t *testing.T) {
	heap := NewMax[int32]()
	heap.Add(4)
	heap.Add(2)
	heap.Add(9)
	heap.Add(11)

	if heap.Len() != 4 {
		t.Error("Heap should have 4 elements")
	}

	expected := []int32{11, 9, 4, 2}
	for _, v := range expected {
		if heap.Next() != v {
			t.Errorf("Heap should have %d as next element", v)
		}
	}

	heap.Add(1)
	if heap.Next() != 1 {
		t.Error("Heap should have 1 as next element")
	}
}

func TestPeek(t *testing.T) {
	heap := NewMax[int32]()
	heap.Add(4)
	heap.Add(2)
	heap.Add(9)
	heap.Add(11)

	value, _ := heap.Peek()
	if value != 11 {
		t.Errorf("Heap should have 11 as next element")
	}
}

func TestKeyHeap(t *testing.T) {
	type Point struct {
		x int32
		y int32
	}

	fn := func(a, b Point) bool {
		return a.x < b.x
	}

	heap := New(fn)
	heap.Add(Point{x: 1, y: 5})
	heap.Add(Point{x: 3, y: 10})
	heap.Add(Point{x: -2, y: 4})

	if heap.Len() != 3 {
		t.Error("Heap should have 3 elements")
	}

	expected := []Point{{x: -2, y: 4}, {x: 1, y: 5}, {x: 3, y: 10}}
	for _, v := range expected {
		if heap.Next() != v {
			t.Errorf("Heap should have %d as next element", v)
		}
	}
}
