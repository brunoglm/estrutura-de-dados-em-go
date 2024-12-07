package hashtable

type Hashtable[T any] struct {
	table map[int]T
}

func NewHashtable[T any]() *Hashtable[T] {
	return &Hashtable[T]{
		table: map[int]T{},
	}
}

func (t *Hashtable[T]) loseloseHashCode(key string) int {
	hash := 0

	for _, char := range key {
		hash += int(char)
	}

	return hash
}

func (t *Hashtable[T]) HashCode(key string) int {
	return t.loseloseHashCode(key) % len(t.table)
}

func (t *Hashtable[T]) Put(key string, value T) bool {
	if key == "" {
		return false
	}
	hash := t.HashCode(key)
	t.table[hash] = value
	return true
}

func (t *Hashtable[T]) Remove(key string) bool {
	hash := t.HashCode(key)

	_, exists := t.table[hash]
	if !exists {
		return false
	}

	delete(t.table, hash)

	return true
}

func (t *Hashtable[T]) Get(key string) T {
	hash := t.HashCode(key)

	_, exists := t.table[hash]
	if !exists {
		var zero T
		return zero
	}

	return t.table[hash]
}
