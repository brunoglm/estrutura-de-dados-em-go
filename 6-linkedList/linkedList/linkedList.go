package linkedlist

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	count    int
	head     *Node[T]
	equalsFn func(a, b T) bool
}

func newNode[T any](element T) *Node[T] {
	return &Node[T]{
		data: element,
	}
}

func NewLinkedList[T any](equalsFn func(a, b T) bool) *LinkedList[T] {
	if equalsFn == nil {
		equalsFn = equalsDefaultFn[T]
	}
	return &LinkedList[T]{
		equalsFn: equalsFn,
	}
}

func (l *LinkedList[T]) Push(element T) {
	node := newNode(element)

	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}

	l.count++
}

func (l *LinkedList[T]) IndexOf(element T) int {
	current := l.head
	for i := 0; i < l.count && current != nil; i++ {
		if l.equalsFn(element, current.data) {
			return i
		}
		current = current.next
	}
	return -1
}

func (l *LinkedList[T]) Remove(element T) T {
	i := l.IndexOf(element)
	return l.RemovedAt(i)
}

func (l *LinkedList[T]) RemovedAt(index int) T {
	if index <= 0 || index >= l.count {
		var zero T
		return zero
	}

	current := l.head

	if index == 0 {
		l.head = current.next
	} else {
		previous := l.getElementAt(index - 1)
		current = previous.next
		previous.next = current.next
	}

	l.count--
	return current.data
}

func (l *LinkedList[T]) Insert(element T, index int) bool {
	if index <= 0 || index >= l.count {
		return false
	}

	node := newNode(element)

	if index == 0 {
		node.next = l.head
		l.head = node
	} else {
		previous := l.getElementAt(index - 1)
		node.next = previous.next
		previous.next = node
	}

	l.count++

	return true
}

func (l *LinkedList[T]) getElementAt(index int) *Node[T] {
	if index <= 0 || index >= l.count {
		return nil
	}

	node := l.head

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (l *LinkedList[T]) Size() int {
	return l.count
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *LinkedList[T]) GetHead() *Node[T] {
	return l.head
}

func (l *LinkedList[T]) ToString() string {
	if l.head == nil {
		return ""
	}
	objString := fmt.Sprintf("%v", l.head.data)
	current := l.head.next
	for i := 1; i < l.Size() && current != nil; i++ {
		objString += fmt.Sprintf(",%v", l.head.data)
		current = current.next
	}
	return objString
}

func equalsDefaultFn[T any](a, b T) bool {
	return fmt.Sprintln(a) == fmt.Sprintln(b)
}
