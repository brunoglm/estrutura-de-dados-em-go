package main

import "fmt"

func main() {
	// criação do slice(array)
	items := []int{}

	// adição de itens
	items = append(items, 10)
	items = append(items, 20)
	items = append(items, 30)

	// verificação se existe
	fmt.Println("Item 10 exists: ", itemExists(items, 10))

	// get
	fmt.Println("Item: ", items[2])

	// update de item
	items[1] = 100

	// deleção de item
	index := 1
	items = append(items[:index], items[index+1:]...)

	// percorrendo itens
	percorrerItems(items)

	fmt.Println("Items: ", items)
}

func percorrerItems(items []int) {
	for i := 0; i < len(items); i++ {
		fmt.Println("percorrerItems = i: ", i, " item: ", items[i])
	}

	for i, v := range items {
		fmt.Println("percorrerItems = i: ", i, " item: ", v)
	}
}

func itemExists(items []int, item int) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
