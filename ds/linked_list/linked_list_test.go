package linkedlist

import (
	"fmt"
	"testing"
)

func TestInsertAtTailWorks(t *testing.T) {
	list := NewLinkedList[int32]()
	secondVal := int32(2)
	list.InsertAtTail(1)
	list.InsertAtTail(secondVal)

	val, ok := list.Get(1)
	if !ok {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}

	if val != secondVal {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}
}

func TestInsertAtHeadWorks(t *testing.T) {
	list := NewLinkedList[int32]()
	secondVal := int32(2)
	list.InsertAtHead(1)
	list.InsertAtHead(secondVal)

	val, ok := list.Get(0)
	if !ok {
		t.Errorf("Expected to get nil, got %v", val)
	}

	if val != secondVal {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}
}

func TestInsertAtIthCanAddToTail(t *testing.T) {
	list := NewLinkedList[int32]()
	secondVal := int32(2)
	list.InsertAtIth(0, 0)
	list.InsertAtIth(1, secondVal)

	val, ok := list.Get(1)
	if !ok {
		t.Errorf("Expected to get nil, got %v", val)
	}

	if val != secondVal {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}
}

func TestInsertAtIthCanAddToHead(t *testing.T) {
	list := NewLinkedList[int32]()
	secondVal := int32(2)
	list.InsertAtIth(0, 1)
	list.InsertAtIth(0, secondVal)

	val, ok := list.Get(0)
	if !ok {
		t.Errorf("Expected to get nil, got %v", val)
	}

	if val != secondVal {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}
}

func TestInsertAtIthCanAddToMiddle(t *testing.T) {
	list := NewLinkedList[int32]()
	secondVal := int32(2)
	thirdVal := int32(3)
	list.InsertAtIth(0, 1)
	list.InsertAtIth(1, secondVal)
	list.InsertAtIth(1, thirdVal)

	fmt.Println(list.head.next.val)

	val, ok := list.Get(1)
	if !ok {
		t.Errorf("Expected to get nil, got %v", val)
	}

	if val != thirdVal {
		t.Errorf("Expected to get %v, got %v", thirdVal, val)
	}

	val, ok = list.Get(2)
	if !ok {
		t.Errorf("Expected to get nil, got %v", val)
	}

	if val != secondVal {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}
}

func TestInsertAtIthAndDeleteIthWorkOverManyIterations(t *testing.T) {
	list := NewLinkedList[int32]()
	for i := 0; i < 100; i++ {
		list.InsertAtIth(i, int32(i))
	}

	// Pop even numbers to 50
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			list.DeleteIth(i)
		}
	}

	if list.length != 75 {
		t.Errorf("Expected length to be 75, got %v", list.length)
	}

	// Insert even numbers back
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			list.InsertAtIth(i, int32(i))
		}
	}

	if list.length != 100 {
		t.Errorf("Expected length to be 100, got %v", list.length)
	}

	// Ensure numbers were adderd back and we're able to traverse nodes
	val, ok := list.Get(78)
	if !ok {
		t.Errorf("Expected to get nil, got %v", val)
	}

	if val != int32(78) {
		t.Errorf("Expected to get %v, got %v", int32(78), val)
	}
}

func TestDeleteTailWorks(t *testing.T) {
	list := NewLinkedList[int32]()
	firstVal := int32(1)
	secondVal := int32(2)
	list.InsertAtTail(firstVal)
	list.InsertAtTail(secondVal)

	val, ok := list.DeleteTail()
	if !ok {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}
	if val != secondVal {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}

	val, ok = list.Get(0)
	if !ok {
		t.Errorf("Expected to get nil, got %v", val)
	}
	if val != firstVal {
		t.Errorf("Expected to get %v, got %v", firstVal, val)
	}
}

func TestDeleteHeadWorks(t *testing.T) {
	list := NewLinkedList[int32]()
	firstVal := int32(1)
	secondVal := int32(2)
	list.InsertAtTail(firstVal)
	list.InsertAtTail(secondVal)

	val, ok := list.DeleteHead()
	if !ok {
		t.Errorf("Expected to get %v, got %v", firstVal, val)
	}
	if val != firstVal {
		t.Errorf("Expected to get %v, got %v", firstVal, val)
	}

	val, ok = list.Get(0)
	if !ok {
		t.Errorf("Expected to get nil, got %v", val)
	}
	if val != secondVal {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}
}

func TestDeleteIthCanDeleteAtTail(t *testing.T) {
	list := NewLinkedList[int32]()
	firstVal := int32(1)
	secondVal := int32(2)
	list.InsertAtTail(firstVal)
	list.InsertAtTail(secondVal)

	val, ok := list.DeleteIth(1)
	if !ok {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}
	if val != secondVal {
		t.Errorf("Expected to get %v, got %v", secondVal, val)
	}
}