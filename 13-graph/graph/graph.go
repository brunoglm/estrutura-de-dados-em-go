package graph

import "fmt"

type Graph[T comparable] struct {
	IsDirected bool
	Vertices   []T
	AdjList    map[T][]T
}

func NewGraph[T comparable](isDirected bool) *Graph[T] {
	return &Graph[T]{
		IsDirected: isDirected,
		Vertices:   make([]T, 0),
		AdjList:    make(map[T][]T),
	}
}

func (g *Graph[T]) hasVertice(v T) bool {
	for _, vertice := range g.Vertices {
		if vertice == v {
			return true
		}
	}
	return false
}

func (g *Graph[T]) hasAdjList(v T) bool {
	_, ok := g.AdjList[v]
	return ok
}

func (g *Graph[T]) AddVertex(v T) {
	if g.hasVertice(v) {
		return
	}
	g.Vertices = append(g.Vertices, v)
	g.AdjList[v] = make([]T, 0)
}

func (g *Graph[T]) AddEdge(v, w T) {
	if !g.hasAdjList(v) {
		g.AddVertex(v)
	}

	if !g.hasAdjList(w) {
		g.AddVertex(w)
	}

	g.AdjList[v] = append(g.AdjList[v], w)
	if !g.IsDirected {
		g.AdjList[w] = append(g.AdjList[w], v)
	}
}

func (g *Graph[T]) GetVertices(v, w T) []T {
	return g.Vertices
}

func (g *Graph[T]) GetAdjList(v, w T) map[T][]T {
	return g.AdjList
}

func (g *Graph[T]) ToString() string {
	s := ""
	for _, v := range g.Vertices {
		s += fmt.Sprintf("%v - > ", v)
		neighbors := g.AdjList[v]
		for _, neighbor := range neighbors {
			s += fmt.Sprintf("%v ", neighbor)
		}
		s += "\n"
	}
	return s
}
