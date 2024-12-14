package stack

import "fmt"

type Stack[T any] struct {
	count int
	itens map[int]T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{itens: make(map[int]T)}
}

func (s *Stack[T]) Push(items ...T) {
	for _, item := range items {
		s.itens[s.count] = item
		s.count++
	}
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	item := s.itens[s.count-1]

	delete(s.itens, s.count-1)

	s.count--

	return item
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	return s.itens[s.count-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.count == 0
}

func (s *Stack[T]) Clear() {
	s.itens = make(map[int]T)
	s.count = 0
}

func (s *Stack[T]) ToString() string {
	values := make([]T, 0, len(s.itens))
	for _, v := range s.itens {
		values = append(values, v)
	}
	return fmt.Sprintf("%v", values)
}

func (s *Stack[T]) Size() int {
	return s.count
}
