package main

import (
	"fmt"
	"queue-lab/queue"
)

func main() {
	qe := queue.NewQueue[string]()

	fmt.Println(qe.IsEmpty()) // Exibe true
	qe.Enqueue("John", "Jack")
	fmt.Println(qe.ToString()) // John Jack
	qe.Enqueue("Camila")
	fmt.Println(qe.ToString()) // John Jack Camila
	fmt.Println(qe.Size())     // Exibe 3
	fmt.Println(qe.IsEmpty())  // Exibe false
	qe.Dequeue()               // Remove John
	qe.Dequeue()               // Remove Jack
	fmt.Println(qe.ToString()) // Camila

}
