package main

import (
	"fmt"
	"sort"
)

type PQ []int

func InitPQ() *PQ {
	var pq PQ
	pq = append(pq, 0)
	return &pq
}

func (pq *PQ) Insert(num int) {
	*pq = append(*pq, num)
	end := len(*pq) - 1
	i := end
	for {
		j := pq.swim(i)
		if j == i {
			break
		}
		i = j
	}
}

func (pq *PQ) DelMax() {
	end := len(*pq) - 1
	(*pq)[1] = (*pq)[end]
	*pq = (*pq)[:end]
	i := 1
	for {
		j := pq.sink(i)
		if j == i {
			break
		}
		i = j
	}
}

func (pq *PQ) Max() (int, bool) {
	if len(*pq) > 1 {
		return (*pq)[1], true
	}
	return 0, false
}

func (pq *PQ) swim(i int) int {
	if i <= 1 {
		return i
	}

	parent := i / 2
	if (*pq)[i] > (*pq)[parent] {
		pq.exchange(i, parent)
		return parent
	}
	return i
}

func (pq *PQ) sink(i int) int {
	lchild := i * 2
	rchild := i*2 + 1
	if lchild > len(*pq)-1 {
		return i
	}

	var maxchild int
	if rchild > len(*pq)-1 {
		maxchild = lchild
	} else if (*pq)[lchild] > (*pq)[rchild] {
		maxchild = lchild
	} else {
		maxchild = rchild
	}
	if (*pq)[maxchild] > (*pq)[i] {
		pq.exchange(maxchild, i)
		return maxchild
	}
	return i
}

func (pq *PQ) exchange(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func main() {
	nums := []int{3, 2, 4, 9, 7, 1, 8, 0, 1}
	ret := make([]int, len(nums))
	copy(ret, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(ret)))

	pq := InitPQ()
	for _, num := range nums {
		pq.Insert(num)
	}

	for _, v := range ret {
		m, ok := pq.Max()
		if !ok {
			fmt.Printf("error: pq is empty\n")
		}
		if v != m {
			fmt.Printf("error: v is %d, but pq.Max is %d\n", v, m)
		} else {
			fmt.Printf("correct: v is %d, pq.Max is %d\n", v, m)
		}
		pq.DelMax()
	}
}
