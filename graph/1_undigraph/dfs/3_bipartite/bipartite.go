package main

import (
	"fmt"
	"math/rand"

	undigraph "com.algorithms/graph/1_undigraph"
)

// 能够用两种颜色将图的所有顶点着色，使得任意一条边的两个端点的颜色都不相同吗
func Bipartite(g *undigraph.Graph) bool {
	visited := make([]bool, g.VertexNum())
	color := make([]bool, g.VertexNum())

	var dfs func(g *undigraph.Graph, start int) bool
	dfs = func(g *undigraph.Graph, start int) bool {
		visited[start] = true // 标记已访问,否则相邻顶点互相递归导致栈溢出
		for _, next := range g.AdjVertex(start) {
			if !visited[next] {
				color[next] = !color[start] // 相邻顶点染相反的颜色
				if !dfs(g, next) {          // 必须传播递归结果,否则深处的冲突会被吞掉
					return false
				}
			} else if color[start] == color[next] {
				return false
			}
		}
		return true
	}

	for i := 0; i < g.VertexNum(); i++ {
		if visited[i] {
			continue
		}
		if !dfs(g, i) {
			return false
		}
	}
	return true
}

// bruteForce 枚举所有 2^n 种二着色方案,判断是否存在合法着色(与实现无关的独立参照)
func bruteForce(n int, edges [][2]int) bool {
	color := make([]int, n)
	var check func(i int) bool
	check = func(i int) bool {
		if i == n {
			for _, e := range edges {
				if color[e[0]] == color[e[1]] {
					return false
				}
			}
			return true
		}
		for c := 0; c < 2; c++ {
			color[i] = c
			if check(i + 1) {
				return true
			}
		}
		return false
	}
	return check(0)
}

func main() {
	tests := []struct {
		name  string
		n     int
		edges [][2]int
		want  bool
	}{
		{"空图", 0, nil, true},
		{"无边图", 5, nil, true},
		{"单条边", 2, [][2]int{{0, 1}}, true},
		{"树(必为二分图)", 6, [][2]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}}, true},
		{"偶环(4顶点)", 4, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, true},
		{"奇环(三角形)", 3, [][2]int{{0, 1}, {1, 2}, {2, 0}}, false},
		{"奇环(五边形)", 5, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 0}}, false},
		{"自环", 1, [][2]int{{0, 0}}, false},
		{"非连通-含奇环分量", 5, [][2]int{{0, 1}, {1, 2}, {2, 0}}, false},
		{"非连通-偶环+奇环", 8, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {4, 5}, {5, 6}, {6, 4}}, false},
		{"非连通-两个偶环", 8, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {4, 5}, {5, 6}, {6, 7}, {7, 4}}, true},
	}

	for _, tt := range tests {
		g := undigraph.NewGraph(tt.n)
		for _, e := range tt.edges {
			g.AddEdge(e[0], e[1])
		}
		got := Bipartite(g)
		status := "PASS"
		if got != tt.want {
			status = "FAIL"
		}
		fmt.Printf("[%s] %s: got=%v, want=%v\n", status, tt.name, got, tt.want)
	}

	// 随机小图与暴力枚举交叉验证(邻接表为 map,遍历顺序随机,可覆盖不同 DFS 顺序)
	rnd := rand.New(rand.NewSource(42))
	fail := 0
	for i := 0; i < 500; i++ {
		n := rnd.Intn(8)
		g := undigraph.NewGraph(n)
		edges := [][2]int{}
		m := rnd.Intn(n*(n-1)/2 + n + 1) // 上限含自环
		for j := 0; j < m; j++ {
			e := [2]int{rnd.Intn(n), rnd.Intn(n)}
			edges = append(edges, e)
			g.AddEdge(e[0], e[1])
		}
		if got, want := Bipartite(g), bruteForce(n, edges); got != want {
			fail++
			fmt.Printf("[FAIL] 随机图 n=%d, edges=%v: got=%v, want=%v\n", n, edges, got, want)
		}
	}
	if fail == 0 {
		fmt.Println("[PASS] 随机交叉验证: 500 组随机图与暴力枚举结果一致")
	} else {
		fmt.Printf("[FAIL] 随机交叉验证: %d 组不一致\n", fail)
	}
}
