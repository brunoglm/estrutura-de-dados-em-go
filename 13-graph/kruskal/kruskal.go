package kruskal

import "math"

const INF = math.MaxInt32

func Kruskal(graph [][]int) []int {
	size := len(graph)
	parent := make([]int, size)
	ne := 0
	var a, b, u, v int
	cost := graph

	for ne < size-1 {
		min := INF
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if cost[i][j] < min {
					min = cost[i][j]
					u = i
					a = u
					v = j
					b = v
				}
			}
		}
		u = find(u, parent)
		v = find(v, parent)
		if union(u, v, parent) {
			ne++
		}
		cost[b][a] = INF
		cost[a][b] = cost[b][a]
	}

	return parent
}

func find(i int, parent []int) int {
	for parent[i] > 0 {
		i = parent[i]
	}
	return i
}

func union(j, i int, parent []int) bool {
	if i != j {
		parent[j] = i
		return true
	}
	return false
}
