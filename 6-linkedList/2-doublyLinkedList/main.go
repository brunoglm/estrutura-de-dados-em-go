package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoublyLinkedList) DeleteNode(value int) {
	if list.head != nil && list.head.data == value {
		list.head = list.head.next
		return
	}

	if list.tail != nil && list.tail.data == value {
		list.tail.prev.next = nil
		list.tail = list.tail.prev
		return
	}

	currentNode := list.head

	for currentNode != nil && currentNode.data != value {
		currentNode = currentNode.next
	}

	if currentNode != nil {
		currentNode.prev.next = currentNode.next
	}

	return
}

func (list *DoublyLinkedList) AddNode(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *DoublyLinkedList) DisplayForward() {
	current := list.head
	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (list *DoublyLinkedList) DisplayReverse() {
	current := list.tail
	fmt.Print("Linked list Reverse: ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.prev
	}
	fmt.Println()
}

func main() {
	list := DoublyLinkedList{}

	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(30)
	list.AddNode(40)

	list.DisplayForward()
	list.DisplayReverse()

	list.DeleteNode(20)

	list.DisplayForward()
}
