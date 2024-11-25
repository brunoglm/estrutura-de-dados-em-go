package main

import "fmt"

type Item struct {
	ID    string
	Value string
}

type HashTable struct {
	items map[string]Item
}

func NewHashTable() *HashTable {
	return &HashTable{
		items: make(map[string]Item),
	}
}

// Create a new item in the hash table
func (ht *HashTable) Create(item Item) {
	ht.items[item.ID] = item
}

// Read an item from the hash table by ID
func (ht *HashTable) Read(id string) (Item, bool) {
	item, exists := ht.items[id]
	return item, exists
}

// Update an existing item in the hash table by ID
func (ht *HashTable) Update(id string, newValue string) bool {
	item, exists := ht.Read(id)
	if exists {
		item.Value = newValue
		ht.items[id] = item
		return true
	}
	return false
}

// Delete an item from the hash table by ID
func (ht *HashTable) Delete(id string) bool {
	_, exists := ht.Read(id)
	if exists {
		delete(ht.items, id)
		return true
	}
	return false
}

func main() {
	ht := NewHashTable()

	// Create items
	ht.Create(Item{ID: "1", Value: "Item 1"})
	ht.Create(Item{ID: "2", Value: "Item 2"})

	// Read items
	if item, exists := ht.Read("1"); exists {
		fmt.Println("Item: ", item)
	} else {
		fmt.Println("Item not found, ", exists)
	}

	if item, exists := ht.Read("3"); exists {
		fmt.Println("Item: ", item)
	} else {
		fmt.Println("Item 3 not found")
	}

	fmt.Println("Items: ", ht.items)

	// Update items
	ht.Update("2", "Updated Item 2")
	ht.Update("3", "This won't work")

	// // Delete items
	ht.Delete("1")
	ht.Delete("3")

	fmt.Println("Items: ", ht.items)
}
