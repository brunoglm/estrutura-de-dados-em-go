package queue

import "fmt"

type Queue[T any] struct {
	itens       map[int]T
	count       int
	lowestCount int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{itens: make(map[int]T)}
}

func (s *Queue[T]) Enqueue(items ...T) {
	for _, item := range items {
		s.itens[s.count] = item
		s.count++
	}
}

func (s *Queue[T]) Dequeue() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	item := s.itens[s.lowestCount]

	delete(s.itens, s.lowestCount)

	s.lowestCount++

	return item
}

func (s *Queue[T]) Peek() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	return s.itens[s.lowestCount]
}

func (s *Queue[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Queue[T]) Clear() {
	s.itens = make(map[int]T)
	s.count = 0
	s.lowestCount = 0
}

func (s *Queue[T]) ToString() string {
	values := make([]T, 0, len(s.itens))
	for _, v := range s.itens {
		values = append(values, v)
	}
	return fmt.Sprintf("%v", values)
}

func (s *Queue[T]) Size() int {
	return s.count - s.lowestCount
}
