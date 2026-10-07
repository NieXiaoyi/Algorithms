package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"strings"

	digraph "com.algorithms/graph/4_weighted_digraph"
)

// 起点到顶点vertex的路径
type PathLen struct {
	vertex int
	len    int
}

type PriorityQueue []PathLen

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

func (q *PriorityQueue) Len() int {
	return len(*q)
}

func (q *PriorityQueue) Less(i, j int) bool {
	return (*q)[i].len < (*q)[j].len
}

func (q *PriorityQueue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *PriorityQueue) Push(x any) {
	if node, ok := x.(PathLen); ok {
		*q = append(*q, node)
	}
}

func (q *PriorityQueue) Pop() any {
	defer func() {
		*q = (*q)[:len(*q)-1]
	}()

	return (*q)[len(*q)-1]
}

// 返回所有顶点距start的最短路径
func dijkstra(g *digraph.Graph, start int) []int {
	ret := make([]int, g.VertexNum())
	for i := range ret {
		ret[i] = math.MaxInt
	}

	q := NewPriorityQueue()

	heap.Push(q, PathLen{start, 0})
	for q.Len() > 0 {
		vLen := heap.Pop(q).(PathLen)
		// 失效掉已经被加入到最短路径的节点
		if ret[vLen.vertex] != math.MaxInt {
			continue
		}
		ret[vLen.vertex] = vLen.len

		for _, next := range g.Adj(vLen.vertex) {
			if ret[next] != math.MaxInt {
				continue
			}
			nextVLen := PathLen{
				vertex: next,
				len:    ret[vLen.vertex] + g.Weight(vLen.vertex, next),
			}
			heap.Push(q, nextVLen)
		}
	}

	return ret
}

// ---------------- 以下为验证 Dijkstra 正确性的 main 方法 ----------------

// addEdges 批量加边，每个元素为 {s, d, weight}
func addEdges(g *digraph.Graph, edges [][3]int) {
	for _, e := range edges {
		g.AddEdge(e[0], e[1], e[2])
	}
}

// fmtDist 将最短路径数组格式化为字符串，不可达顶点（MaxInt）打印为 ∞
func fmtDist(dist []int) string {
	parts := make([]string, len(dist))
	for i, d := range dist {
		if d == math.MaxInt {
			parts[i] = "∞"
		} else {
			parts[i] = fmt.Sprintf("%d", d)
		}
	}
	return strings.Join(parts, " ")
}

// check 用期望值校验 dijkstra 的结果
func check(name string, g *digraph.Graph, start int, want []int) bool {
	got := dijkstra(g, start)

	ok := len(got) == len(want)
	for i := 0; ok && i < len(want); i++ {
		if got[i] != want[i] {
			ok = false
		}
	}

	status := "PASS"
	if !ok {
		status = "FAIL"
	}
	fmt.Printf("[%s] %s: 起点=%d 期望=[%s] 实际=[%s]\n", status, name, start, fmtDist(want), fmtDist(got))
	return ok
}

// bellmanFord 独立实现的最短路径算法（松弛 V-1 轮），作为随机用例的对照结果
func bellmanFord(g *digraph.Graph, start int) []int {
	dist := make([]int, g.VertexNum())
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[start] = 0

	for i := 0; i < g.VertexNum()-1; i++ {
		updated := false
		for v := 0; v < g.VertexNum(); v++ {
			if dist[v] == math.MaxInt {
				continue
			}
			for _, next := range g.Adj(v) {
				if nd := dist[v] + g.Weight(v, next); nd < dist[next] {
					dist[next] = nd
					updated = true
				}
			}
		}
		// 本轮没有任何松弛，说明已经收敛
		if !updated {
			break
		}
	}
	return dist
}

func main() {
	pass := true

	// 用例1：手写图，人工核算起点 0 到各顶点的最短路径为 0,7,3,9,5
	// 0->1 直接边权值 10 大于绕行 0->2->1 的 7，2->3 的 8 大于绕行 2->1->3 的 6
	g1 := digraph.NewGraph(5)
	addEdges(g1, [][3]int{
		{0, 1, 10},
		{0, 2, 3},
		{1, 2, 1},
		{1, 3, 2},
		{2, 1, 4},
		{2, 3, 8},
		{2, 4, 2},
		{3, 4, 7},
		{4, 3, 9},
	})
	fmt.Println("用例1: 手写图，起点 0 到各顶点的最短路径")
	pass = check("用例1: 手写图", g1, 0, []int{0, 7, 3, 9, 5}) && pass

	// 用例1续：换一个起点，顶点 0 没有入边，应保持不可达
	pass = check("用例1: 手写图", g1, 2, []int{math.MaxInt, 4, 0, 6, 2}) && pass

	// 用例2：链式图，顶点 3 不可达
	g2 := digraph.NewGraph(4)
	addEdges(g2, [][3]int{
		{0, 1, 5},
		{1, 2, 5},
	})
	pass = check("用例2: 存在不可达顶点", g2, 0, []int{0, 5, 10, math.MaxInt}) && pass

	// 用例3：只有单个顶点，起点到自身的路径长度为 0
	g3 := digraph.NewGraph(1)
	pass = check("用例3: 单顶点", g3, 0, []int{0}) && pass

	// 用例4：固定种子的随机图，与 Bellman-Ford 对照
	r := rand.New(rand.NewSource(42))
	for t := 0; t < 20; t++ {
		n := 5 + r.Intn(16)
		g := digraph.NewGraph(n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j && r.Intn(3) == 0 { // 约 1/3 概率存在边
					g.AddEdge(i, j, 1+r.Intn(20)) // 权值有重复，覆盖失效项场景
				}
			}
		}
		start := r.Intn(n)
		name := fmt.Sprintf("用例4: 随机图#%d (顶点=%d, 边=%d)", t, n, g.EdgeNum())
		pass = check(name, g, start, bellmanFord(g, start)) && pass
	}

	if pass {
		fmt.Println("\n所有用例通过，Dijkstra 正确")
	} else {
		fmt.Println("\n存在失败用例，Dijkstra 有误")
	}
}
