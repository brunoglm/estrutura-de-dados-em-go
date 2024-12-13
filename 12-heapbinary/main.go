package main

import (
	"fmt"
	"labheap/heapsort"
)

func main() {
	// h := minheap.NewMinHeap[int]()
	// h.Insert(2)
	// h.Insert(3)
	// h.Insert(4)
	// h.Insert(5)
	// h.Insert(1)

	// fmt.Println("Heap size: ", h.Size())
	// fmt.Println("Heap is empty: ", h.IsEmpty())
	// fmt.Println("Heap min value: ", h.FindMinimum())

	array := []int{7, 6, 3, 5, 4, 1, 2}
	fmt.Println("Before sorting: ", array)
	arraySorted := heapsort.Sort(array)
	fmt.Println("After sorting: ", arraySorted)
}
