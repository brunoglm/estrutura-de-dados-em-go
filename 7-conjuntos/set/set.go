package set

type Set[T comparable] struct {
	count int
	items map[T]T
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]T),
	}
}

func (s *Set[T]) Has(key T) bool {
	_, ok := s.items[key]

	if !ok {
		return false
	}

	return true
}

func (s *Set[T]) Add(key T) bool {
	if s.Has(key) {
		return false
	}

	s.items[key] = key
	s.count++

	return true
}

func (s *Set[T]) Delete(key T) bool {
	if s.Has(key) {
		return false
	}

	delete(s.items, key)
	s.count--

	return true
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]T)
}

func (s *Set[T]) Size() int {
	return s.count
}

func (s *Set[T]) Values() []T {
	arr := make([]T, 0, s.count-1)

	for _, value := range s.items {
		arr = append(arr, value)
	}

	return arr
}

func (s *Set[T]) Union(otherSet *Set[T]) *Set[T] {
	unionSet := NewSet[T]()

	for _, value := range s.items {
		unionSet.Add(value)
	}

	for _, value := range otherSet.items {
		unionSet.Add(value)
	}

	return unionSet
}

func (s *Set[T]) Intersection(otherSet *Set[T]) *Set[T] {
	intersectionSet := NewSet[T]()
	var biggerSet, smallerSet *Set[T]

	if s.count > otherSet.count {
		biggerSet = s
		smallerSet = otherSet
	} else {
		biggerSet = otherSet
		smallerSet = s
	}

	for _, value := range smallerSet.items {
		if biggerSet.Has(value) {
			intersectionSet.Add(value)
		}
	}

	return intersectionSet
}

func (s *Set[T]) Diference(otherSet *Set[T]) *Set[T] {
	diferenceSet := NewSet[T]()

	for _, value := range s.items {
		if !otherSet.Has(value) {
			diferenceSet.Add(value)
		}
	}

	return diferenceSet
}

func (s *Set[T]) IsSubsetOf(otherSet *Set[T]) bool {
	if s.count > otherSet.count {
		return false
	}

	for _, value := range s.items {
		if !otherSet.Has(value) {
			return false
		}
	}

	return true
}
