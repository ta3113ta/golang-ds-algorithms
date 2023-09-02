package trie

import "testing"

func TestInsert(t *testing.T) {
	t.Parallel()

	trie := NewTrie()
	trie.Insert("hello")

	if trie.Root.
		Children['h'].
		Children['e'].
		Children['l'].
		Children['l'].
		Children['o'].
		IsEnd != true {
		t.Error("Expected true")
	}
}

func TestSearch(t *testing.T) {
	t.Parallel()

	trie := NewTrie()
	trie.Insert("hello")

	if trie.Search("hello") != true {
		t.Error("Expected true")
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()

	trie := NewTrie()
	trie.Insert("hello")
	trie.Delete("hello")

	if trie.Search("hello") != false {
		t.Error("Expected false")
	}
}

