package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(minMoves(nums))
}

func minMoves(nums []int) int {
	var res int
	sort.Ints(nums)
	min, max := nums[0], nums[len(nums)-1]
	if min != max {
		for _, v := range nums {
			res += v - min
		}
		return res
	}
	return 0
}
