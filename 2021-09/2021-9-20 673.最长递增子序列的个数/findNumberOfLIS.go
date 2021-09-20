package main

import "fmt"

func main() {
	nums := []int{1,2,4,3,5,4,7,2}
	findNumberOfLIS(nums)
}

func findNumberOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int,l)
	count := make([]int,l)
	maxLength,res := 0,0

	for i,value1 := range nums{
		dp[i] = 1
		count[i] = 1
		for j,value2 := range nums[:i]{

			// count[i]是维护最长递增子序列的个数
			// 如果value1 > value2  dp[i] = max(dp[j]+1,dp[i])
			// 如果递增子序列的长度增加 需要重新赋值count[i]
			// 如果dp[i] = dp[j] + 1  count[i]+1
			if value1 > value2{
				if dp[j]+1 > dp[i]{
					dp[i] = dp[j]+1
					count[i] = count[j]
				}else if dp[j]+1 == dp[i]{
					count[i] += count[j]
				}
			}
		}

		// 循环找到最长递增子序列的长度
		if maxLength < dp[i]{
			maxLength = dp[i]
			res = count[i]
		}else if dp[i] == maxLength{
			res += count[i]
		}
	}
	fmt.Println(dp,count)
	return res
}
