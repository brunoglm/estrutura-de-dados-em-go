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
		for j, coluna := range linha {
			fmt.Println("Indice linha: ", i, " Indice coluna: ", j, " Item coluna: ", coluna)
		}
	}

	// deletando um item da matriz
	// para alterar, podemos usar a mesma estrategia
	rowToModify := 1
	collumToRemove := 1
	matrix[rowToModify] = append(matrix[rowToModify][:collumToRemove], matrix[rowToModify][collumToRemove+1:]...)

	// percorrendo matriz com for tradicional
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
