package graph

type Graph[T any] struct {
	IsDirected bool
	Vertices   []T
	AdjList    map[string]T
}
