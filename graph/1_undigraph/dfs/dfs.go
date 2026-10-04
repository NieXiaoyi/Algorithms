package main

/* DFS适合求解单点路径问题，即判断两个顶点间是否连通，也适合求解所有路径 */

import (
	"fmt"

	undigraph "com.algorithms/graph/1_undigraph"
)

func Connected(g *undigraph.Graph, node, target int, visited []bool) bool {
	if visited[node] {
		return false
	}

	if node == target {
		return true
	}

	visited[node] = true
	for _, v := range g.AdjVertex(node) {
		if Connected(g, v, target, visited) {
			return true
		}
	}
	// 因为是求解连通性，而不是枚举所有路径，因此不需要恢复visited[node]
	// visited[node] = false

	return false
}

func main() {
	g := undigraph.NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)

	fmt.Println(Connected(g, 0, 4, make([]bool, g.VertexNum())))
	fmt.Println(Connected(g, 0, 3, make([]bool, g.VertexNum())))
}
