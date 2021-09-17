package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2,1,3,5,6,4}
	res := findPeakElement(nums)
	fmt.Println(res)
}

func findPeakElement(nums []int) int {
	if len(nums) == 1{
		return 0
	}

	for index,num := range nums{
		if index == 0 || index == len(nums)-1{
			continue
		}

		if num > nums[index-1] && num > nums[index+1]{
			return index
		}
	}
	return maxIndex(nums[0],nums[len(nums)-1],len(nums)-1)
}

func maxIndex(a,b,lastIndex int) int {
	if a>b{
		return 0
	}
	return lastIndex
}