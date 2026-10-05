package main

import (
	"fmt"

	undigraph "com.algorithms/graph/1_undigraph"
)

func CheckRing(g *undigraph.Graph) bool {
	visited := make([]bool, g.VertexNum())

	var dfs func(g *undigraph.Graph, vertex int, parent int) bool
	dfs = func(g *undigraph.Graph, cur int, pre int) bool {
		visited[cur] = true
		nextList := g.AdjVertex(cur)
		for _, next := range nextList {
			if !visited[next] {
				if dfs(g, next, cur) {
					return true
				}
			} else if next != pre {
				return true
			}
		}
		return false
	}

	for i := 0; i < g.VertexNum(); i++ {
		if visited[i] {
			continue
		}
		if dfs(g, i, i) {
			return true
		}
	}
	return false
}

func main() {
	g := undigraph.NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)

	fmt.Println(CheckRing(g))
}
