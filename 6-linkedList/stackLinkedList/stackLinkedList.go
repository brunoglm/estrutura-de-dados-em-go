package stackLinkedList

import "fmt"

type Node[T any] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

type StackLinkedList[T any] struct {
	count    int
	head     *Node[T]
	tail     *Node[T]
	equalsFn func(a, b T) bool
}

func newNode[T any](element T) *Node[T] {
	return &Node[T]{
		data: element,
	}
}

func NewDoublyLinkedList[T any](equalsFn func(a, b T) bool) *StackLinkedList[T] {
	if equalsFn == nil {
		equalsFn = equalsDefaultFn[T]
	}
	return &StackLinkedList[T]{
		equalsFn: equalsFn,
	}
}

func (l *StackLinkedList[T]) IndexOf(element T) int {
	current := l.head
	for i := 0; i < l.count && current != nil; i++ {
		if l.equalsFn(element, current.data) {
			return i
		}
		current = current.next
	}
	return -1
}

func (l *StackLinkedList[T]) Remove(element T) T {
	i := l.IndexOf(element)

	if i < 0 {
		var zero T
		return zero
	}

	return l.RemovedAt(i)
}

func (l *StackLinkedList[T]) RemovedAt(index int) T {
	if index < 0 || index > l.count {
		var zero T
		return zero
	}

	current := l.head

	if index == 0 {
		l.head = current.next
		if l.Size() == 1 {
			l.tail = nil
		} else {
			l.head.prev = nil
		}
	} else if index == l.count-1 {
		current = l.tail
		l.tail = current.prev
		l.tail.next = nil
	} else {
		current = l.getElementAt(index)
		previous := current.prev
		nextNode := current.next
		previous.next = nextNode
		nextNode.prev = previous
	}

	l.count--
	return current.data
}

func (l *StackLinkedList[T]) Push(element T) {
	node := newNode(element)

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}

	l.count++
}

func (l *StackLinkedList[T]) Pop(element T) T {
	if l.IsEmpty() {
		var zero T
		return zero
	}

	return l.RemovedAt(l.Size() - 1)
}

func (l *StackLinkedList[T]) Peek(element T) T {
	if l.IsEmpty() {
		var zero T
		return zero
	}

	lastItem := l.getElementAt(l.Size() - 1)

	return lastItem.data
}

func (l *StackLinkedList[T]) Insert(element T, index int) bool {
	if index < 0 || index > l.count {
		return false
	}

	node := newNode(element)

	if index == 0 {
		if l.head == nil {
			l.head = node
			l.tail = node
		} else {
			current := l.head
			node.next = current
			current.prev = node
			l.head = node
		}
	} else if index == l.count {
		current := l.tail
		current.next = node
		node.prev = current
		l.tail = node
	} else {
		previous := l.getElementAt(index - 1)
		current := previous.next
		node.next = current
		node.prev = previous
		previous.next = node
		current.prev = node
	}

	l.count++
	return true
}

func (l *StackLinkedList[T]) Size() int {
	return l.count
}

func (l *StackLinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *StackLinkedList[T]) GetHead() T {
	return l.head.data
}

func (l *StackLinkedList[T]) GetTail() T {
	return l.tail.data
}

func (l *StackLinkedList[T]) ToString() string {
	if l.head == nil {
		return ""
	}
	objString := fmt.Sprintf("%v", l.head.data)
	current := l.head.next
	for i := 1; i < l.Size() && current != nil; i++ {
		objString += fmt.Sprintf(",%v", current.data)
		current = current.next
	}
	return objString
}

func (l *StackLinkedList[T]) getElementAt(index int) *Node[T] {
	if index < 0 || index > l.count {
		return nil
	}

	node := l.head

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func equalsDefaultFn[T any](a, b T) bool {
	return fmt.Sprintln(a) == fmt.Sprintln(b)
}
