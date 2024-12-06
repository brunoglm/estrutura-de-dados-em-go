package main

import (
	"dic/dictionary"
	"fmt"
)

func main() {
	s := dictionary.NewDictionary[int]()
	s.Set("1", 1)
	s.Set("2", 2)
	s.Set("3", 3)
	s.Set("4", 4)
	fmt.Println(s.ToString())
}
