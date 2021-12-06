package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int, 0)
	var stack, res []int
	for _, num := range nums2 {
		for len(stack) != 0 && stack[len(stack)-1] < num {
			hash[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	for _, num := range nums1 {
		if hash[num] == 0 {
			res = append(res, -1)
		} else {
			res = append(res, hash[num])
		}
	}
	return res
}
