package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) DeleteNode(value int) {
	if list.head != nil && list.head.data == value {
		list.head = list.head.next
		return
	}

	currentNode := list.head
	var prevNode *Node

	for currentNode != nil && currentNode.data != value {
		prevNode = currentNode
		currentNode = currentNode.next
	}

	if currentNode != nil {
		prevNode.next = currentNode.next
	}

	return
}

func (list *LinkedList) AddNode(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Display() {
	current := list.head
	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(30)
	list.AddNode(40)

	list.Display()

	list.DeleteNode(20)

	list.Display()
}
