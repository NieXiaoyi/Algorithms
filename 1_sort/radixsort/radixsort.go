package main

import (
	"fmt"
	"math"
)

func RadixSort(nums []int) []int {
	var BitCountSort func(nums []int, divisor int) []int
	BitCountSort = func(nums []int, divisor int) []int {
		bitsCnt := make([]int, 10)
		for _, num := range nums {
			bit := num / divisor % 10
			bitsCnt[bit]++
		}

		// 计算nums排序后，该bit位在排序数组中的最大位置
		for i := 1; i < len(bitsCnt); i++ {
			bitsCnt[i] += bitsCnt[i-1]
		}

		// 第一轮排序后，第0位已经有序
		// 因为对于同一值的bit，ret是先写大位置，因此我们需要倒序遍历nums，这样第二轮排序时，第0位的顺序才不会被搞乱，以此类推后续轮次的排序
		// 示例： nums = [59, 57, 58, 56, 18]
		// 第一轮排序： nums = [56, 57, 58, 18, 59]
		// 第二轮排序： nums = [18, 56, 57, 58, 59]
		ret := make([]int, len(nums))
		for i := len(nums) - 1; i >= 0; i-- {
			bit := nums[i] / divisor % 10
			pos := bitsCnt[bit] - 1
			ret[pos] = nums[i]
			bitsCnt[bit]--
		}
		return ret
	}

	maxNum := math.MinInt
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	for divisor := 1; maxNum >= divisor; divisor *= 10 {
		nums = BitCountSort(nums, divisor)
	}
	return nums
}

func main() {
	nums := []int{16, 777, 216, 16777215, 1, 0, 7, 44, 3636, 4523, 123431}
	fmt.Println(RadixSort(nums))
}
