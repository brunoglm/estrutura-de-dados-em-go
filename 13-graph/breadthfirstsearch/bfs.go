package breadthfirstsearch

import (
	"labgraph/graph"
	"labgraph/queue"
)

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

func BreadthFirstSearch[T comparable](g *graph.Graph[T], startVertex T, callback func(v T)) {
	vertices := g.GetVertices()
	adjList := g.GetAdjList()
	color := initializeColor(vertices)
	q := queue.NewQueue[T]()
	q.Enqueue(startVertex)

	for !q.IsEmpty() {
		u := q.Dequeue()
		neighbors := adjList[u]
		color[u] = GREY
		for _, w := range neighbors {
			if color[w] == WHITE {
				color[w] = GREY
				q.Enqueue(w)
			}
		}
		color[u] = BLACK
		if callback != nil {
			callback(u)
		}
	}
}
