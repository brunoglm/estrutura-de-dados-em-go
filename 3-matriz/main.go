package main

import "fmt"

func main() {
	// criação matriz com array
	matrix := [][]int{
		{0, 0},
		{0, 0},
	}

	// adição de itens
	matrix[0][0] = 10
	matrix[0][1] = 11
	matrix[1][0] = 20
	matrix[1][1] = 21

	// percorrendo matriz com range
	for i, linha := range matrix {
		for j, elemento := range linha {
			fmt.Println("Indice linha: ", i, " Indice elemento: ", j, " Item elemento: ", elemento)
		}
	}

	// deletando um item da matriz
	// para alterar, podemos usar a mesma estrategia
	rowToModify := 1
	elementToRemove := 1
	matrix[rowToModify] = append(matrix[rowToModify][:elementToRemove], matrix[rowToModify][elementToRemove+1:]...)

	// percorrendo matriz com for tradicional
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
