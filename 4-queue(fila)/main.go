package main

import (
	"fmt"
	"queue-lab/queue"
)

func main() {
	s := queue.NewQueue[int]()
	fmt.Println(s.ToString())
	s.Enqueue(10, 20, 30, 40)
	fmt.Println(s.ToString())
	s.Dequeue()
	fmt.Println(s.ToString())
	s.Dequeue()
	fmt.Println(s.ToString())
	s.Enqueue(100, 200)
	fmt.Println(s.ToString())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())
	fmt.Println(s.Peek())
	s.Clear()
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())
	fmt.Println(s.ToString())
}
