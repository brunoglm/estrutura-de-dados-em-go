package main

import (
	"deque-lab/deque"
	"fmt"
)

func main() {
	de := deque.NewDeque[string]()

	fmt.Println(de.IsEmpty()) // exibe true
	de.AddBack("John")
	de.AddBack("Jack")
	fmt.Println(de.ToString()) // John Jack
	de.AddBack("Camila")
	fmt.Println(de.ToString()) // John Jack Camila
	fmt.Println(de.Size())     // Exibe 3
	fmt.Println(de.IsEmpty())  // Exibe false
	de.RemoveFront()           // Remove John
	fmt.Println(de.ToString()) // Jack Camila
	de.RemoveBack()            // Camila decide sair
	fmt.Println(de.ToString()) // Jack
	de.AddFront("John")        // John retorna para pedir uma informação
	fmt.Println(de.ToString()) // John Jack

}
