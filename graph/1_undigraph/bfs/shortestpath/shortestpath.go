package main

import (
	"fmt"

	undigraph "com.algorithms/graph/1_undigraph"
)

type Queue []int

func (q *Queue) Enqueue(n int) {
	*q = append(*q, n)
}

func (q *Queue) Dequeue() int {
	defer func() {
		*q = (*q)[1:]
	}()

	return (*q)[0]
}

func (q *Queue) Size() int {
	return len(*q)
}

func ShortestPath(g *undigraph.Graph, beg, end int) []int {
	q := &Queue{}
	q.Enqueue(beg)
	visited := make([]bool, g.VertexNum())
	visited[beg] = true
	parentNode := make([]int, g.VertexNum())
	pathLen := 0

	for q.Size() > 0 {
		pathLen++
		size := q.Size()
		for i := 0; i < size; i++ {
			cur := q.Dequeue()
			nextNodes := g.AdjVertex(cur)
			for _, node := range nextNodes {
				if visited[node] {
					continue
				}
				visited[node] = true
				parentNode[node] = cur
				if node == end {
					pathLen++
					goto makeret
				}
				q.Enqueue(node)
			}
		}
	}
	return []int{}

makeret:
	ret := make([]int, pathLen)
	i := pathLen - 1
	for cur := end; cur != beg; cur = parentNode[cur] {
		ret[i] = cur
		i--
	}
	ret[0] = beg
	return ret
}

func main() {
	g := undigraph.NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)

	fmt.Println(ShortestPath(g, 0, 4))
}
