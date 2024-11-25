package main

import (
	"fmt"
	"math/rand"
	"queue-lab/queue"
)

func hotPotato(items []string, num int) ([]string, string) {
	queue := queue.NewQueue[string]()
	elimitatedList := make([]string, 0, len(items)-1)

	queue.Enqueue(items...)

	for queue.Size() > 1 {
		for i := 0; i < num; i++ {
			queue.Enqueue(queue.Dequeue())
		}
		elimitatedList = append(elimitatedList, queue.Dequeue())
	}

	return elimitatedList, queue.Dequeue()
}

func main() {
	// qe := queue.NewQueue[string]()

	// fmt.Println(qe.IsEmpty()) // Exibe true
	// qe.Enqueue("John", "Jack")
	// fmt.Println(qe.ToString()) // John Jack
	// qe.Enqueue("Camila")
	// fmt.Println(qe.ToString()) // John Jack Camila
	// fmt.Println(qe.Size())     // Exibe 3
	// fmt.Println(qe.IsEmpty())  // Exibe false
	// qe.Dequeue()               // Remove John
	// qe.Dequeue()               // Remove Jack
	// fmt.Println(qe.ToString()) // Camila

	// fmt.Println()
	// fmt.Println()

	list := []string{"a1", "a2", "a3", "a4"}
	fmt.Println("competidores: ", list)

	randomInt := rand.Intn(20)

	listElimitated, winner := hotPotato(list, randomInt)

	fmt.Println("listElimitated: ", listElimitated)
	fmt.Println("winner: ", winner)

}
