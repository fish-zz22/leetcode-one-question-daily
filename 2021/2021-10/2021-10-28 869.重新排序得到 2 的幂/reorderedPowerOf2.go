package main

import "fmt"

func main() {
	n := 268341
	fmt.Println(reorderedPowerOf2(n))
}

func reorderedPowerOf2(n int) bool {
	// todo:先全排列出所有可能的数,然后验证是否为2的幂
	// 全排列+判断
	var numList []int
	for n > 0 {
		numList = append(numList, n%10)
		n /= 10
	}
	used := make([]bool, len(numList), len(numList))

	var dfs func(num int, n int) bool
	dfs = func(num int, n int) bool {
		if n == len(numList) {
			return isPowerOf2(num)
		}
		for i := 0; i < len(numList); i++ {
			if n == 0 && numList[i] == 0 {
				continue
			} else if !used[i] {
				used[i] = true
				if dfs(10*num+numList[i], n+1) {
					return true
				}
				used[i] = false
			}
		}
		return false
	}
	return dfs(0, 0)
}

func isPowerOf2(n int) bool {
	return n > 0 && n&(n-1) == 0
}
