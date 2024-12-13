package main

import (
	"fmt"
	"labheap/minheap"
)

func main() {
	h := minheap.NewMinHeap[int]()
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)
	h.Insert(1)

	fmt.Println("Heap size: ", h.Size())
	fmt.Println("Heap is empty: ", h.IsEmpty())
	fmt.Println("Heap min value: ", h.FindMinimum())
}
