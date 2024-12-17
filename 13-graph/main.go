package main

import (
	"fmt"
	"labgraph/kruskal"
	"labgraph/prim"
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

	// graph := [][]int{
	// 	{0, 2, 4, 0, 0, 0},
	// 	{0, 0, 2, 4, 2, 0},
	// 	{0, 0, 0, 0, 3, 0},
	// 	{0, 0, 0, 0, 0, 2},
	// 	{0, 0, 0, 3, 0, 2},
	// 	{0, 0, 0, 0, 0, 0},
	// }

	// graph := [][]int{
	// 	{0, 2, 2, 0},
	// 	{0, 0, 0, 2},
	// 	{0, 0, 0, 4},
	// 	{0, 0, 0, 0},
	// }

	// dist := dijkstra.Dijkstra(graph, 0)
	// fmt.Println("dist: ", dist)

	// distances := floydwarshall.FloydWarshall(graph)
	// for _, distance := range distances {
	// 	for _, d := range distance {
	// 		fmt.Printf("%v ", d)
	// 	}
	// 	fmt.Printf("\n")
	// }

	graph := [][]int{
		{0, 2, 4, 0, 0, 0},
		{2, 0, 2, 4, 2, 0},
		{4, 2, 0, 0, 3, 0},
		{0, 4, 0, 0, 3, 2},
		{0, 2, 3, 3, 0, 2},
		{0, 0, 0, 2, 2, 0},
	}

	parent, key := prim.Prim(graph)
	fmt.Println("parent: ", parent)
	fmt.Println("key: ", key)
	fmt.Println()

	mst, totalWeight := kruskal.Kruskal(graph)
	fmt.Println("mst: ", mst)
	fmt.Println("totalWeight: ", totalWeight)
}
