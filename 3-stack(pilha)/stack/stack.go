package stack

type Stack[T any] struct {
	Itens []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{Itens: make([]T, 0)}
}

func (s *Stack[T]) Add(items ...T) {
	s.Itens = append(s.Itens, items...)
}

func (s *Stack[T]) Pop() T {
	if len(s.Itens) == 0 {
		var zero T
		return zero
	}

	lastIndex := len(s.Itens) - 1
	item := s.Itens[lastIndex]

	s.Itens = s.Itens[:lastIndex]

	return item
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Itens) == 0
}
