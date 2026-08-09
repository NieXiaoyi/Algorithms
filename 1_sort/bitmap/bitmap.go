package main

import "fmt"

const MAX_NUM = 1<<24 - 1

// 对一个大数组进行排序，已知数组内最大数不超过 2^24-1，最小数不小于0，数组内无重复数
func BitMapSort(nums []int) []int {
	// 构建位图，每个位的位置对应一个数，如位置2对应数2
	blocksize := 64
	bitvector := make([]int64, (MAX_NUM+1)/blocksize)

	for _, num := range nums {
		blockindex := num / blocksize
		bitvector[blockindex] = bitvector[blockindex] | (int64(1) << int64(num%blocksize))
	}

	ret := make([]int, len(nums))
	retindex := 0
	for blockindex := range bitvector {
		for i := range blocksize {
			if bitvector[blockindex]&(int64(1)<<int64(i)) != 0 {
				ret[retindex] = blocksize*blockindex + i
				retindex++
			}
		}
	}
	return ret
}

func main() {
	nums := []int{16, 777, 216, 16777215, 1, 0, 7, 44, 3636, 4523, 123431}
	fmt.Println(BitMapSort(nums))
}
