package main

import (
	"fmt"
	"lab-set/set"
)

func main() {
	// sa := set.NewSet[int]()
	// sa.Add(1)
	// sa.Add(2)
	// sa.Add(3)
	// sb := set.NewSet[int]()
	// sa.Add(3)
	// sa.Add(4)
	// sa.Add(5)
	// sa.Add(6)
	// union := sa.Union(sb)
	// fmt.Println(union.Values())

	// sa := set.NewSet[int]()
	// sa.Add(1)
	// sa.Add(2)
	// sa.Add(3)
	// sb := set.NewSet[int]()
	// sb.Add(2)
	// sb.Add(3)
	// sb.Add(4)
	// intersectionAB := sa.Intersection(sb)
	// fmt.Println(intersectionAB.Values())

	// sa := set.NewSet[int]()
	// sa.Add(1)
	// sa.Add(2)
	// sa.Add(3)
	// sb := set.NewSet[int]()
	// sb.Add(2)
	// sb.Add(3)
	// sb.Add(4)
	// diferenceAB := sa.Diference(sb)
	// diferenceBA := sb.Diference(sa)
	// fmt.Println(diferenceAB.Values())
	// fmt.Println(diferenceBA.Values())

	sa := set.NewSet[int]()
	sa.Add(1)
	sa.Add(2)
	sb := set.NewSet[int]()
	sb.Add(1)
	sb.Add(2)
	sb.Add(3)
	fmt.Println(sa.IsSubsetOf(sb))
	fmt.Println(sb.IsSubsetOf(sa))
}
