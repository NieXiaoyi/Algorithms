package main

import "fmt"

// nums为一个单调递增且无重复元素的数组，请从中找到导致该数组不连续的缺失整数， 未找到则返回-1
func missinteger(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := -1
	for (left <= right) && (nums[right]-nums[left] > right-left) {
		mid = (left + right) / 2

		if nums[mid]-nums[left] > mid-left {
			right = mid - 1
		} else if nums[right]-nums[mid] > right-mid {
			left = mid + 1
		}
	}

	if (mid > 0) && (nums[mid]-nums[mid-1] > 1) {
		return nums[mid] - 1
	}

	if (mid < len(nums)-1) && (nums[mid+1]-nums[mid] > 1) {
		return nums[mid] + 1
	}

	return -1
}

func main() {
	nums1 := []int{0, 1, 2, 3, 4, 5, 6, 8}
	fmt.Println(missinteger(nums1))
	nums2 := []int{0, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(missinteger(nums2))
}
