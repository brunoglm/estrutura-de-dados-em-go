package main

import (
	"fmt"
	"labgraph/dijkstra"
)

func main() {
	// g := graph.NewGraph[string](false)

	// myVertices := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

	// for _, v := range myVertices {
	// 	g.AddVertex(v)
	// }

	// g.AddEdge("A", "B")
	// g.AddEdge("A", "C")
	// g.AddEdge("A", "D")
	// g.AddEdge("C", "D")
	// g.AddEdge("C", "G")
	// g.AddEdge("D", "G")
	// g.AddEdge("D", "H")
	// g.AddEdge("B", "E")
	// g.AddEdge("B", "F")
	// g.AddEdge("E", "I")

	// fmt.Println(g.ToString())
	// breadthfirstsearch.BreadthFirstSearch(g, "A", func(v string) { fmt.Println("Visited vertex: ", v) })
	// distances, predecessors := breadthfirstsearch.BfsfindingShortestRoute(g, "A", true)
	// fmt.Println("distances: ", distances)
	// fmt.Println("predecessors: ", predecessors)

	// g := graph.NewGraph[string](true)

	// myVertices := []string{"A", "B", "C", "D", "E", "F"}

	// for _, v := range myVertices {
	// 	g.AddVertex(v)
	// }

	// g.AddEdge("A", "C")
	// g.AddEdge("A", "D")
	// g.AddEdge("B", "D")
	// g.AddEdge("B", "E")
	// g.AddEdge("C", "F")
	// g.AddEdge("F", "E")

	// discovery, finished, predecessors := depthfirstsearch.DfsInstantDiscoveryExploration(g)
	// fmt.Println("discovery: ", discovery)
	// fmt.Println("finished: ", finished)
	// fmt.Println("predecessors: ", predecessors)

	// result := depthfirstsearch.DfsTopoSort(myVertices, finished)
	// fmt.Println(result)

	graph := [][]int{
		{0, 2, 4, 0, 0, 0},
		{0, 0, 1, 4, 2, 0},
		{0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0, 2},
		{0, 0, 0, 3, 0, 2},
		{0, 0, 0, 0, 0, 0},
	}

	dist := dijkstra.Dijkstra(graph, 0)
	fmt.Println("dist: ", dist)
}
