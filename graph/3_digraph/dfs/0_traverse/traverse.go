package main

import (
	"fmt"
	"slices"

	digraph "com.algorithms/graph/3_digraph"
)

func preorder_traverse(g *digraph.Graph, start int) []int {
	path := []int{}
	visited := make([]bool, g.VertexNum())

	var dfs func(g *digraph.Graph, start int)
	dfs = func(g *digraph.Graph, start int) {
		if visited[start] {
			return
		}

		visited[start] = true
		path = append(path, start)
		for _, next := range g.Adj(start) {
			dfs(g, next)
		}
	}

	dfs(g, start)
	return path
}

func postorder_traverse(g *digraph.Graph, start int) []int {
	path := []int{}
	visited := make([]bool, g.VertexNum())

	var dfs func(g *digraph.Graph, start int)
	dfs = func(g *digraph.Graph, start int) {
		if visited[start] {
			return
		}

		visited[start] = true
		for _, next := range g.Adj(start) {
			dfs(g, next)
		}
		path = append(path, start)
	}

	dfs(g, start)
	return path
}

// 因为初次访问的顶点未必是整张图的起点，因此有向图的遍历必须使用后序遍历
func traverse(g *digraph.Graph) []int {
	path := []int{}
	visited := make([]bool, g.VertexNum())

	var dfs func(g *digraph.Graph, start int)
	dfs = func(g *digraph.Graph, start int) {
		if visited[start] {
			return
		}

		visited[start] = true
		for _, next := range g.Adj(start) {
			dfs(g, next)
		}
		path = append(path, start)
	}

	for v := range visited {
		if visited[v] {
			continue
		}
		dfs(g, v)
	}

	slices.Reverse(path)
	return path
}

func main() {
	// 只有无环图才能有顺序的遍历（即进行拓扑排序）
	g1 := digraph.NewGraph(6)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 4)
	g1.AddEdge(3, 0)
	g1.AddEdge(0, 5)
	g1.AddEdge(0, 1)
	fmt.Println(preorder_traverse(g1, 3))
	fmt.Println(postorder_traverse(g1, 2))
	fmt.Println(traverse(g1))

	// 有环图遍历得出的结果并非有序的
	g2 := digraph.NewGraph(6)
	g2.AddEdge(2, 3)
	g2.AddEdge(3, 4)
	g2.AddEdge(3, 0)
	g2.AddEdge(0, 5)
	g2.AddEdge(0, 1)
	g2.AddEdge(1, 3)
	// 遍历结果[3 4 0 5 1]，1排在3后面，但是因为存在环，1还应该要排在3的前面（自相矛盾）
	fmt.Println(preorder_traverse(g2, 3))
	// 遍历结果[4 5 1 0 3 2]，1排在3前面，但是因为存在环，1还应该要排在3的后面（自相矛盾）
	fmt.Println(postorder_traverse(g2, 2))
	// 遍历结果[2 0 1 3 4 5]，1排在3前面，但是因为存在环，1还应该要排在3的后面（自相矛盾）
	fmt.Println(traverse(g2))
}
