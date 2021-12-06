package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getSqrt(8))
	fmt.Println(minSteps(8))
}

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i

		for j := 2; j <= getSqrt(n); j++ {
			if i%j == 0 {
				dp[i] = dp[j] + dp[i/j]
				break
			}
		}
	}
	return dp[n]
}

func getSqrt(num int) int {
	return int(math.Ceil(math.Sqrt(float64(num))))
}
