package dictionary

type Dictionary[T comparable] struct {
	Table map[string]T
}

func NewDictionary[T comparable]() *Dictionary[T] {
	return &Dictionary[T]{
		Table: make(map[string]T),
	}
}

func (s *Dictionary[T]) HasKey(key string) bool {
	_, ok := s.Table[key]

	if !ok {
		return false
	}

	return true
}

func (s *Dictionary[T]) Set(key string, value T) bool {
	if key == "" {
		return false
	}

	s.Table[key] = value

	return true
}

func (s *Dictionary[T]) Remove(key string) bool {
	if !s.HasKey(key) {
		return false
	}

	delete(s.Table, key)

	return true
}

func (s *Dictionary[T]) Get(key string) T {
	if !s.HasKey(key) {
		var zero T
		return zero
	}

	return s.Table[key]
}

func (s *Dictionary[T]) Keys() []string {
	keys := []string{}

	for key := range s.Table {
		keys = append(keys, key)
	}

	return keys
}

func (s *Dictionary[T]) Values() []T {
	values := []T{}

	for _, value := range s.Table {
		values = append(values, value)
	}

	return values
}

func (s *Dictionary[T]) ForEach(fn func(key string, value T) bool) {
	for key, value := range s.Table {
		result := fn(key, value)
		if !result {
			break
		}
	}
}
