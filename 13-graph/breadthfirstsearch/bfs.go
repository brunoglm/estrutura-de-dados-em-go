package breadthfirstsearch

import (
	"fmt"
	"labgraph/graph"
	"labgraph/queue"
	"labgraph/stack"
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

func BfsfindingShortestRoute[T comparable](g *graph.Graph[T], startVertex T, printPath bool) (map[T]int, map[T]T) {
	vertices := g.GetVertices()
	adjList := g.GetAdjList()
	color := initializeColor(vertices)
	q := queue.NewQueue[T]()
	distances := make(map[T]int)
	predecessors := make(map[T]T)
	q.Enqueue(startVertex)

	for _, v := range vertices {
		distances[v] = 0
		predecessors[v] = *new(T)
	}

	for !q.IsEmpty() {
		u := q.Dequeue()
		neighbors := adjList[u]
		color[u] = GREY
		for _, w := range neighbors {
			if color[w] == WHITE {
				color[w] = GREY
				distances[w] = distances[u] + 1
				predecessors[w] = u
				q.Enqueue(w)
			}
		}
		color[u] = BLACK
	}

	if printPath {
		bfsPrintPath(predecessors, vertices)
	}

	return distances, predecessors
}

func bfsPrintPath[T comparable](predecessors map[T]T, vertices []T) {
	fromVertex := vertices[0]
	for i := 1; i < len(vertices); i++ {
		toVertex := vertices[i]
		path := stack.NewStack[T]()
		for v := toVertex; v != fromVertex; v = predecessors[v] {
			path.Push(v)
		}
		path.Push(fromVertex)
		s := fmt.Sprintf("%v", path.Pop())
		for !path.IsEmpty() {
			s += fmt.Sprintf(" - %v", path.Pop())
		}
		fmt.Println(s)
	}
}
