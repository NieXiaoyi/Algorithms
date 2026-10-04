package main

import (
	"fmt"

	undigraph "com.algorithms/graph/1_undigraph"
)

type Component struct {
	id []int
}

func NewComponent(g *undigraph.Graph) *Component {
	comp := &Component{
		make([]int, g.VertexNum()),
	}
	visited := make([]bool, g.VertexNum())

	var dfs func(g *undigraph.Graph, vertex int, id int)
	dfs = func(g *undigraph.Graph, vertex int, id int) {
		if visited[vertex] {
			return
		}

		visited[vertex] = true
		comp.id[vertex] = id
		for _, next := range g.AdjVertex(vertex) {
			dfs(g, next, id)
		}
	}

	id := 0
	for vertex := 0; vertex < g.VertexNum(); vertex++ {
		if visited[vertex] {
			continue
		}
		dfs(g, vertex, id)
		id++
	}
	return comp
}

func (c *Component) ID(vertex int) int {
	return c.id[vertex]
}

func main() {
	g := undigraph.NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)

	comp := NewComponent(g)
	for i := 0; i < 5; i++ {
		fmt.Println("vertex: ", i, " comp-id: ", comp.ID(i))
	}

}
