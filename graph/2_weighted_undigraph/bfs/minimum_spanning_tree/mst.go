package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"

	undigraph "com.algorithms/graph/2_weighted_undigraph"
)

/* 图的生成树是它的一棵含有其所有顶点的无环连通子图。一棵生成树必然有V-1条边，其中V指图的顶点个数。
** 一幅加权图的最小生成树（MST）是它的一棵权值（树中所有边的权值之和）最小的生成树。
** 切分定理：在一幅加权图中，给定任意的切分，它的横切边中的权重最小者必然属于图的最小生成树
 */

type Edge struct {
	x      int
	y      int
	weight int
}

type PriorityQueue []Edge

func (q *PriorityQueue) Len() int {
	return len(*q)
}

func (q *PriorityQueue) Less(i, j int) bool {
	return (*q)[i].weight < (*q)[j].weight
}

func (q *PriorityQueue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *PriorityQueue) Push(x any) {
	if edge, ok := x.(Edge); ok {
		*q = append(*q, edge)
	}
}

func (q *PriorityQueue) Pop() any {
	defer func() {
		*q = (*q)[:len(*q)-1]
	}()

	return (*q)[len(*q)-1]
}

// 根据切分定理，我们每次轮询都获取到一条生成树和图中其他节点的横切边的最小边，然后将对应的顶点加入生成树中，直到所有顶点都加入生成树中为止，这样获得的生成树即为最小生成树
func Prime(g *undigraph.Graph) *undigraph.Graph {
	// 保存相邻节点全部已经遍历过的节点
	visited := make([]bool, g.VertexNum())
	tree := undigraph.NewGraph(g.VertexNum())
	// 保存权重边的最小堆
	pq := &PriorityQueue{}

	var visit func(v int)
	visit = func(v int) {
		visited[v] = true
		for _, next := range g.AdjVertex(v) {
			heap.Push(pq, Edge{
				v,
				next,
				g.Weight(v, next),
			})
		}
	}

	for v := range visited {
		if visited[v] {
			continue
		}
		visit(v)

		for pq.Len() > 0 {
			edge, _ := heap.Pop(pq).(Edge)
			// 该边的两个顶点都已经在最小生成树中了，失效边
			// 从A到B有两条路径，路径1的边B1先入堆，路径2的边B2后入堆，但由于B2的权重小于B1，所以B2先出堆并进树，后续B1出堆时要失效
			if visited[edge.x] && visited[edge.y] {
				continue
			}
			tree.AddEdge(edge.x, edge.y, edge.weight)

			visit(edge.x)
			visit(edge.y)
		}
	}

	return tree
}

// ---------------- 以下为验证 Prime 正确性的 main 方法 ----------------

// addEdges 批量加边，每个元素为 {x, y, weight}
func addEdges(g *undigraph.Graph, edges [][3]int) {
	for _, e := range edges {
		g.AddEdge(e[0], e[1], e[2])
	}
}

// totalWeight 计算图中所有边的权值之和
func totalWeight(g *undigraph.Graph) int {
	sum := 0
	for v := 0; v < g.VertexNum(); v++ {
		for _, next := range g.AdjVertex(v) {
			if v < next {
				sum += g.Weight(v, next)
			}
		}
	}
	return sum
}

