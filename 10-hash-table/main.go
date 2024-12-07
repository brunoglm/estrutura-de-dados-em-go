package main

import (
	"fmt"
	"hash/linearProbingHashtable"
)

func main() {
	// ht := hashtable.NewHashtable[string]()
	// ht.Put("Gandalf", "gandalf@email.com")
	// ht.Put("John", "john@email.com")
	// ht.Put("Tyrion", "tyrion@email.com")
	// fmt.Println(ht.HashCode("Gandalf"), " - Gandalf")
	// fmt.Println(ht.HashCode("John"), " - John")
	// fmt.Println(ht.HashCode("Tyrion"), " - Tyrion")
	// fmt.Println(ht.Get("Gandalf"))
	// fmt.Println(ht.Get("John"))
	// fmt.Println(ht.Get("Tyrion"))
	// fmt.Println(ht.Get("teste"))

	ht := linearProbingHashtable.NewHashtable[string]()
	ht.Put("Ygritte", "Ygritte@email.com")
	ht.Put("Jonathan", "Jonathan@email.com")
	ht.Put("Jamie", "Jamie@email.com")
	ht.Put("Jack", "Jack@email.com")
	ht.Put("Jasmine", "Jasmine@email.com")
	ht.Put("Jake", "Jake@email.com")
	ht.Put("Nathan", "Nathan@email.com")
	ht.Put("Athelstan", "Athelstan@email.com")
	ht.Put("Sue", "Sue@email.com")
	ht.Put("Aethelwulf", "Aethelwulf@email.com")
	ht.Put("Sargeras", "Sargeras@email.com")
	fmt.Println(ht.GetTable())
}
