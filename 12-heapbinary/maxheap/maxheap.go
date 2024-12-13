package maxheap

import "golang.org/x/exp/constraints"

type MaxHeap[T constraints.Ordered] struct {
	Heap []T
}

func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{}
}

func (h *MaxHeap[T]) GetLeftIndex(index int) int {
	return (2 * index) + 1
}

func (h *MaxHeap[T]) GetRightIndex(index int) int {
	return (2 * index) + 2
}

func (h *MaxHeap[T]) GetParentIndex(index int) int {
	if index == 0 {
		return 0
	}

	return (index - 1) / 2
}

func (h *MaxHeap[T]) Insert(value T) bool {
	var zero T
	if value == zero {
		return false
	}

	h.Heap = append(h.Heap, value)
	h.siftUp(h.Size() - 1)
	return true
}

func (h *MaxHeap[T]) Extract() T {
	if h.IsEmpty() {
		var zero T
		return zero
	}

	if h.Size() == 1 {
		return h.shift()
	}

	removedValue := h.shift()
	h.siftDown(0)
	return removedValue
}

func (h *MaxHeap[T]) shift() T {
	if h.IsEmpty() {
		var zero T
		return zero
	}

	firstValue := h.Heap[0]

	// Limpar a referência para o objeto na posição 0
	h.Heap[0] = *new(T)

	h.Heap = h.Heap[1:]

	return firstValue
}

func (h *MaxHeap[T]) Size() int {
	return len(h.Heap)
}

func (h *MaxHeap[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *MaxHeap[T]) FindMinimum() T {
	if h.IsEmpty() {
		var zero T
		return zero
	}

	return h.Heap[0]
}

func (h *MaxHeap[T]) siftUp(index int) {
	parent := h.GetParentIndex(index)

	for index > 0 && h.defaultCompare(h.Heap[parent], h.Heap[index]) == -1 {
		h.swap(parent, index)
		index = parent
		parent = h.GetParentIndex(index)
	}
}

func (h *MaxHeap[T]) siftDown(index int) {
	element := index
	left := h.GetLeftIndex(index)
	right := h.GetRightIndex(index)
	size := h.Size()

	if left < size &&
		h.defaultCompare(h.Heap[element], h.Heap[left]) == -1 {
		element = left
	}

	if right < size &&
		h.defaultCompare(h.Heap[element], h.Heap[right]) == -1 {
		element = right
	}
	if index != element {
		h.swap(index, element)
		h.siftDown(element)
	}
}

func (h *MaxHeap[T]) swap(a, b int) {
	parent := h.Heap[a]
	h.Heap[a] = h.Heap[b]
	h.Heap[b] = parent
}

func (h *MaxHeap[T]) defaultCompare(item1, item2 T) int {
	if item1 == item2 {
		return 0
	}

	if item1 < item2 {
		return -1
	}

	return 1
}
