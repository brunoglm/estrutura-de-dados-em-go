package stack

import (
	"fmt"
)

type Stack[T any] struct {
	itens []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{itens: make([]T, 0)}
}

func (s *Stack[T]) Push(items ...T) {
	s.itens = append(s.itens, items...)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	lastIndex, item := s.getLastItem()

	s.itens = s.itens[:lastIndex]

	return item
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	_, item := s.getLastItem()

	return item
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.itens) == 0
}

func (s *Stack[T]) Clear() {
	s.itens = make([]T, 0)
}

func (s *Stack[T]) Size() int {
	return len(s.itens)
}

func (s *Stack[T]) ToString() string {
	return fmt.Sprintf("%v", s.itens)
}

func (s *Stack[T]) getLastItem() (int, T) {
	lastIndex := len(s.itens) - 1
	item := s.itens[lastIndex]

	return lastIndex, item
}
