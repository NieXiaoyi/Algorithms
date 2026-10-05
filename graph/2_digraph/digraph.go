package digraph

type Graph struct {
	// 顶点的出点集合
	adjMap  []map[int]struct{}
	edgeNum int
}

func NewGraph(n int) *Graph {
	adjMap := make([]map[int]struct{}, n)
	for i := range adjMap {
		adjMap[i] = make(map[int]struct{})
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

func (g *Graph) AddEdge(s, d int) {
	if _, ok := g.adjMap[s][d]; !ok {
		g.adjMap[s][d] = struct{}{}
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

// 有向图的逆向图
func (g *Graph) Reverse() *Graph {
	rg := NewGraph(g.VertexNum())

	for s := 0; s < g.VertexNum(); s++ {
		for d := range g.Adj(s) {
			rg.AddEdge(d, s)
		}
	}
	return rg
}
