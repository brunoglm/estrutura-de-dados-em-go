package floydwarshall

const INF = 9223372036854775807

func FloydWarshall(graph [][]int) [][]int {
	size := len(graph)
	dist := make([][]int, size)

	for i := 0; i < size; i++ {
		dist[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				dist[i][j] = 0
			} else if !isFinite(graph[i][j]) {
				dist[i][j] = INF
			} else {
				dist[i][j] = graph[i][j]
			}
		}
	}

	return dist
}

func isFinite(i int) bool {
	return i != 0
}
