package main

import "fmt"

func MergeSortTopDown(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	ret := make([]int, len(nums))
	mid := len(nums) / 2
	leftNums := MergeSortTopDown(nums[:mid])
	rightNums := MergeSortTopDown(nums[mid:])
	leftIndex, rightIndex := 0, 0
	for i := 0; i < len(ret); i++ {
		if leftIndex > len(leftNums)-1 {
			ret[i] = rightNums[rightIndex]
			rightIndex++
			continue
		}
		if rightIndex > len(rightNums)-1 {
			ret[i] = leftNums[leftIndex]
			leftIndex++
			continue
		}
		if leftNums[leftIndex] < rightNums[rightIndex] {
			ret[i] = leftNums[leftIndex]
			leftIndex++
		} else {
			ret[i] = rightNums[rightIndex]
			rightIndex++
		}
	}
	return ret
}

func MinNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MergeSortDownTop(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for windowsize := 1; windowsize <= (len(nums)+1)/2; windowsize *= 2 {
		tmp := make([]int, len(nums))
		leftIndex, leftEnd, rightIndex, rightEnd := 0, 0, 0, 0

		for t := 0; t < len(tmp); t++ {
			if leftIndex == leftEnd && rightIndex == rightEnd {
				leftIndex = rightEnd
				leftEnd = leftIndex + windowsize
				rightIndex = leftEnd
				rightEnd = MinNum(rightIndex+windowsize, len(nums))
			}
			if leftIndex == leftEnd {
				tmp[t] = nums[rightIndex]
				rightIndex++
				continue
			}
			if rightIndex == rightEnd {
				tmp[t] = nums[leftIndex]
				leftIndex++
				continue
			}
			if nums[leftIndex] < nums[rightIndex] {
				tmp[t] = nums[leftIndex]
				leftIndex++
			} else {
				tmp[t] = nums[rightIndex]
				rightIndex++
			}
		}

		nums = tmp
	}

	return nums
}

func main() {
	nums := []int{3, 1, 9, 4, 5, 3, 7, 8}
	MergeSortTopDown(nums)
	fmt.Println("TopDown: ", MergeSortTopDown(nums))
	fmt.Println("DownTop: ", MergeSortDownTop(nums))
}
