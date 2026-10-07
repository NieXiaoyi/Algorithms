package undigraph

type Graph struct {
	// adjMap[i][j]的值为顶点i到顶点j的权重
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

func (g *Graph) AddEdge(i, j, weight int) {
	if _, ok := g.adjMap[i][j]; !ok {
		g.adjMap[i][j] = weight
		g.adjMap[j][i] = weight
		g.edgeNum++
	}
}

func (g *Graph) AdjVertex(i int) []int {
	ret := []int{}
	for v := range g.adjMap[i] {
		ret = append(ret, v)
	}
	return ret
}

func (g *Graph) Weight(i, j int) int {
	w, ok := g.adjMap[i][j]
	if !ok {
		return -1
	}
	return w
}
