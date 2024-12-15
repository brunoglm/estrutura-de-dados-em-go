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

func Depthfirstsearch[T comparable](g *graph.Graph[T], startVertex T, callback func(v T)) {
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
