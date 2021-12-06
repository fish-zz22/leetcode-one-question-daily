package main

import (
	"fmt"
)

//-2
//-2
//2
//2
//3
//3
//4
//4

func main() {
	ax1 := -2
	ay1 := -2
	ax2 := 2
	ay2 := 2
	bx1 := 2
	by1 := 0
	bx2 := 3
	by2 := 1
	computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	var conflictArea int
	// 判断两矩形是否重合
	if bx1 > ax2 || bx2 < ax1 || by1 > ay2 || by2 < ay1 {
		conflictArea = 0
	} else {
		// 算重叠的面积
		conflictArea = (min(ax2, bx2) - max(bx1, ax1)) * (min(by2, ay2) - max(by1, ay1))
	}
	// 求出两个矩形的面积不考虑重叠
	area := ((ax2 - ax1) * (ay2 - ay1)) + ((bx2 - bx1) * (by2 - by1))
	fmt.Println(area, conflictArea)
	return area - conflictArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
