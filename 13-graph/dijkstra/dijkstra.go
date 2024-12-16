package dijkstra

const INF = 9223372036854775807

func Dijkstra(graph [][]int, src int) []int {
	size := len(graph)
	dist := make([]int, size)
	visited := make([]bool, size)

	for i := 0; i < size; i++ {
		dist[i] = INF
		visited[i] = false
	}

	dist[src] = 0

	for i := 0; i < size-1; i++ {
		u := minDistance(dist, visited)
		visited[u] = true

		for v := 0; v < size; v++ {
			if !visited[v] &&
				graph[u][v] != 0 &&
				dist[u] != INF &&
				(dist[u]+graph[u][v]) < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	return dist
}

func minDistance(dist []int, visited []bool) int {
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
