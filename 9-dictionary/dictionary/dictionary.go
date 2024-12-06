package dictionary

import "fmt"

type Dictionary[T comparable] struct {
	table map[string]T
}

func NewDictionary[T comparable]() *Dictionary[T] {
	return &Dictionary[T]{
		table: make(map[string]T),
	}
}

func (s *Dictionary[T]) HasKey(key string) bool {
	_, ok := s.table[key]

	if !ok {
		return false
	}

	return true
}

func (s *Dictionary[T]) Set(key string, value T) bool {
	if key == "" {
		return false
	}

	s.table[key] = value

	return true
}

func (s *Dictionary[T]) Remove(key string) bool {
	if !s.HasKey(key) {
		return false
	}

	delete(s.table, key)

	return true
}

func (s *Dictionary[T]) Get(key string) T {
	if !s.HasKey(key) {
		var zero T
		return zero
	}

	return s.table[key]
}

func (s *Dictionary[T]) Keys() []string {
	keys := []string{}

	for key := range s.table {
		keys = append(keys, key)
	}

	return keys
}

func (s *Dictionary[T]) Values() []T {
	values := []T{}

	for _, value := range s.table {
		values = append(values, value)
	}

	return values
}

func (s *Dictionary[T]) ForEach(fn func(key string, value T) bool) {
	for key, value := range s.table {
		result := fn(key, value)
		if !result {
			break
		}
	}
}

func (s *Dictionary[T]) Size() int {
	return len(s.table)
}

func (s *Dictionary[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Dictionary[T]) Clear() {
	s.table = make(map[string]T)
}

func (s *Dictionary[T]) ToString() string {
	values := s.Values()

	if len(values) == 0 {
		return ""
	}

	objString := fmt.Sprintf("%v", values[0])

	for i := 1; i < len(values); i++ {
		objString += fmt.Sprintf(",%v", values[i])
	}

	return objString
}
