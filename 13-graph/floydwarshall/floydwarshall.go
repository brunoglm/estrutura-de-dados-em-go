package floydwarshall

import "math"

const INF = math.MaxInt32

func FloydWarshall(graph [][]int) [][]int {
	size := len(graph)
	dist := make([][]int, size)

	for i := 0; i < size; i++ {
		dist[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if graph[i][j] == 0 && i != j {
				dist[i][j] = INF
			} else {
				dist[i][j] = graph[i][j]
			}
		}
	}

	for k := 0; k < size; k++ {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if dist[i][k] != INF && dist[k][j] != INF && (dist[i][k]+dist[k][j] < dist[i][j]) {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}
