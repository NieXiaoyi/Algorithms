package undigraph

type Graph struct {
	adjMap  []map[int]struct{}
	edgeNum int
}

// 创建一个有n个顶点的空图
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

func (g *Graph) AddEdge(i, j int) {
	if _, ok := g.adjMap[i][j]; !ok {
		g.adjMap[i][j] = struct{}{}
		g.adjMap[j][i] = struct{}{}
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
