package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 4, 5, 1}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {

	// 时间复杂度O(n)
	//var res int
	//temp := arr[0]
	//for i, v := range arr {
	//	if v > temp {
	//		temp = v
	//		res = i
	//	}else if v < temp{
	//		return res
	//	}
	//}
	//return res

	// 时间复杂度O(log(n)) 二分查找
	//l := len(arr)
	//left, right := 0, l-1
	//for left < right {
	//	// 防止溢出
	//	mid := left + (right-left)/2
	//	if arr[mid] < arr[mid+1] {
	//		left = mid + 1
	//	} else {
	//		right = mid
	//	}
	//}
	//return left

	// 调库版本二分
	return sort.Search(len(arr)-1, func(i int) bool {
		return arr[i] > arr[i+1]
	})
}
