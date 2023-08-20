package set

import "testing"

func TestEmptySet(t *testing.T) {
	set := NewSet[int32]()
	if !set.IsEmpty() {
		t.Error("Set should be empty")
	}
}

func TestInsertSet(t *testing.T) {
	set := NewSet[int32]()
	set.Insert(2)
	set.Insert(3)

	if set.IsEmpty() {
		t.Error("Set should not empty")
	}

	if set.Len() != 2 {
		t.Error("length should be 2")
	}
}

func TestSetContainsValue(t *testing.T) {
	set := NewSet[int32]()
	set.Insert(30)
	set.Insert(2)

	if !set.Contains(30) {
		t.Error("Set should contain 30")
	}
}

func TestRemoveSetValue(t *testing.T) {
	set := NewSet[int32]()
	set.Insert(30)
	set.Insert(2)

	if set.Len() != 2 {
		t.Error("Length of set should be 2")
	}

	set.Remove(2)
	if set.Len() != 1 {
		t.Error("Length of set should be 1")
	}

	if set.Contains(2) {
		t.Error("Set should not contain 2")
	}
}

func TestClearSet(t *testing.T) {
	set := NewSet[int32]()
	set.Insert(30)
	set.Insert(32)

	if set.Len() != 2 {
		t.Error("Length of set should be 2")
	}

	set.Clear()
	if set.Len() != 0 {
		t.Error("Set should be empty")
	}
}
