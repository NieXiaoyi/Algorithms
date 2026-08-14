package main

import "fmt"

func exchange(nums []int, ia int, ib int) {
	tmp := nums[ia]
	nums[ia] = nums[ib]
	nums[ib] = tmp
}

func partition(nums []int, head int, end int) int {
	left := head + 1
	right := end
	for {
		for left <= end && nums[left] < nums[head] {
			left++
		}
		for right > head && nums[right] > nums[head] {
			right--
		}
		if left < right {
			exchange(nums, left, right)
		} else {
			break
		}
	}

	exchange(nums, head, right)
	return right
}

func sort(nums []int, head int, end int) {
	if head >= end {
		return
	}
	index := partition(nums, head, end)
	sort(nums, head, index-1)
	sort(nums, index+1, end)
}

func main() {
	nums := []int{3, 1, 9, 4, 5, 3, 7}
	sort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
