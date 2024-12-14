package stack

import "fmt"

type StackMap[T any] struct {
	count int
	itens map[int]T
}

func NewStackMap[T any]() *StackMap[T] {
	return &StackMap[T]{itens: make(map[int]T)}
}

func (s *StackMap[T]) Push(items ...T) {
	for _, item := range items {
		s.itens[s.count] = item
		s.count++
	}
}

func (s *StackMap[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	item := s.itens[s.count-1]

	delete(s.itens, s.count-1)

	s.count--

	return item
}

func (s *StackMap[T]) Peek() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	return s.itens[s.count-1]
}

func (s *StackMap[T]) IsEmpty() bool {
	return s.count == 0
}

func (s *StackMap[T]) Clear() {
	s.itens = make(map[int]T)
	s.count = 0
}

func (s *StackMap[T]) ToString() string {
	values := make([]T, 0, len(s.itens))
	for _, v := range s.itens {
		values = append(values, v)
	}
	return fmt.Sprintf("%v", values)
}

func (s *StackMap[T]) Size() int {
	return s.count
}
