package weighteddigraph

type Graph struct {
	// adjMap[s][d]的值为顶点s到顶点d的权重
	adjMap  []map[int]int
	edgeNum int
}

// 创建一个有n个顶点的空图
func NewGraph(n int) *Graph {
	adjMap := make([]map[int]int, n)
	for i := range adjMap {
		adjMap[i] = make(map[int]int)
	}

	return &Graph{
		adjMap:  adjMap,
		edgeNum: 0,
	}
}

func (g *Graph) VertexNum() int {
	return len(g.adjMap)
}

func (g *Graph) EdgeNum() int {
	return g.edgeNum
}

func (g *Graph) AddEdge(s, d, weight int) {
	if _, ok := g.adjMap[s][d]; !ok {
		g.adjMap[s][d] = weight
		g.edgeNum++
	}
}

// 计算顶点s的出点
func (g *Graph) Adj(s int) []int {
	adj := make([]int, 0)
	for d := range g.adjMap[s] {
		adj = append(adj, d)
	}
	return adj
}

// 顶点s到顶点d的权重，两顶点间没有边时返回-1
func (g *Graph) Weight(s, d int) int {
	w, ok := g.adjMap[s][d]
	if !ok {
		return -1
	}
	return w
}

// 加权有向图的逆向图，各边的权重保持不变
func (g *Graph) Reverse() *Graph {
	rg := NewGraph(g.VertexNum())

	for s := 0; s < g.VertexNum(); s++ {
		for _, d := range g.Adj(s) {
			rg.AddEdge(d, s, g.Weight(s, d))
		}
	}
	return rg
}
