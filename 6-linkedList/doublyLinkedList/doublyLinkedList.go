package doublyLinkedList

import "fmt"

type Node[T any] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

type DoublyLinkedList[T any] struct {
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

func NewDoublyLinkedList[T any](equalsFn func(a, b T) bool) *DoublyLinkedList[T] {
	if equalsFn == nil {
		equalsFn = equalsDefaultFn[T]
	}
	return &DoublyLinkedList[T]{
		equalsFn: equalsFn,
	}
}

func (l *DoublyLinkedList[T]) IndexOf(element T) int {
	current := l.head
	for i := 0; i < l.count && current != nil; i++ {
		if l.equalsFn(element, current.data) {
			return i
		}
		current = current.next
	}
	return -1
}

func (l *DoublyLinkedList[T]) Remove(element T) T {
	i := l.IndexOf(element)
	return l.RemovedAt(i)
}

func (l *DoublyLinkedList[T]) RemovedAt(index int) T {
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

func (l *DoublyLinkedList[T]) Push(element T) {
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

func (l *DoublyLinkedList[T]) Insert(element T, index int) bool {
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

func (l *DoublyLinkedList[T]) Size() int {
	return l.count
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *DoublyLinkedList[T]) GetHead() T {
	return l.head.data
}

func (l *DoublyLinkedList[T]) GetTail() T {
	return l.tail.data
}

func (l *DoublyLinkedList[T]) ToString() string {
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

func (l *DoublyLinkedList[T]) getElementAt(index int) *Node[T] {
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
