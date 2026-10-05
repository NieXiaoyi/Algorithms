package main

import (
	"fmt"

	digraph "com.algorithms/graph/3_digraph"
)

// 将后序遍历代码增加稍加改造，即可得到有向图的环检测代码
func checkring(g *digraph.Graph) bool {
	path := make([]int, 0)
	inpath := make([]bool, g.VertexNum())
	visited := make([]bool, g.VertexNum())

	var dfs func(g *digraph.Graph, start int) bool
	dfs = func(g *digraph.Graph, start int) bool {
		// visited的值只会在本轮dfs时才会设为true，所有如果visited[start]为true，说明本次搜索已经访问过该节点了，必然存在环
		if visited[start] {
			return true
		}

		// inpath[start]为true，说明之前轮次的dfs已经访问过start节点了，本轮dfs无需再次访问
		if inpath[start] {
			return false
		}

		visited[start] = true
		for _, next := range g.Adj(start) {
			if dfs(g, next) {
				return true
			}
		}
		// 因为visited的值只会在本轮dfs才会设为true，因此在函数返回前，需要将visited还原
		visited[start] = false
		// 后序遍历，函数退出时将顶点存入path中
		inpath[start] = true
		path = append(path, start)
		return false
	}

	for v := range inpath {
		if inpath[v] {
			continue
		}
		if dfs(g, v) {
			return true
		}
	}

	return false
}

func main() {
	g1 := digraph.NewGraph(6)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 4)
	g1.AddEdge(4, 5)
	g1.AddEdge(4, 1)
	g1.AddEdge(3, 0)
	fmt.Println(checkring(g1))

	g2 := digraph.NewGraph(6)
	g2.AddEdge(2, 3)
	g2.AddEdge(3, 4)
	g2.AddEdge(4, 5)
	g2.AddEdge(4, 1)
	g2.AddEdge(3, 0)
	g2.AddEdge(1, 3)
	fmt.Println(checkring(g2))
}
