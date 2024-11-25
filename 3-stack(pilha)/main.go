package main

import (
	"fmt"
	"stack-lab/stack"
	"strconv"
)

func decimalToBinary(num int) {
	remStack := stack.NewStack[int]()
	number := num
	rem := 0
	binaryString := ""

	for number > 0 {
		rem = number % 2
		remStack.Push(rem)
		number = number / 2
	}

	for !remStack.IsEmpty() {
		binaryString += strconv.Itoa(remStack.Pop())
	}

	fmt.Println(binaryString)
}

func baseConverter(num int, base int) {
	remStack := stack.NewStack[int]()
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	number := num
	rem := 0
	binaryString := ""

	if !(base >= 2 && base <= 36) {
		return
	}

	for number > 0 {
		rem = number % base
		remStack.Push(rem)
		number = number / base
	}

	for !remStack.IsEmpty() {
		binaryString += string(digits[remStack.Pop()])
	}

	fmt.Println(binaryString)
}

func main() {
	s := stack.NewStackMap[string]()

	fmt.Println(s.IsEmpty()) // Exibe true
	s.Push("John", "Jack")
	fmt.Println(s.ToString()) // John Jack
	s.Push("Camila")
	fmt.Println(s.ToString()) // John Jack Camila
	fmt.Println(s.Size())     // Exibe 3
	fmt.Println(s.IsEmpty())  // Exibe false
	s.Pop()                   // Remove Camila
	s.Pop()                   // Remove Jack
	fmt.Println(s.ToString()) // John

	fmt.Println()

	decimalToBinary(10)

	fmt.Println()

	baseConverter(10, 2)
	baseConverter(100345, 2)
	baseConverter(100345, 8)
	baseConverter(100345, 16)
	baseConverter(100345, 35)
}
