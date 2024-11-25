package main

import (
	"fmt"
	"stack-lab/stack"
)

func main() {
	s := stack.NewStack[int]()
	fmt.Println(s.ToString())
	s.Push(10, 20, 50, 60, 70)
	fmt.Println(s.ToString())
	s.Pop()
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())
	s.Clear()
	fmt.Println(s.ToString())

	fmt.Println()

	s2 := stack.NewStackMap[int]()
	fmt.Println(s2.ToString())
	s2.Push(10, 20, 50, 60, 70)
	fmt.Println(s2.ToString())
	s2.Pop()
	fmt.Println(s2.ToString())
	fmt.Println(s2.Peek())
	fmt.Println(s2.IsEmpty())
	fmt.Println(s2.Size())
	s2.Clear()
	fmt.Println(s2.ToString())
}
