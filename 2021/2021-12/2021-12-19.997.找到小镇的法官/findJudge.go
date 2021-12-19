package main

import "fmt"

func main() {
	trust := [][]int{
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{4, 3},
	}
	fmt.Println(findJudge(4, trust))
}

func findJudge(n int, trust [][]int) int {
	// 初始化2个数组
	// 1.每个人相信别人的次数
	// 2.每个人被别人相信的次数

	if n == 1 {
		return n
	}

	a1 := make([]int, n+1)
	a2 := make([]int, n+1)

	for _, v := range trust {
		a1[v[0]]++
		a2[v[1]]++
	}

	for i := 0; i <= n; i++ {
		if a1[i] == 0 && a2[i] == n-1 {
			return i
		}
	}

	return -1
}
