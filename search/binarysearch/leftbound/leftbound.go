package main

import (
	"fmt"
)

// 给一个升序数组[1, 2, 2, 2, 5, 7]，请找出最左侧的2的位置，找不到则返回 -1
func LeftBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	ret := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			ret = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ret
}

func main() {
	nums := []int{1, 2, 2, 2, 5, 7}
	target := 2
	fmt.Println("the index is ", LeftBound(nums, target))
}
