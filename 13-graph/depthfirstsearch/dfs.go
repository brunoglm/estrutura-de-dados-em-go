package depthfirstsearch

import "labgraph/graph"

type color int

const (
	WHITE color = iota
	GREY
	BLACK
)

func initializeColor[T comparable](vertices []T) map[T]color {
	color := make(map[T]color)
	for _, v := range vertices {
		color[v] = WHITE
	}
	return color
}

func Depthfirstsearch[T comparable](g *graph.Graph[T], callback func(v T)) {
	vertices := g.GetVertices()
	adjList := g.GetAdjList()
	color := initializeColor(vertices)
	for _, v := range vertices {
		if color[v] == WHITE {
			depthfirstsearchVisit(v, color, adjList, callback)
		}
	}
}

func depthfirstsearchVisit[T comparable](u T, color map[T]color, adjList map[T][]T, callback func(v T)) {
	color[u] = GREY

	if callback != nil {
		callback(u)
	}

	neighbors := adjList[u]

	for _, w := range neighbors {
		if color[w] == WHITE {
			depthfirstsearchVisit(w, color, adjList, callback)
		}
	}

	color[u] = BLACK
}

func DfsInstantDiscoveryExploration[T comparable](g *graph.Graph[T]) (map[T]int, map[T]int, map[T]T) {
	vertices := g.GetVertices()
	adjList := g.GetAdjList()
	color := initializeColor(vertices)
	discovery := make(map[T]int)
	finished := make(map[T]int)
	predecessors := make(map[T]T)
	time := map[string]int{
		"count": 0,
	}

	for _, v := range vertices {
		finished[v] = 0
		discovery[v] = 0
		predecessors[v] = *new(T)
	}

	for _, v := range vertices {
		if color[v] == WHITE {
			dfsInstantDiscoveryExplorationVisit(v, color, discovery, finished, predecessors, time, adjList)
		}
	}

	return discovery, finished, predecessors
}

func dfsInstantDiscoveryExplorationVisit[T comparable](u T, color map[T]color, discovery map[T]int, finished map[T]int, predecessors map[T]T, time map[string]int, adjList map[T][]T) {
	color[u] = GREY
	time["count"]++
	discovery[u] = time["count"]

	neighbors := adjList[u]

	for _, w := range neighbors {
		if color[w] == WHITE {
			predecessors[w] = u
			dfsInstantDiscoveryExplorationVisit(w, color, discovery, finished, predecessors, time, adjList)
		}
	}

	color[u] = BLACK
	time["count"]++
	finished[u] = time["count"]
}
