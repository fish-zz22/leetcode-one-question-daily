package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{16, 17, 18}
	fmt.Println(numFriendRequests(ages))
}

func numFriendRequests(ages []int) (res int) {
	//todo 超时
	//for i := 0; i < len(ages); i++ {
	//	for j := 0; j < len(ages); j++ {
	//		// 不能给自己发好友请求
	//		if i == j {
	//			continue
	//		}
	//
	//		if ages[j] <= int(0.5*float64(ages[i]))+7 || ages[j] > ages[i] || (ages[j] > 100 && ages[i] < 100) {
	//			continue
	//		} else {
	//			res += 1
	//		}
	//	}
	//}
	//return
	sort.Ints(ages)
	left, right := 0, 0
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		res += right - left
	}
	return
}
