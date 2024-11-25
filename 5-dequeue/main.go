package main

import (
	"deque-lab/deque"
	"fmt"
	"strings"
)

func palindromeChecker(aString string) bool {
	if len(aString) == 0 {
		return false
	}

	deque := deque.NewDeque[string]()
	aString = strings.ReplaceAll(strings.ToLower(aString), " ", "")
	isEqual := true

	for _, rune := range aString {
		deque.AddBack(string(rune))
	}

	var firstChar, lastChar string

	for deque.Size() > 1 && isEqual {
		firstChar = deque.RemoveFront()
		lastChar = deque.RemoveBack()
		if firstChar != lastChar {
			isEqual = false
		}
	}

	return isEqual
}

func main() {
	// de := deque.NewDeque[string]()

	// fmt.Println(de.IsEmpty()) // exibe true
	// de.AddBack("John")
	// de.AddBack("Jack")
	// fmt.Println(de.ToString()) // John Jack
	// de.AddBack("Camila")
	// fmt.Println(de.ToString()) // John Jack Camila
	// fmt.Println(de.Size())     // Exibe 3
	// fmt.Println(de.IsEmpty())  // Exibe false
	// de.RemoveFront()           // Remove John
	// fmt.Println(de.ToString()) // Jack Camila
	// de.RemoveBack()            // Camila decide sair
	// fmt.Println(de.ToString()) // Jack
	// de.AddFront("John")        // John retorna para pedir uma informação
	// fmt.Println(de.ToString()) // John Jack

	fmt.Println(palindromeChecker("a"))
	fmt.Println(palindromeChecker("aa"))
	fmt.Println(palindromeChecker("Kayak"))
	fmt.Println(palindromeChecker("Level"))
	fmt.Println(palindromeChecker("Was it a car or a cat I saw"))
}
