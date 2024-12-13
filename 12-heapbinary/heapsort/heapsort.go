package heapsort

import (
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](array []T) []T {
	if array == nil || len(array) <= 1 {
		return array
	}

	heapSize := len(array)

	buildMaxHeap(array)

	for heapSize > 1 {
		heapSize--
		swap(array, 0, heapSize)
		siftDown(array, 0, heapSize)
	}

	return array
}

func buildMaxHeap[T constraints.Ordered](array []T) []T {
	indexInitial := len(array) / 2

	for i := indexInitial; i >= 0; i-- {
		siftDown(array, i, len(array))
	}

	return array
}

func siftDown[T constraints.Ordered](array []T, index int, size int) {
	element := index
	left := getLeftIndex[T](index)
	right := getRightIndex[T](index)

	if left < size &&
		defaultCompare(array[element], array[left]) == -1 {
		element = left
	}

	if right < size &&
		defaultCompare(array[element], array[right]) == -1 {
		element = right
	}

	if index != element {
		swap(array, index, element)
		siftDown(array, element, size)
	}
}

func swap[T constraints.Ordered](array []T, a, b int) {
	parent := array[a]
	array[a] = array[b]
	array[b] = parent
}

func getRightIndex[T constraints.Ordered](index int) int {
	return (2 * index) + 2
}

func getLeftIndex[T constraints.Ordered](index int) int {
	return (2 * index) + 1
}

func defaultCompare[T constraints.Ordered](item1, item2 T) int {
	if item1 == item2 {
		return 0
	}

	if item1 < item2 {
		return -1
	}

	return 1
}
