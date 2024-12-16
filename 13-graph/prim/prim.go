package prim

import "math"

const INF = math.MaxInt32

func Prim(graph [][]int) ([]int, []int) {
	size := len(graph)
	parent := make([]int, size)
	key := make([]int, size)
	visited := make([]bool, size)

	for i := 0; i < size; i++ {
		key[i] = INF
		visited[i] = false
	}

	key[0] = 0
	parent[0] = -1

	for i := 0; i < size-1; i++ {
		u := minKey(key, visited)
		visited[u] = true

		for v := 0; v < size; v++ {
			if graph[u][v] > 0 && !visited[v] && graph[u][v] < key[v] {
				parent[v] = u
				key[v] = graph[u][v]
			}
		}
	}

	return parent, key
}

func minKey(dist []int, visited []bool) int {
	min := INF
	minIndex := -1
	for v := 0; v < len(dist); v++ {
		if !visited[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}
	return minIndex
}
