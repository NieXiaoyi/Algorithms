package main

import (
	"fmt"
	"math"
)

func CountingSort(nums []int) []int {
	minNum := math.MaxInt
	maxNum := math.MinInt
	for _, num := range nums {
		if minNum > num {
			minNum = num
		}
		if maxNum < num {
			maxNum = num
		}
	}

	numsCnt := make([]int, maxNum-minNum+1)
	for _, num := range nums {
		numsCnt[num-minNum]++
	}

	ret := make([]int, len(nums))
	retIndex := 0
	for i := 0; i < len(numsCnt); i++ {
		for numsCnt[i] > 0 {
			ret[retIndex] = i + minNum
			numsCnt[i]--
			retIndex++
		}
	}

	return ret
}

func main() {
	nums := []int{9, 1, 3, 4, 3, 11, 23, 88, 3, 45, 4, 6, 8, 9}
	fmt.Println(CountingSort(nums))
}
