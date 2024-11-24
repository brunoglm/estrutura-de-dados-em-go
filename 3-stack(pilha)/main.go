package main

import (
	"fmt"
	"stack-lab/stack"
)

func main() {
	s := stack.NewStack[int]()
	fmt.Println(s.Itens)
	s.Add(10, 20, 50)
	fmt.Println(s.Itens)
	s.Pop()
	fmt.Println(s.Itens)

	fmt.Println()

	s2 := stack.NewStack[string]()
	fmt.Println(s2.Itens)
	s2.Add("s1", "s2", "s3")
	fmt.Println(s2.Itens)
	s2.Pop()
	fmt.Println(s2.Itens)
}
