package trie

type Node struct {
	Children map[rune]*Node
	IsEnd    bool
	Value    any
}

func NewNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
	}
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewNode(),
	}
}

func (t *Trie) Insert(key string) {
	node := t.Root
	for _, c := range key {
		if node.Children[c] == nil {
			node.Children[c] = NewNode()
		}
		node = node.Children[c]
	}
	node.IsEnd = true
}

func (t *Trie) Search(key string) bool {
	node := t.Root
	for _, c := range key {
		if node.Children[c] == nil {
			return false
		}
		node = node.Children[c]
	}
	return node.IsEnd
}

func (t *Trie) Delete(key string) {
	t.delete(t.Root, key, 0)
}

func (t *Trie) delete(node *Node, key string, depth int) {
	if node == nil {
		return
	}

	if depth == len(key) {
		node.IsEnd = false
		return
	}

	t.delete(node.Children[rune(key[depth])], key, depth+1)
}
