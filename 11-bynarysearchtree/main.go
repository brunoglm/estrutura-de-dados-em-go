package main

import (
	"fmt"
	"labtree/bynarysearchtree"
)

func main() {
	t := bynarysearchtree.NewBynarysearchtree[int](nil)
	t.Insert(11)
	t.Insert(7)
	t.Insert(15)
	t.Insert(5)
	t.Insert(3)
	t.Insert(9)
	t.Insert(8)
	t.Insert(10)
	t.Insert(13)
	t.Insert(12)
	t.Insert(14)
	t.Insert(20)
	t.Insert(18)
	t.Insert(25)
	t.Insert(6)

	t.InOrderTraverse(func(key int) { fmt.Print(key, " ") })
	fmt.Println()
	t.PreOrderTraverse(func(key int) { fmt.Print(key, " ") })
	fmt.Println()
	t.PostOrderTraverse(func(key int) { fmt.Print(key, " ") })
	fmt.Println()
	fmt.Println(t.Min())
	fmt.Println(t.Max())
	fmt.Println(t.Search(3))
	fmt.Println(t.Search(25))
	fmt.Println(t.Search(11))
	fmt.Println(t.Search(10))
	fmt.Println(t.Search(12))
	fmt.Println(t.Search(300))
}
