package linkedlist

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type SortedLinkedList[T constraints.Ordered] struct {
	count    int
	head     *Node[T]
	equalsFn func(a, b T) int
}

func newNode[T any](element T) *Node[T] {
	return &Node[T]{
		data: element,
	}
}

func NewLinkedList[T constraints.Ordered](equalsFn func(a, b T) int) *SortedLinkedList[T] {
	if equalsFn == nil {
		equalsFn = equalsDefaultFn[T]
	}
	return &SortedLinkedList[T]{
		equalsFn: equalsFn,
	}
}

func (l *SortedLinkedList[T]) Push(element T) {
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

func (l *SortedLinkedList[T]) IndexOf(element T) int {
	current := l.head
	for i := 0; i < l.count && current != nil; i++ {
		if l.equalsFn(element, current.data) == 0 {
			return i
		}
		current = current.next
	}
	return -1
}

func (l *SortedLinkedList[T]) Remove(element T) T {
	i := l.IndexOf(element)

	if i < 0 {
		var zero T
		return zero
	}

	return l.RemovedAt(i)
}

func (l *SortedLinkedList[T]) RemovedAt(index int) T {
	if index < 0 || index > l.count {
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

func (l *SortedLinkedList[T]) Insert(element T) bool {
	if l.IsEmpty() {
		return l.insert(element, 0)
	}

	pos := l.getIndexNextSortedElement(element)
	return l.insert(element, pos)
}

func (l *SortedLinkedList[T]) getIndexNextSortedElement(element T) int {
	current := l.head
	i := 0

	for ; i < l.Size() && current != nil; i++ {
		comp := l.equalsFn(element, current.data)
		if comp == -1 {
			return i
		}
		current = current.next
	}

	return i
}

func (l *SortedLinkedList[T]) insert(element T, index int) bool {
	if index < 0 || index > l.count {
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

func (l *SortedLinkedList[T]) getElementAt(index int) *Node[T] {
	if index < 0 || index > l.count {
		return nil
	}

	node := l.head

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (l *SortedLinkedList[T]) Size() int {
	return l.count
}

func (l *SortedLinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *SortedLinkedList[T]) GetHead() *Node[T] {
	return l.head
}

func (l *SortedLinkedList[T]) ToString() string {
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

func equalsDefaultFn[T constraints.Ordered](a, b T) int {
	if a == b {
		return 0
	}

	if a < b {
		return -1
	}

	return 1
}
