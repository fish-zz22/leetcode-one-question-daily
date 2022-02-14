package main

import "fmt"

func main() {
	nums := []int{1, 3, 3, 7, 7, 11, 11}
	fmt.Println(singleNonDuplicate(nums))
}

func singleNonDuplicate(nums []int) int {
	i, j := 0, 1
	for k := 0; k < len(nums)/2; k++ {
		if nums[i] == nums[j] {
			i += 2
			j += 2
		} else {
			break
		}
	}
	return nums[i]
}
