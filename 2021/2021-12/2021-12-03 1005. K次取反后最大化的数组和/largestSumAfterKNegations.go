package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, -3, -1, 5, -4}
	fmt.Println(largestSumAfterKNegations(nums, 4))
}

func largestSumAfterKNegations(nums []int, k int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			k--
			nums[i] = -nums[i]
		}
		sum += nums[i]
	}
	sort.Ints(nums)
	last := 0
	if k%2 != 0 {
		last = 2 * nums[0]
	}
	return sum - last
}
