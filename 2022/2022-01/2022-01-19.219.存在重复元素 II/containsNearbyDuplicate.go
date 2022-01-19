package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int)

	for i, v := range nums {
		if temp, ok := hash[v]; ok && i-temp <= k {
			return true
		}
		hash[v] = i
	}
	return false
}
