package dijkstra

const INF = 9223372036854775807

func Dijkstra(graph [][]int, src int) []int {
	size := len(graph)
	dist := make([]int, 0, size)
	visited := make([]bool, 0, size)

	for i := 0; i < size; i++ {
		dist[i] = INF
		visited[i] = false
	}

	dist[src] = 0

	return dist
}
