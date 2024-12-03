package main

import (
	"fmt"
	"linkedlist/doublyLinkedList"
)

func main() {
	fmt.Println("Testando doublyLinkedList")

	doublyLinkedList := doublyLinkedList.NewDoublyLinkedList[string](nil)

	fmt.Println("IsEmpty: ", doublyLinkedList.IsEmpty())
	fmt.Println("Insert Bruno 0: ", doublyLinkedList.Insert("Bruno", 0))
	fmt.Println("Size: ", doublyLinkedList.Size())
	fmt.Println("GetHead: ", doublyLinkedList.GetHead())
	fmt.Println("Insert Gabriel tail: ", doublyLinkedList.Insert("Gabriel", 1))
	fmt.Println("Insert Lemes tail: ", doublyLinkedList.Insert("Lemes", 2))
	fmt.Println("Insert De tail: ", doublyLinkedList.Insert("De", 3))
	fmt.Println("Insert Magalhães tail: ", doublyLinkedList.Insert("Magalhães", 4))
	fmt.Println("Size: ", doublyLinkedList.Size())
	fmt.Println("GetHead: ", doublyLinkedList.GetHead())
	fmt.Println("GetTail: ", doublyLinkedList.GetTail())

	fmt.Println("toString: ", doublyLinkedList.ToString())

	fmt.Println("Insert vixi tail: ", doublyLinkedList.Insert("vixi", 2))
	fmt.Println("GetHead: ", doublyLinkedList.GetHead())
	fmt.Println("GetTail: ", doublyLinkedList.GetTail())

	fmt.Println("Insert toString: ", doublyLinkedList.ToString())
}
