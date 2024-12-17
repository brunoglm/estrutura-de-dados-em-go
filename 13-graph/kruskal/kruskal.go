package kruskal

import (
	"sort"
)

type Edge struct {
	Source, Target, Weight int
}

func adjacencyMatrixToEdgeList(graph [][]int) []Edge {
	var edges []Edge
	size := len(graph)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ { // Considera apenas a metade superior (grafo não direcionado)
			if graph[i][j] > 0 {
				edges = append(edges, Edge{Source: i, Target: j, Weight: graph[i][j]})
			}
		}
	}
	return edges
}

// Kruskal encontra a Árvore Geradora Mínima (MST) de um grafo.
func Kruskal(graph [][]int) ([]Edge, int) {
	size := len(graph)
	edges := adjacencyMatrixToEdgeList(graph)
	// Ordena as arestas pelo peso
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// Inicializa estruturas do Union-Find
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}

	var mst []Edge
	totalWeight := 0

	for _, edge := range edges {
		uRoot := find(edge.Source, parent)
		vRoot := find(edge.Target, parent)

		// Se não formar um ciclo, adiciona a aresta na MST
		if uRoot != vRoot {
			mst = append(mst, edge)
			totalWeight += edge.Weight
			union(uRoot, vRoot, parent, rank)
		}
	}

	return mst, totalWeight
}

// Encontra o representante (root) de um nó com path compression
func find(node int, parent []int) int {
	if parent[node] != node {
		parent[node] = find(parent[node], parent) // Path compression
	}
	return parent[node]
}

// Realiza a união de dois conjuntos com união por ranking
func union(u, v int, parent, rank []int) {
	if rank[u] < rank[v] {
		parent[u] = v
	} else if rank[u] > rank[v] {
		parent[v] = u
	} else {
		parent[v] = u
		rank[u]++
	}
}
