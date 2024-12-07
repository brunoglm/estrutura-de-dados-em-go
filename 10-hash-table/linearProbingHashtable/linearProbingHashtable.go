package linearProbingHashtable

type KeyValue[T any] struct {
	Key   string
	Value T
}

type LinearProbingHashtable[T any] struct {
	table map[int]KeyValue[T]
}

func NewHashtable[T any]() *LinearProbingHashtable[T] {
	return &LinearProbingHashtable[T]{
		table: map[int]KeyValue[T]{},
	}
}

func (t *LinearProbingHashtable[T]) loseloseHashCode(key string) int {
	hash := 0

	for _, char := range key {
		hash += int(char)
	}

	return hash % 37
}

func (t *LinearProbingHashtable[T]) HashCode(key string) int {
	return t.loseloseHashCode(key)
}

func (t *LinearProbingHashtable[T]) Put(key string, value T) bool {
	if key == "" {
		return false
	}

	hash := t.HashCode(key)
	originalHash := hash

	for t.has(hash) {
		hash++
		if hash == originalHash {
			return false
		}
	}

	t.table[hash] = KeyValue[T]{key, value}

	return true
}

// func (t *LinearProbingHashtable[T]) Remove(key string) bool {
// 	hash := t.HashCode(key)

// 	if !t.has(hash) {
// 		return false
// 	}

// 	delete(t.table, hash)

// 	return true
// }

func (t *LinearProbingHashtable[T]) Get(key string) T {
	var zero T
	hash := t.HashCode(key)

	if !t.has(hash) {
		return zero
	}

	if t.table[hash].Key == key {
		return t.table[hash].Value
	}

	index := hash + 1
	for t.has(index) && t.table[index].Key != key {
		index++
	}

	if t.has(index) && t.table[index].Key != key {
		return t.table[index].Value
	}

	return zero
}

func (t *LinearProbingHashtable[T]) has(hash int) bool {
	_, exists := t.table[hash]
	return exists
}

func (t *LinearProbingHashtable[T]) GetTable() map[int]KeyValue[T] {
	return t.table
}
