package main

import (
	"fmt"
)

// 给一个升序数组[1, 2, 2, 2, 5, 7]，请找出位于最右侧，小于等于3的最大数的位置，不存在则返回-1
func MaxNum(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	ret := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			ret = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ret
}

func main() {
	fmt.Println("the index is ", MaxNum([]int{1, 2, 2, 2, 5, 7}, 4))
	fmt.Println("the index is ", MaxNum([]int{1, 2, 2, 2, 5, 7}, 0))
}
