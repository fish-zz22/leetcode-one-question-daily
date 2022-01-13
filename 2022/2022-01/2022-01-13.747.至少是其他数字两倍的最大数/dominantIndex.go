package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(sort.SearchInts(nums, 0))
	fmt.Println(dominantIndex(nums))
}

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	newNums := make([]int, len(nums))
	copy(newNums, nums)
	sort.Ints(newNums)
	fmt.Println(nums)

	if newNums[len(nums)-1] >= 2*newNums[len(nums)-2] {
		for i, v := range nums {
			if v == newNums[len(nums)-1] {
				return i
			}
		}
	}
	return -1
}
