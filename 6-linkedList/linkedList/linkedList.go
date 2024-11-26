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

func equalsDefaultFn[T any](a, b T) bool {
	return fmt.Sprintln(a) == fmt.Sprintln(b)
}
