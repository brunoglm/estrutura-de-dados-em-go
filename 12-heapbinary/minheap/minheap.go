package minheap

import "golang.org/x/exp/constraints"

type MinHeap[T constraints.Ordered] struct {
	Heap []T
}

func NewMinHeap[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{}
}

func (h *MinHeap[T]) GetLeftIndex(index int) int {
	return (2 * index) + 1
}

func (h *MinHeap[T]) GetRightIndex(index int) int {
	return (2 * index) + 2
}

func (h *MinHeap[T]) GetParentIndex(index int) int {
	if index == 0 {
		return 0
	}

	return (index - 1) / 2
}

func (h *MinHeap[T]) Insert(value T) bool {
	var zero T
	if value == zero {
		return false
	}

	h.Heap = append(h.Heap, value)
	h.siftUp(h.Size() - 1)
	return true
}

// func (h *MinHeap[T]) Extract(value T) T {

// }

// func (h *MinHeap[T]) FindMinimum() T {

// }

func (h *MinHeap[T]) Size() int {
	return len(h.Heap)
}

func (h *MinHeap[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *MinHeap[T]) FindMinimum() T {
	if h.IsEmpty() {
		var zero T
		return zero
	}

	return h.Heap[0]
}

func (h *MinHeap[T]) siftUp(index int) {
	parent := h.GetParentIndex(index)

	for index > 0 && h.defaultCompare(h.Heap[parent], h.Heap[index]) == 1 {
		h.swap(h.Heap, parent, index)
		index = parent
		parent = h.GetParentIndex(index)
	}
}

func (h *MinHeap[T]) swap(heap []T, parentIndex, childIndex int) {
	parent := heap[parentIndex]
	heap[parentIndex] = heap[childIndex]
	heap[childIndex] = parent
}

func (h *MinHeap[T]) defaultCompare(item1, item2 T) int {
	if item1 == item2 {
		return 0
	}

	if item1 < item2 {
		return -1
	}

	return 1
}
