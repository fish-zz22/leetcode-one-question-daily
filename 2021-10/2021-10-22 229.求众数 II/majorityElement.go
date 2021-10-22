package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 2, 2, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) []int {
	//var res []int
	//countMap := make(map[int]int, 0)
	//for _, v := range nums {
	//	if _, ok := countMap[v]; ok {
	//		countMap[v]++
	//	} else {
	//		countMap[v] = 1
	//	}
	//}
	//
	//for v, count := range countMap {
	//	if count > len(nums)/3 {
	//		res = append(res, v)
	//	}
	//}
	//return res

	var res []int
	select1, select2, vote1, vote2 := 0, 0, 0, 0
	for _, v := range nums {
		if vote1 > 0 && select1 == v {
			vote1++
		} else if vote2 > 0 && select2 == v {
			vote2++
		} else if vote1 == 0 {
			select1 = v
			vote1++
		} else if vote2 == 0 {
			select2 = v
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	count1, count2 := 0, 0
	for _, v := range nums {
		if vote1 > 0 && select1 == v {
			count1++
		}
		if vote2 > 0 && select2 == v {
			count2++
		}
	}

	if vote1 > 0 && count1 > len(nums)/3 {
		res = append(res, select1)
	}
	if vote2 > 0 && count2 > len(nums)/3 {
		res = append(res, select2)
	}
	return res
}
