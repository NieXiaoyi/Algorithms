package main

import (
	"fmt"
	"slices"

	digraph "com.algorithms/graph/2_digraph"
)

/* 输出一幅图的拓扑排序数组
** 拓扑排序定义：原图中如果B顶点在A顶点后面，那么在拓扑排序数组中，B也应当在A后面
** 如果有向图有环，则无法进行拓扑排序，输出空数组即可
 */

// 将后序遍历代码及环检测代码结合，即可得到有向图的拓扑排序代码
func topologysort(g *digraph.Graph) []int {
	path := make([]int, 0)
	inpath := make([]bool, g.VertexNum())
	visited := make([]bool, g.VertexNum())

	var checkring func(g *digraph.Graph, start int) bool
	checkring = func(g *digraph.Graph, start int) bool {
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
			if checkring(g, next) {
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
		if checkring(g, v) {
			return []int{}
		}
	}

	slices.Reverse(path)
	return path
}

func main() {
	g1 := digraph.NewGraph(6)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 4)
	g1.AddEdge(4, 5)
	g1.AddEdge(4, 1)
	g1.AddEdge(3, 0)
	fmt.Println(topologysort(g1))

	g2 := digraph.NewGraph(6)
	g2.AddEdge(2, 3)
	g2.AddEdge(3, 4)
	g2.AddEdge(4, 5)
	g2.AddEdge(4, 1)
	g2.AddEdge(3, 0)
	g2.AddEdge(1, 3)
	fmt.Println(topologysort(g2))
}
