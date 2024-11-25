package deque

import "fmt"

type Deque[T any] struct {
	itens       map[int]T
	count       int
	lowestCount int
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{itens: make(map[int]T)}
}

func (s *Deque[T]) AddFront(item T) {
	if s.IsEmpty() {
		s.AddBack(item)
	} else if s.lowestCount > 0 {
		s.lowestCount--
		s.itens[s.lowestCount] = item
	} else {
		for i := s.count; i > 0; i-- {
			s.itens[i] = s.itens[i-1]
		}
		s.count++
		s.lowestCount = 0
		s.itens[0] = item
	}
}

func (s *Deque[T]) AddFrontUsingNegativeNumbers(item T) {
	if s.IsEmpty() {
		s.AddBack(item)
		return
	}

	s.lowestCount--
	s.itens[s.lowestCount] = item
}

func (s *Deque[T]) AddBack(item T) {
	s.itens[s.count] = item
	s.count++
}

func (s *Deque[T]) RemoveFront() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	item := s.itens[s.lowestCount]

	delete(s.itens, s.lowestCount)

	s.lowestCount++

	return item
}

func (s *Deque[T]) RemoveBack() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	item := s.itens[s.count-1]

	delete(s.itens, s.count-1)

	s.count--

	return item
}

func (s *Deque[T]) PeekFront() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	return s.itens[s.lowestCount]
}

func (s *Deque[T]) PeekBack() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	return s.itens[s.count]
}

func (s *Deque[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Deque[T]) Clear() {
	s.itens = make(map[int]T)
	s.count = 0
	s.lowestCount = 0
}

func (s *Deque[T]) ToString() string {
	values := make([]T, 0, len(s.itens))
	for _, v := range s.itens {
		values = append(values, v)
	}
	return fmt.Sprintf("%v", values)
}

func (s *Deque[T]) Size() int {
	return s.count - s.lowestCount
}
