package linkedlist

type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

func NewNode[T any](t T) *Node[T] {
	return &Node[T]{val: t}
}

type LinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{length: 0}
}

// InsertAtHead inserts a value at the head of the list.
func (l *LinkedList[T]) InsertAtHead(obj T) {
	node := NewNode(obj)
	node.next = l.head
	node.prev = nil

	nodePtr := node
	if l.head != nil {
		l.head.prev = nodePtr
	} else {
		l.tail = nodePtr
	}

	l.head = nodePtr
	l.length++
}

// InsertAtTail inserts a value at the tail of the list.
func (l *LinkedList[T]) InsertAtTail(obj T) {
	node := NewNode(obj)
	node.next = nil
	node.prev = l.tail

	nodePtr := node
	if l.tail != nil {
		l.tail.next = nodePtr
	} else {
		l.head = nodePtr
	}

	l.tail = nodePtr
	l.length++
}

// InsertAtIth inserts a value at the ith position of the list.
func (l *LinkedList[T]) InsertAtIth(idx int, obj T) {
	if l.length < idx {
		panic("Index out of bounds")
	}

	if idx == 0 || l.head == nil {
		l.InsertAtHead(obj)
		return
	}

	if idx == l.length {
		l.InsertAtTail(obj)
		return
	}

	if l.head != nil {
		temp := l.head
		for i := 0; i < idx; i++ {
			if temp.next == nil {
				panic("Index out of bounds")
			}
			temp = temp.next
		}

		node := NewNode(obj)
		node.prev = temp.prev
		node.next = temp

		if temp.prev != nil {
			nodePtr := node
			temp.prev.next = nodePtr
			l.length++
		}
	}
}

// DeleteHead deletes the head of the list.
func (l *LinkedList[T]) DeleteHead() (T, bool) {
	if l.head == nil {
		return *new(T), false
	}

	oldHead := l.head
	if l.head.next != nil {
		l.head.next.prev = nil
	} else {
		l.tail = nil
	}

	l.head = oldHead.next
	l.length--
	return oldHead.val, true
}

// DeleteTail deletes the tail of the list.
func (l *LinkedList[T]) DeleteTail() (T, bool) {
	if l.tail == nil {
		return *new(T), false
	}

	oldTail := l.tail
	if l.tail.prev != nil {
		l.tail.prev.next = nil
	} else {
		l.head = nil
	}

	l.tail = oldTail.prev
	l.length--
	return oldTail.val, true
}

// DeleteIth deletes the ith element of the list.
func (l *LinkedList[T]) DeleteIth(idx int) (T, bool) {
	if l.length < idx {
		panic("Index out of bounds")
	}

	if idx == 0 || l.head == nil {
		return l.DeleteHead()
	}

	if idx == l.length {
		return l.DeleteTail()
	}

	if l.head != nil {
		temp := l.head
		for i := 0; i < idx; i++ {
			if temp.next == nil {
				panic("Index out of bounds")
			} else {
				temp = temp.next
			}
		}

		oldIth := temp
		if oldIth.prev != nil {
			oldIth.prev.next = oldIth.next
		} 

		if oldIth.next != nil {
			oldIth.next.prev = oldIth.prev
		}

		l.length--
		return oldIth.val, true
	}

	return *new(T), false
}

// Get returns the ith element of the list.
func (l *LinkedList[T]) Get(idx int) (T, bool) {
	return l.getIthNode(l.head, idx)
}

func (l *LinkedList[T]) getIthNode(node *Node[T], idx int) (T, bool) {
	if node == nil {
		return *new(T), false
	}

	if idx == 0 {
		return node.val, true
	}

	return l.getIthNode(node.next, idx-1)
}