// components 用 BFS 统计图的连通分量个数
func components(g *undigraph.Graph) int {
	visited := make([]bool, g.VertexNum())
	count := 0
	for s := 0; s < g.VertexNum(); s++ {
		if visited[s] {
			continue
		}
		count++
		visited[s] = true
		queue := []int{s}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			for _, next := range g.AdjVertex(v) {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	return count
}

// kruskalWeight 用 Kruskal + 并查集独立计算最小生成树（森林）的权值和，作为对照结果
func kruskalWeight(g *undigraph.Graph) int {
	type edge struct{ x, y, w int }

	edges := []edge{}
	for v := 0; v < g.VertexNum(); v++ {
		for _, next := range g.AdjVertex(v) {
			if v < next {
				edges = append(edges, edge{v, next, g.Weight(v, next)})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].w < edges[j].w })

	parent := make([]int, g.VertexNum())
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	sum := 0
	for _, e := range edges {
		rx, ry := find(e.x), find(e.y)
		if rx == ry {
			continue
		}
		parent[rx] = ry
		sum += e.w
	}
	return sum
}

// printEdges 按 x 从小到大打印图中的所有边
func printEdges(g *undigraph.Graph) {
	type edge struct{ x, y, w int }

	edges := []edge{}
	for v := 0; v < g.VertexNum(); v++ {
		for _, next := range g.AdjVertex(v) {
			if v < next {
				edges = append(edges, edge{v, next, g.Weight(v, next)})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].x < edges[j].x })

	for _, e := range edges {
		fmt.Printf("    %d -- %d  权值=%d\n", e.x, e.y, e.w)
	}
}

// verify 校验 Prime 的结果：
//  1. 边数 == 顶点数 - 连通分量数（树/森林恰好不含环且不遗漏顶点）
//  2. 结果的连通分量数与原图一致（原图连通则结果也连通）
//  3. 权值和 == Kruskal 独立算出的最小生成树权值和
func verify(name string, g, mst *undigraph.Graph) bool {
	comp := components(g)
	wantEdges := g.VertexNum() - comp
	wantWeight := kruskalWeight(g)
	gotWeight := totalWeight(mst)

	edgeOK := mst.EdgeNum() == wantEdges
	compOK := components(mst) == comp
	weightOK := gotWeight == wantWeight
	ok := edgeOK && compOK && weightOK

	status := "PASS"
	if !ok {
		status = "FAIL"
	}
	fmt.Printf("[%s] %s: 期望(边数=%d, 权值和=%d, 连通分量=%d) 实际(边数=%d, 权值和=%d, 连通分量=%d)\n",
		status, name, wantEdges, wantWeight, comp, mst.EdgeNum(), gotWeight, components(mst))
	return ok
}

func main() {
	pass := true

	// 用例1：手写图，可人工核算最小生成树为 0-2(1) + 1-2(2) + 3-4(3) + 1-3(5) = 11
	g1 := undigraph.NewGraph(5)
	addEdges(g1, [][3]int{
		{0, 1, 4},
		{0, 2, 1},
		{1, 2, 2},
		{1, 3, 5},
		{2, 3, 8},
		{2, 4, 10},
		{3, 4, 3},
	})
	mst1 := Prime(g1)
	fmt.Println("用例1: 手写图的最小生成树")
	printEdges(mst1)
	pass = verify("用例1: 手写图", g1, mst1) && pass

	// 用例2：非连通图，期望得到最小生成森林（含 3 个连通分量，共 7-3=4 条边）
	g2 := undigraph.NewGraph(7)
	addEdges(g2, [][3]int{
		{0, 1, 1},
		{1, 2, 2},
		{0, 2, 4},
		{4, 5, 3},
		{5, 6, 6},
		{4, 6, 9},
	})
	mst2 := Prime(g2)
	fmt.Println("用例2: 非连通图的最小生成森林")
	printEdges(mst2)
	pass = verify("用例2: 非连通图", g2, mst2) && pass

	// 用例3：只有单个顶点
	g3 := undigraph.NewGraph(1)
	pass = verify("用例3: 单顶点", g3, Prime(g3)) && pass

	// 用例4：链式图（含环、含重复权值，强制走失效边分支）
	g4 := undigraph.NewGraph(4)
	addEdges(g4, [][3]int{
		{0, 1, 3},
		{1, 2, 3},
		{2, 3, 3},
		{0, 3, 3},
		{0, 2, 7},
	})
	pass = verify("用例4: 全等权值图", g4, Prime(g4)) && pass

	// 用例5：固定种子的随机图，与 Kruskal 对照
	r := rand.New(rand.NewSource(42))
	for t := 0; t < 20; t++ {
		n := 5 + r.Intn(16)
		g := undigraph.NewGraph(n)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if r.Intn(3) == 0 { // 约 1/3 概率存在边
					g.AddEdge(i, j, 1+r.Intn(20)) // 权值有重复，覆盖失效边场景
				}
			}
		}
		pass = verify(fmt.Sprintf("用例5: 随机图#%d (顶点=%d, 边=%d)", t, n, g.EdgeNum()), g, Prime(g)) && pass
	}

	if pass {
		fmt.Println("\n所有用例通过，Prime 正确")
	} else {
		fmt.Println("\n存在失败用例，Prime 有误")
	}
}
